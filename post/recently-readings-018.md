{"title":"最近読んだもの 18","date":"2021-10-03T23:30:00+09:00","tags":["readings"]}

## 記事

- [We analyzed 30 SRE job postings to find top SRE responsibilities\.\.](https://spike.sh/blog/sre-role-2021-analysed-30-job-postings/)
	- 30 社の JD を集計して SRE という役割の仕事をまとめたもの
	- インフラの構築運用、SLOSLIの定義と管理、モニタリングとアラートの設定、オンコール対応、自動化
- [10 questions teams should be asking for faster incident response \| PagerDuty](https://www.pagerduty.com/blog/how-to-resolve-incidents-faster/)
	- pagerduty の統計では 2019-2020 比でインシデント数は増えているが MTTA, MTTR は改善していてるらしい。より効率的に障害に対処できるようになってきている
	- 障害対応を Detect, Prevent customer impact, Diagnose, Resolve, Learn のステップに分けてそれぞれの概説
	- まだ読んでないがリンクされている pagerduty のそれぞれのガイドが良さそうだった
- [Reduce Cloud SQL costs with optimizations by Active Assist \| Google Cloud Blog](https://cloud.google.com/blog/products/databases/reduce-cloud-sql-costs-with-optimizations-by-active-assist)
	- Cloud SQL のコスト最適化についてアドバイスしてくれる機能がでたらしい
- [The value of in\-house expertise](https://danluu.com/in-house/)
	- ある程度のスケールの組織では、高度に専門的なチームを内部で持った方が有利になる
		- 例えばカーネルや JVM の専門家チーム
	- 規模が大きいのでその分野の問題が起こる頻度が高い
	- また少しの割合の改善でも絶対値が大きいのでペイしやすい
	- 買ってくるよりも自作した方が、深く使いこなせるので有利なことも多い
- [Lessons learned in a major Rails upgrade: tooling \| by Luan Vieira \| Aug, 2021 \| Invoca Engineering Blog](https://engineering.invoca.com/lessons-learned-in-a-major-rails-upgrade-tooling-67e48838f3b6)
	- Rails 4 から 5 へアップグレードした際に使ったツールなど
	- 複数バージョンを並行できるようにして、細かくリリースしていく戦略が良い

## ドキュメント

- [Memory management best practices  \|  Memorystore for Redis](https://cloud.google.com/memorystore/docs/redis/memory-management-best-practices)
	- redis プロセスが使っているメモリ使用量 `redis.googleapis.com/stats/memory/usage` とインスタンスのメモリ使用量 `redis.googleapis.com/stats/memory/system_memory_usage_ratio` が別概念としてある、というのがポイント
		- 前者は `maxmemory-gb` というパラメータが上限で、それに対してのメモリ使用量。`maxmemory-gb` を超えると eviction の処理が走る
		- 後者はそのインスタンスの実際のメモリ使用量
	- GCP の推奨は `system_memory_usage_ratio` 80% で検知して何らかの対処をすること
	- あまりにもメモリ使用量が高まると `-OOM command not allowed under OOM prevention` というエラメッセージとともに、write がブロックされる
		- これが発生した期間は `System memory overload duration	(redis.googleapis.com/stats/memory/system_memory_overload_duration)` という metric に記録されている
	- またデータのエクスポート (BGSAVE) やスケールアップ、バージョンのアップグレードを行う際、理論上 2 倍のメモリを使いうるので 50% の空き領域が必要になる
- [Moving Forward From Beta \| Kubernetes](https://kubernetes.io/blog/2020/08/21/moving-forward-from-beta/)
	- beta である期間に制限時間を設けているのが面白い
		- rest api だけ
