{"title":"アドテクゑびす会 #2 〜やっときた動画広告元年","date":"2014-10-16T00:48:41+09:00","tags":["ad"]}

[アドテクゑびす会 #2 〜やっときた動画広告元年｜EventRegist（イベントレジスト）](http://eventregist.com/e/at-yebisu2)

mobile safari での動画表示に関する発表が気になったので参加。

## CSS スプライトによるパラパラ動画 (DAC)

- iOS Safari は動画をタップしないと再生されない
- 広告での利用としては望ましくないので、自前実装でがんばる

### CSS スプライトによる動画

- 動画前処理
  - 動画を 0.1 秒ごとに画像化 (10 fps)
  - 1秒をつなげて1枚にする
  - CSS スプライトでパラパラ動画に
  - adobe media encoder & image magik で前処理
- multiple background images にして高速化
- 複数枠同期にも対応
- VAST 3.0 対応

multiple background images は知らなかったんだけど、複数画像を重ね合わせて要素の背景画像にするというものらしい。指定は単純にカンマ区切りで `background` の値を指定するだけで簡単だった。

[Using CSS multiple backgrounds - Web developer guide | MDN](https://developer.mozilla.org/en-US/docs/Web/Guide/CSS/Using_multiple_backgrounds)

複数枠同期は cookpad の真似をしたかったんだと思う。 [クックパッド、動画広告をスマートフォンで開始 〜業界初・上下広告枠間の動画再生位置を連携し、再生完了率の高い動画広告を実現〜 | クックパッド株式会社](https://info.cookpad.com/press/2014/0808)

### その他

- android で video タグの自動再生 (`autoplay`) が素直にできないらしい。workaround で対応したらしいが詳細は不明。

[Android webview html5 video autoplay not working on android 4.0.3 - Stack Overflow](http://stackoverflow.com/questions/15946183/android-webview-html5-video-autoplay-not-working-on-android-4-0-3) によると `onload` でもいいので別のイベント時に `play()` を呼び出せばよいとある。これかな。

## Jストリーム紹介 (JStream)

- 動画配信専門系の事業者
- トランスコスモスの子会社

### 動画ストリーミング方式

-rtmp
  - flash player
- hls (http live streaming)
  - ios 標準
- hds (http dynamic streaming)
  - adobe
- pd (プログレッシブダウンロード)

### equipmedia

http://www.jstream.jp/service/delivery/equipmedia/

- 動画をあげるとマルチデバイスに変換して適切なものを返す
- クローズドなサイトでの利用に強い
  - この点が youtube や ustream との差別化

### medialize

http://www.jstream.jp/service/delivery/medialize/

- vast にも対応している

### mpeg-dash

http://en.wikipedia.org/wiki/Dynamic_Adaptive_Streaming_over_HTTP

- 次世代の統一規格を目指す
- apple の対応が遅い

### safari でデフォルトプレイヤーを立ち上げない技術

- DAC とほぼ同じだが、音声にも対応している
- 10 fps なのも同様

### 動画の著作権と映画と広告について (livepass)

- スマホの push 通知に動画をのせるサービスをやっているらしい。[TOPページ - ActAds - livepass PUSH](http://www.actads.com/)
- [Idomoo Personalized Videos | Videos that get personal](http://idomoo.com/) というパーソナライズされた動画コンテンツ作成技術の代理店をしているらしい
