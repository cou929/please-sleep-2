{"title":"最近読んだもの 5","date":"2021-06-20T23:00:00+09:00","tags":["readings"]}

## 記事

- [Things I Wished More Developers Knew About Databases \| by Jaana Dogan \| Medium](https://rakyll.medium.com/things-i-wished-more-developers-knew-about-databases-2d0178464f78)
    - データストアに関するいろいろなトピック。各論が列挙されている感じだけどいい内容で面白かった
    - トランザクション分離レベルの話とか、分散システムでネットワークやタイムスタンプが信用ならない話とか、そういう [データ志向アプリケーションデザイン](http://www.amazon.co.jp/exec/obidos/ASIN/4873118700/pleasesleep-22/ref=nosim/) で出てきそうなちょっと理論寄りの話から、トランザクションがネストしないようトランザクションを開始するレイヤーを決めておくとか、性能検証のコツとか、実運用の話まであり、筆者の高い力量が垣間見える
    - 面白かった
        - トランザクション分離レベルは SQL 標準で定義されている 4 つよりも更に細分化できるよという話
            - データ志向アプリケーションデザイン筆者の [プロジェクト](https://github.com/ept/hermitage) も登場していた
        - オートインクリメントな pk がベストではないケース
            - 分散環境にあるデータストアの場合 (分散システムでシーケンスを生成するのは大変)
            - パーティションアルゴリズムに pk を使うデータストアで、シーケンシャルなことで偏りが生まれる場合
            - シーケンシャルな番号以外で pk になりうるカラムがある場合 (計算の分が無駄になる)
        - 性能検証はシビアなワークロードでやったほうが良いという話
            - 例えば、50M 行あるテーブルに関連レコードを伴う INSERT、平均 500 人友達が居るグラフで友達の友達までを取得、500 フォローしているタイムラインの上位 100 レコードの取得、など

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873118700/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://images-na.ssl-images-amazon.com/images/I/51T+k4VRzpL._SX389_BO1,204,203,200_.jpg" alt="データ指向アプリケーションデザイン ―信頼性、拡張性、保守性の高い分散システム設計の原理" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873118700/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">データ指向アプリケーションデザイン ―信頼性、拡張性、保守性の高い分散システム設計の原理</a></div><div class="amazlet-detail">Martin Kleppmann (著), 斉藤 太郎 (監修), 玉川 竜司 (翻訳)<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873118700/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

- [Internal Tech Emails on Twitter: "Mark Zuckerberg: "speed and strategy" February 14, 2008 https://t\.co/zx6m54tWr6" / Twitter](https://mobile.twitter.com/TechEmails/status/1404489425372000258)
    - Mark Zuckerberg が Fb 社内に送った、事業戦略を説明するメール
    - こんなに明快な説明はあんまり見たことなくて、素直にすごいと思った
    - 出どころは、Facebook と Six4Three という会社間の訴訟の際の証拠物件として提出された Fb 社内の内部文章らしい
        - [Thousands of Facebook Internal Documents, Emails Made Public in Leak](https://www.businessinsider.com/facebook-internal-documents-executive-emails-published-six4three-court-leak-2019-11)
            - [こちらがその生の pdf](https://dataviz.nbcnews.com/projects/20191104-facebook-leaked-documents/assets/facebook-sealed-exhibits.pdf) (627 MB あるので注意)
            - 529 ページ目にある
    - 出処となった事件はアレだけど、このメール単体だとめちゃくちゃいい内容
- [Google Testing Blog: How Much Testing is Enough?](https://testing.googleblog.com/2021/06/how-much-testing-is-enough.html?m=1)
    - テストどこまでやる？できてる？を考える上での分類、枠組み
    - とてもよい
- [Data denormalization is broken \| Hacker Noon](https://hackernoon.com/data-denormalization-is-broken-7b697352f405)
    - パフォーマンスのために非正規化する話から始まるが発展がすごい
    - 非正規化したカラムの更新方法として pure function 的に副作用なしに再計算するか、差分更新的に前の値を使って更新していくか
    - こうしたニーズを完全に満たすデータストアはまだない
        - 限定的なものなら materialized view が該当
- [How Lowe’s leverages Google SRE practices \| Google Cloud Blog](https://cloud.google.com/blog/products/devops-sre/how-lowes-leverages-google-sre-practices)
    - `reducing toil` の活動例として、1 番にアラートのトリアージを機械学習でやる、というのが出てきてびっくりした
        - いきなり高度な気がしたけど、そうでもないのかな。ちょっとした分類タスクを機械学習でやるクラはもう結構当たり前になってる？
    - Black Monday とかの対応の際に Google の CRE チームにコンサルしてもらうのは面白そう
        - GCP への寄稿なのでそこは割り引く必要はあるけど
- [Implementing ChatOps into our Incident Management Procedure — Infrastructure](https://shopify.engineering/implementing-chatops-into-our-incident-management-procedure)
    - Shopify の slack にある障害対応チャンネルのスクショがのっている。だいたい 400 人強のユーザーが入ったいた。
        - 2018 年の記事だけど、プロダクトに関わる人が大体このくらい居たと言える
    - その規模でこのくらいの作業プロセスの標準化、明文化がされているというイメージができた
    - どういうタイミングで誰にどうコミュニケーションとるかは障害対応時に大事だけど、それは確かに俗人的な知識になりうるし、この規模だとそこを標準化、自動化してもペイするんだな。とか
- [The MTTR that matters \| FireHydrant](https://firehydrant.io/blog/the-mttr-that-matters)
    - MTTR (Mean time to recover) だと、障害の状況に応じて分散が大きすぎて意味のある数値になりづらい
    - 発生から検知までの時間とか、解決から振り返り実施までの時間とかを測った方がよい
- [Toward Vagrant 3\.0](https://www.hashicorp.com/blog/toward-vagrant-3-0)
    - vagrant を ruby から go にポートするらしい
    - しかも後方互換性を保ちながらアーキテクチャもクライアントサーバ型に変える
    - 大変だろうけど楽しみ
- [Nat Friedman on Twitter: "GitHub processes 2\.8 billion API requests per day, peaking at 55k rps\. Lots of busy bots\. 🤖" / Twitter](https://mobile.twitter.com/natfriedman/status/1404835709278580739)
    - GitHub のピーク時 api トラフィックは 55k rps
    - [Rails, memcached, & MySQL running on Kubernetes and served via haproxy and GLB](https://mobile.twitter.com/AaronBBrown777/status/1404984775291592709) という構成
    - これとは別に[同規模の git オペレーション](https://mobile.twitter.com/natfriedman/status/1405295566217564166)もある
    - 参考値としてなんかで使えそう
