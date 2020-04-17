{"title":"Jasmine の HtmlReporter を読む","date":"2013-06-30T21:18:03+09:00","tags":["memo"]}

![](/images/jasmine_logo.png)
[pivotal/jasmine · GitHub](https://github.com/pivotal/jasmine)

jasmine をブラウザ上で実行する際につかう jasmine-html.js を読む。jasmine のコアロジックは jasmine.js で定義されていて、ブラウザ上で実行され html を出力する、MVC でいう View に当たる部分が jasmine-html.js で実装されている。

jasmine-html.js には 2 つの名前空間が登場する。ひとつは `HtmlReporter`, もうひとつは `HtmlReporterHelpers` だ。両方 jasmine 名前空間の下につけられている。

`HtmlReporterHelpers` は単なるオブジェクトで、ユーティリティ用のメソッドが詰め込まれている。`HtmlReporter` は `jasmine.Reporter` と同じインタフェースを実装したクラスだ。jasmine には `jasmine.Reporter` というクラスが提供されており、これと同じインタフェースを実装すれば独自の出力ができるようになっている。

まずは SpecRunner.html で htmlReporter がどのようにつかわれているのか。

<pre><code data-language="javascript">      var htmlReporter = new jasmine.HtmlReporter();

      jasmineEnv.addReporter(htmlReporter);

      jasmineEnv.specFilter = function(spec) {
        return htmlReporter.specFilter(spec);
      };</code></pre>

こんなふうに jasmineEnv に対して `addReporter` と `specFilter` している。では jasmine.js を読んでいって、この 2 つのメソッドの役割をみていこう。

### jasmine.html

まず addReporter は jasmine のステータスアップデートを受け取ってレポートするレポーターを登録するものだ。

<pre><code data-language="javascript">/**
 * Register a reporter to receive status updates from Jasmine.
 * @param {jasmine.Reporter} reporter An object which will receive status updates.
 */
jasmine.Env.prototype.addReporter = function(reporter) {
  this.reporter.addReporter(reporter);
};</code></pre>

この Reporter というクラスにはどういうふるまいが求められているのか。Reporter のベースクラスにインタフェースが定義されている。

<pre><code data-language="javascript">/** No-op base class for Jasmine reporters.
 *
 * @constructor
 */
jasmine.Reporter = function() {
};

//noinspection JSUnusedLocalSymbols
jasmine.Reporter.prototype.reportRunnerStarting = function(runner) {
};

//noinspection JSUnusedLocalSymbols
jasmine.Reporter.prototype.reportRunnerResults = function(runner) {
};

//noinspection JSUnusedLocalSymbols
jasmine.Reporter.prototype.reportSuiteResults = function(suite) {
};

//noinspection JSUnusedLocalSymbols
jasmine.Reporter.prototype.reportSpecStarting = function(spec) {
};

//noinspection JSUnusedLocalSymbols
jasmine.Reporter.prototype.reportSpecResults = function(spec) {
};

//noinspection JSUnusedLocalSymbols
jasmine.Reporter.prototype.log = function(str) {
};</code></pre>

この 6 つのメソッドが必要なようだ。それぞれの役割は `jasmine.Runner` の部分を読んでいくとわかる。

- reportRunnerStarting
  - execute 直後、start 前によばれる。引数は Runner インスタンス
- reportRunnerResults
  - すべての完了後によばれる。引数は Runner インスタンス
- reportSuiteResults
  - Suite の完了時によばれる。引数は Suite インスタンス
- reportSpecStarting
  - spec 開始時によばれる。 specFilter でフィルタされた場合はよばれない。引数は Spec インスタンス
- reportSpecResults
  - spec 完了時によばれる。引数は Spec インスタンス
- log
  - ほとんど使われていない? 引数としてログ文字列をうけとりよしなにする

ちなみに Jasmine のコンストラクタでは MultiReporter というのを reporter として登録している。

<pre><code data-language="jasmine">  this.reporter = new jasmine.MultiReporter();</code></pre>

`specFilter` は `jasmineEnv` に登録される関数で、各 Spec 実行前にこのフィルタに通し、戻り値が偽ならば実行されない。引数として Spec インスタンスを受け取る。

<pre><code data-language="javascript">jasmine.Spec.prototype.execute = function(onComplete) {
  var spec = this;
  if (!spec.env.specFilter(spec)) {
    spec.results_.skipped = true;
    spec.finish(onComplete);
    return;
  }

  // spec 実行
}</code></pre>

### jasmine-html.js

`jasmine-html.js` では `ReporterView` と `SpecView` の 2 つのクラスが登場する。それぞれテスト全体の状態と spec ごとの状態を管理する。各 spec を実行するごとに jasmine 本体から渡される Spec インスタンスをうけとり、SpecView が結果を取り出し、それぞれの Spec の結果を ReporterView インスタンスに加算していく。これらの結果をそれぞれの `*Results` メソッドが dom に反映させていくという仕組みだ。

では上記のインタフェースを `jasmine-html.js` はどう実装しているかをみていく。

`reportRunnerStarting` では当然ながら出力の前準備をしている。出力先の dom をつくったり、結果を格納する ReporterView クラスを準備したり、パラメーター初期設定をしたりだ。引数である `Runner` のインスタンスからは spec の一覧を取得している。spec 数などの情報を得るためだけで、spec の編集はしていない。

`reportRunnerResults` ではすべての実行が終わったあとに、実行数や成功・失敗数、失敗した場合はその詳細表示などをおこなう。

`reportSuiteResults` では、各 suite に対応する view クラスのインスタンスをとりだし、リフレッシュする。リフレッシュ処理では発生したアラートに応じた dom 表示など、Suite 単位の表示を扱う。

`reportSpecStarting` では、Suite と Spec の description を表示させるのみ。

`reportSpecResults` では、引数の Spec クラスのインスタンスから SpecView をつくり、テスト結果を取り出し、ReporterView の内部のカウンタをインクリメントする。こうして Spec ごとに結果を記録していき、上位の `*Results` で dom に反映させる。

`specFilter` は実行する spec を指定すると、それだけをフィルタするようになっている。
