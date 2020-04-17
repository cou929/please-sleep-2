{"title":"bash fork retry エラー","date":"2014-06-30T22:48:25+09:00","tags":["nix"]}

`sudo -u user ls` などのようにあるユーザーでコマンドを実行しようとすると、

    bash : fork : retry : リソースが一時的に利用できません

というエラーが出た。

調べてみるとユーザーごとのプロセス数制限に引っかかっているらしい。確かに、不要なプロセスを kill することで問題は解決した。

ちなみにこの閾値は `/etc/security/limits.conf` の `nprocs` という変数で定義されているらしい。

現在値は `ulimit` で確認できる

    $ ulimit -a | grep proc
    max user processes              (-u) 1024

### 参考

- [[SE@HUMAN]$ reboot：So-netブログ](http://silvervine-ninelives.blog.so-net.ne.jp/archive/201009-1)
- [LinuxのPAM認証の設定入門](http://www.geocities.jp/sugachan1973/doc/funto47.html)
- [/etc/security/limits.confに関するメモ - OpenGroove](http://open-groove.net/linux/memo-etcsecuritylimits-conf/)
