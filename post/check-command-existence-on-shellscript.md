{"title":"シェルスクリプトでコマンドの有無を確かめる","date":"2012-12-09T18:52:54+09:00","tags":["nix"]}

### `which` を使う

最初に思いついた方法.

    if [ `which SOME_COMMAND` ]; then
        echo 'found'
    fi

    # 一行で
    which SOME_COMMAND > /dev/null 2>&1 && echo 'found'

### `type` を使う

[OSに付属するシェルスクリプトを読んで技術を盗む（1/2） － ＠IT](http://www.atmarkit.co.jp/flinux/rensai/smart_shell/03/01.html) で紹介されていた方法. コマンドは違うが考え方は上記と同じ.

    if type logger > /dev/null 2>&1; then
            LOGGER="logger -s -p user.notice -t dhclient"
    else
            LOGGER=echo
    fi

### `command -v` を使う

rvm がこういうふうにやっていた. なぜ `builtin command` というふうにわざわざやっているのかがよくわからない. command っていうコマンドがビルトイン以外にもあってかぶるケースがあるのかな.

    __rvm_sha256_for()
    {
      if builtin command -v sha256sum > /dev/null ; then
        echo "$1" | sha256sum | awk '{print $1}'
      elif builtin command -v sha256 > /dev/null ; then
        echo "$1" | sha256 | awk '{print $1}'
      elif builtin command -v shasum > /dev/null ; then
        echo "$1" | shasum -a256 | awk '{print $1}'
      else
        rvm_error "Neither sha256sum nor shasum found in the PATH"
        return 1
      fi
    
      return 0
    }

[rvm/scripts/functions/rvmrc at master · wayneeseguin/rvm · GitHub](https://github.com/wayneeseguin/rvm/blob/master/scripts/functions/rvmrc)
