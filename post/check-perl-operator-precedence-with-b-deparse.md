{"title":"perl で演算子の優先順位をしらべる","date":"2013-06-21T20:03:45+09:00","tags":["perl"]}

`perl -MO=Deparse,-p` してあげればよい。

今回 `not defined $a->{foo} && $a->{a}` というコードがあった。意図としては `(not defined $a->{foo}) && $a->{a}` としたかったようだが、実際には `not (defined $a->{foo} && $a->{a})` という挙動になる。つぎのようにすれば簡単にしらべられる。

    % perl -MO=Deparse,-p -le '$a = { a => 1 }; not defined $a->{foo} && $a->{a}; !defined $a->{foo} && $a->{a}'
    BEGIN { $/ = "\n"; $\ = "\n"; }
    ($a = {'a', 1});
    (not (defined($$a{'foo'}) && $$a{'a'}));
    (defined($$a{'foo'}) or $$a{'a'});
    -e syntax OK

`not defined $a->{foo} && $a->{a}` は `(not (defined($$a{'foo'}) && $$a{'a'}))` と解釈され、`!defined $a->{foo} && $a->{a}` は `(defined($$a{'foo'}) or $$a{'a'})` と解釈されていることがわかる。

### 参考

[404 Blog Not Found:perl - B::Deparse](http://blog.livedoor.jp/dankogai/archives/50761629.html)
