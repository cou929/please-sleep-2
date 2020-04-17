{"title":"ソーシャルプラットフォームでの signed request のフロー","date":"2013-06-09T22:39:30+09:00","tags":["memo"]}

いわゆるソーシャルゲームなど、あるプラットフォームの iframe 内で動くサードパーティ web アプリ、OAuth を使ってプラットフォームの API をたたいてエンドユーザーにサービスを提供するようなもの。この手のアプリでは、まずプラットフォームにアプリケーション登録し、アプリの ID や OAuth の consumer_key, consumer_secret などの必要な情報を払いだしてもらう。ユーザーがプラットフォーム内から特定のアプリにアクセスしようとすると、プラットフォーム側は iframe にサードパーティのアプリを読み込んでユーザーに提供する。iframe からリクエストをおくるサードパーティ側のエンドポイントは事前のアプリケーション登録時に設定しておく。

![](/images/signed_request_flow_P0.png)

サードパーティアプリ側では当然、プラットフォームのユーザーごとに情報を保存して別の挙動をさせたり、あるいはバッチ系の操作が必要なら OAuth のアクセストークンもデータストアに持っている可能性もある。プラットフォーム側からリクエストしているのはどのユーザーか、これはユーザーのログインセッションを持つプラットフォーム側だけが知り得る情報だ。誰がリクエストしているかをプラットフォーム側がアプリケーション側に伝えないとまともなアプリは作れない。

ではプラットフォーム側の iframe からは単にユーザー ID をクエリストリングにつけてサードパーティアプリを呼び出すことにしよう。当然ながらこの方式には脆弱性がある。悪意のある第三者が適当にクエリストリングを偽装してアプリ側にリクエストを送ると、簡単にそのユーザーになりすましてアプリを利用することができてしまう。

![](/images/signed_request_flow_P1.png)

対策もストレートに、プラットフォーム側とアプリ側の 2 者しか知り得ない情報でリクエストをチェックすればよい。この場合最も簡単なのはプラットフォームが払いだしている OAuth の consumer\_secret だ。一般的なのは、プラットフォーム側はアプリケーション側にユーザー ID などの必要な情報とともに、それを SHA1 でハッシュ化したシグネチャを送る。ハッシュ化の際には対象のアプリに払いだした consumer\_secret をキーに用いる。リクエストをうけたアプリ側では、送られた内容からシグネチャとそうでない部分 (データ) を分離。データを自分の consumer\_secret でハッシュ化し、シグネチャと一致するか調べる。こうしてリクエスト元の正当性を調査するという仕組みだ。

![](/images/signed_request_flow_P2.png)

このへんの話は、仕様としては opensocial で定められている。

- [Introduction To Signed Requests - OpenSocial Reference Material - OpenSocial Wiki](https://opensocial.atlassian.net/wiki/display/OSREF/Introduction+To+Signed+Requests)

各ソーシャルプラットフォームのベンダーのドキュメントを読むと具体的でわかりやすい。Facebook は、全般的にそうだが今回も例に漏れず、オープンな仕様に若干独自のアレンジを加えた仕様になっている。

- [Using Login with Games on Facebook - Facebook Developers](https://developers.facebook.com/docs/facebook-login/using-login-with-games/)
- [起動時のパラメータとOAuth Signatureの検証 << mixi Developer Center (ミクシィ デベロッパーセンター)](http://developer.mixi.co.jp/appli/ns/pc/oauth_signature/)
- [Gadget サーバーからゲームサーバーへのリクエスト - Smartphone Web - Mobage Developers Documentation Center](https://docs.mobage.com/display/JPSPBP/Gadget+to+GameServer)
- [OAuth認証方法について - GREE Developer Center](https://docs.developer.gree.net/ja/globaltechnicalspecs/webapplication/oauth)
