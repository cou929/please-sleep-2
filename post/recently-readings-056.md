{"title":"最近読んだもの 56 - MySQL の通信を HTTP/3 化、現代のデータベースプロダクト など","date":"2023-01-09T23:20:00+09:00","tags":["readings"]}

- [Faster MySQL with HTTP/3](https://planetscale.com/blog/faster-mysql-with-http3)
    - MySQL のクライアント - サーバ間通信を MySQL のバイナリプロトコルから HTTP/3 に変えてみてベンチをとったところ速度の改善が見られた
    - High Latency のケース (手元の PC からの接続といったケース) はまだわかるが、Low Latency のケース (同一 DC 内でのアプリケーションサーバと DB サーバとの通信といったケース) でも高速化していたのが興味深かった
    - ちなみに今回の HTTP を話す MySQL サーバ側の実装は Planet Scale のプロダクトのため必要で実装されたものらしい
- [Building a database in the 2020s \- me\.0xffff\.me](https://me.0xffff.me/build-database-in-2020s.html)
    - TiDB 創業者による 2020 の記事で、今データベースシステム (プロダクト) を一から作る場合に求められる要件や考慮点など
    - 2 年経っているが良い記事で、データベースシステムというだけでなく、開発者向けの XXX aa a Service を開発する際の勘所が包括的に簡潔にまとまっている
    - computing と storage を分離させる設計 (Aurora や Alloy DB のような設計) を更に発展させると、データベース管理システムのいろいろなコンポーネント (例えばスキーマ変更、データ圧縮、ロギング、etc...) も分離できるようになっていくという話が面白かった
- [The End of Programming \| January 2023 \| Communications of the ACM](https://cacm.acm.org/magazines/2023/1/267976-the-end-of-programming/fulltext)
    - 昨今の AI の進展を受けて、現在のようなプログラミング・CS はもう終わるという話
    - 大きい変化なのはそのとおりだけど、ちょっと感傷的すぎるというか、一旦落ち着いてほしい感じはある
- [H12 \- Request Timeout in Ruby \(MRI\) \| Heroku Dev Center](https://devcenter.heroku.com/articles/h12-request-timeout-in-ruby-mri)
    - 「rack-timeout によるタイムアウトが日に 10 - 20 件以上起こると、レースコンディションが発生している可能性が高い」という経験則が参考になる
- [The Oldest Bug In Ruby \- Why Rack::Timeout Might Hose your Server](https://www.schneems.com/2017/02/21/the-oldest-bug-in-ruby-why-racktimeout-might-hose-your-server/)
    - こちらも rack-timeout の危険性について
- [Why Ruby’s Timeout is dangerous \(and Thread\.raise is terrifying\)](https://jvns.ca/blog/2015/11/27/why-rubys-timeout-is-dangerous-and-thread-dot-raise-is-terrifying/)
    - 同上で Thread.raise 起因のレースコンディションについて
- [Kubernetesで実践するクラウドネイティブDevOps](http://www.amazon.co.jp/exec/obidos/ASIN/4873119014/pleasesleep-22/ref=nosim/)
    - 読み始めたばかりだが 1 章のクラウド => DevOps => コンテナ => Kubernetes という潮流の説明がすっきりしていてわかりやすかった

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873119014/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/51wecBhtIOL._SX389_BO1,204,203,200_.jpg" alt="Kubernetesで実践するクラウドネイティブDevOps" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873119014/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Kubernetesで実践するクラウドネイティブDevOps</a></div><div class="amazlet-detail">John Arundel  (著), Justin Domingus (著), 須田 一輝 (監修), 渡邉 了介 (翻訳)<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873119014/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
