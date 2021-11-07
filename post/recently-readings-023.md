{"title":"最近読んだもの 23","date":"2021-11-08T01:45:00+09:00","tags":["readings"]}

## 記事

- [Static Analysis Fatigue – Embedded in Academia](https://blog.regehr.org/archives/259)
    - 2010 年の記事だが、当時静的解析ツールで言語処理系など有名 OSS プロジェクトのコードを解析しバグ報告をしていたが、一部で疎まれたこと
    - falase positive かどうかの確認がされていなかったりなど actionable でない報告内容だったり、そのような報告が当時多く S/N 比が非常に低かったりという原因だったそう
    - 当時の状況や実際の報告方法をみてみないとなんとも言えないが、ありそうな話だなとは思う
- [Facebook to stop using facial recognition, delete data on over 1 billion people \| Ars Technica](https://arstechnica.com/tech-policy/2021/11/after-tagging-people-for-10-years-facebook-to-stop-most-uses-of-facial-recognition/)
    - fb の顔認識による自動タグ付機能を廃止し過去データも削除するとのこと
    - 文脈としては、まず最近の fb の社会への悪影響 (米大統領戦におけるフェイクニュース拡散、ミャンマーにおける人種差別助長など) により批判が高まっている
    - fb の機能に限らず顔認識一般の問題点として、白人男性以外は認識率に劣るため (訓練データのバイアスによるもの)、誤認識による誤認逮捕が実際に発生している。また中国などは顔認識を人種差別的な政策のために利用している
    - fb の機能がオプトアウト方式なのが EU の一部の国では批判されていたり、個人のバイオメトリクス情報の取り扱いに厳密な州では訴訟が起こされており、実際に fb が和解金を支払う事例も発生しており、おそらく経済的にも割に合わなくなってきている
    - "統計的手法による自動タグ付け" は、スタート時点では素朴に良い機能 (ユーザーメリットも事業メリットもあり、技術でレバレッジしている感もよい) 機能にしか見えないと思うし、自分だってそう考えると思う。それがこのような帰結になるのは恐ろしい。自分だったとしたらこれを避けようがなかったと思うし、近い業界での出来事でもあるので他人事とも思いづらい
- [Introducing container image streaming in GKE \| Google Cloud Blog](https://cloud.google.com/blog/products/containers-kubernetes/introducing-container-image-streaming-in-gke)
    - gke で pod をスケールアウトする際に image を pull する部分が高速化するらしい
    - 通常は image 全体を足元のディスクにダウンロードして、完了したらアプリを起動する
    - streaming では、基本的な考え方としては必要な部分だけ転送が終わればアプリを動かし始めるらしい
    - その際に色々なレイヤーでのキャッシュ戦略も最適化
    - 特に pod 立ち上げ当初はイメージをネットワークからストリーミングで取得しアプリを動かし始める。並行して通常の pull も行っておき、あとで入れ替えるらしい
- [Containers vs\. Pods \- Taking a Deeper Look \- Ivan Velichko](https://iximiuz.com/en/posts/containers-vs-pods/)
    - コンテナと pod の cgroups, namespaces を実際に確認しそれぞれどのように構成しているかをしらべたもの
    - 読みやすくまとまっていてわかりやすい
- [Martin Heinz \| Keeping Kubernetes Clusters Clean and Tidy](https://martinheinz.dev/blog/60)
    - k8s クラスタに残ったゴミリソースのお掃除
    - LimitRange や ResourceQuota といったリソースで上限設定
        - 1.21 で job リソースの ttl が beta にくるらしい。よさげ
    - すでにあるリソースの削除 (マニュアル、自動 (kube-janitor))
    - 監視
        - etcd のサイズやオブジェクト数、クラスタ全体の cpu, mem 使用量合計、ネームスペースごと、ノードごと
    - 包括的で良かった
- [Google BigQuery table snapshots for data backups \| Google Cloud Blog](https://cloud.google.com/blog/products/data-analytics/google-bigquery-table-snapshots-for-data-backups)
    - bq のテーブルのスナップショットを取れるようになったらしい

## ドキュメント

- [Use preemptible VMs to run fault\-tolerant workloads](https://cloud.google.com/kubernetes-engine/docs/how-to/preemptible-vms)
    - 1.20 以降は [graceful node shutdown](https://kubernetes.io/docs/concepts/architecture/nodes/#graceful-node-shutdown) が有効になる
    - 最大 24 時間で停止する
    - 停止の 30 秒まえに [通知](https://cloud.google.com/compute/docs/instances/preemptible#preemption) が来る
    - nodeSelector や taints, tolerations で配置される node を制御する
