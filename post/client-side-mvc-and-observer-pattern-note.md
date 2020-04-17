{"title":"Ovserver pattern と Client-side MVC あたり","date":"2012-02-25T19:07:29+09:00","tags":["javascript"]}

一定以上の大きさの js を書くと 500 行超えたくらいから汚くなる (脳内メモリには入りきる程度なので破綻はしないけど汚くてテンションが下がる) ので対策を整理したい.

### Observer pattern (publisher/subscriber)

要するに jquery のカスタムイベント. `bind()`, `live()` で独自のイベントとそのハンドラを定義しておいて, `trigger()` で発火させる.

たぶん, この方法がいいのは, イベントの発生源とその処理を疎結合にできること. そうするとイベントを受けて何かするレイヤ (MVC だと M?) を一箇所にまとめられるし, M のテストをしやすい.

詳しく見ていないので後で以下を読んでおく.

- [http://api.jquery.com/bind/](http://api.jquery.com/bind/)
- [http://api.jquery.com/live/](http://api.jquery.com/live/)
- [http://api.jquery.com/trigger/](http://api.jquery.com/trigger/)
- [http://blog.rebeccamurphey.com/2009/12/03/demystifying-custom-events-in-jquery](http://blog.rebeccamurphey.com/2009/12/03/demystifying-custom-events-in-jquery)

### Client-side MVC

フレームワークどころか本も出るくらいメジャーなので勉強しやすいはず.

[JavaScript Web Applications](http://shop.oreilly.com/product/0636920018421.do)

とりあえずフレームワークのチュートリアルを読んで感覚をつかもう.

[Hello Backbone.js](http://arturadib.com/hello-backbonejs/)

このチュートリアルが簡潔で雰囲気がわかった.

- UI 上の要素の単位 (実際にはある DOM のひとつ) 志向で, その要素にモデルとイベントとイベントハンドラを集約する. 論理的意味が要素 (backbone.js では View と呼ばれている) にカプセルされるので, 見通しが良くなる. この View がクラスベース OO だとクラスに相当する印象
  - Visual * とか デルファイみたいな考え方. コンポーネントに対応付けてロジックを書く
  - 見ようによってはここを Controller と呼んで, 単純なテンプレートを View と呼んでもいいかもしれない. spine.js はそういう考え方のようだ.
- View.events でその View 内にある DOM 発のイベントを管理している (ハンドラとの対応付けを行なっている)
- View.constructor でその View 内の Model が発生させるイベントとハンドラ (もちろん定義は View 内) の対応付けを行うのがよい習慣のようだ
- View 同士のやりとりはあるメンバ関数のなかで別 View を new して使っている. (移譲っぽいという捉え方で合ってるだろうか. OOP 自信なし)
- View の説明がメインだったけど, たぶん Model もいろいろカスタマイズできるんだと思う. いわゆるデータストア的に, 見た目を持たないグローバルな Model の使い方は出て来なかった. あるいは View に要素を対応付けずに使う? でも名前的にやりたくない
- いわゆる MVC とはちょっと違うかも. あるいは View + Controller が一体になっているイメージ
  - このイベントにこのハンドラが対応しているというのがもっと明示的になったらいい. ただしどうしていいかわからないけど
    - というのも Model に `set()` すると `change` イベントが発火しますってのがあって, そりゃ知らないよ
- 見てて気づいたけどクライアントサイドのコントローラーに当たる部分って 2 種類あって, url とページ遷移 (single page app だと UI の組み換え) を対応付けるルーターと, DOM とイベントを対応付ける (backbone.js でいう) View があると思う. 現に backbone.js も spine.js もそういう機能があった.
- jQuery, zepto, underscore にべっとりなのも気になる. その点 spine はシンプル
- Model はこのチュートリアルだとオブジェクトに便利関数つきましたよ程度なので, もっとリファレンス読んでみないと判断つかない
  - ajax と localstorage をうまく抽象化してくれる層もあるみたいだから期待
- spine.js, javascriptMVC, agility.js, MooTools のプラグイン, SproutCore あたりは競合? (light-weight vs full-brown の区分けはありそう)
  - あと YUI, Google Crosure Lib にも当たり前のようにこういうフレームワークありそう
  - 言っても軽量のフレームワークだから会社で多人数やるようなときは MooTools とか使ったほうがいいのかも
- こういうアーキテクチャのデザインはどんなアプリを作るケースにマッチするのか・しないのかはやってみないとわからない. ので手を動かそう.

いろいろあるけど勢いのある underscore をまず触っておこう. というわけで次は手を動かしてみる.

### further reading

ライブラリによらないクライアントサイドのアーキテクチャ考察なので, あとで読んでおく.

- [Patterns For Large-Scale JavaScript Application Architecture](http://addyosmani.com/largescalejavascript/)
- [Scaling Isomorphic Javascript Code](http://blog.nodejitsu.com/scaling-isomorphic-javascript-code)

単一ページ型の web app に特化した内容になってきてる部分もあるけどそれはそれで
