{"title":"Chrome Tech Talk Night #4","date":"2012-10-25T00:53:40+09:00","tags":["memo"]}

![](images/20121024200614.jpg)

[Chrome Tech Talk Night #4 を開催します - Google Japan Developer Relations Blog](http://googledevjp.blogspot.jp/2012/10/chrome-tech-talk-night-4.html)

Paul Irish はじめ Google の Chrome チームがやってくるとのことで, 話を聞きに行って来ました. インターネットで見たことある人がいろいろいてスター軍団という感じがありありです.

### Variations on the Mobile Web (Boris Smus)

[Variations on the Mobile Web - Google IO 2012](http://smustalks.appspot.com/japan-12/#1)

[Boris Smus (borismus) on Twitter](https://twitter.com/borismus)

モバイルのウェブ開発をするにあたって対処しないといけない問題とその解決方法を紹介するという内容. どういう問題があって, その解決策にどういうアプローチがあって, それぞれにこういう pros, cons があって, というのを丁寧に説明してくれていて非常に面白かったです. トピックは mobile な web dev な方々にはおなじみなのかもしれませんが自分には新鮮でしたし, なにより問題に対処するアプローチを複数出してきてその中でのベストプラクティスを紹介してくれていたのでかなり興味をそそれれました. 遅刻して途中からしか聞けなかったのが悔やまれます.

- responsive-image
  - client side approach
    - js で src を書き換え
      - look ahead parser 対応がダメ
    - media query
      - 回線速度を考慮できない
  - server side apporech
    - ua で反映するしか無い
  - browser approach
    - image-set
      - [4.8 Embedded content — HTML Standard](http://www.whatwg.org/specs/web-apps/current-work/multipage/embedded-content-1.html#processing-the-image-candidates)
      - 回線速度も考慮してくれる
      - まだ chrome, safari だけ
      - [borismus/srcset-polyfill](https://github.com/borismus/srcset-polyfill)
- device variation
  - たくさんの端末にどう対応するか
  - one version to rule them all
    - デバイス間で差異が
  - a version for each evice
    - デバイス多すぎ
  - responsive design approach
    - css で条件わけ
    - 充分でないこともある
    - form factor に応じて切り替えたい
      - js でやる
  - multiple versions
    - tablet, phone などで分ける
    - modernizr の touch サポートチェックなどを使う
  - [borismus/device.js](https://github.com/borismus/device.js)
  - サーバ上でデバイス判定
    - ua でやる
    - device db
- input variation
  - touch != mouse
  - dev tool でタッチイベントをエミュレートできる
  - touch event
  - mouse & touch 両方をサポートしたい
    - pointer event
    - mouse / touch event -> pointer event に変換される
    - [borismus/pointer.js](https://github.com/borismus/pointer.js)
  - ジェスチャーの実装は大変
    - ライブラリがたくさんある

### Wonderous Web Dev Workflow & Yeoman (Paul Irish)

[Wonderous Web Dev Workflow & Yeoman - Google IO 2012](https://dl.dropbox.com/u/39519/talks/tok-workflow/index.html#2)

[Paul Irish (paul_irish) on Twitter](https://twitter.com/paul_irish)

Yeoman の紹介を中心に, web 開発を楽に楽しくするツールの紹介. とにかく良さ気なツールやサービスが大量に出てきてお腹いっぱいになる発表でした. よさげなツールはあとで消化せねば...

- bulding web app with very enjoynable way
- trivia to valuing tools
- your shell
  - [dotfiles/.bash_prompt at master · paulirish/dotfiles](https://github.com/paulirish/dotfiles/blob/master/.bash_prompt)
- deploy on push
  - github の hook で
- yeoman
  - [Yeoman - Modern workflows for modern webapps](http://yeoman.io/)
  - package management
    - yeoman コマンドでインストールなど
    - 依存パッケージに新しいバージョンが出たら notify
  - generators
    - mvc ライブラリと連動していて, yeoman コマンドからコントローラーを足したりできる
  - Scaffold in a snap
  - Live-recompile, Live-refresh
  - Sass, Coffeescript, AMD & ES6 Modules
  - Run unit tests in headless WebKit via PhantomJS
  - Robust build script
- testing
  - in the browser
    - qunit とか
  - in headless webkit
    - pahtomjs
    - yeoman test
  - in cloud mobile browsers
    - [Cross Browser Testing Tool - BrowserStack](http://www.browserstack.com/)
    - [ryanseddon/bunyip](https://github.com/ryanseddon/bunyip)
  - make test more fun
    - nayncat
- style iteration & devtools
  - sass + livereload
  - chrome devtools support for sass
  - webstorm liveedit
- [Setapp | Share and discover insanely great tools](http://setapp.me/)
  - 好きなツールをシェア
- continually learn how to develop better
- styled console message
- questions
  - 環境をどこまでめんどうみてくれるのか (sass だと ruby とか, node とか)
    - 最初に yeoman をインストールするときにチェックして, 必須のものが無ければインストールを促す
    - インストールを自動化するスクリプトもある
  - dependency を定義してインストールしてくれるようなものはあるのか
    - まだ無い
    - いずれ入る
  - dart や typescript などのサポート
    - まだない
    - yeoman は grunt.js 上で実装しているから拡張はわりと簡単
    - 自分で拡張することも可能
  - grunt から yeoman へのマイグレート
    - yeoman チームでも priority の高い issue
  - yeoman のパッケージインストール元. github 以外で, 例えば自前リポジトリを使えるようにできるか
    - bower が対応してくれている

### LT

- opera のダニエルさんの visibility api の話が面白かった
  - [Using the Page Visibility API - Document Object Model (DOM) \| MDN](https://developer.mozilla.org/en-US/docs/DOM/Using_the_Page_Visibility_API)
  - [Page Visibility](http://dvcs.w3.org/hg/webperf/raw-file/tip/specs/PageVisibility/Overview.html)
  - ページが visible かどうかを通知するイベント
  - タブがインアクティブの時はアニメーションを止めて cpu 負荷を下げたり, などの使い道がある
- [Google Developers Live](https://developers.google.com/live/)
  - google の開発関連の生放送チャンネル
  - その場でコミュニケーションしたり投票したり
