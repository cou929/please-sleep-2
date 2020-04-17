{"title":"HTML5 Web Message のイントロダクション","date":"2013-04-16T00:02:29+09:00","tags":["javascript"]}

![](http://farm3.staticflickr.com/2469/3928105188_3e98ca7eb3.jpg)

[An Introduction to HTML5 web messaging - Dev.Opera](http://dev.opera.com/articles/view/window-postmessage-messagechannel/) を読んだメモ。MessageChannel は知らなかったので。

### 導入

Web messaging は異なるブラウジングコンテキスト間で、DOM を介すことなく、データを共有する手段だ。これには cross-document messaging (`window.postMessage()` など) と channel messaging (`MessageChannel`) の 2 種類がある。

### Message イベント

Cross document messaging, channel messaging, server-sent events, web sockets はすべて `message` イベントを発火させる。そのためまずは `message` イベントを見ていこう。

message イベントは [MessageEvent interface] [1] で定義されている。5 つのリードオンリー属性を持つ。

- data
  - string 形式のデータ
- origin
  - データ送信元のオリジン
- lastEventId
  - このメッセージイベントの ID
- source
  - データ送信元の window オブジェクト (より正確には [WindowProxy object] [2])
- ports
  - メッセージとともに送られた [MessagePort オブジェクト] [3] の配列

cross-document messaging と channel messaging の場合 `lasttEventId` は空だ。また `MessageEvent` は DOM の [Event interface] [4] を継承しているが、バブリングしない、キャンセルできない、デフォルトのアクションを持たないという特徴がある。

### Cross-document messaging

cross-document message を送るにはまずブラウジングコンテキストを作る必要がある。方法は新しい window を作るか、iframe の window を refer するかの 2 通りだ。メッセージは `postMessage()` で送信する。引数は次の 2 つ。

- message
  - 送信するメッセージ
- targetOrigin
  - メッセージを送るオリジン

message パラメーターには string だけでなく object や データオブジェクト (File や  ArrayBuffer) や配列も渡すことが可能だ。ただし IE8, 9 と Fx3.6 以下は string しか対応していない。

targetOrigin パラメーターにアスタリクスを設定すると送信するオリジンを絞ることがなくなる。データの漏洩につながるのでこの項目はきちんと設定することが推奨される。`/` を設定すると同一オリジンに限定される。

<pre><code data-language="javascript">var iframe = document.querySelector('iframe');
var button = document.querySelector('button');

var clickHandler = function(){
    // iframe.contentWindow refers to the iframe's window object.
    iframe.contentWindow.postMessage('The message to send.','http://dev.opera.com');
}

button.addEventListener('click',clickHandler,false);</code></pre>

`postMessage()` でメッセージが送信されると、受信側の message イベントが発火する。

<pre><code data-language="javascript">var messageEventHandler = function(event){
    // check that the origin is one we want.
    if(event.origin == 'http://dev.opera.com'){
alert(event.data);
    }
}
window.addEventListener('message', messageEventHandler,false);</code></pre>

受信側のブラウジングコンテキストが `postMessage()` を受け取る準備ができているかどうかをチェックしたい場合、受信側でロードが完了したら送信側へ message を送信、送信側は受信側からのメッセージを listen し、とどいたらメッセージを送るようにすれば良い。

<pre><code data-language="javascript">var clickHandler, messageHandler, button;

button = document.querySelector('button');

clickHandler = function(){
    window.open('otherpage.html','newwin','width=500,height=500');
}

button.addEventListener('click',clickHandler,false);

messageHandler = function(event){
    if(event.origin == 'http://foo.example'){
        event.source.postMessage('This is the message.','http://foo.example');
    }
}

window.addEventListener('message',messageHandler, false);</code></pre>

<pre><code data-language="javascript">var loadHandler = function(event){
    event.currentTarget.opener.postMessage('ready','http://foo.example');
}
window.addEventListener('DOMContentLoaded', loadHandler, false);</code></pre>

### Channel messaging

channel messaging はブラウジングコンテキスト間のダイレクトな双方向通信を提供する。 cross-document messaging 同様に DOM は介さず、 port 間の通信を行う。channel messaging は複数オリジン間の通信に便利だ。

#### MessageChannel と MessagePort オブジェクト

MessageChannel オブジェクトを作ると 2 つの関連するポートが作られる。一つは送信側に、もうひとつは他のブラウジングコンテキストに転送される。

それぞれのポートは [MessagePort] [5] オブジェクトで、次のメソッドを持つ。

- postMessage()
  - チャンネルを通じてメッセージを送る
- start()
  - このポートで受けたメッセージのディスパッチを開始する
- close()
  - ポートを閉じる

`MessagePort` オブジェクトは `onmessage` 属性を持つ。これは message イベントを listen する代わりのものだ。

#### ポートとメッセージの送信

ここでは、ドキュメントの中に 2 つの iframe があり、片方からもう片方へメッセージを送る例を示す。

ひとつめの iframe では以下を行う。

- MessageChannel オブジェクトの作成
- ひとつの MessageChannel ポートを親ドキュメントに転送する。これをもう片方の iframe へ転送してもらうためだ。
- もう片方のポートでイベントを listen する。メッセージを受信するため。
- ポートをオープンして受信可能な状態にする。

<pre><code data-language="javascript">var loadHandler = function(){
    var mc, portMessageHandler;

    mc = new MessageChannel();

    // Send a port to our parent document.
    window.parent.postMessage('documentAHasLoaded','http://foo.example',[mc.port2]);

    // Define our message event handler.
    portMessageHandler = function(portMsgEvent){
        alert( portMsgEvent.data );
    }

    // Set up our port event listener.
    mc.port1.addEventListener('message', portMessageHandler, false);

    // Open the port
    mc.port1.start();
}

window.addEventListener('DOMContentLoaded', loadHandler, false);</code></pre>

親ドキュメントはポートを受け取るとそれをもう一つの iframe へ post する。

<pre><code data-language="javascript">var loadHandler = function(){
    var iframes, messageHandler;

    iframes = window.frames;

    // Define our message handler.
    messageHandler = function(messageEvent){
        if( messageEvent.ports.length > 0 ){
            // transfer the port to iframe[1]
            iframes[1].postMessage('portopen','http://foo.example',messageEvent.ports);
        }
    }

    // Listen for the message from iframe[0]
    window.addEventListener('message',messageHandler,false);
}

window.addEventListener('DOMContentLoaded',loadHandler,false);</code></pre>

最後に、2 つ目の iframe はポートを受け取るとメッセージを送信する。メッセージは最初の iframe の `portMsgHandler` 関数がハンドルする。

<pre><code data-language="javascript">var loadHandler(){
    // Define our message handler function
    var messageHandler = function(messageEvent){

        // Our form submission handler
        var formHandler = function(){
            var msg = 'add <foo@example.com> to game circle.';
            messageEvent.ports[0].postMessage(msg);
        }
        document.forms[0].addEventListener('submit',formHandler,false);
    }
    window.addEventListener('message',messageHandler,false);
}

window.addEventListener('DOMContentLoaded',loadHandler,false);</code></pre>

サンプルコードは簡略化されていて、本来ならば MessageChannel がサポートされているかのチェックや、 origin が予期しているものかどうかのチェックをすべきである。

### 参考

- [W3C の仕様] [6]
- [Mike Taylor によるスライド] [7]

[cover photo by Sergio Aguirre](http://www.flickr.com/photos/sergiodjt/3928105188/)

[1]: http://dev.w3.org/html5/postmsg/#event-definitions
[2]: http://www.whatwg.org/specs/web-apps/current-work/multipage/browsers.html#windowproxy
[3]: http://dev.w3.org/html5/postmsg/#messageport
[4]: https://dvcs.w3.org/hg/domcore/raw-file/tip/Overview.html#interface-event
[5]: http://dev.w3.org/html5/postmsg/#messageport
[6]: http://dev.w3.org/html5/postmsg/
[7]: http://www.slideshare.net/miketaylr/html5-web-messaging-7970364
