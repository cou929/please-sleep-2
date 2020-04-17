{"title":"use feature と use VERSION","date":"2012-10-03T23:49:16+09:00","tags":["perl"]}

あ, perl の話です.

### use feature

新しいバージョンでできた機能を読み込む. たとえば say が使いたかったら.

    use feature qw/say/;

とする.

perl バージョンも渡せて, たとえば perl 5.10 以降の全機能を使いたい場合は

    use feature ':5.10';

とする.

有効範囲はレキシカル. ブロック内だけで使う or 全体で `use feature` して特定ブロックで `no feature` という使い方もできる

    use feature 'say';
    say "say is available here";
    {
        no feature 'say';
        print "But not here.\n";
    }
    say "Yet it is here.";

### use VERSION

特定バージョン以上の perl での実行を保証する

    use 5.012;

として perl 5.12 以前の perl で実行するとエラーになる

    [23:42 kosei@mba /tmp]% cat test.pl
    use 5.012;
    use warnings;
    [23:42 kosei@mba /tmp]% perl5.10 test.pl
    Perl v5.12.0 required--this is only v5.10.1, stopped at test.pl line 1.
    BEGIN failed--compilation aborted at test.pl line 1.

また use VERSION だとそのバージョンからの機能をすべて使えるようになる. 次と同じ意味

    use feature ':<そのバージョン>';

### 5.12 と strict

5.12 からは `use strict` がデフォルトになったそうなので,

    use 5.012;
    use warnings;

は

    use strict;
    use warnings;
    use feature ':5.12';

と同じ意味になる.
