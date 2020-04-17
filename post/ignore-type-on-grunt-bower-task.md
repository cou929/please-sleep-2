{"title":"grunt-bower-task の exportsOverride で type 名のディレクトリを作らないようにする","date":"2014-09-21T12:58:19+09:00","tags":["javascript"]}

<img src="https://camo.githubusercontent.com/8a2024183152023c85dc7124365c1afb721450a4/687474703a2f2f626f7765722e696f2f696d672f626f7765722d6c6f676f2e706e67" width="200"/>

結論としては以下の 2 つの方法があり、おそらく前者のほうが良さそう。

- `layout` オプションに type を無視してディレクトリを作るようなコールバックを渡す
- `exportsOverride` のキー名を空文字にする

### 背景

[Bower](http://bower.io/) は基本的に対象のパッケージを git pull してくるだけで、ファイルをいい感じに配置しれくれたりいらないファイル (そのパッケージのテストやビルドスクリプトなど) を省いたりといった機能は提供していない。そのあたりの面倒はサードパーティのライブラリで見たほうがよさそうで、例えば grunt プラグインの [grunt-bower-task](https://github.com/yatskevich/grunt-bower-task) はいろいろとファイル配置に関する機能を提供している。なかでも [exportsOverride](https://github.com/yatskevich/grunt-bower-task#advanced-usage) というディレクティブで細かくファイル配置を指定できる。bower が [一応公式に推奨している main ディレクティブ](https://github.com/bower/bower.json-spec#main) は単にインストール対象となりうるファイルの一覧 (パターン) を記載するだけだが、`exportsOverride` はより細かい指定ができて実用的だ。

`exportsOverride` には `type` という概念がある。インストール対象のパッケージに含まれるファイルそれぞれに `js`, `css`, `img` といった型がわりあてられ、(`layout` が `byType` の場合は) たとえば `jquery/js/jquery.js` といったように `型名/パッケージ名/ファイル` といった配置になる。

ここで、(1 種類の type のファイルしか提供しないパッケージや既存のプロジェクトを bower 移行する場合などにありそうだが、) わざわざ型別にディレクトリを切らず、パッケージ名ディレクトリの直下にファイルを配置したい。[こちらの qiita の投稿](http://qiita.com/okeyaki/items/f0e009638d353a299f47) によると `exportsOverride` のオブジェクトのキー名 (本来 `type` を入れるべきところ) を空文字列にするといいらしい。

確かにこれで意図した動作になるし、[コードを読んだ限りでも](https://github.com/yatskevich/grunt-bower-task/blob/60dc2181df5945e2e3d94485dbb4751e80cf92da/tasks/lib/asset_copier.js#L43) うまく動きそう。ただしこの方法はドキュメントに記載がないし、設定ファイルをみても挙動が読めないのでトリッキーな印象がある。別の方法として [layout オプションに任意のコールバックが渡せる](https://github.com/yatskevich/grunt-bower-task#optionslayout) ので、こんな感じでそちらで頑張ったほうがまだましな気がする。

<pre><code data-language="javascript">var path = require('path');

grunt.initConfig({
  bower: {
    install: {
      options: {
        layout: function(type, component, source) {
          // type を無視する
          return path.join(component);
        }
      }
    }
  }
});</code></pre>

ちなみに JavaScript のオブジェクトのキーは文字列であればいいので、[空文字列でももちろんいい](http://stackoverflow.com/questions/8343938/should-i-use-an-empty-property-key)。たいていのブラウザでも問題なく動作するようだ。

またこれも余談だけど、`exportsOverride` を `bower.json` 側に書かせるのはなぜなんだろう。あくまで `grunt-bower-task` が提供している機能なので、`Gruntfile.js` に書いたほうがよさそうなんだけれど。
