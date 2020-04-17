{"title":"perl のモジュールインポートまわりの整理","date":"2013-11-17T11:25:44+09:00","tags":["perl"]}

### use と require のちがい

`perldoc -f use` より、

> It is exactly equivalent to
>
>     BEGIN { require Module; Module->import( LIST ); }

とのこと。

- BEGIN ブロックの中で require しているので、つまりコンパイル時に評価されるということ。コード中のどこで use してもコンパイル時に読み込み処理が行われる。逆に conditional にモジュールを読み込みたい場合は require を使う。
- requrie と同時に import も行う。Exporter の import メソッドを想定していて、Exporter によってエクスポートされた関数を呼び出し側の名前空間に import する。import 関数は別に予約語でもないし、特別扱いもされていない。モジュールがたまたま (Exporter とは関係のない) import というメソッドを持っていたとしても、それが呼び出される。
- [use と require](http://tech.bayashi.net/pdmemo/use-require.html) によると、このほかにも use の場合は `INIT`、`CHECK` ルーチンが走るという違いがあるそうだ。

### require と %INC 特殊変数

`perldoc - require` より、require の動きはつぎのようなコードらしい。

<pre><code data-language="perl">sub require {
    my ($filename) = @_;
    if (exists $INC{$filename}) {
        return 1 if $INC{$filename};
        die "Compilation failed in require";
    }
    my ($realfilename,$result);
    ITER: {
        foreach $prefix (@INC) {
            $realfilename = "$prefix/$filename";
            if (-f $realfilename) {
                $INC{$filename} = $realfilename;
                $result = do $realfilename;
                last ITER;
            }
        }
        die "Can't find $filename in \@INC";
    }
    if ($@) {
        $INC{$filename} = undef;
        die $@;
    } elsif (!$result) {
        delete $INC{$filename};
        die "$filename did not return true value";
    } else {
        return $result;
    }
}</code></pre>

`return 1 if $INC{$filename};` とあるように、`%INC` をチェックして読み込み済みのものはロードしないよう制御している。

`%INC` は use/require/do でロードしたファイルをもつハッシュで、キーはモジュール名 (ファイル名)、値はそのパスである。

> The hash %INC contains entries for each filename included via
> the "do", "require", or "use" operators.  The key is the
> filename you specified (with module names converted to
> pathnames), and the value is the location of the file found.

 (`perldoc perlvar` より)

シンボルテーブルをみてみると `%INC` は main 名前空間に属していることがわかる。

    $ perl -le 'use Data::Dumper; print Dumper \%{main::}'
    $VAR1 = {
              'version::' => *{'::version::'},
              '/' => *{'::/'},
              'stderr' => *::stderr,
              ...
              'INC' => *::INC,
              ...
            };

よっておなじプロセスのなかで何度モジュールを読み込んでも、実際にロードされるのははじめの一度だけということになりそうだ。

またさきほどの use の擬似コード、

`BEGIN { require Module; Module->import( LIST ); }`

require は何度呼び出されても実際にモジュールをロードするのは最初の一回だが、import は use をされるたびに呼び出されるようだ。モジュールを読み込んでメモリに展開するのは一度だけにしたいが、メソッドを呼び出し側の名前空間に展開するのは呼び出すたびに行う必要があるためだと思う。

ちなみに `use FooModule qw//;` と空リストを渡すと import の処理が走らない。呼び出し側の名前空間を汚したくない場合はこうするとよい。

> Again, there is a distinction between omitting LIST ("import"
> called with no arguments) and an explicit empty LIST "()"
> ("import" not called).

 (`perldoc -f use` より)

### Exporter モジュール

モジュールのメソッドを呼び出し側の名前空間にエクスポートする機能を提供するコアモジュール。基本的な使い方は、

<pre><code data-language="perl">package YourModule;
require Exporter;
@ISA = qw(Exporter);
@EXPORT = qw(munge);
@EXPORT_OK = qw(frobnicate);</code></pre>

とモジュールに書いておくと、

- `use YourModule;` とした場合、`@EXPORT` に定義されているものがエクスポートされる (呼び出し側の名前空間に展開される)。
- `use YourModule qw/frobnicate/;` とした場合、frobnicate メソッドだけがエクスポートされる。
  - このようにリストを渡した場合は、`@EXPORT`、`@EXPORT_OK` から該当するメソッドを探してそれだけをエクスポートする。

Exporter.pm では import メソッドが定義されており、use が想定しているのはこれ。中をみてみると、

<pre><code data-language="perl">sub import {
  my $pkg = shift;
  my $callpkg = caller($ExportLevel);

  if ($pkg eq "Exporter" and @_ and $_[0] eq "import") {
    *{$callpkg."::import"} = \&import;
    return;
  }</code></pre>

というふうにシンボルテーブルに該当の関数を代入しているという仕組みだった。よって use をうまくつかうとテスト時にメソッドをスタブに差し替えたりが簡単にできそうと一瞬考えたが、同じ方法で直接シンボルテーブルに代入するなり、Test::Double を使うなりしたほうが手っ取り早いと思い直した。
