{"title":"最近読んだもの 47 - 分散 DB でのトランザクション分離レベルなど","date":"2022-06-26T23:00:00+09:00","tags":["readings"]}

- [MySQL "Got an error reading communication packet" \- Percona Database Performance Blog](https://www.percona.com/blog/2016/05/16/mysql-got-an-error-reading-communication-packet-errors/)
    - MySQL での `Got an error reading communication packet` エラーについて
    - そもそもトラブルシューティングが難しいと前提しつつ、泥臭い tips を列挙してくれていて良い
- [Pitfalls of isolation levels in distributed databases](https://planetscale.com/blog/pitfalls-of-isolation-levels-in-distributed-databases)
    - トランザクション分離レベルと分散データベースとの整理
    - Repeatable Read と Snapshot Read を分けて記載したり
    - 一般論として、アプリケーションは必要最低限の分離レベルで動作するように作っておくと、分散データストアを導入する際に良いという結論かな
- [Embracing Critical Voices \| July 2022 \| Communications of the ACM](https://cacm.acm.org/magazines/2022/7/262085-embracing-critical-voices/fulltext)
    - ACM のコミュニティ運営に関するセクションの Chair に就いた方の挨拶的な記事
    - 知らなかったが、論文のピアレビューの際に、その提案手法のネガティブな副作用の可能性についてもレビューすべきという提言が過去にあり、賛否両論を呼んでいたらしい (後述)
        - 背景としてはプライバシーの侵害 (主に AI 系かな) や民主主義の後退 (主に SNS 系かな) といった無視できない副作用が実際に起っていること受けて、コンピュータサイエンス業界としてもちゃんと対応しないとということらしい
        - "move fast and break things" を突き詰めると最終的に個々に到達する構造になっている。日本だと「許可を得るな謝罪せよ」は今どき許されないよねという雰囲気は一部であったと思われるが、アメリカほど具体的なリスクとまではまだなっておらず、緊急度も違いそうだと思った
    - 結構丁寧な・慎重な書き方でそうした意見にも取り組んでいくとのことだった
    - そもそもその提言が 2018 年だったり、提言をした FCA はもう解散していたり、冒頭でチャーチルの言葉を引用して正当性の論拠とするような書きぶりなどを見ると、実際のニュアンスはわからないがなかなかなるほどという感じがした
- [The ethics of computer science: this researcher has a controversial proposal](https://www.nature.com/articles/d41586-018-05791-w)
    - 上記の 2018 年当時のインタビュー
    - 確かに少し攻撃的だったのかもしれない
        - 新薬で 1000 人が救われたけど副作用で 500 人が犠牲になったときに前者だけをアピールしいないでしょとか、タバコやオイル産業みたいになりたいのとか、孫に誇れる仕事をしたいでしょとか
- <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B08VNF8481/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">クララとお日さま</a>
    - AI の一人称視点で話が進むので、叙述トリック的な、先入観を逆手に取るような違和感・不穏さが少しずつ出てくる
    - その違和感って本当はどういうことなんだろうという、謎駆動で話が進んでいくので、個人的にはとても読みやすかった
    - その謎 (違和感、不穏さ) の出し方が本当に巧みで、流石としか言いようがなかった
    - あとはクララが最も呪術的なことをしているという対比的な構造も面白かった

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B08VNF8481/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/51C2GuSq3BL.jpg" alt="クララとお日さま" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B08VNF8481/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">クララとお日さま</a></div><div class="amazlet-detail">カズオ イシグロ  (著), 土屋 政雄 (翻訳)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B08VNF8481/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
