{"title":"最近読んだもの 10","date":"2021-07-26T20:30:00+09:00","tags":["readings"]}

## 記事

- [Adding support for cross\-cluster associations to Rails 7 \| The GitHub Blog](https://github.blog/2021-07-12-adding-support-cross-cluster-associations-rails-7/)
    - GitHub は 15 個のプライマリ、15 個のリプリカという DB 構成
    - functional partitioning (テーブルごとに別の DB インスタンスに載せている) を採用している
        - 反対は horizontal sharding で一テーブルを別 DB インスタンスに分割
    - よって join の際には各テーブルごとにクエリを発行してアプリ側でよしなにマージしている
        - 二つ目のクエリをうまく `id IN (...)` のようにする必要がある
    - 開発者が手書きすると間違いやすいので、AR のアソシエーションをもとにクエリを作成してくれるようにした
    - もともと GitHub 内部 gem として運用されていたが枯れたので upstream にもマージされた
        - ありがたい
- [Migrating Facebook to MySQL 8\.0 \- Facebook Engineering](https://engineering.fb.com/2021/07/22/data-infrastructure/mysql/)
    - 課題が多くあるときにそれを分類して評価してから対応していくアプローチが良い
    - あわせてレプリケーション方式を RBR に変更したことで、トランザクション分離レベルを下げてデッドロックを起きづらくできたらしい
        - `A few applications hit repeatable-read transaction deadlocks involving insert … on duplicate key queries on InnoDB. 5.6 had a bug which was corrected in 8.0, but the fix increased the likelihood of transaction deadlocks. After analyzing our queries, we were able to resolve them by lowering the isolation level. This option was available to us since we had made the switch to row-based replication.`
- [No, we don’t use Kubernetes \| Ably Blog: Data in Motion](https://ably.com/blog/no-we-dont-use-kubernetes)
    - 細部は全然違うけど、ハイレベルでは k8s  でできることは他の既存の技術でもできる
    - なので既存のインフラ資産のあるシステムが k8s に移行するのはなかなかペイしないと思う
        - オンプレとクラウドを共存させたいとか特殊な事情があればハマることもあるのかもしれない
    - 一方で、自分は k8s 初学者かつ現在 pj は既に k8s で構築されているという立場から見ると、あることを実現させるための方法として k8s ならどうやるかを学ぶことはわりと安心感がある
        - 実業務ですぐに役立つのと、k8s が今後もデファクトであり続けると信じるならば、学習した内容の有効期限が長いので、投資として悪くないなという気持ちになる
    - 一からシステムを構築するときに k8s を選ぶかというと悩ましい
        - 詳しい人がメンバーにいることが必要条件のひとつ
        - ある程度の規模以上のトラフィックを受ける見込みがあるかがひとつ
- [Unpacking Observability: A Beginner's Guide \| Medium](https://adri-v.medium.com/unpacking-observability-a-beginners-guide-833258a0591f)
    - observability とは何かの定義として、途中で引用されている内容がよかった
    - `“You can understand the inner workings of a system […] by asking questions from the outside […], without having to ship new code every time. It’s easy to ship new code to answer a specific question that you found that you need to ask. But instrumenting so that you can ask any question and understand any answer is both an art and a science, and your system is observable when you can ask any question of your system and understand the results without having to SSH into a machine.”`
- [Mitchell's New Role at HashiCorp](https://www.hashicorp.com/blog/mitchell-s-new-role-at-hashicorp)
    - Mitchell Hashimoto が hashicorp の IC になるらしい
    - いい話
- [Extend your Golang app with embedded WebAssembly functions in WasmEdge](https://www.secondstate.io/articles/extend-golang-app-with-webassembly-rust/)
    - 手続きが色々あってなかなか面倒そう
