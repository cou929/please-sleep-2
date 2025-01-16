{"title":"証明書のピンニングをしている自アプリのデバッグ","date":"2025-01-16T20:00:00+09:00","tags":["linux", "tls"]}

モバイルアプリの HTTP 通信を見ながらデバッグしたい場合、[mitmproxy](https://mitmproxy.org/) や [Wireshark](https://www.wireshark.org/) といったプロキシを挟む (意図的に MITM を行い通信を見る) という定番の方法がある。

[TLS 通信のパケットキャプチャ \- Please Sleep](https://please-sleep.cou929.nu/decrypting-tls-traffic-packet-capture.html)

ただし証明書のピンニング (ピン留め, Certificate Pinning, SSL Pinning) をしている場合、中間で通信を覗こうと思っても単純にはできない。プロキシデフォルトの証明書をクライアントは信用しないため、通信に失敗する。

対象が自プロジェクトや自社のアプリで、管理している証明書にアクセスできる場合、これを回避できる。プロキシに正しい証明書を渡すことでクライアントはそれを信用し通信を行う。あとは通常通りプロキシから通信内容を参照できる。

ここでは Leaf certificate をピンニングしているケースを想定している。この状況で mitmproxy を使って iOS のネイティブアプリをデバッグする。ピンニングの方式やプロキシソフトによって手順の詳細は変わるので注意。

# 証明書の準備

PEM 形式の証明書を用意する。実際にピンニングしているものを使うので、取り扱いには十分注意する。

PEM 形式は概ね次のようなフォーマット。

```
-----BEGIN PRIVATE KEY-----
<private key>
-----END PRIVATE KEY-----
-----BEGIN CERTIFICATE-----
<cert>
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
<intermediary cert (optional)>
-----END CERTIFICATE-----
```

秘密鍵と証明書のそれぞれのファイルがある場合は次のようにして結合する。

```sh
cat cert.key cert.crt > cert.pem
```

# 証明書を受け渡して Proxy を起動する

[`--certs` オプション](https://docs.mitmproxy.org/stable/concepts-options/#certs) で任意の (Leaf) 証明書を使うよう mitmproxy に指定できる。

```sh
# <domain>=<cert file path> という記法。
mitmproxy --certs 'please-sleep-cou929.nu=/path/to/cert.pem'
```

> https://docs.mitmproxy.org/stable/concepts-options/#certs
>
> SSL certificates of the form "[domain=]path". The domain may include a wildcard, and is equal to "*" if not specified. The file at path is a certificate in PEM format. If a private key is included in the PEM, it is used, else the default key in the conf dir is used. The PEM file should contain the full certificate chain, with the leaf certificate as the first entry.

# 動作確認

モバイルデバイスから試す前に、CLI での確認をしておくと効率的。

mitmproxy 認証局の証明書を使ってアクセスする状況を curl などで確認できる。デフォルトでは mitmproxy は `~/.mitmproxy` に認証局のキーを作成する。こちらを引数として渡せば良い。

```sh
# mitmproxy がローカルホストの 8080 番ポートで待ち受けていると仮定。
curl --proxy 127.0.0.1:8080 --cacert ~/.mitmproxy/mitmproxy-ca-cert.pem https://please-sleep.cou929.nu
```

以上で通信が成功すれば OK。ここまででピンニング以外の部分は確認できる。

返される証明書の内容も確認しておくとよい。

```sh
# host, port に Proxy の接続先を指定する。
openssl s_client -servername please-sleep.cou929.nu -host 127.0.0.1 -port 8080 -showcerts </dev/null 2>/dev/null

# 例えばクライアントでは Leaf 証明書の base64 エンコードした文字列を比較していると行ったケースでは、次のようにしても確認しやすいかもしれない。
## 上記で取得した証明書部分を cert.crt に書き出しておく。
cat cert.crt | openssl x509 -pubkey -noout | openssl rsa -pubin -outform der 2>/dev/null |  openssl dgst -sha256 -binary | openssl enc -base64
```

これらの内容がピンニングしているクライアントが想定しているものと一致すれば OK。ここまででピンニングも含めた通信ができていることが確認できる。

# 端末からのアクセス

以降は通常のプロキシ利用と同じ手順で良い。

- 前述の手順で mitmproxy のプロセスが起動されている
- iPhone 側でプロキシと証明書の設定
    - `Settings > Wifi > 詳細ページ > Configure Proxy` にて、`Manual` をタップ
    - mitmproxy を動かしているホストの ip と port 8080 (デフォルト) を入力し Save
    - Safari で `mitm.it` にアクセスし証明書をダウンロード
    - `Settings > General > Profiles mitmproxy` を選択しインストール
    - `Settings > General > About > Certificate Trust Settings` より mitmproxy のトグルを有効化
- 該当のアプリで適当に通信を行う。mitmproxy でその内容が確認できれば成功

詳細は [TLS 通信のパケットキャプチャ \- Please Sleep](https://please-sleep.cou929.nu/decrypting-tls-traffic-packet-capture.html) に詳しい。

# 参考

- [Certificates - mitmproxy docs](https://docs.mitmproxy.org/stable/concepts-certificates/)
- [Pinning \- OWASP Cheat Sheet Series](https://cheatsheetseries.owasp.org/cheatsheets/Pinning_Cheat_Sheet.html)
- [Identity Pinning: How to configure server certificates for your app \- Discover \- Apple Developer](https://developer.apple.com/news/?id=g9ejcf8y)

# PR

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/490868619X/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/51DGE8zEr6L._SY425_.jpg" alt="プロフェッショナルTLS＆PKI 改題第2版" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/490868619X/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">プロフェッショナルTLS＆PKI 改題第2版</a></div><div class="amazlet-detail">Ivan Ristić (著), 齋藤孝道 (監修)<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/490868619X/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
