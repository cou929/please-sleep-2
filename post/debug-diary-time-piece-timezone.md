{"title":"デバッグ日記: Time::Piece のタイムゾーン依存","date":"2013-06-06T19:00:07+09:00","tags":["debug"]}

cookie の expire をチェックするテスト。ある Plack のウェブアプリには、リクエストを受けた時点から n ヶ月後を expire に設定し cookie をセットする処理がある。この機能のテストとして、日付時刻を固定して処理を走らせ、レスポンスヘッダを見て、expire が想定した日時になっているかをチェックしていた。

今回はこのテストが環境によって通ったり落ちたりする現象。

時間周りの処理で環境依存で落ちるとなると、まず怪しいのはタイムゾーンの設定だが、結果からいうとまさに、テストケースがタイムゾーンに依存した書き方をしていることが原因だった。

まず注目したのは落ちたテストのデータ。got と expected がちょうど 9 時間ずれて落ちていたので、一発でタイムゾーンまわりにあたりを付けられた。GMT と JST の違いに起因する問題だ。

次にロジックを追ってみる。まずはテスト側。

おおまかな流れとして、サーバ側の処理を行うクラスのインスタンスを作り、そのクラスに擬似的な時刻を知らせる。時刻は Time::Piece のオブジェクトとして持っている。そのクラスに擬似的なリクエストを投げてレスポンスを受け取る。レスポンスのヘッダを見て、設定した時刻の n ヶ月後になっているかをチェックする。

`n = 2` とするとこんなかんじだ。Time::Piece::MySQL を使って、時刻を作っている。

<pre><code data-language="perl">use Time::Piece::MySQL;

my $now = localtime->from_mysql_datetime('2013-01-01 11:11:11');
my $server = Foo::Bar::ServerProcess->new( now => $now );

my $res = $server->run;

is_deeply $res->[1], [
...
'Set-Cookie'   => 'foo=bar; domain=.test.com; path=/; expires=Fri, 01-Mar-2013 02:11:11 GMT',
...
];</code></pre>

サーバ側の処理は、Plack::Response の cookies に、expire として now の時刻 + n ヶ月したものを epoch でいれている。

<pre><code data-language="perl">my $res = Plack::Response->new;
my $n = 10;

$res->cookies->{foo} = {
    value   => 'bar',
    expires => $now->add_months($n)->epoch;
    domain  => 'test.com',
    path    => '/',
};
</code></pre>

テストでは localtime から Time::Piece オブジェクトを作っているので、"2013-01-01 11:11:11" はシステムのタイムゾーンとして扱われる。一方 Plack::Response は cookie の expires に epoch が指定された場合は GMT として扱って文字列を作る。

<pre><code data-language="perl">sub _date {
    my($self, $expires) = @_;

    if ($expires =~ /^\d+$/) {
        # all numbers -> epoch date
        # (cookies use '-' as date separator, HTTP uses ' ')
        my($sec, $min, $hour, $mday, $mon, $year, $wday) = gmtime($expires);
        $year += 1900;

        return sprintf("%s, %02d-%s-%04d %02d:%02d:%02d GMT",
                       $WDAY[$wday], $mday, $MON[$mon], $year, $hour, $min, $sec);

    }

    return $expires;
}</code></pre>

テストのほうも GMT として時刻を作ってあげないと、環境のタイムゾーン設定によって結果が変わってしまう。今回の場合はテストのほう、時刻を作る部分を `Time::Piece->from_mysql_datetime('2013-01-01 11:11:11')` とし、expected も GMT で n ヶ月後に指定してあげればよい。これで環境依存を無くせる。

Time::Piece のタイムゾーンまわりはよくはまるところで、はまったひとはだいたいこのエントリのお世話になっていることだろう。

[Time::Piece とタイムゾーンの甘い罠 - (ひ)メモ](http://d.hatena.ne.jp/hirose31/20110210/1297341952)

もちろん今回も参照させていただいた。いつもお世話になっております。
