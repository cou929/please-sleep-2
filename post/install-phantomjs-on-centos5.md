{"title":"CentOS5 系に phantomjs をインストールするのが超簡単だった","date":"2013-04-10T00:58:02+09:00","tags":["nix"]}

[PhantomJS: Download and Install](http://phantomjs.org/download.html)

ここにバイナリがおいてあるのでダウンロードして展開してパスを通すだけで大丈夫。qt とか webkit をインストールする必要すらなくて超簡単。

古いシステムなどでソースから入れたい場合は [PhantomJS: Build Instructions](http://phantomjs.org/build.html) に従う。以下のステップだけでいいらしい

    sudo yum install gcc gcc-c++ make git openssl-devel freetype-devel fontconfig-devel
    git clone git://github.com/ariya/phantomjs.git
    cd phantomjs
    git checkout 1.9
    ./build.sh

ソースから入れる場合も同様に qt などは必要なくて、build.sh を叩くだけでいいそうだ。簡単。

[Phantom JS on Centos5 « rhythmicalmedia.com](http://rhythmicalmedia.com/?p=146) のような古い記事を見ると、昔はいろいろ依存するものを準備しないといけなくて大変そうだけど、phantomjs 1.5 以降ではビルドスクリプトがいろいろめんどうみてくれるように改良されたらしい。

[The evolution of PhantomJS build workflow](http://ariya.ofilabs.com/2012/03/the-evolution-of-phantomjs-build-workflow.html)
