{"title":"ハイパフォーマンスブラウザネットワーキング を読んだ","date":"2021-01-14T01:18:00+09:00","tags":["book"]}

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873116767/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://images-na.ssl-images-amazon.com/images/I/51x2sA8N+TL._SX389_BO1,204,203,200_.jpg" alt="ハイパフォーマンス ブラウザネットワーキング ―ネットワークアプリケーションのためのパフォーマンス最適化" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873116767/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">ハイパフォーマンス ブラウザネットワーキング ―ネットワークアプリケーションのためのパフォーマンス最適化</a></div><div class="amazlet-detail">Ilya Grigorik (著), 和田 祐一郎  (翻訳), 株式会社プログラミングシステム社 (翻訳)<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873116767/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

TCP、UDP からはじまり、TLS、HTTP2、WebSocket、WebRTC まで、各プロトコルの詳解と、その上でアプリケーションを作る際の最適化方針について扱っている本。

各論のテクニックではなく、プロトコルの仕組みと、それを効率的に利用するための考え方を説明している。

そのため Web アプリケーションをチューニングしたい人だけでなく、上記のプロトコルスタックの内容を理解したい人におすすめだと思う。自分はまさに後者の目的だった。

原著は 2013 とちょっと古いが、その時点の最新状況までよく説明されている。時期的に HTTP3 などはカバーされていないが、2013 年時点までの動向をこの本でキャッチアップして、差分は別の方法でカバーするのが良さそう。

余談だが原著は Web に公開されている。

