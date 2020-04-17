{"title":"perl でモジュールを動的にロードする","date":"2013-05-02T19:35:01+09:00","tags":["perl"]}

`UNIVERSAL::require` をつかうと良いそうだ

<pre><code data-language="javascript">use UNIVERSAL::require;
my $package = "Foo::Bar";
$package->require;</code></pre>

ちょっとあれな方法だと

<pre><code data-language="javascript">my $package = "Foo::Bar";
eval "require $package";</code></pre>

ふつうに `require $package;` とやると require が `$package` の内容にファイルパス (`/foo/bar/Foo/Bar.pm` みたいな文字列) を期待してしまうらしくダメらしい.

`Mouse::Util::load_class` など, モジュールがロードの仕組みを提供しているものもある
