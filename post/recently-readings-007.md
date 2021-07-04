{"title":"最近読んだもの 7","date":"2021-07-04T23:30:00+09:00","tags":["readings"]}

## 記事

- [How is software developed at Amazon?  \- High Scalability \-](http://highscalability.com/blog/2019/3/4/how-is-software-developed-at-amazon.html)
    - ピザ2枚のチーム・システムから始まり、それを分割していく（顧客のニーズに合わせて）
        - チームには目標と必要なリソースが与えられ、それぞれがひとつのスタートアップのように振る舞う。マネジメントは複数のスタートアップをみる取締役のような立ち位置になる
    - エンジニアは、開発者、アーキテクト、運用者、テスター、セキュリティスペシャリストとして振る舞うことを期待される
    - チームは職能を跨いだメンバーで構成。エンジニア、PM、マーケターなど
    - 計画はボトムアップに（各チームが一番ユーザーをわかっている）
    - チーム間のコンフリクトはマネジメントの階層構造で調整
        - 何もやらないより、リスクを織り込んで何かを出す方を選ぶ。折り込まれたリスクはのちに別のリファクタチームが解消する。スピードを維持する
        - そもそもとしてチームが分かれるとコミュニケーションと一貫性は難しい
        - 全体方針はトップダウンで、それへの向かい方は各チームで
- [QUIC at Snapchat \- Snap Engineering](https://eng.snap.com/quic-at-snap)
    - QUIC 導入はかっこいい施策だが、あくまでネットワークのレイテンシーとエラー率の改善活動の一環として位置付けられているのがよい
        - プロトコルの変更以外にも `make requests and responses smaller, reduce unnecessary sync, utilize global content distribution partners to bring media close to the people who use it` などやれる事はたくさんある
    - 導入による改善は、いずれの指標も 95 パーセンタイル以上で 5 ~ 25 % くらい
        - この桁感は覚えておくと参考になりそう
- [Security headers quick reference](https://web.dev/security-headers/)
    - まとまってる
- [curl vs Wget](https://daniel.haxx.se/docs/curl-vs-wget.html)
    - curl 作者の [@bagder](https://mobile.twitter.com/bagder) による比較
    - curl は cat、wget は cp の analogue とするのは確かにすんなり納得できた
        - `curl works more like the traditional Unix cat command, it sends more stuff to stdout, and reads more from stdin in a "everything is a pipe" manner. Wget is more like cp, using the same analogue.`
    - wget は busybox に入っていて便利というのはすごいそう思う
    - curl はもう http/3 もサポートしていてすごい
- [Open Source Insights](https://deps.dev/)
    - 各種ライブラリの依存関係などが確認できるツールらしい
    - Google が各種パッケージマネージャの情報をインデックスしているらしい
- [Revisiting the Twelve\-Factor App Methodology — Coder Society](https://codersociety.com/blog/articles/twelve-factor-app-methodology)
    - 12 factor apps が出てからもう 10 年とは...
- [Application Frameworks \| July 2021 \| Communications of the ACM](https://cacm.acm.org/magazines/2021/7/253463-application-frameworks/fulltext)
    -  当たり前のことが書かれているだけに見えた...

## 動画

- [Accelerate networking with HTTP/3 and QUIC \- WWDC21 \- Videos \- Apple Developer](https://developer.apple.com/videos/play/wwdc2021/10094/)
    - もうちょい待てば普通に使えるようになるというのがすごい (サーバ側も対応していれば)
    - 途中で出てくる通信のモニタリングツールが便利そうで気になった
        - これかな?
        - [Analyze HTTP traffic in Instruments \- WWDC21 \- Videos \- Apple Developer](https://developer.apple.com/videos/play/wwdc2021/10212)

## 雑誌

- [WEB\+DB PRESS Vol\.123 \| 後藤 ゆき, 古川 陽介, 吉井 健文, 藤原 涼馬, 雑司ヶ谷, 西山 和広, 五十嵐 進士, 佐藤 歩, 櫻庭 祐一, James Van Dyne, うたがわ きき, 牧 大輔, 池田 拓司, 是澤 太志, 関 満徳, はまちや2, 竹原, WEB\+DB PRESS編集部 \|本 \| 通販 \| Amazon](https://www.amazon.co.jp/dp/4297122073/)
    - 特集 1 の HTTP/3 入門目当て
        - [@flano_yuki](https://twitter.com/flano_yuki) さん著
    - 短い紙面で概要がまとまっていて良かった
        - これからキャッチアップするためのスタート地点として

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B097D7WMDB/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/51-RGvu8GUS.jpg" alt="WEB+DB PRESS Vol.123" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B097D7WMDB/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">WEB+DB PRESS Vol.123</a></div><div class="amazlet-detail">WEB+DB PRESS編集部  (編集)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B097D7WMDB/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

## コード

- [rs/zerolog: Zero Allocation JSON Logger](https://github.com/rs/zerolog)
    - `Zerolog's API is designed to provide both a great developer experience and stunning performance. Its unique chaining API allows zerolog to write JSON (or CBOR) log events by avoiding allocations and reflection.` とのことで `unique chaining API` と `avoiding allocations` とはどういう感じか気になったので流し読み
    - `unique api` に関しては有名な [uber の zap](https://github.com/uber-go/zap) で発明された、メソッド呼び出しを chain できる api のこと
        - chaining 自体は一般的なインタフェースだけど、確かにロギングの api でこのようなものはなかったかもしれない
        - 確かにこれは書きやすそう・読みやすそう
            - ログへの出力要素の追加やメタデータの設定がスッキリ書ける
    - パフォーマンスに関しては [sync.Pool](https://golang.org/pkg/sync/#Pool) を使って確保したメモリを再利用していることと、出力前のログ内容を単なる `[]byte` に追記していっている (内部で構造体などで保持していない) のがポイントだった
        - [Event](https://github.com/rs/zerolog/blob/6ed1127758fabcfdd3c27f2c7b788f2dd9378d84/event.go#L22) がひとつのログを表すクラス
        - ロギングする内容は [`buf []byte`](https://github.com/rs/zerolog/blob/6ed1127758fabcfdd3c27f2c7b788f2dd9378d84/event.go#L23) に保持している
            - `buf` は Event インスタンス作成時に [pool から取得](https://github.com/rs/zerolog/blob/6ed1127758fabcfdd3c27f2c7b788f2dd9378d84/event.go#L59) しログ出力後 (Event インスタンス利用完了後) に [pool に戻している](https://github.com/rs/zerolog/blob/6ed1127758fabcfdd3c27f2c7b788f2dd9378d84/event.go#L81)
            - `buf` のサイズは [500 byte からスタート](https://github.com/rs/zerolog/blob/6ed1127758fabcfdd3c27f2c7b788f2dd9378d84/event.go#L15) する
        - `buf` にはログの要素を諸々処理してから append していっている
            - 例えば `Str()` で要素を足すと、必要なエスケープなどを行ったあとに [buf に append する](https://github.com/rs/zerolog/blob/117cb53bc66413d9a810ebed32383e53416347e3/internal/json/string.go#L58)
            - [同じキーを重複して登録できてしまう](https://github.com/rs/zerolog#caveats) 仕様はここに由来する
        - 以上よりメモリ確保が走るのは次のケース
            - pool に確保済みの buf がない場合、新規で確保する
                - pool の初期化時に複数確保したりはしていない
            - 500 byte を超えるログを作成する場合、buf を延長する
                - ロジックは通常の Go の append と同じ
    - 基本的には json 出力形式のみサポートされているのは今どきでよい
    - ちなみに sync.Pool の各要素は [ほぼ同じメモリコストにしないとにしないと非効率](https://github.com/golang/go/issues/23199) らしく、[buf が一定以上大きいと pool に戻さない](https://github.com/rs/zerolog/blob/6ed1127758fabcfdd3c27f2c7b788f2dd9378d84/event.go#L39-L42) という制御が入っていた
        - sync.Pool の実装は見ずにあてずっぽうだけど、でかいバッファが gc されずに確保され続けることにもなり得そうなので、ハードリミットがあるのは安心感がある
- [kubernetes\-sigs/kustomize: Customization of kubernetes YAML configurations](https://github.com/kubernetes-sigs/kustomize)
    - k8s の yaml を外側から色々操作できる小粋なツール
    - どうやって yaml のパースやその値の操作をやっているのか気になってそこだけさらっと見た
        - 自分で頑張っていた
    - kustomization.yaml の定義は以下
        - https://github.com/kubernetes-sigs/kustomize/blob/d818ccae92e92b0de3b770a7f1e1a0f53b0dd2e0/api/types/kustomization.go#L23
        - この内容に沿って対象を操作する
    - yaml のパース、AST 作成、AST 変換などは [kyaml](https://github.com/kubernetes-sigs/kustomize/blob/d818ccae92e92b0de3b770a7f1e1a0f53b0dd2e0/kyaml/doc.go) という大きなパッケージが担当している
        - 細かくは全く追えていないが、例えば https://github.com/kubernetes-sigs/kustomize/blob/d818ccae92e92b0de3b770a7f1e1a0f53b0dd2e0/kyaml/yaml/example_test.go のような感じ
