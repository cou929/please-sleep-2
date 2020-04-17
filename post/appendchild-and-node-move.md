{"title":"appendChild とノードの移動","date":"2015-02-01T22:56:05+09:00","tags":["javascript"]}

`appendChild` や `insertBefore` でノードを追加すると、元の位置からは削除され、ノードが移動する挙動になる。この例では、ボタンを押すと `appendChild` でノードが移動しているのがわかる。

<iframe width="100%" height="300" src="http://jsfiddle.net/cou929/smf8yzvr/1/embedded/" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

body 以下に直接ついていないノード、例えば動的に `createElement` したノードでも同様だ。

次の例では、3 つのノードを作りそれをループで `appendChild` している。`appendChild` するたびに追加したノードは `nodes` リストから削除される。そのため、1 番目と 3 番目のノードだけが追加されることになる。一度目のループでは 0 番目の要素が append され、それがリストから削除される。2 度目のループの際にはリストの長さは 2 になっており、その 1 番目の要素、つまりもとのリストの最後の要素が append される

<iframe width="100%" height="300" src="http://jsfiddle.net/cou929/jeb1gbve/2/embedded/" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

このようにループと組み合わさることで、一見直感に反する挙動になってしまうことがある。今回の場合は `NodeList` を配列にコピーするとこれを回避できる。この例では `Array.prototype.slice.call` で NodeList を配列にしている。意図したとおり 3 つの要素すべてが append される。

<iframe width="100%" height="300" src="http://jsfiddle.net/cou929/qoo1jp30/embedded/" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

または `cloneNode` でノードをコピーしてもよい。

<iframe width="100%" height="300" src="http://jsfiddle.net/cou929/b3Lvxof2/embedded/" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

`appendChild` だけではなく `insertBefore` などでも同様の挙動になる。

### 参考

- [Document Object Model Core](http://www.w3.org/TR/2000/REC-DOM-Level-2-Core-20001113/core.html#ID-184E7107)
- [Node.appendChild - Web API インターフェイス | MDN](https://developer.mozilla.org/ja/docs/Web/API/Node.appendChild#Notes)

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00P2EG5LC/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="http://ecx.images-amazon.com/images/I/51OtMPYsLwL._SL160_.jpg" alt="パーフェクトJavaScript" style="border: none;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00P2EG5LC/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">パーフェクトJavaScript</a><div class="amazlet-powered-date" style="font-size:80%;margin-top:5px;line-height:120%">posted with <a href="http://www.amazlet.com/" title="amazlet" target="_blank">amazlet</a> at 15.02.01</div></div><div class="amazlet-detail">技術評論社 (2014-10-31)<br />売り上げランキング: 8,486<br /></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00P2EG5LC/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
