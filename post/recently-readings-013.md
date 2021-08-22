{"title":"最近読んだもの 13","date":"2021-08-22T23:50:00+09:00","tags":["readings"]}

## 記事

- [A future for SQL on the web](https://jlongster.com/future-sql-web)
	- wasm でブラウザ上で sqlite を動かす sql.js が以前からあったが、永続化は対応していなかった
	- sqlite のストレージレイヤーとして IndexedDB を使うのとでこれを解決した
	- IndexedDB はパフォーマンスがとても悪いが、数倍オーダーで改善した
		- 通常の環境で sqlite を動かすよりもおそいが、素の IndexedDB よりは早い
		- 仕組みとしては読み書きをバッファリンクしている感じっぽい
		- なお [Storage Foundation API](https://web.dev/storage-foundation/) が準備できればこうしたハックなしで改善する予定
	- ストレージが遅くてもそれより上のレイヤーの工夫で数倍オーダーの性能改善ができるのはなるほどだった
- [Tidying up the Go web experience \- go\.dev](https://go.dev/blog/tidy-web)
	- go.dev に集約していきますとのこと
- [The Incident Review: 4 Times When Typos Brought Down Critical Systems \| Rootly](https://rootly.io/blog/the-incident-review-4-times-when-typos-brought-down-critical-systems)
	- タイポ、というよりも設定ミス、が重大な障害につながった例
	- 回復速度をあげるための日々の積み重ねや、やユーザーコミュニケーションの透明性、バリデーションの重要性と難しさ、など
- [3 Code Metrics Every Developer Should Know \| by Miloš Živković \| Level Up Coding](https://levelup.gitconnected.com/3-code-metrics-every-great-developer-must-measure-499b0b2b31ad)
	- "悪いコード" を検知するための評価指標
	- code churning, ABC metric, cyclomatic complexity
- [Optimizing Cloud Native Development Workflows: Combining Skaffold, Telepresence, and ArgoCD \| by Daniel Bryant \| Ambassador Labs](https://blog.getambassador.io/optimizing-cloud-native-development-workflows-combining-skaffold-telepresence-and-argocd-8774d12bf22f)
	- Opinionated なワークフローとして Skaffold, Telepresence, ArgoCD を用いたものを紹介
	- 一部はマイクロサービス故に必要性が高いものもある気がする
	- k8s そのものもそうだが、こうした周囲のツールにもちゃんと学習コストを投資して行かないと、生産性をあげらないかもなと思うなど

## ドキュメント

- [Metrics, time series, and resources  \|  Cloud Monitoring  \|  Google Cloud](https://cloud.google.com/monitoring/api/v3/metrics)
	- 以下にまとめた
	- [Cloud Monitoring のメトリクスの基本的なコンセプトをざっと理解する \- Please Sleep](https://please-sleep.cou929.nu/cloud-monitoring-metrics-fundamental-concepts.html)

## 本

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B08BNNY877/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/412uOzq6rDL.jpg" alt="村上春樹 雑文集（新潮文庫）" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B08BNNY877/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">村上春樹 雑文集（新潮文庫）</a></div><div class="amazlet-detail">村上春樹  (著)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B08BNNY877/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

- 半年ほど、ほぼ技術周辺のことだけに時間を使ってきたが、さすがにインプットが枯渇して来た感があるので、意識的に仕事から遠い本を読んだ
- なにそれという感想だが、やっぱり文章が段違いにうまくて、読めてよかった
    - まずリズムがいいのと、何かを説明する際にその根拠として思いもつかないような角度の物事を取り出してくる
- 不思議と平成的な時代感覚があった
    - オウムの話があったからかもしれないけど

