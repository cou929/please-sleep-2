{"title":"rpm コマンドでパッケージの削除","date":"2012-12-18T13:03:05+09:00","tags":["nix"]}

`rpm -qa` で一覧取得. `-e` で削除

    $ sudo rpm -qa | grep setuptool
    setuptool-1.19.2-1.el5.centos
    $ sudo rpm -ev setuptool-1.19.2-1.el5.centos
