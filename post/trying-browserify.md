{"title":"browserify をはじめてみる","date":"2014-11-10T01:05:34+09:00","tags":["javascript"]}

![](images/browserify-wizard.png)

[Browserify](http://browserify.org/) を触ってみたメモ。

### Browserify とは

[CommonJS](http://wiki.commonjs.org/wiki/Modules/1.1.1) のモジュールの仕組み、つまり Node.js の `require` をブラウザ上でも使えるようにするもの、ということでいいみたい。Readme を読む限りは、npm にあるモジュールをブラウザ上にもっていくために作られ始めたような印象をうけるが、ちまたのエントリーをみていると AMD に代わりに CommonJS でフロントエンドの依存関係の管理をする ([RequireJS](http://requirejs.org/) ではなく、Node.js 感覚で `require` 関数をフロントエンドで使う) ためのツールとしても使っていいようだ。

### やりたいこと

- 複数の js ファイルの依存関係を記述したい
- 最終的に、依存関係を考慮した順番で、ひとつの js ファイルに結合したい
- 作りたいのは第三者のサイトに埋め込んでもらうスクリプト (サードパーティスクリプト) である。そのため、
  - グローバル変数をひとつだけ予約していて、その中に必要なプロパティを定義していく
  - ライブラリなどをグローバル変数経由で使うことは不可
  - スクリプトのサイズも極力小さくする必要がある

こういう事情のため RequireJS はあまりマッチしなかった。RequireJS は基本的に自分のサイト・サービスで使うことを意図したプロダクトだと思う。そこまで複雑な依存関係がないのに、ファイルを一つに結合するためだけに数 kb あるコードを同梱するのは避けたかった。

そのため現在は、

- グローバル変数をひとつ予約
- 各ファイルの冒頭で、予約された変数に必要なものを追加。他ファイルのモジュールを使う場合はすでに予約されたグローバル変数内にそれがある前提で処理を記述
- [grunt の concat](https://github.com/gruntjs/grunt-contrib-concat) 相当のものでファイルを結合。
  - つまり依存関係をコードではなくビルド設定側 (ここでは `Gruntfile.js`) に記述している

というトラディショナルな方法で対応している。

browserify によって、依存関係をコード側に持ってこられるならば素敵だなというのが期待。また `require` の実装を同梱することになるはずなので、そのサイズが十分に小さいかも確認する必要がある。

### やったこと

試しに、シンプルに jQuery に依存し、また自前の別モジュールも使うアプリを書くというシナリオを試してみる。`var $ = require('jquery');` と `var foo = require('foo');` できるようになれば OK。

まずは必要なモジュールをインストール

    $ npm install browserify grunt grunt-browserify --save-dev

jQuery は bower で管理していることにする

    $ bower install jquery --save

`package.json` に `browser` というディレクティブを追加する。こうしておくと `require('jquery')` で `./bower_components/jquery/dist/jquery.js` を見に行ってくれるらしい。alias のようなもの。

<pre><code data-language="javascript">"browser": {
  "jquery": "./bower_components/jquery/dist/jquery.js"
}</code></pre>

`Gruntfile.js` は次のようにする。`src.js` をコンパイルして `dest.js` を作っている。

<pre><code data-language="javascript">module.exports = function(grunt) {
    grunt.initConfig({
        browserify: {
            sample: {
                files: {
                    "dest.js": ["src.js"]
                },
            }
        }
    });

    grunt.loadNpmTasks("grunt-browserify");
    grunt.registerTask("default", ["browserify"]);
};</code></pre>

`src.js` はこのような内容にした。バージョン番号がログに出る。

<pre><code data-language="javascript">(function() {
    var $ = require('jquery');
    var from_lib = require('./lib.js');
    console.log($().jquery);  // version 番号
    from_lib();
})();</code></pre>

ロードされている `lib.js` の内容。

<pre><code data-language="javascript">module.exports = function() {
    console.log('from lib');
};</code></pre>

これで grunt をキックすれば動作し、`dest.js` が生成される。`<script src="dest.js"></script>` のような html で動作確認できる。

#### 結果

やってみたところ、確かに意図どおり動作した。`$` や `jQuery` 変数はグローバルに出ておらず、ちゃんとスコープの中に閉じられていた。目で見た感じでは browserify によって挿入されたコードはだいたい 500 byte 以下のようだ。よってなんとなく使えそうな気がする。

何もしなければ jQuery は `$` 変数などを `window` オブジェクトに作るはずで、これをどうやって抑制しているのかが気になった。

#### browserify-shim

CommonJS 対応していないライブラリを使うための shim で [browserify-shim](browserify-shim) というものもあるらしい。例えば jQuery に依存している jQuery-UI など、グローバル空間にある依存ライブラリがある前提で作られているライブラリのためのもの。例えばある X ライブラリの `$` 変数にに依存する Y ライブラリがあったとすると `package.json` に

<pre><code data-language="javascript">"browserify-shim": {
  "X": "$",
  "Y": {"exports": "Y", "depends": ["X:$"]}
},</code></pre>

と記述すると、X の `$` をグローバル空間に export してくれる、というものらしい。(たぶん)

今回は用途が違うが、通常のフロントエンド開発では必須な気がする。

### 次にやること

- browserify によって同梱されるコードがどれくらいなのか正確にしらべる
- なぜ jQuery の `$` などがグローバルにできていないのかをしらべる
- [napa](https://github.com/shama/napa) というモジュールを使うと、npm にはないモジュールを `package.json` 経由で `node_modules` にインストールできるらしい
  - これを使うとアプリに必要なモジュールをすべて `package.json` に記述できる気がする
  - そういうものは素直に `bower.json` など別管理にしたほうがいいのか、こういうものを使って `package.json` に寄せたほうがいいのか検討

browserify も browserify-shim もちゃんと構文解析して動作しているらしいので、実装も気になる。

### 参考

- [Browserify](http://browserify.org/)
- [Grunt-Browserify 2.x and Browserify-Shim - ÆFLASH](http://aeflash.com/2014-05/grunt-browserify-2-x-update.html)
- [substack/browserify-handbook](https://github.com/substack/browserify-handbook)
- [thlorenz/browserify-shim](https://github.com/thlorenz/browserify-shim)
- [shama/napa](https://github.com/shama/napa)

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00IOGV3XU/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="http://ecx.images-amazon.com/images/I/517dA3j6YbL._SL160_.jpg" alt="サーバサイドJavaScript　Node.js入門 アスキー書籍" style="border: none;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00IOGV3XU/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">サーバサイドJavaScript　Node.js入門 アスキー書籍</a><div class="amazlet-powered-date" style="font-size:80%;margin-top:5px;line-height:120%">posted with <a href="http://www.amazlet.com/" title="amazlet" target="_blank">amazlet</a> at 14.11.09</div></div><div class="amazlet-detail">KADOKAWA / アスキー・メディアワークス (2014-02-27)<br />売り上げランキング: 9,371<br /></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00IOGV3XU/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
