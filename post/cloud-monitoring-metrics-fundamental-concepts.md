{"title":"Cloud Monitoring のメトリクスの基本的なコンセプトをざっと理解する","date":"2021-08-21T23:30:00+09:00","tags":["gcp"]}

GCP 上に構築されたシステムを運用していると日々 Monitoring Dashbord や Metrics Explorer を見ることになると思う。Metrics の設定項目は結構複雑で、いつもなんとなくやりたいことはできているのだが、理解が曖昧で不安だった。

そんな中 Cloud Monitoring のドキュメントに [Concept > Metrics, time series, and resources](https://cloud.google.com/monitoring/api/v3/metrics) という項目があることに気づいた。読んでみると Metrics の基礎的な概念をざっと理解できてよかった。またその概念の GCP 上での呼び名もわかるり、そこからコンソールのこの項目はこのことだったのかという紐付けができ、それも良かった。以下は読んだメモ。

## [Metrics, time series, and resources  \|  Cloud Monitoring  \|  Google Cloud](https://cloud.google.com/monitoring/api/v3/metrics)

Cloud Monitoring に登場する概念はハイレベルから見ると以下の 3 つ。

- Monitored-resource types
	- モニタリング対象のリソース
	- 例えば `gcs_bucket` など
- Metric types
	- 対象のリソースから取得できる指標
	- 例えば gcs_bucket ごとのリクエスト数など
	- 時系列データの種別や各値の型のバリエーションがある
- Time series
	- そのメトリクスの時系列データ
	- 時系列データは timestamp と値のペア

## [Filtering and aggregation: manipulating time series  \|  Cloud Monitoring](https://cloud.google.com/monitoring/api/v3/aggregation)

時系列データはそのままだと膨大な量なので、通常は加工して用いる。加工は大きく Filtering と Aggrigation の 2 ステップで行う。

- Filtering
	- 不要なデータを取り除く
	- 期間指定や、値のしきい値で外れ値を取り除いたり
- Aggregation
	- 複数のデータをそれより少数の代表値に集約する

特に Aggrigation は多種・複雑な設定ができるので覚えることが多い。今回理解を整理したかったのは主にここ。

Aggregation (あるいは summarization) には 2 つの側面がある。

- Alignment
	- データをひとつの時系列に整列させる
- Reduction
	- 複数の時系列を結合する
	- 事前に alignment されている必要がある

### Alignment

あるいは 1 つの時系列内での正則化。ある時系列内のデータポイントはばらばらの間隔で記録されているが、それを一定間隔ごとに整列させる。次のステップで整列をさせる。

- 整列の間隔を指定する
	- 間隔は interval, period, alignment period, alignment window などと呼ばれる
- 指定した interval 内での代表値を計算する
	- 例えば平均値、最大値、最小値など、たくさんの関数がある
		- [Aligner](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.alertPolicies#Aligner) と呼ばれる

Aligner によって、各 interval ごとにひとつの代表値が計算される。こうして正規化された interval ごとに値が並んだ時系列が生成される。

### Reduction

複数の時系列を一つにまとめるステップ。例えば [Memorystore redis の redis/stats/cpu_utilization という metrics](https://cloud.google.com/monitoring/api/metrics_gcp#redis/stats/cpu_utilization) は user, sys それぞれの cpu 使用率を記録しているが、これを総合した cpu 使用率を見たいときなどがユースケースだと思う。前提として Alignemt された時系列でないと Reduction できない。

数値をまとめる関数は [Reducer](https://cloud.google.com/monitoring/api/v3/aggregation) と呼ばれる。内容は Aligner と対になっている。(ここは同じ関数が別名で二度登場するので、所見で Metrics Explorer を見て混乱するポイントだと思う)

複数の系列をまとめる際、Grouping で指定した軸ごとにまとめることができる。指定しなければすべての系列が一つの系列にまとめられる。例えばすべての Pod のメトリクスの系列を production cluster と staging cluster でグルーピングし 2 系列にまとめる、といったユースケースが考えられる。

## [Value types and metric kinds  \|  Cloud Monitoring  \|  Google Cloud](https://cloud.google.com/monitoring/api/v3/kinds-and-types)

- Value types (各データの値の型)
	- BOOL, INT64, DOUBLE, STRING などの基本的な型と [Distribution](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/TypedValue#Distribution) がある
- Metric kind (指標の種別)
	- gauge
		- 単純な値
	- delta
		- 前の値からの差分
	- cumulative
		- 累積
- [有効な組み合わせとそうでないものがある](https://cloud.google.com/monitoring/api/v3/kinds-and-types)
	- 例えば BOOL の delta はありえない

## [Retention and latency of metric data  \|  Cloud Monitoring  \|  Google Cloud](https://cloud.google.com/monitoring/api/v3/latency-n-retention)

- Data retension (保持期間) は以下で定義されている
	- [Quotas and limits  \|  Cloud Monitoring  \|  Google Cloud](https://cloud.google.com/monitoring/quotas#data_retention_policy)
	- たいていのものは 6 週間
- Latency (イベントが発生してから参照可能になるまでの時間)
	- [各指標のドキュメント](https://cloud.google.com/monitoring/api/metrics) に `Sampled every 60 seconds. After sampling, data is not visible for up to 240 seconds.` と言った説明がある

## PR

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873119014/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://images-na.ssl-images-amazon.com/images/I/51wecBhtIOL._SX389_BO1,204,203,200_.jpg" alt="Kubernetesで実践するクラウドネイティブDevOps" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873119014/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Kubernetesで実践するクラウドネイティブDevOps</a></div><div class="amazlet-detail">John Arundel  (著), Justin Domingus (著), 須田 一輝 (監修), 渡邉 了介 (翻訳)<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873119014/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

