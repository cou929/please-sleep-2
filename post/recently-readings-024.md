{"title":"最近読んだもの 24","date":"2021-11-14T21:00:00+09:00","tags":["readings"]}

## 記事

- [Debugging Action Callbacks \(aka Filters\) in Rails \| Hashrocket](https://hashrocket.com/blog/posts/debugging-action-callbacks-aka-filters-in-rails)
	- Rails のフィルタ、コールバックをデバッグするには、デバッガを起動して _process_action_callbacks を見るといいらしい
	- 筆者の恨み節も伝わってくる。気持ちはわかる
- [Dockershim removal is coming\. Are you ready? \| Kubernetes](https://kubernetes.io/blog/2021/11/12/are-you-ready-for-dockershim-removal/)
	- 事前の準備状況のサーベイの締め切りがあと一ヶ月というアナウンス
- [Faster debugging with traces and logs together \| Google Cloud Blog](https://cloud.google.com/blog/products/operations/faster-debugging-with-traces-and-logs-together)
	- 各システムで対応する必要があるが、GCP の Logging から各システムのレイテンシなどのトレースが見られるようになるらしい

## ドキュメント

- [General best practices  \|  Cloud SQL for MySQL  \|  Google Cloud](https://cloud.google.com/sql/docs/mysql/best-practices), [Operational guidelines for MySQL instances  \|  Cloud SQL for MySQL](https://cloud.google.com/sql/docs/mysql/operational-guidelines), [Managing database connections  \|  Cloud SQL for MySQL  \|  Google Cloud](https://cloud.google.com/sql/docs/mysql/manage-connections#ruby)
	- Cloud SQL (MySQL) の運用に関して、こういう使い方をすると SLA の範囲外になるというガイダンス
		- ある意味免責事項?
		- 設定の自由度を提供する代わりにこういうことをするとサポートしきれんよみたいな話
	- 前提単独 zone 構成は SLA 対象外だったり、1 インスタンスに 10,000 テーブルまでとか、アプリケーションも接続プールと exponential backoff での接続リトライが推奨とか
	- 一部 MySQL を選択しても PostgreSQL の内容が出てきて困った
- [What is OpenTelemetry?  \|  Google Cloud](https://cloud.google.com/learn/what-is-opentelemetry)
	- 上記の traces の記事に関連して
	- アプリケーションやそれが使用するサービスから、各種監視システムに送信するメトリクスについてのオープンな規格
	- 確かに、コンピューティングリソースの標準化が進んでいる中で、メトリクスも標準化しないと目的は達成できないと思うので、なるほど感があった
