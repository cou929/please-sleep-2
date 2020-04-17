{"title":"perl の do 関数","date":"2012-10-06T13:05:15+09:00","tags":["perl"]}

### do BLOCK

- そのブロックで最終的に評価された値を返す
- ようは関数の戻り値と一緒
- その場で無名関数を作って即実行するような意味になるようだ

        $ perl -le 'print sub { 'foo' }->()'
        foo
        $ perl -le 'print do { 'foo' }'
        foo

- 使い分けはよくわからない

### do EXPR

- `EXPR` のファイルを読み込んで eval する
- 以下は同じ意味らしい

        do 'stat.pl';
        
        # is just like
        
        eval `cat stat.pl`;

- わりと昔ながらの手法らしい
  - [モジュールの読み込みの仕組みを理解する - サンプルコードによるPerl入門](http://d.hatena.ne.jp/perlcodesample/20090208/1232890021)
- ちょっとしたツールを書く場合には便利そう

### 見かけた例

コンフィグファイルを読み込むモジュール

    use strict;
    use warnings;
    use File::Basename qw(dirname);
    use Exporter qw(import);
    our @EXPORT_OK = qw(config);
    
    {
        my $config;
        sub config() {
            $config //= do {
                my $script_dir = dirname $0;
                do "$script_dir/config.pl";
            };
        }
    }
