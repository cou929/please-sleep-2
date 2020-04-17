{"title":"node.js の build には python2.6 or 2.7 が必要","date":"2013-12-27T08:19:12+09:00","tags":["javascript"]}

新調した macbook、nodebrew で node.js をインストールしようとおもったら、configure スクリプトの syntax エラーで死んでしまってできない。

    ######################################################################## 100.0%
      File "./configure", line 319
        '''
          ^
    SyntaxError: invalid syntax

どうも python のバージョンが合っていないようだ。そういえば手元の環境は python3.3 をデフォルトにしていた。

node.js の Prerequisites をみると `python 2.6 or 2.7` とあった。通りで。

[Installation · joyent/node Wiki](https://github.com/joyent/node/wiki/installation#prerequisites)

システムの python2.7 を使って解決

    $ mkvirtualenv -p `which python2.7` py27
    $ nodebrew install v0.10.24

