{"title":"src のない iframe を動的生成した際の location の謎","date":"2013-08-02T00:19:41+09:00","tags":["html"]}

createElement で iframe を生成し、body に appendChild するなど適当に DOM に挿入する。挿入した iframe 要素の document オブジェクトを `iframe.contentWindow.document` などとして取得。その document の write メソッドでコンテンツを作成する。サンプルコードはこのようなものだ。

<script src="https://gist.github.com/cou929/6119775.js"></script>

動的に生成した iframe の中で `window.location` を参照する。この時の疑問が 2 点。

- window.location がホスト html の内容になる
  - この例の場合 `hash_sample.html` になる
- IE の場合 window.location.hash の内容がとれない
  - `hash_sample.html#foo` とアクセスした場合、IE の場合 hash がとれない
  - IE 8, 10 で確認
  - chrome, fx は hash がとれる

自分の認識では、このように src なしで動的に生成した iframe 要素の場合、location は空という扱いになると思っていた。たとえば chrome の developer tool で試すと次のようになる。

    > iframe = document.createElement('iframe')
    <iframe>​</iframe>​
    > document.body.appendChild(iframe)
    <iframe>​…​</iframe>​
    > iframe.contentWindow.location.origin
    "null"
    > iframe.contentWindow.location.href
    "about:blank"

origin は null、href は 'about:blank' だ。

ところが前述のサンプルコードでは、この内容がホスト html のものになっている。さらに IE では hash の内容だけとれない。

ちなみに IE の場合でも、iframe から `parent.window.location.hash` とすると hash も取得することができる。`parent.window` を参照しているのでこの挙動は納得がいく。

どうしてこのようになるのか、あるいは動作を定義している仕様はあるのか、わかっていないので調べたい。
