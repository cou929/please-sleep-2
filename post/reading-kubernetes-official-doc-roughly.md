{"title":"Kubernetes の公式ドキュメントを流し読みした","date":"2021-01-02T23:38:58+09:00","tags":["nix"]}

[Concepts \| Kubernetes](https://kubernetes.io/docs/concepts/)

[イラストでわかる Docker と Kubernetes を読んだ \- Please Sleep](https://please-sleep.cou929.nu/docker-kubenetes-book.html) の続きで、kubernetes の公式ドキュメントの主に `Concepts` の部分を流し読みした。ざっくり頭の中に目次を作りたいだけなので、ほぼ読み飛ばしている。

感想としては、まずもってコンセプトや設計が先進的なのは流石としか言いようがない。k8s の単なるユーザーとして、コンセプトに沿ったアプリケーションを作るだけでも、自然とクラウド環境でのベストプラクティスに近いものになりそうな気がする。(ここでのベストプラクティスとは [The Twelve\-Factor App](https://12factor.net/) のようなイメージ)。分散システム上のデータ整合性 (etcd など)、スケジューラのアルゴリズム、kube-proxy の実装など、技術的にチャレンジングな部分がなんとなく察せられた。機会があれば調べてみたい。

一方でどうかなと思ったこと。コンテナベースの実行環境を非常に上手く抽象化しているのだが、反面やりたいことをシンプルに実現するには、k8s が抽象化している部分が大きすぎる恐れがあるなと思った。AWS など特定のクラウドプロバイダにガッツリ依存して必要最低限のサービスに絞って構築したほうがシンプルにできそうなのと、それで十分なケースは多そうに思われる。多くのクラウドプロバイダやオンプレのケースなど、k8s は幅広い環境に対応するようにうまく抽象化しているが、それゆえの遠回りは発生しそう。すべてを AWS サービスに統一するなど特定環境に依存した作りにしたほうが無駄が少ないかもしれない (AWS 独自サービス間のほうが連携性が良いと思われるので)。なお、自分は実際に本番で運用したことは無く、これはドキュメントを読んだだけの現時点での印象でしかないので、初心者の適当な物言いとしてスルーしてほしいです。

以下メモ。

## [Concepts \| Kubernetes](https://kubernetes.io/docs/concepts/)

### [Overview \| Kubernetes](https://kubernetes.io/docs/concepts/overview/)

- 基本
    - desired state を定義する
    - kubectl が kubernetes api を使って設定する
    - controll plane が desired state に保つ
    - controll plane
        - master node
            - kube-apiserver、kube-controller-manager、kube-scheduler
            - master node はクラスタ内に 1 ノードある
        - 非 master node
            - kubelet、kube-proxy
    - ![](https://d33wubrfki0l68.cloudfront.net/7016517375d10c702489167e704dcb99e570df85/7bb53/images/docs/components-of-kubernetes.png)
    - 基本的なオブジェクト
        - Pod, Service, Volume, Namespace
    - オブジェクトより高レベルの抽象化
        - Deployment, DaemonSet, StatefulSet, ReplicaSet, Job
- Components
	- Control plane
		- Kube-apiserver
			- 水平スケール可能
		- etcd
			-クラスタ情報を保管する kvs
			- 要バックアップ
		- kube-scheduler
			- 空いている pod を保持して job を割り振り
		- kube-controller-manager
			- 各種リソース管理
		- cloud-controller-manager
			- クラウドプロバイダーとのやりとり
	- Node
		- Kubelet
			- ノード内で pod とコンテナを管理
			- api とやりとり
		- kube-proxy
			- ネットワークのルールを実現
		- container runtime
			- containerd, cri-o など
	- addon
		- Dns, web ui, etc
- Kubernetes objects
	- ユーザーは desired state を表明
	- api を通じて操作
		- kubectl or calling api directly
	- Spec and state
		- ユーザーは spec を指定し、control plane が state を更新
	- kubectl が yaml をもとに api を呼び出す
		- apiVersion, kind, metadata, spec というフィールド構成
- object management
	- 直接コマンド発行、設定ファイル適用、設定ファイル群の directory 適用の3種類
- 名前
	- ユーザーがつける name とシステムがつけるuuid はクラスタ内のオブジェクトごとに一意
	- label, attribute はそうではない
- namespace
	- クラスタを複数ユーザー（チームなど）で共有する際の分離手段
		- デプロイタグなど同じユーザーが使うなら label で分離した方が良い
		- コマンドごとにオプションで namespace を指定したり、kubectl の config に書いてそれを省略したり
	- default, kube-node-lease, kube-public, kube-system は initial からある
	- namespace, node, volume といったリソースは namespace に属しない
		- `kubectl api-resources --namespaced=false` とかで調べられる
- label
	- リソースを任意の方法で分類
	- あるアーキテクチャでだけ pod が起動するようになどといった使い方
		- オブジェクトのグルーピング
- attribute
	- label がグルーピングに対して、それ以外の用途のメタデータ、外部ツール用など
- field selector
	- metadata.name や status といったオブジェクトのフィールドでフィルタできる
- pod, service etc は最低限のセットで、それ以上のユースケースは label や attribute で対応する設計ぽい
	- label と attribute を分けているのも興味深い（tag のようなものでなく）
	- label の方が公式 api から公式に使われている
    
### [Cluster Architecture \| Kubernetes](https://kubernetes.io/docs/concepts/architecture/)

- Node
    - Management
        - kubelet の自己登録 or manual 登録
        - 登録され、チェックが通ったら追加される
    - Status
        - Addresses
        - Conditions
        - Capacity and Allocatable
            - cpu, mem, 作成できる pod 上限など
        - Info
            - 各種バージョンなど
    - Node Controller
        - Control Plane の 1 モジュール
        - node 登録時の CIDR 割当
        - Node Controller が持つノードリストの管理 (クラウドプロバイダーとの整合)
        - ノードのヘルスチェック
            - `NodeStatus` の更新 (デフォルトでは 5 分ごと)
            - Lease object の更新 (デフォルトでは 10 秒ごと)
        - eviction
            - eviction-rate は秒間に evict するノード数 (デフォルト 0.1 = 10 秒に 1 ノード以上は外れない)
            - az 内の unhealthy なノードの割合が一定以上だと eviction-rate をさげたり、ノード数が少なすぎると evict しないといった挙動をする
    - graceful shutdown
        - アルゴリズムは [inhibit](https://www.freedesktop.org/wiki/Software/systemd/inhibit/) らしい
- Control Plane-Node Communication
    - [Hub\-and\-Spoke Architecture \| Polarising](https://www.polarising.com/2016/09/hub-spoke-architecture/)
- Controllers
    - control loop の disred state の仕組みをサーモスタットに例えている
    - Control plane の各コントローラーは自分の責務に集中
    - 通信は apiserver を通じて行う
- Cloud control manager
    - クラウドプロバイダーとのインタフェース、プラグイン形式
    - 本体とクラウドプロバイダーとを疎結合にする設計

## [Workloads \| Kubernetes](https://kubernetes.io/docs/concepts/workloads/)

- Pod の集合の表現が workloads
- Deployment, ReolicaSet
	- Stateless
- StatefulSet
- DaemonSet
- Job, CronJob
- サードパーティのカスタムもある

## [Services, Load Balancing, and Networking \| Kubernetes](https://kubernetes.io/docs/concepts/services-networking/)

- 扱うのは以下
	- コンテナからの通信
	- pod 間の通信
	- service による外部公開
	- service によるクラスタ内公開
- 個人的には iptables, nat, dns あたりの理解が怪しいのがわかった

### [Service \| Kubernetes](https://kubernetes.io/docs/concepts/services-networking/service/)

- workload で定義された pod の集合をバックエンドとして、それを呼び出すフロントエンドとは疎結合にしたい
	- フロントはバックエンドの pod の一覧と状態管理やアクセス先の管理をしたくない
	- バックエンドは任意のタイミングで pod を入れ替えたい
- service で実現する
	- 対象の pod をセレクタなどで指定
	- そこにアクセスするための ip と名前の割り当て
		- clusterIP (= virtual ip)
- kube-proxy
	- ノードからの通信の調整
		- apiserver に service の情報を聞く
		- いくつかの方式
			- user space proxy, iptables proxy etc
	- dns ラウンドロビンは使わない理由
		- 歴史的経緯で理想的でない実装がありうる
		- ttl0 だと負荷
- Service discovery
	- 環境変数
		- kubelet がサービスごとの接続先が入っている環境変数を設定する
	- dns
		- アドオン
		- apiserver と通信し更新
		- https://github.com/coredns/coredns
- service type
	- クラスタ内のみ、外部公開、クラウドプロバイダの lb 使用など
	- 外部公開の仕組みにはほかに ingress がある
	- クラウドプロバイダごとにアノテーションで細かい設定がある
	- 外部用の名前をつける externalName
- virtual ip の実装
    - etcd で ip 振り出しの重複排除
    - iptable の利用
        - client 側で接続先を振り分ける (kube-proxy が apiserver とやりとりしつつ iptables を活用)
    - クラスタの大きさによってスケールするよう複数のアルゴリズムを選択可能

### [Connecting Applications with Services \| Kubernetes](https://kubernetes.io/docs/concepts/services-networking/connect-applications-service/)

- 普通にあるホストで Docker コンテナを動作させる場合
    - 他のコンテナとのやりとりで、ポート番号やホストの ip の衝突に注意して管理する必要がある
- k8s の場合 pod 単位で ip (クラスタプライベートの) を割り振る
    - 「同一のノード内で複数の nginx pod を立ち上げ、それぞれのコンテナは 80 番を待ち受ける」というようなユースケースを、クラスタのユーザーの負担低く実現できる
        - ノードの ip とその中の pod (コンテナ) 単位のポートに対して通信する方式 (普通はこっちのほうが一般的) よりも、ユーザー視点ではかなりシンプルになる
    - 大規模なクラスタ、チームでの運用負荷を下げられる
- endpoints
    - service に所属する pod のセット
    - controller が定期的に更新

### [Ingress \| Kubernetes](https://kubernetes.io/docs/concepts/services-networking/ingress/)

- クラスタ全体の外部との GW
- 外部からの接続をどう内部の Service に割り振るかルールを記載
- ロードバランス、SSL 終端、名前ベースのバーチャルホストなども
- 実態は各クラウドプロバイダの LB や nginx などのミドルウェアで、ユーザーが設定する必要がある
    - ingress controller

### [Network Policies \| Kubernetes](https://kubernetes.io/docs/concepts/services-networking/network-policies/)

- pod, namespace 単位でのアクセス制御や ip ブロックなど
    - ip, port 番号単位での管理 (OSI layer 3, 4)

## [Storage \| Kubernetes](https://kubernetes.io/docs/concepts/storage/)

- 通常は ephemeral だが、永続化用途のために
- EBS などを指定した位置にマウントできる
- k8s では PersistentVolume (ストレージのセット) と PersistentVolumeClaim (ユーザーが必要とするストレージ要件の定義。コンピューティングリソースに対する pod のようなもの) というリソースで表現している

## [Configuration \| Kubernetes](https://kubernetes.io/docs/concepts/configuration/)

- configMap や secrets に設定などを切り離す
- resource (cpu, memなど) の request, limit の設定
- kubectl 用の設定ファイル kubeconfig

## [Security \| Kubernetes](https://kubernetes.io/docs/concepts/security/)

- The 4C's of Cloud Native Security
	- Cloud, cluster, container, code layer
	- 下の層の上にセキュリティを構築していく
		- 下の層がセキュアじゃなかったら上で頑張っても限定的
- Cloud layer
	- クラウドプロバイダの責任範囲
	- ネットワークの設定
- Cluster layer
	- クラスタのリソース操作をセキュアに
		- 操作への認証や認可
	- アプリケーションそのもの
		- 認証、認可、秘匿情報管理など
- Container layer
	- 脆弱性のスキャン、イメージのサイン、特権排除など
- Code layer
	- tls の利用、依存ライブラリの信頼性など

## [Policies \| Kubernetes](https://kubernetes.io/docs/concepts/policy/)

- namespace  ごとにクオータ設定可
	- ResourceQuota オブジェクト

## [Scheduling and Eviction \| Kubernetes](https://kubernetes.io/docs/concepts/scheduling-eviction/)

- Pod をどのノードに割り当てるかを決めるロジック
- 候補となるノードをフィルタし、残りにスコアをつけ最も高いものが選ばれる
- taints and tolerations で割り当ての調整
- Node affinity で、できればこう割り当てて欲しい、欲しくないの表現
- eviction policy
    - kubelet がリソース不足のノードを検知すると pod を落とす
    - deployment などは desired state に合わせて対応

## [Cluster Administration \| Kubernetes](https://kubernetes.io/docs/concepts/cluster-administration/)

- todo
- マネージドサービスではどの程度カバーされるがわからないが、実際に運用する際には読んだほうが良さそう

## その他

- [kubernetes/examples: Kubernetes application example tutorials](https://github.com/kubernetes/examples)
    - example
- [Learn Kubernetes Basics \| Kubernetes](https://kubernetes.io/docs/tutorials/kubernetes-basics/)
    - [katacoda](https://www.katacoda.com/) がすごい
        - チュートリアルとともにブラウザ上で実際に試せる環境
        - 複数ターミナル開ける
        - 日本からだと遅いかも
    - デフォルトでは pods は kube クラスタ内からしかアクセスできないが、kubectl proxy でプロキシ経由でアクセスできる
        - `localhost:8001`

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B08FZX8PYW/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/51C+pft8SJL.jpg" alt="Kubernetes完全ガイド 第2版 impress top gearシリーズ" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B08FZX8PYW/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Kubernetes完全ガイド 第2版 impress top gearシリーズ</a></div><div class="amazlet-detail">青山真也  (著)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B08FZX8PYW/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
