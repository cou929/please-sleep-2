{"title":"Mac のローカルに WordPress をたてる","date":"2015-02-01T23:06:31+09:00","tags":["nix"]}

<img src="/images/wordpress-logo-stacked-rgb.png" style="width:30%"/>

検証に必要だったので数年ぶりにセットアップしたのでメモ。

### Apache のセットアップ

Apache 自体はすでにインストール済みのはずなので、設定ファイルで php を有効化するだけでよい。

設定ファイルは `/etc/apache2/httpd.conf` にあるので、`#LoadModule php5_module libexec/apache2/libphp5.so` の行のコメントを外す。

プロセスの起動は `apachectl` で良い。

    $ sudo apachectl start

ブラウザから `http://localhost/` にアクセスして、いつもの `It works!` が表示されれば動作は OK。

ドキュメントルートの位置は、デフォルトでは `/Library/WebServer/Documents` だと思う。`httpd.conf` の `DocumentRoot` ディレクティブで確認できる。またエラーログの位置もデフォルトは `/private/var/log/apache2/error_log` で、`ErrorLog` ディレクティブで確認できる。

php の動作確認は次のような内容のファイルで。

    <?
    phpinfo();
    ?>

### mysql のセットアップ

インストールは homebrew で。

    $ brew install mysql

設定ファイルの位置やプロセスの起動方法は `brew info mysql` で表示される。おそらくそれぞれ `/etc/my.cnf` と `launchctl load ~/Library/LaunchAgents/homebrew.mxcl.mysql.plist` あたりがデフォルトだと思う。

最初はパスワード無しの root ユーザーができているので、必要に応じて wordpress 用のユーザーも作っておく。

    $ mysql -uroot -p
    mysql> GRANT ALL ON *.* TO wpuser@"localhost" IDENTIFIED BY "PASSWORD";
    mysql> FLUSH PRIVILEGES;

### php のセットアップ

mysql との接続設定が必要。php の設定ファイルは `/etc/php.ini.default` にある。これを `php.ini` にリネームし、`mysql.default_socket` に `/tmp/mysql.sock` を設定する。

    $ sudo mv /etc/php.ini.default /etc/php.ini
    $ sudo vim /etc/php.ini
    # mysql.default_socket = /tmp/mysql.sock と設定する
    $ sudo apachectl start

### WordPress のセットアップ

[WordPress › Download WordPress](https://wordpress.org/download/)

こちらからダウンロードし展開、ドキュメントルート ``/Library/WebServer/Documents` に置き、ブラウザからアクセスする。ウィザードが始まるはずなので、あとはそれに従えば良い。

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00KYKYC8W/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="http://ecx.images-amazon.com/images/I/51xdO2y5LAL._SL160_.jpg" alt="WordPress Perfect GuideBook 3.x対応版" style="border: none;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00KYKYC8W/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">WordPress Perfect GuideBook 3.x対応版</a><div class="amazlet-powered-date" style="font-size:80%;margin-top:5px;line-height:120%">posted with <a href="http://www.amazlet.com/" title="amazlet" target="_blank">amazlet</a> at 15.02.01</div></div><div class="amazlet-detail">ソーテック社 (2014-06-15)<br />売り上げランキング: 22,601<br /></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00KYKYC8W/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
