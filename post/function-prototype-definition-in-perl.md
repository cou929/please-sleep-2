{"title":"perl のプロトタイプ宣言","date":"2012-10-06T13:06:24+09:00","tags":["perl"]}

- あんまり使われていない理由は意図しない挙動になっちゃうことが多いから
- `&foo()` と呼び出すとプロトタイプ宣言は無視される
- プロトタイプ宣言すると関数呼び出しのカッコを省ける
- main の名前空間で使うと将来的に組み込み関数とバッティングする可能性はあるけど, モジュール内で使う分には大丈夫なので, dsl っぽい記法にしたい時などに限定的に使うとよさそう
  - `&` を指定して, map とか grep みたいにブロックを渡すような構文も作れる

プログラミング perl にあるらしい例

    sub try(&$){
      my ($try, $catch) = @_;
      eval { &$try };
      if ($@) {
        local $_ = $@;
        &$catch;
      }
    }
    sub catch(&){ $_[0] }
    #実行
    try{
      die "phooey";
    }
    catch{
      /phooey/ and print "unphooey\n";
    };

### Ref.
[おそらくはそれさえも平凡な日々: Perlのサブルーチンプロトタイプについて](http://www.songmu.jp/riji/archives/2009/03/perl_1.html)
