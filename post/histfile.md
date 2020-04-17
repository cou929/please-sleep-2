{"title":"history の保存先","date":"2012-12-09T18:52:00+09:00","tags":["nix"]}

history コマンドで見られるようなシェルのコマンド履歴はどこに保存されているかというと, `$HISTFILE` 環境変数に定義されているファイルにある

    $ tail -n 5 `echo $HISTFILE`
    : 1354803827:0;history
    : 1354803855:0;which history
    : 1354803891:0;echo $HISTFILE
    : 1354803904:0;cat `echo $HISTFILE`
    : 1354803943:0;tail -n 10 `echo $HISTFILE`
