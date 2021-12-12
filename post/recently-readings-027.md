{"title":"最近読んだもの 27","date":"2021-11-29T01:00:00+09:00","tags":["readings"]}

体調を崩していたので一周飛ばし & 少なめ。

## 記事

- [Property\-Based Testing In Go \- Earthly Blog](https://earthly.dev/blog/property-based-testing/)
	- Go での pbt 事例
	- 例があまり実用的なケースでない
	- Go は標準ライブラリに pbt のパッケージが入っているのは知らなかった。流石
- [Shopify and Google Cloud team up for an epic BFCM weekend \| Google Cloud Blog](https://cloud.google.com/blog/topics/retail/shopify-and-google-cloud-team-up-for-an-epic-bfcm-weekend)
	- shpify の black friday / cyber monday は規模がすごいですよ。そしてそれは gcp でさばいていますよという宣伝
	- [ghostferry](https://shopify.github.io/ghostferry/master/index.html) という内製ツールが面白そうだった
		- mysql から他へのデータ移行を downtime 5 sec で実現
		- shopify オンプレから gcp へのショップ移行で使われているらしい

## ドキュメント

- [MySQL :: MySQL 5\.7 Reference Manual :: 14\.11 InnoDB Row Formats](https://dev.mysql.com/doc/refman/5.7/en/innodb-row-format.html)
	- あまりここを変更することは無さそうだけど、innodb の raw format の説明
	- 関係ないけど clustered index と secondary index の説明はこのページが一番端的でわかりやすい気がした
- [ActiveRecord::ConnectionAdapters::ConnectionPool](https://api.rubyonrails.org/classes/ActiveRecord/ConnectionAdapters/ConnectionPool.html)
	- ActiveRecord::ConnectionAdapters::ConnectionPool のドキュメント
	- スレッドセーフなプールを提供している
	- コネクションの実態は AbstractAdapter
	- コネクションは遅延評価的に必要になってから貼られる
		- プールにコネクションを要求すると、アイドルのものがあればそれを、なくて上限未満なら新規確立、そうでなければ空きを待つ
	- オプション
		- pool: 保持するコネクションの最大数
		- idle_timeout: この秒数以上アイドルだと切断される。デフォルトは 300 秒
		- checkout_timeout: プールから接続を取得する際のタイムアウト。デフォルトは 5 秒
	- stst() で次のような統計値を取得できる
		- `{ size: 15, connections: 1, busy: 1, dead: 0, idle: 0, waiting: 0, checkout_timeout: 5 }`
	- プールからコネクションを取り出すことをチェックアウト、戻すことをチェックインと命名されている
- [Concurrency and Database Connections in Ruby with ActiveRecord \| Heroku Dev Center](https://devcenter.heroku.com/articles/concurrency-and-database-connections)
	- ActiveRecord のコネクションプールの設定についての heroku のドキュメント
	- AR のコネクションプールはプロセスごとに作られ、その中のスレッドで共有される
	- puma のようなスレッド型のサーバーの場合は puma のワーカー数に AR の pool オプションを設定するのがよい
		- pool の値 x プロセス数 x サーバ数 が api 全体から db に貼られるコネクション数になる
	- worker の場合は、例えば sidekiq の場合は sidekiq の concurrency (= ワーカースレッド数) に合わせて pool オプションを設定するとよい
		- api 同様に worker 全体のコネクション数が計算できる
	- api と worker のコネクション数が db の maxConnection を超えないように注意する
	- 超えそうな場合の対策の一つとして、プロキシ型のコネクションプールを提供するミドルウェアの導入がある
		- AR が提供する db へのクライアント側でのプールではなく、クライアントと db の間でプロキシとして動作するタイプのもの。プロセス単位ではなくサーバ単位でコネクションを共有できる
		- postgresql だと PgBouncer が有名
		- プリペアードステートメントなど使えない機能もある
	- コネクションが bad state (= 接続済みだが使えなくなっている接続) の場合、db へのコネクション数が増える
	- そのため pool 値をスレッド数近くの値に揃えておくことを検討する
		- 例えば pool 5 で 1 スレッドのサーバ設定の場合、普段は 1 接続だが最大 5 接続まで増える可能性がある
	- reaping_frequency という bad な接続を定期的に確認して閉じるオプションもある
		- https://github.com/rails/rails/issues/9907 の issue よりデフォルトではオフになっているらしい
