{"title":"PhantomJS でネットワークのデバッグと SSL handshake failed","date":"2014-11-01T16:08:47+09:00","tags":["javascript"]}

PhantomJS で itunes の url に接続しようとするとエラーになり、デバッグしたいので方法を調べた話。

まず、SSL の通信なので、別のプロセスから tcpdump などで通信をキャプチャすることは難しい。

方法の一つとして、`--debug=yes` というコマンドラインオプションを渡すと、詳細なログをだしてくれるようだ。([ドキュメント](http://phantomjs.org/api/command-line.html) には載っていないオプションだった)

<pre><code>$ phantomjs --debug=yes crawl.js
2014-11-01T15:53:28 [DEBUG] CookieJar - Created but will not store cookies (use option '--cookies-file=<filename>' to enable persisten cookie storage)
2014-11-01T15:53:28 [DEBUG] Phantom - execute: Configuration
2014-11-01T15:53:28 [DEBUG]      0 objectName : ""
...
</code></pre>

あるいは面倒だけど、`onResourceRequested` などのイベントを片っ端から Listen して自分でログに落とすこともできる。

こんな感じ:

<pre><code data-language="javascript">var page = new WebPage(),
    system = require('system');

page.onResourceRequested = function (requestData, networkRequest) {
    system.stderr.writeLine('[onResourceRequested]');
    system.stderr.writeLine('Request (#' + requestData.id + '): ' + JSON.stringify(requestData));
};

page.onResourceReceived = function(response) {
    system.stderr.writeLine('[onResourceReceived]');
    system.stderr.writeLine('Response (#' + response.id + ', stage "' + response.stage + '"): ' + JSON.stringify(response));
};

page.onLoadStarted = function() {
    var currentUrl = page.evaluate(function() {
        return window.location.href;
    });
    system.stderr.writeLine('Current page ' + currentUrl + ' will gone...');
    system.stderr.writeLine('Now loading a new page...');
};

page.onLoadFinished = function(status) {
    system.stderr.writeLine('[onLoadFinished]');
    system.stderr.writeLine('Status: ' + status);
};

page.onNavigationRequested = function(url, type, willNavigate, main) {
    system.stderr.writeLine('[onNavigationRequested]');
    system.stderr.writeLine('Trying to navigate to: ' + url);
    system.stderr.writeLine('Caused by: ' + type);
    system.stderr.writeLine('Will actually navigate: ' + willNavigate);
    system.stderr.writeLine('Sent from the page\'s main frame: ' + main);
};

page.onResourceTimeout = function(request) {
    system.stderr.writeLine('[onResourceTimeout]');
    console.log('Response (#' + request.id + '): ' + JSON.stringify(request));
};

page.onResourceError = function(resourceError) {
    system.stderr.writeLine('[onResourceError]');
    system.stderr.writeLine('Unable to load resource (#' + resourceError.id + 'URL:' + resourceError.url + ')');
    system.stderr.writeLine('Error code: ' + resourceError.errorCode + '. Description: ' + resourceError.errorString);
};

page.onError = function(msg, trace) {
    var msgStack = ['ERROR: ' + msg];
    if (trace && trace.length) {
        msgStack.push('TRACE:');
        trace.forEach(function(t) {
            msgStack.push(' -> ' + t.file + ': ' + t.line + (t.function ? ' (in function "' + t.function +'")' : ''));
        });
    }

    system.stderr.writeLine('[onError]');
    system.stderr.writeLine(msgStack.join('\n'));
};

page.open('https://itunes.apple.com/en/app/instagram/id389801252?mt=8', function (status) {
    console.log('Finished');
});
</code></pre>

[Web Page Module | PhantomJS](http://phantomjs.org/api/webpage/)

ちゃんと調べていないが、CasperJS や Nightmarejs で同等のことをやろうとすると、後者じゃないとだめかもしれない。

### SSL handshake failed

ちなみに itunes の url に接続できなかったのは SSL Protocol が原因だった。上記のスクリプトを実行するとこんなメッセージが出る。

<pre><code>2014-11-01T15:59:40 [DEBUG] Network - Resource request error: 6 ( "SSL handshake failed" ) URL: "https://itunes.apple.com/en/app/instagram/id389801252?mt=8"

または

[onResourceError]
Unable to load resource (#1URL:https://itunes.apple.com/en/app/instagram/id389801252?mt=8)
Error code: 6. Description: SSL handshake failed
</code></pre>

curl でアクセスしてみると TLS 1.2 を使っているのがわかる

<pre><code>$ curl -v 'https://itunes.apple.com/en/app/instagram/id389801252?mt=8'
...
* TLS 1.2 connection using TLS_RSA_WITH_AES_256_CBC_SHA
* Server certificate: itunes.apple.com
...
</code></pre>

で、PhantomJS のデフォルトのプロトコルは SSLv3 らしい。

[Command Line Interface | PhantomJS](http://phantomjs.org/api/command-line.html)

`--ssl-protocol=tlsv1` (または `any`) とオプション指定してあげると解決。

<pre><code>$ phantomjs --ssl-protocol=tlsv1 crawl.js
</code></pre>

これも、CasperJS や Nightmarejs でも同様だと思う。

もちろんブラウザでも、curl でも問題ないのに PhantomJS でだけ接続できないし、エラーメッセージもほぼ無いしで結構はまった。

ちなみに POODLE 問題をうけて、[1.9.8 以降はデフォルトが TLSv1 に変更になった](https://github.com/ariya/phantomjs/pull/12663) ようだ。

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00KYMCWPU/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="http://ecx.images-amazon.com/images/I/51qTXH7DyJL._SL160_.jpg" alt="PhantomJS Cookbook" style="border: none;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00KYMCWPU/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">PhantomJS Cookbook</a><div class="amazlet-powered-date" style="font-size:80%;margin-top:5px;line-height:120%">posted with <a href="http://www.amazlet.com/" title="amazlet" target="_blank">amazlet</a> at 14.11.01</div></div><div class="amazlet-detail">Packt Publishing (2014-06-12)<br /></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00KYMCWPU/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
