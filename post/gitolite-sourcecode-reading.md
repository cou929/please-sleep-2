{"title":"gitolite ってどういう仕組","date":"2012-02-25T19:07:29+09:00","tags":["nix"]}

### 予想

- インストールスクリプト郡
- gitolite-admin へのフックスクリプト群

で構成されてる

### 結論
肝心なところは全然読めていない感じだけど, 上記の予想はだいたいあってそう.

- gl-setup, gl-install 的なスクリプト郡がファイルの設置とコンフィグを行う
- 数々のフックスクリプトがパーミッションのチェック, リポジトリの作成, ユーザの追加などを行う

実質的な処理 (アクセスコントロールも) が全部 hook で実現されているので, git 専用ユーザを一人作っておけばあとは全部 OK になる. 運用効率がいい.
アドミンの設定ファイルもリポジトリ管理というのが筋がいい.

結局のところ gitolite が解決しているのは.

- ユーザ管理, アクセスコントロール (普通にやるとかなりめんどうだと思う)
- リポジトリ管理
- 設定変更履歴の管理

あたりだろうか. 面倒なユーザ管理が簡単にできて, 各種設定関係も一元化できる.
perl だったのは意外だった.
git サーバを運用した経験, 必要から生まれたという感じがする

### Makefile
最新のタグを tar でアーカイブしている

- Makefile ではこうやって変数をかけるらしい. shell expr で expr の実行結果になる

        branch := $(shell git rev-parse --abbrev-ref HEAD)

- git describe で最新のタグ名
- git archive で tar で固める
- tar -r -f foo.tar bar とすると foo.tar に bar を追加

### src

#### gl-system-install
実行ユーザのホームディレクトリに gitolite 本体をインストールするシェルスクリプト. root の場合は system-wide (/usr/local/bin など) にインストール
インストールとは実際には bin/* などを適当なディレクトリにコピーするだけ

- die は $@ と定型文を出して exit 1 する. echo するたびに >&2 しているのは意味がわからない

        die() { echo >&2; echo "$@" >&2; echo >&2; echo Run "$0 -h" for a detailed usage message. >&2; exit 1; }

- 1 ラインでチェック & usage は "[ "$1" = "-h" ] && usage" こういう書き方をするといいのか
- ディレクトリのバリデーション
  - "||" と "&&" を $? や test と組み合わせると簡潔に書けて良い感じ

         validate_dir() {
            echo $1 | grep '^/' >/dev/null || die "$1 should be an absolute path"
            [ -d $1 ] || mkdir -p $1 || die "$1 does not exist and could not be created"
         }

- 実効ユーザ id によってパスを変える
  - [ test ] && { } でブロックをとれるらしいけど,  test && foo のイディオムは foo がワンラインの時だけにしたほうがいいなじゃないか. foo が数行あるブロックになる場合は普通に if で書いたほうが意味が伝わる
  - perl だと $> で euid とれるらしい
  - set はそのスクリプトの引数 ($@ や $*, $1, $2 ... で参照できるもの) をセットする
  - 複数の値の取得につかえそう


        [ -z "$1" ] && {
          euid=`perl -e 'print $>'`
          if [ "$euid" = "0" ]
          then
            set /usr/local/bin /var/gitolite/conf /var/gitolite/hooks
          else
            set $HOME/bin $HOME/share/gitolite/conf $HOME/share/gitolite/hooks
          fi
          echo "using default values for EUID=$euid:" >&2
          echo "$@" >&2
        }

- だいたい L67 くらいまでで引数のバリデーションや取得.
- ここから, まずは src/ 以下を bin_dir にコピー. セットアップスクリプトの引数を置換する

        cp src/* $buildroot$gl_bin_dir || die "cp src/* to $buildroot$gl_bin_dir failed"
        perl -lpi -e "s(^GL_PACKAGE_CONF=.*)(GL_PACKAGE_CONF=$gl_conf_dir)" $buildroot$gl_bin_dir/gl-setup

- 同様に conf/*, hook/* もコピー. テンプレートファイルを置換する
- bin をインストールした (cp した) 対象ディレクトリが PATH に入っているかを, "which gl-setup 2> /dev/null" して見つかったかどうかでチェックしてる
  - ただし OSX だと which forrbar すると stdout に "foobar not found" と出るので良くないかも
  - env を見るのが妥当か
  - こんなイメージ

        target="/ausr/local/bin"
        in_path=0
        for p in `env | grep PATH | cut -d '=' -f 2 | sed -e 's/:/ /g'`; do
            [ "$p" = "$target" ] && in_path=1
        done
        echo "result: $in_path"

#### gl-setup
admin の名前, その公開鍵をインプットとし, もろもろのインストール作業 (gl-install が行う) と gitolite-admin リポジトリの初期設定までを行うシェルスクリプト

- trap でシグナルハンドラが書ける. trap "処理" SIGNUM. シグナル 0 はプロセスが終了時に自分自身に送る EXIT というシグナルらしい
  - こんな感じで終了時にテンポラリファイルを削除している

        TEMPDIR=`mktemp -d -t tmp.XXXXXXXXXX`
        export TEMPDIR
        trap "/bin/rm -rf $TEMPDIR" 0

- $GITOLITE_HTTP_HOME という変数で分岐させてるけどこれがどこから来たのかよくわからない (L38)
- basename path suffix で path からパスと suffix を抜いた文字列が返される. suffix は知らなかった便利
- L54 から 30 行ほどで rc ファイルのセットアップ
  - comm という diff ツールは知らなかった
- get_rc_val は gl-query-rc を読んで設定値を読み取る
- そのあとは "gl-setup -q" して, 初期の gitolite.conf ファイルを作って pubkey があればそれをコピー, "gl-compile-conf -q", ここまでの設定を gitolite-admin にコミット (その際ユーザ名とかもよしなに), もう一回 "gl-install -q", 最後に sshkeys-lint っていうので authorized_keys のチェックををやってる.
  - こういう, "cat << EOF | ..." といった具合にヒアドキュメント読ませてそのアウトプットで続きをやる, みたいな書き方は新鮮だった. 確かにこうするのが早い
  - ssh issue がたくさん来て単変だったのかなあと伺わせる...

#### gl-install
インターナルコマンド
必要なディレクトリを作ったり, src や doc dir をコピーしたり, 対象の全 repo の hook に hook/common のリンクを貼ったりする
gl-setup が gitolite-admin レベルの初期化, gl-install がその配下の repo に対して, という切り分けっぽい

### hooks

#### common/update
主にアップデートしようとしている対象のパーミッションチェックをしているようだ

#### gitolite-admin/post-update
gitolite-admin にアップデートがかかった際に呼ばれるシェルスクリプト
gitolite-admin は hooks/common にあるフック + これ. この 1 ファイルのみが他の repo と異なる
ディレクトリ名のチェック (src, hooks は使えない), conf ファイルのコンパイル (gl-compile-conf) のあと, hooks/post-update.secondary (デフォルトでは) を実行. 設定されてた場合だけかも.

#### common/update.secondary.sample
フックスクリプトのエントリーポイントのようなシェルスクリプト
update.secondary.d 以下のスクリプトをすべて実行する. 失敗したらログる

- exec >&2 とすると, そのスクリプトの実行結果が全部 stderr にいくらしい. (理解できていない)
