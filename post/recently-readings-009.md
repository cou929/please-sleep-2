{"title":"最近読んだもの 9","date":"2021-07-19T20:30:00+09:00","tags":["readings"]}

## 記事

- [The Right Way to Ship Software \| First Round Review](https://review.firstround.com/the-right-way-to-ship-software)
    - 最近考えていたことと近い話題で、かなり面白かった
    - チームにあった開発方針、文化
        - 前提全てに適合するベストはない
        - チームごとの最適解をどう探るか
    - 枠組みとして
        - まず顧客を知る
            - 事業のビジョンやビジネスモデルを起点に
        - 受け入れられるリスクとデプロイ戦略を考える
            - 単純に頻度だけでなく、QA方針や信頼性の定義など広義
            - スタートアップ・大企業、toB・toC、事業ドメインなどに応じて変わる
    - これは本当にそう
        - `But if shipping so much software has taught me one thing, it’s to be an agnostic. Different methodologies optimize for different goals, and all of them have downsides.`
    - web からモバイルファーストに移行する時に、fb のエンジニアも価値観の変更をするのが大変だった。web のリリースコストは 0 に近いのに比べ、モバイルは高いので、コードフリーズ、QA、ドキュメントでの spec 定義などが必要になってくる。move fast and break things はモバイルには適さなかった
        - `Learning new programming languages and frameworks wasn’t what made it hard for Facebook engineers to pivot to mobile. It was hard because they had to undo all their assumptions about how to make software.`
        - `At the heart of that debate were different assumptions about tolerance for risk. Appetite for risk was baked into Facebook’s culture — after all, this company brought you the slogan “Move Fast and Break Things!” Longtime Facebook engineers viewed embracing risk as an essential cultural trait — and at the time, did not realize that mode of operating relied on assumptions about the universe that were true for the web but not for mobile.`
    - どちらが正しいエンジニアリングをしているとかではなく、ビジネスモデルに規定された中で最適解を選んでいるだけ
        - `It’s probably obvious to the world that VMware is substantially more risk averse than Facebook. Realize that it is not because Diane Greene and Mark Zuckerberg have different personalities, nor because one of them is right and one of them is wrong. They are different businesses with different technology stacks and they appropriately ship software in completely different ways based on how much risk their customers can absorb.`
    - デプロイの戦略はプロセスではなくカルチャー
        - `But how you ship is not just process, it’s culture and identity. Swapping out a process is easy. Changing culture is hard. And it’s even harder for a small company to embrace different cultures for different teams.`
    - 組織内に複数のデプロイ文化があるのは大変
        - `If you find multiple “shipping cultures” in tension in your company, you’re dealing with one of the fundamentally hard execution challenges of building and shipping software. There are no easy answers when people stake out positions grounded on emotions rather than reason. On the plus side: your team’s emotions are engaged!`
    - 変数を減らしている
        - リスクとそれに合わせたデプロイ戦略
- [Benchmarks in GO can be surprising](https://leveluppp.ghost.io/benchmarks-in-go-can-be-surprising/)
    - 途中からついていけなくなったけど、アセンブリまでわかってないと解決できない問題が起こり得るということがわかった
- [Write a time-series database engine from scratch](https://nakabonne.dev/posts/write-tsdb-from-scratch/)
    - 力作
    - 差分だけ保存することで空間効率を高めるのはなるほど
- [What Is WebAssembly — and Why Are You Hearing So Much About It? – The New Stack](https://thenewstack.io/what-is-webassembly/)
    - 単にブラウザ上での実行環境というだけでは無かった
    - これがうまくいけば確かにパラダイムも変わりそう
- [Cybersecurity and the Curse of Binary Thinking](https://www.philvenables.com/post/cybersecurity-and-the-curse-of-binary-thinking)
    - わかりやすい結論に流れちゃダメだよみたいな話。例がたくさん
    - 締めの文がよい
        - `if you see something presented as an absolute or a binary choice then use that as a red flag to do some critical thinking. Ask what would need to be also true for this to be true and what would challenge the assertion. It’s fun when you start to do this. `
- [Going Serverless \(on AWS\)\. In order to successfully go Serverless… \| by Payam Moghaddam \| Build Galvanize \| Jul, 2021 \| Medium](https://medium.com/galvanize/going-serverless-on-aws-116a04a0defd)
    - サーバーレスアーキテクチャのみに移行した意思決定はすごい。真似はできない
    - 実際どうなの？という痛いところにちゃんと向き合ってそうなのは良かった
- [Kubernetes Essential Tools: 2021\. Review of the best tools for Kubernetes \| by Javier Ramos \| Jul, 2021 \| ITNEXT](https://itnext.io/kubernetes-essential-tools-2021-def12e84c572)
    - へーと思いながら流し読み
    - 深く知らなくても課題とツールの名前を聞いたことがあるかどうかで結構違う気がする
- [Capturing logs at scale with Fluent Bit and Amazon EKS \| Containers](https://aws.amazon.com/blogs/containers/capturing-logs-at-scale-with-fluent-bit-and-amazon-eks/)
    - 数千 pod 規模のクラスタでは Fluent Bit からの kube-apiserver へのリクエストが増加して負荷が高まる
    - Use_Kubelet オプションで回避した

## コード

- [mhenrixon/sidekiq\-unique\-jobs: Ensure uniqueness of your Sidekiq jobs](https://github.com/mhenrixon/sidekiq-unique-jobs)
    - サイドキックのワーカーをユニークに一度だけ動かすよう制御するミドルウェアらしい
    - どうやっているのかとちょっと見たところ、redis の [EVAL (EVALSHA)](https://redis.io/commands/eval) を使っていた
        - クエリとして lua スクリプトを送ることができる
        - スクリプトはアトミックに実行されるらしい
        - サーバ側にはクエリのキャッシュ機構もあり、EVALSHA というサーバ側にキャッシュがあれば使い、無ければエラー（クライアントはクエリを再送する）という仕組みもある
            - 帯域幅の節約
    - EVAL (EVALSHA) を使って [ロックの確認と取得をアトミックに行っている](https://github.com/mhenrixon/sidekiq-unique-jobs/blob/8c8d54c8b9dea363a7d8b8aeaceb2e82966b8503/lib/sidekiq_unique_jobs/lua/lock.lua) ようだった
        - `locked` という hash で job_id ごとのロックを管理している
        - いろいろなデータ構造が登場して結構複雑になっているが、それぞれ何をしているかまでは追えていない
            - job_id の queue や digest の sorted set などがある
    - これはこれとして、ワーカーを冪等に実行できるよう実装するのが一般的な気はしている
        - 実行済みの場合はスキップするなど、そのジョブが何度呼び出されても壊れないように
