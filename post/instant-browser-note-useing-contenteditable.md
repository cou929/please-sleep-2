{"title":"contenteditable でブラウザでメモを書く","date":"2013-01-30T19:55:23+09:00","tags":["tips, browser"]}

[Jose Jesus Perez Aguinaga : One line browser notepad](https://coderwall.com/p/lhsrcq)

ここで紹介されていた, Data URI と contenteditable を組み合わせてブラウザのタブを即席のメモ帳にするアイデアが面白かった. やりかたはブラウザのロケーションバーに以下をいれるだけ:

<pre><code data-language="python">data:text/html, &lt;html contenteditable&gt;</code></pre>

黒背景にしたかったので, すこしだけスタイルを足して使っている. 以下をブックマークレットとしてブラウザに保存した.

<script src="https://gist.github.com/4672258.js"></script>

ちなみにこれで書いたメモは html で保存できる. 中身はぐちゃぐちゃの html なのでスクリプトから読んだりするのは難しいけど, ブラウザで開けば読むことはできる.

`contenteditable` 属性を DOM のある要素に指定すると, その中身が自由に編集できるようになる. 変更された DOM の中身は js から読める. ブラウザ上での wysiwyg エディタ開発を想定している機能らしい.

[The contenteditable attribute \| HTML5 Doctor](http://html5doctor.com/the-contenteditable-attribute/)

もういっこちなみに, data uri の syntax

    data:[<mediatype>][;base64],<data>
