{"title":"ssh port forward 整理","date":"2012-04-04T23:58:30+09:00","tags":["nix"]}

たとえばこんなことができる

    (local [port XXX]) ---- ssh ---- ([port 22] target [port YYY])
    
    (local [port XXX]) ---- ssh ---- ([port 22] remote) -------- ([port YYY] target)

- ローカルの XXX ポートへの接続は, ssh を経由して target の YYY ポートへの接続になる
- 外部から proxy を通して内部ネットワークにアクセスするときの常套手段

コマンドはこういう感じになる

    $ ssh -L 10100:example.com:5000 example.com

上記はローカル 10100 port への接続は example.com への 5000 port に転送される. local - example.com 間は普通に ssh 接続で, example.com から example.com:5000 へ転送される.

`-L` はローカルへの接続をリモートへ転送するオプション. 逆の `-R` というのもあって, リモートのあるポートへの接続をローカルへ転送する.

    $ ssh -L 8080:remote.example.com:80 proxy.example.com

`proxy.example.com` が proxy サーバで, `remote.example.com` は proxy を介さないと接続できないとする. 上記のようにすると, ローカル 8080 への接続は ssh を介して proxy に行き, そこから remote:80 へ転送される.

    $ ssh -C -f -N -L 1234:remote.example.com:1234 remote.example.com

このへんのオプションをつけると便利

- `-C` は転送するデータを圧縮するオプション
- `-f` は ssh 接続をバックグラウンドで実行するオプション
- `-N` はリモートでコマンドを実行しないというオプション. ポートフォワード専用で ssh 接続するときはリモートでコマンドは実行しないので, これを設定しておくと良い (しなかったらどうなるかよくわからない)

### Refs
- [ssh ポートフォワード機能の活用](http://www.turbolinux.co.jp/products/server/11s/user_guide/x9016.html)
- [ssh ポート転送](http://www.geocities.jp/ko_tyche/linux/port.html)
- [SSH ポートフォワーディング](http://www14.plala.or.jp/campus-note/vine_linux/server_ssh/ssh_portforwarding.html)
