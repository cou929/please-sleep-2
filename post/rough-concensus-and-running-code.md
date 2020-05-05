{"title":"Rough Concensus and Running Code、IETF での合意形成について調べたこと","date":"2020-05-05T15:02:00+09:00","tags":["management"]}

[IETF](https://www.ietf.org/) には `rough consensus and running code` というモットーがあるらしい。

> rough consensus and running code
>
> [IETF \| Running Code](https://www.ietf.org/how/runningcode/)

<div></div>

> We reject: kings, presidents and voting.
> We believe in: rough consensus and running code.
>
> [https://www\.ietf\.org/proceedings/24\.pdf](https://www.ietf.org/proceedings/24.pdf)

きっかけは、仕事で技術的な判断のファシリテーションがうまくいかず困っていて、標準策定をする組織のベストプラクティスが参考になるかもと思ったこと。彼らは構造上トップダウンで方針決定をできないはずだし、加入時の選考がないのでメンバーの多様性もより大きいはず。そんな環境での合意形成は普通の企業よりも難しそうなので、参考になるのではと考えた。

軽く調べたところ、手頃なドキュメントが複数あり、また上記のモットーがキャッチーで気になったので IETF について調べることにした。

このようなモチベーションのため、この記事は IETF の標準化プロセスの詳細を説明するものではないし、特定の技術標準についてでもないので、注意してください。

## Rough Concensus

まずは Rough Concensus について、[RFC 2418 \- IETF Working Group Guidelines and Procedures](https://tools.ietf.org/html/rfc2418) には次のように説明されている。

> Working groups make decisions through a "rough consensus" process. IETF consensus does not require that all participants agree although this is, of course, preferred.

<div></div>

> In general, the dominant view of the working group shall prevail. (However, it must be noted that "dominance" is not to be determined on the basis of volume or persistence, but rather a more general sense of agreement.)

<div></div>

> Consensus can be determined by a show of hands, humming, or any other means on which the WG agrees (by rough consensus, of course).

- ワーキンググループは、全会一致を理想としながらも、`rough concensus` が取れた場合は議論を先にすすめることにしている
- 単純な投票数ではなく、大多数が合意できていると考えられているかどうか (`genearl sense of agreement`)
- `rough concensus` に達したかどうかは議長が判断する

客観的な方法ではなく、あえて主観的に判断しているということらしい。

[RFC 7282 \- On Consensus and Humming in the IETF](https://tools.ietf.org/html/rfc7282) ではこの考え方についてもっと詳しく説明している。

ここでは投票ベースとコンセンサスベースという切り口で議論の進め方を分類している。IETF はコンセンサスベースを採用している。

- 投票ベース
    - 挙手なり投票なりで多数決をとる
    - わかりやすいし機械的に進めやすい
- コンセンサスベース
    - 参加者間での合意ができたら進む
    - IETF では 100% の全会一致ではなく、それよりもラフな合意で良いとしている

なぜコンセンサスベースなのかというと、アウトプットの品質を上げるためには、投票ベースの進め方にはデメリットが多かったようだ。

標準化の議論ではいかに情報量を増やすかが重要となる。例えばちゃんとコーナーケースまで考慮されているかなど、具体的な課題が発見できればアウトプットの品質が向上する。課題の発見には、その指摘が多数派からか少数派からかは関係ない。投票制では少数派の指摘が取りこぼされてしまうし、単に賛成者がおおいからといって良いものができるわけではない。後から見落としが見つかると大変だし、多少議論が長引いても先にリスクを潰せたほうがよい。

> Lack of disagreement is more important than agreement.

一方で 100% 全会一致を目指すと、議論が進捗しづらくなる。例えば消費電力が低く高速な CPU は無いように、技術的なトレードオフはそもそも避けようがない。そうでなくても妨害的に反対する人が一人でもいると、全く進まなくなってしまうという弱さにもなる。

このようなジレンマへの対応が Rough Concensus というコンセプト。たとえ反対者が反対し続けたとしても、反対の根拠になっている「どのようなケースで問題になるか」が十分に検討されていれば先に進む。

大事なのは反対という声ではなく、それが指摘している「問題」に集中すること。技術的にどのような課題があるのかを明確にし、その検討を行うこと。そのうえでアウトプットを変更することもあるし、トレードオフとして対応しないこともある。その後の同様の異論があったとしても、その課題は検討済みとして前に進むことができる。

> Rough consensus is achieved when all issues are addressed, but not necessarily accommodated.

投票ベースでもコンセンサスベースでも、議長が参加者に挙手等 (IETF では [ハミング](https://tools.ietf.org/html/rfc2418#section-3.3)) を求めるという手続きが似ている。そのためコンセンサスの確認を投票のように扱わないように注意する。投票ベースでは議論の最後に投票し議論を決定づけるが、コンセンサスベースでは議論の出発点として確認し、そこから論点を深堀りするイメージだ。決定するためのツールではなく、論点を出し尽くすためのツールという感覚だと思う。

> Humming should be the start of a conversation, not the end

有名なスローガンとして次のものがあるらしい。

> We reject: kings, presidents and voting.
> We believe in: rough consensus and running code.

ここでの `kings, presidents` は個人に権限を集中させてのトップダウン、voting は投票ベースの多数決で、それよりもコンセンサスと実装を重視していると読み取れる。

IETF という歴史のある組織で、客観的・手続き的な投票ではなく、議長という特定の権限の主観によるコンセンサス方式をとっているのは興味深い。適切な (ラフ) コンセンサスの形成は議長の技量に依りそうだし、主観なので再現性も担保しづらそう。それでもこの方法が生き残っているのは、これがうまく働くからだと思う。理論上は投票ベースでも適切な議論はできるかもしれないけれど、実際には多数派の意見だけが取り沙汰されてしまい議論を深めるのが難しいのだと思う。よいアウトプットが出すことが最大の目的というのも納得度が高かった。

## And Running Code

ラフコンセンサスに比べ、こちらはわかりやすい。(ので、さらっとした説明しかない)

[IETF \| The Tao of IETF: A Novice's Guide to the Internet Engineering Task Force](https://www.ietf.org/about/participate/tao/#contribute) には、実装による仕様へのフィードバックが、重要な貢献のひとつとして説明されている。実装により仕様の正しさ、見落としていたバグがわかる。

個人的な経験から、設計と実装を相互に行き来したほうがよりよい設計ができると感じている。標準化においても同じというのは理解しやすい。

> You can help the development of protocols before they become standards by implementing (but not deploying) from I-Ds to ensure that the authors have done a good job. If you find errors or omissions, offer improvements based on your implementation experience.

また実際にその仕様が使われていればいるほど、その仕様の重要性があがる。`"running code wins"` というフレーズがわかりやすい。

> One of the oft-quoted tenets of the IETF is "running code wins", so you can help support the standards you want to become more widespread by creating more running code.

## そのほかの工夫

Rough Concensus は独自性の高いコンセプトだけど、それ以外にもいろいろなプロジェクト運営の工夫がされている。
内容自体は一般的かもしれないが、個人的に参考になったものをピックアップした。

### プロジェクトの目的・スコープ・マイルストーンを明確化すること

ワーキンググループを開始するときには Charter (憲章) を必ず作成する。憲章はプロジェクトの目的、スコープ、マイルストーンを定めているもの。

議論を効率的にすすめるために、これらを定義しておくことは重要だと思う。

- 議論が脇道にそれてしまうのを防ぐ (防ぐ際の根拠になる) 
- 目標を明確にすることで判断基準も明確になり、対立を抑制できる
- マイルストーンにより何から手を付ければよいのかが明確になる

IETF では憲章がないとワーキンググループををもそも開始できないし、憲章を変更の際は承認が必要。割と厳重なプロセスをしいていて、憲章を重く扱っているのが伺える。

- 憲章のサンプル
    - [IPv6 over the TSCH mode of IEEE 802\.15\.4e \(6tisch\) \-](https://datatracker.ietf.org/wg/6tisch/about/)

### チーム内に一定のヒエラルキーを作ること

ワーキンググループはフラットな個人の集まりではなく、役割分担があり、役割に応じて一定の権限が与えられている。例えばこんな役割がある。

- 議長
    - コンセンサスの判断をする権限がある
    - 他の役割のアサインをする権限がある
    - 最終的なアウトプットに責任を持つ
- エディタ
    - ドラフト (標準化作業の成果物にあたる) を書く権限がある
    - 議論内容をドラフトに適切に反映させる責任がある

それぞれ役職名を見ると権限が大きそうだが、彼らがやりたい方針をトップダウンで強要するほどの権限は無い (少なくとも直接的には)。あくまで判断の主体はワーキンググループ全体にあり、議論を効率化・活性化させるためにこうした役割を設置しているように見える。例えば全員がフラットな権限しかなければ、(たとえラフだったとしても) コンセンサスの判断ができず、議論が進まないか、あるいは投票制で行こうという結論になってしまうかもしれない。

### 議論のアウトプットの形が明確なこと

ワーキンググループはドラフトを書くことがアウトプットだ。ドラフトとは RFC になる前の文章で、議論はドラフトの内容に対して行い、議論の結果はドラフトに反映される。

成果物のフォーマットが明確で、かつそれが一箇所にあることが強いと思う。例えば普段の開発で仕様を決める際、BTS のコメントで議論し Fix したらプルリクエストを作成という流れがよくあると思う。こういうチケットをあとから見ると、現在の仕様 (あるいは結局どういう結論になったのか) が分かりづらい。標準策定作業とは前提が違うので、普段の開発で RFC のような文章をメンテするのは現実的ではないかもしれないが、それでも理想形として理解できる。

### 議論前後の資料を充実させること

ドラフトの周辺にもいろいろなドキュメントが作られる。

- アジェンダ
    - 事前に必ず読んでおかないといけない資料がリンクされている
    - そのために十分な準備期間をとっている
        - 2週間前までにアジェンダを submit しないと会議が開催できない、など
    - サンプル
        - [Agenda IETF106: 6tisch](https://datatracker.ietf.org/doc/agenda-106-6tisch/)
- 議事録
    - 議事録を出すことは議長の責任として明記されている
        - それだけ重要視しているということだと思う
    - 議論結果をドラフトに反映させる
        - 「結局どうなってるんだっけ」をなくし、同じ議論が繰り返されることを防ぐのが大事
    - サンプル
        - [Minutes IETF106: 6tisch](https://datatracker.ietf.org/doc/minutes-106-6tisch/)

事前・事後のドキュメントをしっかり整備することで、関係ない議論がされていまったり、同じ議論が繰り返されたりすることを防いでいる。それぞれ明確にプロセス化されていて、それだけ重要なことなんだなと思う。

### オープンで公平な原則に基づいて運営すること

当たり前かもしれないが大切なこととして、オープン性と公平性に基づいた運営をする。

> WG should be open to any participant 

> Qualifications for a WG chair: 
> You have to balance progress and fairness 

逆を考えるとわかりやすくて、一部のひとだけでクローズに進められた意思決定はその外の人には納得しづらい。判断が公平でないと少数派が指摘をするインセンティブが働かない。結果として (本来的な) 大多数のコンセンサスに達するまでに余計時間がかかってしまう。

## 参考

- [RFC 2418 \- IETF Working Group Guidelines and Procedures](https://tools.ietf.org/html/rfc2418)
    - ワーキンググループ運営の全般的な解説
- [RFC 7282 \- On Consensus and Humming in the IETF](https://tools.ietf.org/html/rfc7282)
    - Rough Consensus の考え方について説明している
    - [RFC 7282 – On Consensus and Humming in the IETF \(2014\) \| Hacker News](https://news.ycombinator.com/item?id=19606657)
        - ハミングって何? というのは Hacker News でもコメントされていた
- [IETF \| The Tao of IETF: A Novice's Guide to the Internet Engineering Task Force](https://www.ietf.org/about/participate/tao/)
    - RFC ではなくサイト上にあった記事
    - 全体の概要をつかむのにとてもよいまとめだった。とりあえずこれだけでも読めば良いと思う
    - [IETFのタオ：初心者のためのインターネット技術タスクフォースガイド](https://www6.ietf.org/tao-translated-ja.html)
        - 日本語訳もある
- [IETF \| Working Group leadership training](https://www.ietf.org/about/participate/tutorials/process/working-group-leadership-training/)
    - いつかのセッションでの発表スライド
    - ワーキンググループの議長・エディタ向けのトレーニング資料
- [IETF Datatracker](https://datatracker.ietf.org/)
    - 各種ドキュメントはここから辿れる。実例を見たいときに。

<iframe style="width:120px;height:240px;" marginwidth="0" marginheight="0" scrolling="no" frameborder="0" src="//rcm-fe.amazon-adsystem.com/e/cm?lt1=_blank&bc1=000000&IS2=1&bg1=FFFFFF&fc1=000000&lc1=0000FF&t=pleasesleep-22&language=ja_JP&o=9&p=8&l=as4&m=amazon&f=ifr&ref=as_ss_li_til&asins=B00JO8MR28&linkId=ee357a09cfb84b81f19184fee92e8199"></iframe>
