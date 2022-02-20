{"title":"最近読んだもの 35 - Netflix のインベント通知、スタートアップでの障害対応、NewRelic の SLO/SLI ベストプラクティスなど","date":"2022-02-20T22:00:00+09:00","tags":["readings"]}

- [Rapid Event Notification System at Netflix \| by Netflix Technology Blog \| Feb, 2022 \| Netflix TechBlog](https://netflixtechblog.com/rapid-event-notification-system-at-netflix-6deb1d2b57d1)
	- Netflixのイベント通知システム
	- このレベルになると db に単純に crud とは当然なっていなくて、ユーザーアクションや clockをトリガーにイベントが生成されら最終的にクライアントに通知される仕組みになっている
	- イベントの priorityが設定できたり通知も push, pull両方に対応しているなど、さすがの規模という感じ
- [The startup guide to sensible incident management \| incident\.io \| incident\.io](https://incident.io/blog/the-startup-guide-to-sensible-incident-management)
	- SRE 本をはじめベストプラクティスは数多くあるが、スタートアップにそのまま適用するとオーバーキルなので、どの程度の塩梅がよいかのおすすめの紹介
	- severity は major, minor の2種類で十分とか、フォーマルなポストモーテムは ROI が合わないとか
- [SLOs and SLIs Best Practices for Modern, Complex Systems \| New Relic](https://newrelic.com/blog/best-practices/best-practices-for-setting-slos-and-slis-for-modern-complex-systems)
	- NewRelic による SLO/SLI 設定の tips
	- system boundary ごとに設定すると良い
		- システムの内部的な単位であるコンポーネントごとではなく、ユーザーから見た境界で区切るとシンプルで良いらしい
	- まず自然言語でそのシステムが満たすべき目標を書いてから、それを SLI に変換して、組み合わせて SLO にする
- [The Cloud Spanner team busts some of the most common myths \| Google Cloud Blog](https://cloud.google.com/blog/products/databases/cloud-spanner-myths-busted)
	- spanner のよくある誤解を解くという記事
	- 上位の項目が「高負荷のワークロードで 使う物」「費用が高い」「一貫性を高めるとレイテンシが遅くなる」というもので、とはいえ解答を読んでもそんなに外れてない気がしてしまった
	- レイテンシに関しては stale read というオプションがあることが紹介されていた
- [Cloud SQL launches IAM Conditions and Tags \| Google Cloud Blog](https://cloud.google.com/blog/products/databases/cloud-sql-launches-iam-conditions-and-tags)
	- cloudsql でdb instance にタグをつけて、iam から特定のタグにだけ許可と言った、条件付きの権限設定ができるようになったらしい
