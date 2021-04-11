{"title":"Rails on GKE なプロジェクトにキャッチアップするためにやったこと","date":"2021-04-11T17:30:00+09:00","tags":["ruby", "k8s", "book"]}

gke 上で rails アプリケーションを動かしているプロジェクトに、新たに参加することになった際に読んだ本などのまとめ。

## はじめる前の知識

いわゆる Web サービスの開発運用経験はあるが、Ruby, Rails, k8s を仕事で使うのは初めて。GCP も BigQuery とそれに連動する gcs, stackdriver を多少業務で使ったことがあったが、それ以外は初めて。

前提としてこの記事は、こういう経験値の視点からとなっている。

## 入門書を多読するアプローチ

今回はこの記事で紹介されているアプローチを試してみた。

[効率的に新しいことを学ぶ方法 – 栗林健太郎](https://kentarokuribayashi.com/journal/2020/07/31/2020-07-31-003804)

> 1. 新しく学びたい領域について、入門書を5冊〜10冊ほど買う（技術書なら1万〜2万ぐらいか）
> 2. ひとつひとつを精読するのではなく、ただ文字を追うぐらいの感じでわからないところは読み流しつつ、読み切る
> 3. 1冊1時間と時間を決めて、必ず時間を守る
本を読んでいる時にコードを書いたりコマンドを実行したりなど、試したりすることはしない。ただ読むだけ
> 4. それを、買った冊数分（5冊〜10冊）くりかえす。そうすれば5時間〜10時間、すなわち1日で学習できる
> 5. 上記により、その領域の入門的な全体像は頭の中に入るので、あとは簡単なタスクについて手を動かしながら、公式ドキュメントなどを読みながら自分で進める

今回の自分の状況はこれらにまさに当てはまる。最速でこのプロジェクトの固有の知識（製品のドメイン知識だったり、コードベースの特有の運用方法だったり）を身につけられる状態になりたかった。ただ、Ruby / Rails にしても k8s にしても、知っていれば効率的だけど初見では読み解きづらい構文や、そもそも各リソースのざっくりとした意味や概念などは、最初に理解しておかないとさすがに非効率。そのため最初に一気に基礎知識をインプットしてしまうのは良いアプローチに思えた。

## 感想

全体としてはよかった。たぶん理想的には入門書5冊を急いで読むのと、同じ時間で良書を1冊じっくり読むのとはそれほど変わらないと思う。ただ実際には "わからないところは飛ばす" "手は動かさない" "とにかくざっと読み通す" という方針のおかげで、妙に引っかかって時間を使うことがなかったのが
とても良かった。yak shaving に時間を取られすぎずに、本来やるべきこと素早く行ける気がした。

ただ、1冊1時間では全然読めなかった。その速度で文を眺めても頭に一切残らないので、それぞれ数時間かけて読むことになった。他の本で出たトピックの繰り返しや、現在では明らかに古くなっている部分は意識的に飛ばしたが、それでも1時間には及ばない。たぶんこのアプローチには、初見だと良書と悪書の区別がつかないリスクを分散する利点もあると思う。それでもある程度冊数を絞った方が自分には向いていたかもしれない。

# 実際に読んだ本やドキュメント

## Ruby と Rails

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873113679/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://images-na.ssl-images-amazon.com/images/I/51BI+oUtJCL._SX371_BO1,204,203,200_.jpg" alt="初めてのRuby" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873113679/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">初めてのRuby</a></div><div class="amazlet-detail">Yugui  (著)<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873113679/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

過去に購入して積読になっていた。古いのでどうかなと思ったけど、今回読んだものの中で一番良かった。ニーズに一番フィットしていた。

他言語経験者に向けの解説で、必要十分・簡潔に説明されていて素晴らしかった。発行からだいぶ時間が経っていることが今の自分では評価しづらいが、その点を無視すればお薦めできる本だと思う。

Go のときは Tour of go をざっと眺めると、とりあえずコードを何となく読めるようになって、良い体験だった。これと近い。

[Ruby on Rails チュートリアル：プロダクト開発の０→１を学ぼう](https://railstutorial.jp/)

次に Rails チュートリアルを読んだ。初めての Ruby とはうってかわって、プログラミング自体が初めての人をターゲットにしているようで、git の使い方なども説明されていた。自分には冗長だった（このチュートリアルが悪いわけでは無い）

Rails の規約を知りたかったので、それとは無関係なところと、ビューとフロントまわりは全部飛ばして読んだ。

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B077Q8BXHC/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/51tyOVaM2HL.jpg" alt="プロを目指す人のためのRuby入門 言語仕様からテスト駆動開発・デバッグ技法まで Software Design plus" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B077Q8BXHC/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">プロを目指す人のためのRuby入門 言語仕様からテスト駆動開発・デバッグ技法まで Software Design plus</a></div><div class="amazlet-detail">伊藤 淳一  (著)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B077Q8BXHC/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

初めての Ruby に比べもっと初心者向けに書かれているようで、Ruby に限らないソフトウェア開発一般の概念も丁寧に説明している印象。説明が丁寧でわかりやすくて感心した。初心者の人にも、そうで無い人にも最初の一冊としておすすめしやすい。初めての Ruby との差分を意識しながら読んだ。

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00P0UR1RU/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/51YrQYJdmPL.jpg" alt="パーフェクトRuby on Rails" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00P0UR1RU/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">パーフェクトRuby on Rails</a></div><div class="amazlet-detail">すがわらまさのり  (著), 前島真一  (著), 近藤宇智朗  (著), & 1 その他  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00P0UR1RU/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

[改訂版が 2020 年に出ている](http://www.amazon.co.jp/exec/obidos/ASIN/B08D3DW7LP/pleasesleep-22/ref=nosim/)ので、今はそちらを読んだ方が良い。流し読み目的だったのと初版が手に入りやすいところにあったので今回はとりあえずこちらにした。Rails チュートリアルとの差分と今は古くなっていそうな部分を意識しつつ飛ばせるところは飛ばした。

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B071K5WM6P/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/517to0qVtmL.jpg" alt="改訂2版 パーフェクトRuby" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B071K5WM6P/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">改訂2版 パーフェクトRuby</a></div><div class="amazlet-detail">Rubyサポーターズ (著)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B071K5WM6P/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

構文は前2冊でOKということにして、それ以外のところに目を通すようにした。ざっくりどういう標準ライブラリがあるのかなどが俯瞰できた気がする。メタプログラミングのトピックも全部飛ばして、必要になってからかえってくることにした。

# k8s と GCP

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B08PNMRXKN/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/41g+F7WohJL.jpg" alt="イラストでわかるDockerとKubernetes Software Design plus" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B08PNMRXKN/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">イラストでわかるDockerとKubernetes Software Design plus</a></div><div class="amazlet-detail">徳永 航平  (著)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B08PNMRXKN/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

概念部分の説明に振り切っているので、今回の自分のニーズにかなり合っていた。タイトルの通り図解が多く、実際の yaml を読み始めてこれなんだっけとなった際に、さっと読み返せて便利だった。短いので流し読みを意識しないでもすぐに読めるし、最初の一冊としておすすめできると思う。

[Kubernetes ドキュメントの Concepts の章](https://kubernetes.io/docs/concepts/)

ドキュメントのコンセプトの部分だけをざっと目を通した。イラストでわかるDockerとKubernetes の補完としてのイメージ。細かいところは深入りせず、そういうトピックがあるんだなということだけ覚えておく感じで読み進めた。

[Learn Kubernetes Basics \| Kubernetes](https://kubernetes.io/docs/tutorials/kubernetes-basics/)

こちらは手を動かしながらやってみた。ブラウザで動作するインタラクティブな環境でハンズオン的なことができ、そこにも感心した。

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B08FZX8PYW/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/51C+pft8SJL.jpg" alt="Kubernetes完全ガイド 第2版 impress top gearシリーズ" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B08FZX8PYW/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Kubernetes完全ガイド 第2版 impress top gearシリーズ</a></div><div class="amazlet-detail">青山真也  (著)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B08FZX8PYW/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

こちらは通読していない。わからないことのリファレンス的に読んでいる。

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B07S1LG1Y1/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/51o4lhZcgXL.jpg" alt="GCPの教科書" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B07S1LG1Y1/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">GCPの教科書</a></div><div class="amazlet-detail">吉積礼敏  (著)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B07S1LG1Y1/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B088LZGPM5/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/51xvCD7nX8L.jpg" alt="GCPの教科書II 【コンテナ開発編】" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B088LZGPM5/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">GCPの教科書II 【コンテナ開発編】</a></div><div class="amazlet-detail">クラウドエース株式会社 (著), 飯島宏太  (著), 高木亮太郎 (著), & 2 その他  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B088LZGPM5/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

こちらは初心者向け系。とりあえずどういうキーワードが出てくるのかざっと見るくらいにはなったと思う。

# プロジェクト固有の知識

ここまでは一般的な知識で、ここからはプロジェクト固有の知識。例えばマーケットや製品のドメイン知識や、そのシステムやチームの現状や運用方針など。むしろここからが本番だと思う。

実業務をやりながら、コードの読み書き、データの試行錯誤、issue のフィードを追いながらドキュメントを読んでいくなど、基本的にはじっくりやらないといけない部分だと思う。

そんな中で、今のプロジェクトで素敵だなと思ったのは、ログが充実している点。アプリケーションログからインフラ、アプリの計測ツール経由のイベント、RDBMS の中身などなどが全て BigQuery に保存されている。例えばアプリケーションログは構造化して入っているので、非常にクエリしやすい。一箇所に集まることで、他のソースのデータとも Join できるのも良い。これがシステムの動きをキャッチアップするにあたって、とても助けになっている。

リーダーが最初にこれらのログを活用するようなタスクを振ってくれた。レビューでも丁寧にデータの位置や実際に使っているクエリを提示してくれて、お陰で最初のうちにログの見方を身につけることができたと思う。これがキャッチアップの効率にかなり効いている感覚があり、とてもありがたかった。
