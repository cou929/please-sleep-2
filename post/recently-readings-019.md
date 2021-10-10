{"title":"最近読んだもの 19","date":"2021-10-10T23:30:00+09:00","tags":["readings"]}

## 記事

- [Reverse Proxy, HTTP Keep\-Alive Timeout, and sporadic HTTP 502s \- Ivan Velichko](https://iximiuz.com/en/posts/reverse-proxy-http-keep-alive-and-502s/)
	- TCP の状態は分散システムという指摘が目から鱗だった
	- CAP 定理を当てはめると、AP を保障するために C は結果整合性になっている
	- 具体例として、reverse proxy とその後ろにいる upstream server 間で接続をプールし keep alive しているケースでの 502
	- idle connection のタイムアウトが proxy > upstream server だった場合、proxy が  request を送っている最中に upstream が fin を返すケースが起こりうる
	- この時クライアントには 502 が reverse proxy から返され、upstream server はエラーに全く気づいていない
- [Update about the October 4th outage \- Facebook Engineering](https://engineering.fb.com/2021/10/04/networking-traffic/outage/)
	- 一定以上高度化したソフトウェアシステムの障害は、もうほぼ設定ミス起因ばかりな気がする。体感だけど
	- 運用的なブレイクスルーは、カオスエンジニアリングの次はこの辺の発明が来るのかな。他の分野、軍事航空宇宙とかの知見が取り入れられたりとかもありそう
- [Lessons learned in a major Rails upgrade: strategy \| by Luan Vieira \| Aug, 2021 \| Invoca Engineering Blog](https://engineering.invoca.com/lessons-learned-in-a-major-rails-upgrade-strategy-1b0463ec41fc)
	- 先週読んだ rails アップグレードの話の関連記事
	- 細かくマージ、デプロイできる仕組みが良い
- [Spot GKE cost optimizations in the console \| Google Cloud Blog](https://cloud.google.com/blog/products/containers-kubernetes/spot-gke-cost-optimizations-in-the-console)
	- Cloud SQL に続き GKE のコスト最適化サジェスト機能らしい
- [How \(some\) good corporate engineering blogs are written](https://danluu.com/corp-eng-blogs/)
	- 面白い企業テックブログの共通点はレビュープロセスが短いことかもという話
	- 確かにふわっとした具体性のない記事は結構あるが、それはリスクヘッジのためレビュープロセスが長くなっているゆえなのかもしれない
	- 例に挙げられている面白いブログのなかでは、cloudfrare は読んだことがある。とても面白い
	- 今でもすぐに思い出せるのは https://blog.cloudflare.com/syn-packet-handling-in-the-wild/ 。いい記事
- [Inside the Lab: Expanding connectivity by sea, land, and air](https://tech.fb.com/inside-the-lab-connectivity/)
	- 大陸間のケーブルのため太陽光で電力供給するブイの発明、光ファイバーを敷設するロボット、電線から各家庭への無線通信
	- スケールがでかくて面白い

## ドキュメント

- [Using pipelining to speedup Redis queries – Redis](https://redis.io/topics/pipelining)
	- 読み取りリクエストをパイプラインでまとめた際の、レスポンスのエラーハンドリングがどうなるか知りたかったがそう言う話は出てこなかった。クライアントの実装依存なのかな
	- ループバック経由で redis server に接続する場合でも、ベンチマークで五倍ほど差が出ていたのは意外だった。こんなに効くのか
