{"title":"Sinon.js Code Reading","date":"2013-11-03T20:33:04+09:00","tags":["javascript"]}

### モジュールのロードまわり

- `lib/sinon.js` がモジュールのエンドポイント
  - sinon object の作成、環境に応じた初期化、ユーティリティメソッドの定義を行う
- spy や mock などの機能毎にファイルが分かれる
- `lib/sinon/*.js` に配置
  - `lib/sinon/spy.js` など

#### sinon.js

大きくは以下のように sinon object を作って返す。

<pre><code data-language="javascript">var sinon = (function() {
    function somePrivateFunction() {};

    var sinon = {
        foo: function foo() {}
    };

    return sinon;
}());</code></pre>

node の場合、ブラウザの場合、busterjs の場合で異なった初期化を行う。

node の環境かどうかの判定は `module.exports` の有無で行う。

<pre><code data-language="javascript">var isNode = typeof module !== "undefined" && module.exports;</code></pre>

node の場合は個別のモジュールを exports する。

<pre><code data-language="javascript">    if (isNode) {
        try {
            buster = { format: require("buster-format") };
        } catch (e) {}
        module.exports = sinon;
        module.exports.spy = require("./sinon/spy");
        module.exports.spyCall = require("./sinon/call");
        module.exports.stub = require("./sinon/stub");
        // ...
    }</code></pre>

ブラウザの場合は特になにもせず sinon オブジェクトを返す。ブラウザの場合はそれぞれのファイルを結合する前提なのでこれでよい。

#### lib/sinon/*.js

個別のモジュールのロード周りは要約するとこうなっている。これは spy.js の例。

<pre><code data-language="javascript">(function (sinon) {
    var commonJSModule = typeof module !== 'undefined' && module.exports;

    if (!sinon && commonJSModule) {
        sinon = require("../sinon");
    }

    if (!sinon) {
        return;
    }

    function spy(object, property) {
        // ...
    }

    if (commonJSModule) {
        module.exports = spy;
    } else {
        sinon.spy = spy;
    }
}(typeof sinon == "object" && sinon || null));
</code></pre>

まず sinon オブジェクト。ブラウザの場合は引数として渡されるもの、node の場合 (commonJSModule の場合) は `../sinon.js` ファイルを require する。そのあとこのファイルが担当するクラスの定義 (今回は spy) がつづき、それを export する。commonJSModule の場合は単純に `module.exports` する。先に見たようにここで export したものは sinon.js ファイル内で改めてそれぞれの名前空間に export される。ブラウザの場合は sinon オブジェクトにプロパティとして挿入する。

### spy のコンストラクタまわり

spy を使うともともとある関数を呼び出し回数などの各種記録処理でラップして、あとから記録結果を参照できるようにするものだ。これで意図した通りの引数で呼び出されているか。意図した回数呼び出されているか。意図した例外を投げているかなどをチェックできる。使い方には三通りある。

- コンストラクタの引数として関数を渡す場合
  - その関数はそのままに、呼び出し回数や引数を記録する処理でラップする。あとから記録した内容を参照できる。
- 引数なしでインスタンスを作成した場合
  - 上記と同じく各種記録を行うが、それだけの無名関数を返す。無名関数は呼び出し回数などを記録するが、戻り値として undefined を返すだけ。
  - あるメソッドのコールバック呼び出しをテストしたい、その結果はいらない、という時に使う
- オブジェクトとプロパティ名を渡す場合
  - そのプロパティの中身を spy の処理でラップしたものに置き換える


#### コンストラクタの引数として関数を渡す場合

その関数を単に spy.create に放り込む。

<pre><code data-language="javascript">        create: function create(func) {
            var name;

            if (typeof func != "function") {
                func = function () { };
            } else {
                name = sinon.functionName(func);
            }

            var proxy = createProxy(func);

            sinon.extend(proxy, spy);
            delete proxy.create;
            sinon.extend(proxy, func);

            proxy.reset();
            proxy.prototype = func.prototype;
            proxy.displayName = name || "spy";
            proxy.toString = sinon.functionToString;
            proxy._create = sinon.spy.create;
            proxy.id = "spy#" + uuid++;

            return proxy;
        },</code></pre>

- createProxy で proxy オブジェクトを作る
- proxy オブジェクトは spy オブジェクト、もともとの (引数で渡された) func オブジェクトと同じプロパティを持つ
  - sinon.exted で実現
- proxy を reset したり、一意な id をふったりして初期化

createProxy はどんなのか

- 最終的に proxy 関数を返す
  - spy.invoke した結果を返す関数
