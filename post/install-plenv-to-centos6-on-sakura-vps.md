{"title":"さくら VPS CentOS6 に plenv を導入する","date":"2013-12-20T21:56:59+09:00","tags":["perl"]}

perl 版 rbenv であるところの plenv を CentoOS6 環境に導入する。

[tokuhirom/plenv](https://github.com/tokuhirom/plenv)

README に書いてあるとおりにするだけでよい。

- clone する

        $ git clone git://github.com/tokuhirom/plenv.git ~/.plenv

- パスをとおす

        $ echo 'export PATH="$HOME/.plenv/bin:$PATH"' >> ~/.bash_profile

- plenv init するよう設定

        $ echo 'eval "$(plenv init -)"' >> ~/.bash_profile

- shell を読み込み直す

        $ exec $SHELL -l

- perl-build を入れる。

        $ git clone git://github.com/tokuhirom/Perl-Build.git ~/.plenv/plugins/perl-build/

- インストールできる perl の一覧

        $ plenv install -l

- perl をインストールする

        $ plenv install 5.18.0

- rehash

        $ plenv rehash

- global のデフォルトに設定

        $ plenv global 5.18.0
        $ plenv global
        5.18.0

- cpanm もいれておく

        $ plenv install-cpanm

