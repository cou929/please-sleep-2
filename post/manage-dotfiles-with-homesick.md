{"title":"homesick を導入してみたが","date":"2014-03-18T23:04:06+09:00","tags":["blog"]}

[technicalpickles/homesick](https://github.com/technicalpickles/homesick)

ドットファイルをはじめホームディレクトリ以下の設定ファイルの管理。いままでは Dropbox においてリンクを張って使っていた。homesick という gem でその辺の管理をしているひとも多いと聞いて、ちょっと試してみた。

homesick は、要は github 上にある設定ファイルを pull してリンクしているだけ。homesick コマンドを通じて統一的に commit や push は一応可能。

使い方としては、

- github に dotfiles などを管理するリポジトリを作って push しておく
  - home というディレクトリの下にファイルを置く規約
  - home 以下のすべてのファイル・ディレクトリの symlink をホームディレクトリにはってくれる。よって dotfile だけではなくディレクトリなどホームディレクトリ直下に置きたいものならなんでも対応可能
- homesick をインストール

        gem install homesick

- dotfile 管理リポジトリをクローンする

        gem clone cou929/dotfiles

- ホームディレクトリへ symlink をはる

        homesick symlink dotfiles

準備はこれで OK

何かファイルの管理は以下のように行う

- pull

        homesick pull

- 手元の diff を確認

        homesick diff

- 手元の差分をコミット

        homesick commit

- push

        homesick push

リポジトリは `~/.homesick/repos/` に clone されている。何かのミスで手元でコンフリクトした場合などは、最悪ここに入って直接 git コマンドをたたけばよい。

### 感想

このように、この方法はそんなにシンプルではない。dropbox or github + 手製の管理スクリプト (symlink はるだけ) とそれほど手間は変わらない。(もちろんそういう野良管理よりは少し楽だけど、劇的ではない。) 一応今のところ使い続けているが、このような使い方をする分には、そんなに旨味を感じないツールだった。

