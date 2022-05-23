{"title":"最近読んだもの 42 - MySQL とその周辺技術の歴史と将来など","date":"2022-05-22T23:00:00+09:00","tags":["readings"]}

- [Long Live MySQL: Good Old MySQL Should Be Rejuvenated \| PingCAP](https://en.pingcap.com/blog/long-live-mysql-good-old-mysql-should-be-rejuvenated/)
    - データベース技術全般 (RDBMS ~ NewSQL) の歴史を踏まえての現状認識 by TiDB 設計者
        - どのようなニーズの変化があり、技術の変化があり、現在なにが求められているか
    - 納得感も高くとてもいい記事
- [Long Live MySQL: Kudos to the Ecosystem Innovators \| PingCAP](https://en.pingcap.com/blog/long-live-mysql-kudos-to-the-ecosystem-innovators/)
    - 上記の続きで、それを踏まえて MySQL 周辺の各技術の解説と比較
    - MariaDB, Percona から Vitess, TiDB まで
    - それぞれ端的に特徴と pros, cons が述べられていて非常にわかりやすかった
    - ポジショントークも入っているだろうし、実際に技術選定をする際はここに記載のない要素（既存資産の活用度、ラーニングカーブ、採用実績など）も検討しないといけないが、とてもわかりやすい素晴らしい記事だった
- [The operational relational schema paradigm](https://planetscale.com/blog/the-operational-relational-schema-paradigm)
    - 理想的な Schema change (migration) の要件リスト
- [Shopify Invests in Research for Ruby at Scale \(2022\)](https://shopify.engineering/shopify-ruby-at-scale-research-investment)
    - Shopify は ruby 関連のアカデミックな研究に $500k/year 規模の出資をしているらしい
- [Lets improve GitOps usability \| Google Cloud Blog](https://cloud.google.com/blog/products/containers-kubernetes/lets-improve-gitops-usability)
    - [GoogleContainerTools/kpt](https://github.com/GoogleContainerTools/kpt) の紹介
    - IaC は良いけれど学習コストも高く職人芸が必要になりがちで、Infrastructure as a Data ということで GUI, CLI, API などから設定を触れるようにできると良いよねというものっぽい
    - Platform Team - Production Team 間のやりとりも最低限にしたいというモチベーション
    - kustomize での知見が活かされているらしい
