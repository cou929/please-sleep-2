{"title":"最近読んだもの 30","date":"2022-01-16T23:00:00+09:00","tags":["readings"]}

## 記事

- [The James Webb Space Telescope — making 300 points of failure reliable \| by Robert Barron \| Jan, 2022 \| Medium](https://flyingbarron.medium.com/the-james-webb-space-telescope-making-300-points-of-failure-reliable-db669810a9d8)
    - DevOps/SRE のソフトウェアエンジニアが James Webb 宇宙望遠鏡の今回の打ち上げについて記載したもの
    - 通常このような宇宙ミッションの成功は (1) 冗長性で担保 (同じものを複数飛ばす)、(2) 信頼性で担保、(3) 修復性で担保 (あとから宇宙飛行士が修正・拡張する) するらしいが、今回は予算や技術制約のため (2) しか実施できなかったらしい
        - 予算 = 9B USD
    - デプロイの工程には 300 の SPOF があるという非常に複雑なもの。実環境でのテストができずに (2) を達成するためにひたすら地上でのテストが行われた
    - なんでそんな複雑なのかというと、デザイン上代替案よりそっちのがマシだったから
        - 既存のスペースシャトルに搭載できないサイズの機器なので、それを搭載できるサイズのスペースシャトルを新規開発するか、部品を打ち上げて宇宙空間で組み立てるかの二択で、後者のほうがリスクが低いと判断された
    - 以降の記事でさらに詳細が書かれるらしいので楽しみ
- [The Mightiest Monolith\. Shuttle To SRE — STS Lesson 1 \| by Robert Barron \| IBM Garage \| Medium](https://medium.com/ibm-garage/the-mightiest-monolith-7aa67ab8d84f)
    - 上記と同じシリーズの過去記事
    - スペースシャトルの歴史からの学び
    - スペースシャトルはそれまでのミッションと異なり、再利用可能な汎用の宇宙船が志向された
    - そのためトライアンドエラーができず、目的も特化していないのでフルスタックな、モノリス的な開発フローになった
    - それまでの使い捨ての船と異なり再利用必須なので、大気圏再突入に耐える設計が複雑になったのと、それをメンテナンスし次回のフライトを行う工数もかさんだ
    - 結果としてフライトの頻度も費用も以前の使い捨て船よりも悪化したらしい
- [4th of July Fireworks — A Balanced Action Plan\. \| by Robert Barron \| IBM Garage \| Medium](https://medium.com/ibm-garage/4th-of-july-fireworks-solving-technical-debt-5d54bbea833)
    - 同じシリーズ
    - コロンビア号の事故後の対策と SRE が行うポストモーテムとを照らし合わせる
    - 対策は以下のタイプに分類でき、それぞれ組み合わせるのが大事
        - observability の増加
        - 問題の切り分けを速くする改善
        - 問題自体の改修
        - 再発防止のための根本的な対策
- [Testing and Proofs of Concept — the Shuttle Approach and Landing Tests \| by Robert Barron \| Medium](https://flyingbarron.medium.com/shuttle-approach-and-landing-tests-b2efe79aa927)
    - 同じシリーズ。シャトル編はこれで最後
    - エンタープライズの失敗について
    - 試作機として学びを蓄積する意味において正しかったが、周りから「完成品」とみなされてしまっていたのが問題だった
        - 少しづつ複雑にしていきながら、プロダクションに近づけながら、試作品をトライアンドエラーしていくのが良いのと、そうであると関係者の認識が揃っていること
    - 途中で出てくる ibm のドキュメントが良さげだった
        - [IBM's principles of chaos engineering \- IBM Cloud Architecture Center](https://www.ibm.com/cloud/architecture/architecture/practices/chaos-engineering-principles/)
    - このシリーズは面白いんだけど、出典を書いてくれないかな
        - 宇宙開発の歴史が興味深いのが十分伝わったので、本で読みたい
- [Observability in Chaos](https://world.hey.com/minglei/observability-in-chaos-88d43d84)
    - 現代のソフトウェアは複数のマイクロサービスの相互作用のため、不確実性が高いカオスで、簡単に観測することができない
    - ある切り口で可視化し人間がパターン認識できるようになると有用なので、今後はそういうツールがいろいろと生まれてきそう
- [Designing Instagram \- High Scalability \-](http://highscalability.com/blog/2022/1/11/designing-instagram.html)
    - Designing シリーズの instagram 編
    - ストリーム処理やグラフ DB、列指向 DB が組み合わされていてお手本のような大規模システム設計な気がした
