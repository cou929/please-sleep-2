{"title":"redis-cli でメモリ使用量の多いキーを探すには `--memkeys` を使う","date":"2021-10-17T02:10:00+09:00","tags":["redis"]}

`redis-cli` には `--bigkeys` というオプションがあり、要素数の大きいキーや、型ごとのキー数の割合、平均サイズなどを集計してくれる。また `--memkeys` というオプションもあり、こちらは要素数ではなく実際のメモリ使用量をベースに集計する。いずれも調査時などに自分でスクリプトで集計しなくても良いので便利。

`--memkeys` はおそらく [5.0.4 で導入された](https://github.com/redis/redis/blob/f72f4ea311d31f7ce209218a96afb97490971d39/00-RELEASENOTES#L190) オプションで、[MEMORY USAGE](https://redis.io/commands/memory-usage) コマンドの結果を集計してくれる。実際にどのキーがメモリを多く使っているかなどの調査には `--bigkeys` ではなくこちらを使ったほうがよい。

`--memkeys` は [redis\-cli, the Redis command line interface – Redis](https://redis.io/topics/rediscli) などの Web のドキュメントには記載が見つからず (cli のヘルプには記載があるが)、たどり着くまでに時間がかかったのでメモ。

## 出力のサンプルと見方

集計結果はこんな感じで出力される。

```sh
# bigkeys の例。memkeys も似たようなフォーマットで、要素数の代わりにメモリ使用量が出力される。

$ redis-cli --bigkeys

# Scanning the entire keyspace to find biggest keys as well as
# average sizes per key type.  You can use -i 0.1 to sleep 0.1 sec
# per 100 SCAN commands (not usually needed).

[00.00%] Biggest string found so far 'key-419' with 3 bytes
[05.14%] Biggest list   found so far 'mylist' with 100004 items
[35.77%] Biggest string found so far 'counter:__rand_int__' with 6 bytes
[73.91%] Biggest hash   found so far 'myobject' with 3 fields

-------- summary -------

Sampled 506 keys in the keyspace!
Total key length in bytes is 3452 (avg len 6.82)

Biggest string found 'counter:__rand_int__' has 6 bytes
Biggest   list found 'mylist' has 100004 items
Biggest   hash found 'myobject' has 3 fields

504 strings with 1403 bytes (99.60% of keys, avg size 2.78)
1 lists with 100004 items (00.20% of keys, avg size 100004.00)
0 sets with 0 members (00.00% of keys, avg size 0.00)
1 hashs with 3 fields (00.20% of keys, avg size 3.00)
0 zsets with 0 members (00.00% of keys, avg size 0.00)


# memkeys の場合はこんな感じですべて `with xxx bytes` 表記になる。

0 hashs with 0 bytes (00.00% of keys, avg size 0.00)
0 lists with 0 bytes (00.00% of keys, avg size 0.00)
0 strings with 0 bytes (00.00% of keys, avg size 0.00)
0 streams with 0 bytes (00.00% of keys, avg size 0.00)
0 sets with 0 bytes (00.00% of keys, avg size 0.00)
0 zsets with 0 bytes (00.00% of keys, avg size 0.00)
```

基本的には見たとおりだが、注意が必要なのは `avg size`。この項目はその行で何を集計しているかに依存していて、常に平均のバイト数などが表示されているわけではない。例えば以下の `--bigkeys` での hash 型の集計結果について、hash キーのフィールド数の平均が 3 という意味になる。(hash キーの平均メモリ使用量が 3 バイトと言う意味ではない)

```
1 hashs with 3 fields (00.20% of keys, avg size 3.00)
```

## 実装の詳細

実装は [このへん](https://github.com/redis/redis/blob/4be2dd6ab98a66e5e2cb92b66ac93d3b49dc4219/src/redis-cli.c#L7564)。おおまかには [SCAN](https://redis.io/commands/scan) で全キーを走査し、それぞれのキーの要素数やサイズなどを取得し集計している。

- まずキーの総数を [DBSIZE](https://redis.io/commands/scan) コマンドで [取得](https://github.com/redis/redis/blob/4be2dd6ab98a66e5e2cb92b66ac93d3b49dc4219/src/redis-cli.c#L7399)
- [SCAN](https://redis.io/commands/scan) でキーを走査しながら、[それぞれについて要素数やメモリ使用量を取得](https://github.com/redis/redis/blob/4be2dd6ab98a66e5e2cb92b66ac93d3b49dc4219/src/redis-cli.c#L7518-L7532) する
    - それぞれ最大のキーと、これまでに走査した要素数、サイズの累計を計算していく
    - `--memkeys` の場合、[MEMORY USAGE](https://redis.io/commands/memory-usage) コマンドの結果を [集計する](https://github.com/redis/redis/blob/4be2dd6ab98a66e5e2cb92b66ac93d3b49dc4219/src/redis-cli.c#L7523)
    - `--bigkeys` の場合、キーの型によって [それぞれ以下のコマンドを使う](https://github.com/redis/redis/blob/4be2dd6ab98a66e5e2cb92b66ac93d3b49dc4219/src/redis-cli.c#L7429-L7435)
        - 例えば hash の場合 `HLEN` の値を取得する

```c
# type, コマンド、単位
typeinfo type_string = { "string", "STRLEN", "bytes" };
typeinfo type_list = { "list", "LLEN", "items" };
typeinfo type_set = { "set", "SCARD", "members" };
typeinfo type_hash = { "hash", "HLEN", "fields" };
typeinfo type_zset = { "zset", "ZCARD", "members" };
typeinfo type_stream = { "stream", "XLEN", "entries" };
typeinfo type_other = { "other", NULL, "?" };
```

- 収集したデータを出力する
    - キーの型ごとに最大の要素数・メモリ使用量だったキーを [表示する](https://github.com/redis/redis/blob/4be2dd6ab98a66e5e2cb92b66ac93d3b49dc4219/src/redis-cli.c#L7672-L7677)
    - キーの型ごとにキー数、要素数・メモリ資料用、キー数の全体に占める割合、平均サイズを [出力する](https://github.com/redis/redis/blob/4be2dd6ab98a66e5e2cb92b66ac93d3b49dc4219/src/redis-cli.c#L7686-L7689)
        - `--bigkeys` の場合、キーの型によって出力されるものが違う。例えば hash ならフィールド数など (上記参照)
        - `--memkeys` の場合すべて `MEMORY USAGE` で調べたバイト数が表示される
        - 平均サイズ (`avg size`) は、単に [合計値をキー数で割っただけの値](https://github.com/redis/redis/blob/4be2dd6ab98a66e5e2cb92b66ac93d3b49dc4219/src/redis-cli.c#L7689)。例えば `--bigkeys` の場合は平均の要素数 (hash のフィールド数など) であり、バイト数の平均ではないので注意

## 背景

- 運用している redis プロセスのメモリサイズが増加したことがあり、その調査をしていた
- 自分で集計スクリプトを書こうかと思っていたら、`redis-cli` にちょうどよく `--bigkeys` というオプションが有ることに気づいてそれをまず使ってみた
- 出力の見方がよくわからず、特に `avg size` にハマった
    - `avg size` はそのキーが使っている平均バイト数かと思ったが、計算してもプロセス全体のメモリ使用量と一致しなかった
- 軽く実装から追ってみると実は `--memkeys` という別のオプションがあとから追加されていることに気づき、こちらを使うのが正解だとわかった

## PR

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873119545/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://images-na.ssl-images-amazon.com/images/I/510rXBOSUTS._SX389_BO1,204,203,200_.jpg" alt="詳説 データベース ―ストレージエンジンと分散データシステムの仕組み" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873119545/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">詳説 データベース ―ストレージエンジンと分散データシステムの仕組み</a></div><div class="amazlet-detail">Alex Petrov (著), 小林 隆浩 (監修), 成田 昇司 (翻訳)<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873119545/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