- つまりもとのメソッドが呼び出されると、その引数をそのまま `spy.invoke` に proxy している
- もとの関数の引数は function.length と eval で処理している
  - length にはこの関数が期待している引数の数が入る

よくわからないのは function.length が 1 以上の関数の proxy の仕方。最終的には invoke に arguments オブジェクトを渡せばいいので、proxy 関数には仮引数の宣言はいらないはず。それがわざわざ文字列を頑張って操作して関数宣言をつくり、それを eval するというトリッキーな方法をとってまで仮引数の宣言を作っている。一体なぜ。

話を戻して、spy.invoke はこんな感じだ。

<pre><code data-language="javascript">        invoke: function invoke(func, thisValue, args) {
            var matching = matchingFake(this.fakes, args);
            var exception, returnValue;

            incrementCallCount.call(this);
            push.call(this.thisValues, thisValue);
            push.call(this.args, args);
            push.call(this.callIds, callId++);

            try {
                if (matching) {
                    returnValue = matching.invoke(func, thisValue, args);
                } else {
                    returnValue = (this.func || func).apply(thisValue, args);
                }
            } catch (e) {
                push.call(this.returnValues, undefined);
                exception = e;
                throw e;
            } finally {
                push.call(this.exceptions, exception);
            }

            push.call(this.returnValues, returnValue);

            createCallProperties.call(this);

            return returnValue;
        },</code></pre>

おおまかには、その関数の引数、戻り値、呼び出し回数などの記録処理を、もとの関数の呼び出しにかぶせていることになる。このへんで保管している呼び出し回数などは spy.reset で初期化している。

createProxy で新しいプロキシ関数をつくって、そいつを extend して必要な機能をもたせている。プロキシ関数はひとつのインスタンスなようなもので、独立した環境をもっており、その this の中で各種記録を保管している。明示的に new したりせず、このように新しい環境と関数でラップするような書き方は、関数型的な考え方ということでいいのだろうか。

#### 引数なしでインスタンスを作成した場合

単に空っぽの無名関数を spy.create に渡しているだけ。あとは先ほどと同じだ。

#### オブジェクトとプロパティ名を渡す場合

まず対象のオブジェクトのプロパティを取り出し、それで spy.create する。こうしてプロキシをはさんだメソッドをもとのオブジェクトで使えるようにするのだが、それには sinon.wrapMethod を使う。

wrapMethod のしごとは主に 2 つ。対象プロパティを proxy 関数で置き換えること、もとの状態に戻す restore メソッドの提供だ。

メソッドの置換えは単純に `object[property] = method;` というふうに代入するだけ。当然その際にもとのメソッドは別の変数に退避してある。

restore メソッドはこのように実装されている。

<pre><code data-language="javascript">            method.restore = function () {
                // For prototype properties try to reset by delete first.
                // If this fails (ex: localStorage on mobile safari) then force a reset
                // via direct assignment.
                if (!owned) {
                    delete object[property];
                }
                if (object[property] === method) {
                    object[property] = wrappedMethod;
                }
            };</code></pre>

owned には、指定されたプロパティがオブジェクトのものかどうかが真偽値で入っている。`!owned` はそのオブジェクトの持ち物ではない、つまりプロトタイプに入っているプロパティということだ。その場合はまずプロパティの delete を試みる。なぜそうしているかはよくわかっていない。

次は wrappedMethod 変数に退避していあるもとのメソッドを単純な代入で復帰させている。`object[property] === method` (method は置き換えした proxy 関数) という条件式で置き換えてよいかをチェックしている。例えば spy が object[property] を置き換えたあとにユーザーがそれを上書きした場合などは何もしない。

wrapMethod は呼び出し方に不備があった場合にエラーを例外で通知する。引数の型チェックや、対象のメソッドがすでに wrap 済みかどうかのチェックなどを行う。他のメソッドではここまで細かいチェックはしていない。エラーが有った場合は即座に undefined を return するくらいだ。wrapMethod だけ特別厳重にバリデーションをしているのはなぜだろう。

### spy api

spy オブジェクトにはいくつものメソッド・プロパティが準備されている (spy api と呼ばれる)。多くは spy.callCount、spy.calledWith、spy.calledAfter など、spy した結果をあとで assert する際につかうものだ。前述したように spy オブジェクトは呼び出しに関わる情報、引数・this の値・返り値・例外など、はすべて保持している。これらのメソッドは基本的にその内容を読みだして真偽値を返すような役割を担っている。

