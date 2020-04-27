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

```html
<!DOCTYPE html>
<html>
<head>
<script>
// for caching
(new Image()).src = 'http://upload.wikimedia.org/wikipedia/commons/2/23/1x1.GIF?a=1';
</script>
</head>
<body>
<div>
  <h2>set src attribute before attaching onload listener</h2>
  <p>is onload called: <span id="result1"></span></p>
  <p>complete attribute: <span id="result2"></span></p>
</div>

<div>
  <h2>set src attribute after attaching onload listener</h2>
  <p>is onload called: <span id="result3"></span></p>
  <p>complete attribute: <span id="result4"></span></p>
</div>

<script>
setTimeout(function(){
    // set src attribute before attaching onload listener
    var img = new Image();
    img.src = 'http://upload.wikimedia.org/wikipedia/commons/2/23/1x1.GIF?a=2';
    img.onload = function() { document.getElementById('result1').innerHTML = 'called'; };
    document.getElementById('result2').innerHTML = img.complete;
}, 500);

setTimeout(function(){
    // set src attribute after attaching onload listener
    var img = new Image();
    img.onload = function() { document.getElementById('result3').innerHTML = 'called'; };
    img.src = 'http://upload.wikimedia.org/wikipedia/commons/2/23/1x1.GIF?a=3';
    document.getElementById('result4').innerHTML = img.complete;
}, 500);
</script>
</body>
</html>
```

<iframe style="width:120px;height:240px;" marginwidth="0" marginheight="0" scrolling="no" frameborder="0" src="//rcm-fe.amazon-adsystem.com/e/cm?lt1=_blank&bc1=000000&IS2=1&bg1=FFFFFF&fc1=000000&lc1=0000FF&t=pleasesleep-22&language=ja_JP&o=9&p=8&l=as4&m=amazon&f=ifr&ref=as_ss_li_til&asins=4048930737&linkId=ff6b7c48c954ab6b7faa261f503919f4"></iframe>
