{"title":"XDomainRequest の通信が中断する","date":"2015-02-01T22:56:49+09:00","tags":["javascript"]}

再現条件がわからないが、XDomainRequest を使った GET リクエストが必ず abort してしまうことがあった。このようなコードで発生した。

<pre><code data-language=“javascript">var url = '...';  // url to send request
var callback = function() {
    // callback function of http request
}

var xdr = new XDomainRequest();

xdr.onload = function() {
    var result;
    try {
        // do something with result
    } catch (error) {
        callback(error);
        return;
    }
   
    callback(null, result);
};
xdr.onerror = function() {
    callback(new Error('Network error'));
};
xdr.open('GET', url);
xdr.withCredentials = true;
xdr.send();</code></pre>

こちらの StackOverflow のスレッドによると、`XDomainRequest` の各イベントハンドラに関数を登録することで回避できるらしい。

[ajax - XDomainRequest aborts POST on IE 9 - Stack Overflow](http://stackoverflow.com/questions/15786966/xdomainrequest-aborts-post-on-ie-9)

よって以下のように `ontimeout` と `onprogress` にもハンドラを登録すると解消した。

<pre><code data-language=“javascript”>var url = '...';  // url to send request
var callback = function() {
    // callback function of http request
}

var xdr = new XDomainRequest();

xdr.onload = function() {
    var result;
    try {
        // do something with result
    } catch (error) {
        callback(error);
        return;
    }
   
    callback(null, result);
};
xdr.onerror = function() {
    callback(new Error('Network error'));
};

// ハンドラを追加
xdr.ontimeout = function () {
    callback(new Error('Network timeout'));
};
xdr.onprogress = function() {};  // 内容は空でもいいらしい

xdr.open('GET', url);
xdr.withCredentials = true;
xdr.send();</code></pre>

StackOverflow は POST についての話題だったが、今回のように GET でも同様だった。

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00EESW7JQ/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="http://ecx.images-amazon.com/images/I/517tJs%2B%2B%2BnL._SL160_.jpg" alt="Effective JavaScript　JavaScriptを使うときに知っておきたい68の冴えたやり方" style="border: none;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00EESW7JQ/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Effective JavaScript　JavaScriptを使うときに知っておきたい68の冴えたやり方</a><div class="amazlet-powered-date" style="font-size:80%;margin-top:5px;line-height:120%">posted with <a href="http://www.amazlet.com/" title="amazlet" target="_blank">amazlet</a> at 15.02.01</div></div><div class="amazlet-detail">翔泳社 (2013-04-13)<br />売り上げランキング: 4,709<br /></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00EESW7JQ/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