spy.getCall というメソッドがある。これは spy の呼び出しのうち特定のものだけを取り出すためのものだ。3 回目の呼び出しは `spy.getCall(3)` で取得する。戻り値は `spyCall` というオブジェクトで、こちらも calledWith などの便利なメソッドが用意されている。spyCall は `sinon/call.js` で定義されている。

getCall の実装はこうなっている。

<pre><code data-language="javascript">        getCall: function getCall(i) {
            if (i < 0 || i >= this.callCount) {
                return null;
            }

            return sinon.spyCall(this, this.thisValues[i], this.args[i],
                                    this.returnValues[i], this.exceptions[i],
                                    this.callIds[i]);
        },</code></pre>

前述したように他のファイルに分かれている処理はすべて、node 環境・ブラウザ環境ともに、sinon オブジェクトにストアされている。よって `sinon.spyCall( ... )` というふうに sinon オブジェクトを経由して spyCall オブジェクトを作成している。

spy api の中でもう一つ特殊なメソッドは `withArgs` だ。これは特定の引数での呼び出しの場合のみ記録するという設定を spy オブジェクトに対して行うものだ。

インタフェースとしては次のように、事前に `spy.withArgs(ARG)` で登録。その後同じように `spy.withArgs(ARG)` とすると、指定した引数での呼び出しのみを記録している spy オブジェクトが返される。

<pre><code data-language="javascript">var object = { method: function () {} };
var spy = sinon.spy(object, "method");
spy.withArgs(42);
spy.withArgs(1);

object.method(42);
object.method(1);

assert(spy.withArgs(42).calledOnce);
assert(spy.withArgs(1).calledOnce);</code></pre>

withArgs の処理の流れはシンプルだ。

- まずすでに同様の設定がないかを調べる。合った場合はそのオブジェクトを返す。withArgs で作られた特定引数の場合のみ記録を行う spy オブジェクトは `fakes` プロパティに保存されている。
- 通常の spy オブジェクト同様に `spy.create()` メソッドでオブジェクトを作る。ただし `matchingArguments` プロパティに設定したい引数を持つことと、こうして作られたオブジェクトを `fakes` プロパティに保存する点が異なる。
- すでにこの spy インスタンスに関して何度か呼び出しが行われているかもしれない。そのため今までの呼び出し履歴 (`spy.args` プロパティ) を全て調べ、マッチする呼び出しが合った場合は自身 (fake オブジェクト) に記録する。

### stub, mock

力尽きたのでここは概要だけ。

spy は関数呼び出しの挙動を記録して、あとでその結果を assert するものだった。stub はそれに加えて関数の戻り値を変更することができる。テスト対象の関数が本番と同様の動きをするのではなく、テスト用のダミーの値を返すように事前に設定できる。

mock は事前にダミーの挙動を設定するだけでなく、その期待する結果も事前に設定できる。事前にその mock は何度呼び出されるべきかなどを指定しておき、動作終了後に `mock.verify()` を呼ぶと期待通りだったかをチェックしてくれる。spy や stub を使い、呼び出し回数をあとで assert しても同様のことができる。

### Fake timers

spy, stub, mock のように assertion のために挙動を記録するようなものではなく、テストのために単にダミーの挙動をするだけの Fake というものも Sinon にはある。特にテストを書く上で問題になりやすい timer や XHR といったオブジェクトの fake が提供されている。

ここでは Fake timer をみてみる。Fake timer は setTimeout や Date オブジェクトがからむ処理のテストで力を発揮する。次のように任意の期間の時間をすすめることができる。

<pre><code data-language="javascript">{
    setUp: function () {
        this.clock = sinon.useFakeTimers();
    },

    tearDown: function () {
        this.clock.restore();
    },

    "test should animate element over 500ms" : function(){
        var el = jQuery("<div></div>");
        el.appendTo(document.body);

        el.animate({ height: "200px", width: "200px" });
        this.clock.tick(510);

        assertEquals("200px", el.css("height"));
        assertEquals("200px", el.css("width"));
    }
}</code></pre>

fake timer が置き換えるオブジェクト、メソッドは以下の 5 つだ。(もちろん明示すれば特定のものだけを置き換えることもできる)

- Date
- setTimeout
- setIntervl
- clearTimeout
- clearInterval

これらを置き換えた関数と `tick()` 関数で、処理系の時間軸とは独立して、任意の時間を進めたり、それに応じて timeout / interval で設定してあるコールバックを呼び出したりする。

fake timer のこの挙動は次のような戦略で実現されている。

