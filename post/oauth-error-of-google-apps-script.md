{"title":"Google Apps Script の OAuth エラー","date":"2018-07-01T21:30:01+09:00","tags":["etc"]}

Google Sheets を自動更新していた Google Apps Script が以下のエラーで動かなくなっていた。

    The OAuth identity of this script has been deleted or disabled. This may be due to a Terms of Service violation.

    このスクリプトの OAuth ID は削除されたか、無効になっています。利用規約違反が原因である可能性があります。

[OAuth Error \- script deleted or disabled \- Stack Overflow](https://stackoverflow.com/questions/44270918/oauth-error-script-deleted-or-disabled) を参考に、結論としては、スクリプト (プロジェクト) を作り直すと解消した。

- スクリプトをローカルなどに一旦退避させる (コピペなどで)。
- スクリプトエディタからプロジェクトを削除する。
    - `ファイル > プロジェクトを削除`
- スクリプトエディタを開き直し、新しいプロジェクトを作成する。
- スクリプトの内容は先程退避させた内容をコピペする。

この問題の原因は、スクリプトが紐付いている Cloud Platform プロジェクトや呼び出している拡張サービスになにかの変更や問題、例えば Term of Service に更新があり同意が必要など、があったことらしい。なので普通は個別にそれらの対応をしたほうが良い。ただしエラーメッセージからは具体的にどこが原因かはわからず、また今回は試行錯誤しても解消しなかったので、手っ取り早くプロジェクトを作り直した。このへんを個別に確認するには以下の手順となる。

- スクリプトエディタを開く
- `リソース > Cloud Platform プロジェクト`
- 紐付けられている Cloud Platform プロジェクトを開く
- 規約同意などの画面が出ると思われるので、同意する
- 解消しなければ `リソース > Google の拡張サービス` から使用しているサービスの部分を確認しても良いかも

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B07BNB1Z9L/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://images-fe.ssl-images-amazon.com/images/I/61-EfNRESDL._SL160_.jpg" alt="詳解！ Google Apps Script完全入門 ～Google Apps & G Suiteの最新プログラミングガイド～" style="border: none;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B07BNB1Z9L/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">詳解！ Google Apps Script完全入門 ～Google Apps & G Suiteの最新プログラミングガイド～</a><div class="amazlet-powered-date" style="font-size:80%;margin-top:5px;line-height:120%">posted with <a href="http://www.amazlet.com/" title="amazlet" target="_blank">amazlet</a> at 18.07.01</div></div><div class="amazlet-detail">秀和システム (2018-03-23)<br />売り上げランキング: 440<br /></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B07BNB1Z9L/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
