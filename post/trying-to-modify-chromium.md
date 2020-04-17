{"title":"Chromium のコードを触ってみる","date":"2015-03-03T00:14:58+09:00","tags":["browser"]}

![](/images/chromium-logo.png)

[Contributing to Chromium: an illustrated guide – Monica Dinculescu](http://meowni.ca/posts/chromium-101/) という記事が面白かったのでやってみた。Chromium へのコントリビューションのガイドということで、環境構築、コードへの手の入れ方から、ビルド、差分の送り方まで解説している記事。

題材としてはこのように、アバターのメニューに項目を追加するというもの。

![](/images/trying-to-modify-chromium-result.png)

さすがにパッチを送るわけにはいかないので、コードに差分をいれてビルドするところまでをやってみる。

### 環境構築

記事にもあるが、基本的に [Get the Code: Checkout, Build, Run & Submit - The Chromium Projects](http://www.chromium.org/developers/how-tos/get-the-code) の内容が最新の状態に保たれているので、これに従う。

記事中に説明があるが、この wiki に無い情報としては、以下のビルドオプションの設定がおすすめらしい。

    export GYP_DEFINES="component=shared_library dcheck_always_on=1"

基本的に書いてあるとおりにするだけでよかったが、一点自分の環境では、以下のように `CoreFoundation` がないよというエラーが出た。

    % ninja -C out/Release chrome && out/Release/Chromium.app/Contents/MacOS/Chromium
    ninja: Entering directory `out/Release'
    [396/17862] MACTOOL copy-bundle-resource ../../breakpad/src/client/mac/sender/English.lproj/Localizable.strings
    FAILED:  ./gyp-mac-tool copy-bundle-resource ../../breakpad/src/client/mac/sender/English.lproj/Localizable.strings crash_report_sender.app/Contents/Resources/English.lproj/Localizable.strings
    Traceback (most recent call last):
      File "./gyp-mac-tool", line 601, in <module>
        sys.exit(main(sys.argv[1:]))
      File "./gyp-mac-tool", line 28, in main
        exit_code = executor.Dispatch(args)
      File "./gyp-mac-tool", line 43, in Dispatch
        return getattr(self, method)(*args[1:])
      File "./gyp-mac-tool", line 66, in ExecCopyBundleResource
        self._CopyStringsFile(source, dest)
      File "./gyp-mac-tool", line 105, in _CopyStringsFile
        import CoreFoundation
    ImportError: No module named CoreFoundation
    [396/17862] ACTION Generating resources from app/generated_resources.grd
    ninja: build stopped: subcommand failed.

[PyObjC](https://pythonhosted.org/pyobjc/index.html) が必要という内容だったので pip でインストールした。

    pip install -U pyobjc

また、Chromium は非常に大きいコードベースなので、空間と時間の余裕には注意が必要。特に容量に関しては、リポジトリだけで 7 GB (`--no-history` オプション付きの場合。そうでなければ倍くらいになる) ほどにもなるし、ビルド時にもさらに容量が必要になる。一度手元の MBA の SSD が一杯になり、環境をかえてやり直すはめになってしまった。

### コードに手を入れる

記事の内容に沿ってやってみる。

まずは、アバターのメニューを実装しているファイルを探す。`git grep` なり `ack` なりツールは何でもいいが、ここでは [google code の code search](https://code.google.com/p/chromium/codesearch) を使う。

こういう時の常套手段は、ユニークっぽい文字列で検索して対象のファイルをみつけることだ。今回は `switch person` という文字列で検索してみると、以下のように `generated_resources.grd` というファイルが見つかる。これは文字列のリソースを管理しているファイルらしい。

![](/images/trying-to-modify-chromium-search-switch-person.png)

ここから `IDS_PROFILES_SWITCH_USERS_BUTTON` という ID が見つかる。コードからはこの ID を参照していると思われるので、次はこれで検索してみると、以下のファイルが見つかる。

![](/images/trying-to-modify-chromium-search-id.png)

`src/chrome/browser/ui/views/profiles/profile_chooser_view.cc` と
`src/chrome/browser/ui/cocoa/profiles/profile_chooser_controller.mm` が見つかるが、今回は Mac で試しているので、後者を見ていく。

`createOptionsViewWithRect:` という関数の中で、既存の `Switch Person` のボタンを作って `container` に `addSubView` している箇所がみつかる。ここを真似するとなんとなくできそうとわかる。

以下のように `switchUsersButton` を作ってコンテナに追加している部分をコピペして、適当に書き換えてみる。

<script src="https://gist.github.com/cou929/2a98672804fddca0dcde.js"></script>

ボタンをクリックすると `onMyButtonClicked` が呼び出され、`ShowSingletonTab` で `http://please-sleep.cou929.nu/` を新しいタブで開くようにした。新しいタブで url を開くメソッド `ShowSingletonTab` は、本当はちゃんとコードを検索して調べるべきだったけれど、今回は横着して記事を参考にした。

これでビルドし直すと、意図したとおりにボタンを追加できた。

    (py27)[22:50 kosei@localmba src (b684380...)]% ninja -C out/Release chrome && out/Release/Chromium.app/Contents/MacOS/Chromium --enable-new-avatar-menu
    ninja: Entering directory `out/Release'
    [4/12] SOLINK "Chromium Framework.framework/Versions/A/Chromium Framework", POSTBUILDS
    [12/12] STAMP Chromium.app
    [64001:1299:0302/232010:ERROR:profile_chooser_controller.mm(1873)] Hello, Chromium!

![](/images/trying-to-modify-chromium-result.png)

記事ではこのあとレビューの出し方などの説明が続くが、ここでは割愛。

### 参考

- [Contributing to Chromium: an illustrated guide – Monica Dinculescu](http://meowni.ca/posts/chromium-101/)
- [Get the Code: Checkout, Build, Run & Submit - The Chromium Projects](http://www.chromium.org/developers/how-tos/get-the-code)
- [Chromiumのビルドでつまった話 on MAC OSX 10.8 - Qiita](http://qiita.com/iseki-masaya@github/items/9bad09279749acdb1cf1)
