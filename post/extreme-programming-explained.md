{"title":"エクストリームプログラミング","date":"2018-11-11T00:13:48+09:00","tags":["book"]}

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012UWOLOQ/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://images-fe.ssl-images-amazon.com/images/I/51QDy-s%2BFFL._SL160_.jpg" alt="エクストリームプログラミング" style="border: none;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012UWOLOQ/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">エクストリームプログラミング</a><div class="amazlet-powered-date" style="font-size:80%;margin-top:5px;line-height:120%">posted with <a href="http://www.amazlet.com/" title="amazlet" target="_blank">amazlet</a> at 18.11.10</div></div><div class="amazlet-detail">オーム社 (2017-07-14)<br />売り上げランキング: 30,962<br /></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012UWOLOQ/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

ちょっと読みづらいけれど、良い本だった。

ソフトウェア開発手法の、方法論でなく原則からよく説明されている。良い開発スタイル・チームに共通する点をこの本では「価値」と呼んでいて、そこから具体的にどうすればよいか (「プラクティス」) の説明を展開している。ここでの「プラクティス」は今のなっては定着しきっていたり古臭いものもあるが、この本自体はアジャイル開発の基礎となっているものらしく、その原則となる「価値」は今でも間違っていない。なぜそれが大事なのかという根本まで立ち返って考えることは、技術を低レイヤーまで深掘りすることと同じで、しっかりと問題を理解しいろいろなケースに対応するために重要だ。開発フローやチームビルディングの方法論を学び・実践していくなかでこの本を読むと、今までの知識をすっきりと整理できて良いと思う。

この本で説明されている「価値」は、例えば「シンプルな方がいい」とか「適切な頻度のフィードバックを得る」などといった、抽象度の高く、このままなにかに使えるものではない。その分、この「価値」に反するフローや習慣があれば、「なにかおかしい」と感じるための指針として使えると思った。具体的な方法論自体は時と場合によって変わっていくので、その軸としての「価値」を抑えておきたい。

内容が古いせいか、あるいは「パターン」という仰々しい書き方のせいか、読みやすい本ではない。それでも、開発手法やチームを改善したい人にとっては、読む価値のある本だと思う。

以下は読書メモ。

### 価値

- コミュニケーション
	- `いちばん大事なのはコミュニケーション` という認識をもつこと
	- 情報共有も含める
- シンプリシティ
	- システムも方法もシンプルな方がいい
- フィードバック
	- 適切な頻度のフィードバック
		- 従来よりも頻度を上げることを意識することのほうが多そう
		- ノイジーなフィードバックが多すぎてもいけない
	- フィードバックを元にした改善活動も含める
- 勇気
	- 何かをやるときに奮い立たせるために入れている項目かな
	- これだけではだめ
- 尊敬
	- 基本姿勢として
	- 採用時にも

### プラクティス

- 全員同席 (Sit Together)
	- 同じ場所で働く
- チーム全体 (Whole Team)
	- クロスファンクションで 1 チームとする。帰属意識を持つところまで
- 情報満載のワークスペース (Informative Workspace)
	- 物理的なかんばんやダッシュボードなどをオフィスに作る
- いきいきとした仕事 (Energized Work)
	- 長時間労働をしない
- ペアプログラミング (Pair Programming)
- ストーリー (Stories)
	- ユーザーに見える機能単位で計画する
- 週次サイクル (Weekly Cycle)
	- 振り返り、計画のリズム
- 四半期サイクル (Quarterly Cycle)
	- 振り返り、計画のリズム
- ゆとり (Slack)
	- 計画にはゆとりをもたせておく
- 10 分ビルド (Tne-Minute Build)
	- ビルド、テストにかかる時間は短いほどよい
- 継続的インテグレーション (Continuous Integration)
- テストファーストプログラミング (Test-First Programming)
- インクリメンタルな設計 (Incremental Design)
	- システムの設計に毎日手を入れること

導出プラクティス

- 本物の顧客参加 (Real Customer Involvement)
	- 本当のユーザーをチームに入れる
		- あるいはドッグフーディング?
- インクリメンタルなデプロイ (Incremental Deployment)
	- 一気に大きなリリースをしない
- チームの継続 (Team Continuity)
	- 小さいチームを保つ
		- 適切な分割の話?
- チームの縮小 (Shrinking Teams)
	- 効率をよくするということ?
- 根本原因分析 (Root-Cause Analysis)
	- 問題が起こったら深く追うこと
- コードの共有 (Shared Code)
	- チームの誰でもどこにでもアクセス・修正できる
- コードとテスト (Code and Tests)
	- コードを正として、コードからドキュメントを生成する
- 単一のコードベース (Single Code Base)
	- 複数のブランチをメンテするようなことを避ける
- デイリーデプロイ (Daily Deployment)
	- 頻繁にリリースする
- 交渉によるスコープ契約 (Negotiated Scope Contract)
	- 期間・費用・品質を固定し、スコープを調整すること
- 利用都度課金 (Pay-Per-Use)
	- 従量課金のプロダクト設計にするとフィードバックが得やすいということ

### チームの役割

- テスター
	- つまらないミスはプログラマーが自分で見つけるべき
	- テスターは通常見つけられないようなケースを発見するなど
	- 自動化をプログラマーと一緒に推進する
- インタラクションデザイナー
	- ストーリーを書いたり、リリース後の状況から新たなストーリーの機会を探したりする
- アーキテクト
	- 大規模なリファクタの調査・実施や、パフォーマンス・チューニングなど
	- 小さく継続的に実施していく
- プロジェクトマネージャ
	- チーム内のコミュニケーションの円滑化、顧客・サプライヤー・その他のチーム外組織とのコミュニケーション調整
	- 進捗把握、報告、ファシリテーション
- プロダクトマネージャ
	- ストーリーを書いたり、四半期のロードマップを引いたり、週次のストーリーを選択したりする
- 経営幹部
	- 方針を示す、チームに的確な説明を求める
	- 改善の監視、促進、円滑化
- テクニカルライター
	- フィーチャーのフィードバックを早期に早期に提供、ユーザーと密接な関係性を築く
	- CS・CE に近そう
- ユーザー
	- ストーリー記述の支援、専門領域の意思決定
- プログラマー
	- 見積もり、ストーリーをタスクに分解、テスト、実装
- 人事
	- 人事評価と雇用

### 気になる本

テイラー主義はソフトウェア開発にはマッチしないと紹介されていた

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/447800983X/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://images-fe.ssl-images-amazon.com/images/I/41uevpgSWUL._SL160_.jpg" alt="|新訳|科学的管理法" style="border: none;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/447800983X/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">|新訳|科学的管理法</a><div class="amazlet-powered-date" style="font-size:80%;margin-top:5px;line-height:120%">posted with <a href="http://www.amazlet.com/" title="amazlet" target="_blank">amazlet</a> at 18.11.10</div></div><div class="amazlet-detail">フレデリック W.テイラー <br />ダイヤモンド社 <br />売り上げランキング: 29,773<br /></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/447800983X/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
