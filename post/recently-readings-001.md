{"title":"最近読んだもの 1","date":"2021-05-24T22:55:00+09:00","tags":["radings"]}

最近読んだ記事や本をログしてみます。目的は自分のモチベアップ、プレッシャー、etc のため。

初回なのでここ 1 ヶ月くらいで面白かったものを。今後は週次か隔週くらいで続けられると良いなと思っています。

## 記事

- [Why puma workers constantly hung, and how we fixed by discovering the bug of Ruby v2\.5\.8 and v2\.6\.6 \| by Yohei Yoshimuta \| Apr, 2021 \| ITNEXT](https://itnext.io/why-puma-workers-constantly-hung-and-how-we-fixed-by-discovering-the-bug-of-ruby-v2-5-8-and-v2-6-6-7fa0fd0a1958)
    - 会社の人が書いた記事。すごい。いろいろなところで取り上げられていた
- [Seeing Like an SRE: Site Reliability Engineering as High Modernism \| USENIX](https://www.usenix.org/publications/loginonline/seeing-sre-site-reliability-engineering-high-modernism)
    - 紹介されている `Seeing Like A State` という本が面白そうだった
        - [ダイアン・コイル 「ジェームズ・スコット（著）『Seeing Like A State』を再読して」（2015年9月6日） — 経済学101](https://econ101.jp/%E3%83%80%E3%82%A4%E3%82%A2%E3%83%B3%E3%83%BB%E3%82%B3%E3%82%A4%E3%83%AB-%E3%80%8C%E3%82%B8%E3%82%A7%E3%83%BC%E3%83%A0%E3%82%BA%E3%83%BB%E3%82%B9%E3%82%B3%E3%83%83%E3%83%88%EF%BC%88%E8%91%97%EF%BC%89/)
    - ハイモダニズムという考え方を SRE にあてはめてみること、techne と metis という観点での整理がとてもなるほどだった

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00D8JJYWA/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/51zy6SoQUsL.jpg" alt="Seeing Like a State: How Certain Schemes to Improve the Human Condition Have Failed (The Institution for Social and Policy St) (English Edition)" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00D8JJYWA/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Seeing Like a State: How Certain Schemes to Improve the Human Condition Have Failed (The Institution for Social and Policy St) (English Edition)</a></div><div class="amazlet-detail">英語版  James C. Scott  (著)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00D8JJYWA/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

- [What Serverless Computing Is and Should Become: The Next Phase of Cloud Computing \| May 2021 \| Communications of the ACM](https://cacm.acm.org/magazines/2021/5/252179-what-serverless-computing-is-and-should-become/fulltext)
    - いわゆる "サーバレス" は今までほぼ気にしてこなかったけど、ここでは自分が認識していたより大きなビジョンを示していて興味深い
        - クラウドコンピューティングの第一段階はオペレーションの抽象化で、これはだいぶ成熟してきた。次はプログラミングの抽象化が進むので、アプリケーションを書くエンジニアにとってのパラダイムシフトになるよという話
    - ちょうど最近アナウンスされた [AWS App Runner](https://aws.amazon.com/blogs/aws/app-runner-from-code-to-scalable-secure-web-apps/) を触ってみようという気になった。サブプロジェクトで試してみたい
- [Site Reliability Engineering for Native Mobile Apps](https://www.infoq.com/articles/site-reliability-engineering-mobile-apps/)
    - SRE の考え方をモバイルアプリに適用という話
    - 実際にはアプリ開発の "基盤" っぽい取り組みをうまく SRE という観点で整理した、という感じだけど面白かった
    - 個人的には最近ネイティブクライアント側に仕事で携わっていて、改めてその深さを実感していたところだったのでタイムリーだった
        - なんというか、メインの UI スレッドとユーザー入力がある、他のアプリと共存する世界観でどう戦うか、そのために登場するリアクティブプログラミングとか並行処理とかのトピック、が新鮮
        - サーバサイドの開発は、並行性に関する部分はおおよそサーバなり FW が抽象化しているので、プログラマーはその中でシングルスレッドの処理を書く事が多い。難しさはいかに効率よくミドルウェア・データストアを使うかとか、分散システムでどう一貫性を担保するかとかに寄る印象
    - [Behind the scenes of 1Password for Linux \| by Dave Teare \| May, 2021 \| Medium](https://dteare.medium.com/behind-the-scenes-of-1password-for-linux-d59b19143a23)
        - これなんかは、1 Password レベルになるとクライアントアプリの中だけで一定上複雑なアーキテクチャになりうるという例がみれて興味深かった
            - クライアントの中でさらにバックエンドとフロントエンドに分かれていて、バックエンドは Rust で、フロントは React で、みたいな話
            - この規模になるとこうした抽象化は必須なんだなという学び
- [SRE Case Study: Mysterious Traffic Imbalance](https://tech.ebayinc.com/engineering/sre-case-study-mysterious-traffic-imbalance/)
    - RFC 3484 によって DNS ラウンドロビンで偏りが生まれる事があるよという話。常識だったのかもしれないけど、知らなかった
- [Fast and flexible observability with canonical log lines](https://stripe.com/blog/canonical-log-lines)
    - リクエストごとにメタデータを構造化してロギングすると便利という話だけど、これに Canonical log というかっこいい名前をつけているのがうまいなと思った (流行ってはなさそうだけど)
- [How to Successfully Hand Over Systems \| SoundCloud Backstage Blog](https://developers.soundcloud.com/blog/how-to-successfully-hand-over-systems)
    - [Lightweight Architecture Decision Records](https://www.thoughtworks.com/de/radar/techniques/lightweight-architecture-decision-records) というドキュメントの書き方を初めて聞いた (流行ってはなさそうだけど)
- [Deep Dive into Database Timeouts in Rails](https://engineering.grab.com/deep-dive-into-database-timeouts-in-rails)
    - `read_timeout` オプションについて調べていて行き当たった記事だけど、ちゃんと動作検証してていいなと思った
- [Thundering herds, noisy neighbours, and retry storms \| Mads Hartmann](https://mads-hartmann.com/sre/2021/05/14/thundering-herd.html)
    - `Thundering herd` みたいな用語に operational patterns という名前をつけて集めているらしい。ほとんど知らなかった
- [Humanity wastes about 500 years per day on CAPTCHAs\. It’s time to end this madness](https://blog.cloudflare.com/introducing-cryptographic-attestation-of-personhood/)
    - 内容もよいとして、`500 human years wasted every single day — just for us to prove our humanity.` という煽り文句が流石だなと思った

## 読み終わった本

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B088BLSH9V/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/419mypQVhHL.jpg" alt="React Native　～JavaScriptによるiOS／Androidアプリ開発の実践" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B088BLSH9V/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">React Native　～JavaScriptによるiOS／Androidアプリ開発の実践</a></div><div class="amazlet-detail">髙木 健介  (著), ユタマこたろう (著), 仁田脇 理史 (著)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B088BLSH9V/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

- 比較的最近の情報を反映している RN の日本語の書籍は貴重だったと思う
- コードレベルの解説に紙面をけっこう割いていたが、もっと図などで概念の説明をしてくれたほうがありがたい

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4815607117/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://images-na.ssl-images-amazon.com/images/I/51CYaQOVd3L._SX322_BO1,204,203,200_.jpg" alt="最新の脳研究でわかった! 自律する子の育て方 (SB新書)" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4815607117/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">最新の脳研究でわかった! 自律する子の育て方 (SB新書)</a></div><div class="amazlet-detail">工藤勇一 (著), 青砥瑞人  (著)<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4815607117/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

- 普段なら避ける書名と表紙の感じだが、著者の方の記事を前に読んだことがあって気になっていたので読んでみたら面白かった
    - [子どもが主体的に動くようになる「3つの言葉」横浜創英・工藤勇一校長インタビュー＜前編＞ \| リセマム](https://resemom.jp/article/2020/09/16/58137.html)
- 平易でよい。納得感もある
- 主張の出典が示されていればもっとよかった
