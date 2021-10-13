{"title":"MySQL の Generated Columns のキャッチアップ","date":"2021-04-07T02:00:00+09:00","tags":["mysql"]}

個人的に Generated Columns を使ったことがなかったが、今のプロジェクトで使用していたのでキャッチアップするメモ。

## Generated Columns とは

- Generated Columns は他のカラムから計算した結果を値を持つカラム
- 通常のカラムの型定義のうしろに `[GENERATED ALWAYS] AS (式)` という書式で定義する

```sql
-- カラム b は a に 1 を加えた数を値として持つ
mysql> CREATE TABLE gctest(a INT, b INT GENERATED ALWAYS AS (a + 1));

-- a だけ INSERT すると b も SELECT できている
mysql> insert into gctest (a) values (1), (2), (3);
mysql> select * from gctest;
+------+------+
| a    | b    |
+------+------+
|    1 |    2 |
|    2 |    3 |
|    3 |    4 |
+------+------+
```

- insert, update, replace で直接操作できない
    - `DEFAULT` という値を変わりに指定することは可能

```sql
-- generated column を指定して INSERT はできない
mysql> insert into gctest (a, b) values (10, 11);
ERROR 3105 (HY000): The value specified for generated column 'b' in table 'gctest' is not allowed.

-- b = DEFAULT と指定すれば通る
mysql> insert into gctest (a, b) values (10, DEFAULT);
Query OK, 1 row affected (0.00 sec)

mysql> select * from gctest where a = 10;
+------+------+
| a    | b    |
+------+------+
|   10 |   11 |
+------+------+

-- update も同様
mysql> update gctest set b = 100 where a = 10;
ERROR 3105 (HY000): The value specified for generated column 'b' in table 'gctest' is not allowed.

mysql> update gctest set a = 100, b = DEFAULT where a = 10;
Query OK, 1 row affected (0.00 sec)
Rows matched: 1  Changed: 1  Warnings: 0

mysql> select * from gctest where a = 100;
+------+------+
| a    | b    |
+------+------+
|  100 |  101 |
+------+------+
```

- expr には非決定的な関数やサブクエリは使えないなどの制約がある

```sql
mysql> CREATE TABLE gctest2(a INT, b DATETIME GENERATED ALWAYS AS (NOW()));
ERROR 3763 (HY000): Expression of generated column 'b' contains a disallowed function: now.
```

- 式の結果とカラムの型が違う場合、mysql デフォルトの型変換がかかる

## 値の格納方式 (VIRTUAL / STORED)

- 計算結果の格納方式として VIRTUAL と STORED という 2 種類がある
    - その名の通り、VIRTUAL は実際には値を保持せず毎回計算、STORED は計算結果を実際にデータ領域に保存 (計算もとのカラムが更新されるたびに導出先も更新) する 
    - デフォルトは VIRTUAL で、1 テーブル内で両方式は混在できる
- いずれもインデックスを張ることができる
    - STORED の場合は通常のカラムと同様
    - VIRTUAL には Secondary Index のみ対応
        - virtual index とも呼ばれるらしい
        - 5.7.8 から対応されたらしい

```sql
mysql> CREATE TABLE gctest2(a INT, b INT GENERATED ALWAYS AS (a + 1) STORED PRIMARY KEY);
Query OK, 0 rows affected (0.01 sec)

mysql> CREATE TABLE gctest3(a INT, b INT GENERATED ALWAYS AS (a + 1) VIRTUAL UNIQUE);
Query OK, 0 rows affected (0.02 sec)

mysql> SHOW CREATE TABLE gctest2\G
*************************** 1. row ***************************
       Table: gctest2
Create Table: CREATE TABLE `gctest2` (
  `a` int DEFAULT NULL,
  `b` int GENERATED ALWAYS AS ((`a` + 1)) STORED NOT NULL,
  PRIMARY KEY (`b`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
1 row in set (0.00 sec)

mysql> SHOW CREATE TABLE gctest3\G
*************************** 1. row ***************************
       Table: gctest3
Create Table: CREATE TABLE `gctest3` (
  `a` int DEFAULT NULL,
  `b` int GENERATED ALWAYS AS ((`a` + 1)) VIRTUAL,
  UNIQUE KEY `b` (`b`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
1 row in set (0.00 sec)

-- VIRTUAL のカラムを PK にはできない
mysql> CREATE TABLE gctest4(a INT, b INT GENERATED ALWAYS AS (a + 1) VIRTUAL PRIMARY KEY);
ERROR 3106 (HY000): 'Defining a virtual generated column as primary key' is not supported for generated columns.
```

