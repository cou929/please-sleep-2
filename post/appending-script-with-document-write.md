{"title":"jquery の html() と append を追う","date":"2012-10-17T13:03:00+09:00","tags":["javascript"]}

script 要素を含む dom を jquery の `$.html()` で動的に挿入するようなケースで意図した動作をしなかった.

- 動的に追加された script 要素のなかの js のコードや外部 js が実行されることを期待
  - それらのコードは他の人が書いたものでいろんな種類があり中身は把握しきれない
- 実際には新しいドキュメントにそれらの script が実行された結果がでたりする
  - document.write を使ってるスクリプトはあるからそれのせいっぽい

### $.html()

- $.html() は manipuration.js
  - 基本的に innerHTML につっこんでいる
  - ただし引数の文字列に script などの要素が入っている場合は $.append() を呼ぶ

            rnoInnerhtml = /<(?:script|style|link)/i,
            ....

            if ( typeof value === "string" && !rnoInnerhtml.test( value ) &&
                  ... ) {
            
                ...
            
                try {
                            ...
                            elem.innerHTML = value;
                            ...
                    }
            
                // If using innerHTML throws an exception, use the fallback method
                } catch(e) {}
            }
            
            if ( elem ) {
                this.empty().append( value );
            }

    - script とか style とかは評価してあげる必要があるから.
      - 逆に言うと innerHTML で script 要素などを動的に追加してもなかのコードは実行されない

### $.append()

appendChild してるだけっぽい

        append: function() {
            return this.domManip(arguments, true, function( elem ) {
                if ( this.nodeType === 1 ) {
                    this.appendChild( elem );
                }
            });
        },

`domManip` は文字列から dom を作っている

### appendChild と innerHTML とで script 要素などの実行の違い

実験してみると予想通り innerHTML だと script タグの内容は実行されずそのままノードに置き換わる. appendChild の場合は中身が実行される

#### test.html

    <script type="text/javascript" src="test.js"></script>
    <div id="target"></div>

#### test.js

    (function() {
        function inner_html_test() {
            document.querySelector('#target').innerHTML = '<script type="text/javascript">console.log("with innerHTML")</script>';
        }
    
        function append_child_test() {
            var script = document.createElement('script');
            script.setAttribute('type', 'text/javascript');
            var content = document.createTextNode('console.log("with appendChild")');
            script.appendChild(content);
            document.querySelector('#target').appendChild(script);
        }
    
        window.addEventListener('load', function() {
            inner_html_test();
            window.setTimeout(append_child_test, 1000);
        });
    })();

#### 結果
1 秒後に 'with appendChild' と表示される

### document.write
script 要素を appendChild で動的挿入した場合でも, その中身が document.write() だったら実行されない.

#### test.html

    <script type="text/javascript">
    window.addEventListener('load', function() {
        var script = document.createElement('script');
        var url = 'async_content.js';
        script.setAttribute('src', url);
        document.querySelector('body').appendChild(script);
    });
    </script>

#### async_content.js

    console.log('in async_content.js');
    document.write('aaaaaaaaaaaaaaaaaaaaaaaaaa');
    console.log('after document.write');
    
    window.setTimeout(function() {
        console.log('timeouted');
    }, 1000);

#### 結果

    in async_content.js
    after document.write
    timeouted

    // html 上にはなにもでていない

例えば async_content.js のなかで write する前に `document.open()` すると html の内容が `aaaaaaaaaaa...` で全部上書きされる. そうでなくても新しいドキュメントに `aaa...` がでるような気がしていたんだけど, どうなんだろう

### まとめ

- document.write するような外部スクリプトを動的に追加するのは面倒
  - あるていどフォーマットが決まっていればいいが, 今回はいろいろなスクリプトが大量にあるのでむずかしい
- `document.write` と `document.open/close` の挙動がちょっと理解不足
