{"title":"CentOS に tig をインストール","date":"2012-12-18T08:16:56+09:00","tags":["nix"]}

ソースから入れてあげる. 手順はごく普通. ビルドもすぐ終わる

    $ curl http://jonas.nitro.dk/tig/releases/tig-1.1.tar.gz -o tig-1.1.tar.gz
    $ tar -zxvf tig-1.1.tar.gz
    $ cd tig-1.1/
    $ ./configure
    $ make
    $ sudo make install