- VIRTUAL のカラムに Secondary Index を張るった場合、前述のようにデータ領域に計算された値は保存されないが、インデックスには保存される
    - なので Covering Index が効く

> When a secondary index is created on a virtual generated column, generated column values are materialized in the records of the index. If the index is a covering index (one that includes all the columns retrieved by a query), generated column values are retrieved from materialized values in the index structure instead of computed “on the fly”.
>
> https://dev.mysql.com/doc/refman/8.0/en/create-table-secondary-indexes.html より

```sql
-- カラム b は unique な Generated Columns 
mysql> SHOW CREATE TABLE gctest3\G
*************************** 1. row ***************************
       Table: gctest3
Create Table: CREATE TABLE `gctest3` (
  `a` int DEFAULT NULL,
  `b` int GENERATED ALWAYS AS ((`a` + 1)) VIRTUAL,
  UNIQUE KEY `b` (`b`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
1 row in set (0.00 sec)

-- 適当なデータを入れる
mysql> insert into gctest3 (a) values (1), (2), (3), (4), (5);
Query OK, 5 rows affected (0.00 sec)

-- Extra = Using index
mysql> explain select b from gctest3 where b = 3;
+----+-------------+---------+------------+-------+---------------+------+---------+-------+------+----------+-------------+
| id | select_type | table   | partitions | type  | possible_keys | key  | key_len | ref   | rows | filtered | Extra       |
+----+-------------+---------+------------+-------+---------------+------+---------+-------+------+----------+-------------+
|  1 | SIMPLE      | gctest3 | NULL       | const | b             | b    | 5       | const |    1 |   100.00 | Using index |
+----+-------------+---------+------------+-------+---------------+------+---------+-------+------+----------+-------------+
```

ここまでをまとめると次のようになると思われる。(推測のみなので間違っている可能性があります)

| | データ領域<br/>作成・更新時 | データ領域<br/>参照時 | データ領域<br/>値の保持 | インデックス<br/>作成・更新時 | インデックス<br/>参照時 | インデックス<br/>値の保持 |
| --- | --- | --- | --- | --- | --- | --- |
| STORED | 計算が発生 | 計算なし | 保持する | 計算が発生 | 計算なし | 保持する |
| VIRTUAL | 計算なし | 計算が発生 | 保持しない | 計算が発生 | 計算なし | 保持する |

- STORED はデータ領域とインデックスの両方にデータが保持されるのに対して、VIRTUAL はインデックス側にだけ保持されることになる
- VIRTUAL は更新時にインデックスのライトコストがかかるものの、インデックスがうまく作用すればリードコストが STORED と同水準になり、さらに空間効率の分有利になる可能性がありそう
- よって計算式が比較的単純な場合には VIRTUAL はよさそうかも?

### ところで Secondary Index とは

