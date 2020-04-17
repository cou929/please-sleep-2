{"title":"NYTProf でプロファイリング","date":"2014-03-20T11:47:36+09:00","tags":["perl"]}

実行時に `-d:NYTProf` オプションを追加してあげれば良い。

plackup などなにかコマンドをかませている場合は、こういう感じにすればいいはず。

    carton exec -- perl -d:NYTProf -S plackup app.psgi

この場合 NYTProf の設定は `NYTPROF` 環境変数で渡す。例えば、

    NYTPROF='file=/tmp/nytprof.out'

とすれば、結果の出力先を変えられる。

またオプションではなくコード中でやるには、`DB::enable_profile`、`DB::disable_profile` を使えば良い。

plack アプリの場合、psgi ファイルなどで、次のようにしてあげればよい。

<pre><code data-language="perl">use Devel::NYTProf;

my $app = sub {
    my $env = shift;
    DB::enable_profile('/tmp/nytprof.out');
    my $req = Plack::Request($env);
    my $app = YourApp->new(req => $req);
    my $res = $app->run;
    DB::disable_profile();
    return $res;
};</code></pre>
