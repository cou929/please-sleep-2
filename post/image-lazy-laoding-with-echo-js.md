{"title":"Echo.js で画像を遅延ロード","date":"2014-03-22T01:16:16+09:00","tags":["html"]}

画像を遅延ロード、つまりユーザが該当部分までスクロールした時にはじめて画像をロードする UI、を実現したい。[jQuery の Lazy Load Plugin](http://www.appelsiini.net/projects/lazyload) というのがデファクトとしてあるようだが、今回はモバイル向けのサイトがターゲットだったので、このために jQuery をまるまる配信するのは避けたかった。

例えば zepto ベースとかもっと軽いのを探していたところ、 Echo.js がよさげだった。

[Echo.js, simple JavaScript image lazy loading](http://toddmotto.com/echo-js-simple-javascript-image-lazy-loading/)

ライブラリには何も依存せず単体で動き、サイズも 1KB を切るのでいい感じ。使い勝手も lazy load plugin とほぼ同じ。

対象の img タグには、data 属性に画像の url を指定しておき、src にはロードまでの仮画像を指定しておく。

<pre><code data-language="html"><img src="img/blank.gif" alt="" data-echo="img/album-1.jpg"></code></pre>

仮画像には [echo.js のデモページ](http://toddmotto.com/labs/echo/) のようにローディングインジケータのアニメ gif を指定してもいいし、例えば空の gif を data url でこんな風に指定しておいてもいい。

<pre><code data-language="html"><img src="data:image/gif;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw==" data-echo="img/album-1.jpg"/></code></pre>

あとは echo.js を読み込んで実行してあげるだけでよい。

<pre><code data-language="html"><script src="echo.js"></script>
<script>
Echo.init();
</script></code></pre>

