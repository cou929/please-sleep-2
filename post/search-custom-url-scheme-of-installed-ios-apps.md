{"title":"インストール済み iOS アプリのカスタム URL スキームを調べる","date":"2014-03-26T23:59:39+09:00","tags":["mobile"]}

面倒くさいけど一番確実なのは iPhone 構成ユーティリティで確認する方法だと思う。

[アップル - サポート - ダウンロード](http://support.apple.com/ja_JP/downloads/#iphone 構成ユーティリティ)

- こちらからダウンロードし起動
- iPhone を接続
- 左カラムの LIBRARY -> Devices から接続したデバイスを選択
- 上部にある Export ボタンでデバイス情報を出力

これでデバイス情報の XML ファイルが出力される。この中からなんとかみたいアプリの情報を探さなければならない。

`CFBundleDisplayName` という要素にデバイスで表示されるアプリ名が入る。

    <key>CFBundleDisplayName</key>
    <string>FooApp</string>

`CFBundleURLSchemes` にそのアプリの URL スキームが入る

    <key>CFBundleURLSchemes</key>
    <array>
      <string>foo_app_scheme</string>
    </array>

`CFBundleDisplayName` でひっかけてから `CFBundleURLSchemes` を見るのがよさそう。

こちらでいい感じにアプリ名と url scheme を抜き出してくれるブックマークレットが紹介されていたので、これを使うのが簡単。

[iOS6のURLスキームをお手軽に調べる決定版 SchemeTaker - W&R : Jazzと読書の日々](http://d.hatena.ne.jp/wineroses/20121023/p1)

このへんのランチャー系のアプリを使うと端末を PC に接続しなくても調べられるらしいのだが、アプリが有料だったりそもそもストアになかったりして、自分のニーズには合わなかった。

- [URLスキームを知る方法。iPhoneアプリ「Launch Center Pro」で簡単にURLスキームを見つけることができます。 | Cohtaro.com](http://cohtaro.com/2014/02/2336)
- [iPhoneだけでサクッとURLスキームを調べる方法があった。 | オレっち.com](http://ore-ch.com/urlscheme-search.html)

