{"title":"perl の標準デバッガと plack アプリのデバッグ","date":"2012-10-06T13:07:00+09:00","tags":["perl"]}

    $ perl -d foo.pl

とすると gdb-like なデバッガが立ち上がるが, plack なアプリをデバッグする際も同様にできる.

`perl -d <plackup へのフルパス> foo.psgi` などとすればいいんだけど, フルパスは面倒なので `perl -Sd plackup foo.psgi` と省略できる. (plackup へ PATH が通っている前提)

フレームワークを使ったアプリになるとスタックが深くなるので,

    $DB::single = 1;

などとコード側で一旦ブレイクポイントを設定しておいて, そこを足がかりにデバッグしていくといいかもしれない.

### コマンド復習

- h
  - help
- n
  - 一行実行. 一回選ぶと次からは enter でいい
- s
  - 関数の中に入る
- c
  - 次のブレイクポイントまで continue
  - 数値を渡すとその行まで行く
- p
  - 変数をプリント
- x
  - 変数を展開してプリント
- v
  - 周囲のコードを見る
- b
  - ブレイクポイント設置

### Ref.
- [perldebug - Perl のデバッグ - perldoc.jp](http://perldoc.jp/docs/perl/5.8.8/perldebug.pod)
- [Perlデバッガの手引き - サンプルコードによるPerl入門](http://d.hatena.ne.jp/perlcodesample/20100302/1269670120)
