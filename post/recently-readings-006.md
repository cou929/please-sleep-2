{"title":"最近読んだもの 6","date":"2021-06-28T01:10:00+09:00","tags":["readings"]}

## 記事

- [Online migrations at scale](https://stripe.com/blog/online-migrations)
    - データストアの移行。細部が地に足がついている内容でとても良かった
    - 特に旧データストアから新データストアにデータコピーする際に本番に負荷をかけないために hadoop クラスタに入っていたデータを元にして行ったこと、[github/scientist](https://github.com/github/scientist) というライブラリで experiment を記述しレコードレベルのデータ不整合があった際はすぐに例外を投げるようにしたのが良い
- [Why we decided to rewrite our iOS & Android apps from scratch — in React Native \| by Naoya Makino \| Mercari Engineering \| Jun, 2021 \| Medium](https://medium.com/mercari-engineering/why-we-decided-to-rewrite-our-ios-android-apps-from-scratch-in-react-native-9f1737558299)
    - フロントを React エコシステムに集約できると、理想的には生産性は確かにかなりあがりそう
    - Android クラッシュフリーレート低下見積もりなど RN 固有の問題も記載されていて良い
    - 合意形成が大変だったというのがひしひしと伝わる
- [Why Discord is Sticking with React Native \| by Fanghao \(Robin\) Chen \| Discord Blog](https://blog.discord.com/why-discord-is-sticking-with-react-native-ccc34be0d427)
    - 通話やリアルタイム性、同時接続数など要件が厳しそうな Discord のアプリが RN で成り立っているのは示唆的
    - パフォーマンスなど課題はあるとはいえ、それが Discord でも致命的では無いという
- [Consistency Models](https://jepsen.io/consistency)
    - 平行システムのいろいろな整合性モデルの説明
    - 前提となる用語の整理から始まってるのが好感
- [New – AWS Step Functions Workflow Studio – A Low\-Code Visual Tool for Building State Machines \| AWS News Blog](https://aws.amazon.com/blogs/aws/new-aws-step-functions-workflow-studio-a-low-code-visual-tool-for-building-state-machines/)
    - サーバーレス環境でのアプリケーションの設計は通常とは結構変わりそうで、どこかでキャッチアップしたいと思った
- [Want to Debug Latency?\. In the recent decade, our systems got… \| by Jaana Dogan \| Medium](https://rakyll.medium.com/want-to-debug-latency-7aa48ecbe8f7)
    - latency distribution heat map  とそこから trace にも入れるのがすごい
- [Software Estimation Is Hard\. Do It Anyway\. \- Jacob Kaplan\-Moss](https://jacobian.org/2021/may/20/estimation/)
    - 見積もりは避けられないから向き合おうというのは確かに
- [My Software Estimation Technique \- Jacob Kaplan\-Moss](https://jacobian.org/2021/may/25/my-estimation-technique/)
    - 期間の見積もりとともに不確実性も見積もるのがなるほど
    - `choose a model and stick with it`
    - `JFDI`
- [Custom "cops" for RuboCop: an emergency service for your Ruby code — Martian Chronicles, Evil Martians’ team blog](https://evilmartians.com/chronicles/custom-cops-for-rubocop-an-emergency-service-for-your-codebase)
    - やっぱ AST のとこがむずい
- [CPU Utilization is Wrong](http://www.brendangregg.com/blog/2017-05-09/cpu-utilization-is-wrong.html)
    - cpu utilization にはメモリ待ちも含まれているので、実はメモリバウンドな状況と見分けがつかない
- [Kubernetes Failure Stories](https://k8s.af/)
    - k8s の障害事例を集めたらしい
- [Site Reliability Engineering for Kubernetes \| by Tammy Bryant Butow \| Medium](https://tammybutow.medium.com/site-reliability-engineering-for-kubernetes-b52877c70fb7)
    - 上記を集計して障害の傾向を紹介
    - SLO 等を事業の方針と紐付けて設計しているのはなるほどだった

## コード

- [github/scientist: A Ruby library for carefully refactoring critical paths\.](https://github.com/github/scientist)
    - 複数のコードブロックを渡すとそれぞれを実行し結果が一致していなければエラーをあげる
    - 慎重にリファクタリングや移行を行う際に有用
    - 実行順序をランダムにしたり、比較関数のカスタマイズや各種フック、実行条件の指定、実行時間の計測などの機能が揃っていて行き届いている
    - 実装はシンプルで、複数のコードブロックをランダム順に実行し結果を比較している
        - 各コードブロックの実行結果クラス Observation と全結果の統合クラス Result というエンティティを作っているのは綺麗で良い
        - 命名が全体的に良い
            - experiment
            - control
            - observation
            - scientist
    - ruby の構文にまだ不慣れなところがあるのでそのエクササイズにもなった
