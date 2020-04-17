{"title":"いくつかのサービスの embed タグ","date":"2012-10-12T00:03:21+09:00","tags":["javascript"]}

### youtube

    <iframe width="960" height="720" src="http://www.youtube.com/embed/mHtdZgou0qU" frameborder="0" allowfullscreen></iframe>

- iframe 直貼り方式
- パラメータは embed タグを作るところで from が出るので, そこでユーザに選ばせる
  - 再生後におすすめが出るかどうか (クエリパラメータで制御している)
  - https を使うかどうか
  - privacy-enhanced mode
    - www.youtube-nocookie.com というドメインから配信して, youtube.com のクッキーを使った情報収集を行わない (ということをドメインを変えることで明示している)
  - 昔の embed タグ (flash)
  - サイズ

この手のを javascript でやると, その script タグのある場所に確実にコンテンツを差し込むのは若干面倒だし, とはいえ `document.write` を使うのは避けたいので, js よりも iframe + オプションはフォームから というのがベターだと思う

### viemo

    <iframe src="http://player.vimeo.com/video/43382919?title=1&amp;byline=1&amp;portrait=1" width="960" height="306" frameborder="0" webkitAllowFullScreen mozallowfullscreen allowFullScreen></iframe> <p><a href="http://vimeo.com/43382919">Christian Johansen - Pure, functional JavaScript</a> from <a href="http://vimeo.com/webrebels">Web Rebels Conference</a> on <a href="http://vimeo.com">Vimeo</a>.</p>

- iframe 直貼り方式
- フォームで各種パラメータを変動させる
- パラメータは表示関係のものばかり
  - https のページに貼るときはサイトオーナーが自分で注意する必要がある
- iframe の外に動画の説明などが p タグでついてくる. スタイルはサイトオーナーが調整しろということか

### slideshare

    <iframe src="http://www.slideshare.net/slideshow/embed_code/14112177" width="427" height="356" frameborder="0" marginwidth="0" marginheight="0" scrolling="no" style="border:1px solid #CCC;border-width:1px 1px 0;margin-bottom:5px" allowfullscreen> </iframe> <div style="margin-bottom:5px"> <strong> <a href="http://www.slideshare.net/toddanglin/5-tips-for-better-javascript" title="5 Tips for Better JavaScript" target="_blank">5 Tips for Better JavaScript</a> </strong> from <strong><a href="http://www.slideshare.net/toddanglin" target="_blank">Todd Anglin</a></strong> </div>

- iframe 直貼り方式
- フォームでパラメータの変更
  - クエリパラメータで何枚目のスライドからはじめるかなども選べる
- こちらも http 限定. ssl はサイトオーナーが気をつけろ方式
- 説明文が iframe の外にもついていて, この点も vimeo と同様
- wordpress 用のタグもある

        [slideshare id=14112177&doc=tips-for-better-javascriptv1-120829183910-phpapp02]

### speaker deck

    <script async class="speakerdeck-embed" data-id="504dd52fd3696f000203e6b5" data-ratio="1.3333333333333333" src="//speakerdeck.com/assets/embed.js"></script>

- js 方式
- プロトコルは `//` からはじめる方式
  - ブラウザがサイトのプロトコルに応じて切り替えてくれる
- js は長いのであとでちゃんと読もうと思うけど, 結局は iframe を作って埋め込んでいるようだ
  - ただ domcontentloaded に何かしたりしているので, js じゃないとダメな理由がちゃんとありそう. これだけの js を書いて何をしたかったのか, ちゃんと見てみたい

### gist

    <script src="https://gist.github.com/3735053.js?file=notation_checker.py"></script>

- js 方式
- https 限定方式
  - そもそも gist 自体すべて https
- 驚いたことに js からは html と css を `document.write` していた...
  - たしかに, シンタックスハイライトまで考えると, これはなかなか難しいのか..
  - いや, であれば iframe にして, どにかく `document.write` を避けたほうがいい気がするのだが

### 感想

- 第三者コンテンツ提供型のサービスの場合は, js でがんばるよりも iframe タグをそのまま提供するほうがなにかと簡単そう
- embed タグの提供 ui にフォームを付けて各種パラメータをユーザに選んでもらうのがベター
- http or https の切り替えは意外と気にしていないサービスが多い
  - 可能ならばすべて https なサービスにすべきなんだろうけど, そうでないならユーザに選んでもらう or プロトコル省略の url を使うのがベターかな
