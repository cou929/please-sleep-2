{"title":"みんなのGo言語 現場で使える実践テクニック","date":"2019-01-01T22:33:42+09:00","tags":["book"]}

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01LMS7B1O/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://images-fe.ssl-images-amazon.com/images/I/51lqgv%2BWyxL._SL160_.jpg" alt="みんなのGo言語[現場で使える実践テクニック]" style="border: none;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01LMS7B1O/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">みんなのGo言語[現場で使える実践テクニック]</a><div class="amazlet-powered-date" style="font-size:80%;margin-top:5px;line-height:120%">posted with <a href="http://www.amazlet.com/" title="amazlet" target="_blank">amazlet</a> at 19.01.01</div></div><div class="amazlet-detail">技術評論社 (2016-09-09)<br />売り上げランキング: 28,351<br /></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01LMS7B1O/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

Go を使って実務的なシステムを開発するためのベストプラクティス集といった本。他の言語だと、`Effective-x` 系とか、perl だと <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00G9QIN58/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">モダンPerl入門</a> にトーンは近いと思う。

Go の学習のながれとして、まずは [A Tour of Go](https://tour.golang.org/welcome/1) で文法的な部分を学んで、その後 [公式のドキュメント](https://golang.org/doc/) で言語仕様を補完していくのが鉄板だと思っている。その次に読む本としては本書はかなりいいと思う。一日で読み終わる程度のさらっとした分量なのでとっつきやすい。Software Design か Web+DB Press の特集をまとめたムック本くらいの内容なので、中級者以上には内容の深さが足りないかもしれないが、そういう意味でも文法を覚えたあとの初級者に次のステップとしておすすめできる。

個人的には 4 章の CLI ツールをつくるための tips を紹介した章がよかった。たとえばディレクト構成やテストしやすいパッケージの切り方など、Go らしいデファクトなやり方、いわゆるベストプラクティスを紹介してくれていたのが理由。デファクトのライブラリやその使い方の解説は、そのキーワードにさえたどり着ければ (もちろん有用なんだけれど) ある程度調べればわかることなんだけれど、こういうベストプラクティスはまず「Go らしい書き方をしているコード」を見分けられるようになって、それをいくつも読み込んでいかないと身につかないものだと思う。この学習コストを本書でスキップできたのは助かった。

2016 年に発売されたものなので、流石に古くなっている部分もある。たとえば vendoring に関しては glide が紹介されているが、当たり前だがその後の dep, Modules (vgo) の流れはカバーされていない。

- [Modules · golang/go Wiki](https://github.com/golang/go/wiki/Modules)
- [research\!rsc: Go & Versioning](https://research.swtch.com/vgo)

また Web アプリケーション (Web API) の開発運用に関するトピックがまったくないのも物足りない部分だと思った。最近では API を Go で書くことがかなりの市民権を得ている状況なので、この部分がカバーされるとよりよいかなと思った。(もちろんこれも、2016 年に出版された本に求めるものではない)

とはいえまだまだ有用な内容が多く、良書だった。言語仕様をひととおり見たあとに手に取る本としておすすめ。
