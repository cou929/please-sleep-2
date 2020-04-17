{"title":"nginx で ip によるアクセス制限","date":"2012-12-18T17:55:12+09:00","tags":["nix"]}

ふつうに `allow`, `deny` で OK. 上から順に評価していってマッチすると抜けるので, `deny all` を先頭に書くとアクセスできなくなるので注意

    location / {
      deny    192.168.1.1;
      allow   192.168.1.0/24;
      allow   10.1.1.0/16;
      allow   2620:100:e000::8001;
      deny    all;
    }

[HttpAccessModule](http://wiki.nginx.org/HttpAccessModule)
