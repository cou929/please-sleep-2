{"title":"ブラウザキャッシュのまとめ","date":"2015-03-07T20:12:46+09:00","tags":["web"]}

古い文章だけど、[Caching Tutorial for Web Authors and Webmasters](http://www.mnot.net/cache_docs/) を読んだメモ。ウェブマスター向けのブラウザキャッシュの概要。

ブラウザキャッシュは、ブラウザがサーバからレスポンスされたリソースを (場合によっては) ローカルに保存しておいて、二度目以降の表示では (場合によっては) キャッシュから読むという動作。無駄なリクエストが減ることでサーバリソースにも優しいし、ユーザーにとっては表示速度向上というメリットがある。

コンテンツをどのようにキャッシュするかを、サーバが UA に HTTP ヘッダを利用して指示する。`Expires` や `Cache-Control` ヘッダがその例だ。html の meta 要素で、たとえば `<meta http-equiv="Cache-Control" content="no-store" />` というふうに、キャッシュをコントロールする方法もあるが、多くのブラウザは meta 要素での指定に対応していない。原則として HTTP レスポンスヘッダでコントロールするのがベターなようだ。

ブラウザキャッシュの処理フローはつぎのようになっている。

1. レスポンスヘッダがキャッシュを禁止していた場合、キャッシュしない。
2. 認証がひつようなリクエストだったり SSL 通信の場合、キャッシュしない
3. キャッシュが新鮮な場合、キャッシュからリソースをロードする。キャッシュが新鮮と判断されるのは次のケースだ。
   - レスポンスヘッダに明示的にキャッシュの有効期限が指定されている場合。その期間内ならば新鮮と判断される。
   - キャッシュが最近保存されたもので、かつそのコンテンツの最終更新日時が比較的古い場合。
4. キャッシュが新鮮でない場合、サーバに問い合わせ validation が行われる。サーバがコンテンツの更新が行われていないと返答した場合、まだキャッシュが有効と判断されキャッシュからコンテンツが読み込まれる。
5. キャッシュが古いと判断されていても、例えばネットワークが切断されている場合などはキャッシュから読み込まれる。

ブラウザキャッシュに限定せず、一般的な意味でのキャッシュを考えたとき、キャッシュのコントロールで重要なのは、何をキャッシュするかとそのキャッシュの新鮮さをいかにして判定するかの二点だ。一方ブラウザキャッシュでは何をキャッシュするかのコントロールは基本的にブラウザが行い、コンテンツ配信者がコントロールできるのはキャッシュをしてほしくないということを宣言することのみだ。またブラウザキャッシュにはキャッシュが切れた場合にすぐにオリジンにとりにいくのではなく、一度コンテンツの更新有無を問い合わせる validation というステップが存在する。このように一般的なキャッシュとはすこし趣がことなる。またコンテンツ配信者視点で考えると、ブラウザキャッシュコントロールのポイントは新鮮さの指定と validation の 2 点といえるだろう。

では、このフローを上からみていこう。

まずはキャッシュ禁止の指定方法。これには `Cache-Control` レスポンスヘッダを用いる。このヘッダに `no-cache` を指定すると、ブラウザは毎回サーバに問い合わせ、更新がない場合のみキャッシュを利用する。`no-store` を指定するとそもそもそのコンテンツのキャッシュを保存しなくなる。

つぎは新鮮さの指定。HTTP 1.0 では `Expires` レスポンスヘッダが定義されている。Expires にはそのコンテンツが新鮮でなくなる日付時刻が GMT で指定される。この期日を過ぎていると、そのキャッシュは新鮮でないと判定され、サーバに問い合わせがなげられることになる。Expires にはいくつかの問題点もある。ひとつはサーバとクライアントの時間設定が必ず同期されていないといけないということ、もうひとつは日付時刻での指定なので、更新を忘れがちになるということだ。それをふまえ HTTP 1.1 では、先程も登場したが、 `Cache-Control` レスポンスヘッダが定義されている。Cache-Control には `max-age=[seconds]` というふうに、具体的な日付時刻ではなくレスポンスを受け取ってからの期間で指定することができる。

validation の方法は 2 種類ある。ひとつは時間。サーバ側は `Last-Modified` というレスポンスヘッダをつけ、ブラウザはキャッシュの最終更新日時を記録する。つぎにブラウザが `If-Modified-Since` リクエストヘッダで記録した日時を送ることで、サーバ側はキャッシュ更新の可否を判断できるというものだ。一方で HTTP 1.1 で導入されたのは `ETag` だ。これはコンテンツの内容を表すなんらかのハッシュ値で、コンテンツが更新されるたびに別の値になる。ブラウザが validation をリクエストする際に `If-None-Match` リクエストヘッダで ETag を送信すれば、サーバ側は現在の ETag 値とこれを比較し判断することができる。

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/487311361X/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="http://ecx.images-amazon.com/images/I/51hIDIWHmYL._SL160_.jpg" alt="ハイパフォーマンスWebサイト ―高速サイトを実現する14のルール" style="border: none;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/487311361X/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">ハイパフォーマンスWebサイト ―高速サイトを実現する14のルール</a><div class="amazlet-powered-date" style="font-size:80%;margin-top:5px;line-height:120%">posted with <a href="http://www.amazlet.com/" title="amazlet" target="_blank">amazlet</a> at 15.03.07</div></div><div class="amazlet-detail">Steve Souders スティーブ サウダーズ <br />オライリージャパン <br />売り上げランキング: 147,818<br /></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/487311361X/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
