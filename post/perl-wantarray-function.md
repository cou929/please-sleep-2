{"title":"perl の wantarray 関数","date":"2012-10-06T13:06:40+09:00","tags":["perl"]}

- そのブロックのコンテキストに応じて戻り値が変わる
- スカラーコンテキストを要求されていたら偽, リストコンテキストなら真
- 関数から return するときにコンテキストに合わせて配列を返すかスカラーを返すか切り替えるために使う
- たしかにコンテキストがある限りこのような関数は必要なのはわかる
- perldoc には wantlist という名前にすべきと書いてあった
