{"title":"最近読んだもの 16","date":"2021-09-12T23:30:00+09:00","tags":["readings"]}

## 記事

- [The rise of ransomware \- NCSC\.GOV\.UK](https://www.ncsc.gov.uk/blog-post/rise-of-ransomware)
	- ランサムウェアの最近の動向
	- もともとは多くのクライアントに拡散してデータを人質にとって金銭を得るタイプが多かったが、それはデータのバックアップなどの対策により被害が緩和されてきた
	- ここ数年はデータを公開すると脅迫するパターンが増えてきている
	- こうなると発生後の対処は難しくなる。そもそも攻撃者が侵入できた経路を特定し対処する必要があり、それは結構大変
		- ここを怠り何度も身代金要求をされたケースもある
	- 攻撃者の立場から見ると、ROI が下がるので、昔のように無闇矢鱈に拡散させたいものでもない
	- など、知らないことが多く面白かった
- [Kill It With Fire \| USENIX](https://www.usenix.org/publications/loginonline/kill-it-fire)
	- Kill It With Fire という本の紹介
	- レガシーシステムの移行がテーマだが、技術的な側面ではなく、"システムの周りのシステム" つまり組織や慣習などに着目している
	- 運用の視点やレガシーか否かは二値ではなく程度、など、納得感のある記載が多く、面白そう
	- それはそうと、記事の筆者の Laura Nolan  さんは前に読んだ [Seeing Like an SRE: Site Reliability Engineering as High Modernism \| USENIX](https://www.usenix.org/publications/loginonline/seeing-sre-site-reliability-engineering-high-modernism) もかなり面白かった
		- [Google の SRE 本の acknowledge](https://sre.google/sre-book/preface/#acknowledgments-8ksNUe) に名前がクレジットされていた

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B08CTFY4JP/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/41T4DSKUExL.jpg" alt="Kill It with Fire: Manage Aging Computer Systems (and Future Proof Modern Ones) (English Edition)" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B08CTFY4JP/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Kill It with Fire: Manage Aging Computer Systems (and Future Proof Modern Ones) (English Edition)</a></div><div class="amazlet-detail">英語版  Marianne Bellotti  (著)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B08CTFY4JP/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

- [Why every software engineering interview should include ops questions – charity\.wtf](https://charity.wtf/2021/08/21/why-every-software-engineering-interview-should-include-ops-questions/)
	- 運用は大事だし、swe の面接でも運用系の質問をしましょう
	- それは組織が運用に価値を見出しているということの証にもなる、というのがいい言葉だった
	- 質問例もいい感じだった
		- 仕様からアーキテクチャ、インフラ設計を議論する
		- デプロイパイプラインを設計する
		- 障害の事例を聞く
		- visibility や debugging で好きなツールを聞く
		- レイテンシが悪化している際に、どこから調査するか聞く
		- ブラウザに url を打ち込むのと何が起こるか説明してもらう
- [Cascading retries and the sulky applications \- Ayende @ Rahien](https://ayende.com/blog/194626-C/cascading-retries-and-the-sulky-applications)
	- より下のレイヤーがリトライしているのに、呼び出し元でもリトライするのは無駄にリソースを使うし何も解決しないので良くない
	- またエラーを適切に報告できていない
	- 一般化すると予期しないエラーはローカルで処置せず呼び出し元に報告すること
- [Autoloading in Rails 7, get ready\! \| Riding Rails](https://weblog.rubyonrails.org/2021/9/3/autoloading-in-rails-7-get-ready/)
	- Rails7 で autoloader の zeitwerk mode がデフォルトになり、classic からの移行期間が終わる
	- それぞれどんなモードなのかこの記事ではわからず、初学者には辛いところ
- [Real\-time stress: AnyCable, k6, WebSockets, and Yabeda — Martian Chronicles, Evil Martians’ team blog](https://evilmartians.com/chronicles/real-time-stress-anycable-k6-websockets-and-yabeda)
	- AnyCable というライブラリで Rails 上に実装された WebSocket を使うリアルタイムアプリの負荷試験を k6 というツールで行う例
	- どっちも聞いたことがなかったので機会があった際の選択肢になった
