{"title":"最近読んだもの 34","date":"2022-02-13T22:30:00+09:00","tags":["readings"]}

## 記事

- [What is CAP Theorem?](https://softwaremill.com/what-is-cap-theorem/)
	- CAP 定理について
	- P は通常外せない（P を外すということはシングルノードの、分散システムでないということになる）ので、C と A のどちらを選択するかという問題になる
	- 現実のシステムでは CAP 定理だけでは不足している
	- 平常時に整合性とレイテンシのどちらを取るかとか、可用性とレイテンシといかに高いレベルで達成するかといった課題の方が重要
	- それを指摘する PACELC というコンセプトもある
- [New for App Runner – VPC Support \| AWS News Blog](https://aws.amazon.com/blogs/aws/new-for-app-runner-vpc-support/)
    - AWS App Runner のアプリケーションから VPC 内にある DB などのリソースに接続できるようについになった
    - これがないとほぼ実用は不可能だと思われたので、対応されてよかった
        - ref. [AWS App Runner の軽い検証と雑感 \- Please Sleep](https://please-sleep.cou929.nu/try-aws-app-runner.html)
- [Retrospective and Technical Details on the recent Firefox Outage \- Mozilla Hacks \- the Web developer blog](https://hacks.mozilla.org/2022/02/retrospective-and-technical-details-on-the-recent-firefox-outage/)
	- 先日の Firefox の障害レポート
	- GCP の LB が基盤側要因でアナウンスなしで設定が変わり、それ起因で発覚したレアなコードパスでしか発生しないバグだったらしい
	- なかなか厳しい
- [Google Testing Blog: Code Health: Now You're Thinking With Functions](https://testing.googleblog.com/2022/02/code-health-now-youre-thinking-with.html?m=1)
	- 意図が読みづらい、可読性の低いコードのことを `the maintenance burden` と表現するらしい
	- 内容はどうということはないけれど、英語の勉強になった
- [Kubectl auth changes in GKE v1\.25 : gke\-gcloud\-auth\-plugin \| Google Cloud Blog](https://cloud.google.com/blog/products/containers-kubernetes/kubectl-auth-changes-in-gke)
	- k8s 1.25 ではベンダー固有のコードが oss 版から取り除かれる
	- そのため gke のクライアント側でも認証部分が gke-gcloud-auth-plugin というプラグインに切り分けられるので、各自インストールが必要
- [Minna Bank, COLOPL and 7\-Eleven Japan build apps on Spanner\. \| Google Cloud Blog](https://cloud.google.com/blog/products/databases/minna-bank-colopl-and-7-eleven-japan-build-apps-on-spanner)
	- コロプラ、みんなの銀行、セブン&アイの spanner 活用事例
- [A ‘Hello World’ GitOps Example Walkthrough – zwischenzugs](https://zwischenzugs.com/2021/07/31/a-hello-world-gitops-example-walkthrough/)
	- この図の中で FluxCD の部分がよくわかっていないことが分かった
		- 何のための物かのなんとなくイメージはつくが

## ドキュメント

- [Concurrency and Database Connections in Ruby with ActiveRecord \| Heroku Dev Center](https://devcenter.heroku.com/articles/concurrency-and-database-connections)
	- 前にも読んだが、コネクションプールのパラメタ設定に関する heroku のドキュメント
	- ゾンビになっている接続を早く解放する意味で pool (コネクションプールが保持する接続の最大数) をサーバのスレッド数に揃える意味はあるというのがなるほど
	- 接続数が増えた場合に別デーモン型の接続プール（ここでの例は pgBouncer）を使うとのことだが GET_LOCK などが使えなくなるなど問題もある
- [Deploying Rails Applications with the Puma Web Server \| Heroku Dev Center](https://devcenter.heroku.com/articles/deploying-rails-applications-with-the-puma-web-server)
	- on_worker_boot での establish_connection はもう不要とのこと
		- https://github.com/rails/rails/pull/31241
	- thread-safe に関するトピック　
		- https://stackoverflow.com/questions/15184338/how-to-know-what-is-not-thread-safe-in-ruby/15184752#15184752
	- などおもしろそうなリンクも知れた
	- 英語の話だが Loosely speaking という語彙を知った
