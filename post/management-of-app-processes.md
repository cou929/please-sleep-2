{"title":"サーバのプロセスとポート番号がとっちらかっていて困った話","date":"2012-12-03T01:29:03+09:00","tags":["nix"]}

- port 3000, 5000 周辺をすでに誰かが使っていて app を立ちあげられなかったり
- 誰が使っているか調べようと ps しても `node app.js` だけでどいつかわかんなかったり

という困ったことがあったので, ちょっとしたバッチであっても以下のルールで運用しようと思う.

1. 絶対にフルパスで起動する (ps したときわかりやすい)
2. ある程度かたちになってて常駐させるアプリは supervisor or forever 必須
3. /var/app or /var/www に置く
4. アプリログは /var/log/app/name or /var/app/name/foo.log
5. supervisor のログは service/name/log, forever は app_path/forever.log とか
6. 基本 nginx でリバースプロキシする. port は nginx のコンフィグをみればわかるようにする

まだちょっとぶれてるけどサーバ上にいるプロセスを一度に把握できて, どれも同じやり方になっててほしい. capstrano もいれようかなあ. まだ自前の hogehoge.sh で十分ではあるけど.

ちなみに特定のポートを使っているプロセスは lsof とか netstat で調べられる

    $ sudo /usr/sbin/lsof -i:5001
    COMMAND   PID  USER   FD   TYPE  DEVICE SIZE/OFF NODE NAME
    node    27073 kosei    8u  IPv4 2273200      0t0  TCP *:commplex-link (LISTEN)
    $ netstat -anp | grep 5001
    (Not all processes could be identified, non-owned process info
     will not be shown, you would have to be root to see it all.)
    tcp        0      0 0.0.0.0:5001                0.0.0.0:*                   LISTEN      27073/node
