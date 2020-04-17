{"title":"img 要素の src 属性と onload イベントリスナの設定タイミング","date":"2013-12-25T22:54:55+09:00","tags":["javascript"]}

<pre><code data-language="javascript">var img = new Image();
img.onload = function() {
    // event handler
};
img.src = 'http://example.com/foo.png';
document.body.appendChild(img);</code></pre>

こういうふうに動的に画像をロードして、かつ onload のイベントを取りたい場合。三行目で src 属性に url を設定した時点で即座に非同期のリクエストが飛ぶ。画像のリクエストが完了する前に次の行へ処理が移る。よって src 属性に値をセットする前にイベントリスナの設定をしなくてはいけない。

というのも、たぶん src の挙動をちゃんと理解していなくて、以下のようなコードをサンプルとしてあげているブログなどを複数目にしたので気になったという経緯。

<pre><code data-language="javascript">var img = new Image();
img.src = 'http://example.com/foo.png';
img.onload = function() {  // このタイミングではすでにロードが完了している可能性がある
    // event handler
};
document.body.appendChild(img);</code></pre>

この書き方だと潜在的にはすべての環境でレースコンディションの状態になると思っている。そして特に問題になるのが古い (おそらく 8 以前の) IE だ。src に設定している画像がブラウザキャッシュから読み込まれた場合、おそらく onload イベントの設定より前に画像のロードが終わるので、ハンドラの関数が呼ばれないということが高い確率でおこる。

この挙動を IE のバグと言って onlaod を先に書くと説明していたり (結論はあっているが原因の認識が異なる)、あるいは "IE の場合は onreadystatechange にハンドラを設定してその中で readyState を見る" という workaround を紹介しているものもあって、ちょっともやもやした次第。

簡単な検証コードだが、以下の html を IE 8 などで開くて何度かリロードすると現象を再現できると思う。

<script src="https://gist.github.com/cou929/8123027.js"></script>

