{"title":"js の strict mode おさらい","date":"2012-12-11T22:06:14+09:00","tags":["javascript"]}

スクリプト or 関数の先頭に `'use strict';` と書いておくと, 処理系が厳密にコードをチェックしてアラートを出してくれる.

簡単な例. strict mode では `with` は使っちゃいけない. 以下の html をブラウザで開くと,

    <html>
    <body>
    <script>
    'use strict';
    
    with (localhost) {
        alert(href);
    }
    </script>
    </body>
    </html>

chrome だとこんなエラーを出してくれる.

    Uncaught SyntaxError: Strict mode code may not include a with statement

どういうものが strict mode だとエラーになるのかは, 導入は zakas さんのこの記事がよさげ.

[It’s time to start using JavaScript strict mode \| NCZOnline](http://www.nczonline.net/blog/2012/03/13/its-time-to-start-using-javascript-strict-mode/)

より詳細はこちらで仕様にあたるといい.

[ECMA-262 » ECMA-262-5 in detail. Chapter 2. Strict Mode.](http://dmitrysoshnikov.com/ecmascript/es5-chapter-2-strict-mode/)

書き方にも若干注意点があるけれども, 基本的には strict mode にしたいスコープの先頭に `'use strict';` を書いておけば OK. 前にコメントはあってもいいけどコードはだめ. 詳しくは niw さんのブログ

["use strict" - blog.niw.at](http://blog.niw.at/post/26687866336)
