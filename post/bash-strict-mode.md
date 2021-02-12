{"title":"シェルスクリプトを書くときにいつもやるやつを調べた","date":"2021-02-12T22:55:00+09:00","tags":["nix"]}

bash のシェルスクリプトを書くときに、いつも脳死で以下をやっている。(同僚が整備してくれたものをコピペしている)

- エディタなり CI で [shellcheck](https://github.com/koalaman/shellcheck) をまわす
- `set -euxo pipefail` と冒頭に書く
    - こんな感じ

```sh
#!/bin/bash
set -euxo pipefail
```

いつまでもコピペではさすがにアレなので、意味を調べたメモ。

## shellcheck

[koalaman/shellcheck: ShellCheck, a static analysis tool for shell scripts](https://github.com/koalaman/shellcheck)

- イケてない書き方に警告を出してくれる
- それぞれの警告にはエラーコード割り振られていてとても便利
    - エラーコードごとに正誤例、解説が書かれているのでわかりやすい
        - [SC1000 の例](https://github.com/koalaman/shellcheck/wiki/SC1000)
- CI もそうだし、[エディタのプラグインも充実](https://github.com/koalaman/shellcheck#in-your-editor) しているのでとりあえず入れておくと良い

## set

`set` という builtin でフラグを立てることによってスクリプトの挙動を変更できる。

設定の仕方によって、エラー時の abort や間違いやすい挙動を防ぐことで、バグに気づかずなんとなく動いてしまう状況を防げる。以下の記事では上手く設定したものを strict mode のようなものだと紹介していた。

[Bash Strict Mode](http://redsymbol.net/articles/unofficial-bash-strict-mode/)

### -e

- コマンドが失敗した時点でスクリプト全体を即座にエラー終了する
- これがないとエラーがあっても以降の処理を継続してしまう
- スクリプト全体の exit code が最後のコマンドのものになってしまい、CI との相性も悪い

```sh
# -e なし
$ cat test.sh
set -e
grep abc foo  # 存在しないファイルを grep しようとしてエラー
echo 1
$ bash test.sh
grep: foo: No such file or directory
1
$ echo $?
0  # exit code が 0 になってしまっている

# -e あり
$ cat test.sh
set -e
grep abc foo
echo 1
$ bash test.sh
grep: foo: No such file or directory
$ echo $?
2
```

### -u

- 初期化していない変数があるとエラーにしてくれる
- 例は前述の記事より引用

```sh
# -u なし
$ cat test.sh
firstName="Aaron"
fullName="$firstname Maxwell"  # `n` を小文字にタイポ
echo "$fullName"
$ bash test.sh
 Maxwell

# -u あり
$ cat test.sh
set -u
firstName="Aaron"
fullName="$firstname Maxwell"
echo "$fullName"
$ bash test.sh
test.sh: line 3: firstname: unbound variable  # エラーでタイポに気づくことができた
```

### -x

- 実行するコマンドを出力してくれる
- 何をしたらどうなったかがログに残る
    - CI に乗せるようなスクリプトなら入れたほうが良いし、ちょっとした処理でもデバッグしやすい

```sh
# -x なし
$ cat test.sh
echo 123
date
$ bash test.sh
123
Fri Feb 12 13:36:43 UTC 2021

# -x あり
$ cat test.sh
set -x
echo 123
date
$ bash test.sh
+ echo 123
123
+ date
Fri Feb 12 13:37:04 UTC 2021
```

### -o pipefail

- パイプの途中でエラーがあれば exit code がそれになる
    - デフォルトではパイプ最後のコマンドの exit code

```sh
# -o pipefail なし
$ set +o pipefail
$ grep foo bar | sort
grep: bar: No such file or directory
$ echo $?
0  # 標準エラーにはエラーメッセージが出たが、exit code は 0

# -o pipefail あり
$ set -o pipefail
$ grep foo bar | sort
grep: bar: No such file or directory
$ echo $?
2
```

## IFS

[Bash Strict Mode](http://redsymbol.net/articles/unofficial-bash-strict-mode/) によると `IFS=$'\n\t'` という設定もおすすめらしい。

- `IFS = Internal Field Separator` で、区切り文字の指定
- デフォルトは `IFS=$' \n\t'` スペース、改行、タブ
    - 例えば `"a b c"` という文字列を for でまわすと `a` `b` `c` の 3 要素に分割される

```sh
$ cat test.sh
items="a b c"
for x in $items; do
    echo "$x"
done
$ bash test.sh
a
b
c
```

- スペース区切りが意図しない挙動になりがちなので、間違えの元らしい
- 次のような引数のパースをするような際には、確かに空白区切りはバクの温床になる

```sh
# IFS デフォルト
$ cat test.sh
for arg in $@; do
    echo "doing something with file: $arg"
done
$ bash test.sh argFoo argBar 'some file.txt'
doing something with file: argFoo
doing something with file: argBar
doing something with file: some
doing something with file: file.txt

# IFS 改行・タブのみ
$ cat test.sh
IFS=$'\n\t'
for arg in $@; do
    echo "doing something with file: $arg"
done
$ bash test.sh argFoo argBar 'some file.txt'
doing something with file: argFoo
doing something with file: argBar
doing something with file: some file.txt
```

## 参考

- [Bash Strict Mode](http://redsymbol.net/articles/unofficial-bash-strict-mode/)
- [Writing Safe Shell Scripts](https://sipb.mit.edu/doc/safe-shell/)
- [Bash Reference Manual](https://www.gnu.org/savannah-checkouts/gnu/bash/manual/bash.html)
    - set builtin のリファレンスは [4.3.1節](https://www.gnu.org/savannah-checkouts/gnu/bash/manual/bash.html#The-Set-Builtin)
- [koalaman/shellcheck: ShellCheck, a static analysis tool for shell scripts](https://github.com/koalaman/shellcheck)

## PR

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B07JGYV4Q8/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/41PzlCY9+sL.jpg" alt="［改訂第3版］シェルスクリプト基本リファレンス ──#!/bin/shで、ここまでできる WEB+DB PRESS plus" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B07JGYV4Q8/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">［改訂第3版］シェルスクリプト基本リファレンス ──#!/bin/shで、ここまでできる WEB+DB PRESS plus</a></div><div class="amazlet-detail">山森 丈範  (著)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B07JGYV4Q8/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
