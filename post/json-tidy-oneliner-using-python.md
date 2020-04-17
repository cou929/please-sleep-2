{"title":"CLI で json を整形する","date":"2013-08-14T12:19:49+09:00","tags":["nix"]}

`python -mjson.tool` に渡すだけで OK。

    $ curl -v 'http://some.api.com/some/resource.json' | python -mjson.tool
    ...
    * Closing connection #0
    * SSLv3, TLS alert, Client hello (1):
    {
        "foo": [ 1, 2, 3 ],
        "bar": "baz
    }

json.tool は python2.6 以降ならばコアに入っているモジュール。CentOS 5 系などはデフォルトの python が 2.4 なので注意。
