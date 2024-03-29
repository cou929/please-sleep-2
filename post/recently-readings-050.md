{"title":"最近読んだもの 50 - Delivery Lead Time の実践、外部サービス障害への対策など","date":"2022-08-21T23:30:00+09:00","tags":["readings"]}

- [Delivery Lead Time In Practice](https://isthisit.nz/posts/2022/delivery-lead-time-in-practice/)
    - Delivery lead time (そのコミットが本番環境に乗るまでの時間) という指標について
    - 開発生産性を測る指標として Accelerated でも紹介されているもの
    - 定量化しにくい部分なので、数字が一人歩きしないために気をつけないといけないポイントが多い
    - Dora という標準？も知らなかったし、このようなソフトウェア開発のより社会学的な方面の研究も思ったよりも進んでいるんだなと少し驚いた
- [Handling third\-party provider outages \| incident\.io \| incident\.io](https://incident.io/blog/third-party-outages)
    - AWS, GCP, Cloudfrare など依存しているインフラがダウンした時にどう向き合うか
    - できることは少ないがコントローラブルな部分に集中するしかない
        - マルチリージョンにする。どこが正常でどこが異常なのか切り分けるなど
    - マルチクラウドは多くの場合悪い選択
        - それぞれのサービスに共通する機能しか使えない
        - 例えば LB などはそれぞれのサービスの外側に作る必要があり、かつそれぞれのサービスと同等以上の可用性が求められる
        - 結果としてシングルクラウドでの運用より可用性がさがることの方が多い
- [How to Handle Kubernetes Health Checks](https://doordash.engineering/2022/08/09/how-to-handle-kubernetes-health-checks/)
    - Spring bootデフォルトの？ヘルスチェックを使っていたところ、その api とは無関係の Redis に問い合わせるケースがあり、その Redis が不調な際にサービス全体が過負荷になる障害につながった
    - また readiness, liveness probe のエンドポイントを特別扱いしログやトレースを取っておらず、調査に時間がかかったので、他のエンドポイント同様に情報をとった方が良い
- [Understanding basic networking in GKE \- Networking basics \| Google Cloud Blog](https://cloud.google.com/blog/topics/developers-practitioners/understanding-basic-networking-gke-networking-basics)
    - GKE のネットワークについての駆け足の紹介
    - 登場人物紹介といった感じ
- [Top SRE Interview Questions You Should Know](https://www.blameless.com/sre/sre-interview-questions)
    - SRE 職の面接で確認するべき事項
    - 抽象度高め
- [Ruby on Rails: How we disable our Sidekiq jobs](https://planetscale.com/blog/ruby-on-rails-how-we-disable-our-sidekiq-jobs)
    - flipper という gem を利用して、フィーチャーフラグに応じてジョブをキューに入れるかを制御
