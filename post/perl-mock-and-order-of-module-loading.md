{"title":"モジュールをモックに差し替える, あとモジュールの読み込み順","date":"2012-10-03T23:22:50+09:00","tags":["perl"]}

- for testing
- api 叩いたり, データストアにアクセスするなど, 外部に依存するパッケージをまるまるモックに置き換えたい場合
- 同じ名前のモジュールを `t/lib` とか `t/mock` とかに作っておく
- テストを流す前に `use lib "そのパス"` とかでモジュールのサーチパスに追加
- perl のモジュールの読み込み機構がわかってないけど, 先に読み込んだ同名のモジュールを上書きしない or 後から読み込んだモジュールを優先するという仕組みを利用してプロダクションのコードがモックの方を呼ぶようになる
  - あとで調べる. 前者だとサーチパスのリストの先頭にモックを追加すればいいし, 後者なら末尾

### 調べた

まず `lib` モジュールは `@INC` の先頭にパスを追加する. `perldoc lib` によると `use lib LIST` は以下と同じ意味になる.

    BEGIN { unshift(@INC, LIST) }

モジュールは `@INC` のパスで先に出てきたものが優先される (というより, たぶん見つかり次第探索を打ち切っているんだと思う) らしいので, テストスクリプの中でモックモジュールのパスを `use lib` で追加してあげるとそちらで置き換えることができるという仕組みのようだ.

一応以下のコードで検証した.

    [23:13 kosei@mba module_test]% tree
    .
    ├── lib_a
    │   └── Foo.pm
    ├── lib_b
    │   └── Foo.pm
    ├── test01.pl
    └── test02.pl

`lib_a/Foo.pm`, `lib_b/Foo.pm` と同じ名前のパッケージを 2 つ用意しておき, それぞれ読み込み順を変えて検証してみる

以下のように 2 つの `Foo.pm` は `foo` メソッドの戻り値が違っている

    [23:14 kosei@mba module_test]% cat lib_a/Foo.pm
    package Foo;
    
    use strict;
    use warnings;
    
    sub foo { 1 }
    
    1;

1 を返す

    [23:14 kosei@mba module_test]% cat lib_b/Foo.pm
    package Foo;
    
    use strict;
    use warnings;
    
    sub foo { 100 }
    
    1;

こっちは 100 を返している.

`test01.pl` と `test02.pl` はそれぞれ `lib_a`, `lib_b` のパスを追加順序が違う

    [23:15 kosei@mba module_test]% cat test01.pl
    #!/usr/bin/env perl
    
    use 5.012;
    use warnings;
    use lib qw/lib_a lib_b/;
    
    use Foo;
    
    say Foo->foo();

先に `lib_a`

    [23:15 kosei@mba module_test]% cat test02.pl
    #!/usr/bin/env perl
    
    use 5.012;
    use warnings;
    use lib qw/lib_b lib_a/;
    
    use Foo;
    
    say Foo->foo();

こっちは `lib_b` が先.

実行結果. 予想通り先に読み込まれた方の挙動になっている.

    [23:18 kosei@mba module_test]% perl test01.pl
    1
    [23:18 kosei@mba module_test]% perl test02.pl
    100
