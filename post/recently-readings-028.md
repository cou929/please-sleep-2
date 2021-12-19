{"title":"最近読んだもの 28","date":"2021-12-19T23:00:00+09:00","tags":["readings"]}

## 記事

- [Upgrading MySQL at Shopify — Infrastructure \(2021\)](https://shopify.engineering/upgrading-mysql-shopify)
    - shopify での MySQL アップグレード事例
    - ブラックフライデー・サイバーマンデー (BFCM) が年に一度の一大イベントで、それに備えて 8 ヶ月前から準備するらしい
    - MySQL (percona server) はマーチャントごとの水平シャーディングで、replica も合わせて全体では 1000 インスタンス規模らしい
    - アップグレードの準備はすんなり行ったが、ロールバックのテストが上手く行かず、最終的に percona server にパッチをあてる必要がでた
        - 内容もそうだが、ロールバックの検証を重点的にやっているのが良かった
    - もう一つ良かったのが、このプロジェクトの目標の一つが「MySQLのアップデートを定型業務化する」ということ
        - そのためにドキュメンテーションや作業の toil をへらすためのツール開発を頑張ったらしい
- [Understanding Zeitwerk in Rails 6 \| by Marcelo Casiraghi \| Cedarcode \| Medium](https://medium.com/cedarcode/understanding-zeitwerk-in-rails-6-f168a9f09a1f)
    - zeitwerk と classic で挙動が異なるケースに遭遇したので、この 2 つの違いを具体的に知らないので読んでみた
    - `const_missing` が起こったら load/require するのか Module#autoload を使うのかという違いらしい
    - classic には、例えば autoload される順序に依存して同名の違うオブジェクトがロードされてしまうといった問題があったが、zeitwerk ではこれが解消している
    - ただ Module#autoload の挙動を知らないので、これを読んだだけだとなぜその問題が解消するかよくわからなかった
- [Go 1\.18 Beta 1 is available, with generics \- go\.dev](https://go.dev/blog/go1.18beta1)
    - ついにジェネリクスが
- [Ruby on Rails — Rails 7\.0: Fulfilling a vision](https://rubyonrails.org/2021/12/15/Rails-7-fulfilling-a-vision)
    - marginalia がメインラインに取り込まれたのと、load_async が気になった
