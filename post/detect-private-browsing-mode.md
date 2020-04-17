{"title":"プライベートブラウジングモードの判定","date":"2013-12-16T01:19:46+09:00","tags":["browser"]}

<a href="http://www.flickr.com/photos/holster/195031415/" title="Private by Richard Holt®, on Flickr"><img src="http://farm1.staticflickr.com/57/195031415_8702c6e446.jpg" width="500" height="375" alt="Private"></a>

 [プライベートブラウジングモードのいろいろ - Please Sleep](http://please-sleep.cou929.nu/surveys-of-private-browsing-mode.html) の続き。あるブラウザがプライベートブラウジングモード (Private Browsing, InPrivate Browsing, Incognito) であるか JavaScript で判定する方法について。

browser history sniffing の対策が入りプライベートモードの一般的な判定ができなくなったと思われたが、ブラウザによっていくつかの API の挙動が通常・プライベートモードで変化するため、それを利用して判定ができるということを教えていただいた。

<blockquote class="twitter-tweet" lang="en"><p><a href="https://twitter.com/cou929">@cou929</a> <a href="http://t.co/1OE9k3kJv9">http://t.co/1OE9k3kJv9</a> プライベートブラウズ判定方法あります、SafariだとlocalStorage,DNT Firefox Chromeは <a href="https://t.co/mDDRfUeCP9">https://t.co/mDDRfUeCP9</a></p>&mdash; mala (@bulkneets) <a href="https://twitter.com/bulkneets/statuses/410220936600616961">December 10, 2013</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

というわけでプライベートモードを判定するコードスニペットを書いてみた。

<script src="https://gist.github.com/cou929/7973956.js"></script>

detect-private-browsing.js が判定ロジックの実装、example.html が利用例だ。`detectPrivateMode` にコールバック関数を渡すと判定結果を引数にいれて呼び出してくれる。プライベートモードだった場合真、通常モードだった場合偽、判定できなかった場合は undefined が渡される。

この記事の発行時点で以下のブラウザで判定できることを確認した。

- IE 10 以降
- Firefox 16 以降
- Chrome 14 以降 (13 以前が browserstack に残っていなかったので未検証。webkitRequestFIleSystem があれば判定できるはず)
- Opera 15 以降
- Safari (含む IE Safsri) 4.0 以降 (それ以前が browserstack に残っていなかったので未検証)
  - 未検証だが iOS Safari も同様に判定できるはず

ブラウザごとの判定ロジックは次のようになっている。プライベートモードの性質的にストレージ系の挙動は通常時と変えざるをえないので、それを利用するイメージだ。

- IE は `window.indexedDB` の有無で判定する。プライベートモードだとこれが undefined になる。indexedDB は IE10 で実装されたためそれ以前のバージョンはこの方法で判定ができない。以前のバージョンの判定方法は見つけられなかった。
- Firefox も indexedDB で判定する。IE とは異なり indexedDB が open できるかどうかで判定する。特に例外が発生するわけではないので、window.indexedDB.open の戻り値を検査する。今回は result パラメータの有無で成功・失敗を判断することにした。また open メソッドは非同期なので、readyState プロパティが 'done' になっているかどうかをチェックする必要があるので注意。
- Chrome は `window.webkitRequestFileSystem` の呼び出しが成功するかどうかで判定する。webkitRequestFileSystem に成功・失敗時のコールバックを渡すことができるので、そこで結果をみればよい。
- Opera は Presto 時代のバージョンでの判定方法が見つけられなかった。15 以降だと Chrome と同じ方法でよい。
- Safari は localStorage の set で判定する。プライベートモードの場合例外が発生する。これは有名。

たぶん他にも判定方法はあるし、将来にわたって利用できるほど安定したものではないので注意。もし古い IE, Opera での方法も見つかったらアップデートしたい。

