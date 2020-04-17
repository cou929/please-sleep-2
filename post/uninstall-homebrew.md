{"title":"homebrew の uninstall","date":"2014-03-02T19:54:13+09:00","tags":["mac"]}

[公式 FAQ](https://github.com/Homebrew/homebrew/wiki/FAQ#wiki-how-do-i-uninstall-homebrew) によると、以下のシェルスクリプトを使えばよいらしい。

[Uninstall Homebrew](https://gist.github.com/mxcl/1173223)

念のためインストール済みの formula の確認

    $ brew list
    ack                     cmake                   jenkins                 mysql-connector-c       pcre                    python3                 sqlite
    autoconf                emacs                   libyaml                 nginx                   phantomjs               rbenv                   tig
    automake                gdbm                    mercurial               nkf                     pkg-config              readline                xz
    casperjs                git                     mysql                   openssl                 pstree                  ruby-build

`uninstall_homebrew.sh` の実行

    $ curl -O https://gist.githubusercontent.com/mxcl/1173223/raw/a833ba44e7be8428d877e58640720ff43c59dbad/uninstall_homebrew.sh
    $ sh -x uninstall_homebrew.sh
    + set -e
    + /usr/bin/which -s git
    + test -d /usr/local/.git
    ++ brew --prefix
    + cd /usr/local
    + git checkout master
    Already on 'master'
    + git ls-files -z
    + pbcopy
    + rm -rf Cellar
    + bin/brew prune
    Pruned 0 dead formula
    Pruned 720 symbolic links and 28 directories from /usr/local
    + pbpaste
    + xargs -0 rm
    + rm -r Library/Homebrew Library/Aliases Library/Formula Library/Contributions
    + test -d Library/LinkedKegs
    + rmdir -p bin Library share/man/man0

これだけで OK。

余談だがこのスクリプトはクリップボードをつかってデータを受け渡ししている部分がある。一時ファイルをつくらないいい方法だけど、状況によっては大量の文字列がクリップボードに入りっぱなしになってしまうので `echo -n | pbcopy` などとしておくといいかも。

