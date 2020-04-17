{"title":"GitHub にある perl のプロジェクトを Jenkins で CI","date":"2014-02-22T16:45:45+09:00","tags":["nix"]}

GitHub にある perl のプロジェクトのテストを jenkins でまわす。

- perl のテストは prove で実行する
- フロントエンドの lint, test のため grunt も導入してあり、そのジョブもつくる
- perl の環境は plenv + carton
- node は nvm で

### GitHub と Jenkins の接続

GitHub に push した際に hook で Jenkins に通知を送り、それをうけとった Jenkins がリポジトリを pull できるようにする。

いろいろ方法があるようだが、今回は [GitHub Plugin](https://wiki.jenkins-ci.org/display/JENKINS/Github+Plugin) を使う方法。

- Jenkins に GitHub Plugin をインストールする
  - [GitHub Plugin - Jenkins - Jenkins Wiki](https://wiki.jenkins-ci.org/display/JENKINS/Github+Plugin)
- Jenkins ユーザの鍵を作成

        $ grep jenkins /etc/passwd                # jenkins ユーザのホームディレクトリをしらべる
        jenkins:x:497:498:Jenkins Continuous Build server:/var/lib/jenkins:/bin/false
        $ cd /var/lib/jenkins/                    # ホームディレクトリへ移動
        $ sudo -u jenkins -H ssh-keygen -t rsa    # 鍵を作成

- GitHub に公開鍵を登録
  - GitHub の Web UI、ユーザー設定の `SSH Keys` から登録する
- 接続確認

        $ sudo -u jenkins ssh -T git@github.com
        Hi cou929! You've successfully authenticated, but GitHub does not provide shell access.

- jenkins ユーザの git の設定 (念のため)

        $ sudo -u jenkins git config --global user.email "jenkins@example.com"
        $ sudo -u jenkins git config --global user.name "jenkins"

- GitHub の hook に jenkins の url を追加
  - GitHub の Web UI から、リポジトリの `settings` -> `Webhooks & Services` -> `Configure services` -> `Jenkins (GitHub plugin)` を選択
  - `Jenkins Hook Url` に `SERVER_HOST/github-webhook/` を入力する
    - ex. `http://jenkins.exampale.com/github-webhook/`

### jenkins ユーザーの環境設定

jenkins ユーザのホームに perl と node の環境を準備する。

perl の環境の設定

    $ sudo -u jenkins -s
    $ git clone git://github.com/tokuhirom/plenv.git ~/.plenv
    $ echo 'export PATH="$HOME/.plenv/bin:$PATH"' >> ~/.bash_profile
    $ echo 'eval "$(plenv init -)"' >> ~/.bash_profile
    $ exec $SHELL -l
    $ git clone git://github.com/tokuhirom/Perl-Build.git ~/.plenv/plugins/perl-build/
    $ plenv install 5.18.2
    $ plenv global 5.18.2
    $ plenv rehash
    $ plenv install-cpanm
    $ cpanm Carton

node の環境の設定

    $ git clone git://github.com/creationix/nvm.git ~/nvm
    $ . ~/nvm/nvm.sh
    $ nvm install v0.10.26
    $ nvm alias default v0.10.26
    $ nvm use default
    $ npm install -g grunt-cli

### Jenkins の設定

perl 側。ビルドの設定では、 plenv のセットアップ、`carton install`、`carton exec prove` を行う。prove のオプションは proverc ファイルに書いておく。テスト結果は `TAP::Formatter::JUnit` で出力して Jenkins が集計できるようにする。手元で prove コマンドをうつときには TAP で出力して欲しいので、`--formatter` オプションは `proverc` ファイルには書かないようにした。

- 新しい job を作成
- GitHub project にリポジトリの URL を設定
- ソースコード管理
  - Git を選択
  - リポジトリ URL を入力
    - `git@github.com:cou929/blahblah.git`
  - Credentials には先ほど作成した鍵を選択
    - `Add` -> `SSH ユーザー名と秘密鍵`
    - 秘密鍵は `Jenkinsマスター上の~/.sshから` を選択
  - Branches to build に対象のブランチを設定
- ビルド・トリガ
  - `Build when a change is pushed to GitHub` を選択
- ビルド

        export PATH="$HOME/.plenv/bin:$PATH"
        eval "$(plenv init -)"
        carton install
        carton exec -- prove --rc=t/proverc --formatter=TAP::Formatter::JUnit -r t > prove-result.xml

- ビルド後の処理
  - `JUnitテスト結果の集計` を選択
  - `prove-result.xml` を入力

grunt 側の設定も基本的に同様で、ビルドの部分だけを次のようにする。

    source $HOME/.nvm/nvm.sh
    npm install
    grunt test

### 参考

- [GitHubとJenkins連動　自動デプロイ　開発環境設定編 at ITエンジニアmegadreamsの開発日記](http://megadreams14.com/?p=27)

<iframe src="http://rcm-fe.amazon-adsystem.com/e/cm?lt1=_blank&bc1=000000&IS2=1&bg1=FFFFFF&fc1=000000&lc1=0000FF&t=pleasesleep-22&o=9&p=8&l=as4&m=amazon&f=ifr&ref=ss_til&asins=4774148911" style="width:120px;height:240px;" scrolling="no" marginwidth="0" marginheight="0" frameborder="0"></iframe>

