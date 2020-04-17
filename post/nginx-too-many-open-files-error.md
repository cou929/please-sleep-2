{"title":"nginx の Too many open files エラー","date":"2013-05-29T16:52:47+09:00","tags":["nix"]}

同僚が解決してくれた問題。

Web アプリが 500 を返す頻度が増えてきた。nginx のエラーログを見ると次のような行が多発している。

    2013/00/00 00:00:00 [crit] 23907#0: accept4() failed (24: Too many open files)

話としてはこのへんと同じもの。

- [nginx で Too many open files エラーに対処する - Shin x blog](http://www.1x1.jp/blog/2013/02/nginx_too_many_open_files_error.html)

nginx はひとつのプロセスが複数のリクエストをあつかうモデル。1 プロセスが開くことができるファイルディスクリプタの数が足りなくなると上記のエラーが出るそうだ。

よって見るべきはシステム全体のファイルディスクリプタ数上限と nginx のプロセスごとの上限の 2 箇所だ。

### fs.file-max

システム全体のファイルディスクリプタ数上限は fs.file-max というパラメーターで制御されている。

確認するには、

    $ cat /proc/sys/fs/file-max

一時的に変更するには、

    $ sudo -s
    # echo 320000 > /proc/sys/fs/file-max
    # cat /proc/sys/fs/file-max

再起動しても有効なようにするには、sysctl.conf に書いて sysctl -p する。

    $ sudo vim /etc/sysctl.conf
    $ sudo /sbin/sysctl -p

今回のケースではこのパラメーターの値は問題なかった。よって nginx の設定のほうをみていく。

### `worker_rlimit_nofile`

`worker_rlimit_nofile` は 1 プロセスあたりの最大ファイルディスクリプタ数を設定する項目。`worker_connection` (1 プロセスが処理できるコネクション数の上限) の 3 ~ 4 倍にしておくとよいらしい。

今回はここの値が小さかったので修正した。

    -    worker_connections  4096;
    +    worker_connections  10240;

設定ファイルを書き換え後 nginx をリスタート。

サーバ 1 台を設定変更しエラーが落ち着くことを確認後、全体に適用。

適用後の確認は procfs を見るのが確実らしい。

    $ ps aux | grep nginx | grep worker
    $ cat /proc/XXXX/limits
    
    (Max open files をみる)

ちなみに

[軽量高速Webサーバのnginxで静的コンテンツ配信とキャッシュコントロール \| KRAY Inc](http://kray.jp/blog/nginx/)

によると

> nginx ではワーカープロセス初期化処理にて worker_rlimit_nofile で指定した値をsetrlimit(2) で設定するので、ulimit などで指定された値は上書きされます。（ソフトリミット、ハードリミット両方）

とのことらしいので注意。
