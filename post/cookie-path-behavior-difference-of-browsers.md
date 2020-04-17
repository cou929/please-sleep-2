{"title":"ブラウザごとの cookie path の挙動の違い","date":"2013-12-23T10:34:19+09:00","tags":["browser"]}

![](/images/path.png)

Path 付きの Cookie を保持するブラウザがどのような場合に Cookie ヘッダをサーバに送信するのか。その挙動がブラウザによって異なるようなので調査した。

### 調査結果

<table border="1">
<tr><th>ブラウザ</th><th>挙動</th></tr>
<tr><td>Chrome</td><td>RFC 6265 に準拠</td>
<tr><td>Firefox</td><td>RFC 6265 に一部準拠の独自挙動</td>
<tr><td>IE</td><td>ネットスケープの仕様に準拠</td>
<tr><td>Safari</td><td>ネットスケープの仕様に準拠</td>
<tr><td>Opera</td><td>12 以前はネットスケープの仕様に準拠。15 以降は RFC 6265 に準拠</td>
</table>

- スラッシュを考慮したマッチを行っているものを RFC 6265 準拠と呼んでいる。詳細は後述
- 単純に前方一致でマッチしているものをネットスケープの仕様に準拠と呼んでいる

### 仕様

そもそも Cookie は仕様からして揺れている。現在ある文章は以下の 4 つ。

