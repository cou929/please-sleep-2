{"title":"はてなダイアリーからの記事移行","date":"2017-12-17T16:54:32+09:00","tags":["etc"]}

思い立って昔はてなダイアリーで書いていた記事をこのブログに移行した。

昔のブログは [フリーフォーム フリークアウト](http://d.hatena.ne.jp/cou929_la/)。2012 年から更新していない。

以降の方針

- はてなダイアリーの機能で、エントリを movable type 形式でエクスポート
- エクスポートしたファイルを、現在のブログの形式に変換
    - 1 つの mt 形式のファイル => エントリごとに分割された md 形式のファイル
- 過去のはてなダイアリーの各記事の中身を移行後の URL へのリンクに書き換える
    - リダイレクトするのは無理なので

## 記事の移行

- 記事のエクスポート
    - `http://d.hatena.ne.jp/{your id}/port` から `Movable Type 形式` のテキストファイルを取得
- ruby 環境の準備諸々

<pre><code>brew update
brew install rbenv
rbenv install 2.4.1
rbenv exec gem install bundler
rbenv rehash
cat << EOS > Gemfile
source "http://rubygems.org"
gem "mustache"
gem "stringex"
EOS
rbenv exec bundle install --path vendor/bundle
</code></pre>

- 適当に拾ってきた mt to md converter というスクリプトを改造して使う
    - https://gist.github.com/cou929/6f2e3eb40abaea09a9d933d7f946b0e4
    - なお、以下のカスタマイズを行った
        - 出力フォーマットの修正 (元は octpress に合わせたフォーマットだったので)
        - 旧記事のリンクをコメントに埋め込み
        - 1 日に複数記事ある場合、1 ファイルにマージする
        - parse error の微修正
        - 出力ファイルの mtime を記事作成日に変更

<pre><code>rbenv exec bundle exec ruby mt_to_markdown.rb ./cou929_la.txt ./res
</code></pre>

これではてなダイアリーの各エントリを markdown 形式で保存できた。あとはいつもの手順でこれを publish した。

## 旧 URL からのリダイレクト

- 認証
    - OAuth は面倒なので [WSSE 認証](http://developer.hatena.ne.jp/ja/documents/auth/apis/wsse) で。
    - こんな感じでたたける

<pre><code>curl -H 'X-WSSE: UsernameToken Username="cou929_la", PasswordDigest="6iTSm2Gsg6ePCz+Or7UnIv/IARI=", Nonce="5kV1Q1Tq3mK5CItNrdL1Ph4UFB0=", Created="xxxx"' 'http://d.hatena.ne.jp/cou929_la/atom'
</code></pre>

- 過去エントリの取得
    - まずは全エントリの title, date, entry_id を取得する。[こちら](http://developer.hatena.ne.jp/ja/documents/diary/apis/atom) の `日記エントリーの取得` API を使う。
    - レスポンスが xml なのでパースに苦労する。`xmllint` で何故かうまくパースできず (xpath がうまくマッチしない。後述の `xpath` コマンドでなら動作するのだが。)、mac の `xpath` コマンドを使った。それでも出力がかなり微妙 (stdout への出力に改行が入らなかったり)... なので、置換で無理やりなんとかしている。

<pre><code>curl -s -H 'X-WSSE: UsernameToken Username="cou929_la", PasswordDigest="6iTSm2Gsg6ePCz+Or7UnIv/IARI=", Nonce="5kV1Q1Tq3mK5CItNrdL1Ph4UFB0=", Created="2017-08-06T19:43:07Z"' 'http://d.hatena.ne.jp/cou929_la/atom/blog' | xpath -p '//entry/link[contains(@rel,"alternate")]/@href | //entry/title/text()' 2>/dev/null | perl -nle '$_ =~ s/ href="(.*?)"/\n$1\t/g; print $_' | perl -nle '/cou929_la\/(\d+)\/(\d+)/; print "$1\t$2\t" . (split /\t/, $_)[1]' | tee -a entries.tsv
</code></pre>

- こんな感じの出力になる

<pre><code># date, etnry_id, title の tsv
20130420        1366445641       MySQL Casual Talks vol.4
20130121        1358776754      [book] Web API 設計のベストプラクティス集 "Web API Design - Crafting Interfaces that Developers Love"
20121229        1356746993      [javascript] Google JavaScript Style Guilde をリビジョン 2.64 にあわせて修正しました
...
</code></pre>

- `日記エントリーの取得` API はページングされているので、同じ要領ですべてのデータを取得。過去記事のりすとにあたるひとつの tsv ができる。

- 最後に、更新
    - [こちら](http://developer.hatena.ne.jp/ja/documents/diary/apis/atom) の `日記エントリーの編集` API を使う
    - 以下のように xml を投げ込めば良い

<pre><code>curl -s -X PUT -H 'X-WSSE: UsernameToken Username="cou929_la", PasswordDigest="6iTSm2Gsg6ePCz+Or7UnIv/IARI=", Nonce="5kV1Q1Tq3mK5CItNrdL1Ph4UFB0=", Created="2017-08-06T19:43:07Z"' 'http://d.hatena.ne.jp/cou929_la/atom/blog/20071109/1194621111' -d '<?xml version="1.0" encoding="utf-8"?><entry xmlns="http://purl.org/atom/ns#"><title>はじめます</title><content type="text/plain">移転しました http://please-sleep.cou929.nu/</content></entry>'
</code></pre>

## まだやっていないこと

- 画像の置き場所をはてなフォトライフから自前に移行する
- はてなキーワードになっているリンクの修正

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01IQQTCDA/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://images-fe.ssl-images-amazon.com/images/I/51YyioxNlgL._SL160_.jpg" alt="はてなブログ Perfect GuideBook" style="border: none;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01IQQTCDA/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">はてなブログ Perfect GuideBook</a><div class="amazlet-powered-date" style="font-size:80%;margin-top:5px;line-height:120%">posted with <a href="http://www.amazlet.com/" title="amazlet" target="_blank">amazlet</a> at 17.08.06</div></div><div class="amazlet-detail">ソーテック社 (2016-07-22)<br />売り上げランキング: 9,926<br /></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01IQQTCDA/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
