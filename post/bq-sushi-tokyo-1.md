{"title":"bq_sushi tokyo #1","date":"2015-04-24T23:30:29+09:00","tags":["conference"]}

[#bq_sushi tokyo #1 - connpass](http://bq-sushi.connpass.com/event/13219/)

## BigQueryを支える技術 + Ask Me Anything

[BigQueryクエリの処理の流れ // Speaker Deck](https://speakerdeck.com/googlecloudjapan/bigquerykuerifalsechu-li-falseliu-re)

途中から参加

- キャッシュ
  - データ更新時間と参照テーブルの SHA1 ハッシュ
  - NOW() など値が変化する関数を使っている、固定の出力テーブルを指定、参照テーブルがストリーミングバッファを使用していると、キャッシュされない
  - 結果は SHA1 のテーブル名になる
- Dremel クエリの開始
  - Spanner -> BQ API サーバ -> Dremel gateway -> [dremel clusters]
- bq のレプリケーション
  - 地理的に離れた複数のゾーン間でレプリケーション
    - 可用性とディザスタリカバリ
  - dc 内でのレプリケーション
    - 永続性の向上、リード性能の向上、リードソロモンに似たエンコーディングを使用
- streaming insert
  - streaming insert されたデータは bigtable 上にバッファワーカーによって colossus に定期的に保存
  - dremel query 実行時は bigtable 上のバッファも読み込み
  - バッファの内容は durable だが、大きな障害時には一時的に読み込みできなくなる
- ストレージ階層
  - colossus メタデータ
    - ディレクトリ構造
    - ACLs
  - colossus ファイル
    - ほぼフラットな構造
    - 暗号化
  - 分散ストレージ
    - ディスクエンコード (RAID 的なもの)
    - クラスタ内でのレプリケーション

質問タイム

- streaming insert から SELECT できるまでの時間はどれくらいかかるか
  - 0
  - streaming api で 200 が返ればすぐに使える (すぐに bigtable に入るので)
  - bigtable が調子が悪い時は streaming api がエラーを返す
- クエリを投げて webhook で結果が返ってくるような機能の実装予定は
  - すぐには無いけどリクエストは多い
- 普通の DB のように index について考慮する必要はあるか
  - 必要ない
- UTC 以外のタイムゾーンのサポート予定は?
  - すぐには無い
- ユーザ定義関数で新たにできることとユースケースが知りたい
  - sql で書くことが難しい処理
  - 地理情報の取り扱いなど
  - cloud dataflow と組み合わせて複雑な処理も記述できる
- google cloud logging から bq へのインサートについて
  - cloud logging -> bq へは streaming insert している
  - よって quota なども同じになる

## Google Cloud Dataflow解説

[Google Cloud Dataflow — Google Cloud Platform](https://cloud.google.com/dataflow/)

- google ビッグデータ処理の歴史
  - gfs, mapreduce, big tabale, dremel, pregel, flume, colossus, spanner, millwheel
- cloud dataflow とは
  - 並列化されたデータ処理パイプラインを作るための SDK 軍
  - 並列化されたデータ処理をするためのマネージドサービス
- できること
  - 移動、フィルタ、加工、整形、集約、バッチ処理、ストリーム処理、組み合わせ、外部連携、シミュレーション
- メリット
  - 関数型プログラミングモデル
  - バッチ処理とストリーム処理を統合
  - マネージドサービス
  - 実行時間が短い
- スケジュール
  - いまはベータ版
  - この先オープン
- イメージ
  - メソッドチェーン的にデータ処理のパイプラインを記述できる
  - `p.begin().apply(read()).apply(count())...; p.run()` みたいな
- さまざまな runner
  - direct runner
    - ローカルでインメモリで
  - cloud dataflow service runner
    - フルマネージドで
  - その他。spark とか
- バッチ処理とストリーム処理
  - cloud pub/sub でストリーム読み書きする場合
  - さきほどのパイプラインの記述の read/write を pubsub に返るだけ。処理を使いまわす
  - time window 関数で一部の時間の処理にすることも
- データ連携
  - gcs, pub/sub, bq など
  - カスタム記述で任意のデータソースも (まだバッチのみ)
- 今後
  - sdk の python サポート
  - sdk の機能追加

質問タイム

- google io で紹介されていた gui のモニタリング画面は、今のベータ時点で使えるのか
  - もう使えます
  - io 時点より改善している
- サポート言語の予定
  - まず python

## Dive into Google Cloud Dataflow Java SDK and Google BigQuery

bq の事例紹介

- 使い方
  - resource monitoring, developer activity log, application logs, end-user's access logs
- excel + bq で pos データの分析
  - もともとは rdb を使っていた
  - 結果のファイルを excel に出して渡す。非エンジニアがそれを分析・可視化などを行う
  - 分析にかかる時間が 1/12 くらいになった
  - ランニングコストも 95% カット
- bq の cons
  - 安定性
  - レイテンシや quota

dataflow の sdk code reading

- dataflow は社長が io に行って感化されて使わされた
- 足りないもの
  - java 以外の sdk がまだない
  - カスタムの streaming input がまだない
- sdk は open source
  - [GoogleCloudPlatform/DataflowJavaSDK](https://github.com/GoogleCloudPlatform/DataflowJavaSDK)
- [DataflowJavaSDK Weekly](http://dataflow-java-sdk-weekly.hatenablog.com/)

質問タイム

- なぜ pub/sub を使わずにカスタム streaming insert を自前実装しないといけないとか
  - サービス的に認証まわりの問題がある
  - pub/sub を介さず直接 insert したい (1 層少なくなるので)

## Drillを読んでみよう

- apache drill
  - google の dremel 論文が元になっている

## BigQuery in Windows - ログコレクタからLINQによる検索まで

- streaming insert が quota に達した場合
  - storage 経由のインサートに一部切り替えた
  - 件数が増えると金額も変わるので、検討したほうがいい
- テーブルは月で分割
  - 最大テーブルは 750 億件
  - テーブル名に YYYYMMDD いれて TABLE_DATE_RANGE
  - デコレータは必須
  - 垂直分割は避ける
- [neuecc/LINQ-to-BigQuery](https://github.com/neuecc/LINQ-to-BigQuery)

## SmartNewsのBigQuery事例

- embulk
  - マスターデータの import などで使っている。mysql -> bq
- [Cloud Business Intelligence | Chartio](https://chartio.com/)
  - 可視化ツール
  - mysql と bq を join したりできるらしい

## EmbulkのGCS/BQプラグインについて

- guess と preview がいい
  - config の yaml を作るのに便利

## 感想

- [google cloud dataflow](https://cloud.google.com/dataflow/) は面白そうなので調べる
- 「streaming insert してから SELECT できるようになるまでの delay は 0 だ」と言っていたが、以前この delay が 1 時間を超えることがあった
  - そのときは google 社内のサポートに issue を発行して対応してもらった
- 発表者の nagachika さんという方は、追うと決めた OSS は、コミットを [毎日読んで日記に書く](http://dataflow-java-sdk-weekly.hatenablog.com/) ようにしているらしく、良い手法だなと思った
- [Chartio](https://chartio.com/) は調べる

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/487311716X/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="http://ecx.images-amazon.com/images/I/51LYuI72mPL._SL160_.jpg" alt="Google BigQuery" style="border: none;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/487311716X/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Google BigQuery</a><div class="amazlet-powered-date" style="font-size:80%;margin-top:5px;line-height:120%">posted with <a href="http://www.amazlet.com/" title="amazlet" target="_blank">amazlet</a> at 15.04.24</div></div><div class="amazlet-detail">Jordan Tigani Siddartha Naidu <br />オライリージャパン <br />売り上げランキング: 9,024<br /></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/487311716X/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
