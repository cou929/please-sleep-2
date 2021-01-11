{"title":"GCP の教科書 I, II を読んだ","date":"2021-01-11T22:21:00+09:00","tags":["gcp"]}

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B07S1LG1Y1/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/51o4lhZcgXL.jpg" alt="GCPの教科書" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B07S1LG1Y1/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">GCPの教科書</a></div><div class="amazlet-detail">吉積礼敏  (著)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B07S1LG1Y1/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B088LZGPM5/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/51xvCD7nX8L.jpg" alt="GCPの教科書II 【コンテナ開発編】" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B088LZGPM5/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">GCPの教科書II 【コンテナ開発編】</a></div><div class="amazlet-detail">クラウドエース株式会社 (著), 飯島宏太  (著), 高木亮太郎 (著), 妹尾登茂木 (著), 富永裕貴 (著)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B088LZGPM5/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

GCP の超ざっくり概観をさっと頭に入れるために、2 冊とも図書館にあったので借りて一気読みした。意図的にごく簡単に浅く説明されている本なので数時間ほどで目を通せる。正直かなり物足りない部分も多かったし、最近ではこの手の学習には公式ドキュメントを読むのが一番だが、その前にこうした本にざっと目を通すのは悪くないかもしれない。

## 読書メモ

- 情報源
    - [Documentation  \|  Google Cloud](https://cloud.google.com/docs/)
        - [GoogleCloudPlatform/community: This repository holds the content submitted to https://cloud\.google\.com/community\. Files added to the tutorials/ directory will appear at https://cloud\.google\.com/community/tutorials\.](https://github.com/GoogleCloudPlatform/community)
    - [Google Cloud Platform \| Google Cloud Blog](https://cloud.google.com/blog/products/gcp)
        - [Google Cloud Platform \| Google Cloud Blog](https://cloud.google.com/blog/ja/products/gcp)
    - コミュニティ
        - [GCPUG](https://gcpug.jp/)
- ユーザーとプロジェクト・組織
    - ユーザーは個人の google アカウントや G Suite 等
    - プロジェクトごとにリソースが分かれる
    - プロジェクトとユーザーが N:N
    - プロジェクトをまとめる組織
        - 組織とプロジェクトは 0:N
    - 課金ユーザーとは?
- Cloud SDK = CLI コマンド
    - `gcloud <subcommand>`  が基本だが `gsutil (gcs 用)` `bq (bigquery 用)` `kubectl` などは独立したコマンド
    - ローカルなどに任意にもインストールできるし、Web UI 上の Cloud Shell から実行しても良い
    - GCP の各操作は Web UI or SDK で行うことになる
- IAM
    - Identities
        - ユーザー、サービスアカウント
        - グループや G Suite のドメインなどで複数にも
    - Role
        - 権限のセット
        - プロジェクトへの参照・編集・オーナー・閲覧者という基本セットと、リソースごとのより詳細な事前定義済みのロールの大きく 2 種類に分かれる
    - Policy
        - Id + Role のセット
        - リソースのヒエラルキーのどこかの層へのアクセス制御になっている
            - 組織 > プロジェクト > リソース
    - GCS の権限管理
        - バケットに対しては IAM で制御するが、オブジェクト単位での制御をする ACL という仕組みもある
- [Live Migration](https://cloud.google.com/compute/docs/instances/live-migration)
    - メモリを含めあるインスタンスの状態をまるっとコピーしておき、任意のタイミングで切り替えることで、ダウンタイムほぼ無しでマイグレーションできるという機能らしい
    - インフラの移行などをしない限りは使わなそうだが、表面を聞くだけでも過ごそうな技術
- GAE の制約
    - 処理の実行時間上限、リクエストサイズ上限、(恐らくローカル? エフェメラル?への) ファイル書き込み不可、vcpu, mem が少なめのインスタンスしか無い、といった制約がある
    - そもそも PaaS なので提供されている環境以外は動かせない
        - コンテナを簡単に動かす環境としての Cloud Run はある
- PaaS と Caas (Contaner as a Service) の違いの整理
    - 上から下へこういうレイヤーに分けたとする
        - `アプリケーション` `ラインタイム` `コンテナランタイム` `OS` `仮想化` `ハードウェア`
    - CaaS は `ランタイム` より上をユーザーが、`コンテナランタイム` より下をプラットフォーマーが提供する
    - PaaS は `アプリケーション` より上をユーザーが、`ランタイム` より下をプラットフォーマーが提供する
        - この構造上ランタイムの選択肢が CaaS よりも少なくなる
    - heroku や Beanstalk にコンテナをデプロイできるので、実際この 2 つの境界はあいまいかも
    - ちなみに IaaS は `OS` より上をユーザーが、`仮想化` より下をプラットフォーマーが提供するという整理にできる
- [Knative](https://knative.dev/)
    - k8s 上で paas (caas) を実現するための仕組みらしい
    - Google がはじめたプロジェクトだが現在はコミュニティベースで [Pivotal, IBM, RedHat, Cisco, VMWare](https://github.com/knative/serving/blob/master/AUTHORS) なども参加している
    - Cloud Run はこれベースらしい

### わからなかったこと

- ネットワーク
    - VPC や Security Groups 的なもの (そもそもあるのか、抽象化されているのかなど)
    - それに応じたネットワーク設計の考え方
- データストア
    - BigTable, DataStore, FireStore あたりの棲み分けの理解
    - Dataflow, Dataproc, Dataprep あたりの理解
