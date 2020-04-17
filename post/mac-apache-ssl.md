{"title":"Mac の Apache でオレオレ証明書で SSL","date":"2014-01-09T00:57:52+09:00","tags":["mac"]}

    $ openssl dgst -md5 /var/log/system.log > rand.dat
    $ openssl genrsa -des3 -rand rand.dat 1024 > server.pem
    $ openssl rsa -in server.pem -out server.pem
    $ openssl req -new -days 365 -key server.pem -out csr.pem
    $ openssl req -in csr.pem -key server.pem -x509 -out crt.pem

    # 設置場所は config (/private/etc/apache2/extra/httpd-ssl.conf) を確認。デフォルトは以下
    $ sudo cp server.pem /private/etc/apache2/server.key
    $ sudo cp crt.pem /private/etc/apache2/server.crt

    # /private/etc/apache2/httpd.conf の以下がコメントアウトされている場合はコメントを削除
    Include /private/etc/apache2/extra/httpd-ssl.conf

    $ sudo apachectl restart

### 参考

- [Macに最初から入っているApacheでSSL通信する環境を整えた](http://www.karakaram.com/mac-apache-ssl)

