{"title":"最近読んだもの 39 - USD 7 章、MySQL スケールなど","date":"2022-04-03T23:30:00+09:00","tags":["readings"]}

- [Understanding Software Dynamics](http://www.amazon.co.jp/exec/obidos/ASIN/B09H5JB5HC/pleasesleep-22/ref=nosim/)
	- 7 章まで読んだ
	- 今度は 6 章の実験環境に加えて、Disk IO があり、たまにテーブルロックもある状況で、複数クライアントから並行してリクエストするケース
	- 例えばロック待ちで 1 並列時にはなかった遅延が発生したり、バッファからディスクにフラッシュするタイミングでレイテンシが悪化したりする様子を確認していく
		- こう書くと平凡な内容にみえるかもしれないが、本文ではこれをかなり詳細に確認しながら論を進めている
	- 最先端のソフトウェアエンジニアリングはこんな感じなのかなと思わされたり、自分の普段の仕事と比較したりして、圧倒されている
	- どんどん面白くなってきた
- [Scaling MySQL stack, ep\. 1: Timeouts · Kir Shatrov](https://kirshatrov.com/posts/scaling-mysql-stack-part-1-timeouts/)
	- Global なクエリ実行時間上限は短くして (`SET GLOBAL MAX_EXECUTION_TIME=2000` など)、既知のスロークエリはクエリ単位でタイムアウトを伸ばす (`SELECT /*+ MAX_EXECUTION_TIME(1000) */ FROM ` など) のがおすすめ
	- コードベースが大規模な場合への、コントローラーのアクション単位でタイムアウトを管理する仕組みも紹介している
- [Scaling MySQL stack, ep\. 2: Deadlines · Kir Shatrov](https://kirshatrov.com/posts/scaling-mysql-stack-part-2-deadlines/)
	- 次は個別のクエリのタイムアウトではなく、リクエストごとのグローバルな実行時間上限を設ける話
	- ruby では [rack-timeout](https://github.com/sharpstone/rack-timeout) が有名だが問題もある
		- たぶんだけど、制限時間に達したら「割り込み」的な方法で処理を中断し例外を投げ、その際に Go でいう `defer` のような後処理がうまくできずに状態の一貫性が保てないというリスクが本質的にあるらしい
	- 対策として Go でいう context のような `deadline` という概念をおすすめしていた
		- リクエストを受けたタイミングからの経過時間を持っておき、クエリを投げる・API を呼び出すといった CPU が待たされる処理の前に時間切れになっていないかチェックする
		- そのクエリや API 呼び出し自体のタイムアウトは上記の `MAX_EXECUTION_TIME` のような個別のタイムアウトでカバーする
		- ということらしい
	- 欠点として CPU が待たされる処理を行う全箇所に自分で deadline チェック処理を埋め込む必要があるが、たいていは支配的なのは RDBMS へのクエリと外部 API との通信くらいなので、そこだけおさえておけば十分
- [rack\-timeout/risks\.md at master · sharpstone/rack\-timeout](https://github.com/sharpstone/rack-timeout/blob/master/doc/risks.md#risks-and-shortcomings-of-using-racktimeout)
	- rack-timeout のリスク
	- 上記の記事の補足
- [Scaling MySQL stack, ep\. 3: Observability · Kir Shatrov](https://kirshatrov.com/posts/scaling-mysql-stack-part-3-observability/)
	- 次はクエリにコメントで発行元を出しておくと便利だよという話
	- ruby の場合は [basecamp/marginalia: Attach comments to ActiveRecord's SQL queries](https://github.com/basecamp/marginalia) が有名
- [10 Books Shopify’s Tech Talent Think You Should Read — Culture \(2022\)](https://shopify.engineering/shopify-tech-talent-book-recommendations)
	- Shopify のエンジニアによるブックガイド
	- [Cal Newport の Deep Work](http://www.amazon.co.jp/exec/obidos/ASIN/B00X7D8X8S/pleasesleep-22/ref=nosim/)
		- 同じ著者の [デジタルミニマリスト](http://www.amazon.co.jp/exec/obidos/ASIN/B07YG2PH4Q/pleasesleep-22/ref=nosim/) という本が (軽薄な自己啓発本ぽい邦題とは裏腹に) 面白かったので気になった
		- [和訳もされている](http://www.amazon.co.jp/exec/obidos/ASIN/4478068550/pleasesleep-22/ref=nosim/) みたいだけど Amazon での翻訳の評価が低そう
	- [Extreme Ownership: How U.S. Navy SEALs Lead and Win](http://www.amazon.co.jp/exec/obidos/ASIN/B0739PYQSS/pleasesleep-22/ref=nosim/)
		- 自分の中でのあるあるだけど、マネジメント系のテーマで軍事か医療での知見を扱う本は気になりがち
	- [A Philosophy of Software Design, 2nd Edition](http://www.amazon.co.jp/exec/obidos/ASIN/B09B8LFKQL/pleasesleep-22/ref=nosim/)
		- 素直に面白そう

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B09H5JB5HC/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/41aDqfiWNbL.jpg" alt="Understanding Software Dynamics (Addison-Wesley Professional Computing Series) (English Edition)" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B09H5JB5HC/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Understanding Software Dynamics (Addison-Wesley Professional Computing Series) (English Edition)</a></div><div class="amazlet-detail">English Edition  by Richard L Sites  (著)  Format: Kindle Edition<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B09H5JB5HC/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00X7D8X8S/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/51vmivI5KvL.jpg" alt="Deep Work: Rules for Focused Success in a Distracted World (English Edition)" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00X7D8X8S/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Deep Work: Rules for Focused Success in a Distracted World (English Edition)</a></div><div class="amazlet-detail">English Edition  by Cal Newport  (著)  Format: Kindle Edition<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00X7D8X8S/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0739PYQSS/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/51yoHjJDQ3L.jpg" alt="Extreme Ownership: How U.S. Navy SEALs Lead and Win (English Edition)" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0739PYQSS/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Extreme Ownership: How U.S. Navy SEALs Lead and Win (English Edition)</a></div><div class="amazlet-detail">English Edition  by Jocko Willink  (著), Leif Babin  (著)  Format: Kindle Edition<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0739PYQSS/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B09B8LFKQL/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/51MaNQzHTQL.jpg" alt="A Philosophy of Software Design, 2nd Edition (English Edition)" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B09B8LFKQL/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">A Philosophy of Software Design, 2nd Edition (English Edition)</a></div><div class="amazlet-detail">English Edition  by John K. Ousterhout  (著)  Format: Kindle Edition<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B09B8LFKQL/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
