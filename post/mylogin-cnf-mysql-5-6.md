{"title":"MySQL5.6 の .mylogin.cnf","date":"2013-11-01T19:09:25+09:00","tags":["nix"]}

MySQL 5.6 からは .mylogin.cnf に mysql へのログイン情報を暗号化して保持することができるようになったらしい。いままで my.cnf に平文で置いていたのでそれより多少ましになりそうだ。

`mysql_config_editor` というツールがついてくるので、.mylogin.cnf の編集にはこれを使えばよい。

    $ mysql_config_editor set --user=foobar --password

こうするとプロンプトでパスワードの入力が求められる。入力するとホームディレクトリに .mylogin.cnf ができているはずだ。

設定内容の出力のために print というコマンドもある。

    mysql_config_editor print --all
    [client]
    user = foobar
    password = *****
