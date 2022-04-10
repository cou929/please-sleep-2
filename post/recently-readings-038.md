{"title":"最近読んだもの 38 - MySQL スケール、GitHub 障害など","date":"2022-03-27T23:30:00+09:00","tags":["readings"]}

- [Why Don't You Use \.\.\.](https://www.brendangregg.com/blog/2022-03-19/why-dont-you-use.html)
    - 大きな tech company に対して「なぜ技術 X を使わないのか」という質問をするのがナンセンスな理由
    - 言えない事情もあるし、言うメリットもない事が多い
    - そういう質問が多くて辟易しているのかな
- [An update on recent service disruptions \| The GitHub Blog](https://github.blog/2022-03-23-an-update-on-recent-service-disruptions/)
    - まだ原因はわかっていないらしい
    - 発生したら primary をフェイルオーバーしつつ、ログを仕込みつつ、スケールアップも進めている
    - お疲れさまです 🙏
- [Some benefits of simple software architectures \- Wave Blog](https://www.wave.com/en/blog/simple-architecture/)
    - シンプルなモノリスのアーキテクチャで世界トップ100のトラフィックのサービスを運用できてるよという話
    - とはいえ k8s, graphql, 独自の通信プロトコルなど必要なところで複雑なテクノロジーも取り入れていて良い
- [February service disruptions post\-incident analysis \| The GitHub Blog](https://github.blog/2020-03-26-february-service-disruptions-post-incident-analysis/)
    - こちらは 2020 年の障害の記事
    - 当時から mysql1 の運用が大変そう
- [MySQL 5\.7 read\-write benchmarks \- Percona Database Performance Blog](https://www.percona.com/blog/2016/05/17/mysql-5-7-read-write-benchmarks/)
- [Fixing MySQL scalability problems with ProxySQL or thread pool \- Percona Database Performance Blog](https://www.percona.com/blog/2016/05/19/fixing-mysql-scalability-problems-proxysql-thread-pool/)
    - mysql で接続数を変えながら read, write を繰り返しスループットをはかるベンチマーク
    - mysql5.7 で接続数が一定を超えると性能が劣化する
    - ProxySQL などを挟んで接続数を調整すると劣化が抑えられる
- [Scaling MySQL stack, ep\. 4: Proxies · Kir Shatrov](https://kirshatrov.com/posts/scaling-mysql-stack-part-4-proxy/)
    - 目安として 1k 接続をこえたら proxy 導入を検討すると良いとの記載