- まず fake tiemr (sinon.clock オブジェクト) は `timeouts` というプロパティを持っている。setTimeout / setInteval で指定されたコールバック関数はこのプロパティに保管される。このデータ構造が fake timer の肝だ。
- timeouts は sinon.clock 内で管理するユニークな ID をキーにして、以下の値を持つハッシュだ。この値のオブジェクトは timer と名付けられている
  - func: コールバック関数
  - callAt: 次回コールバックを実行する時刻。unix timestamp。
  - interval: インターバル。単位は ms。インターバルがない (setTimeout) の場合は undefined
  - id: コールバック関数の ID。timeouts のキーと同じ値。
  - invokeArgs: コールバック関数呼び出し時に渡される引数。
- (fake の) setTimeout / setInteval は渡されたコールバック関数や待機時間を timeouts につめこむ
- `tick()` が指定された期間時間をすすめる
  - timeouts 各要素の callAt をチェックし、コールバック関数を呼び出す
  - interval が指定されている場合は callAt + interval を callAt に設定する
- clearTimeout / clearInterval / restore はそれぞれ適切に timeouts プロパティをクリアする

少しコードを覗いてみよう。tick メソッドは以下だ。

<pre><code data-language="javascript">        tick: function tick(ms) {
            ms = typeof ms == "number" ? ms : parseTime(ms);
            var tickFrom = this.now, tickTo = this.now + ms, previous = this.now;
            var timer = this.firstTimerInRange(tickFrom, tickTo);

            var firstException;
            while (timer && tickFrom <= tickTo) {
                if (this.timeouts[timer.id]) {
                    tickFrom = this.now = timer.callAt;
                    try {
                      this.callTimer(timer);
                    } catch (e) {
                      firstException = firstException || e;
                    }
                }

                timer = this.firstTimerInRange(previous, tickTo);
                previous = tickFrom;
            }

            this.now = tickTo;

            if (firstException) {
              throw firstException;
            }

            return this.now;
        },</code></pre>

現在時刻からある時間を進めた場合に一番最初に呼び出されるべきコールバック関数を調べて timer オブジェクトを返すのが  `firstTimerInRange` というメソッドだ。tick は while ループの中で firstTimerInRange を呼び出し、順にコールバックを実行していき、最後に sinon.clock.now に進められた時刻を入れる。コールバックを実行する役割を持つのは `callTimer` メソッドだ。

`firstTimerInRange` は次のような実装になっている。

<pre><code data-language="javascript">        firstTimerInRange: function (from, to) {
            var timer, smallest = null, originalTimer;

            for (var id in this.timeouts) {
                if (this.timeouts.hasOwnProperty(id)) {
                    if (this.timeouts[id].callAt < from || this.timeouts[id].callAt > to) {
                        continue;
                    }

                    if (smallest === null || this.timeouts[id].callAt < smallest) {
                        originalTimer = this.timeouts[id];
                        smallest = this.timeouts[id].callAt;

                        timer = {
                            func: this.timeouts[id].func,
                            callAt: this.timeouts[id].callAt,
                            interval: this.timeouts[id].interval,
                            id: this.timeouts[id].id,
                            invokeArgs: this.timeouts[id].invokeArgs
                        };
                    }
                }
            }

            return timer || null;
        },</code></pre>

呼ばれるたびに timeouts を全探索し、from, to の範囲内にあってもっとも callAt が若い timer オブジェクトを返している。正直毎回 timer を線形探索しているのでかっこ悪いが、そんなに大量のコールバックを設定することもなさそうなので問題無いだろう。

次は `callTimer` の実装。

<pre><code data-language="javascript">        callTimer: function (timer) {
            if (typeof timer.interval == "number") {
                this.timeouts[timer.id].callAt += timer.interval;
            } else {
                delete this.timeouts[timer.id];
            }

            try {
                if (typeof timer.func == "function") {
                    timer.func.apply(null, timer.invokeArgs);
                } else {
                    eval(timer.func);
                }
            } catch (e) {
              var exception = e;
            }

            if (!this.timeouts[timer.id]) {
                if (exception) {
                  throw exception;
                }
                return;
            }

            if (exception) {
              throw exception;
            }
        },</code></pre>

内容としては interval の場合は次回実行時刻をセット・そうでない場合は timeouts からオブジェクトを削除し、あとはコールバックを呼び出しているだけだ。

また fake timer は Date オブジェクトも上書きしている。`new Date()` したときに処理系の現在時刻ではなく、テスト時に指定した時刻 (sinon.clock.now) で返すという違いだけで、あとはもともとと同じだ。
