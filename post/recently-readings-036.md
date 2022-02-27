{"title":"最近読んだもの 36 - Understainding Software Dynamics 3 章まで、puma のアーキテクチャなど","date":"2022-02-27T23:30:00+09:00","tags":["readings"]}

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B09H5JB5HC/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/41aDqfiWNbL.jpg" alt="Understanding Software Dynamics (Addison-Wesley Professional Computing Series) (English Edition)" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B09H5JB5HC/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Understanding Software Dynamics (Addison-Wesley Professional Computing Series) (English Edition)</a></div><div class="amazlet-detail">英語版  Richard L Sites  (著)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B09H5JB5HC/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

- [Understanding Software Dynamics](http://www.amazon.co.jp/exec/obidos/ASIN/B09H5JB5HC/pleasesleep-22/ref=nosim/)
    - Chapter 3 まで読んだ
    - 本番環境で一部のリクエストのレイテンシが悪化するケースの原因特定と対処をメインテーマにしている本
        - こうしたケースの原因究明は難しいしツールも出揃っていない
        - 例えば「常に遅い api」はスコープ外。開発環境などで再現して対処できるので
    - まずは cpu とメモリの仕組みと測定方法から話が始まっているが、早くも難しくてついて行けてない
        - 自信がない人は [パタヘネ](http://www.amazon.co.jp/exec/obidos/ASIN/B01M5FMGDL/pleasesleep-22/ref=nosim/) で復習してから読み始めたほうがいいよと冒頭に書いてあり嫌な予感がしたが、そのとおりだった
    - cpu、メモリのハードウェアレベルでの最適化方法の説明を 50 年代のナイーブな実装だった時点から時代を経ながら説明。そのあと C (高級言語) のレイヤから速度を測定する際に、そうした最適化が意図した計測の妨げになるので、どうやって計測処理を実装するのが良いかの説明をしている
        - このあとの章で計測ツールを実装するので、アセンブリではなく　C で頑張っている
    - 個人的には、内容そのものが難しいものの英語で詰まりはせず読めているのはよかった
        - 技術系の英文ならば一定の長さでも集中力が持つようになってきているように感じた。こうして面白そうな本に翻訳を待つことなくアクセスできているのも嬉しいし、「最近読んだもの」を続けてきた甲斐があった
- [puma/architecture\.md at master · puma/puma](https://github.com/puma/puma/blob/master/docs/architecture.md)
    - tcp backlog の他に puma 側でもリクエストの buffer を持っているらしい
    - オプションでオンオフはできる
- [puma/kubernetes\.md at master · puma/puma](https://github.com/puma/puma/blob/master/docs/kubernetes.md)
    - k8s のローリングアップデートの際に、シャットダウン予定の pod がサービスから抜かれる前に、コンテナ内のプロセスに SIGTERM が送られることがあるらしい
    - そうなると SIGTERM を受けて新規リクエストは受け付けない web サーバにリクエストがルーティングされて 5xx が起こってしまう
    - pre-stop hook で sleepするワークアラウンドが必要
- [Kubernetes best practices: terminating with grace \| Google Cloud Blog](https://cloud.google.com/blog/products/containers-kubernetes/kubernetes-best-practices-terminating-with-grace)
    - pod の termination について
    - SIGTERM と preStop hook
- [Container Lifecycle Hooks \| Kubernetes](https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/)
    - hook は ENTRYPOINT と非同期で実行されたり、at least once であるなど、注意点が書かれていた

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01M5FMGDL/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/61vXdBJ3-6L.jpg" alt="コンピュータの構成と設計 第5版 上・下電子合本版" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01M5FMGDL/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">コンピュータの構成と設計 第5版 上・下電子合本版</a></div><div class="amazlet-detail">デイビッド・A・パターソン (著), 成田 光彰 (翻訳)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01M5FMGDL/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
