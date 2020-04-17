{"title":"perl でファイルの内容を一度に読み込む","date":"2013-12-20T21:58:46+09:00","tags":["perl"]}

[Perlでファイルの一括読込 : D-7 <altijd in beweging>](http://lestrrat.ldblog.jp/archives/23209249.html) より、以下のようにするのが簡単だった。

<pre><code data-language="perl">my $content = do { local $/; <$fh> };</code></pre>

ワンライナーで使えそう。

で、`local $/;` の意味がわからなかったので perldoc perlvar した。曰く、`$/` はファイルを読む際の一行の区切りを指定するための特殊変数らしい。デフォルトは "\n"。これに undef を設定すると行の区切りはなく、EOF まで一気に読んでくれる。

    HANDLE->input_record_separator( EXPR )
    $INPUT_RECORD_SEPARATOR
    $RS
    $/      The input record separator, newline by default.  This
            influences Perl's idea of what a "line" is.  Works like awk's
            RS variable, including treating empty lines as a terminator if
            set to the null string (an empty line cannot contain any spaces
            or tabs).  You may set it to a multi-character string to match
            a multi-character terminator, or to "undef" to read through the
            end of file.  Setting it to "\n\n" means something slightly
            different than setting to "", if the file contains consecutive
            empty lines.  Setting to "" will treat two or more consecutive
            empty lines as a single empty line.  Setting to "\n\n" will
            blindly assume that the next input character belongs to the
            next paragraph, even if it's a newline.

                local $/;           # enable "slurp" mode
                local $_ = <FH>;    # whole file now here
                s/\n[ \t]+/ /g;

今回は do ブロックの中で local で指定しているので、ブロックの中でだけ一時的に変更し、ファイルを一気に読み込んでいる。

