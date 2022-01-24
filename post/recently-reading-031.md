{"title":"最近読んだもの 31","date":"2022-01-23T23:30:00+09:00","tags":["readings"]}

## 記事

- [Designing Tinder \- High Scalability \-](http://highscalability.com/blog/2022/1/17/designing-tinder.html)
	- Designing シリーズ Tinder 編
	- 位置的に近い人同士がマッチングするので、ユーザーを位置情報でインデックスしないといけない。こういう要件は経験込みがないので新鮮だった
	- google s2 という位置情報をセルという領域の階層に分類する？仕組みがあるらしい
	- アクティブユーザーの標準偏差が小さくなるようにシャードを分ける
		- 人が少ないシャードは地理的に広く、多いところは狭くなる
	- あとはストリームを挟んだマイクロサービスの連携は大規模システムの定石ぽい
- [Cloud SQL for MySQL launches database auditing \| Google Cloud Blog](https://cloud.google.com/blog/products/databases/cloud-sql-for-mysql-launches-database-auditing)
	- CloudSQL の MySQL 用の監査ログプラグインが出たらしい
	- 特定のテーブルへの参照以外の操作を Cloud Logging に送ったりできる
	- 監査用ではないが、調査でこの手の情報が欲しいことが結構あるので、覚えておこう
- [How to Fix Slow Code in Ruby — Development](https://shopify.engineering/how-fix-slow-code-ruby)
	- Profiling, Benchmarking の丁寧な説明
	- Shopify は app_profiler を使っているらしい
- [Some ways DNS can break](https://jvns.ca/blog/2022/01/15/some-ways-dns-can-break/)
	- DNS にまつわる困った小逸話いろいろ
	- どれも気付きづらい、表層的なエラーメッセージと route cause が遠くなりがちなのが難しいなと思う
- [SRE and the Practice of Practice \| Blameless](https://www.blameless.com/sre/sre-and-the-practice-of-practice)
	- 紹介されていたデレクベイリーのインプロビゼーションに関する本が面白そうだった

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4875022220/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://images-na.ssl-images-amazon.com/images/I/513cP-2Y02L._SX355_BO1,204,203,200_.jpg" alt="インプロヴィゼーション―即興演奏の彼方へ" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4875022220/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">インプロヴィゼーション―即興演奏の彼方へ</a></div><div class="amazlet-detail">デレク ベイリー (著), Derek Bailey (原著), 竹田 賢一 (翻訳), 斉藤 栄一 (翻訳), 木幡 和枝 (翻訳)<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4875022220/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
