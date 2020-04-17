{"title":"Google JavaScript Style Guide 和訳をリビジョン 2.93 にあわせて修正しました","date":"2015-07-05T03:21:49+09:00","tags":["javascript"]}

[Google JavaScript Style Guide 和訳](http://cou929.nu/data/google_javascript_style_guide/)

Google JavaScript Style Guide の本家の更新に和訳も追従した。

### 主な変更点

クリティカルな修正が多かった。そもそもの言語仕様の間違いが二点と、脆弱性につながるルールの修正。

- `NaN == NaN` が `true` になるという [間違った記述](https://github.com/cou929/Japanese-Translation-of-Google-JavaScript-Style-Guide/commit/ddb57a879e8981c7d78f6bad330db1f1adac28d5?diff=split#diff-db4b047b30ffef7855b12b527860ac2cR143) の修正
- セミコロン省略時の自動挿入について。二項演算子の前には自動挿入されないが、[されるという前提のルールになっていた](https://github.com/cou929/Japanese-Translation-of-Google-JavaScript-Style-Guide/commit/ddb57a879e8981c7d78f6bad330db1f1adac28d5?diff=split#diff-db4b047b30ffef7855b12b527860ac2cR1008)。そのためルールの必然性がなくなってしまった。
  - その旨をコメントに記載しつつ、一貫性のため過去と同じルールでこれからも行くことになった。
- `eval()` の利用を [JSON のパースに利用することを禁止](https://github.com/cou929/Japanese-Translation-of-Google-JavaScript-Style-Guide/commit/ddb57a879e8981c7d78f6bad330db1f1adac28d5?diff=split#diff-db4b047b30ffef7855b12b527860ac2cR342)。普通に `JSON.parse()` を推奨するように。
  - JSON を `eval()` でパースすると、悪意のあるコードが実行される脆弱性になる。その旨も追記されている。

ところで、いままで google code 管理だったのが、いつのまにか [GitHub に移行していた](https://github.com/google/styleguide)。[この Issue](https://github.com/google/styleguide/issues/39) によると今年の 5 月頃からのよう。

使い慣れた画面で diff などが見られるので基本的に歓迎なのだけど、過去のコミットの日付が古いように見えるのが気になる。今回更新した 2 コミットも、どちらも 2013 年ということになっていた。

### 参考

- 原文の diff
  - [Update JavaScript style guide to 2.82: · google/styleguide@5684bbc](https://github.com/google/styleguide/commit/5684bbc8b5eda7592399f144eda966846117aed9#diff-45bd843b0d7647017672882356fdcdce)
  - [Update JavaScript style guide to 2.93: · google/styleguide@7b24563](https://github.com/google/styleguide/commit/7b24563e08b7e6a6477eed22bca2eb4e6ac093fa#diff-45bd843b0d7647017672882356fdcdce)
- [訳文の diff](https://github.com/cou929/Japanese-Translation-of-Google-JavaScript-Style-Guide/commit/ddb57a879e8981c7d78f6bad330db1f1adac28d5)


<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873115736/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="http://ecx.images-amazon.com/images/I/51c9uCrhHgL._SL160_.jpg" alt="JavaScript 第6版" style="border: none;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873115736/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">JavaScript 第6版</a><div class="amazlet-powered-date" style="font-size:80%;margin-top:5px;line-height:120%">posted with <a href="http://www.amazlet.com/" title="amazlet" target="_blank">amazlet</a> at 15.07.04</div></div><div class="amazlet-detail">David Flanagan <br />オライリージャパン <br />売り上げランキング: 5,722<br /></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873115736/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
