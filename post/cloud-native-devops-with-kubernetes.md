{"title":"Kubernetesで実践するクラウドネイティブDevOps を読んだ","date":"2023-01-29T22:00:00+09:00","tags":["k8s", "book"]}

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873119014/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/51wecBhtIOL._SX389_BO1,204,203,200_.jpg" alt="Kubernetesで実践するクラウドネイティブDevOps" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873119014/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Kubernetesで実践するクラウドネイティブDevOps</a></div><div class="amazlet-detail">John Arundel  (著), Justin Domingus (著), 須田 一輝 (監修), 渡邉 了介 (翻訳)<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873119014/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

<a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873119014/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Kubernetesで実践するクラウドネイティブDevOps</a> を読んだ。本番環境で Kubernetes クラスタを運用するために必要な知識が体系的かつコンパクトにまとまっていて、自分の知識のインデックスをさっと作るのにとても良い本だった。

少し個人的な経緯を記載しておく。全く知識のない状態で Kubernetes (GKE) を使うプロジェクトに入った際に、自分はまず [CKAD を取得した](https://please-sleep.cou929.nu/ckad.html)。[Udemy のコース](https://www.udemy.com/course/certified-kubernetes-application-developer/) を一通り見て概念を掴み、実際に [資格](https://www.cncf.io/certification/ckad/) の受験・取得まで行った。CKAD が良かったのは、具体的な目標とカリキュラムがあるので必要な情報をスピーディに学べることと、ハンズオン形式なので手を動かしながら理解を深められるのが非常に良かった。特に kubectl でのオペレーションなどはある程度手になじませたほうが良いので、実運用でもこの資格取得で得た知識が役に立った実感が多くあった。ちなみに CKA も同時に準備していたが、出産等で受験する時間がとれないまま今に至っている。ただマネージドの Kubernetes を運用するために必要な最低限の知識をまずざっと得るのには CKAD が良かったと今でも思う。

その後実際に運用しながら、知らないトピック (CKAD ではカバーされていなかったリソースや GKE 側の知識など。例えば PDB、HPA、Pod Anti-Affinity などは自分が受けたときの CKAD ではカバーされていなかったように記憶している) については都度公式ドキュメントを読みながら対応していた。このように必要な部分に対してドキュメントや Issue 読んでいくので良いのだが、だんだんと自分のプロジェクトで利用する部分だけに知識が偏っている感覚が増えてきた。また「X をしたい」という目標に対して、どのような方法をとればよいのかパッと出てこない状態になっていた。

こうしてそろそろ体系的に知識を整理し直したいと思っていたところに、この <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873119014/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Kubernetesで実践するクラウドネイティブDevOps</a> はぴったりの本だった。

まず "本番運用に必要な知識" にフォーカスして、必要以上に細部に立ち入らずに記載されているので、スピーディに読むことができた。運用のうえで必要になる概念だけに集中していることや、数多くあるツールの中で代表的なものを選択肢として提示してくれている。詳細は最新のドキュメントで確認したほうが良く、必要なのは脳内のインデックスだったので、本書のこの方針はまさに今の自分が求めているものだった。

また知っていた知識についても、その重要度がこの本を通して明確になった感覚がある。例えばコンテナリソースの requests, limits はもちろん知識としては知っていたが、どちらかというと単にスペックの中の一要素程度にしか認識していなかった。これが本書では "リソースの管理" の章のトップで扱われていて、確かに自分の認識よりも重要なパラメータだということが理解でき、知識の補正になったと思う。

オブザーバビリティについてもカバーしている。個人的になかなかとらえどころのない概念でよく理解できていなかったのだが、「分散システムの性能はグラデーション的で、普段から常にどこかデグレしている」という前提にたっているという説明がわかりやすかった。とはいえかなり短い紙面しかないので、[Observability Engineering](https://www.oreilly.com/library/view/observability-engineering/9781492076438/) あたりは読んでみたい気になった (以前無料公開されていた PDF が積んである)。

問題があるとすると原著が 2019 年刊行なので、Kubernetes の変化のスピードに対しては少し古くなりつつあると思う。基本的な概念は変わっていないが、紹介されているツールや SaaS 等が一部古くなっているように見えた。間は自分で埋める必要がある。

以下読書メモ。

## 3. Kubernetes 環境の選択

- k8s 環境を準備するにあたって、マネージドサービス、ターンキー方式（コントロールプレーンはマネージドでワーカーノードは自前）、セルフホスティング（その際に有用な kops といったツールの紹介）という方式に大別して、それぞれの特徴や代表的なサービスを検討
- その上で可能な限りはマネージドサービスを利用することを強く推奨している
    - [k8s をスクラッチから本番運用できる状態にするにはエンジニアの給与で 100 万ドルかかるという経験則](https://mobile.twitter.com/copyconstruct/status/1020880388464377856)がある
    - 例えばセットアップは kubeadm といった便利なツールがあるが、運用上のハードで難しい部分を解決してくれるツールは無い

## 5. リソースの管理

- 最も基礎となる情報として、コンテナのリソースの要求 `spec.containers[].resources.requests` と制限 `spec.containers[].resources.limits` がある
    - 原則どちらも設定されているのが望ましい
- また健全性、準備完了の判定として Liveness, Readiness probe がある
    - Liveness probeに失敗した場合 pod は再起動を試みる。Readiness probeに失敗している場合は Service配下に入らない
- PodDisruptionBudget でpod の eviction が多すぎないよう、一定の pod を確保するように指定できる
- プリエンプティブノードを使う場合は Node affinity などを利用して容易に再起動すると困る pod がスケジュールされないようにする
- ノードのマシンタイプを決めるための経験則として、1 ノードあたり典型的な pod を最低 5 つ起動でき、かつ取り残されたリソース (小さすぎて新しい pod をスケジュールできない余りリソース) が 10 % 以下になるようにするとよい
    - ノードのスペックが高い方が費用対効果が良いが一台の退役の可用性への影響が大きくなり、スペックが低いと取り残されたリソースが多くなり費用対効果が悪化する

## 6. クラスタの運用

- クラスタのサイジングは難しい問題で、運用しながら調節していくしかない
- Cluster autoscaler は需要に合わせてクラスタを伸縮させる
    - スパイクの激しく無いサービスでは、利用せず手動でのクラスタ調整で十分なことも多い

## 7. Kubernetes の強力なツール

- マニフェストのヘルプは `kubectl explain`
- `--watch` でオブジェクトの監視
- `kubectl logs --timestamps` でタイムスタンプ付与
- デバッグ用の一時的な使い捨て [busybox](https://busybox.net/downloads/BusyBox.html) コンテナを起動する例
    - `kubectl run nslookup --image=busybox:1.28 --rm -it --restart=Never --command -- nslookup demo`
- `COPY --from=busybox:1.28 /bin/busybox` のようにしてイメージ作成時に最低限のシェルやツールを入れることができる
    - 1MB くらい
- `kubectl completion -h` でシェル自動補完を導入するガイドが表示される
    - kubectl に `k` などのエイリアスを設定している場合は `complete -o default -F __start_kubectl k` が必要

## 8. コンテナの実行

- コンテナランタイムを学習のため Go で自作する
    - https://youtu.be/HPuvDm8IC-4
- imagePullPolicy でイメージをプルする際の挙動を指定できる
    - Always, IfNotPresent, Never
- restartPolicy でコンテナの再起動ポリシーを指定できる
    - Always がデフォルトで OnFailure, Never

## 9. Podの管理

- Node Affinity の `requiredDuringSchedulingIgnoredDuringExecution` と `preferredDuringSchedulingIgnoredDuringExecution` はそれぞれ "ハード"・"ソフト"アフィニティと考えると覚えやすい
- k8s のスケジューラはデフォルトではノード間に pod を分散させる力学はなく、必要なら Pod Anti-Affinity を使う
- Taint, Toleration は Affinity とは反対に特定のノードに Pod がスケジュールされないことを設定する
- 各 Affinity など、いずれもスケジューラの最適化を多少なりとも妨げる動きになるので、運用上必要な微調整で切り札的に指定するイメージがよい
- ざっくり、Pod をどのノードに割り当てるかを決めるのがスケジューラで、Pod のライフサイクルを管理するのが Pod コントローラと考えると理解しやすい。以下はコントローラの一部
- DaemonSet はノードごとに 1 つだけの Pod をデーモン的に起動する
- StatefulSet は Pod のレプリカを特定の順序で起動・終了する
    - Headless Service (clusterIP のタイプが None な Service) で特定のレプリカを指定してアクセスできるので、組み合わせて使うことが多い
- Job は一度きり (または指定された回数) Pod を実行する
- Horizontal Pod Autoscaler (HPA) は Pod を水平にスケーリングさせる

## 10. 設定と機密情報

- ConfigMap を更新した際の Pod の再起動について。Deployment はあくまで自身のスペックが更新されないと再起動しないので、例えば ConfigMap 更新時に Deployment のスペックのアノテーションを更新するといたテクニックが必要になる
- 機密情報の管理はマネージドサービス (Hashicorp Vault や AWS Secret Manager 等) を使うのが良いが、小さいチームでは暗号化した情報をリポジトリにコミットする方法も簡便で良い
- その場合、フレームワーク等に依存しないライブラリとして [mozilla/sops](https://github.com/mozilla/sops) がある
    - 全体を暗号化するのではなく `user: ENC[AES256_GCM,data:CwE4O1s=,iv:2k=,aad:o=,tag:w==]` のようになり、どのようなキーが方式で暗号化されているのかレビューしやすくなっていたり、暗号化処理や保存場所のバックエンドには各種サービスやライブラリに対応していたりと、運用が意識されたよくできたソフトウェアに見えた

## 11. セキュリティとバックアップ

- [Kubernetes Security \[Book\]](https://www.oreilly.com/library/view/kubernetes-security/9781492039075/) を読むのがよい

## 12. Kubernetesアプリケーションのデプロイ

- マニフェストのバリデーションを行う [instrumenta/kubeval](https://github.com/instrumenta/kubeval) ([yannh/kubeconform](https://github.com/yannh/kubeconform)) やマニフェストのテストが書ける [open-policy-agent/conftest](https://github.com/open-policy-agent/conftest) といったツールがある

## 13. 開発ワークフロー

- ローリングアップデート中に停止・起動する pod 数の調整には maxSurge, maxUnavailable を使う
- カナリアデプロイの手順の説明が [公式ドキュメントにある](https://kubernetes.io/docs/concepts/cluster-administration/manage-deployment/#canary-deployments)

## 15. オブザーバビリティと監視

- 分散システムにおいては、サービスがアップ・ダウンどちらなのかを単純に 2 値に分類できない
    - 平時から常にどこかがデグレして動いている可能性が高く、またそれでもユーザー影響のないサービスが提供できていれば良い
    - [Ops: It's everyone's job now \| Opensource\.com](https://opensource.com/article/17/7/state-systems-administration)
    - [Gray failure: the Achilles’ heel of cloud\-scale systems \| the morning paper](https://blog.acolyer.org/2017/06/15/gray-failure-the-achilles-heel-of-cloud-scale-systems/)
- オブザーバビリティはこうした背景を受けて現れた概念で、監視アラートの "ダウン" の判定を複雑高度化していくアプローチではなく、各種のメトリクスを常時取得・適切に可視化しておき、問題の事前検知や発生時の原因究明を科学的に進めやすくするといった活動を行うこと (私個人の理解)

## 16. Kubernetesにおけるメトリクス

- メトリクスは大量にあり、必要なものにフォーカスすることがまず重要になる
- 代表的な指針として RED (リクエスト、エラー、持続時間) や USE (使用率、飽和度、エラー) がある
    - 前者は API 志向、後者はミドルウェア志向なイメージ
- Kubernetes での有用なメトリクスの例
    - クラスタ
        - ノード数、ノードあたりの Pod 数、ノードのリソース使用率
    - Deployment
        - Deployment 数、Deployment あたりの Pod レプリカ数、利用できないレプリカ
    - コンテナ
        - リソース使用率、Liveness/Readiness Probe 状態、再起動回数、ネットワークトラフィック、ネットワークエラー
- サービスが複数ある場合は、ダッシュボードのフォーマットも揃えた方が良い
- Pager 対応もトラッキングして適時レビューすると良い
- Prometheus は Borgmon にインスパイアされたソフトウェア

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873119014/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/51wecBhtIOL._SX389_BO1,204,203,200_.jpg" alt="Kubernetesで実践するクラウドネイティブDevOps" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873119014/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Kubernetesで実践するクラウドネイティブDevOps</a></div><div class="amazlet-detail">John Arundel  (著), Justin Domingus (著), 須田 一輝 (監修), 渡邉 了介 (翻訳)<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873119014/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4814400128/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/41gBiqAEGAL._SX389_BO1,204,203,200_.jpg" alt="オブザーバビリティ・エンジニアリング" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4814400128/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">オブザーバビリティ・エンジニアリング</a></div><div class="amazlet-detail">Charity Majors (著), Liz Fong-Jones (著), George Miranda (著), 大谷 和紀 (翻訳), 山口 能迪 (翻訳)<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4814400128/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
