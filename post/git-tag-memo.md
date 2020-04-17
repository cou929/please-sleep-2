{"title":"git tag のメモ","date":"2015-03-07T20:10:37+09:00","tags":["git"]}

tag はコミットにつけることができる。コミットの human friendly な別名

- `git tag <tag name>` で HEAD にタグ付け
- `git tag` ですべてのタグを参照
- `git log —decorate` でタグを含めたログをだす

![](images/git-tag-log-decorate.png)

- `git tag -d <tag name>` でタグの削除
- `git tag -a` でタグにメッセージを付与できる
  - コミットメッセージと同様に `git tag -am ‘message’` でエディタを開かずに追加もできる
- `git push origin master —tags` などとするとタグをリモートに反映させる
- tig はデフォルトでダグも表示してくれるので便利

![](images/git-tag-tig.png)

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4274068641/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="http://ecx.images-amazon.com/images/I/51WQ7GsnOZL._SL160_.jpg" alt="Gitによるバージョン管理" style="border: none;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4274068641/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Gitによるバージョン管理</a><div class="amazlet-powered-date" style="font-size:80%;margin-top:5px;line-height:120%">posted with <a href="http://www.amazlet.com/" title="amazlet" target="_blank">amazlet</a> at 15.03.07</div></div><div class="amazlet-detail">岩松 信洋 上川 純一 まえだこうへい 小川 伸一郎 <br />オーム社 <br />売り上げランキング: 226,062<br /></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4274068641/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
