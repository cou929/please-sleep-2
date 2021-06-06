{"title":"最近読んだもの 3","date":"2021-06-06T23:30:00+09:00","tags":["radings"]}

## 記事

[Increment](https://increment.com/containers/) という Stripe が出している雑誌があり、2021 年 5 月号がコンテナ特集だった。

- [A primer on containers – Increment](https://increment.com/containers/primer-on-containerization/)
    - 導入すべきでないケースにも言及していて誠実
- [Building on\-demand staging environments at Paystack – Increment](https://increment.com/containers/on-demand-staging-environments-kubernetes/)
    - 具体的な構成も載せてほしかった
        - データどうしてるかとか、QA どうしてるかとか
    - リモートに開発環境を作る事例をちょいちょい目にするようになってきた
        - マイクロサービスによりローカルに環境を作るのが現実的でなくなってきたのに加え、vs code のリモートモードなどの発達もありそう
- [Interview with Joe Beda, Kubernetes co\-creator – Increment](https://increment.com/containers/joe-beda-interview/)
    - k8s 開発時のモチベーションはアプリケーション（コンテナ、またはプロセス）を物理サーバーから分離させる抽象の提供
    - `we want Kubernetes to be boring. Good infrastructure is boring`
        - かっこいい
    - それとアグレッシブな進化を両立させるための拡張システムという設計
- [Exploring an open\-source codebase: Digging into the Docker CLI – Increment](https://increment.com/containers/exploring-open-source-codebase-docker-cli/)
    - コードベースを理解するための、この人がやってる方法の紹介
    - いきなり上から読み始める前に、issue に目を通して、良さげなのをピックアップ、その変更を行う方法を調べるというアプローチが面白かった

その他の記事。

- [Distributed cloud builds for everyone \- Made of Bugs](https://blog.nelhage.com/post/distributed-builds-for-everyone/)
    - gcc or clang 互換のコンパイルを lambda 上で行う
    - 開発環境をまるっとリモートに置く or push したらビルドされるのどちらとも違い、コンピューティングリソースが必要なところだけをクラウドに任せるのを、いかにシームレスにやるか、という思想かな?
    - llamacc 以外のユースケースにも同じ考え方を適用できていくと便利そう
- [QUIC is now RFC 9000 \| Fastly](https://www.fastly.com/blog/quic-is-now-rfc-9000)
    - Oku さんの名前が普通にでてきててかっこいい
- [The 17 Ways to Run Containers on AWS \- Last Week in AWS](https://www.lastweekinaws.com/blog/the-17-ways-to-run-containers-on-aws/)
    - こんなにあるんだ
- [Using The "X\-Amzn\-Trace\-Id" Header For Request Tracing Through Amazon's Load Balancers](https://www.bennadel.com/blog/4054-using-the-x-amzn-trace-id-header-for-request-tracing-through-amazons-load-balancers.htm)
    - AlB が自動で付与、更新する [X-Amzn-Trace-Id](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-request-tracing.html) というヘッダがあり、リクエストのトレーシングに便利らしい
- [Building an SRE Team? How to Hire, Assess, & Manage SREs](https://www.blameless.com/blog/sre-team)
    - SRE の役割や組織パターンの全体観
    - 引用に役立ちそう
- [That Salesforce outage: Global DNS downfall started by one engineer trying a quick fix • The Register](https://www.theregister.com/2021/05/19/salesforce_root_cause/)
    - `It's always dns`

## 読み終わった本

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y0A8TZO/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/41yo4qcSzlL.jpg" alt="小林カツ代のお料理入門 (文春新書)" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y0A8TZO/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">小林カツ代のお料理入門 (文春新書)</a></div><div class="amazlet-detail">小林カツ代  (著)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y0A8TZO/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

レシピ本ではなくエッセイに近い。生活、という感じでとても良かった。実践性、継続性、あとはユーモアと言うか雰囲気というか、が要素かな。
