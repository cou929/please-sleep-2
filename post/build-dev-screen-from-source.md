{"title":"開発版 screen の導入","date":"2012-11-28T22:21:08+09:00","tags":["nix"]}

縦分割がないと業務がつらいので重い腰を上げて開発版を導入. tmux に移るか迷ったんですが今回は screen 継続にしました. 環境は CentOS です.

### インストール

ソースから入れるしかありません.

1. ソースを持ってくる
  - `git clone git://git.savannah.gnu.org/screen.git`
2. ビルドする
  - 今回はホームディレクトリ以下に置いておきたい (/usr などには入れたくない) ので configure 時にそのようにします

            $ cd screen/src
            $ ./autogen.sh
            $ ./configure --prefix=$HOME
            $ make

  - 今回は ncurses-devel が無くて configure 時に怒られたので別途インストールしました.
    - こういうエラーが出た: `configure: error: !!! no tgetent - no screen`
    - `sudo yum install ncurses-devel` でインストール
  - バイナリが screen/src/screen にできているので好きな場所に置いてパスを通せばOK.
    - 今回は `~/bin` に置きました

            $ screen -v
            Screen version 4.00.03 (FAU) 23-Oct-06
            $ bin/screen -v
            Screen version 4.01.00devel (GNUcbaa666) 2-May-06

### 縦分割をやってみる

以下はエスケープキーが `a` の場合

- `C-a |` で縦分割
- `C-a S` で横分割
- `C-a C-i` で画面間を時計回りに移動
- `C-a X` で分割した画面を閉じる
- `C-a Q` で 1 画面に戻す (emacs の `C-x 1` のような動き)

### 感想

キーバインドを覚えるのが面倒なので emacs ふうにカスタマイズするか検討中. 画面がたくさんあるとややこしくて使いこなせないのでいざというとき (複数ホストで tail やら top して見比べたいなど) だけ分割するのが自分にはいいかもしれない. がこれも慣れの問題なのかもしれない. グローバルな画面を分割するんじゃなくて, ある画面の中を分割するようなことってできないのかな. ともかく横分割しかできない頃は不便で仕方なかったのでよかった.

ほかにもいろいろできるらしいので徐々にカスタマイズしていこう.

- [これからの「GNU Screen」の話をしよう - Keep It Simple, Stupid](http://yskwkzhr.blogspot.jp/2011/01/gnu-screen.html#layout)
- [screenを縦に横に分割しまくろう - テックノート＠ama-ch](http://d.hatena.ne.jp/ama-ch/20090129/1233211681)