[High Performance Browser Networking \(O'Reilly\)](https://hpbn.co/)

以下読書メモ。

## レイテンシ・帯域幅入門

- レイテンシと帯域幅
    - バンドウィズは理論値でスループットは実効値
- レイテンシ
    - 高速が現在の物理的な上限
    - 光ファイバーでは6割程度まで出ている
    - どれだけ頑張っても今の二倍以上にはならないくらいのオーダー
    - cdn やプリフェッチといった工夫
- 帯域幅
    - 経路上の最も細いところに律速
    - 理論上はレイテンシよりも上限は高い
        - 機器を増強すれば帯域幅を増やせる
    - 動画視聴など大容量転送で重要度が高い
- 帯域幅の測り方
    - 程よい大きさのファイルの転送にかかった時間を測る
        - http://www.math.kobe-u.ac.jp/HOME/kodama/tips-net-speed.html

## TCPの構成要素

- 2way handshake
    - [RFC 7413 \- TCP Fast Open](https://tools.ietf.org/html/rfc7413)
        - syn でデータ転送開始できるように
        - 初回のコネクション確立時には不可能なのと、サイズ上限があるのが制約
- 輻輳
    - [RFC 896 \- Congestion Control in IP/TCP Internetworks](https://tools.ietf.org/html/rfc896)
        - 輻輳についての報告
        - ネットワーク全体が低パフォーマンスな状態に平衡する特徴がある
    - 対策
        - フロー制御
            - 忙しい場合はウィンドウサイズ (rwnd) を伝えて手加減してもらう
            - サイズのデフォルト値、上限がのちに低くなりすぎて拡張された
                - [RFC 1323 \- TCP Extensions for High Performance](https://tools.ietf.org/html/rfc1323)
        - スロースタート
            - フロー制御はホスト内の処理能力しか考慮せず、ネットワークの状況が考慮されていない
                - ホストは余裕があってもホスト間の帯域が逼迫している場合、フロー制御だけでは全力で通信してしまう
            - 輻輳ウィンドウサイズ (`cwnd`) を小さい値からはじめ、パケットの往復ごとに大きくしていく
                - 指数関数的にサイズを増やしていき、パケロスが発生したら止める
                - `min(rwnd, cwnd)` が使われる
                - cwnd の最適サイズ N に到達する時間は `rtt * log2(N/cwnd初期値)`
                    - rtt 56ms (ロンドンニューヨーク間相当) で cwnd 初期値 4 セグメント (古い値)、rnwd 64kb とすると、ウィンドウサイズが 64kb に到達するには 224ms かかる
            - ストリーミングなど継続的に通信するユースケースでは影響が小さい
                - バースト的に短命のリクエストを複数箇所とやり取りするユースケースでは影響が比較的大きくなる
            - 初期値を大きくする、rtt を小さくするといった回避方法がある
    - Slow Start Restart (SSR)
        - 一定期間アイドルだった接続は cwnd をリセットする
            - その間に経路の状況が変わっている可能性があるので
        - keepalive 接続のように、長期間維持しアクティブとアイドルを交互に繰り返すようなユースケースではパフォーマンスに悪影響なため、無効化が推奨される
    - rtt が大きいほど影響が大きい
        - ある rtt 50ms のケースではアプリケーションが 40ms で返していたとしても、初回接続では 250ms ほどかかる
            - 接続を使い回すと 100ms ほどに低減できる
    - 輻輳回避
        - パケロスが発生した際にウィンドウサイズを減らす仕組み
        - AIMD
            - 半分にして徐々に増やしていく
            - 多くのケースでは保守的
        - Proportional Rate Reduction
            - [RFC 6937 \- Proportional Rate Reduction for TCP](https://tools.ietf.org/html/rfc6937)
            - kernel 3.2 で導入
    - 最適なウィンドウサイズは rtt によって決まる
        - rtt の時間を埋められる程度のウィンドウサイズが無いともったいない
        - 小さすぎるウィンドウサイズの場合、rtt の待ち時間が大きくなる
        - 帯域を使い切れていないようなケースでは、ウィンドウサイズが小さいことがありうる
        - 多くのケースでは帯域幅ではなくレイテンシが TCP のボトルネックになる
    - Head of Line (HoL) ブロッキング
        - TCP はデータ順序を保証し、それを緩和することはできない
        - 途中のデータをロストした場合、その再送が終わるまでは待ちになり、バッファに受信済みデータが溜まった状態になる
        - アプリケーションからは時折ジッターが発生しているように見え、かつその原因は TCP が隠蔽しているからわからない (単なるレイテンシの増加に見える)
    - TCP のチューニング
        - TCP 自体はアップデートが繰り返され複雑化しているが、要点はかわらない
            - 3way handshake は 1 往復分のレイテンシを発生させる
            - スロースタートはすべての新規接続に適用される
            - フロー制御と輻輳制御はすべての接続のスループットを制御する
            - TCP のスループットは cwnd によって制限される
            - 結果としてほとんどのケースでは帯域幅ではなくレイテンシがボトルネックになる
        - チューニング
            - 前提としてカーネルをバージョンアップするだけで改善することが多い
            - 初期ウィンドウサイズの増加
                - 短命でバースト的な接続が多い場合に有効
            - スロースタートリスタートの無効化
                - 持続的で間隔をおいてバーストする接続のパフォーマンス向上
            - ウィンドウスケーリング (ウィンドウサイズ最大値の増加)
                - 高レイテンシの接続の改善度が大きい
            - TCP Fast Open
                - 比較的サーバクライアント両方がサポートする必要があるが効果的
        - アプリケーションレベルのチューニング
            - 不要なパケットを送信しない
            - 物理的距離を近くする
            - 接続を再利用する
- [ss\(8\) \- Linux manual page](https://man7.org/linux/man-pages/man8/ss.8.html)
    - another utility to investigate sockets
- tcpdump でのウィンドウサイズの見方
    - `win` がウィンドウサイズ
    - `wscale` がウィンドウスケール
        - win を `2 ^ wscale` 倍したものが実際に使われる
    - 参考
        - [tcpdumpコマンドの使い方: UNIX/Linuxの部屋](http://x68000.q-e-d.net/~68user/unix/pickup?tcpdump)
        - [How to determine TCP initial window size and scaling option? Which factors affect the determination? \- Red Hat Customer Portal](https://access.redhat.com/solutions/29455)
    - 実際の例
        - クライアントからサーバへの win は `29200` から `65535` まで上昇している
        - wscale は 7

<div></div>

```
vagrant@ubuntu-xenial:~$ sudo tcpdump port 443 -l | tee log
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
listening on enp0s3, link-type EN10MB (Ethernet), capture size 262144 bytes

# 別タブで

vagrant@ubuntu-xenial:~$ curl -v https://please-sleep.cou929.nu/

# 結果

12:53:26.724681 IP 10.0.2.15.33692 > server-99-84-130-117.nrt57.r.cloudfront.net.https: Flags [S], seq 273873366, win 29200, options [mss 1460,sackOK,TS val 5634728 ecr 0,nop,wscale 7], length 0
12:53:26.737876 IP server-99-84-130-117.nrt57.r.cloudfront.net.https > 10.0.2.15.33692: Flags [S.], seq 2859584001, ack 273873367, win 65535, options [mss 1460], length 0
12:53:26.737910 IP 10.0.2.15.33692 > server-99-84-130-117.nrt57.r.cloudfront.net.https: Flags [.], ack 1, win 29200, length 0
12:53:26.773433 IP 10.0.2.15.33692 > server-99-84-130-117.nrt57.r.cloudfront.net.https: Flags [P.], seq 1:247, ack 1, win 29200, length 246
12:53:26.773522 IP server-99-84-130-117.nrt57.r.cloudfront.net.https > 10.0.2.15.33692: Flags [.], ack 247, win 65535, length 0
12:53:26.798057 IP server-99-84-130-117.nrt57.r.cloudfront.net.https > 10.0.2.15.33692: Flags [.], seq 1:2841, ack 247, win 65535, length 2840
12:53:26.798079 IP 10.0.2.15.33692 > server-99-84-130-117.nrt57.r.cloudfront.net.https: Flags [.], ack 2841, win 34080, length 0
12:53:26.798239 IP server-99-84-130-117.nrt57.r.cloudfront.net.https > 10.0.2.15.33692: Flags [P.], seq 2841:4207, ack 247, win 65535, length 1366
12:53:26.798247 IP 10.0.2.15.33692 > server-99-84-130-117.nrt57.r.cloudfront.net.https: Flags [.], ack 4207, win 36920, length 0
12:53:26.798466 IP server-99-84-130-117.nrt57.r.cloudfront.net.https > 10.0.2.15.33692: Flags [P.], seq 4207:5285, ack 247, win 65535, length 1078
12:53:26.798475 IP 10.0.2.15.33692 > server-99-84-130-117.nrt57.r.cloudfront.net.https: Flags [.], ack 5285, win 39760, length 0
...
12:53:27.061258 IP 10.0.2.15.33692 > server-99-84-130-117.nrt57.r.cloudfront.net.https: Flags [.], ack 108567, win 65535, length 0
12:53:27.061270 IP server-99-84-130-117.nrt57.r.cloudfront.net.https > 10.0.2.15.33692: Flags [P.], seq 108567:109969, ack 488, win 65535, length 1402
12:53:27.061466 IP server-99-84-130-117.nrt57.r.cloudfront.net.https > 10.0.2.15.33692: Flags [P.], seq 109969:111850, ack 488, win 65535, length 1881
12:53:27.061476 IP 10.0.2.15.33692 > server-99-84-130-117.nrt57.r.cloudfront.net.https: Flags [.], ack 111850, win 65535, length 0
12:53:27.061822 IP 10.0.2.15.33692 > server-99-84-130-117.nrt57.r.cloudfront.net.https: Flags [P.], seq 488:519, ack 111850, win 65535, length 31
12:53:27.061924 IP server-99-84-130-117.nrt57.r.cloudfront.net.https > 10.0.2.15.33692: Flags [.], ack 519, win 65535, length 0
12:53:27.067261 IP 10.0.2.15.33692 > server-99-84-130-117.nrt57.r.cloudfront.net.https: Flags [F.], seq 519, ack 111850, win 65535, length 0
12:53:27.067419 IP server-99-84-130-117.nrt57.r.cloudfront.net.https > 10.0.2.15.33692: Flags [.], ack 520, win 65535, length 0
12:53:27.079538 IP server-99-84-130-117.nrt57.r.cloudfront.net.https > 10.0.2.15.33692: Flags [F.], seq 111850, ack 520, win 65535, length 0
12:53:27.079561 IP 10.0.2.15.33692 > server-99-84-130-117.nrt57.r.cloudfront.net.https: Flags [.], ack 111851, win 16824, length 0

# tcp_window_scaling は有効化されている
vagrant@ubuntu-xenial:~$ sudo sysctl -a | grep tcp_window_scaling
sysctl: net.ipv4.tcp_window_scaling = 1
```


- SSR の状態をみてみる
    - 雑に起動した vagrant の `ubuntu/xenial64` の初期値は有効だった


```
vagrant@ubuntu-xenial:~$ sudo sysctl -a | grep tcp_slow_start
net.ipv4.tcp_slow_start_after_idle = 1
```

## UDPの構成要素

- 有名なユースケース
    - dns, webrtc
- datagram
    - パケットはフォーマットされたデータブロック全般をいう
    - データグラムはその中でも信頼性が低いサービスによって配信されたもの
- udp
    - 送信元ポート番号、宛先ポート番号、パケット長、チェックサムのみ
    - 送信するだけで、配信保証、順序保証、状態追跡、輻輳制御はない
    - TCP がバイトストリーム指向なのに対して UDP は固定長
    - ステートレス
- udp と nat の問題
    - net テーブルのエントリをいつ消せばよいかわからない
        - udp には状態が無いのでいつ終わったのかわからない
        - タイムアウト + keepalive
    - nat トラバーサル
        - nat 内のノードへのインバウンドトラフィックは、まだ nat テーブルにエントリが無いことがほぼなのでほぼ確実にフィルタされてしまう
            - p2p のように両ホストがサーバにもクライアントにもなるケースで問題
        - nat 内のホスト上のアプリケーションがアプリケーション層に自分のグローバル ip 、ポート番号を入れたいが、それがわからない
            - p2p のアプリケーションだと必要ということだと思われる
        - STUN, TURN, ICE
            - STUN
                - nat デバイスのグローバル ip & port を外のサーバにお互い登録してから通信を始める
                    - 登録時に nat テーブルにエントリもできる
                    - アプリケーション層への ip, port 書き込みは外のサーバを見れば良い
            - TURN
                - 外の共通のリレーサーバにデータを送ってリレーしてもらう
                - もはや p2p ではない
            - ICE
                - 使える手段をフォールバックしながら決める
        - 翻訳のせいか本書の書き方だと udp で発生する問題で tcp では問題ないように読めたが、そうではないはず...?
            - [NAT traversal \- Wikipedia](https://ja.wikipedia.org/wiki/NAT_traversal)
            - [NAT越え（NATトラバーサル\-NAT traversal）｜Web会議・テレビ会議システムならLiveOn（ライブオン）](https://www.liveon.ne.jp/glossary/w/nat_traversal.html)

## TLS

- セッション層で実現 (TCP の上、アプリの下)
- 暗号化、認証、データ整合性
- TLS handshake
    - 暗号化スイート合意、鍵準備など
- 中間装置の影響を割けるため 443 のトンネルを確立して通信する用途もある
- 最適化
    - early termination
        - cdn のように edge を配置するが、edge では tls 終端のみをする
        - edge から origin は接続を使い回しやすいし、cdn のバックボーンのほうが通常のインターネットよりも経路が最適化されているので、これだけ (コンテンツをキャッシュせずに tls 終端だけをする) でもメリットが有る
    - sticky session でセッション再開
        - 過去のハンドシェイクで合意した内容をできるだけ使い回す
    - tls false start
        - ハンドシェイク中に先んじてデータ送信
    - データサイズの調整
    - tls データ圧縮の無効化
        - 脆弱性があるらしい
        - セッション層なので、データの中身を見ずに圧縮するため無駄が多い (画像を再圧縮するなど)
    - 証明書チェーンを短く
    - OSCP Stapling
    - HSTS

openssl コマンドを試してみる。

- このブログは cloudfront で配信している
- ルート認証局は AWS Root CA
- 証明書は https://www.amazontrust.com/repository/AmazonRootCA1.pem にある
- `-servername` オプションで SNI をしていしてあげないといけない (たぶん cloudfront は複数ドメインの配信をしているから)
    - [SSL接続時にHandshakeに失敗する場合はSNIが原因かもしれない \- TODESKING](http://www.todesking.com/blog/2017-02-11-openssl-alert-handshake-failure/)

<div></div>

```
vagrant@ubuntu-xenial:~$ openssl s_client -state -CAfile AmazonRootCA1.pem -connect please-sleep.cou929.nu:443 -servername please-sleep.cou929.nu < /dev/null

CONNECTED(00000003)
SSL_connect:before/connect initialization
SSL_connect:SSLv2/v3 write client hello A
SSL_connect:unknown state
depth=4 C = US, O = "Starfield Technologies, Inc.", OU = Starfield Class 2 Certification Authority
verify return:1
depth=3 C = US, ST = Arizona, L = Scottsdale, O = "Starfield Technologies, Inc.", CN = Starfield Services Root Certificate Authority - G2
verify return:1
depth=2 C = US, O = Amazon, CN = Amazon Root CA 1
verify return:1
depth=1 C = US, O = Amazon, OU = Server CA 1B, CN = Amazon
verify return:1
depth=0 CN = *.cou929.nu
verify return:1
SSL_connect:unknown state
SSL_connect:unknown state
SSL_connect:unknown state
SSL_connect:unknown state
SSL_connect:unknown state
SSL_connect:unknown state
SSL_connect:unknown state
SSL_connect:unknown state
SSL_connect:unknown state
---
Certificate chain
 0 s:/CN=*.cou929.nu
   i:/C=US/O=Amazon/OU=Server CA 1B/CN=Amazon
 1 s:/C=US/O=Amazon/OU=Server CA 1B/CN=Amazon
   i:/C=US/O=Amazon/CN=Amazon Root CA 1
 2 s:/C=US/O=Amazon/CN=Amazon Root CA 1
   i:/C=US/ST=Arizona/L=Scottsdale/O=Starfield Technologies, Inc./CN=Starfield Services Root Certificate Authority - G2
 3 s:/C=US/ST=Arizona/L=Scottsdale/O=Starfield Technologies, Inc./CN=Starfield Services Root Certificate Authority - G2
   i:/C=US/O=Starfield Technologies, Inc./OU=Starfield Class 2 Certification Authority
---
Server certificate
-----BEGIN CERTIFICATE-----
MIIFXjCCBEagAwIBAgIQDaMOfYjNhwp/4Lh9uDVrfDANBgkqhkiG9w0BAQsFADBG
...
iqoq+5uEYsbulNNpu40+BNxm/tdyh3V0LlqR2doznD74QQ==
-----END CERTIFICATE-----
subject=/CN=*.cou929.nu
issuer=/C=US/O=Amazon/OU=Server CA 1B/CN=Amazon
---
No client certificate CA names sent
Peer signing digest: SHA256
Server Temp Key: ECDH, P-256, 256 bits
---
SSL handshake has read 5439 bytes and written 422 bytes
---
New, TLSv1/SSLv3, Cipher is ECDHE-RSA-AES128-GCM-SHA256
Server public key is 2048 bit
Secure Renegotiation IS supported
Compression: NONE
Expansion: NONE
No ALPN negotiated
SSL-Session:
    Protocol  : TLSv1.2
    Cipher    : ECDHE-RSA-AES128-GCM-SHA256
    Session-ID: 17B59D48D96501E1C2E8F230CD5B73ADAE66A184C198D622407953328A8A40D9
    Session-ID-ctx:
    Master-Key: 9F565875EA74D6CB150FE161268D202C97AA8C5DF235737FDD62FB872EBC4E6681BDAB6DDEE475A54C0F02A69AD365A0
    Key-Arg   : None
    PSK identity: None
    PSK identity hint: None
    SRP username: None
    TLS session ticket lifetime hint: 216000 (seconds)
    TLS session ticket:
    0000 - 31 36 30 39 37 38 39 37-36 36 30 30 30 00 00 00   1609789766000...
    0010 - 8f 4f b2 65 c5 52 76 93-da 1f 7b fc 3f c5 c2 19   .O.e.Rv...{.?...
    0020 - 4c b2 55 5e 7a 6e 76 25-e6 d7 71 c9 aa 81 96 50   L.U^znv%..q....P
    0030 - 7d db 8b 8a d9 7b 4c d8-84 cc 24 3b e1 a7 ff 68   }....{L...$;...h
    0040 - ca fc 14 34 47 b7 91 74-80 08 4d e6 2f 18 b4 76   ...4G..t..M./..v
    0050 - f1 ef 01 f3 ae b4 c5 60-c7 e8 1b 10 6d 77 63 43   .......`....mwcC
    0060 - 91 3a ba 8a a4 0f 73 bf-                          .:....s.

    Start Time: 1609859506
    Timeout   : 300 (sec)
    Verify return code: 0 (ok)
---
DONE
SSL3 alert write:warning:close notify
```

ちなみに curl で接続すると `/etc/ssl/certs/ca-certificates.crt` のルート証明書を見ているっぽい。

```
vagrant@ubuntu-xenial:~$ curl -v https://please-sleep.cou929.nu | head
* Connected to please-sleep.cou929.nu (99.84.130.44) port 443 (#0)
* found 138 certificates in /etc/ssl/certs/ca-certificates.crt
* found 552 certificates in /etc/ssl/certs
* ALPN, offering http/1.1
* SSL connection using TLS1.2 / ECDHE_RSA_AES_128_GCM_SHA256
*        server certificate verification OK
*        server certificate status verification SKIPPED
*        common name: *.cou929.nu (matched)
*        server certificate expiration date OK
*        server certificate activation date OK
*        certificate public key: RSA
*        certificate version: #3
*        subject: CN=*.cou929.nu
*        start date: Tue, 21 Apr 2020 00:00:00 GMT
*        expire date: Fri, 21 May 2021 12:00:00 GMT
*        issuer: C=US,O=Amazon,OU=Server CA 1B,CN=Amazon
*        compression: NULL
```

そこには Amazon Root CA のものも含まれていた

```
vagrant@ubuntu-xenial:~$ grep MIIDQTCCAimgAwIBAgITBmyfz5m /etc/ssl/certs/ca-certificates.crt
MIIDQTCCAimgAwIBAgITBmyfz5m/jAo54vB4ikPmljZbyjANBgkqhkiG9w0BAQsF
```

## ワイヤレスネットワーク入門

- シャノン・ハートレーの法則
    - `C = BW x log2(1 + S/N)`
    - C = チャネル容量 = 帯域幅
    - つまり、チャネル容量は周波数帯域幅が大きくなるか、信号強度が強くなるとあがる
- 信号は周波数が低いほど、遠くに届くがデータ量が少ない
    - wifi の 2.5gh と 5gh もそうかも
- 信号強度 (S/N 比)
    - 無線は必ず干渉をうける
        - 周波数帯域は不足しているうえに、デバイスの数が多い
    - S を強めるか、距離を近づけるかしかない
    - 遠近問題
        - 強い信号を受信することにより、弱い信号を受信できなくなる問題
    - セルブリージング
        - N が増えることでカバーできる範囲が狭まる問題
- 変調 (modulation)
    - デジタル - アナログの変換
- まとめ
    - どんな無線通信の仕様も以下の原理に従う
    - 無線通信の帯域幅は割り当てられた帯域幅の大きさと信号強度に依存する
    - 共有媒体 (電波) を利用して通信する
    - 特定の周波数帯に制限されている
    - 出力が制限されている
    - 常に変動するバックグラウンドノイズと干渉に依存する
    - 選択されたワイヤレステクノロジーの技術的制約に依存する
    - デバイスの形状、出力などに依存する

## モバイルネットワークの最適化

- 消費電力
    - モバイルネットワークの場合、無線の利用を最低限にしたほうがベター
    - wifi とモバイルネットワークでは要件が異なる
    - 通信頻度を最低限にしたほうがよい
        - ポーリングよりプッシュのほうがベター
        - リクエストはまとめるほうがよい
            - Nagle アルゴリズム
            - サードパーティのプッシュサービス (メッセージまとめ、オンラインの時まで待つみたいな)
            - W3C Push API
    - tcp, udp の接続状態とデバイスの無線状態は無関係
        - 新しいパケットが外部から送られた際に、事業者の無線ネットワークがデバイスに通知し、接続状態に遷移させ、データ送信を再開する
        - tcp より上でキープアライブをする必要はない (不要なキープアライブがバッテリーを無駄に消費する悪影響がある)
- RCC (Radio Resource Controller)
    - デバイスと無線基地局との接続管理
    - ワイヤレスネットワークごとに異なるステートマシン
    - データ送受信が必要な際に高出力状態に遷移する
    - タイムアウトで省電力状態へ遷移する
    - 状態遷移に 10-100ms から数秒まで
    - アプリケーションはユーザーインタラクションやフィードバックとネットワークコミュニケーションを切り分けて設計する
- 継続的な通信よりもバースト的なほうが有利な設計になっている
    - プリフェッチ、ローカル保持
- Wifi へのオフロード

## Webパフォーマンス入門

- 動画ストリーミングなどでない限りは帯域幅よりレイテンシがボトルネック

## HTTP 1.x

- デフォルトで keepalive
    - `Connection: close` で切断
- パイプライン
    - サポート不足、HoL ブロッキングにより実質的に失敗
    - リクエストをまとめて送り、サーバは処理完了後クライアントの返答を待たずに順次結果を返していく
    - http1.1 はシーケンシャルな通信しかできない (並行、多重化した通信ができない) ので HoL ブロッキング
    - 失敗時の複雑さもありほぼ実装されていない
    - パイプライン自体は有効な手法、サーバとクライアントを両方コントロールできる場面では検討の余地あり
        - itunes のチューニング事例
- 並列 TCP 接続、ドメインシャーディング
    - http1.1 のシーケンシャル性制限の回避のため、クライアントは 6 並列まで接続プールを管理しているが、実装が複雑
    - さらに最適化するためサブドメインをシャーディングするテクニックもある
        - 管理コストは高まる
- js 結合、css スプライト
    - 接続数は減るが、デメリットも
        - そのページには関係ないリソースも含まれる
        - 中身の一部の変更だけでも全体が更新される (キャッシュがクリアされる)
        - js, css は html と違い読み込みが全て終わってからパースされるので、アプリケーションが遅くなる恐れがある
- リソースインライン化
    - data uri に画像などのバイナリを埋め込んでしまう
    - リソースが小さくて汎用的でなければ検討の余地あり

## HTTP 2.0

- HTTP1x とは実装上の互換性はない
    - アプリケーションからは透過的だが、サーバクライアント中間装置はそれぞれ 2 への対応が必要
- テキストではなくバイナリで通信
    - 固定長だが従来に比べると大幅サイズ減 + パースが容易に
- 1 コネクション内での多重化
- リクエストの優先順位付け
- サーバプッシュ
    - data 属性などでのインライン化が不要に
- ヘッダ圧縮
    - 送信済みで変化がないヘッダは送信しなくてもよく、ある意味ステートフルに

## アプリケーション配信最適化

- 常に有効な施策
    - ラウンドトリップの数を減らす
        - DNS ルックアップを減らす
        - TCP 接続の再利用
        - HTTP リダイレクトを減らす
    - 通信時のデータ量を減らす
        - 転送中のリソースの圧縮
        - cookie などリクエストから不要なデータを排除
    - そもそも通信を発生させない
        - 必要のないリソースを削除する
        - リソースをクライアントにキャッシュ
            - `Cache-Control` で生存期間指定、`Last-Modified` `ETag` は検証方法の提供
    - その他
        - リクエストとレスポンス処理の並列化
            - keepalive を有効化した上で HTTP1.1 並列接続の利用、または HTTP2 の利用
        - プロトコル特有の最適化を適用 (HTTP1x, 2)
- HTTP2 の最適化
    - cnwd は 10 セグメントから開始
    - TLS と ALPN ネゴシエーションをサポート
    - TLS セッション再開をサポート
        - ハンドシェイクのレイテンシを最小化
    - HTTP1x 最適化を削除して HTTP2 のアルゴリズムの性能を出し切る
        - ファイル結合、スプライト
        - インライン化を排除してサーバプッシュを利用
- LB
    - LB ではなくアプリケーションサーバで TLS 終端させる場合、そのアプリサーバが HTTP2 に対応していれば良い
    - LB で TLS 終端させる場合、その LB は ALPN ハンドシェイクにも対応し、後ろのサーバに適切にプロキシする必要がある
        - 後ろに HTTP1x 歯科対応できないサーバが合った場合の振り分けなど
        - 後ろの HTTP2 サーバは TLS なしの HTTP2 もサポート必要 (それか LB とアプリサーバも TLS + ALPN で通信する)

## ブラウザネットワーク入門

- ブラウザが行う接続管理
    - オリジンごとにソケットのプールを持つ
        - ページ単位ではない
    - こうすることで、ページを超えた最適化をブラウザがすることができる
        - ブラウザはキューのリクエストを優先度に応じて処理できる
        - ブラウザはソケットを再利用できる
        - ブラウザはリクエストの発生を予測して先にソケット接続を開くことができる
        - ブラウザはアイドル状態のソケットを切断時に最適化できる
            - ソケット終了時に RST を送る。モバイルデバイスの無線がアイドル状態だった場合、アクティブにするので電力を消費する。別の通信が発生するまで RST の送信を待機させるという最適化が可能
        - ブラウザはすべてのソケットの帯域幅割り当てを横断的に最適化できる
- サンドボックス化
    - 接続数の制限
    - リクエスト整形とレスポンス処理
        - 変なリクエストを送らない
    - TLS ネゴシエーション
        - 証明書の検証
    - 同一生成元ポリシー
- キャッシュ
- 各種 API を提供
    - XHR, Server-Sent Events, WebSocket
    - WebRTC

## XMLHttpRequest

- CORS
    - クライアントは Origin ヘッダを送信
    - サーバは Access-Controll-Allow-Origin を返す
        - リクエストの Origin ヘッダを検証して、リクエストを受け付ける場合に返す
    - ブラウザが行うセキュリティ予防策
        - CORS リクエストはクッキーや HTTP 認証などのユーザーデータを取り除く
            - クライアントは withCredentials プロパティをセットして、サーバは Access-Control-Allow-Credentials を含めたレスポンスを返すという追加の準備が必要
            - preflight request でクライアントがサーバに事前確認する
        - simple cross-origin request (GET, POST, HEAD と XHR が扱うことができるヘッダだけに制限)
- ストリーミング対応の不足
    - 接続をつなぎぱなしにしておいて、サーバからストリームとしてデータを受け取るのが厳しい
        - 受け取るデータの管理 (オフセット追跡、データの結合、失敗時のフォローなど) をアプリケーションが実装しないといけない上にブラウザ間の差異の対策も必要
        - ブラウザによって逐次読み込むことができるコンテンツタイプが異なる
    - ポーリング
        - 無駄が多い
    - ロングポーリング
        - サーバはクライアントからのリクエストに即座にレスポンスせず、返すべきものができたときにだけレスポンスする
        - 定期的にレスポンスすべきものが発生する場合、通常のポーリングでまとめて返したほうが効率的だが、それ以外では有効な戦略
        - 初期の fb チャットも使用していたらしい

## Server-Sent Events

- サーバからの push を可能にする
- サーバからはテキストベースでデータを送信
    - id, event 名がオプショナル

```
data: xxx

---

event: foo
data: yyy

---

id: 99
event: bar
data: zzz
```

- クライアントはデータの受取時すべてや、特定のイベントに対してコールバックを設定可能
- 接続が切れ再接続した場合、ブラウザは `Last-Event-ID` を送信できる
    - サーバは必要があればその後の ID から resume すればよい
    - 投げっぱなしでいいなら id 無しで送り続ければよい
- サーバからの一方的な送信しかできない、utf-8 テキストをやり取りするオーバーヘッドというデメリット
    - その分シンプルなのと、通信方法のフォールバックの選択肢としての存在価値

## WebSocket

- 接続確立後、send で送信、onmessage で受信というシンプルで raw なインタフェース
    - ブラウザは接続ネゴシエーション、同一生成元ポリシー保証、メッセージのフレーム化などは担当
    - 状態管理、圧縮、キャッシュなどはアプリケーションでやる必要
    - フォーマットはテキストかバイナリを選択
    - 上位でのプロトコルはアプリごとに独自に定義
        - サブプロトコルを規定のやり方でネゴシエーションできる
- [101 Switching Protocols \- HTTP \| MDN](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/101)
    - リクエストを受けてプロトコルを変更
    - http => websocket など
    - websocket はネゴシエーション完了後に 101 でスイッチする
- XHR, SSE との比較
    - XHR
        - 「トランザクション」的なリクエスト・レスポンス型の通信に最適化
    - SSE
        - サーバからクライアントへのストリーム
    - WebSocket
        - 単一の TCP 接続で双方向伝送
    - SSE と WebSocket はサーバの準備ができ次第即時に送信できる、つまりこれはキューイング遅延の排除
        - 通信路におけるレイテンシ自体が削減されるわけではない
- 圧縮
    - バイナリはすでに圧縮済みかもしれないので、常に圧縮するのが正とは限らない
    - ローレベルの機能しか提供されておらず、圧縮はアプリケーションが行う必要がある
        - メッセージ単位圧縮の WebSocket プロトコル拡張機能は、本書執筆時点では未実装だった (今はどうだろう)
- キャッシュ
    - HTTP ではブラウザや中間装置がよしなにキャッシュしてくれるが、WebSocket はそうではない
    - やり取りするデータに応じて最適な通信方法を選ぶべき
        - リアルタイムデータやメッセージなどは WebSocket
        - キャッシュできるリソースは XHR
- インフラ
    - 中間装置で持続的接続を切るものがあると成り立たない
        - TLS でトンネル化して対策
            - これでも経路上の中間装置がアイドル状態の TCP 接続を短時間で切る場合はどうしようもない
            - ただネゴシエーションの成功確率向上、接続タイムアウトのインターバルを長くできる効果は期待できる
        - フォールバックは必ず必要
    - 自身でコントロールできるインフラは対応していく
        - nginx や HAProxy のタイムアウト設定
        - 例えば HAProxy は tunnel のタイムアウトを伸ばす
            - connect, client, server timeout は HTTP のアップグレードハンドシェイクに有効
        - いずれもデフォルトで 60 秒だったりするので、それぞれ点検していく必要

## WebRTC

- p2p で音声、動画、データ通信
- 複数のプロトコルスタック
- UDP 上に構築
- UDP -> ICE/STUN/TURN/ -> DTLS -> SRTP/SCTP -> RTCPeerConnection/DataChannel
- 接続の確立
    - シグナリングチャネル (SIP, Jingle, ISUP)
        - どれを使うかを合意
    - SDP (Session Description Protocol)
        - p2p 接続パラメータを記述するためのプロトコル
    - ICE / STUN, TURN
        - NAT トラバーサル
- DTLS
    - TLS の UDP 版 (D = Datagram)
- SRTP
    - メディア配信
    - フロー制御、輻輳制御
    - ネットワーク状況に合わせたストリームの品質調整
- SCTP
    - データ通信
    - フロー制御、輻輳制御、メッセージ順序管理、メッセージ指向 (HoL ブロッキング回避) など、TCP 的な機能の提供
        - IP のうえに直接 SCTP を乗せるのがもはや早いが、すべての中間装置が対応しないといけないので実質的に無理
        - 内部ネットワークだけなら可能性あり
- DataChannel
    - WebSocket 互換のインタフェース
- 音声・動画ストリーミング
    - ネットワークの帯域は常に変動する
    - ISP によっては上りと下りの帯域幅が違う
    - そのため常に最高の動画・音声品質ではなく、状況に応じた品質調整が行われる
- 多者間通信
    - 1 対 1 の p2p なら簡単だが、そうでない場合はトポロジを考える必要
    - 全員がつながるメッシュ、スーパーノードを決めるスター型、サーバを準備する方針
    - この部分は WebRTC のプロトコルとしては全くサポート外でアプリケーションが考える必要

## 参考

- [High Performance Browser Networking \(O'Reilly\)](https://hpbn.co/)
- [O'Reilly Japan \- ハイパフォーマンス ブラウザネットワーキング](https://www.oreilly.co.jp/books/9784873116761/)

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873116767/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://images-na.ssl-images-amazon.com/images/I/51x2sA8N+TL._SX389_BO1,204,203,200_.jpg" alt="ハイパフォーマンス ブラウザネットワーキング ―ネットワークアプリケーションのためのパフォーマンス最適化" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873116767/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">ハイパフォーマンス ブラウザネットワーキング ―ネットワークアプリケーションのためのパフォーマンス最適化</a></div><div class="amazlet-detail">Ilya Grigorik (著), 和田 祐一郎  (翻訳), 株式会社プログラミングシステム社 (翻訳)<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873116767/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
