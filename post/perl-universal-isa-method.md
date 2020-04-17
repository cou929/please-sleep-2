{"title":"perl の UNIVERSAL::isa は文字列でも呼べる?","date":"2013-04-05T17:12:56+09:00","tags":["perl"]}

`perldoc perlobj` によると

>    Default UNIVERSAL methods
>        The "UNIVERSAL" package automatically contains the following methods that are inherited by all other classes:
>    
>        isa(CLASS)
>            "isa" returns true if its object is blessed into a subclass of "CLASS"

このように `UNIVERSAL::isa` メソッドはオブジェクトにつくものだ. がたとえば `"文字列"->isa()` という呼び方をしてもエラーにならないらしい.

    $ perl -le 'print "blah"->isa("Foo")'
    

アロー演算子の左辺は bless されたオブジェクトか, あるいはクラス名の文字列スカラーでもよい. よってクラス名に使えない文字列でなければ `isa` が呼べてしまったりする. クラス名に使えない `@`, `%`, `#` などから始まる文字列の場合は当然エラーになる.

    $ perl -le 'print "%blah"->isa("Foo")'
    Can't call method "isa" without a package or object reference at -e line 1.

さて, 今回やりたかったのは例外をキャッチしたときに例外クラスごとに処理を分けること. ただし例外はふつうに `Carp::croak` されることもある.

こういうふうにしたかった:

<pre><code data-language="perl">eval {
    # do something ...
};
if ($@) {
    if ( $@->isa('Error::Class::Foo') ) {
        # handle error foo
    } else {
        # handle normal exception from Carp::croak
    }
}</code></pre>

が上記のように `$@` が bless されているかクラス名文字列でなければならない.

例外クラスのサブクラスが無いならば `ref` でチェックするか:

<pre><code data-language="perl">eval {
    # do something ...
};
if ($@) {
    if ( ref $@ eq 'Error::Class::Foo' ) {
        # handle error foo
    } else {
        # handle normal exception from Carp::croak
    }
}</code></pre>

例外クラスにサブクラスもあり `ref` でのチェックがつらい場合はさらに `eval` で囲むか:

<pre><code data-language="perl">eval {
    # do something ...
};
if ($@) {
    if ( eval { $@->isa('Error::Class::Foo') } ) {
        # handle error foo
    } else {
        # handle normal exception from Carp::croak
    }
}</code></pre>

あるいは `Scalar::Util` の `blessed` で bless されているか事前にチェックする:

<pre><code data-language="perl">use Scalar::Util 'blessed';

eval {
    # do something ...
};
if ($@) {
    if ( blessed($@) && $@->isa('Error::Class::Foo') ) {
        # handle error foo
    } else {
        # handle normal exception from Carp::croak
    }
}</code></pre>

あたりかなと思っている. 最後が一番いいかな.

認識が正しいか不安.
