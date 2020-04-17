{"title":"setImmediate での script yielding","date":"2013-04-11T20:38:40+09:00","tags":["javascript"]}

[Script yielding with setImmediate \| NCZOnline](http://www.nczonline.net/blog/2011/09/19/script-yielding-with-setimmediate/) を読んだメモ. 2011年の記事だが, いまだに `setImmediate()` が IE でしかサポートされていないのはなぜなんだろう.

### 導入

`setTimeout()` を使って処理を小さいチャンクに分けるパターンがある. `setTimeout()` で処理を小分けにし, 適切なタイミングで UI スレッドにキューのタスクを処理させる.

<pre><code data-language="javascript">setTimeout(function(){
   //do  something
}, 50)</code></pre>

このコードは 50ms 後に処理を UI スレッドのキューに登録する. 登録された処理は自分の番がまわってくると即座に実行される.効果的に `setTimeout()` を呼ぶと現在の JavaScript のタスクを終わらせることができ, js にブロクされず次の UI 更新を行うことができる.

### 問題

この方法には2つ問題がある. ひとつは時間解像度がブラウザによって異なることだ. 例えば IE8 以前は 15ms の解像度だが Chrome は 4ms だ. setTimeout で待機時間を 0 に設定した場合の実際の実行時間はこの解像度に依存する.

もう一つの問題は電力消費だ. タイマーはラップトップやモバイルデバイスの電池を消耗する. [Chrome はかつて時間解像度を 1ms にしていたが, 電力消費が激しいため 4ms に戻した経緯がある] [1]. [MicroSoft の調査] [2] では時間解像度を 1ms にするとバッテリーの駆動時間が 25% 減ることがわかった. そのためブラウザはインアクティブなタブの時間解像度を 1s に変更したり, IE9 はマシンが電源につながっているかどうかで解像度を変えたりしている.

### setImmediate() 関数

W3C の Web Performance WG は [Efficient Script Yielding] [3] という仕様で `setImmediate()` 関数を策定している. これは `setTimeout()` などと同じ引数を受け取り, UI スレッドがアイドルになると即座に関数をキューに登録する.

<pre><code data-language="javascript">var id = setImmediate(function(){
    //do something
});</code></pre>

`setImmediate()` の戻り値の ID は `clearImmediate()` で処理をキャンセルさせるときに必要だ.

`setImmediate()` の引数の最後に, コールバック関数に渡す引数を与えることもできる.

<pre><code data-language="javascript">setImmediate(function(doc, win){
    //do something
}, document, window);</code></pre>

### 利点

`setImmediate()` を使うことで, ブラウザがタイマーの管理をしなくて良くなる. タイマーに寄るシステムインタラプトを待つ (これが電力を食う) ことなく, シンプルにキューが空になったらタスクを入れるようになる. これは Node.js の `process.nextTick()` と同じだ.

また次の timer tick を待たなくなるので, `setTimeout(fn, 0)` を使うよりも全体の実行時間は短くなる.

### ブラウザサポート

現在では IE10 のみが `setImmediate()` をサポートしている. [Microsoft は `setImmediate()` のデモサイト] [4]を公開している.

### 将来

私は `setImmediate()` に関しては非常に前向きだ. 従来の script yielding のパターンはあくまでハックだ. このオフィシャルな方法を導入することでパフォーマンスへの貢献が期待できる. 早く他のブラウザも対応することを望んでいる.

[1]: http://www.belshe.com/2010/06/04/chrome-cranking-up-the-clock/                 "Chrome: Cranking Up The Clock « Mike's Lookout"
[2]: http://msdn.microsoft.com/en-us/windows/hardware/gg463266                      "Timers, Timer Resolution, and Development of Efficient Code"
[3]: https://dvcs.w3.org/hg/webperf/raw-file/tip/specs/setImmediate/Overview.html   "Efficient Script Yielding"
[4]: http://ie.microsoft.com/testdrive/Performance/setImmediateSorting/Default.html "setImmediate API"
