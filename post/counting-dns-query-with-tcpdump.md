{"title":"tcpdump で DNS クエリ数を確認する","date":"2021-02-17T17:00:00+09:00","tags":["nix"]}

[VPC DNS スロットリングが、Amazon が提供している DNS サーバーへのDNS クエリの失敗の原因となっているかどうかを判断する](https://aws.amazon.com/jp/premiumsupport/knowledge-center/vpc-find-cause-of-failed-dns-queries/) を読んでいて、tcpdump でキャプチャしたパケットをファイルに落とし、それを読み込んで DNS のクエリ数を分間で集計する例が記載されていた。

このページの本論も興味深いが、苦手な tcpdump の使い方は地道に身に着けて行こうと思い調べた内容をメモ。と言っても man を読んだだけだが、`-w` オプションに加えて便利なファイルのローテーションを行う機能などは知らなかったので勉強になった。

## ネットワークインタフェースの一覧を確認する

`ifconfig`... とついやってしまいそうになるが、`ip` コマンドを使ったほうが良い。手元の Ubuntu 20.04 (Vagrant) には ifconfig はデフォルトでは入ってなさそうだった。

[If you’re still using ifconfig, you’re living in the past \| Ubuntu](https://ubuntu.com/blog/if-youre-still-using-ifconfig-youre-living-in-the-past)

```
$ sudo ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host
       valid_lft forever preferred_lft forever
2: enp0s3: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 02:6b:34:d2:18:ce brd ff:ff:ff:ff:ff:ff
    inet 10.0.2.15/24 brd 10.0.2.255 scope global dynamic enp0s3
       valid_lft 50196sec preferred_lft 50196sec
    inet6 fe80::6b:34ff:fed2:18ce/64 scope link
       valid_lft forever preferred_lft forever
3: docker0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN group default
    link/ether 02:42:77:3f:07:8a brd ff:ff:ff:ff:ff:ff
    inet 172.17.0.1/16 brd 172.17.255.255 scope global docker0
       valid_lft forever preferred_lft forever
```

また tcpdump にも `--list-interfaces` というオプションがあるらしく、これは知らなかった。

```
$ sudo tcpdump --list-interfaces
1.enp0s3 [Up, Running]
2.lo [Up, Running, Loopback]
3.any (Pseudo-device that captures on all interfaces) [Up, Running]
4.docker0 [Up]
5.bluetooth-monitor (Bluetooth Linux Monitor) [none]
6.nflog (Linux netfilter log (NFLOG) interface) [none]
7.nfqueue (Linux netfilter queue (NFQUEUE) interface) [none]
```

## パケットキャプチャ

```
sudo tcpdump -i eth0 -s 350 -C 100 -W 20 -w /var/tmp/$(curl http://169.254.169.254/latest/meta-data/instance-id).$(date +%Y-%m-%d:%H:%M:%S).pcap
```

[man](https://www.tcpdump.org/manpages/tcpdump.1.html) より。

- `-i`
    - キャプチャするネットワークインタフェースを指定する
- `-s`
    - パケットの最初から指定したバイト数だけをキャプチャする
    - 今回は 350 byte を指定
- `-C`
    - 指定したファイルサイズごとに分割してキャプチャ結果を保存してくれる
    - `-w` で指定したファイル名の後ろに通し番号がつく
    - 単位は 1,000,000 バイト (約 1MB だが、1,048,576 バイトではない)
- `-W`
    - `-C` と組み合わせて使う
    - キャプチャ結果を最大何ファイルまで保存するか指定する
    - 古いものはローテーションされる
- `-w`
    - キャプチャした raw パケットをファイルに保存する
    - あとから `tcpdump -r` で読み取ることを想定
    - `-U` オプションを付けない限り、出力はバッファされる
    - キャプチャ結果のダンプファイルの拡張子としては `.pcap` が一般的らしい

## キャプチャ結果のカウント

```
tcpdump  -r <file_name.pcap> -nn dst port 53 | awk -F " " '{ print $1 }' | cut -d"." -f1 | uniq -c
```

- `-r`
    - `-w` で出力されたファイルを読み込む
- `-nn`
    - ホスト名、ポート番号を名前解決しない
    - [Tcpdump Examples \- 22 Tactical Commands \| HackerTarget\.com](https://hackertarget.com/tcpdump-examples/) などの Web の記事によると、`-n` はホスト名の名前解決抑制、`-nn` でポート番号も含めた抑制になるらしい
        - man にはそのような記載はなく `-n` だけで host も port も抑制するらしい
            - どういうことだろう、バージョンの違いなど?
    - 大きいデータをキャプチャする際は名前解決の分遅くなるので、このオプションを付けると良いらしい
        - [Tcpdump Examples \- 22 Tactical Commands \| HackerTarget\.com](https://hackertarget.com/tcpdump-examples/) より
    - また、このプションをつけないと tcpdump が行う名前解決のパケットまで乗ってきてしまうので、つけておくと良いらしい
        - [tcpdump の便利なオプション \- Qiita](https://qiita.com/ngyuki/items/969d1efaddb68acb5313#-nn) より
- `dst port 53`
    - 宛先が 53 番ポートのもののみに絞り込む
    - ここのフィルタの記法は [pcap-filter(7)](https://www.tcpdump.org/manpages/pcap-filter.7.html) を見るとよいらしい

## 参考

- [Man page of TCPDUMP](https://www.tcpdump.org/manpages/tcpdump.1.html)
- [Tcpdump Examples \- 22 Tactical Commands \| HackerTarget\.com](https://hackertarget.com/tcpdump-examples/)
- [tcpdump の便利なオプション \- Qiita](https://qiita.com/ngyuki/items/969d1efaddb68acb5313#-nn)

## PR

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B076CWFDB8/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/51yaiO5FlXL.jpg" alt="パケットキャプチャの教科書" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B076CWFDB8/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">パケットキャプチャの教科書</a></div><div class="amazlet-detail">みやた ひろし  (著)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B076CWFDB8/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
