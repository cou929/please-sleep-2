{"title":"AWS App Runner の軽い検証と雑感","date":"2021-07-17T14:30:00+09:00","tags":["aws"]}

[AWS App Runner – Fully managed container application service \- Amazon Web Services](https://aws.amazon.com/apprunner/?nc1=h_ls)

機会があって、コンテナ化されている小さいアプリケーションをさっと動かしてデモするための環境を相談された。

そういえば最近 [App Runner が出たよというニュース](https://aws.amazon.com/blogs/aws/app-runner-from-code-to-scalable-secure-web-apps/) を見たのを思い出し、今回の用途に使えるかもしれないので、素振りがてら簡単に検証してみたメモ。

結構ブロッカーはあり、本番で使うのはまだ厳しそうな感触だった。ただ自分の考え方をシフトできていないだけかもしれない。

## 前提

- コンテナ化された小さいアプリケーションがすでにある
  - データストアとして MySQL が必要
- GitHub の特定のブランチに push すると GitHub Actions 経由で ECR に push するようにした
- App Runner は ECR 上のリポジトリ特定のタグを監視していて更新があるとデプロイが走るようにした

## 気になったこと

### App Runner から private な VPC 内にある RDS などのリソースに接続できない

例えば RDS の場合、今だと `Public Access: Yes` にしないとダメそう。これはけっこう致命的な気がするが...

- App Runner を使いたいユーザーはある程度 AWS 上に既存の資産があると思われる
    - アプリケーション (ワークロード) をサーバレスに移したくて App Runner に乗せてみたが、既存の RDS に接続できなくてがっかり... ということは起こり得そう
- そうでなければ (新規アプリで他に AWS 上で動いているリソースも無いのであれば) Heroku なりを使えばもっと楽なわけなので
　
要望の 1 番にあがってきているし upvote もたくさんついているので、まあみんなそう思ってるんだと思う。

[Allow App Runner services to talk to AWS resources in a private Amazon VPC · Issue \#1 · aws/apprunner\-roadmap](https://github.com/aws/apprunner-roadmap/issues/1)

今回は検証なのでとりあえず RDS に public access して凌いだ。

それとも、もしかしたら自分の考え方が古いのかもしれない。ゼロトラスト的に必要な認証さえできればどこからでも RDS にアクセスできたほうが良いのだろうか。

- やるとすると少なくとも通信は TLS にしたい
    - 高負荷な環境ではパフォーマンス的に不利になりそう
    - [SSL/TLS を使用した DB インスタンスへの接続の暗号化 \- Amazon Relational Database Service](https://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL.html)
- IAM でも認証することができるらしく、これはありかも
    - [IAM認証によるRDS接続を試してみた \| DevelopersIO](https://dev.classmethod.jp/articles/iam-auth-rds/)
- Secret Manager にパスワード管理を任せると、連動して RDS 側のパスワードも定期更新してくれるらしい
    - [Tutorial: Rotating a secret for an AWS database \- AWS Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/tutorials_db-rotate.html)
    - 一見良さそうだけど、アプリケーション側でクレデンシャルの更新にダウンタイムなしにどう追従するかがよくわからなかった
    - [aws が出しているサンプル](https://github.com/aws-samples/aws-secrets-manager-credential-rotation-without-container-restart) ではアプリケーションが接続確立時に[直接 secret manager に問い合わせている](https://github.com/aws-samples/aws-secrets-manager-credential-rotation-without-container-restart/blob/11ad22e8f1d55bf48af219fecdd4ba208c88dff4/webapp/app/codecompose/db/backends/secretsmanager/mysql/base.py) ようだった
      - クレデンシャルを環境変数でアプリケーションに渡すだけに比べて面倒

### ECR のタグを immutable にできない

まず、そんなに自信なく間違っているかもしれないが、コンテナイメージのリポジトリのタグはイミュータブル（上書き不可）にした方が運用時の事故が減らせる気がしている。

実際 ECR には [Tag Immutability という設定項目](https://aws.amazon.com/about-aws/whats-new/2019/07/amazon-ecr-now-supports-immutable-image-tags/)があり、上書き禁止にすることもできる。

一方で App Runner は一つの固定値でしか参照するリポジトリ、タグを指定できない。また指定したタグに更新があると自動でデプロイする機能はある。つまり Tag は mutable なものとして運用する思想になっているように見える。

まあ確かに latest タグだけは mutable にする、という運用は全然ありえる気はする。（ECR で特定のタグ以外は immutable という設定はできるのだろうか）

あるいは App Runner のサービスはデプロイごとに新しく作る思想なのかもしれない。自分はサービスを一つ作って更新していく方向で構築してしまった。

ちなみに今回は ECR リポジトリは tag mutable  とし、latest タグが常に最新、それとは別に git のコミットハッシュを同じイメージに別タグとして打って履歴も残すことにした。こんな感じで GitHub Actions のワークフローを設定した。

```yaml
    steps:
    - uses: actions/checkout@v1

    - name: Configure AWS Credentials
      uses: aws-actions/configure-aws-credentials@v1
      with:
        aws-access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID }}
        aws-secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
        aws-region: ap-northeast-1

    - name: Login to Amazon ECR
      id: login-ecr
      uses: aws-actions/amazon-ecr-login@v1

    - name: Build, tag, and push image to Amazon ECR
      env:
        ECR_REGISTRY: ${{ steps.login-ecr.outputs.registry }}
        ECR_REPOSITORY: ${{ secrets.ECR_REPOSITORY }}
      run: |
        IMAGE_TAG=$(git rev-parse --short HEAD)-$(date +%Y%m%d%H%M%S)  # コミットハッシュ込みのタグ。運用上は immutable
        LATEST_TAG=latest                                              # mutable な latest タグ
        docker build -t $ECR_REPOSITORY:$IMAGE_TAG -f path/to/Dockerfile .
        docker tag $ECR_REPOSITORY:$IMAGE_TAG $ECR_REGISTRY/$ECR_REPOSITORY:$LATEST_TAG
        docker push $ECR_REGISTRY/$ECR_REPOSITORY:$LATEST_TAG
        docker tag $ECR_REPOSITORY:$IMAGE_TAG $ECR_REGISTRY/$ECR_REPOSITORY:$IMAGE_TAG
        docker push $ECR_REGISTRY/$ECR_REPOSITORY:$IMAGE_TAG
```

### HTTP ヘルスチェックをコンソールから設定できない

TCP (指定したポートで接続を確立できるか）はコンソールから設定できるが、HTTP GET でのヘルスチェック設定がなぜかコンソールからできない。

機能としては存在して、aws cli からは設定できる。

[create\-service — AWS CLI 1\.20\.1 Command Reference](https://docs.aws.amazon.com/cli/latest/reference/apprunner/create-service.html)

既存のアプリを載せる場合、HTTP のヘルスチェック用エンドポイントを持っていることも多いと思うので、コンソールからも設定できた方が便利な気はするが...

[Support configuration of HTTP health checks from the console · Issue \#63 · aws/apprunner\-roadmap](https://github.com/aws/apprunner-roadmap/issues/63)

## PR

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B08SGSD479/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/519DqB8xIvL.jpg" alt="改訂新版 徹底攻略 AWS認定 ソリューションアーキテクト − アソシエイト教科書［SAA-C02］対応 徹底攻略シリーズ" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B08SGSD479/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">改訂新版 徹底攻略 AWS認定 ソリューションアーキテクト − アソシエイト教科書［SAA-C02］対応 徹底攻略シリーズ</a></div><div class="amazlet-detail">鳥谷部 昭寛  (著), 宮口 光平 (著), 菖蒲 淳司 (著)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B08SGSD479/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
