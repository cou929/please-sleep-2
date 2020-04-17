{"title":"perl のマッチ変数のスコープではまった","date":"2013-08-09T20:57:45+09:00","tags":["perl"]}

正規表現でキャプチャした結果にあらぬ文字列が入っていて、そのあとの処理が壊れるというバグに悩まされた。原因はマッチ変数 (`$1`, `$2` などのキャプチャした結果を保持する特殊変数) のスコープの理解が甘かったからだ。perl でははまりがちなポイントだと思う。

たとえば以下のコード。"foo" という文字列がプリントされる。

<pre><code data-language="perl">use strict;
use warnings;

"foo" =~ /(foo)/;

{
    "bar" =~ /(baz)/;
    print $1;   # foo
}</code></pre>

この挙動を理解するには、まずは perlre を見てみる。[Capture Groups] [1] の節より、

> Capture group contents are dynamically scoped and available to you outside the pattern until the end of the enclosing block or until the next successful match, whichever comes first.

ポイントは、

- マッチ変数はダイナミックスコープ
- マッチ変数の中身は次の成功したマッチ結果 (またはスコープが終わった時点) で更新される

ダイナミックスコープとは、perl でいう `local` と思えばいい。いまのスコープよりグローバルなスコープに `foo` という変数があった場合、いまのスコープで foo を local 宣言すると、このスコープ内では foo の値を書き換えて使える。今のスコープを抜けると foo の値はもとに戻る。あるグローバル変数の値を、現在のスコープの間だけ一時的に退避させ変更し、スコープが終わると戻す戻す。イメージとしてはこう考えている。

マッチ変数はマッチしない限り更新されない。よって今回のケースでは、2 回めの正規表現ではマッチしないためマッチ変数に変化が起こらず、前回のマッチ結果が入ったままになっていた。前回のマッチは一段外のスコープで行われたので、より内側のスコープでも参照可能だったということになる。

対処法としては、マッチしたかどうかを確認すればよい。

<pre><code data-language="perl">use strict;
use warnings;

"foo" =~ /(foo)/;

{
    my $res = "not matched";
    if ("bar" =~ /(baz)/) {
        $res = $1;
    }
    print $res;  # not matched
}</code></pre>

キャプチャ結果を取り出す方法にはいくつかあるようなので、もっとスマートな方法もありそうだ。

[perlvar] [2] にも具体的な例を用いたわかりやすい説明がある。サンプルコードを引用。

<pre><code data-language="perl">my $outer = 'Wallace and Grommit';
my $inner = 'Mutt and Jeff';

my $pattern = qr/(\S+) and (\S+)/;

sub show_n { print "\$1 is $1; \$2 is $2\n" }

{
OUTER:
    show_n() if $outer =~ m/$pattern/;

    INNER: {
        show_n() if $inner =~ m/$pattern/;
    }

    show_n();
}</code></pre>

実行結果はこうなる。

    $1 is Wallace; $2 is Grommit
    $1 is Mutt; $2 is Jeff
    $1 is Wallace; $2 is Grommit

[1]: http://perldoc.perl.org/perlre.html#Capture-groups
[2]: http://perldoc.perl.org/perlvar.html#Variables-related-to-regular-expressions
