{"title":"Markdown interpreter on Ruby","date":"2012-02-25T19:07:29+09:00","tags":["ruby"]}

StackOverflow にいい質問があった

[rake - Better ruby markdown interpreter? - Stack Overflow](http://stackoverflow.com/questions/373002/better-ruby-markdown-interpreter)

1. [Old BlueCloth (Version 1 系?)](http://deveiate.org/projects/BlueCloth)

   Pure Ruby. ただし遅い. 同じ pure ruby 実装である後述の maruku よりも[遅いというベンチ](http://deveiate.org/projects/BlueCloth)が出ている.
2. [Maruku](https://github.com/nex3/maruku)

   Pure Ruby. BlueCloth よりもパフォーマンスが良く, 代表的な markdown の ruby 実装だったようだ. 最後のコミットが 2010 で, もうメンテされていない模様.
3. [RDiscount](https://github.com/rtomayko/rdiscount)

   Non Pure Ruby. コア部分が C 実装のため高速.
4. [New BlueCloth (Version 2 系?)](http://deveiate.org/projects/BlueCloth)

   Non Pure Ruby. 高速化したが RDiscount ベースで pure ruby では無くなった.
5. [kramdown](http://kramdown.rubyforge.org/)

   Pure Ruby. Pure Ruby 実装の中では一番速いとのこと. 現在でもメンテされているようだ.
6. [BlueFeather](http://ruby.morphball.net/bluefeather/index_en.html)

   Pure Ruby. kramdown 曰く kramdown のほうが速いらしい. 最終コミットが 2010 年末だったので, こちらもメンテが止まっていそう.
7. [RedCarpet](https://github.com/tanoku/redcarpet/tree/master/ext/redcarpet)

   Non Pure Ruby. by GitHub. GitHub Flavered Markdown を使いたいならこれ.

というわけで結論は,

- pure ruby がいいなら kramdown
- 速さが欲しいなら BlueCloth or RedCarpet
- GitHub Flavered Markdown を使いたいなら RedCarpet

ということのようです
