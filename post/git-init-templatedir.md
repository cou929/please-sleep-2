{"title":"git hook のテンプレート","date":"2013-09-06T21:42:42+09:00","tags":["git"]}

git の pre-commit のフックにちょっとした linter や checker をはさむことはあると思う。リポジトリの `.git/hooks/` にスクリプトを入れておけば良いが、よく使うものはテンプレートとしてリポジトリを問わずグローバルに適用したい。

git の設定には `init.template` というものがあり、これを使えばよい。

このように、適当なディレクトリにフックしたいスクリプトを準備し、

    $ mkdir ~/.git_template
    $ mkdir ~/.git_template/hooks
    $ vim ~/.git_template/hooks/pre-commit       # フックを作成
    $ chmod 777 ~/.git_template/hooks/pre-commit

次のようにテンプレートディレクトリのパスを git に教えてあげる

    $ git config --global init.templatedir ~/.git_template/

こうすると `git init` のたびに指定したテンプレートディレクトリ (今回は `~/.git_template/`) の中身がそのリポジトリの `.git` にコピーされる。すでにあるリポジトリに git init した場合は、既存のファイルを上書きすることはなく、安全に新しいテンプレートだけをコピーしてくれる。

>        Running git init in an existing repository is safe. It will not overwrite things that are already there. The primary reason for rerunning git
>        init is to pick up newly added templates (or to move the repository to another place if --separate-git-dir is given).
>
> git-init(1) より

みてわかるように init.templatedir は git init した際のテンプレートファイルのコピーもとを指定するオプションだ。よってもちろん hooks だけでなく、ほかのファイルもテンプレート化できる。
