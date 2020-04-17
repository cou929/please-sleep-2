{"title":"npm で package.json がカレントディレクトリにない場合","date":"2015-02-28T17:49:30+09:00","tags":["javascript"]}

例えば `package.json` がリポジトリのルートから見て `foo/` においてあり、手元での開発時は `cd foo/` してから `npm install` (パッケージのインストール先は `foo/node_modules/`) なり、 `npm test` しているような場合。

このようなプロジェクトを CI したいという時には、以下のように `-prefix` や `-cwd` オプションを指定するとうまく動作する。`travis.yml` に書く例だと、

    language: node_js
    node_js:
        - "0.12"
    install:
        # foo/package.json に定義されているパッケージを foo/node_modules にインストールする
        - npm —prefix ./foo install ./foo
    script:
        # foo/package.json で指定されているテスト実行コマンドを実行する。その際のカレントディレクトリは foo/ になる
        - npm test —cwd ./foo

もちろん、`cd foo/; npm install` のようにしてもいいけれど、当然だがこうすると実際にカレントディレクトリが変わってしまう。例えばリポジトリのルートに別のビルドの仕組みの起点があるような場合 (サーバサイドを別言語で書いていて、その中のフロントエンドの js 等だけを npm で管理しているようなプロジェクトではあり得る構成だと思う) には、カレントディレクトリがどこにあるのかいちいち注意する必要ができてしまう。それよりはカレントディレクトリを動かさずに、上記のように `-prefix` や `-cwd` オプションで対応したほうがいいと思う。
