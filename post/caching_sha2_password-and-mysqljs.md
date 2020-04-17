{"title":"mysql8 系からデフォルトになった caching_sha2_password に mysqljs/mysql はまだ対応していない","date":"2020-03-08T11:56:26+09:00","tags":["nix"]}

[mysqljs/mysql](https://github.com/mysqljs/mysql) (pure js の mysql クライアント実装) は [caching_sha2_password](https://github.com/mysqljs/mysql/pull/1962) (mysql8 系でデフォルトになった authentication plugin) に現時点では未対応。

homebrew で雑に mysql をセットアップすると、現時点では [8.0.19](https://github.com/Homebrew/homebrew-core/blob/f0f20525081da7e4f471c4c05a20f92c05be1b23/Formula/mysql.rb#L4) が入るため、mysqljs/mysql で接続しようとすると `ER_NOT_SUPPORTED_AUTH_MODE: Client does not support authentication protocol requested by server; consider upgrading MySQL client` というエラーが出る。

    > const mysql = require('mysql');
    undefined
    > const connection = mysql.createConnection({
    ...   host: 'localhost',
    ...   user: 'root',
    ...   password: 'xxx',
    ...   database: 'yyy'
    ... });
    undefined
    > connection.connect();
    undefined
    > Uncaught:
    <ref *2> Error: ER_NOT_SUPPORTED_AUTH_MODE: Client does not support authentication protocol requested by server; consider upgrading MySQL client
        at Handshake.Sequence._packetToError (/Users/cou929/Desktop/bcrypter/node_modules/mysql/lib/protocol/sequences/Sequence.js:47:14)
        at Handshake.ErrorPacket (/Users/cou929/Desktop/bcrypter/node_modules/mysql/lib/protocol/sequences/Handshake.js:123:18)
        at Protocol._parsePacket (/Users/cou929/Desktop/bcrypter/node_modules/mysql/lib/protocol/Protocol.js:291:23)
        at Parser._parsePacket (/Users/cou929/Desktop/bcrypter/node_modules/mysql/lib/protocol/Parser.js:433:10)
        at Parser.write (/Users/cou929/Desktop/bcrypter/node_modules/mysql/lib/protocol/Parser.js:43:10)
        at Protocol.write (/Users/cou929/Desktop/bcrypter/node_modules/mysql/lib/protocol/Protocol.js:38:16)
        at Socket.<anonymous> (/Users/cou929/Desktop/bcrypter/node_modules/mysql/lib/Connection.js:88:28)
        at Socket.<anonymous> (/Users/cou929/Desktop/bcrypter/node_modules/mysql/lib/Connection.js:523:10)
        at Socket.emit (events.js:321:20)
        at Socket.EventEmitter.emit (domain.js:547:15)
    ...

## 対処方法

対応しているクライアントを使う。

[MySQL :: MySQL 8\.0 Reference Manual :: 2\.11\.4 Changes in MySQL 8\.0](https://dev.mysql.com/doc/refman/8.0/en/upgrading-from-previous-series.html#upgrade-caching-sha2-password-compatible-connectors)

あるいは、検証用環境などであれば、とりあえず auth plugin を以前の `mysql_native_password` に戻すのが簡単だと思われる。

    USE mysql;
    UPDATE user SET plugin = 'mysql_native_password' WHERE user = 'root';  -- root の場合の例
    SELECT Host, User, plugin, authentication_string FROM user; -- 確認
    FLUSH PRIVILEGES;

## `caching_sha2_password`

[MySQL :: MySQL 8\.0 Reference Manual :: 2\.11\.4 Changes in MySQL 8\.0](https://dev.mysql.com/doc/refman/8.0/en/upgrading-from-previous-series.html#upgrade-caching-sha2-password)

mysql にクライアントから接続する際の認証方式。いままでは `mysql_native_password` という名前のものだったが、よりセキュアな `sha256_password` `caching_sha2_password` が追加されている。（具体的な内容までは追っていない）

クライアント側は `libmysqlclient` を使うとよいらしいが、mysqljs/mysql は pure js 実装なので、おそらくこの部分を自前で実装しなければならず、対応が追いついていないのかなと思われる。

`default_authentication_plugin` ディレクティブでデフォルトを `mysql_native_password` に設定しておけば過去の挙動を維持できそうだった。なおこれは新規ユーザーを追加する際の `mysql.User.plugin` のデフォルト値に影響する設定なので、すでに `caching_sha2_password` などで作成済みのユーザーは自分で更新する必要がある。

    [mysqld]
    default-authentication-plugin=mysql_native_password

設定反映後に:

    mysql> CREATE USER 'newuser'@'localhost' IDENTIFIED BY 'newpass';
    Query OK, 0 rows affected (0.01 sec)

    mysql> SELECT Host, User, plugin, authentication_string FROM user; -- 確認
    +-----------+------------------+-----------------------+------------------------------------------------------------------------+
    | Host      | User             | plugin                | authentication_string                                                  |
    +-----------+------------------+-----------------------+------------------------------------------------------------------------+
    ...
    | localhost | newuser          | mysql_native_password | xxx                                                                    |
    ...

## この問題にあたった背景

[auth0](https://auth0.com/) という認証機能を提供する SaaS の検証をしていた。
auth0 はよくできていて、既存のデータベースのユーザーデータと接続し、逐次データマイグレーションをしながら既存のユーザーデータ資産を活かすという使い方ができる。
`既存のデータベースとの接続` のための [グルーコードは Node.js 環境で mysqljs/mysql を利用して書く必要](https://auth0.com/docs/connections/database/custom-db/create-db-connection) がある。
この状況で手元の mysql8.0 に auth0 から接続しようとして `ER_NOT_SUPPORTED_AUTH_MODE` が発生、という経緯だった。

## 参考

- [MySQL :: MySQL 8\.0 Reference Manual :: 2\.11\.4 Changes in MySQL 8\.0](https://dev.mysql.com/doc/refman/8.0/en/upgrading-from-previous-series.html#upgrade-caching-sha2-password)
- [MySQL :: MySQL 8\.0 Reference Manual :: 6\.4\.1\.2 Caching SHA\-2 Pluggable Authentication](https://dev.mysql.com/doc/refman/8.0/en/caching-sha2-pluggable-authentication.html)
- [MySQL :: MySQL 8\.0 Reference Manual :: 2\.11\.4 Changes in MySQL 8\.0](https://dev.mysql.com/doc/refman/8.0/en/upgrading-from-previous-series.html#upgrade-caching-sha2-password-compatible-connectors)
- [MySQL :: MySQL 8\.0 Reference Manual :: 6\.4\.1 Authentication Plugins](https://dev.mysql.com/doc/refman/8.0/en/authentication-plugins.html)
- [Fix MySQL 8\.0\.x incompatibilities by ruiquelhas · Pull Request \#1962 · mysqljs/mysql](https://github.com/mysqljs/mysql/pull/1962)
- [mysql \- npm](https://www.npmjs.com/package/mysql)

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873116384/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://images-fe.ssl-images-amazon.com/images/I/51li6EcTU%2BL._SL160_.jpg" alt="実践ハイパフォーマンスMySQL 第3版" style="border: none;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873116384/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">実践ハイパフォーマンスMySQL 第3版</a><div class="amazlet-powered-date" style="font-size:80%;margin-top:5px;line-height:120%">posted with <a href="http://www.amazlet.com/" title="amazlet" target="_blank">amazlet</a> at 20.03.08</div></div><div class="amazlet-detail">Baron Schwartz Peter Zaitsev Vadim Tkachenko <br />オライリージャパン <br />売り上げランキング: 167,250<br /></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873116384/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B07GRPD4ND/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://images-fe.ssl-images-amazon.com/images/I/51zCiz39t-L._SL160_.jpg" alt="「Auth0」で作る！認証付きシングルページアプリケーション (技術の泉シリーズ（NextPublishing）)" style="border: none;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B07GRPD4ND/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">「Auth0」で作る！認証付きシングルページアプリケーション (技術の泉シリーズ（NextPublishing）)</a><div class="amazlet-powered-date" style="font-size:80%;margin-top:5px;line-height:120%">posted with <a href="http://www.amazlet.com/" title="amazlet" target="_blank">amazlet</a> at 20.03.08</div></div><div class="amazlet-detail">インプレスR&D (2018-08-31)<br />売り上げランキング: 80,200<br /></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B07GRPD4ND/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
