{"title":"nginx の error_page ディレクティブでステータスコード書き換え","date":"2012-12-18T12:59:55+09:00","tags":["nix"]}

nginx の `error_page` というディレクティブでエラーの時に返す html を指定できるが, その際にステータスコードの上書きもできる.

    error_page  500 502 504 =200 /error.html;

こうしておくとエラーの際には `/error.html` を返しつつ, 500, 502, 504 の際は 200 を返すようになる.

なんとしてもクライアント側に迷惑を掛けたくない (ブラウザコンソールにエラーが出ることすら避けたい) 場合にこうして 200 にたおしておく設定はありだろう.

[HttpCoreModule](http://wiki.nginx.org/HttpCoreModule#error_page)
