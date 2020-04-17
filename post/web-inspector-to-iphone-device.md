{"title":"iPhone 実機に対して Web Inspector でデバッグ","date":"2014-03-26T23:06:20+09:00","tags":["javascript"]}

確か iOS6 の頃からだけど、iPhone の実機に対して Web Inspector (開発者ツール) でデバッグができる。むかしはそういうツールが必要だった気がするけど、もう現代では Safari だけあればよい。また iPhone Safari だけでなくアプリ内の WebView も対応していて便利。

手順も簡単で、

1. iPhone の Safari の設定で Web Inspector をオンにする
   - Safari -> Advanced -> Web Inspector

<img style="max-width:200px" alt="" src="/images/iphone_wi_conf.png"/>

2. iPhone と Mac を有線で接続
3. Mac で Safari を起動し Develop -> `端末名` -> `ページ名` を選ぶ

<img style="width:90%" alt="" src="/images/iphone_wi_safari.png"/>

たったこれだけで非常に手軽。余計なものを入れなくていいのもうれしい。

