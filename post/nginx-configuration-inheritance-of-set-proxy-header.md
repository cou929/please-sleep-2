{"title":"nginx の proxy_set_header の継承ではまった","date":"2013-05-18T20:59:22+09:00","tags":["nix"]}

> proxy_set_header directives issued at higher levels are only inherited when no proxy_set_header directives have been issued at a given level.

[HttpProxyModule](http://wiki.nginx.org/HttpProxyModule#proxy_set_header)

とのことなので、親設定ファイルで指定されていた `proxy_set_header` の設定は、子設定ファイルのもので完全に上書きされる。差分を継承しようとしてまりがちだ。

例えば `/etc/nginx/nginx.conf` には次の設定が、



    proxy_set_header Host              $host;
    proxy_set_header X-Real-IP         $remote_addr;
    proxy_set_header X-Forwarded-For   $proxy_add_x_forwarded_for;

    include /etc/nginx/conf.d/*.conf;

そして `/etc/nginx/conf.d/sample.conf` に次の設定があった場合、

    proxy_set_header X-Forwarded-Proto $scheme;

`proxy_set_header` で Host, X-Real-IP, X-Forwarded-For, X-Forwarded-Proto の 4 つが有効になりそうにみえる。しかし実際には X-Forwarded-Proto のみが設定されて、のこりの 3 つについては何も指定していないのと同様の挙動になる。
