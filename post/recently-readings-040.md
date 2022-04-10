{"title":"最近読んだもの 40 - Vitess 導入事例、CloudSQL Insights など","date":"2022-04-10T23:30:00+09:00","tags":["readings"]}

- [Scaling Datastores at Slack with Vitess \- Slack Engineering](https://slack.engineering/scaling-datastores-at-slack-with-vitess/)
	- Slack の Vitess 導入事例
	- MySQL の既存の資産を活かしながらプロダクト開発の速度を落とさずにスケールさせるには、やはり NewSQL への移行よりもこうしたソリューションの方が現実的なんだなと思った
- [Square \| Cloud Native Computing Foundation](https://www.cncf.io/case-studies/square/)
	- Square の Vitess 導入事例
	- `Only 5% of the system had to be changed` いい数字
	- こういう企業たちが upstream に還元してくれているのがとても心強い
- [Partitioning GitHub’s relational databases to handle scale \| The GitHub Blog](https://github.blog/2021-09-27-partitioning-githubs-relational-databases-scale/)
	- 前にも読んだことがあったけれど、データの cutover の部分を再読
	- テーブルを分割する際に別クラスタにダウンタイムなしででーたを移動させるプロセス
	- Vitess を利用する方式と、導入して間もないときは独自のスクリプトでの移行も行ったらしい
	- 後者は primary を一度リードオンリーにし、レプリカへのレプリケーション完了を確認し、レプリカをマスターにプロモートし、SQLProxy の向き先を変更するというもの
	- github の規模でもオフピークに実施すれば十分少ないエラーで切り替えられれたらしい
		- primary で 50,000 queries/s くらい？
- [MySQL database performance monitoring \| Google Cloud Blog](https://cloud.google.com/blog/products/databases/mysql-database-performance-monitoring)
	- CloudSQL insights が MySQL にも来るらしい
	- 早速申し込んだ
- [Selecting a Startup Stack for Scale](https://www.cockroachlabs.com/blog/selecting-startup-stack-for-scale/)
	- ポジショントークぽい部分もありつつ、スタートアップにおけるスケール対応のタイミングとして、実際に問題になってからやるといいというのは割と聞く話かなと思うけど、理想的には問題になる前に必要十分な時間をとって対応した方が良いと言っているのは、確かにその通りだなと思った
	- 技術選定に関しても、現状のコードベースへの変更が小さく、将来の拡張性を損なわす、学習コストが高すぎず、採用コストが高すぎないものが（あくまで理想だが）良いので、そこを出来るだけ目指すべきというのも、判断軸の整理になって良かった
- [GitHub Availability Report: March 2022 \| The GitHub Blog](https://github.blog/2022-04-06-github-availability-report-march-2022/)
	- mysql1 の負荷対応を引き続き地道にやっているらしい
	- 今回は actions のテーブルを functional sharding しているらしい
	- また特定の機能にメンテナンスウィンドウを設けて、高負荷時に止めたりしていたらしい
		- ユーザーとしては気づかなかった
- [Understanding Software Dynamics](http://www.amazon.co.jp/exec/obidos/ASIN/0137589735/pleasesleep-22/ref=nosim/) 9 章まで
	- observability の章に突入
	- ログの出し方やサマリーの仕方など、納得感がある

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/0137589735/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://images-na.ssl-images-amazon.com/images/I/51h17DJx7aL._SX382_BO1,204,203,200_.jpg" alt="Understanding Software Dynamics (Addison-Wesley Professional Computing Series)" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/0137589735/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Understanding Software Dynamics (Addison-Wesley Professional Computing Series)</a></div><div class="amazlet-detail">英語版  Richard Sites (著)<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/0137589735/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
