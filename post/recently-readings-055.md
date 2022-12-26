{"title":"最近読んだもの 055 - Go 並行プログラミングでやりがちなミス、ボルヘスの講義録など","date":"2022-12-26T23:20:00+09:00","tags":["readings"]}

- [Shopify Embraces Rust for Systems Programming — Development \(2022\)](https://shopify.engineering/shopify-rust-systems-programming)
    - Shopify のシステムプログラミング領域の実装言語として Rust が採用された
    - ユースケースは `high-performance network servers` や Ruby のネイティブエクステンションなどらしい
        - もうちょっと詳しく知りたい
- [Data Race Patterns in Go \| Uber Blog](https://www.uber.com/en-US/blog/data-race-patterns-in-go/)
    - Go の並行処理実装で混入しやすいバグのパターン
    - 言語仕様上、意図せず同じメモリを複数スレッドから触ってしまったりなど
    - Uber の大量の Go コードベース (`2,100 unique Go services` (!)) をもとに集計・分類されているので、網羅性も高そうで面白い
- [Dynamic Data Race Detection in Go Code \| Uber Blog](https://www.uber.com/en-US/blog/dynamic-data-race-detection-in-go-code/)
    - 上と同じシリーズの記事
    - Race Detection を継続的に回して検知・可視化・修正する体制をどう構築しているか
    - 検出された Race を重複排除・分類している部分がキモに見える
- [Best pactices for Cloud Memorystore for Redis on Google Cloud \| Google Cloud Blog](https://cloud.google.com/blog/products/databases/best-pactices-for-cloud-memorystore-for-redis/)
    - GCP の Memorystore for Redis のベストプラクティスのざっとしたまとめ
    - メモリ使用量が 80% を超えないように気をつけること
    - Redis6 にアップグレードすると I/O multithreading が使えるのでおすすめ
        - その場合複数コアが割り当てられるスペック (M3 以上) が最低限必要
    - メンテナンスウィンドウを設定すること
    - なくなると困るデータを扱うなら Standard Tier (HA 構成) を選ぶこと
    - 重いコマンドに注意すること
        - KEYS、LRANGE, HGETALL などの戻り値の量が可変のもの
    - クライアントのリトライには exponential backoff を導入すること
    - おすすめの監視メトリクス
- <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01M0CEF5U/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">詩という仕事について - J．L．ボルヘス</a>
    - 内容をちゃんとは咀嚼できていないけれど、面白く読めてしまった
    - 読ませる (聞かせる) 語り口がさすがすぎるし、異国語の古語にここまで精通しているのもすごい
    - 読んでいて中二心をくすぐられる本だった。ハーバードの図書館に数十年眠っていた講義録の録音を書き起こしたという出自のエピソードがもうすでにかっこいい

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01M0CEF5U/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/41aGjevKJaL.jpg" alt="詩という仕事について (岩波文庫)" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01M0CEF5U/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">詩という仕事について (岩波文庫)</a></div><div class="amazlet-detail">J．L．ボルヘス (著), 鼓 直 (著)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01M0CEF5U/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
