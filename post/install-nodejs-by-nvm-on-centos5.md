{"title":"CentOS 5 に nvm で node をインストールする際の wrokaround","date":"2012-12-04T22:53:54+09:00","tags":["javascript, nix"]}

いまどき CentOS 5 系, しかも nvm は無いだろと思うかもしれませんが, そういう人もいるかもしれないので, インストールするためのもろもろの手順を残しておきます. とてもひどい内容なので注意のこと.

### nvm をいれる

ここは困ることはないです

    $ git clone git://github.com/creationix/nvm.git ~/nvm
    $ . ~/nvm/nvm.sh

いまさら nvm を使おうと思っている人に zsh 使いはいないと思いますが, zsh だと `nvm_ls:1: no matches found: vdefault*` みたいな warning がでたり `remote-ls` が動かなかったりしますが, インストールやバージョンの切り替えは一応大丈夫です. ただ nvm は zsh 対応することは無いと思うので素直に [nodebrew](https://github.com/hokaccha/nodebrew) を使ったほうが良いと思います.

### node をインストール

インストールの前に shasum を準備する必要があります. nvm はダウンロードしたバイナリ・コードのチェックに shasum を使っていますが, CentOS には入っていません. 準備しておかないとコマンドが無いよと怒られてしまします.

shasum を新たにインストールしてもいいんですが, sha1sum にエイリアスをはるのが楽だと思います.

    $ alias shasum='sha1sum'

この件いちおう pull request は投げてるんですが, たぶん取り込まれないだろうなあ.

次は glibc 対応です. 最近のバージョンの node のバイナリを持ってきてもうまく動きません. glibc が古いからで, こんなエラーがでます

    $ nvm install v0.8.14
    ######################################################################## 100.0%
    Now using node v0.8.14
    $ node
    node: /lib64/libc.so.6: version `GLIBC_2.9' not found (required by node)
    node: /lib64/libc.so.6: version `GLIBC_2.6' not found (required by node)
    node: /lib64/libc.so.6: version `GLIBC_2.7' not found (required by node)

glibc は 2.5 が入っているようです

    $ yum info glibc
    
    Name       : glibc
    Arch       : x86_64
    Version    : 2.5
    Release    : 65.el5_7.1
    Size       : 11 M
    ...

glibc を新しくしてもいいのですが, 今回はバイナリではなくてソースからコンパイルして対応したいと思います. ものすごくしょうもないですが nvm.sh からバイナリをダウンロードしている処理を一時的に除いてあるげことにしました. 203 行目の if で始まるブロックに入らないようにすればなんでも OK です.

    diff --git a/nvm.sh b/nvm.sh
    index bc42df2..83ad795 100755
    --- a/nvm.sh
    +++ b/nvm.sh
    @@ -201,6 +201,7 @@ nvm()
     
           # shortcut - try the binary if possible.
           if [ -n "$os" ]; then
    +      else
             binavail=
             # binaries started with node 0.8.6
             case "$VERSION" in

これでソースからコンパイルしてのインストールが強制されるようになりました.

次は python です. CentOS 5x に入っている python は例のごとく 2.4 です. node の configure スクリプトは python で書かれているんですが, `A if B else c` という三項演算子の書き方をしているところがあります. この機能は python 2.5 以降でしか使えないんわけなんですね

    $ nvm install v0.8.14
    Additional options while compiling:
    ######################################################################## 100.0%
      File "./configure", line 355
        1 if options.unsafe_optimizations else 0)
           ^
    SyntaxError: invalid syntax
    nvm: install v0.8.14 failed!

epel に python2.6 があるのでそいつを入れます. まだ epel の設定がない場合は http://ftp-srv2.kddilabs.jp/Linux/distributions/fedora/epel/5/x86_64/epel-release-5-4.noarch.rpm などから rpm をもってきて入れておいて,

    $ sudo yum install python26
    $ python26 -V
    Python 2.6.8
    $ python -V
    Python2.4.3

configure スクリプトは `/usr/bin/env python` とやっているので, `/usr/bin/python` のシンボリックリンクを一旦切り替えてあげることにしました. (またひどい対応ですが)

    $ ls -l /usr/bin/python*
    -rwxr-xr-x 2 root root 8304  9月 22  2011 /usr/bin/python
    lrwxrwxrwx 1 root root    6  3月  1  2012 /usr/bin/python2 -> python
    -rwxr-xr-x 2 root root 8304  9月 22  2011 /usr/bin/python2.4
    -rwxr-xr-x 2 root root 4736 11月  7 23:48 /usr/bin/python2.6
    -rwxr-xr-x 2 root root 4736 11月  7 23:48 /usr/bin/python26
    $ sudo rm /usr/bin/python
    $ sudo ln -s /usr/bin/python26 /usr/bin/python

これでビルドが通るようになるかと思います.

    $ nvm install v0.8.14
    $ node -v
    v0.8.14

終わったあとは python のシンボリックリンクをもとに戻しておきましょう

    $ sudo rm /usr/bin/python
    $ sudo ln -s /usr/bin/python2 /usr/bin/python

### まとめ

かなりしょうもない対応をたくさんすることになるので, 素直に, CentOS は 6 系にしましょう. [nodebrew](https://github.com/hokaccha/nodebrew) をつかいましょう.
