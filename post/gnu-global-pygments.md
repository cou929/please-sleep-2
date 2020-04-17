{"title":"gnu global と pygments の導入","date":"2015-05-09T23:58:44+09:00","tags":["nix"]}

[gnu global (gtags)](http://www.gnu.org/software/global/global.html) のサポートしている言語が増えたことを知った。正確には [Pygments](http://pygments.org/) という python の syntax higlighter を利用してパースするプラグインが追加されたらしい。というわけで改めてインストールと設定を行った。

### gnu global ができること

- 関数の定義元にジャンプする
- その関数を参照している場所にジャンプする
- 変数などの名前を元にジャンプする

など。ネイティブでは `C, C++, Yacc, Java, PHP4 and assembly` しかサポートされていないが、前述のように Pygments を使ったプラグインによって他の言語も対応できるようになった。(とはいえ [結構前から](http://lists.gnu.org/archive/html/info-global/2014-09/msg00000.html) だけど)

### インストール

    brew install global --with-exuberant-ctags --with-pygments

だけで OK

### 設定

emacs から使う。

まずは global についてくる `gtags.el` を load-path が通ったところに置く。デフォルトでは `/usr/local/Cellar/global/6.4/share/gtags/gtags.el` などにあるはず。

    cp /usr/local/Cellar/global/6.4/share/gtags/gtags.el ~/.emacs.d/elisp/

つぎに `.emacs` に以下を記述

    (setq gtags-prefix-key "\C-c")
    (require 'gtags)
    (require 'anything-gtags)
    (setq gtags-mode-hook
          '(lambda ()
             (define-key gtags-mode-map "\C-cs" 'gtags-find-symbol)
             (define-key gtags-mode-map "\C-cr" 'gtags-find-rtag)
             (define-key gtags-mode-map "\C-ct" 'gtags-find-tag)
             (define-key gtags-mode-map "\C-cf" 'gtags-parse-file)
             (define-key gtags-mode-map "\C-cb" 'gtags-pop-stack)))
    (add-hook 'c-mode-common-hook
              '(lambda()
                 (gtags-mode 1)))

`defined-key` のあたりはキーバインドの指定で、`add-hook` は `tags-mode` を有効にするモードの指定。それぞれ必要に応じて変更する。

基本的にはこれでいいが、[anything-gtags.el](http://emacswiki.org/emacs/anything-gtags.el) なども追加でインストールしておくと便利かもしれない。

### 簡単な使い方

`gtags` コマンドで事前にインデックスを作る。ネイティブでサポートしている言語であればオプションなしで良いが、そうでないなら以下のようにする。

    gtags --gtagslabel=pygments

これで `GPATH`、`GRTAGS`、`GTAGS` というファイルが生成される。

次に該当の読みたいコードを emacs で開く。`C-c s` で定義元へジャンプする。引数を与えなければ、カーソルのある位置の関数名で検索する。同様に、`C-c r` で参照元へ飛ぶ。ジャンプ前の位置に戻るには `C-c b`。という感じ。

### 注意点

pygments プラグインを使ったインデックス作成処理は時間がかかる。自分の場合、[iconv-lite](https://github.com/ashtuchkin/iconv-lite) というモジュールの古いバージョンにある [big5.js](https://github.com/ashtuchkin/iconv-lite/blob/523ad20cf0cf9b4bc4823406e83767c6e753d064/encodings/table/big5.js) という大きな json を含むファイルのパースが一向に終わらなかった。

また言語によってはパースがうまくいかないものもあるらしい。もとがそもそもシンタックスハイライトのためのモジュールなので仕方ない部分がある。

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774150029/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="http://ecx.images-amazon.com/images/I/51M3ahu1q8L._SL160_.jpg" alt="Emacs実践入門　～思考を直感的にコード化し、開発を加速する (WEB+DB PRESS plus)" style="border: none;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774150029/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Emacs実践入門　～思考を直感的にコード化し、開発を加速する (WEB+DB PRESS plus)</a><div class="amazlet-powered-date" style="font-size:80%;margin-top:5px;line-height:120%">posted with <a href="http://www.amazlet.com/" title="amazlet" target="_blank">amazlet</a> at 15.05.09</div></div><div class="amazlet-detail">大竹 智也 <br />技術評論社 <br />売り上げランキング: 48,654<br /></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774150029/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
