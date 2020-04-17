{"title":"perl で関数にハッシュを渡す","date":"2012-10-06T13:06:02+09:00","tags":["perl"]}

- 配列の引数は全部 `@_` にリストで入っている
- ハッシュの場合も `key, val, key, val ...` と入っているだけ
- `%hash` に `@_` を渡すとよしなにストアしてくれるだけ
  - 関数の引数の場合にかぎらず perl でのハッシュ一般に言えることだけど

        $ perl -le 'sub test { print join ", ", @_;  } test(1); test(qw/foo bar baz/); test([qw/foo bar baz/]); test(foo => 1, bar => 2, baz => 3); test( { foo => 1, bar => 2, baz => 3 } );'
        1
        foo, bar, baz
        ARRAY(0x1800a220)
        foo, 1, bar, 2, baz, 3
        HASH(0x1801c918)
        
        $ perl -MData::Dumper -le 'sub test { my %h = @_; print Dumper \%h;  } test(foo => 1, bar => 2, baz => 3);'
        $VAR1 = {
                  'bar' => 2,
                  'baz' => 3,
                  'foo' => 1
                };

- ハッシュで受け取るようにしておくと, 名前付きで引数を渡せて便利. python っぽい

        package Foo::Bar
        
        sub new {
            my ($class, %params) = @_;
            my $basedir = delete $params{basedir} or die "no basedir";
        
            bless {basedir => $basedir}, $class;
        }
        
        # こうやって使う
        my $instance = Foo::Bar->new(basedir => '/tmp/');
