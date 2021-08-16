{"title":"最近読んだもの 12","date":"2021-08-16T20:30:00+09:00","tags":["readings"]}

先週出せなかったので二週分。

## 記事

- [draft\-iab\-protocol\-maintenance\-05 \- The Harmful Consequences of the Robustness Principle](https://datatracker.ietf.org/doc/draft-iab-protocol-maintenance/)
    - robust principle、寛容に受信して厳密に送信するという、今のインターネットの成功に寄与したという有名な原則、に対する批判
    - 初期の相互通信性の確保にはよかったかもしれないが、バグった挙動が維持され、メンテコストはかさみ、新規実装はしづらくなり、標準の改定もしづらくなり、セキュリティリスクも高まる
- [Disks from the Perspective of a File System \- ACM Queue](https://queue.acm.org/detail.cfm?id=2367378)
    - fsync しても永続化されていないケースもあるという話
- [Troubleshoot high replica lag with Amazon RDS for MySQL](https://aws.amazon.com/premiumsupport/knowledge-center/rds-mysql-high-replica-lag/)
    - `SHOW SLAVE STATUS` の `Master_Log_File` と `Relay_Master_Log_File` を `SHOW MASTER STATES` の `File` と見比べて、IO スレッドか SQL スレッドのどちらで詰まっているかを調べるのがなるほどだった
- [Pass secure information for building Docker images \| by Rafael Natali \| Marionete \| Jul, 2021 \| Medium](https://medium.com/marionete/pass-secure-information-for-building-docker-images-8adeafe08355)
    - BuildKit を使って docker build 時に使う秘匿情報をうまく扱う方法
    - k8s の secret と似た方式
- [How does Go calculate len\(\)\.\.? – tpaschalis – software, systems](https://tpaschalis.github.io/golang-len/)
    - Go の len を題材にしてコンパイラの処理を追う
    - 昔 append の実装がどこにあるかわからなくた困ったのをおもいだした
    - 題材の割に簡潔でよかった
    - ジェネリクスなどが来るとこの辺も多分変わるとのことで、確かになと
- [ibraheemdev/modern\-unix: A collection of modern/faster/saner alternatives to common unix commands\.](https://github.com/ibraheemdev/modern-unix)
    - unix コマンドのカッコイイ代替のリンク集
    - netstat じゃなくて ss 使おう、みたいな話ではなかった
- [Backbone Management at Facebook \- Facebook Engineering](https://engineering.fb.com/2021/08/09/connectivity/backbone-management/)
    - facebook のバックボーンネットワークのリスクを予測するモデルを作って、それをもとに増強などの対応の優先度を決めているらしい
    - 仮に増強しようと思ってもリードタイムが数ヶ月から年のオーダーなので
    - 論文は読んでいない
- [An Introduction to Pattern Matching in Ruby \| AppSignal Blog](https://blog.appsignal.com/2021/07/28/introduction-to-pattern-matching-in-ruby.html)
    - ruby2.7からパターンマッチングが導入されていたらしい
    - これまでパターンマッチングがある言語を通らなかったので興味深く読んだ
    - もうちょっとできることを絞ってもいいんじゃないかなとも思った
- [Verify GKE Services are up with dedicated uptime checks \| Google Cloud Blog](https://cloud.google.com/blog/products/operations/verify-gke-services-are-up-with-dedicated-uptime-checks)
    - GCP には uptime check という外形監視機能がある
    - これが最近 GKE Load Balancer に対応した
- [Chris's Wiki :: blog/programming/GoCarefulDesign](https://utcc.utoronto.ca/~cks/space/blog/programming/GoCarefulDesign)
    - Go の言語仕様は注意深く設計されててすごい
    - 他の言語と比べてもそうらしい
- [I hate almost all software](https://tinyclouds.org/rant.html)
    - 10 年前の、いわゆる rant 記事
    - コンテキストはわからないけれど、いいたいことはわかる
    - けど `the experience of the user` に全振りしたソフトウェア作りもそれはそれでおかしなことになると思う
    - 二律背反をなんとかバランスさせながら前に進む人は実際にいるし、そういう行為こそアートだと思う
