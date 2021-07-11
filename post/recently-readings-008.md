{"title":"最近読んだもの 8","date":"2021-07-11T22:30:00+09:00","tags":["readings"]}

## 記事

- [mnot’s blog: On HTTP Load Testing](https://www.mnot.net/blog/2011/05/18/http_benchmark_rules)
    - 10 年前の記事だけど有用
    - 計測方法が悪いものはミスリーディングなので有害ですらある、というのは確かに
- [r/WallStreetBets Incident Anthology: More Data, More Problems : RedditEng](https://www.reddit.com/r/RedditEng/comments/o4yb4z/rwallstreetbets_incident_anthology_more_data_more/)
    - 特定のオブジェクトが巨大で連鎖的に障害に
- [The Python Paradox](http://www.paulgraham.com/pypar.html)
    - 仕事を得るためにそれをやっている人が多い技術なのか、技術そのものが好きでそれをやっている人が多い技術なのか
    - それ以外にもいろいろ要素はあるし、最近では単に言語だけでは人をフィルタできないんじゃないかなあ。傾向はあるにせよ
    - 現象の名前として python paradox というのを覚えておくと会話のネタにはなりそう
- [How to use CoreDNS Effectively with Kubernetes](https://www.infracloud.io/blogs/using-coredns-effectively-kubernetes/)
    - resolv.conf の ndots オプションを知らなかったので勉強になった
    - [resolv\.conf\(5\): resolver config file \- Linux man page](https://linux.die.net/man/5/resolv.conf)
        - そのドメインに含まれるドット数が ndots より少なければ search する
- [13 Best Practices for using Helm — Coder Society](https://codersociety.com/blog/articles/helm-best-practices)
    - helm の何が嬉しいかなどがちょっとわかった
    - こんなに自由度が高くて大丈夫だろうかとも思った
- [Understanding How Rbenv, RubyGems And Bundler Work Together \- Honeybadger Developer Blog](https://www.honeybadger.io/blog/rbenv-rubygems-bundler-path/)
    - 理解が曖昧なまま嵌るとほんとに時間を溶かすので、あらためて理解を整理できて良かった
    - `RubyGems monkey-patches Kernel.require`
        - これは知らなくてびっくりした
- [Understanding Factory Bot syntax by coding your own Factory Bot \- Code with Jason](https://www.codewithjason.com/understanding-factory-bot-syntax-coding-factory-bot/)
    - メタプログラミングの理解にいいチュートリアル
    - まだ慣れていなくてライブラリのコードを追う時に苦労しているので、読めて良かった

## 動画

- [Reduce network delays for your app \- WWDC21 \- Videos \- Apple Developer](https://developer.apple.com/videos/play/wwdc2021/10239/)
    - ネットワーク検査ツール `/usr/bin/networkQuality`
    - 帯域が太くてもレイテンシーが悪いこともある
        - 昨今のスピードテストは帯域を測っている
        - レイテンシーは中間装置のバッファの大きさ、サーバ処理時間、伝播時間、光速などによって制約されている
    - 対策
        - 一般の開発者にとってコントローラブルなのはリクエスト数の削減
            - 例えばレイテンシが 600 ms という悪いコンディションで接続確立時に4回のラウンドトリップが必要な場合、2.4 sec もかかってしまう
            - モダンなプロトコルを採用しラウンドトリップ 1 回で済むようにしたら 600 ms で済む
            - モダンなプロトコル
                - http/3 over quic, tcp fast open, tos1.3, multipath tcp
                - すべてサーバ側もサポートする必要
            - サーバ側を冪等な作りにしておくのも大事
        - バックグラウンドで通信する
            - ユーザー操作起点でないプリフェッチリクエストや、大きい非同期的なファイル転送などは、バックグラウンドで行うよう os に指示できる
                - URLRequest の networkServiceType パラメタ
            - キューが埋まるのを緩和できる
        - テストする
            - network link conditioner を使って調べると良い
- [Analyze HTTP traffic in Instruments \- WWDC21 \- Videos \- Apple Developer](https://developer.apple.com/videos/play/wwdc2021/10212)
    - 新しい通信の解析ツールの解説
    - 後で試したい

## コード

- [jamesmoriarty/forward\-proxy: Minimal forward proxy using 150LOC and only standard libraries\.](https://github.com/jamesmoriarty/forward-proxy)
    - 150 行くらいで実装されたフォワードプロキシという触れ込み
    - via ヘッダは知らなかった
    - スレッドプールの実装がいまいち理解できていない
        - ノリはわかるけど本当か理解が積み重なってない
        - スレッドに accept 後のハンドリング処理を登録しているので、最初にプールにスレッドをいくつか準備しておき、accept するごとに引き渡していってるのかなとは思うけど、詳細がわかっていない
