{"title":"centos に gitolite 導入, あとロケール問題","date":"2013-10-05T15:41:54+09:00","tags":["nix"]}

このとおりにやる

[gitolite 導入 - Please Sleep](http://please-sleep.cou929.nu/intorducing-gitolite.html)

まず gitolite の `gl-system-install` が通らない. `File::Path` の `make_path` が無いと言われる. 調べると perl のバージョンがすごく古い

    cou929:kosei% perl -v
    
    This is perl, v5.8.8 built for x86_64-linux-thread-multi
    cou929:kosei% perl -MFile::Path -le 'print $File::Path::VERSION'
    1.08

少なくともこのバージョンの File::Path には `make_path` は無いようだ

ちょっとぐぐってみたけど手軽にシステム perl のバージョンをあげる方法はなさそう. yum のリポジトリ追加していけないかなと思ったけど, 少なくとも epel, rpmforge にはなかった.

どうしよう

1. 自分でビルド
2. perlbrew
3. gitolite 使わない (素でリポジトリをたてる)

1 は嫌. 手っ取り早いのは 3 だけど, 後学のために 2 を試してみる. ただし perlbrew はたぶんローカルに perl を入れるわけだし, そこをどう対応しようか. gitolite ユーザで perlbrew を導入するのがよさそうか

    cou929:kosei% su gitolite
    Password:
    sh-3.2$ id
    uid=503(gitolite) gid=503(gitolite) groups=503(gitolite)
    sh-3.2$ bash
    [gitolite@cou929 kosei]$ cd
    [gitolite@cou929 ~]$ pwd
    /home/gitolite
    [gitolite@cou929 ~]$ curl -kL http://xrl.us/perlbrew > perlbrew
    [gitolite@cou929 ~]$ perl ./perlbrew init
    [gitolite@cou929 ~]$ perl ./perlbrew install
    [gitolite@cou929 ~]$ vi .bash_profile
    [gitolite@cou929 ~]$ cat .bash_profile
    ...
    # perlbrew
    source ~/perl5/perlbrew/etc/bashrc
    [gitolite@cou929 ~]$ source ~/.bash_profile
    [gitolite@cou929 ~]$ which perlbrew
    ~/perl5/perlbrew/bin/perlbrew
    [gitolite@cou929 ~]$ perlbrew install perl-5.12.3

テストで落ちたのでとりあえず force

    [gitolite@cou929 ~]$ perlbrew install --force perl-5.12.3
    [gitolite@cou929 ~]$ perlbrew switch perl-5.12.3
    [gitolite@cou929 ~]$ perl -v
    
    This is perl 5, version 12, subversion 3 (v5.12.3) built for x86_64-linux
    
    Copyright 1987-2010, Larry Wall
    
    Perl may be copied only under the terms of either the Artistic License or the
    GNU General Public License, which may be found in the Perl 5 source kit.
    
    Complete documentation for Perl, including FAQ lists, should be found on
    this system using "man perl" or "perldoc perl".  If you have access to the
    Internet, point your browser at http://www.perl.org/, the Perl Home Page.
    
    [gitolite@cou929 ~]$ perl -MFile::Path -le 'print $File::Path::VERSION'
    2.08_01

gitolite セットアップを続行

    [gitolite@cou929 ~]$ sudo perl gitolite/src/gl-system-install
    using default values for EUID=0:
    /usr/local/bin, /var/gitolite/conf, /var/gitolite/hooks
                    ***** WARNING *****
        gl-setup is not in your $PATH.
    
        Since gl-setup MUST be run from the PATH (and not as src/gl-setup or
        such), you must fix this before running gl-setup.  Just add
    
            PATH=/usr/local/bin:$PATH
    
        to the end of your bashrc or similar file.  You can even simply do that
        manually each time you log in and want to run a gitolite command.
    [gitolite@cou929 ~]$ vi .bashrc
    [gitolite@cou929 ~]$ cat .bashrc
    # .bashrc
    
    # Source global definitions
    if [ -f /etc/bashrc ]; then
            . /etc/bashrc
    fi
    
    # User specific aliases and functions
    PATH=/usr/local/bin:$PATH
    [gitolite@cou929 ~]$ ls -l /var/gitolite/
    total 8
    drwxr-xr-x 2 root root 4096 Feb 25 10:39 conf
    drwxr-xr-x 4 root root 4096 Feb 25 10:39 hooks
    [gitolite@cou929 ~]$ ls -l /usr/local/bin/gl*
    -rwxr-xr-x 1 root root  1358 Feb 25 10:39 /usr/local/bin/gl-admin-push
    -rwxr-xr-x 1 root root  8199 Feb 25 10:39 /usr/local/bin/gl-auth-command
    -rwxr-xr-x 1 root root 23792 Feb 25 10:39 /usr/local/bin/gl-compile-conf
    -rwxr-xr-x 1 root root  3396 Feb 25 10:39 /usr/local/bin/gl-conf-convert
    -rwxr-xr-x 1 root root  3392 Feb 25 10:39 /usr/local/bin/gl-dryrun
    -rwxr-xr-x 1 root root  2213 Feb 25 10:39 /usr/local/bin/gl-install
    -rwxr-xr-x 1 root root  2307 Feb 25 10:39 /usr/local/bin/gl-mirror-push
    -rwxr-xr-x 1 root root  6082 Feb 25 10:39 /usr/local/bin/gl-mirror-shell
    -rwxr-xr-x 1 root root   594 Feb 25 10:39 /usr/local/bin/gl-query-rc
    -rwxr-xr-x 1 root root  4509 Feb 25 10:39 /usr/local/bin/gl-setup
    -rwxr-xr-x 1 root root  1723 Feb 25 10:39 /usr/local/bin/gl-setup-authkeys
    -rwxr-xr-x 1 root root  4173 Feb 25 10:39 /usr/local/bin/gl-system-install
    -rwxr-xr-x 1 root root  1507 Feb 25 10:39 /usr/local/bin/gl-time
    -rwxr-xr-x 1 root root  3019 Feb 25 10:39 /usr/local/bin/gl-tool
    [gitolite@cou929 ~]$ gl-setup -q /tmp/kosei.pub
    creating gitolite-admin...
    Initialized empty Git repository in /home/gitolite/repositories/gitolite-admin.git/
    creating testing...
    Initialized empty Git repository in /home/gitolite/repositories/testing.git/
    [master (root-commit) 97d9447] start
     2 files changed, 8 insertions(+), 0 deletions(-)
     create mode 100644 conf/gitolite.conf
     create mode 100644 keydir/kosei.pub

リポジトリ追加してみる

    Koseis-MBA:kosei% git clone ssh://gitolite@cou929.nu/gitolite-admin.git
    Koseis-MBA:kosei% cd gitolite-admin/
    
    # 編集する
    
    Koseis-MBA:kosei% git diff
    diff --git a/conf/gitolite.conf b/conf/gitolite.conf
    index 9719ad2..3ab2985 100644
    --- a/conf/gitolite.conf
    +++ b/conf/gitolite.conf
    @@ -5,3 +5,5 @@ repo    gitolite-admin
     repo    testing
             RW+     =   @all
    
    +repo   please-sleep
    +       RW+     =       kosei
    Koseis-MBA:kosei% git commit -am 'add please-sleep repo'
    [master 9abe6a0] add please-sleep repo
     1 files changed, 2 insertions(+), 0 deletions(-)
    Koseis-MBA:kosei% git push origin master
    Counting objects: 7, done.
    Delta compression using up to 2 threads.
    Compressing objects: 100% (3/3), done.
    Writing objects: 100% (4/4), 381 bytes, done.
    Total 4 (delta 1), reused 0 (delta 0)
    remote: creating please-sleep...
    remote: Initialized empty Git repository in /home/gitolite/repositories/please-sleep.git/
    To ssh://gitolite@cou929.nu/gitolite-admin.git
       97d9447..9abe6a0  master -> master

クローン

    Koseis-MBA:kosei% git clone ssh://gitolite@cou929.nu/please-sleep.git
    Cloning into please-sleep...
    warning: You appear to have cloned an empty repository.

ok

### locale 問題

#### TL;DR

結論としては `/etc/ssh_config` の `SendEnv LANG LC_*` をコメントアウトすればよい。

#### 経緯

この間実はずっと perl のロケールのワーニングが出ていた. これに対応したい. こういうワーニング

    perl: warning: Setting locale failed.
    perl: warning: Please check that your locale settings:
            LANGUAGE = (unset),
            LC_ALL = (unset),
            LC_CTYPE = "UTF-8",
            LANG = "ja_JP.UTF-8"
        are supported and installed on your system.
    perl: warning: Falling back to the standard locale ("C").

そもそもデフォルトが `ja_JP.UTF-8` になっている.

    cou929:kosei% locale
    locale: Cannot set LC_CTYPE to default locale: No such file or directory
    locale: LC_ALL?????????????????????: ??????????????????????
    LANG=ja_JP.UTF-8
    LC_CTYPE=UTF-8
    LC_NUMERIC="ja_JP.UTF-8"
    LC_TIME="ja_JP.UTF-8"
    LC_COLLATE="ja_JP.UTF-8"
    LC_MONETARY="ja_JP.UTF-8"
    LC_MESSAGES="ja_JP.UTF-8"
    LC_PAPER="ja_JP.UTF-8"
    LC_NAME="ja_JP.UTF-8"
    LC_ADDRESS="ja_JP.UTF-8"
    LC_TELEPHONE="ja_JP.UTF-8"
    LC_MEASUREMENT="ja_JP.UTF-8"
    LC_IDENTIFICATION="ja_JP.UTF-8"
    LC_ALL=

これを C にすればよさそう. あるいは `locale-gen` で ja_JP.UTF-8 を作ってもいいそうだけど, そもそも日本語である必要が無いので (man とか文字化けしてるし) 変えたい

システムデフォルト設定は C になっているぽい

    cou929:kosei% cat /etc/sysconfig/i18n
    LANG="C"
    SYSFONT="latarcyrheb-sun16"

`ja_JP.utf8` てのはある

    cou929:kosei% locale -a | grep ja_JP
    locale: Cannot set LC_CTYPE to default locale: No such file or directory
    ja_JP
    ja_JP.eucjp
    ja_JP.ujis
    ja_JP.utf8

あとそもそも C ってなんなの

- [1 Entry per Day: What's "LANG=C" ?](http://mstssk.blogspot.com/2009/04/whats-langc.html)
- [UNIXのLANGの設定で，"C"と"us"って何が違うのでしょうか． そも.. - 人力検索はてな](http://q.hatena.ne.jp/1230614531)

C だと翻訳せずにソースコードそのまま出すという意味で, たいていのソフトは英語で書かれてるから実質英語になるらしい.

[第08回 「ロケールを正しく設定する」](http://landisk.kororo.jp/diary/08_locale.php)

ここが素晴らしくわかりやすかった

あの perl のエラーはやはり `ja_JP.UTF-8` がちゃんとインストールされていないことが原因みたいだ. `ja_JP.utf8`  てのしかない

    cou929:kosei% ls /usr/lib/locale/ja_JP*
    /usr/lib/locale/ja_JP.eucjp:
    LC_ADDRESS  LC_COLLATE  LC_CTYPE  LC_IDENTIFICATION  LC_MEASUREMENT  LC_MESSAGES  LC_MONETARY  LC_NAME  LC_NUMERIC  LC_PAPER  LC_TELEPHONE  LC_TIME
    
    /usr/lib/locale/ja_JP.utf8:
    LC_ADDRESS  LC_COLLATE  LC_CTYPE  LC_IDENTIFICATION  LC_MEASUREMENT  LC_MESSAGES  LC_MONETARY  LC_NAME  LC_NUMERIC  LC_PAPER  LC_TELEPHONE  LC_TIME

locale の追加は `localedef` ってコマンドでいいらしい (debian などは locale-gen らしい)

localedef の例

    $ localedef -f UTF-8 -i ja_JP ja_JP.UTF-8

ここで, -f には charmapfile のパス, つまり文字コードごとの表の場所を指定する. フルパスである必要はい. デフォルトのパスは `localedef --help` すると書いてある. 今回の場合は `/usr/share/i18n/charmaps/` だった.

    cou929:kosei% ls /usr/share/i18n/charmaps/UTF*
    /usr/share/i18n/charmaps/UTF-8.gz
    cou929:kosei% zcat /usr/share/i18n/charmaps/UTF-8.gz | head
    <code_set_name> UTF-8
    <comment_char> %
    <escape_char> /
    <mb_cur_min> 1
    <mb_cur_max> 6
    
    % alias ISO-10646/UTF-8
    CHARMAP
    <U0000>     /x00         NULL
    <U0001>     /x01         START OF HEADING

-i には inputfile というものを指定するらしい. `/usr/share/i18n/locales` 以下のファイルを指定する.

    cou929:kosei% ls /usr/share/i18n/locales/ja_JP
    /usr/share/i18n/locales/ja_JP

とりあえずやってみる

    cou929:kosei% sudo localedef -f UTF-8 -i ja_JP ja_JP.UTF-8

だめっぽい

    cou929:kosei% locale -a | grep ja_JP
    locale: Cannot set LC_CTYPE to default locale: No such file or directory
    ja_JP
    ja_JP.eucjp
    ja_JP.ujis
    ja_JP.utf8

よく見ると `LC_CTYPE` がエラーを吐いているように見える

    cou929:kosei% locale
    locale: Cannot set LC_CTYPE to default locale: No such file or directory
    locale: LC_ALL?????????????????????: ??????????????????????
    LANG=ja_JP.UTF-8
    LC_CTYPE=UTF-8
    LC_NUMERIC="ja_JP.UTF-8"
    LC_TIME="ja_JP.UTF-8"
    LC_COLLATE="ja_JP.UTF-8"
    LC_MONETARY="ja_JP.UTF-8"
    LC_MESSAGES="ja_JP.UTF-8"
    LC_PAPER="ja_JP.UTF-8"
    LC_NAME="ja_JP.UTF-8"
    LC_ADDRESS="ja_JP.UTF-8"
    LC_TELEPHONE="ja_JP.UTF-8"
    LC_MEASUREMENT="ja_JP.UTF-8"
    LC_IDENTIFICATION="ja_JP.UTF-8"
    LC_ALL=

export したら治った

    cou929:kosei% export LC_CTYPE='ja_JP.UTF-8'
    cou929:kosei% locale
    LANG=ja_JP.UTF-8
    LC_CTYPE=ja_JP.UTF-8
    LC_NUMERIC="ja_JP.UTF-8"
    LC_TIME="ja_JP.UTF-8"
    LC_COLLATE="ja_JP.UTF-8"
    LC_MONETARY="ja_JP.UTF-8"
    LC_MESSAGES="ja_JP.UTF-8"
    LC_PAPER="ja_JP.UTF-8"
    LC_NAME="ja_JP.UTF-8"
    LC_ADDRESS="ja_JP.UTF-8"
    LC_TELEPHONE="ja_JP.UTF-8"
    LC_MEASUREMENT="ja_JP.UTF-8"
    LC_IDENTIFICATION="ja_JP.UTF-8"
    LC_ALL=
    cou929:kosei% perl -v
    
    This is perl, v5.8.8 built for x86_64-linux-thread-multi
    
    Copyright 1987-2006, Larry Wall
    
    Perl may be copied only under the terms of either the Artistic License or the
    GNU General Public License, which may be found in the Perl 5 source kit.
    
    Complete documentation for Perl, including FAQ lists, should be found on
    this system using "man perl" or "perldoc perl".  If you have access to the
    Internet, point your browser at http://www.perl.org/, the Perl Home Page.

とりあえず `LC_CTYPE` を設定するとワーニングは解消できそう.

あとは `/etc/sysconfig/i18n` じゃないところでロケールの設定がされているようなので, それがどこかを探る必要がある. それさえわかれば全部 C にするなり ja_JP にするなりできる

`/etc/profile.d/lang.sh` があやしい (/etc/profile から実行されている)

うーん, i18n ファイルに書いてあるものを export しているだけに見える...

ssh -v でログインすると

    debug1: Sending environment.
    debug1: Sending env LC_CTYPE = UTF-8
    debug1: Sending env LANG = ja_JP.UTF-8

となっていることに気づいた

リモート側の `/etc/ssh/ssh_config` を見ると次のようになっている

    # Send locale-related environment variables
            SendEnv LANG LC_CTYPE LC_NUMERIC LC_TIME LC_COLLATE LC_MONETARY LC_MESSAGES
            SendEnv LC_PAPER LC_NAME LC_ADDRESS LC_TELEPHONE LC_MEASUREMENT
            SendEnv LC_IDENTIFICATION LC_ALL

確かにローカルマシンでは `LANG=ja_JP.UTF-8` だし, `LC_CTYPE=UTF-8` だった

    Koseis-MBA:kosei% locale
    LANG="ja_JP.UTF-8"
    LC_COLLATE="ja_JP.UTF-8"
    LC_CTYPE="UTF-8"
    LC_MESSAGES="ja_JP.UTF-8"
    LC_MONETARY="ja_JP.UTF-8"
    LC_NUMERIC="ja_JP.UTF-8"
    LC_TIME="ja_JP.UTF-8"
    LC_ALL=

いったん上記をコメントアウトしてみる => だめだった

ssh -v の結果をよく見てみるとローカル側の設定ファイルを読んでいるようだった. むしろ `/etc/ssh/ssh_config` は `~/.ssh/config` のシステムワイド版で, そのホストから別ホストへ ssh する際にロードされるものらしい. よってリモート側の設定を復旧して, ローカルの方を変えてみる

    Koseis-MBA:kosei% sudo vi /etc/ssh_config
    以下をコメントアウト
    #   SendEnv LANG LC_*

OK だった

    cou929:kosei% locale
    LANG=C
    LC_CTYPE="C"
    LC_NUMERIC="C"
    LC_TIME="C"
    LC_COLLATE="C"
    LC_MONETARY="C"
    LC_MESSAGES="C"
    LC_PAPER="C"
    LC_NAME="C"
    LC_ADDRESS="C"
    LC_TELEPHONE="C"
    LC_MEASUREMENT="C"
    LC_IDENTIFICATION="C"
    LC_ALL=

### 疑問

- ja_JP.UTF-8 が localedef 出来なかったのはなぜか
  - ja_JP.utf8 と同じと解釈されているから?
- mac も centos もシステムワイドの ssh_config でローカルのロケールをリモートでも使うようにしていたけど, これが一般的なんだろうか
  - ローカルも英語系のロケールにしとけという話ではあるけど, 普段使いの環境だと日本語が使えたほうがいい場面もあるし
