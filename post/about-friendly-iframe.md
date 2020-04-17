{"title":"Friendly iFrame とサードパーティスクリプトのロード","date":"2014-11-02T12:13:25+09:00","tags":["javascript"]}

Friendly iFrame という、サードパーティスクリプトの呼び出し方法について調べた。手法自体はかなり古いもので、後述する iAB のドキュメントは 2008 年に出ているものだ。Google の DFP などでは現在でも使われている。

### Friendly iFrame (FiF) とは

サードパーティのスクリプト (広告タグや SNS のシェアボタン、ブログパーツなど) をページに埋め込む手法のひとつ。より狭義には、iAB が出す [Best Practices for Rich Media Ads
in Asynchronous Ad Environments (pdf)](http://www.iab.net/media/file/rich_media_ajax_best_practices.pdf) で説明されている、(主に) リッチアドのアドタグを設計・設置する際のベストプラクティス。

### 利点と他手法との比較

Friendly iFrame を他のサードパーティスクリプトの埋め込み手法と比較すると、ページのレンダリングを妨げないこと、どのようなスクリプトにも対応できること、機能的なディスアドバンテージ (他の手法ではできるのに FiF でできないこと) がないという点が優れている。

ここではサードパーティスクリプトの埋め込み手法を、ざっくりと次の 4 つに分類して、それぞれみていくことにする。

- 同期的 script 要素
- 通常の iframe
- 非同期 script 要素
- Friendly iFrame

#### 同期的 script 要素

普通に script 要素を設置し、呼び出されたサードパーティスクリプトは `document.write()` なり DOM 操作なりで動作するという、最もナイーブな実装。

当然、スクリプトの読み込みと実行がページレンダリングをブロックする。これに対応するため、`</body>` タグの直前にスクリプトを配置するというプラクティスが広く実施されてきた。単純だがどんなに古いブラウザでも確実に一定の効果が得られるので、現在でも意味があるノウハウだと思う。

一方で body の末尾に置いたとしても `window.onload` や `DOMContentLoaded` の発火は当然遅れるので、ブラウザのローディングインジケーターが回りっぱなしになったり、あるいはそれらのイベントを Listen している処理がそのぶん遅れて実行されるという問題は残る。ページロードを待ってからスクリプトを呼び出すといった工夫をしてもいいが、複雑さが増すし、サードパーティスクリプトが `document.write()` を使っている場合のハンドリングが依然大変だったりする。

#### 通常の iframe

iframe にサードパーティのコンテンツを読み込む方法。単純に src 属性がサードパーティのサービスを向いている iframe を埋め込むというもの。サードパーティのサービスはその中にコンテンツを読みこむなり、必要な処理をする。広告タグが最もポピュラーな例だと思う。広告を貼りたい媒体は広告業者の iframe タグを自サイトに設置し、広告サーバーはその iframe の中に広告の画像などを配信する。

iframe の設置さえできてしまえば、以降の処理はページ本体のレンダリングとは独立して動くので、レンダリングを妨げない。iframe の中では `document.write()` を使おうが何をしようが、親ドキュメントの邪魔をすることはない。

一方で、まずは当然だが、そのサードパーティのサービスが iframe に対応している必要がある。また、これも当然だが、別ドメインの iframe なので Same Origin Policy に縛られる。エキスパンド型のリッチアドやページコンテンツを解析するたぐいの計測ツールなど、親ドキュメントにアクセスする必要があるサービスの場合、根本的にこの方法は使えない。

#### 非同期 script 要素

スクリプト要素に `async` や `defer` 属性をつけて呼び出す方法。スクリプトの読み込み・実行ともにレンダリングをブロックしなくなる。属性をつけるだけなのでシンプルなのも魅力だ。

この方針のひとつの問題は、呼び出されるスクリプトが `document.write()` を使ってはいけないということだ。仮にページロード後に `async` で読み込まれたスクリプトが `document.write()` すると、ページ全体が書き換わり真っ白になってしまうなど、致命的な問題になる。よってレガシーな広告タグやブログパーツなどにこの方法を適用することができない。

またこの方法でも、[DOMContentLoaded は大丈夫だが window.onload が遅れるブラウザがある](https://www.facebook.com/notes/facebook-engineering/under-the-hood-the-javascript-sdk-truly-asynchronous-loading/10151176218703920) らしい。些細かもしれないがこの問題はまだ残る。

余談だが `async` 属性は準備ができ次第すぐにロード・実行され、`defer` 属性はドキュメントのパース完了後の実行キューにそのスクリプトを入れる。つまり、`async` は実行順序が保証されないので、依存関係のあるスクリプトの読み込みには注意する必要がある。

#### Friendly iFrame

親と同一ドメインの iframe を動的に作り、その中にサードパーティスクリプトを読み込む方法。

この方法は iframe なので当然レンダリングを妨げないし、`document.write()` を使うようなレガシーな実装でも問題ない。さらに Same Origin Policy に縛られないので、サードパーティ側が対応しさえすれば、親ドキュメントにアクセスする必要があるサービスでも利用できる点が優れている。

欠点としては、実装がやや複雑になる点、エキスパンド広告等の場合はサードパーティ側の対応が必要な点、iframe の src に指定するための空 html を自ドメインでホストする必要がある点 (これは後述する実装方法によっては不要) があげられる。

### 実装方法

ここでは [iAB のベストプラクティス (pdf)](http://www.iab.net/media/file/rich_media_ajax_best_practices.pdf) に記載されている方法を取り上げる。リッチアドなど広告を主眼にしているが、ほとんど一般的な内容だ。

- src 属性を `about:self` にした iframe 要素を動的に生成する
- サードパーティの url を指した script 要素を作成し、iframe 内のドキュメントに設置する
- `inDapIF = true` というフラグを iframe 内に作る

`inDapIF` は呼び出されたサードパーティのスクリプトが Friendly iFrame で呼び出されたことを知るためのフラグだ。

これで IE6 以上のほとんどのブラウザをサポートできるが、唯一 Firefox2.0 が対応できない。これをサポートするためには、次のように自ドメインから空 html をホストする必要がある。

- iframe 要素を生成する。ただし src 属性は自ドメインでホスティングしている小さな html ファイルに向ける
- 以降は先ほどと同様で、サードパーティを向いた script 要素と `inDapIF` フラグを iframe 内に作成する

### 参考

- [Best Practices for Rich Media Ads
in Asynchronous Ad Environments (pdf)](http://www.iab.net/media/file/rich_media_ajax_best_practices.pdf)
- [Under the Hood: The JavaScript SDK – Truly Asynchronous Loading](https://www.facebook.com/notes/facebook-engineering/under-the-hood-the-javascript-sdk-truly-asynchronous-loading/10151176218703920)
- [Friendly Iframes [Cxense Display Classroom]](http://classroom.emediate.com/doku.php/technical:integration:friendly_iframes:start)
- [Choosing the right GPT mode and tag type - DoubleClick for Publishers Help](https://support.google.com/dfp_premium/answer/183282?hl=en)
- [scriptのdefer/asyncを理解し、ページの高速化方法を探る | ゆっくりと…](http://tokkono.cute.coocan.jp/blog/slow/index.php/xhtmlcss/how-to-use-script-defer-and-async-for-performance-enhancement/)
- [document.writeを使った遅いブログパーツ（例えばzenback）を非同期化してサイトを高速表示する方法 | ゆっくりと…](http://tokkono.cute.coocan.jp/blog/slow/index.php/xhtmlcss/script-async-loading-in-iframe/)

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00I5BNSNC/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="http://ecx.images-amazon.com/images/I/51E0Xgh1pDL._SL160_.jpg" alt="サードパーティJavaScript (アスキー書籍)" style="border: none;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00I5BNSNC/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">サードパーティJavaScript (アスキー書籍)</a><div class="amazlet-powered-date" style="font-size:80%;margin-top:5px;line-height:120%">posted with <a href="http://www.amazlet.com/" title="amazlet" target="_blank">amazlet</a> at 14.11.02</div></div><div class="amazlet-detail">KADOKAWA / アスキー・メディアワークス (2014-02-10)<br />売り上げランキング: 19,591<br /></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00I5BNSNC/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
