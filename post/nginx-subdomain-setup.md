{"title":"nginx でサブドメイン","date":"2012-02-25T19:07:29+09:00","tags":["nginx"]}

### nginx 側の設定
/var/www/please-sleep に静的なファイルを置く. please-sleep.cou929.nu でアクセスできるようにする.

    cou929:kosei% cat /etc/nginx/conf.d/virtual.conf
    #
    # A virtual host using mix of IP-, name-, and port-based configuration
    #
    
    server {
        listen       80;
        server_name  please-sleep.cou929.nu;
    
        location / {
            root   /var/www/please-sleep/;
            index  index.html index.htm;
        }
    }
    cou929:kosei% sudo /etc/init.d/nginx restart
    Stopping nginx:                                            [  OK  ]
    Starting nginx:                                            [  OK  ]

### DNS の設定
please-sleep.cou929.nu をサーバ (49.212.15.82) に向ける. A レコード.
