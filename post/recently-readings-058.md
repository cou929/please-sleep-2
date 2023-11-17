{"title":"最近読んだもの 58 - SQLite の過去現在未来など","date":"2023-04-23T23:20:00+09:00","tags":["readings"]}

- [SQLite: Past, Present, and Future](https://www.vldb.org/pvldb/vol15/p3535-gaffney.pdf)
- [The tail at scale](https://www.barroso.org/publications/TheTailAtScale.pdf)
- [How Big Tech Runs Tech Projects and the Curious Absence of Scrum \- The Pragmatic Engineer](https://blog.pragmaticengineer.com/project-management-at-big-tech/)
- [MySQL Data Caching Efficiency](https://www.percona.com/blog/mysql-data-caching-efficiency/)
    - `Innodb_buffer_pool_read_requests` と `Innodb_buffer_pool_reads` からいわゆるバッファプールのキャッシュヒット率を計算できる
- [GNU/Linux shell related internals \| Viacheslav Biriukov](https://biriukov.dev/docs/fd-pipe-session-terminal/0-sre-should-know-about-gnu-linux-shell-related-internals-file-descriptors-pipes-terminals-user-sessions-process-groups-and-daemons/)
    - File descripter, pipe, process groups, terminal といった話題を解説する一連の記事
    - リチャードスティーブンスの APUE のファイルに関する章のような内容を説明しているが、より現代の視点から書かれているのと、説明のテンポが良く読みやすいのがいい
- [MySQL 5\.7 Upgrade Issue: Reserved Words](https://www.percona.com/blog/mysql-5-7-upgrade-issue-reserved-words/)
    - MySQL 8 系にアップグレードすると、予約語が増える
    - `rank`, `system`, `skip`, `lead` などは引っかかりやすそう
- [COMMIT Latency: Aurora vs\. RDS MySQL 8\.0](https://hackmysql.com/post/commit-latency-aurora-vs-rds-mysql-8.0/)
    - `SELECT` や `UPDATE` はおおむねメモリ上での操作ナノに対して、`COMMIT` はディスクへの書き込みを伴う
    - Aurora と RDS で `COMMIT` のレイテンシを調査したところ、Aurora のほうが安定して良い結果だった
- [MySQL :: What is the "\(scanning\)" variant of a loose index scan?](https://dev.mysql.com/blog-archive/what-is-the-scanning-variant-of-a-loose-index-scan/)
    - `GROUP BY` などでインデックスを走査する際に、インデックスを全走査せず不要な部分を飛ばす最適化をした際に、実行計画に loose index scan と表示される
    - 例えば `GROUP BY` でグループごとの最小値を集計する場合、インデックスのソート順がうまく使える場合は、GROUP BY の各キーの最初の値だけを見てそれ移行は飛ばすことができる
- [Announcing Blip: A New MySQL Monitor](https://hackmysql.com/post/announcing-blip-mysql-monitor/)
    - 新しい MySQL のモニタリングツール
- [How I became a machine learning practitioner](https://blog.gregbrockman.com/how-i-became-a-machine-learning-practitioner?ref=matt-rickard.com)
    - OpenAI の共同創業者で元 Stripe の CTO である Greg Brockman が、OpenAI に参画してから、自分の専門性とは異なる AI 分野に対して試行錯誤した記録
    - このレベルの人でも泥臭く失敗しながらやっていることが綴られていて興味深い
- [Connection pooling in Vitess](https://planetscale.com/blog/connection-pooling)
    - Vitess のコネクションプールについて、その意義、セッション単位の設定とコネクションプールの相性の悪さ、それを回避するための各種方法など
    - セッション単位の設定があるクエリでコネクションプールの旨味をできるだけ活かすために、書き換えられるものは SET_VAR に書き換えて、それが難しければ専用の接続をそのクライアント用に確保する
    - Vitess15 からは `Settings Pool` という新しい仕組みが導入された。特定の設定に変更されたセッションをプールに保持するというアプローチ
- [Comparisons of Proxies for MySQL](https://www.percona.com/blog/comparisons-of-proxies-for-mysql/)
    - MySQL Proxy の比較
    - 性能は haproxy > Porxy SQL だが、機能的には後者が多機能なので使い分け
    - MySQLRouter はかなり悪い結果となっている
- [A Healthy Bundle](https://thoughtbot.com/blog/a-healthy-bundle)
    - Gemfile などパッケージ管理システムで定義するライブラリのバージョン指定は、特定のバージョン固定やメジャーバージョンだけ指定するなど柔軟にできるが、どのような方針で指定すべきか
    - 例えば Rails といった影響範囲が大きいものは細かくバージョンを指定する、まだ安定していないライブラリ (後方互換性を壊す変更が入ったり、ブレイキングチェンジがまだまだあり得るもの) はメジャーバージョン・マイナーバージョンまで固定する悲観的な指定を、安定していて後方互換性を壊す変更が入ることもないものはバージョン指定をなくすなど楽観的な指定をする