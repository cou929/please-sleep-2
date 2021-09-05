{"title":"最近読んだもの 15","date":"2021-09-05T22:30:00+09:00","tags":["readings"]}

## 記事

- [Computers are the easy part \| Mailchimp Developer](https://mailchimp.com/developer/blog/computers-are-the-easy-part/)
	- 思い込みによってシンプルな原因をみんな見落としていたという事例。興味深い
	- 外野からみると簡単そうだけど、合理的に思考する何人ものベテランなエンジニアが見落とすことになったという事実
	- 紹介されていたこの本も面白そう

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Q8XCSFI/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/51o5UQmwcZL.jpg" alt="The Field Guide to Understanding 'Human Error' (English Edition)" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Q8XCSFI/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">The Field Guide to Understanding 'Human Error' (English Edition)</a></div><div class="amazlet-detail">英語版  Sidney Dekker  (著)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Q8XCSFI/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

- [A defense of boring languages](https://danluu.com/boring-languages/)
	- つまらない言語、ここでは c, cpp, java を、なめるなよと。大事なソフトウェアは大体これらで書かれているよと
	- 言語の知識よりドメイン知識のほうが大事、学びがあるし、結局働きたい人たちやシステムが何を書くか、何で書かれているかに合わせるといのは共感
- [My Recipe for Optimizing Performance of Rails Apps](https://pawelurbanek.com/optimize-rails-performance)
	- Rails アプリ（と書いてあるが一般的な内容だった）のパフォーマンス改善の際に、レイヤーごとに分類して対応するアプローチ
	- フロントエンド
	- バックエンド
		- エンドポイントごとのレイテンシの可視化とか
	- クエリ
		- N+1、スロークエリ、実行計画とか
	- データベース
		- パラメータチューニング
	- サーバー
		- とりあえず puma
	- という区分け
- [Previewing Rails 7 upcoming changes \| Stefan's Blog](https://www.stefanwienert.de/blog/2021/08/24/upcoming-rails-7-changes-active-record/)
	- Rails 7 で導入予定の機能紹介
	- `config.active_record.query_log_tags_enabled` がなにげに良さそう
- [Chris's Wiki :: blog/programming/GoWorkspacesComing](https://utcc.utoronto.ca/~cks/space/blog/programming/GoWorkspacesComing)
	- 1.18 で検討されている workspace mode について
