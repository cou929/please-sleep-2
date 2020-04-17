{"title":"Charles と modern.IE で本番環境での動作確認","date":"2015-02-01T22:55:45+09:00","tags":["javascript"]}

Web サイト上で動作する JavaScript を修正したので、本番環境で動作確認したい。しかも IE で。こんな時は [Charles](http://www.charlesproxy.com/) というプロキシと [modern.IE](https://www.modern.ie/ja-jp/virtualization-tools#downloads) の仮想マシンを使うと、さくっとできる。proxy をたてて js へのアクセスをローカルに向け、modern.IE の Windows 仮想環境で動作確認できる。

以下はその手順。

### Charles

本当はどんな手段でもいいんだけど、Charles は操作が簡単だった。

[Charles Web Debugging Proxy • HTTP Monitor / HTTP Proxy / HTTPS & SSL Proxy / Reverse Proxy](http://www.charlesproxy.com/)

起動後、`Tools` -> `Map Local` -> `Enable Map Local` をオンにすると、ローカルファイルへの置き換えのルールを追加できる。

![](/images/charles-map-local.png)

また `Proxy` -> `Proxy Setting` からプロキシの設定・確認ができる。デフォルトでは 8888 ポートなので、ネットワークの環境設定から http プロキシの向き先を `localhost` の `8888` にすればよい。

![](/images/charles-system-proxy-setting.png)

Charles は一応有償のソフトウェアだけど、制限つきで無料でも使える。その制限というのが、30 分ごとに強制終了するというもの。再起動すれば引き続き利用できる。面倒だけどちょっと確認する程度ならば十分。

### moern.IE

次に modern.IE の仮想マシンのイメージをセットアップする。[ダウンロードページ](https://www.modern.ie/ja-jp/virtualization-tools#downloads) から必要な OS/Browser バージョンと仮想化のプラットフォームを選びイメージをダウンロードする。今回は [VirtualBox](https://www.virtualbox.org/) を例として説明する。

VirutalBox がインストール済みであれば、ダウンロードしたイメージを展開して読みこむだけで、特になにもしないでも Windows の仮想環境が起動できると思う。ネットワークの設定はデフォルトでは NAT になっているので、このままでよい。

あとはこの仮想環境での http proxy を設定し、ホスト OS の Charles に向ければ良い。ホスト OS の ip は `ipconfig` などで調べられる。コマンドプロンプトを起動し、

```
ipconfig /all
```

これで、`Default Gateway` に設定されている ip がそれになる。デフォルトでは `10.0.2.2` になるはず。

![](/images/modernie-ipconfig.png)

インターネットオプションからプロキシとしてこの ip を設定する。

![](/images/modernie-internet-option-proxy.png)

これで検証したいサイトにブラウザでアクセスすればよい。もちろん開発者ツールなども通常通り使える。

### その他

[BrowserStack](http://www.browserstack.com/) などのホスティング型のサービスのほうが、もちろんサポートしている OS/Browser のバリエーションが多い。けれども、今回のようなローカル型のほうが設定がが簡単なのと、なにより検証環境が遅くないのが大きいメリットだと思う。BrowserStack 等はたいてい海外に接続することになるので、非常に動作が重い。特にデバッガやプロファイラを使いたい場合はローカルでやったほうがいい。

自前のプロダクトで動作する js なら、ステージング環境を作ってそこで確認したほうがいい。一方で js を他サイトに提供するようなプロダクトの場合、proxy を使って本番のサイト上でも動作確認をしたほうがいいと思う。

今回のやり方は、自動テストをやった上での手動の最終確認や、ちょっとした動作検証などでは便利な方法だ。

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00ME9TTMA/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="http://ecx.images-amazon.com/images/I/51K%2BQlzoQlL._SL160_.jpg" alt="フロントエンドエンジニア養成読本［HTML ，CSS，JavaScriptの基本から現場で役立つ技術まで満載！］ Software Design Plus" style="border: none;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00ME9TTMA/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">フロントエンドエンジニア養成読本［HTML ，CSS，JavaScriptの基本から現場で役立つ技術まで満載！］ Software Design Plus</a><div class="amazlet-powered-date" style="font-size:80%;margin-top:5px;line-height:120%">posted with <a href="http://www.amazlet.com/" title="amazlet" target="_blank">amazlet</a> at 15.02.01</div></div><div class="amazlet-detail">技術評論社 (2014-08-04)<br />売り上げランキング: 7,020<br /></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00ME9TTMA/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774166146/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="http://ecx.images-amazon.com/images/I/51xpq1-cODL._SL160_.jpg" alt="フロントエンド開発徹底攻略 (WEB+DB PRESS plus)" style="border: none;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774166146/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">フロントエンド開発徹底攻略 (WEB+DB PRESS plus)</a><div class="amazlet-powered-date" style="font-size:80%;margin-top:5px;line-height:120%">posted with <a href="http://www.amazlet.com/" title="amazlet" target="_blank">amazlet</a> at 15.02.01</div></div><div class="amazlet-detail">cho45(さとう) 五十嵐 啓人 伊野 亘輝 須藤 耕平 片山 育美 池田 拓司 高津戸 壮 石本 光司 竹迫 良範 伊藤 直也 若原 祥正 沢渡 真雪 <br />技術評論社 <br />売り上げランキング: 160,879<br /></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774166146/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
