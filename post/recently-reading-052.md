{"title":"最近読んだもの 52 -  TCP Syn Queue と Accept Queue の実装、誰のために法は生まれた など","date":"2022-09-19T23:30:00+09:00","tags":["readings"]}

- [TCP Socket Listen: A Tale of Two Queues](http://arthurchiao.art/blog/tcp-listen-a-tale-of-two-queues/)
    - Linux の TCP Syn queue と Accept Queue について
    - Syn queue は実際には hash + queue len で、また Syn cookie を有効化すると実質上限は無くなる
        - `SYN_RECV` が `ESTABLISHED` になるのはクライアントからの ACK を受け取った順なので、直列のキューである必要はなく、ハッシュでよい (Syn cookie オフの場合)
        - tcp_max_syn_backlog が上限値 (Syn cookie オフの場合)
        - Syn queue の状態を見る直接的な方法は無いが、`sudo netstat -antp | grep SYN_RECV` `ss -n state syn-recv sport :80` などで現在の SYN_RECV 件数はわかる
        - `netstat -s | grep -i listen` でオーバーフローの発生状況がわかる
    - Accept queue は `listen(2)` の backlog オプションの値と somaxconn の小さい方が上限になる
        - `ss -ntl` の `Recv-Q` `Send-Q` で状況がわかる
            - それぞれキューの現状と最大の値
            - LISTEN 状態以外のソケットではこれらのカラムの意味が変わるので注意
    - コードの解説もあり
- [One million queries per second with MySQL](https://planetscale.com/blog/one-million-queries-per-second-with-mysql)
    - Vitess による horizontal sharding で高い QPS を実現
    - ベンチマーク付き
- [How we store and process millions of orders daily](https://engineering.grab.com/how-we-store-millions-orders)
    - grab のデータストアの設計
    - ワークロードを OLTP と OLAP に分類
    - OLTP 用には DynemoDB
    - OLAP 用には kafka を用いたパイプラインを構築
- [GitHub Availability Report: August 2022 \| The GitHub Blog](https://github.blog/2022-09-07-github-availability-report-august-2022/)
    - 今月の Availability Report
- [Vitess \| A database clustering system for horizontal scaling of MySQL](https://vitess.io/blog/2021-12-16-rails-that-scales/)
    - Vitess の Rails 対応状況
    - Rails Guide のクエリが問題なく通ることを確認しているらしい
- [誰のために法は生まれた](http://www.amazon.co.jp/exec/obidos/ASIN/B07J54DKK9/pleasesleep-22/ref=nosim/)
    - (広義の) 組織の都合で個人か犠牲になるケースを防ぐため
    - という主張を始め、実力行使は一旦ブロック必要など、興味深いコンセプトが色々登場して面白かった
    - また題材となっている古典の作品はかなり面白そうで、読んで・観てみたくなった
    - ただ客観的な根拠が無いのに主張が断定的だったり、そもそも論理の展開がよくわからなかったりなど、ひっかかる部分が多く素直には楽しめなかった
        - 恐らく学生 (中高生?) 向けの口頭の講義の文字起こしであること、題材の作品を自分が読んでいないことが原因の一部だとは思う
    - 同じテーマを扱っている講義録ではない書籍があるならそれも読んでみたい

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B07J54DKK9/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/41Xdr-WjACL.jpg" alt="誰のために法は生まれた" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B07J54DKK9/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">誰のために法は生まれた</a></div><div class="amazlet-detail">by 木庭顕  (著)  Format: Kindle Edition<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B07J54DKK9/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
