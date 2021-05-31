{"title":"最近読んだもの 2","date":"2021-05-31T19:30:00+09:00","tags":["radings"]}

## 記事

- [A robust distributed locking algorithm based on Google Cloud Storage – Joyful Bikeshedding](https://www.joyfulbikeshedding.com/blog/2021-05-19-robust-distributed-locking-algorithm-based-on-google-cloud-storage.html)
    - ちょっとしたロックを実装するのに gcs にロックオブジェクトを置くという例
    - ここで言う「ちょっとしたロック」のユースケースは、例えば CICD プロセスが複数走っている中でワークフローの本体処理自体を更新したい (何も実行されていないときに更新したい) といったもの
    - こういうユースケースに対してかなりしっかり実装されていて (TTL とかロックの延長とか etc)、全然手軽ではなかった
- [Biases in AI Systems \- ACM Queue](https://queue.acm.org/detail.cfm?ref=rss&id=3466134)
    - 門外漢なのでへーという感想しかないが
    - 教師データの偏りなど種々の理由で、AI がいわゆる差別的な挙動をしてしまう問題について
    - 未解決の問題だが、問題の要因を分類・整理するアプローチは良いなと思った (問題解決一般として)
- [A Second Conversation with Werner Vogels \- ACM Queue](https://queue.acm.org/detail.cfm?id=3434573)
    - Amazon CTO Werner Vogels のインタビュー
    - Amazon 規模になると十年以上ずっと事業もシステムもスケールし続けたわけで、そうなるとシステムをいかに進化させやすくするかが最重要になる
    - また一般的に API のインタフェースは一度定義すると容易に変えられない
    - この 2 点から導き出されるのは、システムをシンプルにつくり拡張していくことがいかに大事かということ
    - 単に機能不足の製品を作るのではなく、本質的にユーザーが求めるとこをだけを実現するそぎ落とした mvp からはじめること、それを突き詰めることが大事になる
    - というようなことが書かれていた気がする
- [John Gall \(author\) \- Wikipedia](https://en.wikipedia.org/wiki/John_Gall_(author))
    - Gall's law という経験則「複雑なシステムはシンプルなシステムからしか作れない。最初から複雑なシステムは作れない」
    - 何かの本で見かけて出典がわからず一年くらい探していた。やっと見つけられて嬉しい
- [よくないエンジニアリングカルチャー \- Google Docs](https://docs.google.com/document/d/1v_EoYtnMY9-jBuGME6uQMNIQ6XIIkrO8uPO2Lc9M1zs/edit)
    - `職人信仰` はほんとにそうで、不具合や障害が最初から起きないのが地味だけど一番すごい。自分もできていない
    - `レビュー遅れ` の不味そうなコードをレビューをしっかりやって何も起こらないよりは、ノールックでマージして壊れて直したほうがまだマシと言う考え方は全く新鮮だった
        - 後者はその分試行回数が増やせて学習機会が増えてよいということ
- [How to Monitor Sidekiq Process Uptime in Rails Apps](https://pawelurbanek.com/rails-sidekiq-monitoring)
    - 定期的にある URL にリクエスト、リクエストが来なかったら通知してくれる監視サービスを [cron monitoring](https://www.google.com/search?q=cron+monitoring) というらしい
    - 面白いアイデア
- [How to look at the stack with gdb](https://jvns.ca/blog/2021/05/17/how-to-look-at-the-stack-in-gdb/)
    - gdb を使ってスタックとヒープにそれぞれ格納したデータがメモリ上どうなっているか、ステップバイステップで見ていく
    - `gets` であえてバッファオーバーフローを起こしたり
- [Understanding Ruby Method Lookup \- Honeybadger Developer Blog](https://www.honeybadger.io/blog/ruby-method-lookup/)
    - ruby のメソッド呼び出しの解決ロジックの解説
    - 特異メソッド => mix in された module => インスタンスメソッド => 親クラスを同じ順で => 最終的に Object, Kernel, BasicObject へ
        - prepend は include と似ているが、メソッド解決チェーン? の手前に挿入される点が違う

## 読み終わった本

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B07GBS9XRW/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/41G3-v8o+UL.jpg" alt="舞踏会へ向かう三人の農夫　上 (河出文庫)" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B07GBS9XRW/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">舞踏会へ向かう三人の農夫　上 (河出文庫)</a></div><div class="amazlet-detail">リチャード・パワーズ (著), 柴田元幸 (翻訳)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B07GBS9XRW/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

- まだ上巻を読み終えたところだけど面白い
- 最初は真面目に内容を理解しようとしてなかなかページが進まなかったけど、アメリカ？の文化資本に深く触れていないとよくわからないところが多かった。諦めて流し読みにするといい感じ

