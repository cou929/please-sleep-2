{"title":"AWSプロダクトシリーズ よくわかるAmazon DynamoDB","date":"2014-10-21T22:27:04+09:00","tags":["etc"]}

[AWSプロダクトシリーズ｜よくわかるAmazon DynamoDB](http://kokucheese.com/event/index/226052/) に参加したメモ。

所感として、広告配信のオーディエンスデータストアとして使用している実績が聞けてよかった。単なる KVS 以上の使い方の事例を作ろうとしているようで (急進的なところでは RDB をすべて置き換えようとしているとか)、今そうする有機は全くないので、そのへんは続報を待ちたい。

## Amazon DynamoDB最新情報 (Khawaja Shams)

### Dynamo

- 2006 年の論文
- [http://www.allthingsdistributed.com/files/amazon-dynamo-sosp2007.pdf](http://www.allthingsdistributed.com/files/amazon-dynamo-sosp2007.pdf)

### free tier

- 25/sec read
- 25/sec write
- 25 GB

### 用語

- Table
  - Item が入る
- Item
  - RDB の行みたいな
  - Attribute を持つ
  - pk で分散
- Attribute
  - key value のデータ

### 型

- string, number, set, map, boolean ...

### Beyond key-value: Range key

- hash key で特定した item を絞り込む
- secondary index
  - local secondary index
     - 2 つ目のレンジキー
     - user で絞って、かつ日付で絞るなど
  - global secondary index

### JSON

- native full json support
- JSON データ型
- json をどう格納するかを dynamodb が面倒をみてくれる
- アプリ側で json encode/decode 不要
- 深い階層のキーにもアクセスできる
- 特定のキーの UPDATE

### Optimistic concurrency control

- non rdbmd で高いスケーラビリティを実現するために重要な考え方
- global lock -> ロック待ちが長くなる
- row level lock -> 直列 get しかできない。同時に平行しての get ができない
  - 悲観的なロック。RDBMS でよく使われる
- この解決が楽観的な同時実行性制御
- 条件付き put を使う
  - 失敗したら現在値を取得しなおして再度 put する
  - つぎの alex の講演で例を示します

## Amazon DynamoDBアプリケーション開発のポイント (Alexander Patrikalakis)

### DynamoDB Local

- [http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Tools.DynamoDBLocal.html](http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Tools.DynamoDBLocal.html)
- オフラインでも手元で開発できる
- javascript shell もある

### サンプルコード

- [https://s3-ap-northeast-1.amazonaws.com/dynamodb/dorayaki.txt](https://s3-ap-northeast-1.amazonaws.com/dynamodb/dorayaki.txt)
- これを javascript shell につっこめば動く

### tips

- scan できるのは最大 1MB まで
- getItem
  - hash key で item を引く
- conditionalUpdateItem
  - 条件に合えば更新し、合わなければ修正した値 (この場合はあんこを一個追加する) で再度更新を試みる
     - 一定回数まで再帰する

<pre><code data-language="javascript">
    var updateHistoryOfCookiesEaten = function(item) {
        var ankolist = ['餡子']; //追加するどら焼きの種類
        //食歴に餡子を追加する
        item['食歴']['給仕1'] = item['食歴']['給仕1'].concat(ankolist);
        var params = {
            TableName: 'dorayaki',
            Key: { '氏名': item['氏名'] },
            AttributeUpdates: {
                '個数': { Action: 'ADD', Value: 1 },
                '食歴': { Action: 'PUT', Value: item['食歴'] }
            },
            Expected: [ dynamodb.Condition('個数', 'EQ', item['個数']) ],
            ReturnValues: 'UPDATED_NEW'
        };

        dynamodb.updateItem(params, function(err, data) {
            //更新APIの呼び出しが成功すれば
            if (!err) {
                //その戻り値を出力する
                print(data);
            } else {
                //同時に別ユーザーが読み込み時と更新時の間に同アイテムを更新
                //したりすると、ConditionalCheckFailedExceptionの例外が発生することがある
                if(err.code === 'ConditionalCheckFailedException'){
                    //その時は条件付き更新を帰納的に試みる
                    conditionalUpdateItem(item, ++nRetries);
                } else {
                    //予期しない例外が発生したのでとりあえず出力する
                    print('UpdateItem failed: ' + err);
                }
            }
        });
    };
</code></pre>

- expressionUpdateItem
  - クエリのようなもの (expression) で複数キーを一度に更新する例

<pre><code data-language="javascript">
//アイテムを読み込まずに、一度にサーバー側で複数の更新をする
var expressionUpdateItem = function(name) {
   var params = {
       TableName: 'dorayaki',
       // The primary key of the item (a map of attribute name to AttributeValue)
       Key: { '氏名': name },
       // first, set the increment the count (kosuu)
       // second, append anko to the list of dorayaki that server1 served doraemon
       UpdateExpression: 'SET #count = #count + :one, #history_of_cookies_eaten.#server1 = list_append(#history_of_cookies_eaten.#server1, :ankolist)',
       ExpressionAttributeNames : { '#count':'個数',
                                    '#server1':'給仕1',
                                    '#history_of_cookies_eaten':'食歴'
       },
       ExpressionAttributeValues : {
            ':one': 1,
            ':ankolist' :  ['餡子']

       },
       ReturnValues: 'UPDATED_NEW' // optional (NONE | ALL_OLD | UPDATED_OLD | ALL_NEW | UPDATED_NEW)
   };
   dynamodb.updateItem(params, printDynamoDBResult);
};
</code></pre>

## 質疑応答

- global secondary index の更新速度
  - 10 ms くらいでいける
- dynamodb javascript shell はどこにある?
  - dynamodb local を落として、`http://localohst:8000/shell/`
- primary と secondary の違い。どっちのカラムを primary にすべきかなど
  - base table primary key = hash key, hash range key
  - local secondary index = 同じテーブルの hash key にレンジキーを付ける機能
     - データが格納されているノードに、ローカルにつけられる index
     - 同期的に index が更新される
  - global secondary index = 全く別のレンジキー
     - base table から projection して非同期でつく
     - 結果整合性
     - もうすぐ、テーブルを作った後に gsi を変更できるようになる
  - strong consistency - eventual consistency
     - gsi は eventual consistency (結果整合性) だけが担保される
     - [参考](http://www.publickey1.jp/blog/09/eventual_consistency.html)
- 使いどころ
  - 新しいアプリケーションでスケールさせたいものがあれば使って欲しい
     - 既存のデータの移行よりも新しいプロジェクトで使って欲しい
  - [AdRoll](https://www.adroll.com/) での採用実績
     - [Real-Time Ad Impression Bids Using DynamoDB](http://aws.amazon.com/jp/blogs/aws/real-time-ad-impression-bids-using-dynamodb/)
     - リタゲのためのオーディエンスデータに使っているようだ
     - RTB をやっている
     - 7 billion impressions per day
     - more than a billion cookie profiles stored in DynamoDB
  - [battle camp](https://itunes.apple.com/us/app/battle-camp/id569929985?mt=8) は dynamodb だけで実装されている
  - エンタープライズで RDB を DynamoDB にリプレイスしようとしている人達もいる
     - まじで..
  - 可用性を重視するのであれば、DynamoDB はよい選択肢
- バックアップ、スナップショット
  - インクリメンタルなバックアップがとれる機能を検討中なので期待してね
     - 別の人からも要望があった
  - 整合性を気にしなければオンラインバックアップ可能
- gsi の個数
  - 8 個くらいはりたいっていうリクエスト。開発にフィードバックします
- バックアップからのリストア時間の目安
  - そのテーブルに対してどのくらいの IOPS に設定しているかに、一番おおきく依存する
- 最新情報は [twiiter](https://twitter.com/ksshams) でフォローしてね。インクリメンタルなバックアップを実装したらすぐにつぶやくよ

※ 質問はいくつかメモ漏れあり
