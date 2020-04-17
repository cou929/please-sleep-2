{"title":"jQuery オブジェクトと DOM エレメントの変換","date":"2012-04-06T22:48:00+09:00","tags":["javascript"]}

ジェイクエリーすぐ忘れる

### DOM エレメント -> jQuery オブジェクト

`$()` 関数に入れてあげれば OK

    var sample = document.querySelector('#sample');
    var jq_obj = $(sample);

### jQuery オブジェクト -> DOM エレメント

jQuery オブジェクトのインデックス 0 らしい

    var dom_element = jq_obj[0];

### 検証

    var body = document.querySelector('body');
    var jq_obj = $(body);
    var dom_elm = jq_obj[0];
    body === dom_elm   // true

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117836/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://images-fe.ssl-images-amazon.com/images/I/51U44SJi3jL._SL160_.jpg" alt="初めてのJavaScript 第3版 ―ES2015以降の最新ウェブ開発" style="border: none;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117836/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">初めてのJavaScript 第3版 ―ES2015以降の最新ウェブ開発</a><div class="amazlet-powered-date" style="font-size:80%;margin-top:5px;line-height:120%">posted with <a href="http://www.amazlet.com/" title="amazlet" target="_blank">amazlet</a> at 20.03.08</div></div><div class="amazlet-detail">Ethan Brown <br />オライリージャパン <br />売り上げランキング: 17,881<br /></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117836/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
