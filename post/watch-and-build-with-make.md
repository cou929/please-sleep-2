{"title":"make でナイーブなファイル変更監視とビルド","date":"2020-04-17T20:47:00+09:00","tags":["tips", "nix"]}

Makefile を使っているプロジェクトの場合、以下のようにしておくと簡易的なファイル変更監視 & ビルドの仕組みになる。

```
while true; do make build --silent; sleep 1; done
```

要は定期的に make コマンドを実行しているだけ。

めちゃくちゃ単純なので、linux 系であればほぼどんな環境でも何もせずに動くのと、見れば何をやっているか一目瞭然なのがメリット。一方で毎秒ポーリングしているので無駄が多いのがデメリットと言えると思う。

## Makefile のおさらい

Makefile の基本的なルールの書き方は、

```
ターゲット (targets): 依存するファイル (prerequisites)
    依存からターゲットを生成するためのコマンド (recipe)
```

ターゲットは呼び出しごとに毎回ビルドされるわけではなく、以下のいずれかの条件にマッチした時にだけビルドされる。

- ターゲットが無い
- またはターゲットよりも依存するファイルの方が新しい

> How to decide whether foo.o is out of date: it is out of date if it does not exist, or if either foo.c or defs.h is more recent than it.
> 
> [GNU make](https://www.gnu.org/software/make/manual/make.html#Rules)

こうしたルールを適切に設定しておくと、あとはそれを定期的に呼び出すだけで、簡易的な変更監視&ビルドの仕組みになる。

## [cou929/please-sleep-2](https://github.com/cou929/please-sleep-2) での例

[cou929/please-sleep-2](https://github.com/cou929/please-sleep-2) はこのブログのコンテンツを管理しているリポジトリ。中にはブログのコンテンツとそれを変換するヘルパーツールが入っている。

やりたいことは、ブログのコンテンツを追加したりヘルパーツールを修正した際に、自動でビルド (ヘルパーツールをコンパイルし、それを使ってコンテンツを再生成) してほしいというもの。特にブログ本文は生成した html を見ながら何度も推敲するので、変更検知してくれるとありがたい。

大した処理をしているわけでもないので、特別な仕組みを導入せずにこれを実現したかった。もともと Makefile を使っていたこともあり、今回の方法を採用した。

[Makefile](https://github.com/cou929/please-sleep-2/blob/f7501e7642d7729aa0c61d8a83b4cc752e98663b/Makefile) は次のようにした。

```
SRCS := $(shell find . -type f -name '*.go')  # ヘルパーツールのソースファイル
POSTS := $(shell find ./post -type f)         # ブログコンテンツ

dist: $(POSTS) $(SRCS)  # 結果を ｀dist/` に格納する
	make clean
	make run
	make asset
```

成果物は `dist/` ディレクトリに全部突っ込まれる。make に dist と依存ファイルの SRCS, POSTS を見てもらっている。

ついでに `watch` というコマンドも追加し、`make dist` ルールを定期的に呼び出すようにした。

```
.PHONY: watch
watch:
	while true; do \
		make dist --silent; \
		sleep 3; \
	done
```

これで `make watch` とするだけでやりたいことができた。

この仕組みはまだ使い始めたばかりだが、この程度の規模だと十分実用に耐えられそう。

## 発展

この方法は make が毎回対象の全ファイルをポーリングするような動きになるので、当然無駄が多い。問題になる場合は、変更があった場合にだけ通知を受ける仕組みにしたい。

[今回参考にした記事](https://fromanegg.com/post/2015/08/26/automatically-build-files-when-they-change-with-make/) では [inotify](https://ja.wikipedia.org/wiki/Inotify) ベースのツール ([inotifywait](https://linux.die.net/man/1/inotifywait), [fswatch](https://emcrisostomo.github.io/fswatch/)) を使う方法が解説されていた。

```
.PHONY watch
watch:
	while true; do \
		inotifywait -qr -e modify -e create -e delete -e move app/src; \
		make build; \
	done
```

> inotify (inode notify) とは、ファイルシステムへの変更を通知するようファイルシステムを拡張して、その変更をアプリケーションに報告するLinuxカーネルサブシステムである。
> 
> [inotify \- Wikipedia](https://ja.wikipedia.org/wiki/Inotify)

そもそもがこの手のツールはだいたい inotify に乗っかっているらしい。

## 参考

- [GNU make](https://www.gnu.org/software/make/manual/make.html)
- [Automatically build files when they change with Make · From An Egg](https://fromanegg.com/post/2015/08/26/automatically-build-files-when-they-change-with-make/)
- [inotify \- Wikipedia](https://ja.wikipedia.org/wiki/Inotify)
- [inotifywait\(1\) \- Linux man page](https://linux.die.net/man/1/inotifywait)
- [fswatch](https://emcrisostomo.github.io/fswatch/)

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B075ST51Y5/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/41HbBZyeJ3L.jpg" alt="ふつうのLinuxプログラミング 第2版　Linuxの仕組みから学べるgccプログラミングの王道" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B075ST51Y5/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">ふつうのLinuxプログラミング 第2版　Linuxの仕組みから学べるgccプログラミングの王道</a></div><div class="amazlet-detail">青木 峰郎  (著)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B075ST51Y5/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
