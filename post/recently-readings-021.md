{"title":"最近読んだもの 21","date":"2021-10-25T00:00:00+09:00","tags":["readings"]}

## 記事

- [A viable solution for Python concurrency \[LWN\.net\]](https://lwn.net/SubscriberLink/872869/0e62bba2db51ec7a/)
	- python のインタプリタから GIL を除く取り組み
	- そもそも GIL が必要とされる理由の大部分が GC のための参照カウントの更新
	- 参照カウントのインクリメント、デクリメント処理をスレッドセーフにするか、同時に一スレッドしか実行されないようにインタプリタで制御することで実現するか
		- カウンターが汚染されるよりはマシなのと、多くのケースではシングルスレッドで十分なので後者のアプローチがとられた
	- 単純にカウンターをスレッドセーフに実装すると、パフォーマンスが劣化する
	- そこで提案手法では、カウンターをオーナー用と他スレッド用の二つに分ける
		- そのオブジェクトのオーナーはロックなどなしにカウンターを更新できる
		- 他スレッドはスレッドセーフな方法で更新する
	- ほとんどの処理はオーナーによって行われるのでパフォーマンス劣化を緩和できる
- [Building and Testing Resilient Ruby on Rails Applications](https://shopifyengineering.myshopify.com/blogs/engineering/building-and-testing-resilient-ruby-on-rails-applications)
	- 肝は最初に最初にでてくる影響範囲のマトリクスかもしれない
		- データストア A が落ちたら、サービス B はダウン、サービス C は性能劣化するがサービス継続、などのあるべき姿をまとめたもの
	- それを実現するためのツールが二つ
		- Toxiproxy は外側から tcp レベルのネットワークの遅延や遮断などを実現し、外部リソースが落ちている状況を再現、自動テストを可能にする
			- semian はサーキットブレイカー
	- どちらも良いツール
- [NGINX 502 Bad Gateway: Gunicorn \| Datadog](https://www.datadoghq.com/blog/nginx-502-bad-gateway-errors-gunicorn/)
	- 先週読んだ [Reverse Proxy, HTTP Keep\-Alive Timeout, and sporadic HTTP 502s \- Ivan Velichko](https://iximiuz.com/en/posts/reverse-proxy-http-keep-alive-and-502s/) の関連記事
	- nginx - gunicorn 構成での 502 発生時のトラブルシューティング
	- nginx のコンフィグの書き方やエラーメッセージ例もあり具体性が高い
- [sync: ExampleWaitGroup includes an porn website url · Issue \#48886 · golang/go](https://github.com/golang/go/issues/48886)
	- Go のコアでテストケースに適当に書いていたドメインがその後実際に使われてポルノサイトになっていた
	- 笑えるけど怖くもある
	- [rfc にもなっているテスト用に確保されたドメイン](https://datatracker.ietf.org/doc/html/rfc2606) を使うようにした方が安心そう
- [The Sidekiq job flow](https://longliveruby.com/articles/sidekiq-job-flow)
	- とても基礎的な内容だった
	- [Job Lifecycle · mperham/sidekiq Wiki](https://github.com/mperham/sidekiq/wiki/Job-Lifecycle) を読めば十分
- [In\-depth explanation of operational metrics at Google Cloud \| Google Cloud Blog](https://cloud.google.com/blog/products/operations/in-depth-explanation-of-operational-metrics-at-google-cloud)
	- Google Cloud Monitoring の入門記事
	- 収集するメトリクスは System metrics, Agent metrics, custom metrics, Log based metrics の四つに分類される
- [SRE top interview questions to land an SRE role](https://www.opsera.io/learn/sre-top-interview-questions)
	- SRE 職の面接での質問例が列挙されていてへーという感じだった
	- カテゴリとしては programming, incident response, support, architecture, networking, problem solving くらい
