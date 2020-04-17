{"title":"bower パッケージの登録を解除する","date":"2015-02-28T17:06:43+09:00","tags":["javascript"]}

![](/images/bower-unregister-logo.png)

名前を間違えて登録してしまった場合など、パッケージの登録を解除したい場合。

ドキュメントによると `bower unregister` というコマンドを実装予定らしいが、現在はまだない。よって以下のように直接 `DELETE` リクエストを発行する。

    curl -X DELETE "https://bower.herokuapp.com/packages/<package>?access_token=<token>"

`<package>` は解除したいパッケージ名、`<token>` は GitHub の Personal Access Token に置き換える。GitHub の Personal Access Token は [こちらの設定画面](https://github.com/settings/applications) から `Generate New Token` で発行できる。発行時にそのトークンに与える権限 `Scope` の選択を求められるが、ここはデフォルトのままでOK。

![](/images/bower-unregister-generate-gh-token.png)

解除後は `bower cache clean` でローカルキャッシュをクリアすると良い。

この API での解除がうまくいかない場合は以下の issue に解除依頼を出す必要がある。

[Unregister package requests · Issue #120 · bower/bower](https://github.com/bower/bower/issues/120)

### 参考

[Creating packages · Bower](http://bower.io/docs/creating-packages/#unregister)
