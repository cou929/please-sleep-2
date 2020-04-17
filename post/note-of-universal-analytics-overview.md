{"title":"GA の新機能 Universal Analytics の概要","date":"2012-11-10T10:25:36+09:00","tags":["memo"]}

![logo](http://gyazo.com/48a8860f4271e64b08157c183efdadd4.png?1352260105)

[About Universal Analytics - Analytics Help](http://support.google.com/analytics/bin/answer.py?hl=en&answer=2790010)

Universal Analytics という Google Analytics の新機能でマルチデバイスのトラッキングができるようになったらしいので概要を訳した。

http リクエストで GA にデータを送れるようになったのが最大のポイントで、Web に限らずどこからでもデータを送られるようになった。アクセス解析以外にも使えそうだ。

一番気になっていたクロスデバイスでどうやってユーザーの同一性を特定するのかがこの文章だとわからなかったのでさらに調べよう。

### About Universal Analytics. New features, better insights

> Universal Analytics introduces a set of features that change the way data is collected and organized in your Google Analytics account.

Universal Analytics (UA) はGoogle analytics のアカウントでデータ収集と構成を行う新しい方法を提供します。

> With Universal Analytics, you can collect more types of data and improve your data quality, so you can get a better understanding of how visitors interact with your business at every stage ― advertising, sales, product usage, support, and retention.

UA ではより多くの種類のデータを収集でき、データの質を高められます。それによって事業のステージに応じてビジターとどうインタラクトすべきか (広告、営業、製品説明、サポート、リテンション) のより良い理解につながります。

>  With Universal Analytics (UA), you can:

UA によって以下のことができるようになります。

> - Use the Measurement Protocol to integrate data across multiple devices and platforms.
>  The Measurement Protocol introduced by UA lets you collect and send incoming data from
any device to your Analytics account, so you can track more than just websites. Leverage the new analytics.js code and new developer reference libraries from the Measurement Protocol to see how users interact with all of your devices ― smartphones, tablets, game consoles, and even digital appliances.

Measurement protocol で様々なデバイスやプラットフォームからのデータを統合できます。
UA で導入される Measurement protocol を使うと様々なデバイスからあなたの GA のアカウントへデータを送信・収集できます。よって Web サイト以外のトラッキングも可能になります。新たな analytics.js と Measurement protocol のライブラリを活用し、ユーザーがあらゆるデバイス (スマートフォン、タブレット、ゲーム機、デジタルアプライアンス) とどうインタラクトしているかを見られます。

> - Improve lead generation: Sync offline and online data.
>  With UA, you can track data from all your online and offline customer contact points, like marketing campaigns, sales calls, and store visits, so you can discover relationships between the channels that drive conversions. Because UA is primarily an innovation in data collection methods, there are no new reports showing cross-device data.

リードジェネレーションの改善: オフラインとオンラインのデータ同期
UA によって、顧客のすべてのオフラインとオンラインのコンタクトポイント、例えばマーケティングキャンペーン、営業電話、来店、をトラッキングできます。その中からコンバージョンに貢献する関連を見出せます。UA のイノベーションはデータの収集方法なので、クロスデバイスのトラッキングレポートが別々になることはありません。

> - Define your own dimensions & custom metrics.
>  Custom dimensions and custom metrics are like default dimensions and metrics in your Analytics account, except you create them yourself. Use them to collect data that Google Analytics doesn’t automatically track.

独自のディメンション、メトリクスの定義。
カスタムディメンション・メトリクスは自分でデフォルトのもののようなものを作れます。GA がデフォルトではトラックしないデータの収集が可能です。

> - Understand how well your mobile apps perform.
>  Mobile App Analytics captures mobile app-specific usage data and integrates it with your Google Analytics account, where you can reapply your knowledge of web analytics to dedicated app reports. (Currently in beta.)

モバイルアプリプラットフォームへのより良い理解。
Mobile App Analytics はモバイルアプリに特化した使用状況のデータを GA に統合します。そのため GA で培ったノウハウをそのまま適用できます。

[Mobile App Analytics](http://support.google.com/analytics/bin/answer.py?answer=2709828)

#### Set up UA

> Universal Analytics is only available for a limited number of beta users. Request access to the Universal Analytics beta.

UA は限られたベータユーザーに提供しています。[こちら](https://services.google.com/fb/forms/analyticspreview/)からリクエストを出してください。

> Setting up UA is two step process that requires administrator access to a Google Analytics account and technical knowledge of your development environment. Follow each steps outlined in our setup UA overview to get started. If you don’t have a Google Analytics account yet, you can sign up now.

UA の設定は 2 ステップで、GA アカウントのアドミン権限とあなたの開発環境への技術的知識が必要です。まずはオンラインの setup UA overview にしたがってセットアップしてください。GA のアカウントがない場合はサインアップが必要です。

[Setup UA Overview](http://support.google.com/analytics/bin/answer.py?hl=en&answer=2817075)
[Sign up](http://support.google.com/analytics/bin/answer.py?answer=1008015)

#### Usage guidelines

> UA is developing technology, and not all existing features of Google Analytics are fully supported in UA. Because UA is primarily an innovation of data collection methods, there are no new reports at this time, but any custom dimensions & metrics will appear in existing reports. Learn more about the UA usage guidelines and review the security and privacy information.

UA は開発中の機能で、UA にある機能がまだフルではサポートされていません。UA のイノベーションはデータ収集なので、新しいレポート画面はありませんが、カスタムディメンション・メトリクスが既存のレポートに追加されます。より詳しくは UA usage guidelines を、セキュリティとプライバシーの情報は security and privacy information を参照してください。

[UA usage guidelines](http://support.google.com/analytics/bin/answer.py?answer=2795983)
[security and privacy information](http://support.google.com/analytics/bin/answer.py?answer=2838718)

#### Server side configuration options

> In addition to the new collection methodsj introduced in the Measurement Protocol, UA exposes options that let you configure server side features. See what you can control from your Analytics account interface

UA の Measurement Protocol で追加された機能で、オプションをつけることでサーバサイドの機能を設定できます。

> - Customize organic search sources

> - Session and campaign timeout handling

> - Referral exclusions

> - Search term exclusions

> - About Mobile App Analytics

> - About custom dimensions & metrics

- オーガニックサーチソースのカスタマイズ
- セッションとキャンペーンのタイムアウトの扱い
- リファラの例外設定
- 検索ワードの例外設定
- モバイルアプリ解析の設定
- カスタムディメンション・メトリクスの設定

> Visit our developer site for detailed information on implementation and configuration.

開発者向けドキュメントにより詳細な情報があります。

[developer site](https://developers.google.com/analytics/devguides/collection/protocol/v1/)
