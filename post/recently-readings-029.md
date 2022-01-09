{"title":"最近読んだもの 29","date":"2022-01-09T23:00:00+09:00","tags":["readings"]}

## 記事

- [Does a Container Image Have an OS Inside](https://iximiuz.com/en/posts/not-every-container-has-an-operating-system-inside/)
	- コンテナの仕組みの基礎
	- コンテナは cgroups とか namespaces とかで分離されたプロセスだよという説明書
	- [scratch](https://hub.docker.com/_/scratch) という空のベースイメージがあるのは知らなかった
	- このシリーズおもしろそうなので他の記事も読んでみる
- [Designing WhatsApp - High Scalability](http://highscalability.com/blog/2022/1/3/designing-whatsapp.html)
	- WhatsApp のアーキテクチャを調査してまとめたもの（たぶん）。過不足なく網羅的で面白い
	- チャットメッセージをサーバサイドのデータストアに全件保存するのではなく、クライアントだけが持っているという構成なのがユニークなポイントと思われた（たぶん LINE も同じ設計思想？）
	- 大きくはオンラインのユーザーへのメッセージ配信を行うチャットサービスと、オフラインのユーザーようにメッセージを溜めておくトランジットサービスという二つのサービスが主な登場人物になる
	- またチャットサービスはユーザーごとにシャーディングすることでスケールできる設計になっている
	- 同じシリーズの Netflix の記事もあるようなので読んでみよう
- [Designing Netflix - High Scalability](http://highscalability.com/blog/2021/12/13/designing-netflix.html)
	- 上記と同じシリーズの Netflix 版
	- Netflix の場合このレベルのアーキテクチャデザインよりも、より実装、運用に近い部分の知見の方が面白いのかもと思った
		- 有名なカオスエンジニアリングをはじめとするマイクロサービスの設計運用とか、CDN、動画エンコーディングまわりとか
	- そのへんの資料へのポインタも記事にはあるので掘ることはできそう
- [MySQL Query Optimization: Top 3 Tips](https://blogs.oracle.com/mysql/post/mysql-query-optimization-top-3-tips)
	- MySQL のオプティマイザの tips
	- 全く同じクエリが過去にあっても毎回オプティマイザの全ての処理が行われる
		- Oracle などはそうではないらしい
	- MySQL の 8 系から[Explain の tree format](https://dev.mysql.com/doc/refman/8.0/en/explain.html) というものが追加された
		- 例えば index からカラムを取り出した後の絞り込み (Extra にかいてある内容) がツリーでわかりやすく表現されている
		- hash join の情報はこのフォーマットでしか見れないらしい
- [How to deal with MySQL deadlocks](https://www.percona.com/blog/2014/10/28/how-to-deal-with-mysql-deadlocks/)
	- show engine innodb status の LATEST DETECTED DEADLOCK の読み方と対処方法
	- トランザクションを小さくする、直列になるようにする、分離レベルを変えるなど
- [Five things you did not know about Rails transactions](https://longliveruby.com/articles/five-things-you-did-not-know-about-rails-transactions)
	- rails のトランザクションに関する小ネタ集
	- ネストしたトランザクションへの対応 (checkpoint) とか、トランザクションの状態 (`fully_committed`) など知らないことが多かった
	- ドキュメント・実装も確認したい

## 本

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B09NBZLC7J/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/41l0NwfJDDL.jpg" alt="プロジェクト・ヘイル・メアリー　上" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B09NBZLC7J/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">プロジェクト・ヘイル・メアリー　上</a></div><div class="amazlet-detail">アンディ ウィアー  (著), 小野田 和子 (翻訳)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B09NBZLC7J/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B09NBZ4Z3S/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/412N17w45GL.jpg" alt="プロジェクト・ヘイル・メアリー　下" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B09NBZ4Z3S/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">プロジェクト・ヘイル・メアリー　下</a></div><div class="amazlet-detail">アンディ ウィアー  (著), 小野田 和子 (翻訳)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B09NBZ4Z3S/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

- 面白くて一気に読んでしまった
- 火星の人 (オデッセイ) の作者の新作
- 火星のワトニーと同じく、絶望的な状況でも科学とユーモアで一歩ずつ進む爽やかな前向きさは顕在で、かつ今回はエンタメ度がパワーアップしていてよかった
- 個人的には三体 (まだ 1 巻しか読んでいない) 並に良かった
