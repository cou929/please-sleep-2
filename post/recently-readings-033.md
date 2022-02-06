{"title":"最近読んだもの 33","date":"2022-02-06T23:30:00+09:00","tags":["readings"]}

## 記事

- [Fixing Performance Regressions Before they Happen \| by Netflix Technology Blog \| Jan, 2022 \| Netflix TechBlog](https://netflixtechblog.com/fixing-performance-regressions-before-they-happen-eab2602b86fe)
	- netflix の TV クライアント開発のパフォーマンスモニタリングでの偽陽性を減らす話
	- ここでいうパフォーマンスとはクライアントのメモリ使用量とレイテンシのことだが、扱っている話題はモニタリング一般に適用できる内容
	- 最初は静的な閾値を設けていたが、誤検知が多いので、anomaly detectionと changelings の検出を導入して改善した
	- 前者は過去数回分の試行の平均からある標準偏差以上ずれたらアノマリーと見做して報告する
	- 後者は [e-divisive](https://arxiv.org/pdf/1306.4933.pdf) という方法で傾向が変わったポイントを見つけることができるらしい
	- こうした時系列データの統計処理が本論だが、その前提として、PR ごとにパフォーマンステストを走らせていることや、それをかなりのバリエーションの実機、仮想端末で実行していること、そのインフラがそもそもすごい
	- そのうえで、実機は電源やネットワーク状況などのノイズが多いので試行回数を増やすことと、こうした例にはじまる統計処理を施して結果を解析するのが大事そう
- [Google Testing Blog: A Tale of Two Features](https://testing.googleblog.com/2022/02/a-tale-of-two-features.html)
	- App Engine で過去に発生したデータ消失バグの振り返り
	- わかってしまえば明確な原因だったが、認知の隙間にはまって運悪くすり抜けてしまった系
	- 教訓: 機能単体ではなくワークフロー全体をテストすること、設計レビューでの what-if question の大事さ
- [Scale your Ruby applications with Active Record support for Cloud Spanner \| Google Cloud Blog](https://cloud.google.com/blog/topics/developers-practitioners/scale-your-ruby-applications-active-record-support-cloud-spanner)
	- 同僚に教えてもらった記事
	- ActiveRecord の Google Spanner adapter が GA になった
	- 基本的な操作はできるよう
- [How to stop running out of ephemeral ports and start to love long\-lived connections](https://blog.cloudflare.com/how-to-stop-running-out-of-ephemeral-ports-and-start-to-love-long-lived-connections/)
	- connect の使い方によっては ephemeral port の数が確立できる接続数の上限になってしまうらしい
	- 色々とテクニックを使ってそれを迂回する話。特に UDP はすごい
