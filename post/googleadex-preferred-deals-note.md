{"title":"Preferred Deals (Google Ad Exchange)","date":"2012-11-06T22:31:51+09:00","tags":["ad"]}

[Preferred Deal FAQ (beta) - DoubleClick Ad Exchange Help](http://support.google.com/adxbuyer/bin/answer.py?hl=en&answer=1619735)

Google Ad Exchange に Preferred Deals という新機能ができたらしい. ちょうど純広告と RTB の中間のような商品. 理解のために faq を訳した.

バイヤー側はどういうロジックで応札すべきなんだろう. Preferred Deals じゃないふつうのオークションで入札する金額より Preferred Deal が高くなければ応札くらいかな?

### Preferred Deal FAQ (beta)

> Preferred Deals is currently a beta offering, so you won't see some of the options mentioned below if you're not a participant. If you're interested in learning more, sellers can contact us here, and buyers can contact us here.

Preferred Deals は現在ベータ版での提供です. そのため参加者でない方は, 以下で言及しているオプションを見ることができません. 興味がある方は, セラーの方は[こちら](http://www.google.com/support/adxseller/bin/request.py?contactus=1), バイヤーの方は[こちら](http://www.google.com/support/adxbuyer/bin/request.py?contactus=1)にご連絡ください.

#### Fundamentals of Preferred Deals

> This feature provides the ability for publishers to offer inventory to one or more buyers on a fixed price basis. Deals are pre-auction, meaning that fixed price deals will transact if the buyer accepts the inventory, even if the price is below that of the auction. Publishers can offer a Preferred Deal to more than one buyer.

この機能で, パブリッシャーは在庫を特定の 1 つ以上のバイヤーに固定価格でオファーすることができます. 取引はプレ・オークションです. つまりバイヤーが固定価格での在庫の買い付けを承認した場合, その価格がオークションでの価格を下回っていたとしても取引成立となるということです. パブリッシャーはひとつ以上のバイヤーに対して Preferred Deal をオファーできます.

#### Preferred Deals serving logic

> If two buyers are offered a Preferred Deal at $3 on the same ad unit, and if both buyers bid on the impression with the same targeting, who wins?

2 つのバイヤーがある ad unit に $3 で Preferred Deal でオファーされており, 両者が同時に同じターゲティングのインプレッションに対して入札していた場合, 勝つのはどちらか?

> The winner is decided randomly if both buyers bid at least the fixed deal price.

2 つのバイヤーが Preferred Deals の固定価格以上で入札していた場合, 勝者はランダムに選ばれます.

> If two buyers are offered a Preferred Deal at $5 on the same ad unit, and Buyer A bids $6 and Buyer B bids $5, what will the price be?

2 つのバイヤーが $5 で Preferred Deal をオファーされており, バイヤー A が $6, バイヤー B が $5 で入札していた場合, 最終的に価格はどうなるか?

> The two bids are treated equally because both buyers bid over or at the minimum. The winner is then randomly decided between the two and the higher price does not have an impact. The winner stills pay $5.

両者は固定価格以上の額で入札しているので対等になります. よって勝者はランダムに選ばれ, 高い入札価格に設定していても意味がありません. 勝者は $5 を支払います.

> How does a buyer know that inventory is being offered to them as a fixed price deal?

バイヤーはその在庫が自分に固定額の取引をオファーされていると, どのようにして知ることができますか?

> The seller would have to let the buyer know directly.

セラーがバイヤーにダイレクトに知らせます.

> If Buyer A is offered a Preferred Deal at $2 and Buyer B is offered a Preferred Deal on the same ad unit at $3, how are bids handled?

もし同じ ad unit に対してバイヤー A が $2, バイヤー B が $3 でオファーされていた場合, 入札はどのように処理されますか?

> The higher Preferred Deal takes priority.
> * If Buyer B bids $3 or more, it automatically wins the bid. If Buyer A bids more than Buyer B, Buyer B still wins.
> * If Buyer A bids $2 or more and Buyer B bids less than $3, then Buyer A wins because Buyer B didn’t meet its Preferred Deal negotiated price.

高い価格のほうが優先されます.

- もしバイヤー B が $3 以上で入札していた場合, 自動的にバイヤー B が勝ちます. バイヤー A がそれ以上の価格で入札していたとしても B の勝利です.
- もしバイヤー A が $2 以上で入札し, バイヤー B が $3 未満だった場合, バイヤー B は Preferred Deal の交渉価格にマッチしていないのでバイヤー A が勝ちます.

> If Buyer B ($3 Preferred Deal) chooses not to bid, does Buyer A ($2 Preferred Deal) have a shot at winning the impression?

もしバイヤー B ($3 の Preferred Deal) が入札しなかった場合, バイヤー A ($2 の Preferred Deal) はそのインプレッションに勝ちますか?

> Yes, Buyer A can win the impression it bids $2 or more.

はい, バイヤー A と $2 で取引成立します.

> Can the predetermined pre-auction CPM for a Preferred Deal be overbid by the buyer? Example: Preferred Deal between a publisher and a buyer at $5. The buyer bids $10. Does the buyer pay $5 or $10?

事前に設定していた CPM が Preferred Deal 以上の価格だった場合, バイヤーは多い価格で入札できますか? 例えば, Preferred Deal が $5 でバイヤーが $10 で入札した場合, $5 または $10 どちらになりますか.

> To transact a Preferred Deal, the bid must be higher or equal to the agreed fixed CPM. In this example, despite over bidding, the buyer wins the impression and pays the fixed CPM of $5. Bidding above the fixed price simply signals to the auction that the buyer accepts the fixed price; the value of the bid is not used in any other way. If you believe a case can be made for a different behavior, please contact us.

Preferred Deal の処理が行われるのは, 入札額が合意された固定 CPM 以上の場合です. この例ではそれより高く入札していますが, バイヤーは入札勝利後に固定の CPM $5 を支払います. 固定額以上の入札は単純に固定額での取引に応じるという通知です. 入札金額はそれ以外の用途には使われません. もしそうでない挙動があった場合は[ご連絡ください](http://www.google.com/support/adxbuyer/bin/request.py?contactus=1).

#### Information for publishers

> How do I set up a Preferred Deal?

Preferred Deal の設定方法を教えて下さい.

> Learn how to set up a Preferred Deal.

[こちらを参照してください](http://www.google.com/support/adxseller/bin/answer.py?answer=1314059)

> How do Preferred Deals interact with per buyer min CPM deals?

Preferred Deals はどのように各バイヤーの最小 CPM の取引を扱いますか?

> Publishers can offer inventory either as a Preferred Deal or as part of the auction, but not as both. If a Preferred Deal impression fails to get the agreed-upon bids from any of its direct buyers, then the impression goes to the open auction.

パブリッシャーは在庫を各バイヤーに対して, Preferred Deals としてかオークションの一部としてかどちらかでオファーします. もし Preferred Deal でのインプレッションの取引が成立しなかった (提示した額以上の入札がなかった) 場合は, そのインプレッションはオープンオークションに移行します.

> Why should I use Ad Exchange rather than trafficking my ad unit directly in DART?

自分の ad unit を直接  DART に配信するのではなく, アドエクスチェンジを使わなければいけないのですか?.

> Ad Exchange gives publishers the opportunity to set up Preferred Deals with buyers, giving the latter a “first look” at the impression. If the impression remains unsold via the Preferred Deal, it can then be made available in an open auction, helping to maximize yield.

パブリッシャーはアドエクスチェンジを使うことでバイヤーとの Preferred Deals での取引が設定できます (後述のインプレッション時の "first look"). もしインプレッションが Preferred Deal で売れなかった場合はオープンオークションに移行でき, 収益を最大化できます.

> What happens if I offer inventory at the same fixed price to more than one buyer?

2 つ以上のバイヤーに同額でオファーした場合はどうなりますか.

> If the same inventory (ad unit) is offered to multiple buyers at the same price and all buyers bid on that inventory at the fixed price, Ad Exchange will randomly rotate between the different buyer’s bids.

ある在庫 (ad unit) を複数のバイヤーに同額でオファーし成立した場合, ランダムにバイヤーが選ばれます.

> How does a buyer know that inventory is being offered to them as a preferred deal?

バイヤーはどのようにしてその在庫が Preferred Deal でオファーされていると知りますか?

> The seller has to contact the buyer directly.

セラーが直接バイヤーにコンタクトします.

> If a publisher is already seeing impressions, what are the implications for implementing Preferred Deals?

パブリッシャーがインプレッションをすでに見ている場合, Preferred Deal を実施するとどのような影響がありますか.

> If a publisher implements Preferred Deals, buyers selected for the "first look" can’t transact on open auction for that specific ad unit.

パブリッシャーが Preferred Deal を実施している場合, "first look" を選択しているバイヤーはその ad unit に対してオープンオークションに参加できません.

> How does a preferred deal interact with Dynamic Allocation?

Preferred Deal は Dynamic Allocation に対してどのように動作しますか.

> Dynamic Allocation currently applies to all deal types in Ad Exchange. If a DART-booked ad unit has a higher rate than what is available in Ad Exchange (through any deal type), the DART-booked ad unit serves. If the DART-booked ad unit has an equal or lower booked rate, Ad Exchange serves the impression, assuming it meets all targeting requirements.

[Dynamic Allocation](http://www.google.com/support/adxseller/bin/answer.py?answer=178331) はすべてのアドエクスチェンジの取引に適用できます. DART-booked ad unit がアドエクスチェンジのものより高値だった場合, (取引の種類によらず) DART-booked ad unit が配信されます. もし同額かそれ以下だった場合, アドエクスチェンジがインプレッションを配信します (そのインプレッションがすべてのターゲティングの設定に合致していた場合).

> What’s the best practice when setting up multiple buyers for Preferred Deals on the same inventory?

ある在庫を複数のバイヤーに対して Preferred Deal に設定する場合のベストプラクティスは?

> In general, you should offer the ad unit to multiple buyers at the same fixed price. If you offer Buyer A the ad unit at $2 and Buyer B the ad unit at $3, Buyer B will always outbid Buyer A (assuming the targeting is the same). Buyer A’s bids are unlikely to transact, and it will therefore not have an opportunity to access your inventory. But when you offer Buyer A and B the inventory at $2, they’ll split the inventory, all things being equal.

一般的には複数のバイヤーに同額でオファーすべきです. A に $2, B に $3 でオファーする場合, ターゲティングが同じならば, B がかならず A に勝ちます. A の入札は処理されず, あなたの在庫にアクセスすることはありません. 両者に同額の $2 でオファーした場合は在庫を平等にそれぞれ分割します.

> Can I override an ad unit restriction via a Preferred Deal? For example, I’ve blocked Advertiser Alpha from my ad unit, but can I permit Advertiser Alpha to participate in a Preferred Deal?

ある ad unit へかけている制限を Preferred Deal で上書きできますか? 例えば代理店 A の ある ad unit への広告表示をブロックしていますが, Preferred Deal は許可するということはできますか?

> No. Currently, any restrictions you set up for an ad unit apply across all transactions.

いいえ. すべての制限はすべての取引に適用されます.

> Can I set a Preferred Deal for AdWords?

AdWords で Preferred Deals を使えますか?

> No. Support for a preferred deal with AdWords will be available in a future release. However, you can setup Preferred Deals with Google Reserve. Contact your sales rep for more details.

いいえ. AdWords は将来的にサポートされます. しかしながら Google Reserve には Preferred Deals を設定できます. 詳しくは slaes rep にお問い合わせください.

> What is the difference between Preferred Deals and Private Ad Units?

Preferred Deal と Private ad units との違いを教えて下さい.

> Can I set up a Preferred Deal with a buyer set at $3, and also set the min CPM at $6 for the auction? Yes this is fine. However, if the buyer doesn’t win the first look, they can’t bid in the auction.
> Can I offer inventory through a Preferred Deal and then offer it anonymously via the open auction? Yes. Although the Preferred Deal can only transact on the branded version of your ad unit, once it moves to the open auction it's available on both a branded and anonymous basis, assuming you've offered it as both. Remember that buyers for the Preferred Deal can't transact on the impression in the open auction.
> * Private ad units is a feature that lets sellers set a CPM minimum that specific buyers must bid to enter the auction (Learn more). This feature allows publishers to negotiate a distinct floor price directly with a buyer, but that buyer must still compete with all other buyers in the open auction in order to win the impression.
> * Preferred Deals allows publishers to offer select buyers access to the inventory before it goes out to the open auction, and the ability to buy it at a fixed, negotiated, sell price.

Preferred Dael をあるバイヤーに $3 で設定し, オークションの最低 CPM を $6 に設定することはできますか? -> はい可能です. ただし first look でバイヤーが勝たなかった場合, そのバイヤーはオークションには参加しません.
在庫を Preferred Deal でオファーしたあとに, また匿名でオープンオークションをオファーすることはできますか? -> はい. Preferred Deal をブランドバージョンの ad unit としてオファーしたとしても, オープンオークションはブランド・匿名どちらでも可能です. ただし Preferred Deal でオファーしたバイヤーはオープンオークションに参加できません.

- Private ad unit はオークションの入札に最低 CPM を設定する機能です ([詳細](http://www.google.com/support/adxseller/bin/answer.py?answer=1047199)). この機能では最低入札価格をパブリッシャーがバイヤーに交渉します. その最低入札価格を超えていたとしても, オークションで競売は行われるのが違いです.
- Preferred Deal はパブリッシャーが選択したバイヤーに対して, オープンオークションが始まる前に在庫へのアクセスをオファーします. そしてバイヤーはそれを交渉した売値である固定額で買い付けできます.

#### Information for buyers

> How do I set up a Preferred Deal?

どのように Preferred Deals を設定すれば良いですか？

> Buyers can learn how to set up a Preferred Deal.

こちらを参照してください。

> What changes do I need to make to my bidder?

入札者に対してどのような変更がありますか？

> An additional field has been added to the protocol. The changes are easy to make and allow easy identification of Preferred Deals. The field indicates the fixed deal price. Later in Q3 2012 we'll implement an additional field, a "deal ID," which maps to all the specifics of a Preferred Deal or private auction.
> Please refer to the Buyside RTB Preferred Deal Guidelines for specific changes you need to implement.

プロトコルに新しいフィールどが追加されます。Preferred Deals であると容易に特定できます。そのフィールドには固定の取引額が入っています。 2012 Q3 にはさらに deal ID という項目を追加する予定です。これは Preferred Deals と Private Auction を紐づけるものです。
詳しくは Buyside RTB Preferred Deal Guidelines をご覧ください。

> How long does it take for Preferred Deals to start working?

Preferred Deals が開始されるまでどのくらいかかりますか？

> Preferred deals are active as soon as the publisher updates an ad unit to send the Preferred Deal traffic to the buyer.

パブリッシャーが設定するとすぐにバイヤーに送信されるようになります。

> What kind of reporting is available for Preferred Deals?

どのようなレポートがありますか?

> Currently, reporting specifically for buyers is not available, but we plan to provide this feature in the future.

現在バイヤー向けのレポートはありません。将来的には提供予定です。

> If I set up a new ad group to target inventory I have negotiated a Preferred Deal for, will all traffic I receive for this ad group be Preferred Deal enabled?

もしある広告グループに Preferred Deals で交渉した在庫を指定した場合、この広告グループが受けるトラフィックはすべて Preferred Deals が有効なものになりますか？

> No, Preferred Deals are set up by the publisher on slices of their inventory (ad units). When any of your ad groups target this inventory, they'll be participating in Preferred Deals, and when they don’t, they will not. Often, the publisher will not opt in all inventory coming from a given domain/site to Preferred Deals, so targeting a domain will see both Preferred Deal and non-Preferred Deal requests.

いいえ、Preferred Deals はパブリッシャーによって在庫の一部 (ad unit) に設定されます。もしある広告グループがその在庫にターゲティングしていても、Preferred Deals でない入札になる場合があります。パブリッシャーが特定のドメイン/サイトのすべての在庫に Preferred Deals を設定することは少ないので、そのドメインにターゲティングしていても両方のリクエストがくることが予想されます。

> Is there a list of publishers that I can work with, or a sample of the total opportunity?

取引できるパブリッシャーの一覧やトータルの機会のサンプルを見ることはできますか？

> Yes, the list is available on the Ad Exchange beta site. Please contact us for the current list.

はい、アドエクスチェンジのベータサイトで可能です。現在のリストは私たちにご連絡ください。

> What kind of support will I receive?

どのようなサポートを受けることができますか？

> You'll continue to receive the same support levels from the Act Mgr and TAMs. Feel free to contact us with any questions.

Act Mgr や TAMs と同様のサポートが受けられます。お気軽にご連絡ください。

> Can you help me set up meetings with publishers?

パブリッシャーとのミーティングをセットアップしてもらえますか？

> Yes, for this Beta offering we can put you in touch with publishers interested and get some conversations going.

はい。このベータオファーからあなたとパブリッシャーがコンタクトをできるようにします。
