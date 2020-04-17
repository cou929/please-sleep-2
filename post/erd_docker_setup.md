{"title":"テキストベースの ER 図作成環境の準備","date":"2018-01-09T20:07:01+09:00","tags":["nix"]}

ER 図をテキストベースで作成する方法を調べていたところ、[BurntSushi/erd](https://github.com/BurntSushi/erd) という Haskell 製のツールがあった。Docker を使ってこのツールの環境を整えたのでやったことをメモしておく。

### Dockerfile

以下の Dockerfile を準備した。Docker の環境は整っている前提。

	FROM haskell:8

	RUN apt-get update && apt-get install -y graphviz
	RUN cabal update && cabal install graphviz parsec
	RUN git clone https://github.com/BurntSushi/erd.git
	WORKDIR /erd
	RUN cabal configure
	RUN cabal build
	RUN cp dist/build/erd/erd /usr/bin

	ENTRYPOINT [ "erd" ]
	CMD [ "--help" ]

- 最近はあまりアクティブでないプロジェクトなので、README 通りの方法では導入できなかったし、p-r も溜まっている
- [Dockerfile を追加する p-r](https://github.com/BurntSushi/erd/pull/16) が出ていたのでそちらをベースにした
- CMD でヘルプメッセージを表示
- ENTRYPOINT でコンテナの erd コマンドをホスト側から呼び出せるようにしている

どこかのディレクトリに上記の Docker ファイルを準備して `docker run` すればよい。

#### 使い方の例


	# 入力ファイルを指定
	# `docker run -v` で同期したディレクトリを `erd -i` オプションに渡しているが、他に良い方法はある?
	$ docker run --rm -v $(pwd):/tmp erd -i /tmp/sample.er > result.pdf
	# 出力フォーマットの変更
	$ docker run --rm -v $(pwd):/tmp erd -i /tmp/sample.er -f png > result.png

### er ファイルと生成される図の例

	# Entities are declared in '[' ... ']'. All attributes after the entity header
	# up until the end of the file (or the next entity declaration) correspond
	# to this entity.
	[Person]
	*name
	height
	weight
	+birth_location_id

	[Location]
	*id
	city
	state
	country

	# Each relationship must be between exactly two entities, which need not
	# be distinct. Each entity in the relationship has exactly one of four
	# possible cardinalities:
	#
	# Cardinality    Syntax
	# 0 or 1         0
	# exactly 1      1
	# 0 or more      *
	# 1 or more      +
	Person *--1 Location

このファイルから以下の図を生成できる。

	docker run --rm -v $(pwd):/tmp erd -i /tmp/simple.er -f png > erd_sample.png

![](/images/erd_sample.png)

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774189200/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://images-fe.ssl-images-amazon.com/images/I/61iX9PgNq7L._SL160_.jpg" alt="WEB+DB PRESS Vol.98" style="border: none;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774189200/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">WEB+DB PRESS Vol.98</a><div class="amazlet-powered-date" style="font-size:80%;margin-top:5px;line-height:120%">posted with <a href="http://www.amazlet.com/" title="amazlet" target="_blank">amazlet</a> at 18.01.09</div></div><div class="amazlet-detail">丸山 晋平 前佛 雅人 横田 真俊 小原 薫 小笠原 空宙 高橋 征義 牧 大輔 大沢 和宏(Yappo) 久田 真寛 のざき ひろふみ うらがみ 池田 拓司 ひげぽん 遠藤 雅伸 海野 弘成 はまちや2 竹原 日高 正博 <br />技術評論社 <br />売り上げランキング: 39,429<br /></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774189200/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
