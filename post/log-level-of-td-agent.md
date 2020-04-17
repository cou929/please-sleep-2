{"title":"td-agent のログレベルを設定する","date":"2014-03-18T23:11:00+09:00","tags":["nix"]}

設定ファイルのデバッグのために一時的に td-agent のログレベルを詳細にしたい。

[fluentd(td-agent)使ってサーバ間でjsonログを転送してみる＆詳細ログも出してみるテスト - ためしにやってみる系](http://d.hatena.ne.jp/tweeeety/20131205/1386252268)

こちらによると、`/etc/init.d/td-agent` を直接編集するという方法が公式サイトで説明されているという。探してみてもその記述は見つけられなかったけれど、一時的に書き換えるだけなら手っ取り早いので、この方法でいくことにした。

バージョンによって場所は違うが、`/etc/init.d/td-agent` 内の、次のようなスクリプトの実行コマンドを定義しているところに

    TD_AGENT_ARGS="${TD_AGENT_ARGS-/usr/sbin/td-agent --group td-agent --log /var/log/td-agent/td-agent.log}"

このように `-vv` フラグを足せば良い。

    TD_AGENT_ARGS="${TD_AGENT_ARGS-/usr/sbin/td-agent -vv --group td-agent --log /var/log/td-agent/td-agent.log}"

`-vv` をフラグによって trace レベルのログも出るようになる。

    $ sudo /etc/init.d/td-agent restart
    $ tail -f /var/log/td-agent/td-agent.log

