{"title":"最近読んだもの 4","date":"2021-06-13T23:00:00+09:00","tags":["readings"]}

## 記事

- [Summary of June 8 outage \| Fastly](https://www.fastly.com/blog/summary-of-june-8-outage)
    - 原因特定から復旧までが早い。ある程度予測できたバグだったのかな。全く予期もしないバグだったら流石にこの速さで対応できない気がしたので
        - そうだとして、それでこれほどの影響範囲の広さなのもすごい
    - 障害検知が早いし、外部コミュニケーションもしっかりしてるのは、日頃の積み重ねからだろうし流石としか
- [Event Sourcing pattern \- Cloud Design Patterns \| Microsoft Docs](https://docs.microsoft.com/en-us/azure/architecture/patterns/event-sourcing)
    - めちゃめちゃよくまとまってた
    - Cloud architecture のデザインパターンとして他にもたくさんページがあった。どれもこのクオリティで記載されてるんだとしたら、それらも読む価値ある
    - 会社の GitHub Issue に参考文献としてリンクされていたので読んだもの
- [Minimizing ossification risk is everyone’s responsibility \| Fastly](https://www.fastly.com/blog/minimizing-ossification-risk-is-everyones-responsibility)
    - ある前提が変わらないと思って書いたコードが、その前提が変わってしまい、その後のシステムの進化の妨げになる現象
        - 極端な例だと 2000 年問題。「年」は必ず "19" から始まるという誤った前提が変わった
    - Web は漸進的に変わっていくので、fastly みたいなプレイヤーだとこの問題は大変そうだなと思う
    - ossification とは [骨化、(思想・信仰などの)硬直化、固定化](https://ejje.weblio.jp/content/ossification) という意味らしい
- [Kubernetes: A Pod's Life](https://www.openshift.com/blog/kubernetes-pods-life)
    - pod のライフサイクルとその周辺の tips
    - とてもわかりやすかった
- [An incomplete list of skills senior engineers need, beyond coding \| by Camille Fournier \| Jun, 2021 \| Medium](https://skamille.medium.com/an-incomplete-list-of-skills-senior-engineers-need-beyond-coding-8ed4a521b29f)
    - 自分の過去を振り返ってしまった
    - `How to lead a project even though you don’t manage any of the people working on the project`
        - これが一番できなかった気がする
    - `How to give up your baby, that project that you built into something great, so you can do something else`
        - baby = 実子という意味ではなく、(あるいはそれも含めた) 自分自身で主体的に取り組むプロジェクトという意味で捉えた。これもできなかった
    - このへんは笑った
        - `How to indulge a senior manager who wants to talk about technical stuff that they don’t really understand, without rolling your eyes or making them feel stupid`
        - `How to explain a technical concept behind closed doors to a senior person too embarrassed to openly admit that they don’t understand it`
- [How Netflix uses eBPF flow logs at scale for network insight \| by Netflix Technology Blog \| Jun, 2021 \| Netflix TechBlog](https://netflixtechblog.com/how-netflix-uses-ebpf-flow-logs-at-scale-for-network-insight-e3ea997dca96)
    - すごそうだけど具体的に何をやっているのかよくわからない（主に自分の前提知識不足による）
    - ちなみに Brendan Gregg は最近 [Nitfilx に転職してオーストラリアに移住した](http://www.brendangregg.com/blog/2021-05-29/moving-to-australia.html) らしい
- [Building a Healthy On\-Call Culture \| SoundCloud Backstage Blog](https://developers.soundcloud.com/blog/building-a-healthy-on-call-culture)
    - Soundcloud でのオンコール当番事例
    - 一人当たり月に三日ほどが良いという。定量的な目安は初めてみた気がする
- [Fuzzing is Beta Ready \- The Go Blog](https://blog.golang.org/fuzz-beta)
    - 自分が fuzzing という概念を知ったのは go の標準に入れようという go blog の記事だった気がする
- [Understanding RBS, Ruby's new Type Annotation System \- Honeybadger Developer Blog](https://www.honeybadger.io/blog/ruby-rbs-type-annotation/)
    - 公式にこういう試みがされているとは知らなかった。面白い
- [\[MONITORING\] How to build your monitoring dashboards? \| by Daniel Moldovan \| May, 2021 \| Medium](https://dmoldovan.medium.com/monitoring-how-to-build-your-monitoring-dashboards-e11f89918dd1)
    - 監視ダッシュボードの細かい tips
    - 一番上に一番抽象度の高い、一目でわかるメトリクスを大きく出しておくのはなるほどと思った
- [Incident Management vs\. Incident Response \- What's the Difference? \| Rootly](https://rootly.io/blog/incident-management-vs-incident-response-what-s-the-difference)
    - 障害の特定、解消などを行う活動と、情報集約や外部コミュニケーション含めた非技術だがビジネスに重要な活動について、それぞれ Incident Response / Management と呼ばれるとここで説明している
    - ただ定義も呼び名も世間ではばらばららしい
- [Please don't count outages \(or SEVs, or whatever\)](https://rachelbythebay.com/w/2021/06/01/count/)
    - 測りやすいからといって変な kpi を設定するとダメになるケース
    - それによって促進されるべき行動の逆になっちゃうというか、心理的安全性が下がるというか
    - 行動とは関係が薄くて全体の傾向、ユーザーがハッピーかどうかがわかる指標にしないと危ない
    - 以前読んだ測りすぎという本を思い出した

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B07RL7RGRW/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/314RErfLhuL.jpg" alt="測りすぎ――なぜパフォーマンス評価は失敗するのか？" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B07RL7RGRW/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">測りすぎ――なぜパフォーマンス評価は失敗するのか？</a></div><div class="amazlet-detail">ジェリー・Z・ミュラー (著), 松本裕 (翻訳)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B07RL7RGRW/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
