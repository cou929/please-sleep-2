{"title":"最近読んだもの 44 - Square の Vitess 導入事例、TiKV のコアコンセプトなど","date":"2022-06-05T23:00:00+09:00","tags":["readings"]}

- [Sharding Cash \| Square Corner Blog](https://developer.squareup.com/blog/sharding-cash/)
    - Square の cash での Vitess 導入事例
    - 一連のエントリの親記事
- [Remodeling Cash App Payments \| Square Corner Blog](https://developer.squareup.com/blog/remodeling-cash-app-payments/)
    - シャーディング導入のためにデータモデルを変更
    - かなり大掛かり
- [Cross\-Shard Queries & Lookup Tables \| Square Corner Blog](https://developer.squareup.com/blog/cross-shard-queries-lookup-tables/)
    - Lookup vindex の導入
    - 考えてみれば当たり前だけど、Lookup vindex は必ず別の keyspace にも必要になるので更新はアトミックにできない（2pc などで保障しない限りは）のは考慮が必要
- [The TiKV blog \| Building a Large\-scale Distributed Storage System Based on Raft](https://tikv.org/blog/building-distributed-storage-system-on-raft/)
    - TiKV (TiDB のバックエンドにある分散ストレージ）のコアコンセプト
    - 分散ストレージ実装の要点は、行ってしまえば分散方法とメタデータ管理に尽きるとのこと
    - レンジベースの分散と Raft でのメタデータ管理
- [AddyOsmani\.com \- Software Engineering Insights From 10 Years At Google](https://addyosmani.com/blog/software-eng-10-years/)
    - ポリシーの問題で削除されている模様
    - [Web Archive](https://web.archive.org/web/20220519020040/https://addyosmani.com/blog/software-eng-10-years/)
    - 内容はそうだねという納得感あり、実践するのが大変という感じ
