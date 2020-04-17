{"title":"UNIX という考え方 - その設計思想と哲学","date":"2019-01-26T21:24:25+09:00","tags":["book"]}

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4274064069/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://images-fe.ssl-images-amazon.com/images/I/518ME653H3L._SL160_.jpg" alt="UNIXという考え方―その設計思想と哲学" style="border: none;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4274064069/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">UNIXという考え方―その設計思想と哲学</a><div class="amazlet-powered-date" style="font-size:80%;margin-top:5px;line-height:120%">posted with <a href="http://www.amazlet.com/" title="amazlet" target="_blank">amazlet</a> at 19.01.26</div></div><div class="amazlet-detail">Mike Gancarz <br />オーム社 <br />売り上げランキング: 6,126<br /></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4274064069/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

UNIX の設計思想について解説した本。原作者が思想を語る系ではなく、歴史学者が振り返ってまとめた系だった。UNIX 自体がコミュニティによって改善されてきたものなので、後から歴史を整理する方が必然性があると思う。

UNIX は長い時間と多くの人の貢献を経て確立してきたので、その設計思想はソフトウェア開発の多くの場面でうまく適用できる。少なくとも Web 系のソフトウェアエンジニアには文句なくお勧めできると思った。広い意味での「設計」やあるいは技術選定をうまくやるためには、良い設計とは何かを判断できる審美眼、あるいはセンスが必要。こうしたセンスを磨くためのインプットとして本書はぴったりだ。

本書では UNIX の設計思想を 10 の定理で紹介している。その内容を一言でいうと、「一つのことをうまくやるツールを組み合わせて、レバレッジの効いたアウトプットを生み出す」ということだ。「一つのことをうまくやる」「良い開発者は盗む (車輪の再発明をしない)」「できるだけ早く試作する (それは問題の正確な理解と、本当に必要なものへの集中を促す)」といった考えでシンプルなツールを作ること。そして「移植性の重視」「すべてのプログラムはフィルタ (あるいは標準入出力の存在)」という考えでそれらのツールを柔軟に組み合わせる。一つ一つは非力なツールでも、最大限にレバレッジが効いて、最終的には高度なアウトプットが実現できる。そしてこれが、実現できることの多さ、開発への参加のしやすさ、メンテナンスのしやすさにつながり、多人数で長期間にわたって開発・成長するシステムにはマッチするのだと思う。

最終章では他の OS (Atari, Microsoft, DEC) と UNIX を比較して考察しているのが興味深い。当然だけど他の OS もそれぞれの設計思想があり、その分野ではうまくやっている。ここからは個人的な浅い考察だけれども、UNIX の設計思想が適用できるのは「リテラシーの高いユーザーに限定できる」「複数人で開発する」「長期間にわたっての改善が必要」な環境だと思った。

単機能の処理を組み合わせて実現できることのレバレッジを効かせるには、(多機能でチュートリアルが手厚い単一のプログラムと比べて) ユーザーに学習コストを負わせるし、すべての開発者にそのマナーに従ってもらう必要がある。よってある程度以上のリテラシー、例えばソフトウェア開発者しか使わない、がないと厳しい。そのためこの思想が適用しやすいのは、UNIX 上で動作する CLI ツールを作る際は当然として、例えば開発者向けのライブラリの設計や、Web アプリケーションの設計には活かせると思う。いずれもユーザーはソフトウェア開発者になるはずなので。逆に、例えば一般のエンドユーザーへの UX 設計には適用できないと思う。学習コストをさげ、使い方のミスに先回りする、ある意味性悪説的な設計が必要になりそうだ。

また一度作って売り切りのメンテナンスのないシステムや、短期間のスケジュールがマストなプロジェクトでも当てはまらないかもしれない。こうしたシステムでは移植性やレバレッジよりも、パフォーマンスや短期的な開発スピードの方が優先されるはずだ。一方で一般的なライブラリ開発や Web 開発は、むしろ長期的な改善こそが重要なので、とても当てはまるのだと思う。
