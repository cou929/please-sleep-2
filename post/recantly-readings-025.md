{"title":"最近読んだもの 25","date":"2021-11-21T23:30:00+09:00","tags":["readings"]}

## 記事

- [Memorystore for Redis supports Read Replicas \| Google Cloud Blog](https://cloud.google.com/blog/products/databases/memorystore-for-redis-supports-read-replicas)
	- GCP の memorustore redis でリードレプリカを作成できるようになったらしい。嬉しいニュース
	- あわせて Redis6 にバージョンアップすると、IO がマルチスレッド化されるのでさらに早くなるよとのこと
		- 直接関係ないことに言及していて謎だけど、アップグレードを促したいのかな
- [Diving Into Redis 6\.0 \| Redis](https://redis.com/blog/diving-into-redis-6/)
	- 上記のニュースで Redis6 のスレッド対応について知らなかったので読んでみた
	- コアの処理はシングルプロセルのままだけど、IO だけ別スレッドに逃すらしい
		- IO というと、ネットワークへの読み書きが主なのだろうか？
- [Horizontal Pod Autoscaling with Custom Metrics in Kubernetes \| Pixie Labs Blog](https://blog.px.dev/autoscaling-custom-k8s-metric/)
	- カスタムメトリクスで hpa を動かすサンプルの紹介
	- cpu やメモリではなく、ユーザー定義の指標をもとに pod をスケールさせられる
		- こういうところがプラガブルな設計になっているのが k8s ぽさなのかな
	- 本旨からそれるが、horizontal だけでなく vertical の auto scaler もあることや、スケーリングの反応のしやすさを調整するパラメタがあることは知らなかったので参考になった 
- [kube\-lineage: A CLI tool for visualizing Kubernetes object relationships \- tohjustin's blog](https://tohjustin.github.io/posts/2021-11-01-kube-lineage/)
	- k8s のリソースの依存関係を可視化するツール
	- モチベーションから、既存ツールの調査、実装方針の流れが説明されていてよかった
- [Rails feature that you've never heard about: schema cache · Kir Shatrov](https://kirshatrov.com/posts/schema-cache/)
	- Rails は初期化時に [SHOW FULL FIELDS](https://dev.mysql.com/doc/refman/8.0/en/show-columns.html) を発行している
	- 毎回やると重いのでキャッシュもしている

## 記事

- [Horizontal Pod Autoscaler \| Kubernetes](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/)
	- hpa のドキュメントをあらためて読んでみたが参考になるトピックがいくつかあった
	- アルゴリズムの概要
		- 現在の pod 数 * 現在のメトリクス / desired メトリクス
		- あくまで全体の平均などで動作するので、特定の 1 pod だけ高負荷ということが起こり得るが、known limitation
		- よって恐らく pod によって偏りが出るような設計はアンチパターン
	- 急激な増減を防ぐための stabilizationWindow というヒューリスティックな仕組み
	- スケーリングポリシーを設定できる
		- どの程度の pod が増減できるかや、stabilization の期間の変更
	- desired 数を 0 にすることで設定を変えずにオートスケールを止められる。一時的なメンテなどで使えるテクニック
- [MySQL :: MySQL 8\.0 Reference Manual :: 17\.2\.3 Replication Threads](https://dev.mysql.com/doc/refman/8.0/en/replication-implementation-details.html)
	- レプリケーションのための主なスレッド 3 つ
		- primary が bin log を送る `Binary log dump thread`
		- replica が bin log を受け取る `Replication I/O receiver thread`
		- replica が受け取った relay log から適用する `Replication SQL applier thread`
	- replica が受け取る処理と適用する処理が別スレッドに分かれているのがポイント
		- 適用処理が遅かったり一時停止してもスムーズに動作できる
