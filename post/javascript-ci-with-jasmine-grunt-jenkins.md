{"title":"Jenkins, Grunt, Jasmine で JavaScript のテストを CI","date":"2014-02-23T23:29:01+09:00","tags":["javascript"]}

jasmine で書いた js のテストを grunt で実行できるようにして jenkins で継続的に動作させる。

### Gruntfile.js

まず Gruntfile.js には以下のようにする。

<pre><code data-language="javascript">module.exports = function(grunt) {
    grunt.initConfig({
        jasmine: {
            all: {
                src: 'js/source/**/*.js',
                options: {
                    specs: 'js/spec/**/*.spec.js',
                    helpers: ['js/spec/helpers/jquery-1.11.0.min.js',
                              'js/spec/helpers/jasmine-jquery.js'],
                    keepRunner: true,
                    junit: {
                        path: 'build/jasmine-test/'
                    }
                }
            }
        }
    });
    grunt.loadNpmTasks('grunt-contrib-jasmine');
};</code></pre>

- `js/source/**/*.js` にテスト対象のコード
- `js/spec` 以下にテスト関連のファイル
   - `js/spec/*.spec.js` にテストコード
   - `js/spec/helpers` にテストコードのヘルパー
   - `js/spec/fixtures` に fixture。今回はテストデータとなる html。
- `build/jasmine-test/` にテスト結果を出力

というディレクトリ構造にした。今回のテスト対象は外部のライブラリに依存していないが、たとえば jquery に依存しているなど外部のライブラリを使っているならば、`js/spec/vendors` 以下に依存ライブラリを入れるべきだ。

同様の理由で `jasmine-jquery` とそれが依存している `jquery` は、あくまでテストのためのライブラリで、本番のコードは依存していないので、`helpers` の下に配置している。

また今回は歴史的な事情でこのようなディレクトリの命名になっているが、`spec/javascripts/*` のような rails 文化に沿った規約を採用すると、よりわかりやすいかもしれない。

### テストコード

テストコードの例。これを `js/spec/your_method.spec.js` などとして保存。

<pre><code data-language="javascript">describe("yourMethod", function() {
    beforeEach(function() {
        jasmine.getFixtures().fixturesPath = 'js/spec/fixtures';
        loadFixtures('your_method.fixture.html');
    });

    it("should work fine", function (){
        // do something
    });
});</code></pre>

fixture のロードには `jasmine-jquery` の `loadFixtures` を使っている。`loadFixtures` はデフォルトで `spec/javascripts/fixtures` 以下のファイルを探しに行く。今回は `js/spec/fixtures` にフィクスチャを入れているので、`fixturesPath` プロパティを設定している。

fixture の取り扱いが便利なので、`jasmine-jquery` の fixture 系の機能だけを使っている。その他の jquery 的な機能は個買いは使っていない。

### jenkins

jenkins のビルドの設定を以下のようにしておく。

    npm install
    grunt jasmine

ビルド後の処理として `JUnitテスト結果の集計` で `build/jasmine-test/*.xml` を指定する。

### 参考

- [gruntjs/grunt-contrib-jasmine](https://github.com/gruntjs/grunt-contrib-jasmine)
- [velesin/jasmine-jquery](https://github.com/velesin/jasmine-jquery)
- [marionettejs/backbone.marionette](https://github.com/marionettejs/backbone.marionette)
  - テストの書き方と Gruntfile.js の設定が `jasmine-jquery` を使った fixture の取り扱いの参考になった

<iframe src="http://rcm-fe.amazon-adsystem.com/e/cm?lt1=_blank&bc1=000000&IS2=1&bg1=FFFFFF&fc1=000000&lc1=0000FF&t=pleasesleep-22&o=9&p=8&l=as4&m=amazon&f=ifr&ref=ss_til&asins=477415489X" style="width:120px;height:240px;" scrolling="no" marginwidth="0" marginheight="0" frameborder="0"></iframe>

