{"title":"html/css のスマホ対応 tips","date":"2014-03-22T01:03:45+09:00","tags":["html"]}

モバイル向けの html を書く機会があった。PC とはちがう tips をいろいろと知ることができたのでメモ。

### viewport

表示幅やピンチでの拡大縮小を制御する meta 要素。ここで指定したサイズ等でドキュメントを表示する。例えば `width` をここで指定しないと、iPhone では横幅 980px とみなして表示されてしまうので、見た目がすごく小さくなってしまう。

`width=device-width` とすると端末にあわせてよしなに出してくれる。`scale` 系の設定は 1 にしておくとピンチでの拡大縮小を抑制できる。スマホの画面に最適化した html は、ピンチでの拡大縮小ができないほうがむしろ使いやすいので、こうしておくのも選択肢。

<pre><code data-language="html"><meta name="viewport" content="width=device-width, initial-scale=1, minimum-scale=1, maximum-scale=1, user-scalable=no"></code></pre>

とりあえずこう書いておくと間違いない。

### format-detection

<pre><code data-language="html"><meta name="format-detection" content="telephone=no"></code></pre>

こんなふうに書いておくと、電話番号っぽい文字列をかってにリンクにする動作を抑制してくれる。電話番号が表示されないページはリンクにならないほうがいいので、指定しておくとよい。

### user-select, -webkit-touch-callout, -webkit-tap-highlight-color

ユーザーによる html 上の文字列選択やリンク長押し時の挙動を制御する css 要素。次のようにすると、ユーザは文字列を選択できず、リンクを長押ししても何も出なくなる。

<pre><code data-language="css">body {
    -webkit-user-select: none;
    user-select: none;
    -webkit-touch-callout: none
    -webkit-tap-highlight-color: rgba(0, 0, 0, 0);
}</code></pre>

`user-select` が選択の制御、`touch-callout` がリンク長押し時の制御、`tap-highlight-color` が選択時のハイライト色の制御でここでは透明にしている。

こうすることで誤タップで広範囲が選択され青くなってしまうような挙動を防ぐことができる。一方でユーザーの選択肢を奪うことになるので、導入時はちゃんと検討したほうがいい。例えばアプリ内 webview では、このスタイルを指定することで、ネイティブっぽい動きになって UI に一貫性を持たせることができる。

### media-query

これは結構有名な、画面幅に応じて適用させるスタイルを切り替える css の機能。

例えば `@media screen and (min-width: 500px) { ... }` とすると、横幅が 500px 以上の時だけに適用されるスタイルを記述できる。

また幅だけではなくて、次のようにすると、

<pre><code data-language="css">@media only screen and (orientation : landscape) {
    // 横向き
}
@media only screen and (orientation : portrait) {
    // 縦向き
}</pre></code>

デバイスの横向き・縦向きを検知してそれぞれに応じたスタイルを指定できる。

