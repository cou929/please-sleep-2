{"title":"mac で python 環境構築 (2014年3月版)","date":"2014-03-18T23:10:18+09:00","tags":["mac"]}

今までは手作業で `easy_install pip` して `pip install virtualenv` してなんだかんだとやっていたのを、他の言語に合わせて pyenv 使ったらいいなじゃないかと思い、Brewfile の整理ついでにそっちに寄せてみていた。で、会社の python の人たちにモダンな python 環境について聞いてみたところ、「pyenv はない」「python はマイナーバージョンまで気にするシチュエーションが少ないからプラットフォームのパッケージ管理ツールで 2.7 系と 3 系の最新を入れておけば十分」「virtualenv でやるのが普通」という意見を賜った。なるほどそうかと思い virtualenv ベースの環境に戻ることにする。とはいえできるだけインストールまわりは Brewfile に寄せていき、手で打つコマンドを減らしたいので、そのへんはアップデート。

### 手順

- python, python3 は brew で入れる
  - Brewfile に以下を書いておいて `brew bundle` で OK

            install python || true
            install python3 || true

  - 2.7 系についてシステム python ではなく brew の python を使う理由は、brew の pythnon には pip がついてくるので、それを手でインストールする手間が減ること。また一応マイナーバージョンも最新の 2.7 系が入るため。
- PATH の `/usr/local/bin` を `/usr/bin` より前に持ってくる。`.zshrc` (`.bashrc`) には以下を書いておく

        export PATH=/usr/local/bin:$PATH

  - screen か tmux を使っている場合、こうしておかないとシステム python が見えてしまうようだ
- virtualenv と virtualenvwrapper をインストールする

        `pip install virtualenv virtualenvwrapper`

- virtualenvwrapper の activate
  - `.zshrc` (`.bashrc`) に以下を書いておく

            source /usr/local/bin/virtualenvwrapper.sh

あとは `mkvirtualenv NAME` などとすればよい。python3 の環境が作りたければ `mkvirtualenv NAME --python /usr/local/bin/python3` とする。

思えば virtualenv を使うと python の処理系を含めて仮想環境化してくれるので、virtualenv を動かす python と virtualenv 環境内の python は違っても構わない。この仕組があるかぎり python バイナリそのものの管理はそこそこ適当でいいし、いろいろなバージョンの python バイナリをインストールする作業はそう多くはないので brew などで十分だ。

あと自分の環境では、はじめに `brew install python` したあと、付属のコマンド (`pip` や `easy_install`) などが動かない事があった。具体的には以下のような `pkg_resources` がインポートできないというエラーが出ていた。

    Traceback (most recent call last):
      File "/usr/local/bin/pip", line 5, in <module>
      from pkg_resources import load_entry_point
    ImportError: No module named `pkg_resources`

いまいち理由はわからないけれど、python をインストールし直す (brew が python をビルドし直す) ことで解決した

    brew uninstall python
    brew install python

