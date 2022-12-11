{"title":"最近読んだもの 54 - PlanetSclae Boost, 複合主キーで locality 改善など","date":"2022-12-11T23:30:00+09:00","tags":["readings"]}

- [How to Introduce Composite Primary Keys in Rails](https://shopify.engineering/how-to-introduce-composite-primary-keys-in-rails)
    - 複合主キーにすることでデータの locality を高めて最適化する話
    - shopify の場合マルチテナントなので、テナントの ID ともともとの主キー (id) の複合 PK にした
        - 殆どのクエリはあるテナント内での検索なので、それらのレコードが近い場所に配置されていたほうが効率的
    - Rails で `id` 以外の主キーを扱うためのテクニック
        - 複合 PK とは別に `KEY id (id)` を足す
    - 読み取りは 5-6x 高速化したが、書き込みが 10x 遅くなった
    - SSD でもここまで違いが出るものなのか気になる
- [37signals Dev — Faster pagination in HEY](https://dev.37signals.com/faster-paging-in-hey/)
    - 上記の shopify のテクニックをページングをしているクエリに応用した話
    - 複合 PK で locality を改善したことに加えて、データフェッチをなくした (`Using where` => `Using index` にした) のもなるほどだった
        - もとのクエリは、例えば 400 レコードフェッチしてソートし、そのうち 30 件を返していた
        - 改善後のクエリはソート後に id だけを返すようにし、その後返された 30 件を改めて別クエリで SELECT するようにした
        - 改善後のソートはカバリングインデックスになるようにした
        - データのフェッチ回数が 400 => 30 に減って高速化
    - こういう話は実行計画だけでなく CREATE TABLE も載せてくれないとわかりにくいなと思った
- [How PlanetScale Boost serves your SQL queries instantly](https://planetscale.com/blog/how-planetscale-boost-serves-your-sql-queries-instantly)
    - PlanetScale Boost というマテリアライズドビューのようなものを提供する新機能の内側
    - 元のテーブルから最終的なビューまでの間の中間状態をそれぞれ持っておく
    - マテリアライズドビューとは違い全レコードのビューを事前計算しないのでリソース効率がよいが、その分キャッシュミスがありえる
    - キャッシュミスがあったとしても、計算が終わっている中間状態をうまく利用してくれるので、フルで元テーブルに重いクエリが向くことは少ない
    - `partially materialized views` という方式らしく、根拠となる論文は [Noria: dynamic, partially\-stateful data\-flow for high\-performance web applications \| USENIX](https://www.usenix.org/conference/osdi18/presentation/gjengset) らしい
- [Vitess \| Scalable\. Reliable\. MySQL\-compatible\. Cloud\-native\. Database\.](https://vitess.io/blog/2021-11-02-why-write-new-planner/)
    - 現在のデフォルトである Gen4 Planner に書き換えた時の話
    - はじめはサイドプロジェクトとして始まったというのがかっこいい
- [Enabling static analysis of SQL queries at Meta \-](https://engineering.fb.com/2022/11/30/data-infrastructure/static-analysis-sql-queries/)
    - フロントエンドからの SQL を統一してパースする UPM というレイヤーがある
    - パース結果は次のように使われる
        - lint や rewrite
        - ユーザー定義型の型チェック
            - カラムには integer といった物理的な型の他に、ミリ秒・ナノ秒といったユーザー定義型も、この仕組で付与できるようになる
        - data lineage
            - あるカラムがどこから更新されたか、その連鎖を可視化して、データ不整合の検知や修正、リファクタリングをしやすくするための活動
            - このような概念を data lineage と呼ぶことが知れたことも、個人的にかなり収穫
        - パース結果の抽象構文木を直接バックエンド (presto や spark) に渡すこともできるらしい

## PR

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B09N5NWKR1/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/4151TBrJq1L.jpg" alt="Efficient MySQL Performance (English Edition)" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B09N5NWKR1/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Efficient MySQL Performance (English Edition)</a></div><div class="amazlet-detail">English Edition  by Daniel Nichter  (著)  Format: Kindle Edition<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B09N5NWKR1/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
