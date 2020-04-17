{"title":"readlists-anonymous で High-Performance Browser Networking を電子書籍形式で読む","date":"2014-01-11T23:28:41+09:00","tags":["tips"]}

[アドベントカレンダーを電子書籍で読めるサービスを作った - 2nd life](http://secondlife.hatenablog.jp/entry/2014/01/11/170103)

こちらが素敵だった。

そもそも [Readlists](http://readlists.com/) というサービスを知らなかったが、任意の url のリストを渡すと、それらを [ Readability](https://www.readability.com/) にかけて、結果を kindle やデバイスに送ったり epub を dropbox に保存したりできるものらしい。

コマンドラインから readlists をたたけるようにしたツールが [readlists-anonymous](https://github.com/hotchpotch/readlists-anonymous)。これを使わせてもらって、ちょうどよもうと思っていた [High Performance Browser Networking](http://chimera.labs.oreilly.com/books/1230000000545/index.html) の web で公開されている無料版を epub に変換した。

    $ gem install readlists-anonymous
    $ readlists-anonymous -t 'High-Performance Browser Networking' http://chimera.labs.oreilly.com/books/1230000000545/index.html http://chimera.labs.oreilly.com/books/1230000000545/pr01.html http://chimera.labs.oreilly.com/books/1230000000545/pr02.html http://chimera.labs.oreilly.com/books/1230000000545/ch01.html http://chimera.labs.oreilly.com/books/1230000000545/ch02.html http://chimera.labs.oreilly.com/books/1230000000545/ch03.html http://chimera.labs.oreilly.com/books/1230000000545/ch04.html http://chimera.labs.oreilly.com/books/1230000000545/ch05.html http://chimera.labs.oreilly.com/books/1230000000545/ch06.html http://chimera.labs.oreilly.com/books/1230000000545/ch07.html http://chimera.labs.oreilly.com/books/1230000000545/ch08.html http://chimera.labs.oreilly.com/books/1230000000545/ch09.html http://chimera.labs.oreilly.com/books/1230000000545/ch10.html http://chimera.labs.oreilly.com/books/1230000000545/ch11.html http://chimera.labs.oreilly.com/books/1230000000545/ch12.html http://chimera.labs.oreilly.com/books/1230000000545/ch13.html http://chimera.labs.oreilly.com/books/1230000000545/ch14.html http://chimera.labs.oreilly.com/books/1230000000545/ch15.html http://chimera.labs.oreilly.com/books/1230000000545/ch16.html http://chimera.labs.oreilly.com/books/1230000000545/ch17.html http://chimera.labs.oreilly.com/books/1230000000545/ch18.html http://chimera.labs.oreilly.com/books/1230000000545/ix01.html

こんなかんじに URL を列挙するだけで良い。非常に簡単。できあがったものはこちら。

[High-Performance Browser Networking - Readlists](http://readlists.com/750e80cc/)

ところで、readlists の API について Web 上にはどこにも記述がなくて、ただ readlists をよく見てみると js が単純に `http://readlists.com/api/v1/*` のような api をたたいてできているサービスなので、readlists-anonymous はここをみて作ったのかなと思った。

