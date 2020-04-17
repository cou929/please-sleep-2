{"title":"document.cookie から読み書きできるのはなにか","date":"2013-03-25T12:43:42+09:00","tags":["browser"]}

- document.cookie はそのドキュメントに関連するクッキーの読み書きができる
  - document.cookie を呼ぶ script がサードパーティーのドメインからロードしたものでも同様
  - ブラウザのクッキー設定を辛くしている場合はその限りではない
  - `create*` メソッドで動的に作られた場合など, ドキュメントのロケーションが空だったり定かでない場合も自由な読み書きは不可?
- 書き込みの場合任意のドメインのクッキーを設定できるが, 実際に送信されるか・保存されるかはブラウザの設定次第

こう思っているのだけど本当にそうなのか確信がない (仕様で定義されている or コードにそう書いてあった). また `そのドキュメントに関連するクッキー` の定義がいまいち曖昧. ファーストパーティのドメインのクッキーだけなのか, 許可している場合はサードパーティーのクッキーも OK なのか (経験則としては後者なのだが).

つまりサードパーティーのスクリプトでも, そのドキュメントに関連するクッキーを読めるし, 任意のドメインのクッキーをセットできると思っているのだが, 本当にそうだっけという確証がない. サードパーティーのスクリプトでもファーストパーティの DOM を自由にさわれるので, クッキーも読み書きできても筋が通っている気がしている.

まずは仕様を読んで調査:

- [HTML Standard](http://www.whatwg.org/specs/web-apps/current-work/#dom-document-cookie)
- [document.cookie - Document Object Model (DOM) \| MDN](https://developer.mozilla.org/en-US/docs/DOM/document.cookie)
- [Document Object Model HTML](http://www.w3.org/TR/DOM-Level-2-HTML/html.html#ID-8747038)
- [[whatwg] A document's cookie context](http://lists.whatwg.org/pipermail/whatwg-whatwg.org/2008-June/015076.html)
