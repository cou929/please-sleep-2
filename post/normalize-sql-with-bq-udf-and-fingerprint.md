{"title":"BigQuery の UDF で MySQL のクエリを正規化する","date":"2022-05-27T00:40:00+09:00","tags":["sql"]}

BigQuery にクエリのログ、例えばスローログなどが貯めてあり、それを分析したいとする。似たようなクエリを集約し、それぞれのクエリの実行時間や走査した行数の avg, min, max などを見たい。「似たようなクエリを集約」するためにはクエリを正規化、つまり引数の違いを無くしたりスペースや改行を統一したりする必要がある。

そこでクエリを正規化する BigQuery の UDF を実装し、正規化したクエリごとに GROUP BY できるようにした。またせっかくなので npm のモジュールにまとめた。

[sql\-fingerprint \- npm](https://www.npmjs.com/package/sql-fingerprint)

なお今回は MySQL のクエリを対象にしている。

## pt-fingerprint

一例として [Percona Toolkit](https://www.percona.com/software/database-tools/percona-toolkit) には [pt-fingerprint](https://www.percona.com/doc/percona-toolkit/LATEST/pt-fingerprint.html) というツールがある。例えば次のようにクエリの正規化をしてくれる。正規化された文字列は fingerprint と呼ばれている。実装はごりごりと [perl と正規表現](https://github.com/percona/percona-toolkit/blob/d881738e2fb1c2936d714604c57db21b18d977d1/bin/pt-fingerprint#L1636-L1704) で行われている。

```sql
SELECT name, password FROM user WHERE id='12823';
select name,   password from user
   where id=5;

-- 上記 2 クエリはどちらも以下の文字列に変換される。
select name, password from user where id=?
```

スローログや general log をそのまま解析してくれるツールとしては [pt-query-digest](https://www.percona.com/doc/percona-toolkit/LATEST/pt-query-digest.html) もある。ただ今回は BigQuery 上のデータを見たい用途だったので、fingerprint の計算だけをしてくれる pt-fingerprint のほうがフィットした。なお pt-query-digest の fingerprint 計算も [同じロジック](https://github.com/percona/percona-toolkit/blob/d881738e2fb1c2936d714604c57db21b18d977d1/bin/pt-query-digest#L2928-L2996) で行われている。

この他には、例えば [TiDB](https://en.pingcap.com/tidb/) のパーサ [pingcap/tidb/parser](https://github.com/pingcap/tidb/tree/master/parser) を使うとクエリをパースし AST を得ることができるので、これを使って自分で正規化処理を実装するアプローチも考えられる。正規表現の実装よりはメンテナンスしやすいしパースも頑強になりそうだけど、細かいエッジケースなどに対応するのは大変なので、そこは pt-fingerprint に一日の長があると思う。

また RDS の場合は [Performance Insights](https://aws.amazon.com/rds/performance-insights/) や、自分は利用したことがないが VividCortex, Monyog, PMM といった [モニタリングツール](https://www.percona.com/blog/2017/03/16/monitoring-databases-a-product-comparison/) は、スロークエリの分析機能などを提供している。今回はたまたま BigQuery にあるクエリの分析をしたいという用途だったのでフィットしなかったが、MySQL の管理全般をしたい場合にはこうしたツールもある。

## sql-fingerprint

今回の目的は pt-fingerprint でも満たせるが、いちいち BigQuery からクエリをエクスポートし、手元などでそれに fingerprint をかけ、その結果をまた BigQuery に戻すなどが必要で、手間がかかる。BigQuery 上だけで正規化できたほうが簡単。

そこで BigQuery の UDF は [JavaScript でロジックを書ける](https://cloud.google.com/bigquery/docs/reference/standard-sql/user-defined-functions#javascript-udf-structure) のでこれを利用することにした。pt-fingerprint のロジックを JavaScript に移植した。

次のような感じで BigQuery のクエリだけで正規化と集計ができる。

```sql
CREATE TEMP FUNCTION fingerprint(sql STRING, matchMD5Checksum BOOL, matchEmbeddedNumbers BOOL)
RETURNS STRING
LANGUAGE js AS r"""
function fingerprint(sql, matchMD5Checksum, matchEmbeddedNumbers) {
  let query = sql;

  // special cases
  if (/^SELECT \/\*!40001 SQL_NO_CACHE \*\/ \* FROM `/.test(query)) {
    return 'mysqldump';
  }
  if (/\/\*\w+\.\w+:[0-9]\/[0-9]\*\//.test(query)) {
    return 'percona-toolkit';
  }
  if (/^administrator command: /.test(query)) {
    return query;
  }
  const matchedCallStatement = query.match(/^\s*(call\s+\S+)\(/i);
  if (matchedCallStatement) {
    return matchedCallStatement[1].toLowerCase();
  }

  // shorten multi-value INSERT statement
  const matchedMultiValueInsert = query.match(/^((?:INSERT|REPLACE)(?: IGNORE)?\s+INTO.+?VALUES\s*\(.*?\))\s*,\s*\(/is);
  if (matchedMultiValueInsert) {
    // eslint-disable-next-line prefer-destructuring
    query = matchedMultiValueInsert[1];
  }

  // multi line comment
  query = query.replace(/\/\*[^!].*?\*\//g, '');

  // one_line_comment
  query = query.replace(/(?:--|#)[^'"\r\n]*(?=[\r\n]|$)/g, '');

  // USE statement
  if (/^use \S+$/i.test(query)) {
    return 'use ?';
  }

  // literals
  query = query.replace(/([^\\])(\\')/sg, '$1');
  query = query.replace(/([^\\])(\\")/sg, '$1');
  query = query.replace(/\\\\/sg, '');
  query = query.replace(/\\'/sg, '');
  query = query.replace(/\\"/sg, '');
  query = query.replace(/([^\\])(".*?[^\\]?")/sg, '$1?');
  query = query.replace(/([^\\])('.*?[^\\]?')/sg, '$1?');

  query = query.replace(/\bfalse\b|\btrue\b/isg, '?');

  if (matchMD5Checksum) {
    query = query.replace(/([._-])[a-f0-9]{32}/g, '$1?');
  }

  if (!matchEmbeddedNumbers) {
    query = query.replace(/[0-9+-][0-9a-f.xb+-]*/g, '?');
  } else {
    query = query.replace(/\b[0-9+-][0-9a-f.xb+-]*/g, '?');
  }

  if (matchMD5Checksum) {
    query = query.replace(/[xb+-]\?/g, '?');
  } else {
    query = query.replace(/[xb.+-]\?/g, '?');
  }

  // collapse whitespace
  query = query.replace(/^\s+/, '');
  query = query.replace(/[\r\n]+$/, '');
  query = query.replace(/[ \n\t\r\f]+/g, ' ');

  // to lower case
  query = query.toLowerCase();

  // get rid of null
  query = query.replace(/\bnull\b/g, '?');

  // collapse IN and VALUES lists
  query = query.replace(/\b(in|values?)(?:[\s,]*\([\s?,]*\))+/g, '$1(?+)');

  // collapse UNION
  query = query.replace(/\b(select\s.*?)(?:(\sunion(?:\sall)?)\s\1)+/g, '$1 /*repeat$2*/');

  // limit
  query = query.replace(/\blimit \?(?:, ?\?| offset \?)?/, 'limit ?');

  // order by
  query = query.replace(/\b(.+?)\s+ASC/gi, '$1');

  return query;
}

return fingerprint(sql, true, true);
""";

SELECT
  fingerprint(query, true, true) fp,
  count(*) as num,
  max(query) as raw_query_sample,
  avg(exec_time) as avg_exec_time
FROM
  `your_table`
WHERE
  DATE(timestamp, "Asia/Tokyo") = "2022-05-22"
group by
  fp
order by
  num desc
```

UDF には [最大サイズなどの制限がある](https://cloud.google.com/bigquery/quotas#udf_limits) ので、そういう意味でもパースするアプローチよりもこの正規表現のアプローチのほうが適していたと思う。

実装したものはせっかくなので npm モジュールにまとめておいた。

- [cou929/sql\-fingerprint\-js: Converts a SQL into a fingerprint\. A JavaScript port of pt\-fingerprint\.](https://github.com/cou929/sql-fingerprint-js)
- [sql\-fingerprint \- npm](https://www.npmjs.com/package/sql-fingerprint)

次のようにモジュールとしての利用や、一応 CLI からも呼び出せるようになっている。

```js
// Module
import fingerprint from 'sql-fingerprint';

console.log(fingerprint('SELECT * FROM users WHERE id = 1', false, false));
```

```sh
# CLI

npm install -g sql-fingerprint

fingerprint --query="your query"
```

## PR

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0824F8ZZD/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/41zHk1tFyML.jpg" alt="Google Cloud Platform 実践 ビッグデータ分析基盤開発ストーリーで学ぶGoogle BigQuery" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0824F8ZZD/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Google Cloud Platform 実践 ビッグデータ分析基盤開発ストーリーで学ぶGoogle BigQuery</a></div><div class="amazlet-detail">株式会社トップゲート  (著)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0824F8ZZD/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
