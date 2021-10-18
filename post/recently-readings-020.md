{"title":"最近読んだもの 20","date":"2021-10-18T00:30:00+09:00","tags":["readings"]}

## 記事

- [Partitioning GitHub’s relational databases to handle scale \| The GitHub Blog](https://github.blog/2021-09-27-partitioning-githubs-relational-databases-scale/)
	- 分離予定のテーブル同士の join やトランザクションを検知する linter をまず導入したのがとても良い
		- 具体的にどう実装したのかは気になる
	- 物理的なテーブルの移動が説明されていた方法で数ミリ秒でできたのは意外だった
	- 今ではクラスタ全体で 120 万クエリ/秒を捌いているそう
- [Coalescing Connections to Improve Network Privacy and Performance](https://blog.cloudflare.com/connection-coalescing-experiments/)
	- Connection Coalescing という、サーバのホスト名が異なっても実態が同じ場合（ip と証明書）に、接続を使い回す仕様がある
	- これを本番に適用してみたという記事
	- 接続が使いまわされることでプライバシーの有用性が確認された
	- 意外にもページロードのパフォーマンスは変わらなかった
- [Google’s State of DevOps 2021 Report: What SREs Need to Know \| Rootly](https://rootly.com/blog/google-s-state-of-devops-2021-report-what-sres-need-to-know)
	- Google の State of DevOps 2021 というレポートの概略
	- devops と sre、devsecops といった分野融合的な考え方の大切さ。計算された中程度のリスクを積極的に取ることのできる文化がある組織はハイパフォーマンス
- [Verica \- Announcing the VOID](https://www.verica.io/blog/announcing-the-void/)
	- 障害レポートのパブリックなデータベースを作っているらしい

## ドキュメント

- [semian/README\.md at master · Shopify/semian](https://github.com/Shopify/semian/blob/master/README.md#understanding-semian)
	- Circuit breaker と bulkheads というアルゴリズムの比較と用途の説明を、Web サーバの busy thread 数 (のようなものの) のグラフを使って説明している部分がとてもわかりやすかった
	- circuit breaker は問題発生時にリクエストを止めて回復させる、bulkheads はチケット制を導入することで busy なスレッド数の数を一定数におさえる
