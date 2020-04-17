{"title":"DMP とアドエクスチェンジのオーディエンスセグメント情報のやりとり","date":"2013-04-03T23:48:43+09:00","tags":["ad"]}

以下の quora の質問を読んだメモ。

[How are the audience segments mapped between a DMP and an ad exchange? - Quora](http://www.quora.com/How-are-the-audience-segments-mapped-between-a-DMP-and-an-ad-exchange)

DMP とエクスチェンジの間でオーディエンスセグメントはどのようにしてマッピングされているのか。

### 質問

Bluekai のような DMP 業者はエクスチェンジとどのようにしてセグメントをマッピングしていますか? おそらく pixel piggyback と server to server integration が答えになると予想しています。

### 回答

エクスチェンジでクッキーを有効化するために、BlueKai を含む多くの DMP はエクスチェンジと cookie sync を剃る必要があります。典型的には、まず BlueKai がクライアントに渡しているコンテナタグに pixel を設定します。この pixel が発火すると同時にエクスチェンジにリクエストを飛ばし、その際に DMP のユーザー ID を伝えます。これがいわゆる pixel piggyback と呼ばれる手法です。

この処理によってエクスチェンジは DMP のユーザー 123 が自分のシステムのユーザー abc であるということを知ることができます。そして server to server データフィードで、ユーザーを巻き込むことなく、 DMP はユーザーのデータを渡すことができます。普通の pixel によるインテグレーションではなく server to server 方式を使うのは、単に pixel が不要という理由ではなく、ユーザーのデータと identification を分離できるというのが理由です。つまり、pixel でのインテグレーションはユーザーの識別子だけでなくその属性も同時に長いクエリストリングに乗せて渡すことになります。これだと一度に渡すことができるデータ量の制限、url にデータが含まれるのでプライバシーの問題などが起こります。

以下のリンクも参照してください。

* [Data Management Part III: Syncing Online Data to a Data Management Platform - Ad Ops Insider](http://www.adopsinsider.com/online-ad-measurement-tracking/data-management-platforms/syncing-online-data-to-a-data-management-platform/)
* [SSP to DSP Cookie-Synching Explained - Ad Ops Insider](http://www.adopsinsider.com/ad-exchanges/ssp-to-dsp-cookie-synching-explained/)
