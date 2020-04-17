{"title":"インデックスが効くクエリ・効かないクエリ","date":"2013-12-09T23:17:01+09:00","tags":["nix"]}

MySQL を使っていて、datetime 型のカラムにインデックスを追加する検証をしていたが、いまいちパフォーマンスが向上しない。クエリをみてみると WHERE 句に DATE 関数を使っているものは、追加したインデックスが効いていなかった。よく考えてみると当然だが、datetime 型のインデックスを違う型で検索することはできない。例えば DATE 関数を使ったクエリが多いのであれば、関数を適用したあとの値でインデックスを作成することも原理的にはできそうだが、MySQL は対応していなかった。

ググると、RDBMS 一般でインデックスが効く演算子のことを Sargable、そうでないものを Non-Sargable と呼ぶことがあるそうだ。

[Sargable - Wikipedia, the free encyclopedia](http://en.wikipedia.org/wiki/Sargable)

あまり広く使われている用語には見えないが、内容は納得のいくもの。

- Sargable なオペレーター
  - `=`, `>`, `<`, `>=`, `<=`, `BETWEEN`, % からはじまらない `LIKE`
- Sargable だが必ずパフォーマンスが向上するわけではないもの
  - `<>`, `IN`, `OR`, `NOT IN`, ` NOT EXISTS`, `NOT LIKE`
- Non-sargable なもの
  - % ではじまる `LIKE`, 条件句での関数利用

Non-sargable なクエリを sargable に書き換えることができるため、基本的にそうすべき。例えば `Select ... WHERE Year(date) = 2012` は `Select ... WHERE date >= '01-01-2012' AND date < '01-01-2013'` とする。

