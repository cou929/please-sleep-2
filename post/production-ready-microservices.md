{"title":"プロダクションレディマイクロサービス","date":"2020-03-13T22:53:00+09:00","tags":["book"]}

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873118158/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://images-fe.ssl-images-amazon.com/images/I/51n-OwJUG8L._SL160_.jpg" alt="プロダクションレディマイクロサービス ―運用に強い本番対応システムの実装と標準化" style="border: none;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873118158/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">プロダクションレディマイクロサービス ―運用に強い本番対応システムの実装と標準化</a><div class="amazlet-powered-date" style="font-size:80%;margin-top:5px;line-height:120%">posted with <a href="http://www.amazlet.com/" title="amazlet" target="_blank">amazlet</a> at 20.03.13</div></div><div class="amazlet-detail">Susan J. Fowler <br />オライリージャパン <br />売り上げランキング: 28,298<br /></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873118158/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

マイクロサービスを実運用するにあたって、抑えるべきポイントをカタログした本。具体的な how ではなく、何をすべきかを網羅的に書いてあるので、抜け漏れのチェックに使うと良さそう。訳が若干読みづらいがなんとかなるレベル。

読み終わって印象的だったのは、マイクロサービスだからこそなトピックは実は多くなくて、どんなアーキテクチャのシステム運用でも大切なことの割合の方が、実は多かったことだった。レビューやデプロイから、監視やドキュメンテーションまで、一見地味だけど大切なことをしっかりやっていくのが大事なんだなと改めて思った。

そんな中で、マイクロサービスゆえの難しさは以下の二点があったように思う。

- 標準化が実は重要だということ
  - マイクロサービスの良さとして、各サービスの独立性により例えば技術選定が自由にできることなどが喧伝されることがある
  - 一方で複数のサービスが連動するので、問題が発生すると可用性が乗算で下がっていく
  - そのためどこまで標準化し/どこまで自由にすることで、可用性を担保するかのデザインが重要になってくる
- 依存サービスの多さにより、問題が連鎖することへの対策
  - そのため、サーキットブレイカーや防御的キャッシュ、障害時の代替やフォールバックなどの重要性は、明らかにモノリスよりも高い

マイクロサービスでのシステム全体を次のような 4 層モデルで整理したこと。各サービスが満たすべき標準を Production-ready と名付け、Production-Readiness という標準の形で網羅的に定義したことは、今後自分が実践する場合にはとても役に立ちそうだった。

## マイクロサービスの 4 層モデル

マイクロサービスを動作させるために必要なものを次の 4 層にモデリングしている。

- レイヤ 4: マイクロサービス
  - マイクロサービスそのものと、その構成情報
- レイヤ 3: アプリケーションプラットフォーム
  - 開発環境、テスト、パッケージング、ビルド、リリースツール、デプロイパイプライン、サービスレベルの監視・ロギングなど
- レイヤ 2: 通信
  - ネットワーク、DNS、RPC、サービス検出、負荷分散など
- レイヤ 1: ハードウェア
  - 物理サーバ、データベース、OS、構成管理、ホストレベルの監視・ロギングなど

個々の定義がどうこうというよりは、このように漏れなく分類してくれたおかげで、チーム（分業）を設計する際に指針になりそうだと思った。

## Production-ready なサービスの要件

まず、標準化する理由・目標として `可用性の向上` とシンプルに定義しているのが良かった。そしてそのために必要な要件 = Production-Readiness として、ざっくり次のようなトピックが取り上げられている。

- 安定性・信頼性
  - 開発サイクル (コード管理、レビュー、デプロイ)
  - デプロイパイプライン (ステージング、カナリア、本番環境)
  - 依存関係の把握、可視化
  - ルーティング、ヘルスチェック
- スケーラビリティ・パフォーマンス
  - 量的・質的な指標の定義
  - リソースの把握
  - キャパシティプランニング
  - ボトルネックを排除
- 耐障害性・大惨事対応 (Catastrophe-Preparedness)
  - 単一障害点の除去
  - 障害シナリオの洗い出しとプロセス化
  - テストとカオスエンジニアリング
- 監視
  - ロギング
  - ダッシュボード
  - アラート
  - オンコールローテーション
- ドキュメント

これらを満たしたサービスは Production-ready な状態となる。このネーミングもわかりやすくて良いと思う。
