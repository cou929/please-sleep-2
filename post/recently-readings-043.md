{"title":"最近読んだもの 43 - Vitess の Transaction, Gitee の検閲など","date":"2022-05-29T23:00:00+09:00","tags":["readings"]}

- [SREcon 2022 Americas Wrap Up](https://www.blameless.com/blog/srecon-2022-americas-wrap-up)
    - SRECon2022 の wrap up
- [Vitess \| A database clustering system for horizontal scaling of MySQL](https://vitess.io/blog/2016-06-07-distributed-transactions-in-vitess/)
    - Vitess の 2PC について
- [ACID Transactions are not just for banks \- the Vitess approach](https://planetscale.com/blog/acid-transactions-are-not-just-for-banks-vitess-approach)
    - redolog + binary log と semi-synchronous replication で冗長性を確保
- [Abstracting Sharding with Vitess and Distributed Deadlocks \| Square Corner Blog](https://developer.squareup.com/blog/abstracting-sharding-with-vitess-and-distributed-deadlocks/)
    - 複数 shard にまたがるトランザクション間でのデッドロック
    - 基本的な内容だけど気づきづらそう
- [Gitee, China’s answer to GitHub, to review all code by temporarily closing open\-source projects to the public \| South China Morning Post](https://www.scmp.com/tech/big-tech/article/3178323/gitee-chinas-answer-github-review-all-code-temporarily-closing-open)
    - 無料部分しか読めていないが、中国の検閲強化のため gitee という中国版 GitHub で公開リポジトリが全て非公開になり再公開にはレビューが必要になったとのこと
