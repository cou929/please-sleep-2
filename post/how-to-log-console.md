{"title":"端末出力 (コンソール) のログを取る方法","date":"2012-02-28T11:14:19+09:00","tags":["nix"]}

## 端末ソフトのログ機能

たいていついている. 例えば iTerm だと Shell -> Log -> Start/Stop.

- たいてい screen を使うとごちゃごちゃになる
- 制御コードとかもそのまま出て見づらい

<div></div>

```
ESC[36mDropboxESC[39m/ESC[16CESC[36mMusicESC[39m/ESC[18CaESC[23CESC[36mgitolite-adminESC[39m/ESC[9CESC[35mprojectsESC[39m@ESC[15CESC[36msrcESC[39m/
Koseis-MBA:kosei%^MESC[35BESC[7m[11:12] 0 zsh  1 zsh  2 zsh  ESC[27mESC[1mESC[37mESC[44m3 zshESC[0mESC[7m...omit...
zsh: correct 'lll' to 'll' [nyae]? y
total 168
drwx------+  4 kosei  staff    136  2 25 15:52 ESC[36mDesktopESC[39m/
drwx------+ 13 kosei  staff    442  7 23  2011 ESC[36mDocumentsESC[39m/
drwx------+ 67 kosei  staff   2278  2 24 22:07 ESC[36mDownloadsESC[39m/
drwx------@ 12 kosei  staff    408  2 23 22:34 ESC[36mDropboxESC[39m/
drwx------@ 48 kosei  staff   1632  2  9 08:20 ESC[36mLibraryESC[39m/
drwxr-xr-x   3 kosei  staff    102 12  9  2010 ESC[36mMailESC[39m/
drwx------+  8 kosei  staff    272  4  8  2011 ESC[36mMoviesESC[39m/
drwx------+  5 kosei  staff    170  1  9  2011 ESC[36mMusicESC[39m/
drwx------+  8 kosei  staff    272  9 25 15:33 ESC[36mPicturesESC[39m/
drwxr-xr-x+  5 kosei  staff    170 12  7  2010 ESC[36mPublicESC[39m/
drwxr-xr-x+  9 kosei  staff    306  7 31  2011 ESC[36mSitesESC[39m/
-rw-r--r--   1 kosei  staff   1362  2 10 19:30 a
drwxr-xr-x   9 kosei  staff    306  4 17  2011 ESC[36mbinESC[39m/
-rw-r--r--   1 kosei  staff   3001 12 22 12:54 contestapplet.conf
-rw-r--r--   1 kosei  staff   3001  2  8 12:57 contestapplet.conf.bak
drwxr-xr-x   5 kosei  staff    170  2 25 10:47 ESC[36mgitolite-adminESC[39m/
lrwxr-xr-x   1 kosei  staff     17  1 10  2011 ESC[35mjunkESC[39m@ -> Dropbox/memo/junk
-rw-r--r--   1 kosei  staff   1760  2 27 11:12 logtest
lrwxr-xr-x   1 kosei  staff     25 12  9  2010 ESC[35mmemoESC[39m@ -> /Users/kosei/Dropbox/memo
drwxr-xr-x   3 kosei  staff    102  1  8 17:27 ESC[36mnode_modulesESC[39m/
lrwxr-xr-x   1 kosei  staff     29 12  8  2010 ESC[35mprojectsESC[39m@ -> /Users/kosei/Dropbox/projects
-rw-r--r--   1 kosei  staff  10866  2 27 11:11 screenlog.0
-rw-r--r--   1 kosei  staff  41237  2 27 11:10 screenlog.2
drwxr-xr-x   5 kosei  staff    170  5  8  2011 ESC[36mshareESC[39m/
drwxr-xr-x  16 kosei  staff    544  2 14 16:46 ESC[36msrcESC[39m/
ESC[K
ESC[K
ESC[K
```

## screen のログ機能

`C-a H` でログ開始, その状態で `C-a H` でログ終了 (デフォルトのキーバインドの場合). ホームディレクトリに `screen.n` (n は screen のウィンドウ番号) というログファイルができる.

- screen のウィンドウ切り替えの影響は受けない
- 制御コードがやはり見づらい

<div></div>

```
ESC[0mESC[23mESC[24mESC[JKoseis-MBA:kosei% ESC[Kid^M
uid=501(kosei) gid=20(staff) groups=20(staff),402(com.apple.sharepoint.group.1),...omit...
ESC[0mESC[23mESC[24mESC[JKoseis-MBA:kosei% ESC[Kdate^M
2012年 2月27日 月曜日 11時19分50秒 JST
ESC[0mESC[23mESC[24mESC[JKoseis-MBA:kosei% ESC[Kll^M
total 288
drwx------+  4 kosei  staff    136  2 25 15:52 ESC[36mDesktopESC[39;49mESC[0m/
drwx------+ 13 kosei  staff    442  7 23  2011 ESC[36mDocumentsESC[39;49mESC[0m/
drwx------+ 67 kosei  staff   2278  2 24 22:07 ESC[36mDownloadsESC[39;49mESC[0m/
drwx------@ 12 kosei  staff    408  2 23 22:34 ESC[36mDropboxESC[39;49mESC[0m/
drwx------@ 48 kosei  staff   1632  2  9 08:20 ESC[36mLibraryESC[39;49mESC[0m/
drwxr-xr-x   3 kosei  staff    102 12  9  2010 ESC[36mMailESC[39;49mESC[0m/
drwx------+  8 kosei  staff    272  4  8  2011 ESC[36mMoviesESC[39;49mESC[0m/
```

## script コマンド

`$ script` でログ開始, exit などで抜けると終了.

- 保存ファイル名などを引数で指定できる
- あとは screen のログと同じ

## col コマンド

line feed などをフィルタしてくれるコマンド. いくつかのエスケープシーケンスはこれでフィルタすればいい感じになりそう

```
$ cat logfile | col -bx
```

## zsh

ちゃんと調べてないけど, zsh だといろいろエスケープシーケンスが入っちゃってるから, bash のほうがエスケープシーケンスが少ない

## 結論

作業をログしたいときは

1. bash にする
2. `script [LOG_FILE_NAME]`
3. `col -xb` でフィルタ

が, 一行もコードを書かなくて良くてよさそう.

コード書く場合は script を拡張して col のようなフィルタを入れる感じか.
