{"title":"最近読んだもの 37 - Understainding Software Dynamics 6 章まで、TCP/IP 基礎など","date":"2022-03-20T23:30:00+09:00","tags":["readings"]}

コロナ関連のもろもろ (保育園の休園、ワクチン 3 回目摂取の副反応、家族の副反応中のワンオペ) で時間がとれず、体調も悪かったりして、前回から 3 週間スキップしてしまった。ようやく生活が普段のペースにもどりつつあるので再開していく。

- [Understanding Software Dynamics](http://www.amazon.co.jp/exec/obidos/ASIN/B09H5JB5HC/pleasesleep-22/ref=nosim/)
    - 6 章まで読んだ
    - サーバ・クライアント間の通信の計測まで話が進んで少し読みやすくなってきた (より関心領域に近づいてきたので)
    - 超大枠だとある意味当たり前のことしか言っていないが、細部の深さが異常な面白い本だと改めて思う
    - ただ気合を入れないとページが進まないので、この三週間は進捗が出せなかった
- [Linuxで動かしながら学ぶTCP/IPネットワーク入門](http://www.amazon.co.jp/exec/obidos/ASIN/B085BG8CH5/pleasesleep-22/ref=nosim/)
    - ネットワークの特に L4 から下にずっと自信がなかったので読んでみたが良い本だった
    - コマンドを手打ちして理解を固めながら知識を積んでいけるので気持ちがよい
- [Memorystore for Redis \| Google Cloud Blog](https://cloud.google.com/blog/products/databases/memorystore-for-redis)
    - read replica が GA
    - snapshot 取得機能が preview
    - basic tier で定期メンテ時にデータが飛ばなくなった
- [Scaling Google Cloud Memorystore for high performance \| Google Cloud Blog](https://cloud.google.com/blog/products/databases/scaling-google-cloud-memorystore-for-high-performance)
    - GCP の Memorystore Redis の write パフォーマンスをスケールさせるためにシャーディングを行う例
    - proxy として Envoy をはさむことでクライアント側のコード変更を最小化
    - Memtier という kvs のベンチマークツールは知らなかった
- [Application exceptions surfaced automatically \| Google Cloud Blog](https://cloud.google.com/blog/products/devops-sre/application-exceptions-surfaced-automatically)
    - 普通に Error Reporting の紹介をしている記事
    - 新機能の告知などかと思って読んだがそうではなかった
    - Error reporting は便利なツールだけど、検出ロジックをもうちょっと何とかカスタマイズしたい
- [The Truth About “MEH\-TRICS” \- Honeycomb](https://www.honeycomb.io/blog/truth-about-meh-trics-metrics/)
    - observability とその対になっている概念らしい metrics の比較とそれぞれの適材適所の説明
    - 前提となっている observability の定義がよくわからなかった
    - この会社は observability tool を作っていて、metrics ツールの最適解は Prometheus で、そんな中で新たに metrics tool を出す理由の説明をしていた
- [So You Want To Build An Observability Tool\.\.\. \- Honeycomb](https://www.honeycomb.io/blog/so-you-want-to-build-an-observability-tool/)
    - 上の記事の補足として読んでみた
    - とにかくいろいろな情報を構造化して出しておくことで、事前に予想し得ない問題もあとから観測できるようにしておくみたいなことらしい
    - 言いたいことはわかる気はするけど、observability vs metrics のように独自の言葉の定義論争みたいなものに持ち込もうとするのは何とも言えなかった

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B09H5JB5HC/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/41aDqfiWNbL.jpg" alt="Understanding Software Dynamics (Addison-Wesley Professional Computing Series) (English Edition)" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B09H5JB5HC/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Understanding Software Dynamics (Addison-Wesley Professional Computing Series) (English Edition)</a></div><div class="amazlet-detail">英語版  Richard L Sites  (著)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B09H5JB5HC/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B085BG8CH5/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/51kU2EFP5UL.jpg" alt="Linuxで動かしながら学ぶTCP/IPネットワーク入門" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B085BG8CH5/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Linuxで動かしながら学ぶTCP/IPネットワーク入門</a></div><div class="amazlet-detail">もみじあめ  (著)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B085BG8CH5/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
