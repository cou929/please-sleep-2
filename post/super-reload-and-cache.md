{"title":"リソースの種類・タイミングとスーパーリロード時のブラウザキャッシュ","date":"2013-08-02T22:14:07+09:00","tags":["html"]}

たとえば Chrome で Shift + Command + R というキーでページをリロードすると、普段キャッシュが効く部分でも強制的にサーバからリソースを取得するようになる。正式名称かわからないが、スーパーリロードとよばれる動作だ。Chrome だけでなく最近のメジャーブラウザにはほとんどこの機能があるはずだ。

Command + R キーでの通常のリロードをした場合。おそらくほとんどのリソースは、1) まずはサーバへ問い合わせ 2) 304 でレスポンスされる 3) 実際のデータはローカルのキャッシュを使う、という挙動になるはずだ。

![](/images/normal_reload.png)

またロケーションバーにフォーカスを移してエンターキーを押した場合、多くのリソースはキャッシュから読み込まれる。通常のリロードとの違いはサーバへリクエストを飛ばしていない点だ。

![](/images/page_refresh.png)

これがスーパーリロードになると、すべてのリソースをサーバからとりなおす。

![](/images/super_reload.png)

### リソースのタイプとタイミングによる違い

今回このサンプルコードを考える。img (`new Image` でのリクエスト送信), iframe, script の 3 種類のリソースについて、それぞれ通常の html タグ、js での動的生成、onload 後の DOM へ挿入の 3 つのタイミングでロードするサンプルだ。

<script src="https://gist.github.com/cou929/6139492.js"></script>

リソースを配信しているサーバはデフォルト設定の nginx。またロードしている html にはキャッシュを制御する meta 要素などの宣言はない。

このサンプル html を読み込み、スーパーリロードすると次のことがわかる。

- iframe の html リソースは 3 つのタイミングすべてについて、 cache から読み込まれる可能性が高い
- onload 後に送信される画像リクエストは cache から読み込まれる可能性が高い
- 通常のリロードでも前述のパターンはキャッシュから読み込まれる可能性が高い
- 手元の Chrome, Firefox (いずれも canary, nightly の最新) で確認

![](/images/reload_test_html.png)

ブラウザキャッシュの動作として、このような傾向があることがわかった。開発時には注意する必要がある。一般的なブラウザキャッシュの効きやすさとの関連は不明だが、自サービスでホストする js や widget の変更時にも考慮するといいかもしれない。

html での宣言やサーバの設定によってまた違った制御も可能なはずなので、ひきつづき調査したい。

### 参考

- [君は3つのリロードを知っているか？ - os0x.blog](http://os0x.hatenablog.com/entry/20110617/1308280740)