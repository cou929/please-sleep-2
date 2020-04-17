{"title":"twitter bootstrap のテスト","date":"2012-12-11T23:42:14+09:00","tags":["javascript"]}

エントリポイントは Makefile なのでここから掘り下げる.

### Makefile

    test:
    	jshint js/*.js --config js/.jshintrc
    	jshint js/tests/unit/*.js --config js/.jshintrc
    	node js/tests/server.js &
    	phantomjs js/tests/phantom.js "http://localhost:3000/js/tests"
    	kill -9 `cat js/tests/pid.txt`
    	rm js/tests/pid.txt

- js とテストスクリプトそのものを jshint で lint にかける
- node でサーバを立ち上げる
  - `js/tests/server.js`
  - バックグラウンドプロセスとして実行. プロセス番号をファイルに落としておく
- phantomjs でテスト実行
  - エントリポイントは `js/tests/phantom.js`
  - 各々のテストは qunit で書かれている
  - 上記の node のサーバが配信する `js/tests/index.html` にアクセスを飛ばしてテストと実行する.
- サーバプロセスの kill と pid ファイルの削除

### lint

`js/.jshintrc` に設定が書いてある中身はこちら

    {
        "validthis": true,
        "laxcomma" : true,
        "laxbreak" : true,
        "browser"  : true,
        "debug"    : true,
        "boss"     : true,
        "expr"     : true,
        "asi"      : true
    }

それぞれのオプションについて

- validthis
  - strict モードかつ関数のコンストラクタ以外で this を使っている場合でも警告を出さない. 通常の strict mode の動作だと, 例えばコンストラクタなのに new をつけ忘れて実行して global を汚染する可能性があるので, コンストラクタじゃない関数で this を使うと警告してくれる.
- laxcomma
  - 前置コンマでも警告を出さないようにする
- laxbreak
  - 安全でない改行でも警告を出さない. どんなケースが該当するのかは種類が多いようなのでわからないが, コーディングスタイルによって改行位置は変わるのでいろいろと議論はあるようだ.
- browser
  - ブラウザ環境で動作するスクリプトであると明示する. つまりグローバル変数として document などを使っても警告されない.
- debug
  - `debugger` 文が出てきても警告しない
- boss
  - for 文の比較部など, 比較式が期待される部分で代入を行なっていても警告をださない. 確かにそういうループの書き方をすることもある.
- expr
  - 代入や関数呼び出しが期待されている部分に式があっても警告しない.
- asi
  - セミコロンが行末になくても警告を出さない. そういえばセミコロンをかたくなに付けない人たちの主張をちゃんと読んだことないし今度調べてみよう.

### server.js

ポート 3000 番で http サーバをたてて, プロジェクトホーム以下のファイルを静的に返すようにしている. またプロセス番号を `pid.txt` に落とす.

内容もこれだけ.

    /*
     * Simple connect server for phantom.js
     * Adapted from Modernizr
     */
    
    var connect = require('connect')
      , http = require('http')
      , fs   = require('fs')
      , app = connect()
          .use(connect.static(__dirname + '/../../'));
    
    http.createServer(app).listen(3000);
    
    fs.writeFileSync(__dirname + '/pid.txt', process.pid, 'utf-8')

### phantom.js

これがテストランナーにあたる.

- 引数で渡された url へアクセスする
- `waitFor()` はテストの実行完了を待つ関数. タイムアウトを設定できる.
- waitFor の第一引数の無名関数がテスト完了をチェックしている. 単純に dom 中の innerText を見てテスト完了のメッセージが出ているかをチェックする.
- waitFor の第二匹数の無名関数はテスト結果に応じて終了コードを設定してプロセスを抜ける. テスト結果は例のごとく dom から取得. テストが落ちていたら 1 で exit する.

テストがコケた場合終了ステータスが 1 になるので, make がエラーメッセージを出してくれる & make 自体の終了コードも非 0 になる.

    make: *** [test] Error 1

また, こんなかんじで `page.onConsoleMessage` でページのコンテキストで出された console.log のメッセージを phantomjs のコンテキストに持ってきている.

    page.onConsoleMessage = function(msg) {
      console.log(msg)
    };

よばれる (テストを実行する) index.html 側ではテストの進捗を console.log でも出すようにしている (後述). 通常 qunit は dom 上にテストの進捗や結果を出すが, 今回は無理やりコマンドラインから実行している. そのためコマンドラインからテストの進捗・結果が見えないので, このように console.log をうまくページ -> phantomjs のコンテキストへルートしてあげている. コメントによるとこのへんの仕組みは modernizr のテストを参考にしているそう.

[Modernizr/test at master · Modernizr/Modernizr · GitHub](https://github.com/Modernizr/Modernizr/tree/master/test)

### js/tests/index.html

ここから先はふつうの qunit のテストと同じ. jquery, qunit.js などなど必要なパッケージを読み込んでから, `js/tests/unit/` 以下のテストを読み込んで走らせている.

`js/tests/unit/bootstrap-phantom.js` で `QUnit.begin` (いわゆる setup, teardown) などにログを残す関数を仕込んでいる. ここでテストの進捗・結果を console.log して, それをさきほどの phantomjs 内の onConsoleMessage が拾って表示させている.

    // Logging setup for phantom integration
    // adapted from Modernizr
    
    QUnit.begin = function () {
      console.log("Starting test suite")
      console.log("================================================\n")
    }
    
    QUnit.moduleDone = function (opts) {
      if (opts.failed === 0) {
        console.log("\u2714 All tests passed in '" + opts.name + "' module")
      } else {
        console.log("\u2716 " + opts.failed + " tests failed in '" + opts.name + "' module")
      }
    }
    
    QUnit.done = function (opts) {
      console.log("\n================================================")
      console.log("Tests completed in " + opts.runtime + " milliseconds")
      console.log(opts.passed + " tests of " + opts.total + " passed, " + opts.failed + " failed.")
    }

### あと片付け

node のサーバプロセスを kill し, pid ファイルを削除します

### テストを実行した結果

実際実行してみるとこんなかんじ.

    $ make test
    jshint js/*.js --config js/.jshintrc
    jshint js/tests/unit/*.js --config js/.jshintrc
    node js/tests/server.js &
    phantomjs js/tests/phantom.js "http://localhost:3000/js/tests"
    Starting test suite
    ================================================
    
    ✔ All tests passed in 'bootstrap-transition' module
    ✔ All tests passed in 'bootstrap-alerts' module
    ✔ All tests passed in 'bootstrap-buttons' module
    ✔ All tests passed in 'bootstrap-carousel' module
    ✔ All tests passed in 'bootstrap-collapse' module
    ✔ All tests passed in 'bootstrap-dropdowns' module
    ✔ All tests passed in 'bootstrap-modal' module
    ✔ All tests passed in 'bootstrap-scrollspy' module
    ✔ All tests passed in 'bootstrap-tabs' module
    ✔ All tests passed in 'bootstrap-tooltip' module
    ✔ All tests passed in 'bootstrap-popover' module
    ✔ All tests passed in 'bootstrap-typeahead' module
    
    ================================================
    Tests completed in 2026 milliseconds
    142 tests of 142 passed, 0 failed.
    kill -9 `cat js/tests/pid.txt`
    rm js/tests/pid.txt