- [Client Side State - HTTP Cookies](http://curl.haxx.se/rfc/cookie_spec.html)
  - Netscape が定めた文章。現在ではもとのリソースはなくなっており、第三者によってアーカイブされているものしかない。古いがものによっては依然デファクト。
- [RFC 2109 - HTTP State Management Mechanism](http://tools.ietf.org/html/rfc2109)
  - Obsolated かつ Historical。1997 の文章。
- [RFC 2965 - HTTP State Management Mechanism](http://tools.ietf.org/html/rfc2965)
  - Obsolated かつ Historical。2000 年の文章。Cookie2 / Set-Cookie2 という名前のヘッダを定義しようとしていたらしいが、当然実装は広まっていない。
- [RFC 6265 - HTTP State Management Mechanism](http://tools.ietf.org/html/rfc6265)
  - 2011 年の文章でこれが最新。

最も古いネットスケープの文章がデファクトスタンダードとなっており、今でも一部の実装はこれに従っているそうだ。後発の RFC の仕様も完全な上位互換にはなっていないようで、ただ単に最新の仕様を読めば OK とはいかず、実装の差異や歴史的経緯を考慮しないと行けない状況らしい。

ここでは最新の RFC 6265 と最古のネットスケープの文章をみてみよう。

まずは RFC 6265 から。5.4 Cookie Header の節に Cookie ヘッダを送信する際のアルゴリズムが記載されている。パスに関しては

> *  The request-uri's path path-matches the cookie's path.

リクエスト url のパスがクッキーのパスに "path-matches" した場合にクッキーを送信する、と書かれている。

"path-matches" の定義は 5.1.4. Paths and Path-Match の節にある。以下の条件をひとつ以上満たす場合に "path-matches" とみなされる。

- cookie-path と request-path が同一である場合
- cookie-path が request-path の prefix (前方一致) で、かつ cookie-path の最後の文字がスラッシュの場合。
- cookie-path が request-path の prefix (前方一致) で、かつ request-path の cookie-path に含まれない最初の一文字がスラッシュの場合

つまり '/foo' というパスのクッキーがあった場合、"/foo", "/foo/", "/foo/bar" といった url のケースではクッキーが送信されるが、"/foobar" などには送信されない。

一方でネットスケープの仕様は明確に異なっている。

> path=PATH
>
> The path attribute is used to specify the subset of URLs in a domain for which the cookie is valid. If a cookie has already passed domain matching, then the pathname component of the URL is compared with the path attribute, and if there is a match, the cookie is considered valid and is sent along with the URL request. The path "/foo" would match "/foobar" and "/foo/bar.html". The path "/" is the most general path.

とあるように、単純に前方一致でのみ判定する。"/foo" というパスのクッキーは "/foobar", "/foo/bar" にマッチすべきとう例まであげている。

"path" という属性名から、おそらく多くの人は URL なりディレクトリなりのパスを想像するので、最新の RFC 6265 の仕様はその意味で直感的だと思う。一方で以前は一ベンダーのものとはいえ明文化されたものがあるため、後方互換性のために古くからあるブラウザは挙動を変えにくいことも理解できる。こうしてブラウザ間の挙動の違いが生まれてしまっているようだ。

### ブラウザごとの挙動

#### Chrome

RFC 6265 と同様の挙動。

- パスが "/foo" の場合
  - "/foo", "/foo/", "/foo/bar" の場合送信
  - "/foobar" の場合送信しない
- パスが "/foo/" の場合
  - "/foo/", "/foo/bar" の場合送信
  - "/foo", "/foobar" の場合送信しない

Chrome 14 以降で確認。

#### Firefox

RFC 6265 の動きに近いが、クッキーのパス末尾がスラッシュの場合にアレンジが加えられている。

- パスが "/foo" の場合
  - "/foo", "/foo/", "/foo/bar" の場合送信
  - "/foobar" の場合送信しない
- パスが "/foo/" の場合
  - "/foo", "/foo/", "/foo/bar" の場合送信
    - "/foo" にマッチするのが RFC 6265 には無い挙動
  - "/foobar" の場合送信しない

firefox 3.6 以降で確認。

#### IE

ネットスケープの挙動に準拠。単純にパスの前方一致で判定している。

- パスが "/foo" の場合
  - "/foo", "/foo/", "/foobar", "/foo/bar" いずれの場合も送信
- パスが "/foo/" の場合
  - "/foo/", "/foobar" の場合送信
  - "/foo", "/foobar" の場合送信しない

IE 6 以降で確認。IE 11 でも同様だった。

#### Safari

ネットスケープの挙動に準拠。単純にパスの前方一致で判定している。

- パスが "/foo" の場合
  - "/foo", "/foo/", "/foobar", "/foo/bar" いずれの場合も送信
- パスが "/foo/" の場合
  - "/foo/", "/foobar" の場合送信
  - "/foo", "/foobar" の場合送信しない

Safari 5 以降で確認。

#### Opera

Opera 12 はネットスケープの仕様。

- パスが "/foo" の場合
  - "/foo", "/foo/", "/foobar", "/foo/bar" いずれの場合も送信
- パスが "/foo/" の場合
  - "/foo/", "/foobar" の場合送信
  - "/foo", "/foobar" の場合送信しない

Opera 15 以降は RFC 6265 の仕様。

- パスが "/foo" の場合
  - "/foo", "/foo/", "/foo/bar" の場合送信
  - "/foobar" の場合送信しない
- パスが "/foo/" の場合
  - "/foo/", "/foo/bar" の場合送信
  - "/foo", "/foobar" の場合送信しない

エンジンを変えたタイミングで Chrome と挙動が一致しているので納得。

### 調査につかったコード

[cou929/cookie-baker](https://github.com/cou929/cookie-baker)

Set-Cookie するだけの機能をもつアプリを heroku において、BrowserStack + Selenium WebDriver でテストケースをまわした。テストのシナリオは以下。

- /foo というパスのクッキーを食わせる。以下の URL にアクセスしクッキーが送信されるかチェック
  - /foo
  - /foo/
  - /foobar
  - /foo/bar
- /foo/ というパスのクッキーを食わせる。以下の URL にアクセスしクッキーが送信されるかチェック
  - /foo
  - /foo/
  - /foobar
  - /foo/bar
- パスが /foo と /foo/ の両方のクッキーを食わせる。以下の URL にアクセスしクッキーが送信されるかチェック
  - /foo
  - /foo/
  - /foobar
  - /foo/bar

ちなみに、WebDriver の `all_cookies` メソッドでなぜかクッキーの値がとれず、しょうがないので発行・受信した Set-Cookie / Cookie ヘッダの内容を html に出力してそれを読み取るようにした。html の解析も WebDriver の `find_element` ではなぜかうまく行かなかったのでページ全体のソースを取り出して nokogiri に食わせてパースした。けっこうはまったが、いまだに原因はわかっていない。

今回の調査で BrowserStack の自動テスト無料枠を使いきってしまったので、途中からは [Sauce Labs](https://saucelabs.com/) も使った。こちらもすぐ上限になるので、なにか別の手を考えたい。ほとんどの場合おさえておきたいブラウザバージョンは、IE の各バージョンとそれ以外のブラウザの最新版 (Opera は 12 と 最新でチェックできると理想的) だけで事足りる。[Modern.IE](http://www.modern.ie/) の VM で IE をカバーすれば手元の Mac 上で完結できないかなと妄想している。

### 参考

- [どさにっき](http://ya.maya.st/d/201110b.html#s20111014_1)

