{"title":"python の range のようなことを JavaScript で実現するには","date":"2013-10-07T20:09:49+09:00","tags":["javascript"]}

python の `range` 関数や perl の `..` 演算子のような機能は JavaScript のコアでは提供されていない。普通にループを自分で書くか、あるいは underscore の range メソッドを使うのもよいかもしれない。(意外にも jQuery にはそういうメソッドはなさそうだった)

[Underscore.js](http://underscorejs.org/#range)

ここで同僚に変な呪文を教えてもらった。

<pre><code data-language="javascript">Array.apply(null, {length: 5}).map(Number.call, Number)</code></pre>

[Twitter / cowboy: Create a number list in ...](https://twitter.com/cowboy/status/288702647479984128)

 (ちなみに、このツイートを同僚に教えてもらったということ。この人が同僚なわけではない。)

なぜこれが動くのかぱっとは説明できないのであとでちゃんと調べる。

### パフォーマンス

せっかくなのでパフォーマンスを比較してみた。

underscore は素直。最初に空の配列を作り、ループして、それぞれのインデックスに値をセットする。

[underscore/underscore.js at master · jashkenas/underscore](https://github.com/jashkenas/underscore/blob/master/underscore.js#L577-L594)

これと、上記の map を使った実装。あとは range のアルゴリズムをちょっと変えて、インデックスに値を入れるのではなく Array.push を使う方法の 3 つを比較してみた。

[Array initialization with range · jsPerf](http://jsperf.com/array-initialization-with-range)

結果は、なんとなく予想通りだけど、range の実装が一番はやくて、黒魔術の方法は桁違いに遅かった。
