{"title":"InnoDB のフラグメンテーションがよくわからなかった","date":"2021-08-15T17:15:00+09:00","tags":["mysql"]}

InnoDB のフラグメンテーションについてドキュメントを読んだメモ。

なおいずれも MySQL8 系のドキュメントを参照している。また Issue Tracker やソースコードまでの深堀りはしていおらず、基本的にドキュメントから分かる範囲だけをまとめている。

## フラグメンテーションについて

[MySQL :: MySQL 8\.0 Reference Manual :: 15\.11\.4 Defragmenting a Table](https://dev.mysql.com/doc/refman/8.0/en/innodb-file-defragmenting.html) より。

- ランダムな INSERT や DELETE をしているうちに、だんだんと page のなかで「確保されているが使用されていない」領域が増えていく
- フラグメンテーションが大きくなると読み取り性能が劣化する可能性がある
- 次のような場合に偏っているいると考えられる
    - 「本来あるべきデータサイズ」よりも「実際使われているデータサイズ」が大きい場合
        - 「本来」とは何かが難しい。なぜなら B-tree インデックスの fill factor は 50% ~ 100% の間でのバリエーションがあるため
    - あるいは次のようなフルスキャンするクエリが「本来かかる時間 (とは何?)」よりも長い時間かかった場合、フラグメンテーションの悪影響の可能性がある

```sql
SELECT COUNT(*) FROM t WHERE non_indexed_column <> 12345;
```

### fill factor ?

[MySQL :: MySQL 8\.0 Reference Manual :: MySQL Glossary](https://dev.mysql.com/doc/refman/8.0/en/glossary.html#glos_fill_factor) より。

- fill factor とは page 分割のしきい値
    - 前提として innodb の各種データは [page という固定長の単位で管理](https://dev.mysql.com/doc/refman/8.0/en/innodb-init-startup-configuration.html#innodb-startup-page-size) されている
    - あるインデックスのデータが入っている page について、データが大きくなると page を分割して保存する
    - page のうち何割に達したら分割するかのしきい値が fill factor
- fill factor がある理由は更新処理の効率化のため
    - ある行がより大きいデータに更新された場合、インデックスのデータ未使用領域が足りていれば、一旦更新してコストが高い page の分割処理は別にまわすことができる
- つまり
    - fill factor が大きすぎる (なかなか page 分割されない) 設定の場合、インデックスの更新コストが高まる可能性がある
    - fill factor が小さすぎる (すぐに page 分割される) 設定の場合、インデックスの読み取りコストが高まる可能性がある。また空間効率も悪い

## フラグメンテーションの確認方法

前述の記載と少しかぶるが、[MySQL :: MySQL 8\.0 Reference Manual :: 15\.11\.4 Defragmenting a Table](https://dev.mysql.com/doc/refman/8.0/en/innodb-file-defragmenting.html) には確実な方法はなく、以下の曖昧な方法で兆候を探るよう書かれている。

- データ使用量が「本来」よりも多い場合
- フルスキャンのクエリが「本来」よりも長くかかる場合

ただ、ネット上には `INFORMATION_SCHEMA.TABLES.DATA_FREE` を参照してフラグメンテーションの判断をしている記事が散見された。([1](https://variable.jp/2020/04/29/mysql%E3%81%A8postgresql%E3%81%AB%E3%81%8A%E3%81%91%E3%82%8B%E3%82%A4%E3%83%B3%E3%83%87%E3%83%83%E3%82%AF%E3%82%B9%E3%81%AE%E3%83%A1%E3%83%B3%E3%83%86%E3%83%8A%E3%83%B3%E3%82%B9/), [2](https://serverfault.com/questions/202000/how-to-find-and-fix-fragmented-mysql-tables), [3](https://lefred.be/content/overview-of-fragmented-mysql-innodb-tables/))。次のように出した数値のうち、DATA_FREE 対 (DATA_LENGTH + INDEX_LENGTH) の割合がフラグメント率なのだそうだ。

```sql
SELECT
    ENGINE,
    TABLE_NAME,
    ROUND(DATA_LENGTH/1024/1024) AS data_length_in_MiB,
    ROUND(INDEX_LENGTH/1024/1024) AS index_length_in_MiB,
    ROUND(DATA_FREE/ 1024/1024) AS data_free_in_MiB
FROM
    INFORMATION_SCHEMA.TABLES
WHERE
    DATA_FREE > 0
;
```

### INFORMATION_SCHEMA.TABLES.DATA_FREE

[MySQL :: MySQL 8\.0 Reference Manual :: 26\.3\.38 The INFORMATION\_SCHEMA TABLES Table](https://dev.mysql.com/doc/refman/8.0/en/information-schema-tables-table.html) より。

- 確保されているが未使用領域。バイト単位
- InnoDB の場合そのテーブルが属する tablespace の未使用領域が表示される
    - system tablespace や general tablespaces のような共有の tablespace の場合、その共有のスペースでの未使用領域となる
    - テーブルが独自の tablespace に属する場合はそこの未使用領域となる

### innodb_file_per_table

[MySQL :: MySQL 8\.0 Reference Manual :: 15\.6\.3\.2 File\-Per\-Table Tablespaces](https://dev.mysql.com/doc/refman/8.0/en/innodb-file-per-table-tablespaces.html) より。

- `innodb_file_per_table` が有効な場合テーブルごとにファイル (filespace) が作られる
- デフォルトは有効

よって DATA_FREE も一定の参考にはしてもいいのかもしれない。

### フラグメンテーション確認方法まとめ

- 「本来」よりもデータサイズや実行時間が長いかどうかは、他に似たような DB インスタンスがある場合、比較して判断の参考にすることはできそう
- MySQL のドキュメントには記載されない方法だが、`innodb_file_per_table` がデフォルトオンである事も踏まえて、`INFORMATION_SCHEMA.TABLES.DATA_FREE` を参考にすることもできそう

## フラグメンテーションの解消方法 (Defragmenting)

[MySQL :: MySQL 8\.0 Reference Manual :: 15\.11\.4 Defragmenting a Table](https://dev.mysql.com/doc/refman/8.0/en/innodb-file-defragmenting.html) より。

- 次のような "null" ALTER でテーブルファイルを再作成する
    - `ALTER TABLE tbl_name ENGINE=INNODB`
    - または `ALTER TABLE tbl_name FORCE`
- いずれも Online DDL なので、ほかの DML をブロックしない
- あるいは mysqldump で text に落としてから、DROP TABLE しリロードしてもよい

これでインデックスのスキャンが早くなる可能性があるとのこと。

### OPTIMIZE TABLE Statement

上記のページには記載が無いが、OPTIMIZE TABLE でも同じことができそうだった。

以下は [MySQL :: MySQL 8\.0 Reference Manual :: 13\.7\.3\.4 OPTIMIZE TABLE Statement](https://dev.mysql.com/doc/refman/8.0/en/optimize-table.html#optimize-table-innodb-details) より。

- `OPTIMIZE TABLE ...` は ` ALTER TABLE ... FORCE` にマップされる
    - なので Online DDL を使う
- ただし Mysql 5.6.17 以前は Online DDL を使わず、他の DML をブロックするらしいので注意が必要そうだった
    - [MySQL :: MySQL 5\.6 Reference Manual :: 13\.7\.2\.4 OPTIMIZE TABLE Statement](https://dev.mysql.com/doc/refman/5.6/en/optimize-table.html#optimize-table-innodb-details) (ver 5.6 のドキュメント)
    - `Prior to Mysql 5.6.17, OPTIMIZE TABLE does not use online DDL. Consequently, concurrent DML (INSERT, UPDATE, DELETE) is not permitted on a table while OPTIMIZE TABLE is running, and secondary indexes are not created as efficiently.`

## デフラグはすべきなのか

ドキュメントを読むだけでは良くわからなかった。

- フラグメンテーションが起こっているのか、それがどの程度なのかの明確な指標がなく、デフラグを行う判断がしづらそう
    - `fill factor` に記載があったように、極端な偏りは問題だが、ある程度は更新性能のために未使用領域は必要
    - 仮に `FREE_LENGTH / (DATE_LENGTH + INDEX_LENGTH)` の割合を指標にするとしても、これがどの程度なら問題なのかわからない
- デフラグをするにしても、Online DDL を使うとはいえ、軽い処理ではない
    - 小さいテーブルならフラグメンテーションは気にならないと思われるし、大きいテーブルならデフラグは重いはず
- フラグメンテーションは DB の性能に対する影響の一部でしかなさそう
    - 最終的に読み書きの性能が出れば良いが、その場合フラグメンテーション以前に他の要因の影響であることも多そう
    - 解決方法として、クラウド環境ならば、デフラグをするよりもスケールアップやスケールアウトをしたほうが早い場合も多そう

## PR

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01LCJRCYE/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/51GD7yZsLVL.jpg" alt="詳解MySQL 5.7 止まらぬ進化に乗り遅れないためのテクニカルガイド" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01LCJRCYE/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">詳解MySQL 5.7 止まらぬ進化に乗り遅れないためのテクニカルガイド</a></div><div class="amazlet-detail">奥野幹也 (著)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01LCJRCYE/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
