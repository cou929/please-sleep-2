{"title":"Ubuntu 16.04 (Xenial) 以降のタイムゾーン設定方法","date":"2019-01-20T17:11:43+09:00","tags":["nix"]}

debconf を使う場合は、次のようにする。

    sudo ln -fs /usr/share/zoneinfo/Europe/Berlin /etc/localtime
    sudo dpkg-reconfigure -f noninteractive tzdata

### 16.04 以降の推奨方法

これらのチケットによると、

- [Bug \#1554806 “Change of behavior: “dpkg\-reconfigure \-f nonintera\.\.\.” : Bugs : tzdata package : Ubuntu](https://bugs.launchpad.net/ubuntu/+source/tzdata/+bug/1554806)
- [\#813226 \- tzdata config script ignores /etc/timezone on non\-interactive configuration \- Debian Bug report logs](https://bugs.debian.org/cgi-bin/bugreport.cgi?bug=813226#10)

> Thus the preferred approach for changing the timzone is indeed the change of the /etc/localtime symlink, followed by dpkg-reconfigure.

とのことで、ubuntu 16.04 以降 (tzdata 2016a-1 以降) は、`/etc/localtime` に symlink をはり `dpkg-reconfigure` をするのが推奨される方法らしい。よって巷の古い記事にある `/etc/timezone` を変更するような方法では意図した挙動にならない。

### 背景を少し調べたメモ

- `debconf` について
    - Debian (Ubuntu) のパッケージは `.deb` フォーマットだが、`debconf` はパッケージごとに必要な設定を管理する統合的なフレームワーク。
    - ユーザーとインタラクティブに設定値をやり取りしたり、それをデーターベースに保存してり読み込んだりする。
    - [debconf 論](https://tokyodebian-team.pages.debian.net/html2005/debianmeetingresume2005-fuyuse2.html) の内容がわかりやすい
    - ソース
        - [Debconf Maintainers / debconf · GitLab](https://salsa.debian.org/pkg-debconf/debconf)
    - トラッカー
        - [Debian Package Tracker \- debconf](https://tracker.debian.org/pkg/debconf)
- `dpkg-reconfigure`
    - debconf でインストール済みのパッケージの再設定と反映をさせるコマンド
    - [dpkg\-reconfigure · master · Debconf Maintainers / debconf · GitLab](https://salsa.debian.org/pkg-debconf/debconf/blob/master/dpkg-reconfigure) にソースあある
    - [dpkg-query --control-path your_pkg](https://salsa.debian.org/pkg-debconf/debconf/blob/master/dpkg-reconfigure#L174) でそのパッケージの管理スクリプトのパスが返されるが、[dpkg-reconfigure はそれらを再度実行するような処理をしている](https://salsa.debian.org/pkg-debconf/debconf/blob/master/dpkg-reconfigure#L191)
        - 設定は `config`、そのほか `postrm` はパッケージ削除後の処理、`postinit` はインストール後の処理っぽい
        - postrm, config, postinit の順で処理を呼び出している
- `tzdata`
    - Linux はローカルタイムゾーンとして `/etc/localtime` をみにいく。`/etc/localtime` には `/usr/share/zoneinfo/*` への symlink を置く必要がある。
    - `tzdata` は `/usr/share/zoneinfo` を debian のマナーで管理する。debconf でデフォルトを保存したりしている。(ので、`/etc/localtime` への symlink はっただけで `dpkg-reconfigure` をしないと中途半端になる)
    - ソース
        - [GNU Libc Maintainers / tzdata · GitLab](https://salsa.debian.org/glibc-team/tzdata)
    - トラッカー
        - [Debian Package Tracker \- tzdata](https://tracker.debian.org/pkg/tzdata)
    - 参考
        - [What is the purpose of '/etc/localtime' file on Linux Systems? \- Quora](https://www.quora.com/What-is-the-purpose-of-etc-localtime-file-on-Linux-Systems)

### systemd は?

[timedatectl](https://www.freedesktop.org/software/systemd/man/timedatectl.html) という systemd 由来のコマンドもある。

    # タイムゾーン名確認
    $ timedatectl list-timezones | grep Tokyo
    Asia/Tokyo

    # 設定
    $ sudo timedatectl set-timezone Asia/Tokyo

    # 確認
    $ timedatectl
          Local time: ...snip

systemd を使える環境ならばこちらの方法に寄せるべきなのか、debconf を使ったほうがよいのか、どちらが良いのかはよくわからない。
