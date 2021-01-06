{"title":"leetcode の Web エディタの Ctrl-D キーバインドを変えたい (未解決)","date":"2021-01-06T23:13:00+09:00","tags":["etc"]}

emacs 系統のキーバインドでは `Ctrl-D` は前方の一文字削除だが、leetcode の Web 上のコードエディタでは一行削除になっている。オプションで emacs キーバインドに変更できるが、ここは変わらない。このキーバインドだけがどうしても煩わしいのでなんとかしたかったが、今の所解決できていない。

[Help Center](https://support.leetcode.com/hc/en-us/articles/360012016874-Start-your-Coding-Practice) によるとエディタには [CodeMirror](https://codemirror.net/) という OSS を使っているらしい。CodeMirror の組み込みのキーバインドは以下のようになっていた。

- デフォルトのキーマップでは、Ctrl-D は一行削除にバインドされている
    - https://github.com/codemirror/CodeMirror/blob/8bc57f76383e62e1a03c7d97c9eac74493fdbedc/src/input/keymap.js#L20
- emacs キーマップがオプションで選択でき、その場合は前方一文字削除にバインドされている
    - https://github.com/codemirror/CodeMirror/blob/36c786bcca35c0650e78ab65ac8afb9d71abb89c/keymap/emacs.js#L304
    - 今回ほしいのはこれ
    - 実際に emacs キーバインドのデモでも確認できる。
        - [CodeMirror: Emacs bindings demo](https://codemirror.net/demo/emacs.html)
    - が、Leetcode のほうではそうはなっていない。

フォーラムにも要望が出ていたが、2 年弱動きが無いようだ。

[Emacs Ctrl\-D keybinding problem \- LeetCode Discuss](https://leetcode.com/discuss/general-discussion/257321/Emacs-Ctrl-D-keybinding-problem)

とりあえず upvote & comment しておいたが、期待薄かな...

サードパーティの cli ツールがあるようなので、試すとしたらこれかな。ただキーバインド一つのためにあまり頑張りたく無いので、悩ましいところ。

- [skygragon/leetcode\-cli: A cli tool to enjoy leetcode\!](https://github.com/skygragon/leetcode-cli)

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4065128447/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://images-na.ssl-images-amazon.com/images/I/51OK2x1LvbL._SX348_BO1,204,203,200_.jpg" alt="問題解決力を鍛える!アルゴリズムとデータ構造 (KS情報科学専門書)" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4065128447/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">問題解決力を鍛える!アルゴリズムとデータ構造 (KS情報科学専門書)</a></div><div class="amazlet-detail">大槻 兼資  (著), 秋葉 拓哉 (監修)<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4065128447/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
