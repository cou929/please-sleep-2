{"title":"ユーザースクリプトを userscript.org から chrome にインストールできなくなっていた","date":"2012-11-04T22:37:05+09:00","tags":["javascript"]}

知らなかったんですが, userscript.org から chrome にユーザースクリプトをインストールできなくなっています.

![userscript のインストールが chrome にブロックされている図](http://gyazo.com/fba4db62d67b55465a6352e0944914bb.png?1352034300)

セキュリティ上の理由より, chrome 21 からは google のストア以外からのエクステンションやユーザースクリプトのインストールはできないようになっているらしいです.

[Chrome can't (directly) install userscripts anymore - Userscripts.org](http://userscripts.org/topics/113176)

インストールしたい場合は,

1. スクリプトをローカルに落として
2. chrome のエクステンションの画面を開いて
3. ローカルのファイルをドラッグアンドドロップ

という手順でないといけなくなってしまいました.

以前の挙動にもどすには,

1. chrome の起動時に `--enable-easy-off-store-extension-install` というフラグを付ける
2. ポリシーリストの ExtensionInstallSources に許可する url を追加する
  - [Policy List - The Chromium Projects](http://www.chromium.org/administrators/policy-list-3#ExtensionInstallSources)

試していないのですが, これらの方法でいいけるそうです. chrome://flags にも設定項目はないそうです.

詳しくないんですがそこまでするほどに悪意を持ったエクステンション / ユーザースクリプトなどが最近増えているのか, はたまたストアのマーケティング的な意図なのか. エクステンションはともかくユーザースクリプトはオフィシャルなマーケットプレイスのような場が無いので, もういまどき流行らないんでしょうかね...

### 参考

- [Adding extensions from other websites - Chrome Web Store Help](http://support.google.com/chrome_webstore/bin/answer.py?hl=en&answer=2664769&p=crx_warning)
