{"title":"CentOS 5.6 で zsh をバージョンアップ","date":"2012-12-18T08:16:32+09:00","tags":["nix"]}

ちゃんと調べてないけど, [zsh-syntax-highlighting](https://github.com/zsh-users/zsh-syntax-highlighting) がうまく動かず, 原因はバージョンのようだったのでアップデートしてみた.

yum にあるのは 4.2.6

    $ zsh --version
    zsh 4.2.6 (x86_64-redhat-linux-gnu)

まあソースからやるしか無い. 最新のソースはここでチェックのこと [ZSH - Source download](http://zsh.sourceforge.net/Arc/source.html)

    $ wget http://downloads.sourceforge.net/project/zsh/zsh/5.0.0/zsh-5.0.0.tar.gz
    $ tar zxvf zsh-5.0.0.tar.gz
    $ cd zsh-5.0.0
    $ ./configure --enable-multibyte
    $ make
    $ sudo make install
    $ ll /usr/local/bin/zsh*
    -rwxr-xr-x 2 root root 666112 12月 17 14:56 /usr/local/bin/zsh
    -rwxr-xr-x 2 root root 666112 12月 17 14:56 /usr/local/bin/zsh-5.0.0

リンクを貼り直してあげた

    $ ll /usr/local/bin/zsh*
    lrwxrwxrwx 1 root root     24 12月 17 14:57 /usr/local/bin/zsh -> /usr/local/bin/zsh-5.0.0
    -rwxr-xr-x 1 root root 666112 12月 17 14:57 /usr/local/bin/zsh-4.2.6
    -rwxr-xr-x 1 root root 666112 12月 17 14:56 /usr/local/bin/zsh-5.0.0

    $ zsh --version
    zsh 5.0.0 (x86_64-unknown-linux-gnu)

これでバージョンアップできたけど, zsh-syntax-highlighting は依然うごかなかったので別途しらべる
