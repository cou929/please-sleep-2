{"title":"\"はじめよう！ 要件定義 ～ビギナーからベテランまで\" 読書メモ","date":"2018-01-14T22:19:27+09:00","tags":["book"]}

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00WHUP7UE/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://images-fe.ssl-images-amazon.com/images/I/51QVR3y3sOL._SL160_.jpg" alt="はじめよう！ 要件定義 ～ビギナーからベテランまで" style="border: none;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00WHUP7UE/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">はじめよう！ 要件定義 ～ビギナーからベテランまで</a><div class="amazlet-powered-date" style="font-size:80%;margin-top:5px;line-height:120%">posted with <a href="http://www.amazlet.com/" title="amazlet" target="_blank">amazlet</a> at 18.01.14</div></div><div class="amazlet-detail">技術評論社 (2015-04-24)<br />売り上げランキング: 9,497<br /></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00WHUP7UE/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

読んだのはかなり前なんだけど、読書メモが出てきたのであげておく。

- 要件定義の進め方、成果物 (ドキュメント) について説明された本
- UML ほど形式的ではなくよりカジュアルで、実用的でバランスの良い内容だと思う。
- そのまま適用するのではなく、必要に応じて必要な部分だけ使えば良い
- 受託開発だけでなく、Web 系の自社プロダクト開発にも通じる話し。実体験としてそうだったのだが、例えばエンジニアとしてプロダクトオーナーのやりたいことを定義する場面では有用だった。
- 自分一人で作るものと作り方を決めて進めるフェーズでは不要かもしれないが、プロダクトオーナーや同じチームのエンジニアが居るケースなど、チームで開発を行う場合には役に立つと思う。
- 当然、新規開発の場合だけではなく、既存システムへの機能追加でも使える。

Web 系であれば一定規模以上のシステムを扱うエンジニアは一読の価値があると思う。量的にも必要最低限でさくっと読めるのでおすすめ。

### 読書メモ

#### 成果物

要件定義に必要なドキュメント

- 企画書
    - 一目で理解できるもの
    - プロジェクト名
    - なぜ (目的)
    - 何を (目的達成のために作るもの、それを利用する人、利用者が得られる便益)
    - どのように (体制、期限)
- 全体像 (オーバービュー)
    - 何をに着目した図。プロジェクトのゴールをぱっと見でわかるようにする
    - システムを中心におき周りにユーザーを記載する。
    - ユーザーごとにできることを記載する
    - 連携システムとその内容もかく
    - システム管理者もユーザーとして書く
- 利用する実装技術
- 実現したいこと一覧 (要求一覧)
    - プロダクトオーナーからの要求のリスト。やりたいことの一覧
- 行動シナリオ
    - システムが利用されるタイミング、利用される理由、利用されることでユーザーが達成する仕事を明らかにし、仕様検討に役立てる
    - 「仕事」を最小粒度にし、発生タイミング、前提、内容、成果物、仕事の依存関係を図にする
    - システム仕様検討とは切り離し、あくまでユースケースに注目して作成する
    - UML 業務フロー図の簡易版。WBS、BPR も関連ワード
- 行動シナリオ一覧
    - 発注者やチームメンバーにわかるよう一覧化する。その際構造化して一目でわかるようにする
- 概念データモデル
    - 主要なテーブルだけの図。詳細まで書いてはいけない

要件定義成果物

- ワークセット一覧
    - 一つの仕事を実現するシステムの機能単位
    - この単位で後の分析を行う
    - 一覧にして分かりやすくする
    - 構造化する
- ラフイメージまたはモックアップ
    - ワークセットごとにペーパープロトタイピングをする
    - データ項目、操作項目、レイアウトを決める
- 項目の説明
    - モックのワークセットごとの詳細
    - 表示項目、入力、操作のリスト
- 画面遷移図
    - 詳細をつなげる
    - 操作内容、行われる処理を記載する
    - トリガーがユーザー操作でない処理も記載する
    - UI からは直接現れないデータを導出する
    - 項目を詳細までここで決めきる
- 機能の入出力定義、処理定義
    - 画面遷移図にひもづいて昨日の詳細を決めていく
    - 入出力を定義する
    - 処理を定義する
    - 入れ子にして構造化する
- 結合 ERD
    - ワークセットごとに検討し最後にマージする

下にいくにつれて抽象度が下がる

#### その他

- 一目で伝わるサマリーが発注者とのやりとりやチーム内での共有のため重要
    - 読み解く努力を相手だけに求めない。伝える努力をするのが大切
    - そのためには単なる一覧ではなく構造化すると良い
