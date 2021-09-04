{"title":"MySQL の外部キーのため自動作成されたインデックスは不要になると自動で削除される","date":"2021-09-04T16:00:00+09:00","tags":["mysql"]}

知らなかったのでメモ。

- MySQL の外部キーは相当するカラムのインデックスが必要
- 必要なインデックスがない場合は外部キー作成時に自動でインデックスも作成される
- こうして自動で作成されたインデックスは、不要になった際に自動で drop される
    - 不要になった際 = その外部キーをカバーできるような別のインデックスが追加された場合

[MySQL :: MySQL 5\.7 Reference Manual :: 1\.7\.3\.2 FOREIGN KEY Constraints](https://dev.mysql.com/doc/refman/5.7/en/constraint-foreign-key.html)

> MySQL requires that foreign key columns be indexed; if you create a table with a foreign key constraint but no index on a given column, an index is created.

[MySQL :: MySQL 5\.7 Reference Manual :: 13\.1\.18\.5 FOREIGN KEY Constraints](https://dev.mysql.com/doc/refman/5.7/en/create-table-foreign-keys.html#foreign-key-restrictions)

> MySQL requires indexes on foreign keys and referenced keys so that foreign key checks can be fast and not require a table scan. In the referencing table, there must be an index where the foreign key columns are listed as the first columns in the same order. Such an index is created on the referencing table automatically if it does not exist. This index might be silently dropped later if you create another index that can be used to enforce the foreign key constraint. index_name, if given, is used as described previously.

## 試してみる

- 必要なインデックスが自動で作成される

```sql
-- table1, 2 を作成
mysql> create table table1 (
    ->     id int not null auto_increment,
    ->     primary key (id)
    -> );
Query OK, 0 rows affected (0.02 sec)

mysql> create table table2 (
    ->     id int not null auto_increment,
    ->     table1_id int not null,
    ->     primary key (id)
    -> );
Query OK, 0 rows affected (0.02 sec)

-- table2.table1_id => table1.id という外部キーを追加
mysql> alter table table2 add foreign key (table1_id) references table1 (id);
Query OK, 0 rows affected (0.04 sec)
Records: 0  Duplicates: 0  Warnings: 0

-- `KEY `table1_id` (`table1_id`)` が自動で作成されている
mysql> show create table table2\G
*************************** 1. row ***************************
       Table: table2
Create Table: CREATE TABLE `table2` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `table1_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `table1_id` (`table1_id`),
  CONSTRAINT `table2_ibfk_1` FOREIGN KEY (`table1_id`) REFERENCES `table1` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
1 row in set (0.00 sec)
```

- 自動インデックスが不要になると自動で drop される

```sql
-- table2 に table1_id, id の複合インデックスを追加する (これがあると table1_id のみのインデックスは不要)
mysql> alter table table2 add index test_idx (table1_id, id);
Query OK, 0 rows affected (0.03 sec)
Records: 0  Duplicates: 0  Warnings: 0

-- `KEY `table1_id` (`table1_id`)` が無くなっている
mysql> show create table table2\G
*************************** 1. row ***************************
       Table: table2
Create Table: CREATE TABLE `table2` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `table1_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `test_idx` (`table1_id`,`id`),
  CONSTRAINT `table2_ibfk_1` FOREIGN KEY (`table1_id`) REFERENCES `table1` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
1 row in set (0.00 sec)
```

- 同じ内容でも、自動で作成されたインデックスでなければ、自動で drop されない

```sql
-- table3 として外部キー、そのためのインデックスがある状態で crete table する
mysql> create table table3 (
    ->   id int not null auto_increment,
    ->   table1_id int not null,
    ->   primary key (id),
    ->   key my_table1_id_key (table1_id),
    ->   constraint my_table3_fk foreign key (table1_id) references table1 (id)
    -> );
Query OK, 0 rows affected (0.02 sec)

-- table2 同様に table1_id, id の複合インデックスを追加する
mysql> alter table table3 add index test_idx (table1_id, id);
Query OK, 0 rows affected (0.02 sec)
Records: 0  Duplicates: 0  Warnings: 0

-- table1_id のインデックスが残っている
mysql> show create table table3\G
*************************** 1. row ***************************
       Table: table3
Create Table: CREATE TABLE `table3` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `table1_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `my_table1_id_key` (`table1_id`),
  KEY `test_idx` (`table1_id`,`id`),
  CONSTRAINT `my_table3_fk` FOREIGN KEY (`table1_id`) REFERENCES `table1` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
1 row in set (0.00 sec)
```

## PR

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01LCJRCYE/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/51GD7yZsLVL.jpg" alt="詳解MySQL 5.7 止まらぬ進化に乗り遅れないためのテクニカルガイド" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01LCJRCYE/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">詳解MySQL 5.7 止まらぬ進化に乗り遅れないためのテクニカルガイド</a></div><div class="amazlet-detail">奥野幹也 (著)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01LCJRCYE/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
