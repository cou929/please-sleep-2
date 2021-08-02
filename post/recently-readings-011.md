{"title":"最近読んだもの 11","date":"2021-08-02T20:30:00+09:00","tags":["readings"]}

## 記事

- [Redis persistence demystified](http://oldblog.antirez.com/post/redis-persistence-demystified.html)
    - fsync でカーネルバッファにあるデータがディスクに書き出されると思っていたけど、実はもっとステップがあった
        - ただアンコントローラブルなのでユーザーは write (ユーザー空間からカーネルバッファへのデータ移動) と fsync (カーネルバッファからディスクへの書き出し) と抽象化して認識しておけば良い
    - 説明の論理立てが上手でわかりやすく楽しく読めた
        - 特に序盤、各レイヤーのファイルキャッシュとデータ汚染を整理してから rdb, aof の説明に入っていく流れ
- [7 Redis Worst Practices \| Redis Labs](https://redislabs.com/blog/7-redis-worst-practices/)
    - `Numbered databases/SELECT` は機能の存在も知らなかった
- [\(2\) Solving The Three Stooges Problem : RedditEng](https://www.reddit.com/r/RedditEng/comments/obqtfm/solving_the_three_stooges_problem/)
    - thundering herd 対策
    - 同種のリクエストごとにロックをとって、ひとつだけ処理するようにする。処理結果はキャッシュし、待っている他のリクエストはキャッシュからレスポンスする
- [Making Rails run just a few tests faster](https://world.hey.com/jorge/making-rails-run-just-a-few-tests-faster-2c82dc4b)
    - テストの並列実行は、それぞれに db の初期化と fixture のロードをするので、少数のテストだけ流す場合はオーバーヘッドが大きい
    - 閾値以下のテスト数の場合はシングルプロセスで流すパッチがマージされた
    - 筆者の環境では 70, 80 テストくらいまではシングルプロセスの方が早かったらしく、思ったよりも多い
- [Netcat \- All you need to know :: ikuamike](https://blog.ikuamike.io/posts/2021/netcat/)
    - いろんな実装があるのは知らなかった。歴史
    - 実際この中で使いそうなのは多段サーバーでコマンドを中継するやつかな
        - 実際やったことがあるのもこれだけ
        - 他のは別の方法がある
        - 最近は踏み台サーバーを置いて多段 ssh みたいな必要性も減ってきている
    - pentester (という言葉も初めて聞いた。脆弱性を探す人) とか ctf やる時などセキュリティ系の文脈のようなので、また別のニーズがあるんだと思う
