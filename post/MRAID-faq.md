{"title":"MRAID FAQ","date":"2013-04-09T00:10:45+09:00","tags":["ad"]}

[Mobile Rich Media Ad Interface Definitions (MRAID)](http://www.iab.net/mraid) を読んだメモ. IAB にある MRAID のサマリ文章。

### Executive Summary

近年モバイルアプリ・モバイル Web のリッチメディアディスプレイ広告は重要度を増しており、多くの会社がモバイル広告のエコシステムを構築すべくチャレンジしている。モバイルリッチメディア広告のイノベーションによってパブリッシャー、広告主のエキサイティングな可能性が導かれるが、一方でコンテンツマネタイゼーションの遅れや阻害の要因にもなっている。

クリエイティブのデザインプロセスをシンプルにすることは、代理店、広告主に大きなメリットだ。

### What is MRAID?

MRAID (Mobile Rich Media Ad Interface Definitions) は IAB によるプロジェクトでモバイルアプリ上で動くリッチメディアの共通 API を定義するものだ。仕様では HTML5 と JavaScript を使ったコマンドを定義している。開発者はコマンドを通じて広告を操作する。操作には広告のエキスパンド、リサイズ、デバイスファンクションへのアクセスなどがある。

現在各ベンダーの SDK はそれぞれ API が異なっていて、開発者は同じクリエイティブでもそれぞれのベンダーに応じて書きなおす必要がある。こうした状況に対して MRAID のゴールは統一した API を提供することだ。

つまり MRAID 対応のリッチメディア広告は MRAID 対応の SDK で開発されたアプリで動作する。また代理店は異なるパブリッシャーのアプリ上でもひとつのクリエイティブでよくなる。

### Back up a minute: What is an API, and how does it relate to an SDK?

API とはプログラマーがあるソフトウェアやハードウェアを操作するためのコマンドのセットのことだ。MRAID API はリッチメディア広告のクリエイターがその広告が動いているアプリとコミュニケートできるようにする JavaScript ベースのインターフェースだ。MRAID API は HTML/JS 経験者には簡単に習得でき、リッチメディア広告を作る開発者の習慣に沿ったものになるよう意図されている。

SDK とは開発者が一から実装しなくていいようにコードをパッケージしたものである。モバイル広告ではアプリ開発者は専用の SDK を使って広告のハンドリングを行う。広告はアプリに埋め込まれた SDK と API でやり取りする。

MRAID の目的はクリエイターの人生を簡単にすることだ。もしアプリが MRAID に対応していれば、どんな SDK がアプリに組み込まれていても MRAID API の上で書いたとおりに広告が動作する。IAB と MRAID ワーキンググループは SDK そのものを開発しない。あくまでインターフェースの仕様を定義する。

### So do I need to speak code?

それはあなた次第だ。MRAID の仕様は直接的にはクリエイティブの開発者、間接的にはリッチメディアツールセット開発者に向けられている。また MRAID はまだ策定中だが、いくつかのベンダーはすでにドラフトに対応させ始めている。

### Why is the IAB running the MRAID project?

アプリ内広告は拡大している分野でパブリッシャー・広告主双方の関心の的だ。初期のプレイヤーはそれぞれ独自の API・SDK を提供しているためバイヤーとセラーの間の摩擦となっていた。昨年 (2011年) 数社があつまりモバイルリッチメディアの API を作成し始めた。このオープンソースプロジェクトは Open Rich Media for Mobile Advertising (ORMMA) と呼ばれ、API・SDKの良いスタートとなった。

しかしながらより多くの業界からの参加者の合意を得る上で、IAB での策定作業は重要だ。

MRAID のイニシアチブは ORMMA や他のベンダーの独自 SDK とは独立だが、これらの成果物の要素を含んでいる。

また IAB の MRAID には次のメリットもある。

- 他の仕様の API と一貫性が持たれる。たとえば VPAID などの IAB の仕様。
- 代理店コミュニティとの結びつきが強い。

### How are MRAID and ORMMA working together going forward?

ORMMA の仕様は MRAID に行こうし、ORMMA は標準 API のサポートにフォーカスする。また ORMMA の参加者は MRAID の仕様の主要なメンバーでもある。とはいえ ORMMA は MRAID とは独立であり、MRAID はすべての成果物が MRAID 対応することを願っている。

IAB は MRAID の標準策定を行い、ORMMA はそれに追従しオープンなツール開発などのサポートを行う。

### Does MRAID limit innovation?

いいえ。ベンダーやパブリッシャーは独自の SDK を作ることができる。MRAID の API は SDK が実装すべき最小のセットを定義している。ベンダーが MRAID の基礎のうえに拡張を作ることも可能だ。

### Who are the companies participating in drafting the MRAID spec?


- 24/7 Real Media, Inc.
- AccuWeather.com
- AdMarvel
- AdMeld
- ADTECH
- Adobe Systems Inc.
- AOL
- CBS Interactive
- Celtra
- Crisp Media
- Dow Jones & Company
- ESPN
- FreeWheel
- Goldspot Media
- Google
- Greystripe
- IDG
- InMobi
- Innovid
- Jumptap
- Medialets
- MediaMind
- Microsoft Advertising
- Mixpo
- MobClix, A Velti Company
- Mocean Mobile
- NBC Universal Digital Media
- New York Times Co.
- Nexage
- Pandora
- PointRoll
- Rhythm NewMedia
- Spongecell
- Sprout
- TargetSpot
- Time Inc.
- Turner Broadcasting System, Inc.
- Univision
- The Weather Channel
- Yahoo!, Inc.
