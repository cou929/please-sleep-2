{"title":"MySQL5.6 と sql_mode","date":"2014-01-06T19:11:00+09:00","tags":["nix"]}

MySQL5.6 を RPM から入れるなどすると `/usr/my.cnf` ができている。これは 5.6 で変更の入った `mysql_install_db` が作っているらしい。(すでに同名のファイルがある場合は `my-new.cnf` を作るらしい)

この my.cnf の内容はほぼコメントアウトだけで実質的に sql_mode を設定しているだけだ。かつ読み込み順的に `/etc/my.cnf` の設定を上書きする。

`mysql_install_db` が作る my.cnf では sql_mode を次のように設定している。

    sql_mode=NO_ENGINE_SUBSTITUTION,STRICT_TRANS_TABLES

よって 5.6 にバージョンアップしたとたんに sql_mode が strict になっているように見えてはまるケースが多いようだ。自分もはまった。

対応は `mysql_install_db` が作った `/usr/my.cnf` 内の sql_mode の設定を削除するか、ファイル自体を消せばよい。オンラインだと

    SET @@global.sql_mode='';

などとすれば OK

### 参考

- [日々の覚書: MySQL5.6が勝手にsql_modeを書き換えてくれる話](http://yoku0825.blogspot.jp/2013/03/mysql56sqlmode.html)

