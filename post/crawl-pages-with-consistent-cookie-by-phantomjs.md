{"title":"js と Cookie を有効にしつつクローリングしたい","date":"2013-06-30T21:21:47+09:00","tags":["javascript"]}

- phantomjs コマンドに `--cookies-file` オプションでいわゆる cookie_jar のようにクッキーを保存するファイルを指定できる
  - ただこれをプログラムファイル側から扱う方法は無いみたい
- `phantom.open` はいわゆる onload のタイミングでコールバックを呼ぶが, それだと後から呼ぶ処理やサードパーティのスクリプトが動く前だったりするので, とりあえず setTimeout で逃げた
- phantom.exit でプロセスは終わらない? プロセスを抜けるにはどうしたらいいんだろう
- リクエスト・レスポンスヘッダとかステータスコードをリクエストごとに出したいけどとれないっぽい
- phantomjs をバックエンドにした curl って作れないかなあ

<script src="https://gist.github.com/cou929/4116887.js"></script>
