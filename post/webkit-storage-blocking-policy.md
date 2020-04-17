{"title":"webkit ではサードパーティドメインの localStorage が sessionStorage になる","date":"2013-12-10T11:58:18+09:00","tags":["browser"]}

[iOS7 で 3rd party domain での localStorage の挙動 - Please Sleep](http://please-sleep.cou929.nu/ios7-localstorage-behavior-with-privacy.html) のつづき。

iOS7 の Safari でサードパーティドメインの localStorage のデータがブラウザの再起動でクリアされているような挙動を発見した。W3C の [Web Storage](http://www.w3.org/TR/webstorage/) の仕様を読んでみると、今回の挙動と一致するものが "MAY" で規定されていた。というのがここまでのあらすじ。

webkit の bugzilla をみているとそれっぽいバグが見つかった。差分をみるとやはり今回の挙動と一致。webkit のコードレベルで確認がとれたという話。

### webkit のバグチケット

見つけたバグはこちら。

[Bug 115004 – Change approach to third-party blocking for LocalStorage](https://bugs.webkit.org/show_bug.cgi?id=115004)

経緯として、この対応の前にこちらのバグがあった。

[Bug 93390 – Allow blocking of third-party localStorage and sessionStorage](https://bugs.webkit.org/show_bug.cgi?id=93390)

bug 93390 はクッキーの設定に応じて localStorage の挙動も変更するというもの。"サードパーティクッキー拒否" の設定になっている場合に、サードパーティドメインの localStorage を操作しようとすると、セキュリティ例外を出すようになっていた。

そして今回の bug 115004 で、例外を出すのではなく sessionStorage に保存するようポリシーが変更された。bug 93390 の変更で壊れてしまったサイトが多かったためアプローチを変更したらしい。

> The approach to blocking LocalStorage introduce in r125335 has broken many sites that either depend on LocalStorage sufficiently that disabling it has broken functionality, or sites that blindly try to access it without attempting to catch an exception.

### コードを追う

webkit のコードはこちらの github にあるミラーをクローンした。もちろん本家の svn から持ってきてもよい。

[WebKit/webkit](https://github.com/WebKit/webkit)

#### テストケース

まずはテストケースから追う。

テストシナリオとしては

- ある html が iframe にオリジンが違うドキュメントをよみこむ
- iframe 内では localStorage に適当な値をセットする
  - いままではここで例外が送出されていたが、今回は sessionStorage にセットされる
- そのサードパーティオリジンの html に遷移
- localStorage から値を get する
  - ここで別セッションとみなされ、ストレージから値がとれなくなっていることを確認する

となっている。

個別にテストコードを見ていくと、`LayoutTests/http/tests/security/cross-origin-local-storage.html` は

- iframe に `security/resources/cross-origin-iframe-for-local-storage.html` を読み込む
  - これがサードパーティドメインの扱い
- iframe のロードが終わると `security/resources/load-local-storage.html` に遷移する
- ポリシーは "サードパーティをブロックする" に設定

というものだ。

iframe に読み込まれる 'cross-origin-iframe-for-local-storage.html' は localStorage の適当なキーに適当な値をセットする。

その後ロードされる `security/resources/load-local-storage.html` は同様の値で localStorage から get しその値を dom に表示させる。

このテストの期待結果は `LayoutTests/http/tests/security/cross-origin-local-storage-expected.txt` に定義されていて、"No value"、つまり値が get できないということを期待している。

#### コード

`Source/WebCore/page/DOMWindow.cpp` の `DOMWindow::localStorage` メソッドはいわゆる js の `window.localStorage` に対応している。この中の以下のコードで storageArea とよばれる、ストレージの保存先を決定している。

<pre><code data-language="c">
    RefPtr&lt;StorageArea&gt; storageArea;
    if (!document->securityOrigin()->canAccessLocalStorage(document->topOrigin()))
        storageArea = page->group().transientLocalStorage(document->topOrigin())->storageArea(document->securityOrigin());
    else
        storageArea = page->group().localStorage()->storageArea(document->securityOrigin());
</code></pre>

`document->securityOrigin()->canAccessLocalStorage(document->topOrigin())` で現在のドキュメントが topOrigin (iframe の場合親のドキュメントのオリジン) の localStorage にアクセスできるかをチェックしている。

`canAccessLocalStorage` は `canAccessStorage` へのエイリアス。`canAccessStorage` は `Source/WebCore/page/SecurityOrigin.cpp` で定義されている。

<pre><code data-language="c">
bool SecurityOrigin::canAccessStorage(const SecurityOrigin* topOrigin, ShouldAllowFromThirdParty shouldAllowFromThirdParty) const
{
    if (isUnique())
        return false;

    if (m_storageBlockingPolicy == BlockAllStorage)
        return false;

    // FIXME: This check should be replaced with an ASSERT once we can guarantee that topOrigin is not null.
    if (!topOrigin)
        return true;

    if (topOrigin->m_storageBlockingPolicy == BlockAllStorage)
        return false;

    if (shouldAllowFromThirdParty == AlwaysAllowFromThirdParty)
        return true;

    if ((m_storageBlockingPolicy == BlockThirdPartyStorage || topOrigin->m_storageBlockingPolicy == BlockThirdPartyStorage) && topOrigin->isThirdParty(this))
        return false;

    return true;
}
</code></pre>

サードパーティをブロックしている場合は一番したの if 文がポイントになる。ポリシーが `BlockThirdPartyStorage` かつ topOrigin と this の関係がサードパーティと判定された場合、false が返される。

話を storageArea の選択に戻すと、canAccessStorage が偽だった場合 storageArea には `transientLocalStorage` が設定される。 transientLocalStorage の内容について `Source/WebCore/storage/StorageNamespaceImpl.cpp` を見てみると、sessionStorage を返していることがわかる。

<pre><code data-language="c">
PassRefPtr<StorageNamespace> StorageNamespaceImpl::transientLocalStorageNamespace(PageGroup* pageGroup, SecurityOrigin*)
{
    // FIXME: A smarter implementation would create a special namespace type instead of just piggy-backing off
    // SessionStorageNamespace here.
    return StorageNamespaceImpl::sessionStorageNamespace(*pageGroup->pages().begin());
}
</pre></code>

以上は、コードベースを理解しきっているわけではないのでディティールは間違っていると思うが、大筋としてはサードパーティをブロックに設定している場合は sessionStorage にデータが保存されると考えて良さそうだ。

### まとめ

ブラウザがデフォルト設定の場合、[Bug 115004](https://bugs.webkit.org/show_bug.cgi?id=115004) のコミットから、サードパーティオリジンの localStorage は sessionStorage になったと考えて間違いがなさそう。それ以前は例外が送出されていて ([Bug 93390](https://bugs.webkit.org/show_bug.cgi?id=93390))、そのさらに前はブロックを指定なかった模様。

