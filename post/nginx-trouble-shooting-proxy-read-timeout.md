{"title":"nginx 今日あったトラブルシューティング","date":"2012-12-19T20:33:42+09:00","tags":["nginx"]}

### 現象
とあるページだけ 404 になる. 大丈夫な時もある

### 原因
nginx をリバースプロキシとして使い, 後ろに starman でアプリ (UI) のプロセスが立っている構成. ある 1 ページは処理が重く, nginx 側の proxy_read_timeout に引っかかって nginx 側で切られていた.

- その重い処理のページは過去しばらくの期間のレポートを取得するようなもので, 時間がかかることは承知のもの
- proxy_read_timeout は他のアプリと共通設定のため, かなり辛めの数値になっていた

### 対応
- proxy_read_timeout を伸ばす
- puppet に今回のアプリ用の設定をたてて, デフォルトの設定を使わないように

### ポイント
特定のページだけ 404 だったので, アプリのルーティングまわりがおかしいのかと推定, コントローラーあたりのコミットログを見始めていたが, 今回はそこではなかった. まっさきに nginx のエラーログを見ればもっと早く解決できそうだった.

    2012/12/19 19:57:13 [error] 20613#0: *4065966 upstream timed out (110: Connection timed out) while reading response header from upstream, client: XXX.XXX.XXX.XXX, server: example.com, request: "GET /heavy/page HTTP/1.1", upstream: "http://XXXXXXX/heavy/page", host: "example.com"

こんな感じのエラーが出ていた. まさに今回のページへのアクセスのエラーで, upstream のタイムアウト.

エラーログをひと通り見てからコードを調べるようにしたほうがよさそう.
