{"title":"JavaScript での Mixin メモ","date":"2012-04-29T11:05:03+09:00","tags":["javascript"]}

[peter.michaux.ca - Mixins and Constructor Functions](http://peter.michaux.ca/articles/mixins-and-constructor-functions)

- Mixin 自体はなるほど
  - よく使うものをまとめておいて, あとで別のオブジェクトにその能力を付与できる
  - Spine とかでもやってた
- コメント欄で Angus Call が書いてた functional な方法の方がすっきりしているように感じた
  - [A fresh look at JavaScript Mixins « JavaScript, JavaScript…](http://javascriptweblog.wordpress.com/2011/05/31/a-fresh-look-at-javascript-mixins/)
  - Mixin オブジェクトじゃなくて関数にしておく
  - なかで this にプロパティやメソッドをつけてく
  - これを Mixin したい時は `MixinFn.call(Obj)` というふうにやる
- Addy Osmani のパターン集でもちゃんと取り上げられていた
  - [Essential JavaScript And jQuery Design Patterns](http://addyosmani.com/resources/essentialjsdesignpatterns/book/#mixinpatternjavascript)
  - こっちは `augment()` という関数で mixin してるけど, `call()` の方法のほうがスマートな気がする
    - あーでも関数名的に意図が明確になるのは `augment()` か
    - でも処理の流れはこっちの方が自然
  - cache のくだりと curry 化を使って option を使えるようにする話が面白い
  - cache バージョンの例

            var asRectangle = (function() {
              function area() {
                return this.length * this.width;
              }
              function grow() {
                this.length++, this.width++;
              }
              function shrink() {
                this.length--, this.width--;
              }
              return function() {
                this.area = area;
                this.grow = grow;
                this.shrink = shrink;
                return this;
              };
            })();
            
            var RectangularButton = function(length, width, label, action) {
              this.length = length;
              this.width = width;
              this.label = label;
              this.action = action;
            }
            
            asButton.call(RectangularButton.prototype);
            asRectangle.call(RectangularButton.prototype);
            
            var button3 =
              new RectangularButton(4, 2, 'delete', function() {return 'deleted'});
            button3.area(); //8
            button3.grow();
            button3.area(); //15
            button3.fire(); //'deleted'

- そもそも Mixin のメリットは
  - 同じようなことは prototype をいじればできるが
    - その子クラスにも影響がおよぶ
    - 他のインスタンスに影響が及ぶかどうか
            - mixin の場合, mixin したあとにもとの Mixin 元にメソッドを追加したりしても, 事前に Mixin したオブジェクトには影響がない (prototype だとあとから追加したら全部変わる)
  - Further reading
    - [まつもと直伝　プログラミングのオキテ 第3回（1） - まつもと直伝 プログラミングのオキテ：ITpro](http://itpro.nikkeibp.co.jp/article/COLUMN/20050912/220974/?ST=oss)
    - [まつもと直伝　プログラミングのオキテ 第3回（2） - まつもと直伝 プログラミングのオキテ：ITpro](http://itpro.nikkeibp.co.jp/article/COLUMN/20050915/221233/?ST=oss)
    - [まつもと直伝　プログラミングのオキテ 第3回（3） - まつもと直伝 プログラミングのオキテ：ITpro](http://itpro.nikkeibp.co.jp/article/COLUMN/20050915/221232/)
  - js の場合 prototype を使った単一継承になるので, prototype にいろいろ突っ込むのはそもそも mixin みたいな挙動になると考えてよさそう
    - なのでここで言ってる mixin と prototype chaining の違いは, 他のインスタンスや下のクラスに影響がでるかどうか (そのインスタンスにだけ追加するか prototype に追加するかどうか) しかないように思う
  - angus call のキャッシュバージョンの functional mixin の場合, クロージャの中のプロパティは共有になっている
    - 一方キャッシュしない mixin の場合は毎回新しいプロパティになる
  - mixin をどうすべきかは以下を考慮してプログラマが決める必要がある
    1. 変更が他のクラス・インスタンスに及んで欲しいのか (そのクラスにそもそもあるべき機能なのか, あるインスタンスだけにくっつけたい機能なのかの見極め)
    2. パフォーマンスと分離性. 何度も呼ぶならパフォーマンス, 共有されても困らないならクロージャでキャッシュ
