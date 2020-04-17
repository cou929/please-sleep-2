{"title":"プライベートブラウジングモードのいろいろ","date":"2013-12-16T01:12:49+09:00","tags":["browser"]}

ブラウザのプライベートブラウジングモードについていろいろと調べた。知りたかったのは UA がプライベートモードなのか否かを判定する方法があるのかと、プライベートモードがどのくらい使われているのかのデータがあるのか。

### 結論

- <del>プライベートモードの判定はできない (過去は方法があった)</del>
  - 別のアプローチで判定できる。こちらを参照: [プライベートブラウジングモードの判定](http://please-sleep.cou929.nu/detect-private-browsing-mode.html)
- 2010 年の調査では約 10 % のユーザがプライベートモードを使用している

### 判定方法

この論文はプライベートブラウジングモードについていろいろと調査を行ったもの。

[An Analysis of Private Browsing Modes in Modern Browsers](http://elie.im/publication/an-analysis-of-private-browsing-modes-in-modern-browsers#.Up8abZCzsVk)

この中でプライベートモードにしている割合とサイトジャンルの関係を調査している部分がある。三種類のジャンルのサイトにアドネットワーク経由で広告を配信。その広告の中にモードを判定するスクリプトを含ませて集計したらしい。

判定方法は、配信しているサイトへのリンクのスタイルを getComputedStyle で読み取り、リンクの色から訪問済みかどうかをチェック。訪問済みでないならプライベートモードと判定できる、というもの。

このテクニックはわりと問題になっていたもので、`browser history sniffing` とも呼ばれているようだ。そして現在では以下のように対策がとられている。

[CSS の :visited に行われるプライバシー対策 | Mozilla Developer Street (modest)](https://dev.mozilla.jp/2010/04/privacy-related-changes-coming-to-css-vistited/)

getComputedStyle でも偽の情報が返されるようになっていて、スクリプトからの不正利用を防いでいる。現在ではすべてのブラウザが対応していて、もうこの方法で判定することはできない。

#### 2013-12-16 追記

いくつかのブラウザ・バージョンでは JavaScript API の挙動の変化によって判定する手段がありました。

[プライベートブラウジングモードの判定](http://please-sleep.cou929.nu/detect-private-browsing-mode.html)

### 利用状況調査

先ほどの 2010 年の論文によると、全体の約 10 % のインプレッションがプライベートモードだったそうだ。これはサイトのジャンルによってもばらつきがあり、アダルトや E コマースサイトではその割合が高くなっている。

こちらでは同じ著者が 2012 年に [Google Consumer Surveys](http://www.google.com/insights/consumersurveys/home) でアンケート調査を行った結果が紹介されている。

[19% of users use their browser private mode](http://elie.im/blog/privacy/19-of-users-use-their-browser-private-mode/#.Up8Y0ZCzsVk)

`Do you use private browsing ?` という質問に 19 % のが yes と答えたそうだ。論文の方のデータは実際にプライベートモードでアクセスがあった割合なので、その違いには注意。

もうひとつ、こちらは Mozilla が [Test Pilot](https://testpilot.mozillalabs.com/) で集めたデータをもとに集計したものだ。2010 年の調査。

[Understanding Private Browsing | Blog of Metrics](http://blog.mozilla.org/metrics/2010/08/23/understanding-private-browsing/)

プライベートモードがよく使われている時間帯は 12 時前後で (職場で昼休みに使っているのかな)、プライベートモードを継続している時間は 10 分間が最も多かったそうだ。

総合すると、全体の 1、2 割のユーザーはプライベートモードを利用しており、一回あたりは 10 分程度。よって総リクエストのうち〜1割くらいは、実際には同じなのに別のユーザーとして計測されていると考えることもできるかもしれない。各ユーザーの利用頻度のデータがないので正確ではないし、サイトのジャンルによるばらつきも大きいので、正確ではないが、頭にとどめておきたい。

