{"title":"GNU Make は空白を含むファイル名をうまく扱えない","date":"2020-05-18T19:45:00+09:00","tags":["nix"]}

[GNU Make](https://www.gnu.org/software/make/) は名前やパスに半角スペースを含むファイルをうまく扱えない。古くから報告されている課題のようだが、この記事の時点でもまだ未解決だった。現状では空白を `\` でエスケープするなりのワークアラウンドが必要。

## 今回の事象

例えば次のような Makefile があるとする。(`post/*.md` をもとにビルドするようなタスクと思ってください)

```
POSTS := $(shell find ./post/*.md -type f)

dist: $(POSTS)
	make clean
	make run

# 例の簡略化のため POSTS を上記の実装にしているが、実際はこの程度のことに find を使うまでも無いと思います。
```

このとき `post/foo bar.md` のようなファイルが混ざると、次のエラーになる。

```
make: *** No rule to make target 'post/foo', needed by 'dist'.  Stop.
```

名前の中の空白がデリミタと勘違いされ、`foo` と `bar.md` というふたつのターゲットが必要だと解釈されている。そのため `post/foo` を作るルールは無いと怒られている。

この挙動は make 3.81 で確認できた (現時点の MacOS で XCode 経由でインストールされたもの)。また make 4.3 ([今年の1月にリリース](https://lists.gnu.org/archive/html/info-gnu/2020-01/msg00004.html) されたみたい。MacOS の場合 [homebrew で手軽にインストールできる](https://github.com/Homebrew/homebrew-core/blob/master/Formula/make.rb)) でも再現した。

[ドキュメント](https://www.gnu.org/software/make/manual/make.html) には、ぱっと見た範囲では make が扱える文字種について説明されている箇所はなかった。通読すれば空白やコロンが実質的に使えないことはわかるかもしれないが、そこまでは見ていない。

## 挙動をもうすこし調べてみる

まず、以降は Makefile のルール各部の名前を以下のように呼びます。

```
# おさらい: Makefile のルール各部の名称
# https://www.gnu.org/software/make/manual/make.html#Rule-Syntax

target: prerequisites
    recipe
```

先に結論をまとめると、

- シェルスクリプトのようにクオートをつけてもうまくパースしてくれない
- バックスラッシュでエスケープすると、target, prerequisites はうまくパースしてくれる
    - recipe にはエスケープなしの文字列が渡ってくるので、それをまた自分でなんとかする必要がある

### クオート

シェルスクリプトの感覚で、クオートでファイル名の範囲を指定することはできないようだ。次の例では `"ba r.txt` は `"ba` と `r.txt"` という 2 要素と認識されてしまっている。

```
# Makefile

fo\ o: "ba r.txt"
	@echo $?
	@echo $@
	ls -l "$?"
```
```
$ make
make: *** No rule to make target `"ba', needed by `fo o'.  Stop.
```

### バックスラッシュでのエスケープ

次のようにバックスラッシュでエスケープすると、target, prerequisites で意図通りに動くようだ。(エスケープされた空白はファイルのデリミタとはみなされない)。ただ recipe にはバックスラッシュなしの文字列が届くので、適宜クオートなどで囲んでからシェルに渡したほうがよい。

```
# Makefile

fo\ o: ba\ r.txt
	@echo $?
	@echo $@
	ls -l "$?"
```

```
$ ls
Makefile  ba r.txt

$ make
ba r.txt
fo o
ls -l "ba r.txt"
-rw-r--r--  1 cou929  staff  0  5 18 15:01 ba r.txt
```

wildcard を使うと自分でエスケープする必要はないらしい。recipe にはエスケープ無しで渡ってくるので自分でなんとかする必要がある。

```
# Makefile

fo\ o: *.txt
	@echo $?
	@echo $@
	ls -l $?
```
```
$ make
ba r.txt baz.txt
fo o
ls -l ba r.txt baz.txt
ls: ba: No such file or directory
ls: r.txt: No such file or directory
-rw-r--r--  1 cou929  staff  0  5 18 15:14 baz.txt
make: *** [fo o] Error 1
```

shell 関数で変数に結果を入れる場合、自分でエスケープする必要がある。

```
# Makefile

SRCS := $(shell ls *.txt | sed -e 's/ /\\ /g')

fo\ o: $(SRCS)
	@echo $?
	@echo $@
	ls -l $?
```

### 実装

(あまり自信はなく、全然違うかもしれません)

実装はおそらくこのあたり (GitHub のミラーより)。

https://github.com/mirror/make/blob/453334882668f7e21a85491965f9d369cdd762c4/src/read.c#L2881

Makefile のパースを行う eval 関数の中で、バックスラッシュの場合はカーソルを前に進めている。クオート系の処理は登場しないので、調査した挙動通りの実装なんだと思われる。関数系は function.c に実装があるので、また別だと思われる。

## savannah での報告

この問題は 2002 年に報告されていて、ときどきコメントがついているものの、未解決のようだ。

[make \- Bugs: bug \#712, GNU make can't handle spaces in\.\.\. \[Savannah\]](https://savannah.gnu.org/bugs/index.php?712)

- 名前やパスに空白を許容する変更は全体に影響する
- 後方互換性を維持しながらこの変更を入れるのはなかなか大変な作業
    - そういうフラグを追加する提案もされていたが、きれいなデザインではないので腰が重くなるのもわかるなと個人的には思った

という理由のようだ。

## 回避策

今回の自分のケースは、

- prerequisites に空白が含まれうるファイルが混在している
- recipe で prerequisites のファイル名ひとつひとつを使うことがない

ので、prerequisites をリストアップする部分で空白をエスケープする方法で対応した。

```
POSTS := $(shell find ./post/*.md -type f | sed -e 's/ /\\ /g')

dist: $(POSTS)
	make clean
	make run
```

また [BTS](https://savannah.gnu.org/bugs/index.php?712) では他にも回避策が紹介されていた。

空白を `+` に置き換え、使う際に subst でもとに戻す方法 (自分のアプローチに近いもの)。

```
SRC_NOSP=AA+Mgr/cpp1.cpp AA+Mgr/cpp2.cpp
$(info total $(words $(SRC_NOSP)) source files!)
OBJ_NOSP=$(patsubst %.cpp,$(ObjDir)/%.o,$(SRC_NOSP))
SRC=$(subst +,\ ,$(SRC_NOSP))
OBJ=$(subst +,\ ,$(OBJ_NOSP))
all: $(OBJ)
	...
```

make の関数で問題が起こっている場合、かわりに shell を使う方法。

```
# Don't use:
File=AA\ Mgr/cpp1.cpp
SrcDir=$(dir $(File))

# insteads,
SrcDir=$(shell dirname '$(File)')
```

空白を含まないファイル名の symlink を作成しそちらを使う方法。

```
.PRECIOUS: %.link
%.link:
ln -s "$$(echo $@ | sed -e 's/_dot_/./g' -e 's/_slash_/\//g' -e 's/_space_/ /g' -e 's/.link$$//')" $@

# should depend on ../Plugin Source/metadata/monitor.xml but gnumake doesn't support spaces!
.DELETE_ON_ERROR:
%.sql: %.sql.haml Makefile $(wildcard _*.sql.haml) _dot__dot__slash_Plugin_space_Source_slash_metadata_slash_monitor.xml.link
```

また、可能な場合はそもそも空白が含まれないように前処理することも考えられる。Makefile にそのようなルールを追加する場合は次のような感じ。(そもそもこれが可能なら最初から空白が含まれるファイル名はつけないはずだが...)

```
.PHONY: ensure-no-space
ensure-no-space:
	for f in ./post/*.md; do \
		echo $$f; \
		mv "$$f" `echo $$f | sed -e 's/ /_/g'`; \
	done

# 複数行のコマンドをバックスラッシュでエスケープしたり、`$$` とエスケープするなどのお作法に注意。
```

とはいえどれも面倒なので、本質的には本体で対応されると嬉しい。少なくともドキュメントにもう少し明示的に書かれているとハマる時間を減らせたかもしれない。慣習的に自分から進んで空白を含むファイル名をつけることはないが、unix のファイルシステムが空白を許容している以上、make が対応してくれてもいいのになとも思う。

## 参考

- [make \- Bugs: bug \#712, GNU make can't handle spaces in\.\.\. \[Savannah\]](https://savannah.gnu.org/bugs/index.php?712)
- [GNU make](https://www.gnu.org/software/make/manual/make.html)
- [GNU Makeと空白を含むファイル名 \- 檜山正幸のキマイラ飼育記 \(はてなBlog\)](https://m-hiyama.hatenablog.com/entry/20140920/1411186147)
- [makeで半角スペースを含むファイルを扱いたい\.\.\.。 \- ここにタイトルを入力\|](https://goth-wrist-cut.hatenadiary.org/entry/20080709/1215606353)

<div class="amazlet-box" style="margin-bottom:0px;">
    <div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;">
        <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873112699/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">
            <img src="https://images-fe.ssl-images-amazon.com/images/I/81IQYnz491L._AC_UL115_.jpg" alt="GNU Make 第3版 (日本語)" style="border: none;" width=113 />
        </a>
    </div>
    <div class="amazlet-info" style="line-height:120%; margin-bottom: 10px">
        <div class="amazlet-name" style="margin-bottom:10px;line-height:120%">
            <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873112699/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">
                GNU Make 第3版 (日本語)
            </a>
        </div>
        <div class="amazlet-detail">
            Robert Mecklenburg (著), 矢吹 道郎 (監訳) (翻訳), 菊池 彰 (翻訳) <br />オライリージャパン <br />
        </div>
        <div class="amazlet-sub-info" style="float: left;">
            <div class="amazlet-link" style="margin-top: 5px">
                <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873112699/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a>
            </div>
        </div>
    </div>
    <div class="amazlet-footer" style="clear: left"></div>
</div>
