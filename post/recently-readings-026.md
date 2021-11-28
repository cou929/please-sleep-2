{"title":"最近読んだもの 26","date":"2021-11-29T01:00:00+09:00","tags":["readings"]}

## 記事

- [Amazon MemoryDB for Redis – Where speed meets consistency](https://www.allthingsdistributed.com/2021/11/amazon-memorydb-for-redis-speed-consistency.html)
	- MemoryDB for Redis の設計の紹介
	- Redis は高負荷環境では一貫性が保てない場合がある (weak consistency)
		- 主にはレプリケーションが非同期、つまり primary でコミットした後に replica に送られるので（レプリケーションが非同期）
		- その分書き込み時のレスポンスが速い
	- MemoryDB は primary でのディスク上の transaction log (WAL 的な物) への sync が終わった時点でそのコマンドが committed となる
		- このときスループットを維持するためにクライアントからのコマンドは受け付けて、sync が終わるまでレスポンスは待機させる
			- よって通常の Redis に比べレイテンシーは下がる
		- レプリケーションは transaction log ベースで行われる
	- ハイレベルではRDBMS ぽい構成といえる？
		- これを Aurora のようにプラントのエンジンとしての Redis と、パックエンドの永続化・冗長化層を切り離して独自実装することで実現している
		- Aurora も Spanner もそうだけどインターフェースは既存のプロダクト互換を提供して背後にこのようなレイヤーを置くのが今よく売れる製品の一スタンダードなのかな
	- レプリケーションが非同期なことの問題は GCP の Memorystore ではどうなってるんだろ
		- 特に先週読んだ [replica サポートの新リリース](https://cloud.google.com/blog/products/databases/memorystore-for-redis-supports-read-replicas)はどうなるんだろう
- [Planet MySQL :: Planet MySQL \- Archives \- Externally Stored Fields in InnoDB](https://planet.mysql.com/entry/?id=698268)
	- Innodb にて varchar や text 型のデータは、一定のサイズを超えるとクラスタードインデックスに直接ではなく、別の場所に保存される。その場合クラスタードインデックスにはその別箇所へのポインタが入る
	- 前提として雑にいうとInnodb はクラスタードインデックスに全てのデータを保持している（≒ pk をキーにしたインデックスに値としてそのレコードの全データが入っている）
        - [MySQL :: MySQL 5\.7 Reference Manual :: 14\.6\.2\.1 Clustered and Secondary Indexes](https://dev.mysql.com/doc/refman/5.7/en/innodb-index-types.html)
	- innodb では varchar や text や varbinary などは共通してこの仕組みらしい
	- どのくらいのサイズで別箇所保存になるかは row format に依存する
- [Introducing Relational Database Connectors](https://blog.cloudflare.com/relational-database-connectors/)
	- Cloudflare の worker から rdbms に接続できるようになった
	- もともと worker のランタイムは tcp socket をサポートしていないので、db のドライバーからは tcp socket に見えるが実際は websocket で通信するような仕組みを提供して、既存のライブラリがほとんどそのまま使えるようにしたらしい
	- すごそう
- [What you need to know about cluster logging in Kubernetes \| Opensource\.com](https://opensource.com/article/21/11/cluster-logging-kubernetes)
	- k8s でログをどのようにログバッエンドに送るかのパターン
	- ds などで node ごとの常駐 pod を作成、サイドカーとして pod 内に常駐、アプリケーションから直接バックエンドに送る
	- 後者ほどフレキシブル
