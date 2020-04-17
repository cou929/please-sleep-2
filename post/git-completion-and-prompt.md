{"title":"ブランチ名の補完とプロンプトにブランチ名表示","date":"2013-09-07T11:50:33+09:00","tags":["git"]}

git のコマンドやらブランチ名やらの補完には git-completion.bash だと思っていたが、最近は git-completion.zsh というものを使うらしい。brew の場合 `$(brew --prefix)/share/zsh/site-functions` に適切にシンボリックリンクを張って設置してくれる。git-completion.zsh は _git という名前にしているようだ。

    % ls $(brew --prefix)/share/zsh/site-functions
    _ack@                _git@                git-completion.bash@

`~/.zshrc` にはいかのように書いておく。

    if which brew > /dev/null; then
        fpath=($(brew --prefix)/share/zsh/site-functions $fpath)
    else
        fpath=(~/.zsh/completion $fpath)
    fi
    
    autoload -U compinit
    compinit -u

mac でない場合は `~/.zsh/completion` に同じファイル名で設置することにしている。どちらのファイルも git 本体リポジトリの contrib 以下にあるので、それをもってくればよい。

ところで、普段は `alias g='git'` というエイリアスを張っている。このままでは g で操作している時には補完が効かない。その場合は `~/.zshrc` に以下を追加すれば良い。

    setopt no_complete_aliases

名前が機能を表していないように思えるが、これで意図した動作になる。

ブランチ名をプロンプトに表示するには、次のようにする。この書き方だと git だけでなく hg など vcs を問わず表示できるらしい。

    autoload -Uz vcs_info
    zstyle ':vcs_info:*' formats ' (%b)'
    zstyle ':vcs_info:*' actionformats ' (%b|%a)'
    precmd () {
        psvar=()
        LANG=en_US.UTF-8 vcs_info
        [[ -n "$vcs_info_msg_0_" ]] && psvar[1]="$vcs_info_msg_0_"
    }
    
    PROMPT="%1(v|%F{green}%1v%f|)"

サンプルなので `PROMPT` にはブランチ情報だけを出す例になっている。`%1(v|%F{green}%1v%f|)` を任意の場所に入れれば良い。

以上、詳細はまったく理解してないけどこれでうまくいった。