- [MySQL :: MySQL 8\.0 Reference Manual :: 15\.6\.2\.1 Clustered and Secondary Indexes](https://dev.mysql.com/doc/refman/8.0/en/innodb-index-types.html) や [mysql \- How secondary index scan works in InnoDB? \- Stack Overflow](https://stackoverflow.com/questions/4764693/how-secondary-index-scan-works-in-innodb) より
- Clustered (primary) index と Secondary Index
    - Clustered index は行のデータを保持するインデックスで、通常は主キーと同義
    - Clustered index 以外は Secondary Index
- Secondary index にはそれぞれの行に対する primary key も一緒に保存されている
    - イメージだが、Secondary Index だけですべてのデータが取得できないクエリは、Secondary Index に同梱されている主キーを経由して Clustered Index を参照してデータを取り出すという感じだろうか
- ユーザー目線では、VIRTUAL なカラムは主キーにはできないくらいでイメージしておけば良さそう?

## オプティマイザ

- クエリで直接 Genarated Columns を指定していなくても、計算式と同じ条件があれば、オプティマイザはうまくインデックスを使ってくれる

```sql
-- カラム b は `a + 1` という計算式
mysql> SHOW CREATE TABLE gctest3\G
*************************** 1. row ***************************
       Table: gctest3
Create Table: CREATE TABLE `gctest3` (
  `a` int DEFAULT NULL,
  `b` int GENERATED ALWAYS AS ((`a` + 1)) VIRTUAL,
  UNIQUE KEY `b` (`b`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

-- `a + 1 > 5` という条件句で、`b` インデックスが使われている
mysql> explain select * from gctest3 where a + 1 > 5;
+----+-------------+---------+------------+-------+---------------+------+---------+------+------+----------+-------------+
| id | select_type | table   | partitions | type  | possible_keys | key  | key_len | ref  | rows | filtered | Extra       |
+----+-------------+---------+------------+-------+---------------+------+---------+------+------+----------+-------------+
|  1 | SIMPLE      | gctest3 | NULL       | range | b             | b    | 5       | NULL |    1 |   100.00 | Using where |
+----+-------------+---------+------------+-------+---------------+------+---------+------+------+----------+-------------+

-- 書き換え後のクエリは show warnings で確認できる
-- where 句が `b > 5` となっている
mysql> show warnings;
+-------+------+--------------------------------------------------------------------------------------------------------------------------------------+
| Level | Code | Message                                                                                                                              |
+-------+------+--------------------------------------------------------------------------------------------------------------------------------------+
| Note  | 1003 | /* select#1 */ select `test`.`gctest3`.`a` AS `a`,`test`.`gctest3`.`b` AS `b` from `test`.`gctest3` where (`test`.`gctest3`.`b` > 5) |
+-------+------+--------------------------------------------------------------------------------------------------------------------------------------+
```

- オペランドの順番が違うとだめとか、対象の演算子は ` =, <, <=, >, >=, BETWEEN, IN()` という制約はある 


## Generated Columns のユースケース

- 複雑な条件句をシンプルにできる
    - 条件区の定義を一箇所でできるので、例えばアプリ側で複数箇所で同じ条件くを毎回書くリスクを減らせる
        - アプリ側でもやりようはあるが
        - 個人的には条件式が DDL に明記されているのは良いと思う
- STORED やインデックス付きの VIRTUAL の場合、毎回計算するコストが高い条件のキャッシュという使い方ができる
- functional index のシミュレーション
    - 例えば json 型のカラムの特定のキーにインデックスを張るといったことができる

```sql
-- json の id プロパティにインデックスを張る
mysql> CREATE TABLE gctest4 (a JSON, b INT GENERATED ALWAYS AS (a->"$.id") UNIQUE);

mysql> INSERT INTO gctest4 (a) VALUES ('{"id": "1", "name": "Alice"}'), ('{"id": "2", "name": "Bob"}');

mysql> select a->>"$.name" from gctest4 where b >= 2;
+--------------+
| a->>"$.name" |
+--------------+
| Bob          |
+--------------+

-- b が使われている
mysql> explain select a->>"$.name" from gctest4 where b >= 2;
+----+-------------+---------+------------+-------+---------------+------+---------+------+------+----------+-------------+
| id | select_type | table   | partitions | type  | possible_keys | key  | key_len | ref  | rows | filtered | Extra       |
+----+-------------+---------+------------+-------+---------------+------+---------+------+------+----------+-------------+
|  1 | SIMPLE      | gctest4 | NULL       | range | b             | b    | 5       | NULL |    1 |   100.00 | Using where |
+----+-------------+---------+------------+-------+---------------+------+---------+------+------+----------+-------------+

mysql> show warnings;
+-------+------+-----------------------------------------------------------------------------------------------------------------------------------------------------------+
| Level | Code | Message                                                                                                                                                   |
+-------+------+-----------------------------------------------------------------------------------------------------------------------------------------------------------+
| Note  | 1003 | /* select#1 */ select json_unquote(json_extract(`test`.`gctest4`.`a`,'$.name')) AS `a->>"$.name"` from `test`.`gctest4` where (`test`.`gctest4`.`b` >= 2) |
+-------+------+-----------------------------------------------------------------------------------------------------------------------------------------------------------+
```

## 今のプロジェクトの例

- 今のプロジェクトでは、論理削除を行うテーブルでユニーク制約をかけるために Generated Columns を使っていた
- 例えばこのようなテーブル
    - deleted_at で論理削除を実現する
        - NULL なら未削除、日時が入っていれば削除済

```sql
CREATE TABLE sample (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    deleted_at DATETIME NULL
)
```

- name にユニーク制約をつけたいが、このままだと論理削除したレコードも含めたユニークとなってしまう
    - 削除済みの名前は再利用したいという要件
- name と deleted_at の複合インデックスにすると、同じ name を複数 INSERT できてしまう
    - MySQL ではユニークインデックスに複数の NULL カラムを入れられる
        - [MySQL :: MySQL 8\.0 Reference Manual :: 13\.1\.15 CREATE INDEX Statement](https://dev.mysql.com/doc/refman/8.0/en/create-index.html#create-index-unique)
            - `A UNIQUE index permits multiple NULL values for columns that can contain NULL.`

```sql
mysql> CREATE TABLE table1 (x INT NULL UNIQUE);
Query OK, 0 rows affected (0.03 sec)

mysql> INSERT table1 VALUES (1);
Query OK, 1 row affected (0.01 sec)

mysql> INSERT table1 VALUES (1);
ERROR 1062 (23000): Duplicate entry '1' for key 'table1.x'
mysql> INSERT table1 VALUES (NULL);
Query OK, 1 row affected (0.01 sec)

mysql> INSERT table1 VALUES (NULL);
Query OK, 1 row affected (0.01 sec)

mysql> INSERT table1 VALUES (NULL);
Query OK, 1 row affected (0.00 sec)

mysql> SELECT * FROM table1;
+------+
| x    |
+------+
| NULL |
| NULL |
| NULL |
|    1 |
+------+
4 rows in set (0.00 sec)
```

- そこで Generated Columns を使いそちらにユニークインデックスを張ることで要件を満たす
    - `not_archived TINYINT GENERATED ALWAYS AS (IF(deleted_at IS NULL,  1, NULL))` というカラムを作り、name との複合インデックスにする
    - この時 `deleted_at IS NOT NULL` の場合、`not_archived` は NULL になるのがポイント
        - NULL ならば複数レコードが共存できるので、「削除済みのものは重複可」「未削除のものは重複不可」を実現できる

```sql
mysql> ALTER TABLE sample ADD COLUMN not_archived TINYINT GENERATED ALWAYS AS (IF(deleted_at IS NULL,  1, NULL));

mysql> ALTER TABLE sample ADD UNIQUE (name, not_archived);

mysql> show create table sample\G
*************************** 1. row ***************************
       Table: sample
Create Table: CREATE TABLE `sample` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `not_archived` tinyint GENERATED ALWAYS AS (if((`deleted_at` is null),1,NULL)) VIRTUAL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`,`not_archived`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
```

- 使っているフレームワークやミドルウェアの仕様によって deleted_at がこの例のように実装されているケースは多いと思うので、その際には使えるテクニックだと思われる

## 参考

- [MySQL :: MySQL 8\.0 Reference Manual :: 13\.1\.20\.8 CREATE TABLE and Generated Columns](https://dev.mysql.com/doc/refman/8.0/en/create-table-generated-columns.html)
- [MySQL :: MySQL 8\.0 Reference Manual :: 13\.1\.20\.9 Secondary Indexes and Generated Columns](https://dev.mysql.com/doc/refman/8.0/en/create-table-secondary-indexes.html)
[MySQL :: MySQL 8\.0 Reference Manual :: 8\.3\.11 Optimizer Use of Generated Column Indexes](https://dev.mysql.com/doc/refman/8.0/en/generated-column-index-optimizations.html)
- [Dealing with MySQL nulls and unique constraint \| by Aleksandra Sikora \| Medium](https://medium.com/@aleksandrasays/dealing-with-mysql-nulls-and-unique-constraint-d260f6b40e60)
- [database \- Does MySQL ignore null values on unique constraints? \- Stack Overflow](https://stackoverflow.com/questions/3712222/does-mysql-ignore-null-values-on-unique-constraints)
- [MySQL :: MySQL 8\.0 Reference Manual :: 13\.1\.15 CREATE INDEX Statement](https://dev.mysql.com/doc/refman/8.0/en/create-index.html#create-index-unique)
[MySQL :: MySQL 8\.0 Reference Manual :: 15\.6\.2\.1 Clustered and Secondary Indexes](https://dev.mysql.com/doc/refman/8.0/en/innodb-index-types.html)
- [mysql \- How secondary index scan works in InnoDB? \- Stack Overflow](https://stackoverflow.com/questions/4764693/how-secondary-index-scan-works-in-innodb)

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774142948/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://images-na.ssl-images-amazon.com/images/I/41oqE-9dM2L._SX394_BO1,204,203,200_.jpg" alt="エキスパートのためのMySQL[運用+管理]トラブルシューティングガイド" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774142948/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">エキスパートのためのMySQL[運用+管理]トラブルシューティングガイド</a></div><div class="amazlet-detail">奥野 幹也  (著, 編集)<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774142948/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
