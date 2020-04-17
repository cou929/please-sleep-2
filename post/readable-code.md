{"title":"リーダブルコードを読んだ","date":"2014-05-17T20:03:01+09:00","tags":["book"]}

<iframe src="http://rcm-fe.amazon-adsystem.com/e/cm?lt1=_blank&bc1=000000&IS2=1&bg1=FFFFFF&fc1=000000&lc1=0000FF&t=pleasesleep-22&o=9&p=8&l=as4&m=amazon&f=ifr&ref=ss_til&asins=4873115655" style="width:120px;height:240px;" scrolling="no" marginwidth="0" marginheight="0" frameborder="0"></iframe>

リーダブルコードを読んだ。評判がいい本だけあってとてもいい内容だった。

良いコード == 読みやすいコードとして、読みやすいコードを書くための実践的な tips がまとめられている。そうだよねと納得するところもあり、なるほどと感心するところもあり。

特に面白かったのは中盤の、複雑な関数をシンプルに分割するための方法論。「無関係の下位問題を抽出する」「一度に一つのことをする」「言葉で説明してみる」という 3 つの方法を紹介している。それぞれ別角度からの、コードを改善できる考え方のヒントだ。内容にはもちろん納得。それに、このような抽象的で経験的なことを、わかりやすい言葉にまで落として説明できるのはすごいとも思った。

以下は内容のメモ。

### 表面的な変更

#### 命名

- 曖昧さを避け、カラフルな単語を使う。別のニュアンスでとられないか再考する。
- スコープの広さと名前の長さは比例する

#### レイアウト

レイアウトの原則はそのままコードにも適用できる。

- 一貫性をもたせる
- 関連したものは近くに配置する
- 似たもには似たレイアウトにする

#### コメント

読み手に意図を説明するのが目的。読み手のコードの理解をサポートするのが目的。

- 関数名などで明確に伝わることはコメントにしない。むしろ名前付けが大事
- その実装を選んだ経緯など、コード外のこと簡潔に
- 読みての立場にたつ。はまりそうなところを先回りしてコメント。
- とにかく簡潔に。

#### 制御フロー

- 条件式は 調査対象 - 比較対象
- if else に書く条件は、「関心の対象」を先にする。
- 関数からは早く返すそうすることでネストを浅くできる。深いネストは良くない

#### 大きい式は分割する

- 説明変数、要約変数の導入
- 共通部分のマクロ化

こういった手法で式に名前をつけて、文章のようなコードにする。

#### 変数についてさらに

- いらない変数はなくす。
  - dry に貢献していない。説明力があがらない。中間結果を保存するものなど
- 変数のスコープは最小に
- 変数宣言は使用場所の直前に
- 変数への書き込みは一度に (mutable)

### 構造的な変更

#### 無関係の下位問題の抽出

あるコードの

- 高度な目的
- 無関係の下位問題

を分けて考えて、後者は抽象化する。汎用的ならばユーティリティとして取り出す。

やりすぎは禁物。抽象化は少なからず可読性を下げる。天秤にかけてメリットがあると判断できれば実施する。

マーティンファウラーやケントベックの本にも出る話題。リファクタリングの「メソッドの抽出」パターンや、「メソッドないの処理は同じ抽象レベルにする原則」も参考に。

#### 一度に一つのことをする

関数が複雑になったら、

- その関数のタスクを洗い出し、
- そのタスクを切り出せるのならば別関数に
- そうでないならばコード中の “段落” としてまとめる

このようにして、一部分でひとつの処理だけをするように書き換える。

#### 言葉で説明する

処理が複雑な部分は、

- 処理を言葉で説明する
- それをコードに書き下す

ことで、整理でき、また人間の思考に近い理解しやすいコードになる。

#### コードベースを小さく保つ

- 過剰な機能はつけない
- 必要そうな機能を、もっと簡易な実装で済ませられないか考える
- ライブラリ、API でできることを日頃から把握しておく

### 参考文献

#### 高品質のコードを書くための書籍

<iframe src="http://rcm-fe.amazon-adsystem.com/e/cm?lt1=_blank&bc1=000000&IS2=1&bg1=FFFFFF&fc1=000000&lc1=0000FF&t=pleasesleep-22&o=9&p=8&l=as4&m=amazon&f=ifr&ref=ss_til&asins=B00JEYPPOE" style="width:120px;height:240px;" scrolling="no" marginwidth="0" marginheight="0" frameborder="0"></iframe>

<iframe src="http://rcm-fe.amazon-adsystem.com/e/cm?lt1=_blank&bc1=000000&IS2=1&bg1=FFFFFF&fc1=000000&lc1=0000FF&t=pleasesleep-22&o=9&p=8&l=as4&m=amazon&f=ifr&ref=ss_til&asins=4894712288" style="width:120px;height:240px;" scrolling="no" marginwidth="0" marginheight="0" frameborder="0"></iframe>

<iframe src="http://rcm-fe.amazon-adsystem.com/e/cm?lt1=_blank&bc1=000000&IS2=1&bg1=FFFFFF&fc1=000000&lc1=0000FF&t=pleasesleep-22&o=9&p=8&l=as4&m=amazon&f=ifr&ref=ss_til&asins=4756136494" style="width:120px;height:240px;" scrolling="no" marginwidth="0" marginheight="0" frameborder="0"></iframe>

