{"title":"nc の `-z` オプションと tcp パケット","date":"2021-06-13T23:50:00+09:00","tags":["nix"]}

[nc\(1\): arbitrary TCP/UDP connections/listens \- Linux man page](https://linux.die.net/man/1/nc)

```
-z' Specifies that nc should just scan for listening daemons, without sending any data to them. It is an error to use this option in conjunction with the -l option.
```

nc コマンドの `-z` オプションは、何もデータを送らずにただ接続先をスキャンするというものだが、実際にどういうパケットが送られているのか軽く見てみた。

`-z` をつけると 3 way handshake 確立後にすぐに FIN を送っていた。

## tcpdump 出力の見方のおさらい

nc ではなく curl などでふつうに HTTP GET を飛ばすとこんな感じの出力が得られる。

```
23:32:36.171067 IP ip-172-31-13-124.ap-northeast-1.compute.internal.33940 > server-13-224-146-2.nrt51.r.cloudfront.net.http: Flags [S], seq 2376043291, win 62727, options [mss 8961,sackOK,TS val 60311200 ecr 0,nop,wscale 7], length 0
23:32:36.180481 IP server-13-224-146-2.nrt51.r.cloudfront.net.http > ip-172-31-13-124.ap-northeast-1.compute.internal.33940: Flags [S.], seq 1087796548, ack 2376043292, win 65535, options [mss 1440,sackOK,TS val 3624432296 ecr 60311200,nop,wscale 9], length 0
23:32:36.180507 IP ip-172-31-13-124.ap-northeast-1.compute.internal.33940 > server-13-224-146-2.nrt51.r.cloudfront.net.http: Flags [.], ack 1, win 491, options [nop,nop,TS val 60311210 ecr 3624432296], length 0
23:32:36.180589 IP ip-172-31-13-124.ap-northeast-1.compute.internal.33940 > server-13-224-146-2.nrt51.r.cloudfront.net.http: Flags [P.], seq 1:87, ack 1, win 491, options [nop,nop,TS val 60311210 ecr 3624432296], length 86: HTTP: GET / HTTP/1.1
23:32:36.189975 IP server-13-224-146-2.nrt51.r.cloudfront.net.http > ip-172-31-13-124.ap-northeast-1.compute.internal.33940: Flags [.], ack 87, win 128, options [nop,nop,TS val 3624432305 ecr 60311210], length 0
23:32:36.190296 IP server-13-224-146-2.nrt51.r.cloudfront.net.http > ip-172-31-13-124.ap-northeast-1.compute.internal.33940: Flags [P.], seq 1:589, ack 87, win 128, options [nop,nop,TS val 3624432306 ecr 60311210], length 588: HTTP: HTTP/1.1 301 Moved Permanently
23:32:36.190318 IP ip-172-31-13-124.ap-northeast-1.compute.internal.33940 > server-13-224-146-2.nrt51.r.cloudfront.net.http: Flags [.], ack 589, win 487, options [nop,nop,TS val 60311219 ecr 3624432306], length 0
23:32:36.190509 IP ip-172-31-13-124.ap-northeast-1.compute.internal.33940 > server-13-224-146-2.nrt51.r.cloudfront.net.http: Flags [F.], seq 87, ack 589, win 487, options [nop,nop,TS val 60311220 ecr 3624432306], length 0
23:32:36.200331 IP server-13-224-146-2.nrt51.r.cloudfront.net.http > ip-172-31-13-124.ap-northeast-1.compute.internal.33940: Flags [F.], seq 589, ack 88, win 128, options [nop,nop,TS val 3624432316 ecr 60311220], length 0
23:32:36.200354 IP ip-172-31-13-124.ap-northeast-1.compute.internal.33940 > server-13-224-146-2.nrt51.r.cloudfront.net.http: Flags [.], ack 590, win 487, options [nop,nop,TS val 60311229 ecr 3624432316], length 0
```

`Flags [S]` などとなっているところが TCP のフラグで、記法は以下。例えば 3 way handshake だと `[S] (SYN)` => `[S.] (SYN-ACK)` => `[.] (ACK)` となる。

```
Tcpflags are some combination of S (SYN), F (FIN), P (PUSH), R (RST), U (URG), W (ECN CWR), E (ECN-Echo) or `.' (ACK), or `none' if no flags are set
```

そのあとの `seq 1:87` などとなっているのがシーケンス番号で、`開始:終了` という記法。tcpdump が見やすさのために 1 からの相対表記にしてくれている。データを送っていない場合 (length が 0 の場合) は `終了` は省略されるようだ。

ref. [Man page of TCPDUMP](https://www.tcpdump.org/manpages/tcpdump.1.html)

## `-z` あり

- 3 way handshake 後、すぐに FIN を送って接続終了のシーケンスに入っている
- いずれも seq 1 (length 0) のままで何もデータを送っていないことも確認できる

<div></div>

```
$ nc -w5 -v -z please-sleep.cou929.nu 80
Connection to please-sleep.cou929.nu 80 port [tcp/http] succeeded!
```

<div></div>

```
$ sudo tcpdump host please-sleep.cou929.nu and port 80

23:21:42.951233 IP ip-172-31-13-124.ap-northeast-1.compute.internal.55538 > server-13-227-49-100.nrt20.r.cloudfront.net.http: Flags [S], seq 3089239005, win 62727, options [mss 8961,sackOK,TS val 3787945529 ecr 0,nop,wscale 7], length 0
23:21:42.953759 IP server-13-227-49-100.nrt20.r.cloudfront.net.http > ip-172-31-13-124.ap-northeast-1.compute.internal.55538: Flags [S.], seq 3769853557, ack 3089239006, win 65535, options [mss 1440,sackOK,TS val 4097033079 ecr 3787945529,nop,wscale 9], length 0
23:21:42.953791 IP ip-172-31-13-124.ap-northeast-1.compute.internal.55538 > server-13-227-49-100.nrt20.r.cloudfront.net.http: Flags [.], ack 1, win 491, options [nop,nop,TS val 3787945532 ecr 4097033079], length 0
23:21:42.953813 IP ip-172-31-13-124.ap-northeast-1.compute.internal.55538 > server-13-227-49-100.nrt20.r.cloudfront.net.http: Flags [F.], seq 1, ack 1, win 491, options [nop,nop,TS val 3787945532 ecr 4097033079], length 0
23:21:42.956462 IP server-13-227-49-100.nrt20.r.cloudfront.net.http > ip-172-31-13-124.ap-northeast-1.compute.internal.55538: Flags [F.], seq 1, ack 2, win 128, options [nop,nop,TS val 4097033082 ecr 3787945532], length 0
23:21:42.956484 IP ip-172-31-13-124.ap-northeast-1.compute.internal.55538 > server-13-227-49-100.nrt20.r.cloudfront.net.http: Flags [.], ack 2, win 491, options [nop,nop,TS val 3787945535 ecr 4097033082], length 0
```

## `-z` なし

- 接続はするが `-z` が無いのでクライアントから閉じようとはしない
- 今回は `-w 5` と 5 秒でタイムアウトするようにした
    - tcpdump の出力では、5 秒後に FIN を送っている様子がわかる

<div></div>

```
$ nc -w5 -v please-sleep.cou929.nu 80
Connection to please-sleep.cou929.nu 80 port [tcp/http] succeeded!
```

<div></div>

```
$ sudo tcpdump host please-sleep.cou929.nu and port 80

23:23:13.903640 IP ip-172-31-13-124.ap-northeast-1.compute.internal.57814 > server-13-224-146-2.nrt51.r.cloudfront.net.http: Flags [S], seq 3149699449, win 62727, options [mss 8961,sackOK,TS val 59748933 ecr 0,nop,wscale 7], length 0
23:23:13.911947 IP server-13-224-146-2.nrt51.r.cloudfront.net.http > ip-172-31-13-124.ap-northeast-1.compute.internal.57814: Flags [S.], seq 2782607019, ack 3149699450, win 65535, options [mss 1440,sackOK,TS val 2465407672 ecr 59748933,nop,wscale 9], length 0
23:23:13.911972 IP ip-172-31-13-124.ap-northeast-1.compute.internal.57814 > server-13-224-146-2.nrt51.r.cloudfront.net.http: Flags [.], ack 1, win 491, options [nop,nop,TS val 59748941 ecr 2465407672], length 0
23:23:18.912801 IP ip-172-31-13-124.ap-northeast-1.compute.internal.57814 > server-13-224-146-2.nrt51.r.cloudfront.net.http: Flags [F.], seq 1, ack 1, win 491, options [nop,nop,TS val 59753942 ecr 2465407672], length 0
23:23:18.920877 IP server-13-224-146-2.nrt51.r.cloudfront.net.http > ip-172-31-13-124.ap-northeast-1.compute.internal.57814: Flags [.], ack 2, win 128, options [nop,nop,TS val 2465412681 ecr 59753942], length 0
23:23:18.922630 IP server-13-224-146-2.nrt51.r.cloudfront.net.http > ip-172-31-13-124.ap-northeast-1.compute.internal.57814: Flags [F.], seq 1, ack 2, win 128, options [nop,nop,TS val 2465412683 ecr 59753942], length 0
23:23:18.922657 IP ip-172-31-13-124.ap-northeast-1.compute.internal.57814 > server-13-224-146-2.nrt51.r.cloudfront.net.http: Flags [.], ack 2, win 491, options [nop,nop,TS val 59753952 ecr 2465412683], length 0
```

## 参考

- [tcpdump の見方を勉強 \- Please Sleep](https://please-sleep.cou929.nu/tcpdump-study-pt1.html)

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4894712059/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://images-na.ssl-images-amazon.com/images/I/213B9PVJD1L._BO1,204,203,200_.jpg" alt="UNIXネットワークプログラミング〈Vol.1〉ネットワークAPI:ソケットとXTI" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4894712059/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">UNIXネットワークプログラミング〈Vol.1〉ネットワークAPI:ソケットとXTI</a></div><div class="amazlet-detail">W.リチャード スティーヴンス (著), W.Richard Stevens (原著), 篠田 陽一 (翻訳)<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4894712059/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
