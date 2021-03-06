{"title":"yapcasia 2012","date":"2012-09-30T20:31:39+09:00","tags":["memo"]}

![ypacasia 2012 main hall](images/20120929141034.jpg)

[HOME \| YAPC::Asia Tokyo 2012](http://yapcasia.org/2012/)

yapc には初めて行きましたが, perl コミュニティの凄さを感じました. miyagawa さんが他のコミュニティのいけてるところを perl コミュニティにも持ってこようと発表していましたが, 自分も perl コミュニティから学んで活かしていきたいですね.

### 個人的あとで要復習の発表

- [Web::Security beyond HTML5](http://yapcasia.org/2012/talk/show/aaade824-abc0-11e1-8865-57a46aeab6a4)
- [「新しい」を生み出すためのWebアプリ開発とその周辺](http://yapcasia.org/2012/talk/show/dbe56d5c-ac3a-11e1-8ef2-22926aeab6a4)
- [スマートフォン向けサービスにおけるサーバサイド設計入門](http://yapcasia.org/2012/talk/show/3a843f6e-dabc-11e1-9422-0d4e6aeab6a4)

### memo

#### 前夜祭
- [Officeで使うPerl Excel編](http://yapcasia.org/2012/talk/show/c8c6ea6c-d5a6-11e1-aeff-37a36aeab6a4)
  - windows で perl
    - どれでも OK (activeperl, starawberry perl など)
  - module
    - Wind32::OLE
    - SpreadSheet::*
    - Excel::*
- [UV - libuv binding for Perl](http://yapcasia.org/2012/talk/show/9a662188-da47-11e1-82d2-0d4e6aeab6a4)
  - 一般的な async io 抽象化ライブラリ
    - epoll, kqueue, event port
    - なかったら select, poll
    - win だとたいてい select しかなくて遅い
  - win にも IOCP という高速な API がある
    - [Asynchronous I/O in Windows for Unix Programmers](http://tinyclouds.org/iocp-links.html)
    - ただしインタフェースが unix 系とぜんぜん違う
    - そのため libuv は高い抽象レベル
    - ほとんど独自のフレームワークといったレベルで, 逆に syscall に近い既存のライブラリよりもカジュアル
  - kayac での事例
    - スマホアプリの通信モデル構築
      - iOS 対応のパッチをコミット
  - perl binding
    - [typester/p5-UV](https://github.com/typester/p5-UV)
    - libuv アプリのプロトタイピング用途に
    - api がほぼ一対一のデザインで, プロトタイプからの移植が簡単
  - document
    - libuv 本家もドキュメントは uv.h 読め状態
  - AnyEvent との比較
    - perl でやるんならこれでいい
    - 枯れている, バックエンドを選べる, ドキュメント充実, モジュール充実
  - livuv の AnyEvent 対応
    - libuv の抽象度が高すぎて難しい
  - 開発状況
    - 通信周りは done
    - ファイルIO まわりは未実装
    - win でビルド通るか不明
    - 自分が欲しいところだけ作った
- [Sqale の裏側](http://yapcasia.org/2012/talk/show/048d468c-ab9e-11e1-a3b5-2a656aeab6a4)
  - container, lxc
    - それぞれで nginx, sshd, supervisord
  - amazon linux + patched kernel (3.2.16)
    - grsecurity
    - tcp port bind, fork bomb の抑制をする自前パッチ
  - web proxy
    - ELB -> nginx -> container
    - nginx
      - lua-nginx-module
      - redis2-nginx-module
        - リクエスト url がどのコンテナに入っているかを redis で持っている
        - そこのロジックを lua で
  - ssh router
    - [mizzy/openssh-script-auth](https://github.com/mizzy/openssh-script-auth)
    - SSH_ORIGINAL_COMMAND

#### 一日目
- [What Does Your Code Smell Like?](http://yapcasia.org/2012/talk/show/40330888-db88-11e1-8417-0d4e6aeab6a4)
  - ランダムなシーケンスをソートするコードを perl6 に書き換える
  - 関数のプロトタイプ
  - everything is reference
  - explicit reference
- [Web::Security beyond HTML5](http://yapcasia.org/2012/talk/show/aaade824-abc0-11e1-8865-57a46aeab6a4)
  - ajax
    - xss, データの盗聴
  - xss
    - ie の content-type 無視
      - html でないものが html に昇格して xss
      - text/plain が html として解釈されたり
      - ie は最終的にファイルタイプに基づいてコンテンツを処理する
      - そのメカニズムは文章化されていなくて複雑
        - content-type, x-content-type-option, windows レジストリの関連付け, ie の設定, url 自身, コンテンツそのもの
      - 例外も多く, よく挙動も変わる
    - xss が発生しがちな状況
      - json 文字列内, jsonp コールバック名, text/csv はそのそもエスケープできない
  - データの盗聴
    - json array hijacking for android
      - android の古い (2.2, 2.3) で古い攻撃手法が通用する
        - property setter の再定義
    - json hijacking for ie
    - <script src="target.csv">
  - 対策
    - x-content-type-options: nosniff
      - レスポンスヘッダ
      - ie がコンテントタイプだけで判断するようになる (ie8 以降)
    - xhl 以外ははじく
      - x-request-with リクエストヘッダを見る (jquery, prototype)
      - 任意のヘッダでも
  - waf 使ってると, xhr とかに限定できないし, 泥臭く
    - json の過剰エスケープなど
  - xhr2 の問題
    - url チェックが必要
      - めんどうくさい web セキュリティ
    - origin ヘッダを信用しない
      - `xhr.withCredentials = true` してクッキー付きでやりとり
      - リクエストヘッダに秘密の情報を含める
  - ブラウザの保護機構
    - xss フィルタ
      - リクエストとレスポンスに同一のスクリプトが含まれる場合にブロック
      - x-xss-protection: 0 でフィルタ無効化. 誤検出対策に
      - webkit の xss auditor はユーザに通知せず, 勝手に消す
    - x-content-type-options: nosniff
      - 原則つけておくべき
    - クリックジャッキング対策
      - 標的サイトに透明 iframe を重ねて攻撃者のリンクを押させる
      - x-frame-options: deny, x-frame-options: sameorigin
        - モダンブラウザで対応 (ie8+)
        - たいていのケースではいれておけばOK
    - content-security_policy (csp)
      - 想定された以外のリソースが読めない
    - https の強制
      - hsts
        - 両方対応している場合は https を使うようレスポンスヘッダでクライアント側に要請
- [rapid prototyping with Mojolicious::Lite](http://yapcasia.org/2012/talk/show/5da12cf6-adde-11e1-aff4-22926aeab6a4)
  - idea -> prototype -> test / pivot -> overnight success
    - minimize idea <-> success
  - Mojolicious::Lite + CouchDB + Bootstrap
  - package contents
    - login/user handling, gmail/googleapp support, micro cms, cdn storage
- [大きくなったシステムの疎結合化への取り組み](http://yapcasia.org/2012/talk/show/4d98200c-dae4-11e1-b6d8-0d4e6aeab6a4)
  - 複雑化したコードベースへの対応
  - アーキテクチャ
    - ServiceProcedure
- [Perlと出会い、Perlを作る](http://yapcasia.org/2012/talk/show/37e5eabc-d550-11e1-ace3-37a36aeab6a4)
  - KonohaScript
    - 静的型付け
    - オブジェクト指向
  - 深く理解したいなら作る
    - gperl
    - 世界一高速な perl 処理系を目指す
  - ベンチマーク
    - binary-trees
    - fibonacci
    - たらい回し関数
      - JIT いれて v8 より早い
  - 高速化技術
    - 引数オブジェクトの事前構築
    - NaN-boxing による型検査
- [App::LDAP - 管理者と百台のコンピュータ](http://yapcasia.org/2012/talk/show/5425192c-d411-11e1-8310-37a36aeab6a4)
- [DBD::SQLite: Recipes, Issues, and Plans](http://yapcasia.org/2012/talk/show/7dfaa352-afbd-11e1-9133-ef7b6aeab6a4)
- [スマートフォン向けサービスにおけるサーバサイド設計入門](http://yapcasia.org/2012/talk/show/3a843f6e-dabc-11e1-9422-0d4e6aeab6a4)
  - スマホ専用アプリで利用する SNS を想定
    - メッセージ, 写真投稿. いいね. 友達のフィードなど
  - そもそもアプリでなければならない?
    - カメラ・GPS などデバイス使用の有無
    - 機能毎に処理を行うのがサーバ・クライアントどちらか
  - pro/con
    - native
      - デバイス使いやすい
      - アプリマーケット
      - リッチなユーザー体験
    - browser base
      - アップデートが容易
      - システム構成がシンプル
  - 見える側から裏側へと考えていく
    - 画面 / UI
      - 自分で触ってみる
      - ガイドライン
    - Web API
      - method + param/body + response
      - restful
      - 既存サービスを参考 (facebook, twitter, google+)
      - 既存フレームワークを参考 (catalyst, ror)
    - サーバサイド
      - small start
      - 並列に追加できるように
      - データストア
        - mysql
        - redis
        - memcached
      - lb -> [nginx/starman/memcachd] + lb -> dbmaster/dbslave + lb -> kvs master/kvs slave
      - ユーザ管理
        - セッションは cookie 管理
          - oauth に比べシンプルになるのでまずはこれで
        - アプリ削除でそのまま接続できなくなったセッションの処理
  - perl dev
    - 手に馴染んでいた
    - cpan modules
    - perlbrew, carton
    - amon2
      - 標準でセキュリティまわりが充実しているのがいい
  - amon2::Setup::Flavor::Large でスケルトン生成
  - db
    - orm は teng
    - はじめは orm でさっとつくって, 固まってから sql に
  - レスポンスオブジェクト
    - amon2::plugin::wb::json
    - perl -> json は数値の扱いを注意
  - エラー・例外
    - Class::Exception
      - 共通でこの例外クラスで
    - dispatcher, worker レベルで例外キャッチ
  - テスト
    - model, 共通関数はテストを
    - テストを書きすぎてもきりがないのでバランス
  - 環境ごとの切り分け
    - amon2 の昨日で config を切り分け
    - 同一ドメインで全部サーブする場合はの場合は個別の psgi
  - 本番環境
  - 時間切れ...
- [Redmine::Chan で IRC からプロジェクト管理](http://yapcasia.org/2012/talk/show/7b6375aa-bd29-11e1-ad51-b39f6aeab6a4)
  - redmine 機能豊富だが操作は煩雑
  - とにかく簡単にする
  - irc
  - Redmine::Chan
- [Stackato - ActiveState's new host-your-own-PaaS](http://yapcasia.org/2012/talk/show/fdb4fa30-da91-11e1-8db4-0d4e6aeab6a4)
- [半リアルタイム・分散ストリーム処理をperlで](http://yapcasia.org/2012/talk/show/96627a88-ab9d-11e1-a255-2a656aeab6a4)
  - stream
    - => tail -f
  - ストリーム処理
    - => tail -f | grep hoge | sed fuga | ...
    - これをもっとスケーラブルに, 安全に
    - 特徴: EOF を待たない
  - ネットワーク越しのストリーム処理
    - 重い処理だけ切り出したりしたいから. そうしないとつまる
    - edge: tail -f | nc, backend: nc -l | grep hoge | sed fuga | ...
  - なぜストリーム処理
    - 例えばアクセスログ集計
    - 1 時間ごとに処理だとする
    - flush wait (sevral min) -> copy over nw -> convert / parse
    - 16 時台のログ処理結果がわかるのは1時間以上あとになってしまう
    - ストリーム処理ならば準リアルタイムにわかる
    - 突然ログサイズが大きくなっても影響少
      - もちろんネットワーク帯域などの問題はあるが
  - バッファリング
    - どの程度バッファするか

#### 二日目
- [Perl 今昔物語](http://yapcasia.org/2012/talk/show/ea24040c-db90-11e1-8a64-4c8f6aeab6a4)
- [10 more things to be ripped off](http://yapcasia.org/2012/talk/show/bcdf8c9e-da9d-11e1-a79e-0d4e6aeab6a4)
  - 他の言語から取り入れたほうがいいもの
  - vimeo.com/47344146
  - debug, InteractiveDebuger
    - `-e 'enable Debug'`
      - from django
    - `-e 'enable InteractiveDebuger'`
    - `-e 'enable REPL'`
  - rubygem, pypi, npm の勢い
    - more to steal
  - ruby ecosystem
    - bundler
      - carton
    - gemnasium
    - gemfury
  - web ecosystem
    - asset pipleline
  - templating
    - Tilt
    - tiffany
  - continuous integration
    - travis ci
  - testing
    - capybara
  - pry
  - pow
  - new relic
- [Perlハッカーは息をするようにCPANモジュールを書く。](http://yapcasia.org/2012/talk/show/72ef69fe-da45-11e1-abbb-0d4e6aeab6a4)
  - ドキュメントを書こう
    - モチベーションを含めて
  - こつ
    - pod
    - 真似る
    - たくさん書く
  - 書くべき内容
    - synopsis
    - description
      - motivation はここで
    - methods / functions
  - cpan モジュール作成のポイント
    - ドキュメントだけはしっかり
      - モチベーションを書くことで, 他の似たモジュールとの使い分けを伝える
    - メンテナンスをしっかり
- [Perlで始める！初めての機械学習の学習](http://yapcasia.org/2012/talk/show/14ac724c-db17-11e1-b032-0d4e6aeab6a4)
  - 機械学習 =>python
  - PRML
  - herumi/prml
  - PRML の学習
- [中規模ソーシャルゲーム開発に学ぶWebサービス開発と運用ノウハウ。もしくは2012年にPerlでWeb開発をする理由](http://yapcasia.org/2012/talk/show/1cc88df4-da3a-11e1-b11f-0d4e6aeab6a4)
  - ぼくらの甲子園シリーズ
    - 累計160万ユーザ
    - 2年運営
  - codebase
    - source: 10万
    - controller: 80
    - action 500
  - ORM 必須
    - DBIx::Class
  - アプリのドメイン・目的に応じてモジュールを選定する
  - infra
    - web * 2 (nginx)
    - app * 6  (starman, haproxy)
    - db * 5 (mysql)
    - ...
    - chache, session storage
  - web
    - nginx
    - reverse proxy
    - static file
  - app
    - server::starter
    - starman
    - daemontools
    - haproxy
      - parallel::benchmark + furl で手軽にベンチ
    - swfeditor
      - flash 生成
    - fluentd
      - 行動ログ収集
    - Ark (waf)
  - db
    - mysql
    - master * 1, slave * 4
    - partitioning
  - cache
    - kt
      - session storage
      - dual master
      - ulog の削除注意
    - memd
      - flash のキャッシュなど
  - batch
    - gearman
    - daemontools
  - admin
    - archer
      - deploy ツール
    - mongo
      - fluentd から来た行動ログ
  - マスタ系データ管理
    - google spreadsheet
    - Data::GoogleSpreadsheet::Fetcher (Net::Google::Spreadsheets)
    - スプレッドシートの内容を yaml or csv に落としてバリデーション, その後 db へ流す
  - testing & ci
    - Test::mysqld
    - testfile ごとにトランザクション開始, 終わったらロールバック
    - jenkins
      - war を起動するスクリプトを deamontools 管理が楽
    - spreadsheet を fixture
  - cap, chef
  - 2012 年における perl 開発
