{"title":"コードレビューのポイント","date":"2013-10-17T00:24:54+09:00","tags":["blog"]}

すこし古いけどこちらを読んだ。

[pull request を利用した開発ワークフロー // Speaker Deck](https://speakerdeck.com/hotchpotch/pull-request-woli-yong-sitakai-fa-wakuhuro)

非常にわかりやすく読ませてもらったが、本筋とはちょっとずれるレビューの勘所の話がとても面白く参考になった。

### レビューコメントにラベルをつける

- [MUST] 問題があり、必ず治す
- [IMO] 意見、緩やかな指摘。自分ならこう書くけどどう?
- [nits] ほんの小さな指摘。インデントやタイポ

たしかに普段のコードレビューでも、特に意識せずこれらは使い分けていた。こうしてラベルとして明示することでレビュイーに意図が伝わりやすい。レビュアーもこうした観点からレビューを行える。なによりレビュアー・レビュイー双方にとって明確に、指摘事項に優先順位がつけられる。こうすることでレビューとその後のアクションをよりスムーズにすすめる手助けになっている。

いいレビューとだめなレビューというものは確実に存在していて、いいレビューは事前にバグを防いでくれたり、よりエレガントな設計にブラッシュアップしてくれたりする。一方だめなレビューは些細なことに時間がかかり、開発スピードが遅くなったり、かけた手間の割りには見返りが少なかったりする。

よくないレビューがおこってしまう原因のひとつは、指摘事項に対するコストパフォーマンスの意識が少ないことだろう。工数がかかる割にはそんなにコードが良くならない指摘はコスパが悪いといえる。コスパが悪い指摘事項でも、内容自体が間違っているわけではないので、レビュイー側も無視するわけにはいかない。後回しにするにしても客観的な理由の説明が必要だし、そのやりとりにもコストがかかる。

今回のようなラベルがあることで、レビュアーも指摘のコスト感覚を意識するようになるし、指摘事項にどう対応するのか、一定の客観的な基準が与えられることになるので対応もしやすい。

### レビューの視点

- ビジネスロジックの視点
- コードの質の視点

レビュー時の視点をこのように分類するのは新鮮で、かつ納得感があった。レビューが行われる段階や文脈によって、それぞれどの程度やるべきかが変わってくる。たとえば後述の WIP pull-req のような、実装方針の確認のための中間レビューの場合、コードの質の視点からのレビューに力を入れてもしょうがない。ロジック視点もコーナーケースの考慮などは一端おいておき、データの持ち方や処理の全体的な流れにだけ着目すべきだ。一方でリリース前の最終レビューの場合は、おおまかな流れに関してはもう問題ないはずなので、むしろ細かいロジックの整合性やパフォーマンスが悪い書き方をしていないかなど、細部にも気を配る必要がある。

### WIP の pull request

大きめの修正を行う場合、[WIP] (Work in progress) とラベルを付けたプルリクエストを送り、仕様や設計・方針について確認してもらったり議論する。一度に大きな差分を送りつけるよりもレビュアーの負担は少ないし、また設計に問題があってやり直しになった場合のロスが少ない。やはり開発に関してはいろいろな粒度で、このようなリーン的な考え方で進めていったほうが効率的だ。