<iframe src="http://rcm-fe.amazon-adsystem.com/e/cm?lt1=_blank&bc1=000000&IS2=1&bg1=FFFFFF&fc1=000000&lc1=0000FF&t=pleasesleep-22&o=9&p=8&l=as4&m=amazon&f=ifr&ref=ss_til&asins=4894712741" style="width:120px;height:240px;" scrolling="no" marginwidth="0" marginheight="0" frameborder="0"></iframe>

<iframe src="http://rcm-fe.amazon-adsystem.com/e/cm?lt1=_blank&bc1=000000&IS2=1&bg1=FFFFFF&fc1=000000&lc1=0000FF&t=pleasesleep-22&o=9&p=8&l=as4&m=amazon&f=ifr&ref=ss_til&asins=4048676881" style="width:120px;height:240px;" scrolling="no" marginwidth="0" marginheight="0" frameborder="0"></iframe>

#### プログラミングに関する書籍

<iframe src="http://rcm-fe.amazon-adsystem.com/e/cm?lt1=_blank&bc1=000000&IS2=1&bg1=FFFFFF&fc1=000000&lc1=0000FF&t=pleasesleep-22&o=9&p=8&l=as4&m=amazon&f=ifr&ref=ss_til&asins=4873113911" style="width:120px;height:240px;" scrolling="no" marginwidth="0" marginheight="0" frameborder="0"></iframe>

<iframe src="http://rcm-fe.amazon-adsystem.com/e/cm?lt1=_blank&bc1=000000&IS2=1&bg1=FFFFFF&fc1=000000&lc1=0000FF&t=pleasesleep-22&o=9&p=8&l=as4&m=amazon&f=ifr&ref=ss_til&asins=4621066056" style="width:120px;height:240px;" scrolling="no" marginwidth="0" marginheight="0" frameborder="0"></iframe>

<iframe src="http://rcm-fe.amazon-adsystem.com/e/cm?lt1=_blank&bc1=000000&IS2=1&bg1=FFFFFF&fc1=000000&lc1=0000FF&t=pleasesleep-22&o=9&p=8&l=as4&m=amazon&f=ifr&ref=ss_til&asins=4797311126" style="width:120px;height:240px;" scrolling="no" marginwidth="0" marginheight="0" frameborder="0"></iframe>

<iframe src="http://rcm-fe.amazon-adsystem.com/e/cm?lt1=_blank&bc1=000000&IS2=1&bg1=FFFFFF&fc1=000000&lc1=0000FF&t=pleasesleep-22&o=9&p=8&l=as4&m=amazon&f=ifr&ref=ss_til&asins=4621066072" style="width:120px;height:240px;" scrolling="no" marginwidth="0" marginheight="0" frameborder="0"></iframe>

<iframe src="http://rcm-fe.amazon-adsystem.com/e/cm?lt1=_blank&bc1=000000&IS2=1&bg1=FFFFFF&fc1=000000&lc1=0000FF&t=pleasesleep-22&o=9&p=8&l=as4&m=amazon&f=ifr&ref=ss_til&asins=487311361X" style="width:120px;height:240px;" scrolling="no" marginwidth="0" marginheight="0" frameborder="0"></iframe>

<iframe src="http://rcm-fe.amazon-adsystem.com/e/cm?lt1=_blank&bc1=000000&IS2=1&bg1=FFFFFF&fc1=000000&lc1=0000FF&t=pleasesleep-22&o=9&p=8&l=as4&m=amazon&f=ifr&ref=ss_til&asins=4274066304" style="width:120px;height:240px;" scrolling="no" marginwidth="0" marginheight="0" frameborder="0"></iframe>

####  歴史的記録

<iframe src="http://rcm-fe.amazon-adsystem.com/e/cm?lt1=_blank&bc1=000000&IS2=1&bg1=FFFFFF&fc1=000000&lc1=0000FF&t=pleasesleep-22&o=9&p=8&l=as4&m=amazon&f=ifr&ref=ss_til&asins=4756103642" style="width:120px;height:240px;" scrolling="no" marginwidth="0" marginheight="0" frameborder="0"></iframe>

<iframe src="http://rcm-fe.amazon-adsystem.com/e/cm?lt1=_blank&bc1=000000&IS2=1&bg1=FFFFFF&fc1=000000&lc1=0000FF&t=pleasesleep-22&o=9&p=8&l=as4&m=amazon&f=ifr&ref=ss_til&asins=B00BBDLIME" style="width:120px;height:240px;" scrolling="no" marginwidth="0" marginheight="0" frameborder="0"></iframe>

<iframe src="http://rcm-fe.amazon-adsystem.com/e/cm?lt1=_blank&bc1=000000&IS2=1&bg1=FFFFFF&fc1=000000&lc1=0000FF&t=pleasesleep-22&o=9&p=8&l=as4&m=amazon&f=ifr&ref=ss_til&asins=4320020855" style="width:120px;height:240px;" scrolling="no" marginwidth="0" marginheight="0" frameborder="0"></iframe>

<iframe src="http://rcm-fe.amazon-adsystem.com/e/cm?lt1=_blank&bc1=000000&IS2=1&bg1=FFFFFF&fc1=000000&lc1=0000FF&t=pleasesleep-22&o=9&p=8&l=as4&m=amazon&f=ifr&ref=ss_til&asins=4756101909" style="width:120px;height:240px;" scrolling="no" marginwidth="0" marginheight="0" frameborder="0"></iframe>
