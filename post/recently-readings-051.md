{"title":"最近読んだもの 51 - Critical Depndency と可用性、マルチクラウドへの批判、Trilogy など","date":"2022-09-04T23:30:00+09:00","tags":["readings"]}

- [The Calculus of Service Availability \- ACM Queue](https://queue.acm.org/detail.cfm?id=3096459)
    - IaaS ベンダーなどクリティカルな外部依存と障害について
    - 経験則: 外部依存の可用性には目標の一桁上が求められる。99.99% を目指すには外部依存システムは 99.999% 必要
    - その前提で SLO を達成するのひ必要な障害時の対応速度を具体的に計算してみると、数分以内に検知・復旧する必要があり、想像以上に余裕がなく、自動化が必須であることがわかる
    - 対応のベストプラクティス集
        - システム全体を shard に分けて、部分的な障害が全体に波及しないようにする
            - その際の SLO/SLI の計算も、例えば 3 shard のうち一つがダウンしたケースではその期間の三分の一のエラーバジェットが消費されたという考え方になる
        - 依存する各コンポーネントについて、それがダウンした際に、システム全体としてはデグレードして動作するようにできないか、ひとつひとつ設計時に考慮する
        - Google では目標の 10 倍の負荷に耐えられる程度の設計にしておくことが目安とされている
        - Google 規模になると、同じ API を 3 箇所にデプロイしておきクライアントからは 3 並列でリクエストし最初に届いたレスポンスを採用する (そうすることで可用性だけでなくレイテンシの特異値も排除できる) とか、フェイルオーバー・ロールバックの自動化 (合意の仕組みなどもあるっぽい?) もされているらしい
- [Multi\-Cloud is the Worst Practice \- Last Week in AWS Blog](https://www.lastweekinaws.com/blog/multi-cloud-is-the-worst-practice/)
    - いわゆる multi cloud は多くの場合良くない選択だよという話
    - 各 IaaS に共通する部分以外、例えば LB を始め監視や開発環境系など様々、は自前運用が必要
    - 1 ベンダーあたりの使用量は減るのでコスト的にも不利など
    - そもそも事例が世の中に少ない。少ないのにはそれなりの理由がある
    - 一見良さそうに聞こえるアイデアだが、実際は一つの IaaS で複数リージョンに分散させるなどの方が現実的
- [No observability without theory – Dan Slimmon](https://blog.danslimmon.com/2019/05/03/no-observability-without-theory/)
    - observability には data だけでなく theory も大事という話
    - theory とはつまりデータを文脈に沿って解釈し意味を見出すというようなことっぽい
    - 確かにインフラのダッシュボードを初見で見てもうまく解釈できないが、毎日ウォッチすることで見方がわかってくる
    - こうした「解釈」は職人芸になりやすいので、ドキュメントなどで属人性をなくす努力も大事とのこと
- [Introducing Trilogy: a new database adapter for Ruby on Rails \| The GitHub Blog](https://github.blog/2022-08-25-introducing-trilogy-a-new-database-adapter-for-ruby-on-rails/)
    - GitHub 内製の ruby mysql adapter
    - mysql2 の後継でパフォーマンスに優れるらしい
    - blocking IO やメモリ管理に工夫がされていて、また libmysqlclient や openssl への依存がないらしい
- [Ditching Active Record Callbacks](https://engineeringblog.wonolo.com/ditching-active-record-callbacks)
    - ActiveRecord callback の問題点と解決策
    - 例えば save の場合、保存用メソッドをひとつ準備して、コールバックにある処理はそこから明示的に呼び出す
    - 個人的にはこのブログの主張に賛成