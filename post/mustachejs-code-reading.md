{"title":"mustache.js code reading","date":"2012-02-25T19:07:29+09:00","tags":["javascript"]}

[{{ mustache }}](http://mustache.github.com/)

シンプルなテンプレートエンジンが欲しかったので使ってみる. その前に読んでみる. 500 行くらい.

### intro

node などだったら exports して, クライアントサイドだったら (module っていうグローバルオブジェクトがなかったら) グローバルに Mustache というオブジェクトを作る.

    var Mustache = (typeof module !== "undefined" && module.exports) || {};
    ...
    (function (exports) {
        ...
    })(Mustache);

このへんのユーティリティは無かったら自前で定義

    var _isArray = Array.isArray;
    var _forEach = Array.prototype.forEach;
    var _trim = String.prototype.trim;

IE (version 不明) は `\s` に `\xA0` (non-breaking spaces) が含まれないらしい.

    if (isWhitespace("\xA0")) {
    ...
    } else {
      // IE doesn't match non-breaking spaces with \s, thanks jQuery.
      trimLeft = /^[\s\xA0]+/;
      trimRight = /[\s\xA0]+$/;
    }

実体参照をエスケープしているけど, 正規表現の意味がわからない.

      var escapeMap = {
        "&": "&amp;",
        "<": "&lt;",
        ">": "&gt;",
        '"': '&quot;',
        "'": '&#39;'
      };

      function escapeHTML(string) {
        return String(string).replace(/&(?!\w+;)|[<>"']/g, function (s) {
          return escapeMap[s] || s;
        });
      }

さて, export している関数はこれらなので, ここから降りていくことにする.

      exports.parse = parse;
      exports.compile = compile;
      exports.render = render;
      exports.clearCache = clearCache;

- parse
  - template をパース
- compile
  - たぶんテンプレートをキャッシュさせたい時など
- render
  - template と data を組み合わせて html を返す. その場でコンパイルしている
- clearCache
  - compile されたテンプレートのキャッシュクリア

### parse()

template とオプションをとる. オプションはデバッグのためのものと, タグ (デフォルトは {{}}), タグの中にスペースを許すかというフラグ.

    var code = [
      'var buffer = "";', // output buffer
      "\nvar line = 1;", // keep track of source line number
      "\ntry {",
      '\nbuffer += "'
    ];

template をパースして, その内容に応じてこの code 配列の中に js のプログラムとして文字列を push していく. あとで new Function して関数化して使うことになる.

ひたすら関数定義が並び, L339 からメイン処理が始まる.

ループでテンプレートを一文字ずつみていく.

    for (var i = 0, len = template.length; i < len; ++i) {

タグの長さ分の文字数を取り出して, ``open tag`` (``{{}}`` だと ``{{``) かどうかをチェック. js だと slice しないといけないから遅そう. es5 からだと template[i] でいけたと思うけど.

      if (template.slice(i, i + openTag.length) === openTag) {

タグを一旦退避させておく. mustache は動的にタグの種類を変えられるので. この機能, 別にコンパイル時に指定とかでもいい気がするけど, erb とか他のテンプレートエンジンのテンプレートを読み込んで使うときには便利なんだと思う.

        nextOpenTag = openTag;
        nextCloseTag = closeTag;

タグの種類に応じてコールバックをディスパッチするスイッチ文に入る. コールバックは上でいろいろ定義されてた関数群. 必要に応じてフラグを上げ下げしたり, カーソルを進めたりしている.

        switch (c) {
        case "!": // comment
          i++;
          callback = null;
          break;
        case "=": // change open/close tags, e.g. {{=<% %>=}}
          i++;
          closeTag = "=" + closeTag;
          callback = setTags;
          break;
        case ">": // include partial

閉じタグを探して, 見つからなかったら例外

        var end = template.indexOf(closeTag, i);

        if (end === -1) {
          throw debug(new Error('Tag "' + openTag + '" was not closed properly'), template, line, options.file);
        }

テンプレートの中身を取り出して, 対応する関数を呼ぶ

        var source = template.substring(i, end);

        if (callback) {
          callback(source);
        }

このあと, もしタグの中に改行があったら行番号をインクリメントしたり, タグの種類に変更があったらグローバル変数にセットしたりなど.

タグだった場合の処理はここまで. このとおり結構愚直にやってる. タグのネストには対応していない (そんなケースあるのかわからないけど)

タグじゃなかった場合は, 状態をメンテナンスして, code 配列にその文字をプッシュする. `\` の場合は `\\` にしたり (あとで評価するため), 改行があったら行番号を増やしたり, 多少状態をメンテナンスする. また空白だった場合は spaces という配列にも位置をプッシュしている. code のスペースの位置をこれが記憶しているようだ.

        c = template.substr(i, 1);

        switch (c) {
        case '"':
        case "\\":
          nonSpace = true;
          code.push("\\" + c);
          break;
        case "\r":
          // Ignore carriage returns.
          break;
        case "\n":
          spaces.push(code.length);
          code.push("\\n");
          stripSpace(); // Check for whitespace on the current line.
          line++;
          break;
        default:
          if (isWhitespace(c)) {
            spaces.push(code.length);
          } else {
            nonSpace = true;
          }

          code.push(c);

template を一文字ずつ見ているループはここまで.

その後はタグが全部閉じてるかバリデーションして, code に定型文を push. code を `join("")` で文字列にして返している. その際 `buffer += ""` をエスケープして, 無駄な空白行の扱いを省いているようだ.

### compile()

テンプレートとオプションを受け取って, コンパイル済みの関数をかえす. 戻された関数にデータを渡すと html が返って来る. 中でメモ化みたいなキャッシュの仕組みがあって, ヒットしなかったら _compile() というメインの関数を呼ぶ.

    function compile(template, options) {
      options = options || {};
  
      // Use a pre-compiled version from the cache if we have one.
      if (options.cache !== false) {
        if (!_cache[template]) {
          _cache[template] = _compile(template, options);
        }
  
        return _cache[template];
      }
  
      return _compile(template, options);
    }

`_compile()` も短くて, parse して得られた body を使って関数を作って返している.

    function _compile(template, options) {
      var args = "view,partials,stack,lookup,escapeHTML,renderSection,render";
      var body = parse(template, options);
      var fn = new Function(args, body);
  
      // This anonymous function wraps the generated function so we can do
      // argument coercion, setup some variables, and handle any errors
      // encountered while executing it.
      return function (view, partials) {
        partials = partials || {};
  
        var stack = [view]; // context stack
  
        try {
          return fn(view, partials, stack, lookup, escapeHTML, renderSection, render);
        } catch (e) {
          throw debug(e.error, template, e.line, options.file);
        }
      };
    }

処理の大枠はこんな感じで, あとは動かしてみて body の中を覗いてみたほうが早そう. それを踏まえて個別の関数へ潜っていく.

### body の中身

テスト用の html とスクリプト

test.html

    <html>
      <head>
        <script type="text/javascript" src="mustache.js"></script>
      </head>
      <body>
        <textarea id="template"></textarea>
        <textarea id="json"></textarea>
        <button value="submit" id="submit">submit</button>
        <div id="dst">
        </div>
        <script type="text/javascript" src="test.js"></script>
      </body>
    </html>

test.js

    (function() {
         document.querySelector('#submit')
             .addEventListener('click', function() {
                                   var template = document.querySelector('#template').value;
                                   var view = JSON.parse(document.querySelector('#json').value);
                                   var renderer = Mustache.compile(template, {
                                                                       debug: true
                                                                   });
                                   var output = renderer(view);
                                   document.querySelector('#dst').innerHTML = output;
                               });
     })();

template が

    This is a {{thing}}.

view が

    {"thing": "pen"}

だと, body はこんなん.

	var buffer = "";
	var line = 1;
	try {
	buffer += "This is a ";
	line = 1;
	buffer += escapeHTML(lookup("thing",stack,""));
	buffer += ".";
	return buffer;
	} catch (e) { throw {error: e, line: line}; }

`escapeHTML()` は実体参照を replace しているくらいだったので, `lookup()` を見てみる.

### lookup()

      /**
       * Looks up the value of the given `name` in the given context `stack`.
       */
      function lookup(name, stack, defaultValue) {

とのことなので, 名前と view を受け取って適切な値を返すもの.

stack はそのテンプレートに与えられる view を配列にして渡したもののようだ. `_compile()` の中のクロージャから参照されている.

    var stack = [view]; // context stack

やっていることは上述のとおり, オブジェクトの配列と名前が渡されるので, 名前に相当する値が無いか調べて返す. 値が関数だったらコールして返り値を返すし, 見つからなかったらデフォルト値を返す. ただし名前として `foo.bar.baz` という記法が使える点を考慮しないといけない.

対応する値を探している処理がここ

        var names = name.split(".");
        var lastIndex = names.length - 1;
        var target = names[lastIndex];

        var value, context, i = stack.length, j, localStack;
        while (i) {
          localStack = stack.slice(0);
          context = stack[--i];
    
          j = 0;
          while (j < lastIndex) {
            context = context[names[j++]];
    
            if (context == null) {
              break;
            }
    
            localStack.push(context);
          }
    
          if (context && target in context) {
            value = context[target];
            break;
          }
        }

名前を `.` で split して上から見ていく. 見つかったらその中をスタックの末尾に push する. 無駄な走査をしているので, ここは深さ優先でやったほうが良いのでは?

localStack は値が関数だった場合に context として渡すためにとっておいている. arr.slice(0) で配列をコピーしている. これは覚えておこう.

      localStack = stack.slice(0);
      ...

    // If the value is a function, call it in the current context.
    if (typeof value === "function") {
      value = value.call(localStack[localStack.length - 1]);
    }

### テスト

なんとなくわかったので, あとはテストを走らせてみよう (一番最初にやるべきだったかもしれないけど)

やり方は TESTING.md に詳しく書いてあった. rspec だそうだ

いろいろ必要物を準備して rake するだけ

    Koseis-MBA:kosei% rake
    /System/Library/Frameworks/Ruby.framework/Versions/1.8/usr/bin/ruby -S rspec spec/mustache_spec.rb
    Skipping tests in V8 (node not found)
    Skipping tests in SpiderMonkey (js not found)
    Skipping tests in Rhino (JAR org.mozilla.javascript.tools.shell.Main was not found)
    Testing in JavaScriptCore .............................................. Done!
    Verifying that we ran the tests in at least one engine ... OK
    .
    
    Finished in 1.68 seconds
    47 examples, 0 failures

rspec どころか ruby もまともに書いたことないけど, ruby 周りは落ち着いて読めば読めると信じている.

よくわからないけど `spec/*_spec.rb` を全部実行してくれるのだろう.

    desc "Run all specs"
    task :spec do
      require 'rspec/core/rake_task'
      RSpec::Core::RakeTask.new(:spec) do |t|
        #t.spec_opts = ['--options', "\"#{File.dirname(__FILE__)}/spec/spec.opts\""]
        t.pattern = 'spec/*_spec.rb'
      end
    end

`spec/mustache_spec.rb` がテストのエントリポイントで, `spec/_files/` 以下にテストパターンがたくさん入っている. `foo.txt`, `foo.mustache`, `foo.js` のセットが基本で, `foo.mustache` がテンプレート, `foo.js` が view, `foo.txt` が expected.

rspec のこういうインタフェース面白い

    it "should return the same result when invoked multiple times" do
      js = <<-JS
        #{@boilerplate}
        Mustache.render("x")
        print(Mustache.render("x"));
      JS

      run_js(@runner, js).should == "x\n"
    end

あと .travis.yml というファイルもあって, ちゃんと CI も回してて偉い

### おわり

とりあえずここまで.

思ったより愚直だった. あと view の lookup とか, parse とか高速化の余地ありありだと思う (それが必要かは別として). ライブラリとしてのインタフェースはすごくいいしシンプルで流行るのは理解できた.

テスト周りのことは ruby で勉強したい.
