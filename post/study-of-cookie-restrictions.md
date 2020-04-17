{"title":"ドメインあたりのクッキー数上限と上限に達した場合の挙動","date":"2013-12-09T18:32:37+09:00","tags":["browser"]}

![](/images/cookie-monster.jpg)

BrowserStack にてそれぞれデフォルト設定のブラウザで調査。IE 中心に調べたので他のブラウザは網羅的ではない。あとから補完して別途公開したい。

### サマリー

最も一般的な挙動は、追加日時 (おそらく最終更新日時) が古いものを一件削除し、新しいクッキーを受け入れるという、LRU 的なアルゴリズム。ここからブラウザやバージョンによってバリエーションがある。

- Chrome は一件ではなく一度に三十件削除する
  - その代わり受け入れるクッキー数は多め
  - バックエンドを Chromium (Blink) に切り替えてからの Opera も同様
- 古い Opera は、追加しようとしたクッキーを受け入れ、その次に新しいもの一件を削除する。
- Safari は単に追加順ではなく独自のソート順でクッキーを管理。その降順または昇順で一件を削除する。
  - バージョンによって動きがばらばらなので、詳しく調査したい。

### IE

<table border="1">
<tr><th>OS</th><th>Browser</th><th>Max cookie per domain</th><th>Behavior</th></tr>
<tr><td>WinXP</td><td>IE6</td><td>20</td><td>追加されたのが最も古いクッキー一件を削除</td></tr>
<tr><td>WinXP</td><td>IE7</td><td>50</td><td>追加されたのが最も古いクッキー一件を削除</td></tr>
<tr><td>WinXP</td><td>IE8</td><td>50</td><td>追加されたのが最も古いクッキー一件を削除</td></tr>
<tr><td>Win7</td><td>IE8</td><td>50</td><td>追加されたのが最も古いクッキー一件を削除</td></tr>
<tr><td>Win7</td><td>IE9</td><td>50</td><td>追加されたのが最も古いクッキー一件を削除</td></tr>
<tr><td>Win7</td><td>IE10</td><td>50</td><td>追加されたのが最も古いクッキー一件を削除</td></tr>
<tr><td>Win7</td><td>IE11</td><td>50</td><td>追加されたのが最も古いクッキー一件を削除</td></tr>
<tr><td>Win8</td><td>IE10.0</td><td>50</td><td>追加されたのが最も古いクッキー一件を削除</td></tr>
<tr><td>Win8</td><td>IE10.0 Desktop</td><td>50</td><td>追加されたのが最も古いクッキー一件を削除</td></tr>
<tr><td>Win8.1</td><td>IE11.0</td><td>50</td><td>追加されたのが最も古いクッキー一件を削除</td></tr>
<tr><td>Win8.1</td><td>IE11.0 Desktop</td><td>50</td><td>追加されたのが最も古いクッキー一件を削除</td></tr>
</table>

### Firefox

<table border="1">
<tr><th>OS</th><th>Browser</th><th>Max cookie per domain</th><th>Behavior</th></tr>
<tr><td>Win8</td><td>Firefox 16.0</td><td>150</td><td>追加されたのが最も古いクッキー一件を削除</td></tr>
<tr><td>WinXP</td><td>Firefox 25.0</td><td>150</td><td>追加されたのが最も古いクッキー一件を削除</td></tr>
<tr><td>Mac OSX Mavericks</td><td>Firefox 27.0</td><td>150</td><td>追加されたのが最も古いクッキー一件を削除</td></tr>
</table>

### Chrome

<table border="1">
<tr><th>OS</th><th>Browser</th><th>Max cookie per domain</th><th>Behavior</th></tr>
<tr><td>Win8</td><td>Chrome 22.0</td><td>180</td><td>古いものから30件を削除</td></tr>
<tr><td>WinXP</td><td>Chrome 31.0</td><td>180</td><td>古いものから30件を削除</td></tr>
<tr><td>Mac OSX Mavericks</td><td>Chrome 33.0</td><td>180</td><td>古いものから30件を削除</td></tr>
</table>

### Opera

<table border="1">
<tr><th>OS</th><th>Browser</th><th>Max cookie per domain</th><th>Behavior</th></tr>
<tr><td>Mac OSX Mavericks</td><td>Opera 11.0</td><td>60</td><td>追加されたのが新しいクッキー一件を削除 (set しようとしたクッキーは受け入れ、その次に新しいものを削除する)</td></tr>
<tr><td>Win8</td><td>Opera 12.0</td><td>60</td><td>追加されたのが新しいクッキー一件を削除 (set しようとしたクッキーは受け入れ、その次に新しいものを削除する)</td></tr>
<tr><td>WinXP</td><td>Opera 17.0</td><td>180</td><td>古いものから30件を削除</td></tr>
<tr><td>Mac OSX Mavericks</td><td>Opera 19.0</td><td>180</td><td>古いものから30件を削除</td></tr>
</table>

- opera は Version 12 を最後にレンダリングエンジンを Presto から Blink に切り替えている。

### Safari

<table border="1">
<tr><th>OS</th><th>Browser</th><th>Max cookie per domain</th><th>Behavior</th></tr>
<tr><td>Mac OSX Snow Leopard</td><td>Safari 4.0</td><td>なし (容量制限のみ)</td><td>追加されたのが最も古いクッキー一件を削除</td></tr>
<tr><td>Mac OSX Snow Leopard</td><td>Safari 5.0</td><td>なし (容量制限のみ)</td><td>追加されたのが最も古いクッキー一件を削除</td></tr>
<tr><td>Mac OSX Lion</td><td>Safari 5.1</td><td>なし (容量制限のみ)</td><td>保持しているクッキーの並び順で若い方から一件を削除</td></tr>
<tr><td>Mac OSX Lion</td><td>Safari 6.0</td><td>なし (容量制限のみ)</td><td>保持しているクッキーの並び順で若い方から一件を削除</td></tr>
<tr><td>Mac OSX Mavericks</td><td>Safari 7.0</td><td>なし (容量制限のみ)</td><td>保持しているクッキーの並び順で遅い方から一件を削除</td></tr>
</table>

- Safari 4.0 のクッキーは追加日時の降順 (新しいものが先頭、古いものが後方) にならぶ。他のブラウザは昇順だった。
- Safari 5.1〜6.0 のクッキーは単に追加順ではなく、一定の法則でソートし保存されている。削除はこのソート順に行われるようだ。ソートの方法は短時間に cookie が set された場合、そのうちいくつかをまとめて辞書順に登録しているように見える。サーバサイドからの Set-Cookie ヘッダをどう処理しているかはこの調査からはわからないので、別途。
- Safari 7.0 は逆にソートの降順に一件削除している。

Safari の動きは不可解なので、より詳細に調査したい。

### スクリプト

検証に使ったスクリプトはこちら

<script src="https://gist.github.com/cou929/7869555.js"></script>

おおまかな動きは次のようなシンプルなもの

- js で document.cookie に一件ずつクッキーを書き込み
  - 容量制限に引っかからないよう、キーは二文字、値は "1" という一文字にする
  - expire はランダムな未来の日付
- "書き込み試行回数 != document.cookie に入っているクッキー数" になった時点で、処理を終了
  - 書き込み試行回数、現在のクッキー数、document.cookie の中身を表示

