{"title":"bash ビルトインコマンドのヘルプを見たい","date":"2012-12-09T18:56:44+09:00","tags":["nix"]}

man だとビルトインコマンド全体の情報が出てどうすべきかと思ってたんだけど, `man bash` の中にビルトインコマンドの説明も入っている. また bash の `help` というビルトインコマンドでも情報が表示される

    $ help help
    help: help [-s] [pattern ...]
        Display helpful information about builtin commands.  If PATTERN is
        specified, gives detailed help on all commands matching PATTERN,
        otherwise a list of the builtins is printed.  The -s option
        restricts the output for each builtin command matching PATTERN to
        a short usage synopsis.
