{"title":"最近読んだもの 49 - YJIT の最適化方針、GitHub 障害報など","date":"2022-08-16T23:30:00+09:00","tags":["readings"]}

- [When is JIT Faster Than A Compiler? — Development \(2022\)](https://shopify.engineering/when-jit-faster-than-compiler)
    - yjit を例に jit compilerが行う最適化の例を紹介
    - Ruby は演算子の挙動を上書きできる言語だが、上書きされないと仮定してバイトコードにコンパイルしておき、もし演算子のオーバーライドがあればそのバイトコードを invalidate するという方針で高速化
    - 確かにこれは実行時じゃないとできないし、わかりやすくて面白い例だった
- [GitHub Availability Report: July 2022 \| The GitHub Blog](https://github.blog/2022-08-03-github-availability-report-july-2022/)
    - 毎月の障害報
    - dns 関連とジョブキュー関連
    - 特に後者は、全ての例外をキャッチして無限にリトライする作りになってしまっていたところに、バグで例外を投げてしまうコードが混ざり、ジョブ数が急増したというものだった
    - 対策は特定の例外だけリトライするように、リトライ回数上限導入、rate limit 導入
- [Scaling Sidekiq at Gusto](https://engineering.gusto.com/scaling-sidekiq-at-gusto/)
    - 実体があって良い sidekiq 運用知見
    - キューの名前を rapid, default のような抽象的なプライオリティではなく、within30sec のような具体的なレイテンシにして、人間が設定する際のぶれを無くすこと
    - 一部ジョブのための特別レーンや、並列度が高い代わりに読み取り専用のレーンを設けるといった工夫
        - 後者は開発者がジョブを実装する際にリードオンリーの前処理と更新処理を分けて考えることを促すようになったなど
- [Ruby on Rails: 3 tips for deleting data at scale](https://planetscale.com/blog/ruby-on-rails-3-tips-for-deleting-data-at-scale#)
    - railsでのデータ削除 tips
    - dependent: :destroy_async、destroy と delete の違い、大量データ削除は少しずつなど
- [The Slotted Counter Pattern](https://planetscale.com/blog/the-slotted-counter-pattern)
    - RDB でカウンターをスケールさせる方法のひとつの紹介
    - カウンターをスロットに分けて書き込みを分散させる。読み込みは全スロットに対する select
    - 簡単に効果が出せそうななるほどな方法だった
- [How do database indexes work?](https://planetscale.com/blog/how-do-database-indexes-work)
    - RDB のインデックスについて初学者向けの概要