{"title":"HTML attribute と DOM property","date":"2013-10-03T22:40:31+09:00","tags":["javascript"]}

[jQuery Core 1.9 Upgrade Guide \| jQuery](http://jquery.com/upgrade-guide/1.9/#attr-versus-prop-)

jQuery 1.9 以降では `jQuery.attr()` の挙動がかわり、ブール値を扱うような属性には `jQuery.prop()` を使うよう推奨されてる。このあたりがよく理解できていなかったため整理した。

### 挙動の変化

たとえばラジオボタンを動的にチェックするケース。次の HTML を jQuery で操作する。

<pre><code data-language="html">&lt;html&gt;
  &lt;head&gt;
    &lt;title&gt;test&lt;/title&gt;
    &lt;script src="jquery.js"&gt;&lt;/script&gt;
  &lt;/head&gt;
  &lt;body&gt;
    &lt;form method="GET" action="/"&gt;
      &lt;input type="radio" id="target1" name="test" value="1"/&gt;
      &lt;input type="radio" id="target2" name="test" value="2"/&gt;
      &lt;input type="submit"/&gt;
    &lt;/form&gt;
  &lt;/body&gt;
&lt;/html></code></pre>

jQuery 1.9 以前であれば次のようにしていただろう。

<pre><code data-language="javascript">$("#target1").attr("checked", "checked");
$("target2").attr("checked", "checked");
$("target1").attr("checked", "checked");
$("target2").attr("checked", "checked");</code></pre>

このようにすると 2 つのラジオボタンが交互にトグルしながら選択される。あるいは行儀よく、`removeAttr()` で checked 属性をクリアしてから `attr` でセットするかもしれないが、挙動は同じだ。

jQuery 1.9 以降ではこの挙動がかわる。最初の 2 行でそれぞれの要素に checked 属性がつき、UI 上選択もされるが、残り 2 行では何もおこらない。

この場合 `jQuery.prop()` を使うのが正しい。次のように書くと意図したとおりラジオボタンが交互に選択されるはずだ。

<pre><code data-language="javascript">$("#target1").prop("checked", true);
$("target2").prop("checked", true);
$("target1").prop("checked", true);
$("target2").prop("checked", true);</code></pre>

ちなみに prop() は jQuery 1.6 で導入されている。

### attr と prop、または attribute と property

jQuery の `attr()` メソッドと `prop()` メソッドはそれぞれ HTML の attribute と property を操作するものだ。

- [.attr() \| jQuery API Documentation](http://api.jquery.com/attr/)
- [.prop() \| jQuery API Documentation](http://api.jquery.com/prop/)

attribute と property とはなにか。HTML に記述されている属性が attribute、HTML を解釈し構築した DOM 要素のプロパティが property と考えればよさそうだ。

以下の StackOverflow で回答されていた例がわかりやすい。

[Properties and Attributes in HTML - Stack Overflow](http://stackoverflow.com/questions/6003819/properties-and-attributes-in-html)

つぎのような input 要素があるとする。

<pre><code data-language="html">&lt;input type="text" value="Name:"&gt;</code></pre>

この場合 HTML attribute としては type, text が指定されており、またこれをパースした DOM の HTMLInputElement オブジェクトの property としても type と value ができる。それぞれの値は HTML に定義されているものと同様、type は "text"、value は "Name:" だ。

    input.getAttribute("value");  // Name:

ここでユーザーがテキストボックスになにか値を入力したとする。この場合、DOM の value property の内容は、ユーザーが入力した値になる。この値は HTML の attribute には反映されず、`getAttribute("value")` の結果は "Name:" のままだ。attribute は初期値、property にはその後動的に変更された値が反映される。

    input.value;  // dynamic value

checked 属性を例に出すと、厳密に言うと HTML の checked attribute の値は `defaultChecked` というプロパティに対応している。あくまでデフォルト値であり、動的に変更される値には対応していないということになる。

jQuery としては場面に応じた attr と prop の使い分けを推奨している。selectedIndex, tagName, nodeName, nodeType, ownerDocument, defaultChecked, defaultSelected といった HTML 属性にはないプロパティの取得、及び value や checked, selected などを動的に変更したい場合。これらのケースでは `prop` を使用する必要がある。

[.prop() \| jQuery API Documentation](http://api.jquery.com/prop/#prop-propertyName)

ある要素の属性を表すものが attribute と property として 2 つあり、しかもそれらが sync していない。確かにこれは結構いけてない仕様だとは思うし、現に今回はそのせいで混乱した。しかし、例えば一部のブラウザの挙動の違いとか後方互換性のためにこのような仕様になっているとか、動的な property の変更を HTML の attribute にまで伝播させるのはコストが大きいとか、そういう理由があるのかなと想像する。ブラウザ側がそういう挙動をする以上、 jQuery が 2 つのメソッドを提供するのも仕方がないと思う。それに selectIndex のような HTML には登場しない属性は prop で扱うのが妥当だ。

### Boolean Attributes

option 要素の select 属性など、真偽値を値としてもつものは Boolean Attribute と呼ばれる。

[On SGML and HTML](http://www.w3.org/TR/html4/intro/sgmltut.html#h-3.3.4.2)

spec によると、属性が与えられていた場合 true、なければ false として解釈される。つまり selected、selected="selected"、selected=""、selected="false"などはすべて true として扱われる。属性がない場合のみ falseだ。

html としては selected とだけ、値なしで書いておいても valid だが、この記法は (現在ではさほど気にしなくてもよいが) xml として invalid だ。よって値として属性名をいれる selected="selected" という書き方がデファクトになっている。

[HTML - Why boolean attributes do not have boolean value? - Stack Overflow](http://stackoverflow.com/questions/7089584/html-why-boolean-attributes-do-not-have-boolean-value)

#### jQuery.attr の実装の変化

jQuery 1.8.3 の jQuery.attr メソッドについて、操作対象が boolean attribute の場合は最終的に次の処理が行われる。

<pre><code data-language="javascript">	set: function( elem, value, name ) {
	var propName;
	if ( value === false ) {
		// Remove boolean attributes when set to false
		jQuery.removeAttr( elem, name );
	} else {
		// value is true since we know at this point it's type boolean and not false
		// Set boolean attributes to the same name and set the DOM property
		propName = jQuery.propFix[ name ] || name;
		if ( propName in elem ) {
			// Only set the IDL specifically if it already exists on the element
			elem[ propName ] = true;
		}

		elem.setAttribute( name, name.toLowerCase() );
	}
	return name;
}</code></pre>

このように property を変更した上で setAttribute している。よって attr を呼び出すだけでラジオボタンのチェックがトグルしていたことになる。

その後次の修正が入り、attr での set は単に setAttribute だけをするようになった。

[2.0: Remove getSetAttribute and getSetInput and oldIE attroperties hooks · 3f66e92 · jquery/jquery](https://github.com/jquery/jquery/commit/3f66e928c816de49e4c79a3375524c2d0c4e56ce)

attr メソッドはその名の通り単に HTML の attribute を設定するだけの役割になり、プロパティの変更には prop メソッドが準備された。そのため冒頭のような挙動になった。

### まとめ

jQuery の attr と prop メソッドについて理解するには、HTML の attribute と DOM の property を理解する必要があった。混乱を招く仕様なのでなんとかしてほしいところだが、現状こうなっているものは仕方がないと思う。
