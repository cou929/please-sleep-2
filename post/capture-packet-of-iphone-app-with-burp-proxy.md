{"title":"Burp Proxy で iPhone の通信をパケットキャプチャ","date":"2014-03-22T11:08:17+09:00","tags":["mobile"]}

iPhone アプリのデバッグや挙動調査のために通信を見てみたい。[Burp Proxy](http://portswigger.net/burp/proxy.html) というソフトで proxy を mac にたてておき、iPhone の proxy 設定を mac に向けて、いわゆる Man in the Middle 方式で通信を覗いてみる。仮想の SSL 証明書を iPhone にインポートすることで SSL 通信もキャプチャできる。JailBrake は不要。

### 必要なもの

- [Burp Proxy](http://portswigger.net/burp/proxy.html)
- [iPhone 構成ユーティリティ](http://support.apple.com/kb/DL851?viewlocale=ja_JP)

### 手順

#### Burp Proxy の起動と設定

- proxy タブ -> option タブ -> proxy listeners にエントリがひとつあるので選択して edit
- Bind to address -> All interfaces を選択。ダイアログが出るが続行する
- Intercept タブに移って、"Intercept is On" になっていたらこれをオフにしておく

#### CA 証明書のエクスポート

Burp Proxy は SSL 通信をキャプチャするために一時的な CA 証明書を発行する。これを取得し iPhone にインポートする。

まず取得。ブラウザのプロキシ設定を変更し一時的に手元で動いている Burp Proxy に向ける。この状態で適当なサイトに SSL アクセスし (一時的な証明書なので当然エラーがでる)、そのタイミングで証明書をブラウザから保存してしまえば良い。以下は Firefox での手順。

- Firefox の option -> 詳細 -> ネットワーク -> 接続
- proxy の手動設定を選んで SSL proxy を 'localhost:8080' として設定する

<img style="width:90%" alt="" src="/images/burp_fx_proxy.png"/>

- 適当な ssl のサイトにアクセス
  - proxy 設定をしているので、ここでリクエストが burp へ飛ぶはず。burp 側の history タブにリクエスト内容が出ていれば OK。そうでない場合はなんらかの理由でリクエストが firefox から burp へ届いていないことになる。
- 証明書のエラーが出るので 危険性を理解した上で接続するには -> 例外 -> 表示 -> 詳細

<img style="width:90%" alt="" src="/images/burp_ssl_error.png"/>
<img style="width:90%" alt="" src="/images/burp_cert.png"/>

- エクスポート。フォーマットの設定はそのままでいいはず。

<img style="width:90%" alt="" src="/images/burp_export.png"/>

- 保存したファイルを my.cer などとリネームしてダブルクリック。これで Keychain Access に追加される

#### CA 証明書のインポート

こうしてエクスポートした証明書を iPhone にインポートする。iPhone 構成ユーティリティを使う

- 構成ユーティリティを起動して iPhone を接続する
- 左カラムの構成プロファイルを選んで左上の新規追加ボタン
- 資格情報を選択し、先ほどの PortSwiggerCA.cer を選んで読み込む

<img style="width:90%" alt="" src="/images/burp_iphone_conf.png"/>

- 読み込んだ証明書を選択したまた General を選択 -> Name と Identifier に適当な値を設定する

<img style="width:90%" alt="" src="/images/burp_identify.png"/>

- 左タブから接続した iPhone を選択 -> 構成プロファイル
- 追加した証明書をインストール。パスコードロックがかかっている場合は解除を促される

<img style="width:90%" alt="" src="/images/burp_profile_install.png"/>

- iPhone 側の画面が切り替わるのでインストールボタンを押す

### iPhone のプロキシ設定

Wifi が前提。

- 設定 -> Wifi
- 接続中のアクセスポイントの右端の i ボタンを押す
- 詳細画面に移るので一番下の HTTP プロキシでホストとポートを設定
  - ホストは Burp Proxy をたてた Mac の ip を入れる。ifconfig -a などで確認
  - ポートは 8080

設定は以上。あとは iPhone でなんらかの通信をするとすべて Mac の Burp Proxy を経由するようになる。Burp Proxy の History タブにすべての通信が出るのでこちらをみればよい。例えば画像リクエストを出さないなどのフィルタなども可能。

もちろん tcpdump でもいいのだがこうした GUI ソフトを使うとリクエストレスポンスをまとめてくれたりして見やすいし機能もリッチで便利だ。

### キャプチャがうまくいかない場合

#### プライバシーセパレータ

Wifi ルータによってはプライバシーセパレータが有効になっているため、PC <-> iPhone 間の疎通ができない場合がある。

[無線パソコン同士の通信を禁止する（プライバシーセパレータ）](http://buffalo.jp/download/manual/html/air851/router/whrg54s/chapter11.html)

キャプチャがうまくいかない場合、ルーターの設定変更の権限がある場合は、こちらを疑ってみるとよいかも。

#### 起動方法

自分の環境では open コマンドで burp を起動するとキャプチャできず、java コマンドで起動すると問題ないということがあった。

    # OK
    $ java -jar burpsuite_free_v1.5.jar
    
    # NG
    $ open burpsuite_free_v1.5.jar

正確には、iPhone -> pc (proxy) への疎通はできているが、レスポンスを返せていない (tcpdump してみてみると iPhone -> PC 方向のパケットは飛んでいるが逆方向がない)。

ps してみてみるとそれぞれちがう java を使っているようで、

    # java コマンドで起動
    /usr/bin/java -jar Documents/burpsuite_free_v1.5.ja
    # 実態はこちら
    $ ls -l /usr/bin/java
    lrwxr-xr-x  1 root  wheel  74 11  1 06:28 /usr/bin/java@ -> /System/Library/Frameworks/JavaVM.framework/Versions/Current/Commands/java

    # open コマンドで起動
    /Library/Internet Plug-ins/JavaAppletPlugin.plugin/Contents/Home/bin/java -jar /Users/cou929/Documents/burpsuite_free_v1.5.jar

これ以上深追いはしないが、この違いが問題になっていたようだ。

#### 8080 をすでに奪われている場合

意外と気が付きづらいのが、手元に nginx なり apache なりを起動していることを忘れていて、8080 へのリクエストが burp ではなくそれらを Web サーバにいってしまうというケース。地味にはまるので注意。

#### burp 1.5 と java 1.8 (jdk8) の相性が悪い

いまのところ (2014 年 3 月現在) は、burp 1.5 と java 1.8 は相性が悪いようだ。java 1.8 の場合、 burp 起動時に "java のバージョンが新しすぎるよ" という警告が出る。また手元では、はじめの firefox から接続し証明書を取得する手順で、`ssl_error_no_cypher_overlap` というエラーが出た。このエラーは、クライアントとサーバで共通して使える SSL の暗号化アルゴリズムがみつからない、という意味らしい。

自分の場合は java 1.7 系にダウングレードすると問題なく動作した。方法は [こちら](http://docs.oracle.com/javase/7/docs/webnotes/install/mac/mac-jdk.html) にある通り、`/Library/Java/JavaVirtualMachines/jdk.*.jdk` を削除し、[古いバージョンの jdk](http://www.oracle.com/technetwork/java/javase/downloads/jdk7-downloads-1880260.html) を改めてインストールしなおせば良い。

### 参考

- [[改訂版] iPhoneアプリのSSL接続をパケットキャプチャする方法 | [ bROOM.LOG ! ]](http://blog.rocaz.net/2011/02/1167.html)

<iframe src="http://rcm-fe.amazon-adsystem.com/e/cm?lt1=_blank&bc1=000000&IS2=1&bg1=FFFFFF&fc1=000000&lc1=0000FF&t=pleasesleep-22&o=9&p=8&l=as4&m=amazon&f=ifr&ref=ss_til&asins=B00B71KZNI" style="width:120px;height:240px;" scrolling="no" marginwidth="0" marginheight="0" frameborder="0"></iframe>
<iframe src="http://rcm-fe.amazon-adsystem.com/e/cm?lt1=_blank&bc1=000000&IS2=1&bg1=FFFFFF&fc1=000000&lc1=0000FF&t=pleasesleep-22&o=9&p=8&l=as4&m=amazon&f=ifr&ref=ss_til&asins=B00GJGOPDW" style="width:120px;height:240px;" scrolling="no" marginwidth="0" marginheight="0" frameborder="0"></iframe>
