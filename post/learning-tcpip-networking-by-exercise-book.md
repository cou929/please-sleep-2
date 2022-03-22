{"title":"Linuxで動かしながら学ぶTCP/IPネットワーク入門","date":"2022-03-22T21:30:00+09:00","tags":["book"]}

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B085BG8CH5/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/51kU2EFP5UL.jpg" alt="Linuxで動かしながら学ぶTCP/IPネットワーク入門" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B085BG8CH5/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Linuxで動かしながら学ぶTCP/IPネットワーク入門</a></div><div class="amazlet-detail">もみじあめ  (著)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B085BG8CH5/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

[momijiame/linux\-tcpip\-book: 「Linuxで動かしながら学ぶTCP/IPネットワーク入門」のサポートページです。](https://github.com/momijiame/linux-tcpip-book)

手を動かしながら L2 ~ L4 を中心にネットワーク周りの基礎を学べる本 (一応 L7 の章もある)。network namespaces を活用して比較的手軽に実験環境を用意することができ、いい感じ。コマンドを手打ちして、基礎から一歩ずつ理解を固めながら知識を積んでいけるので気持ちがよい。今までその場しのぎで断片的だった知識を体系的に捉えることができるようになってよかった。

仕事でネットワーク周りで不可解な事象があり、ただ個人的に L4 より下が苦手だったので調査に時間がかかってしまった。基本的なところからしっかり知識を整理したいなと思って読んでみたが、このニーズには当たりの本だったと思う。

以下読書メモ。

## 2. TCP/IP とは

- nic に ip が割り振られる

<div></div>

```
ubuntu@learning-sapsucker:~$ ip address show
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host
       valid_lft forever preferred_lft forever
2: enp0s1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 52:54:00:8f:d5:e5 brd ff:ff:ff:ff:ff:ff
    inet 192.168.64.32/24 brd 192.168.64.255 scope global dynamic enp0s1
       valid_lft 48280sec preferred_lft 48280sec
    inet6 fd48:dd89:70dd:e78e:5054:ff:fe8f:d5e5/64 scope global dynamic mngtmpaddr noprefixroute
       valid_lft 2591984sec preferred_lft 604784sec
    inet6 fe80::5054:ff:fe8f:d5e5/64 scope link
       valid_lft forever preferred_lft forever
```

- traceroute
    - ttl が切れたときに中間装置が自信の ip を返すことは MUST ではない
    - 返さなかった場合に `* * *` になる

<div></div>

```
ubuntu@learning-sapsucker:~$ traceroute -n 8.8.8.8
traceroute to 8.8.8.8 (8.8.8.8), 30 hops max, 60 byte packets
 1  192.168.64.1  0.688 ms  0.533 ms  0.490 ms
 2  192.168.10.1  5.001 ms  4.990 ms  4.979 ms
 3  14.0.9.78  10.663 ms  10.733 ms  10.616 ms
 4  14.0.9.77  10.586 ms  10.547 ms  10.588 ms
 5  72.14.222.189  10.355 ms  10.187 ms  10.112 ms
 6  * * *
 7  8.8.8.8  9.131 ms  7.057 ms  6.969 ms
```

- tcpdump
    - `-i` インタフェースの指定
    - `-n` 名前解決しない

<div></div>

```
ubuntu@learning-sapsucker:~$ sudo tcpdump -n -i any icmp
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
listening on any, link-type LINUX_SLL (Linux cooked v1), capture size 262144 bytes
22:22:16.438812 IP 192.168.64.32 > 8.8.8.8: ICMP echo request, id 2, seq 1, length 64
22:22:16.447594 IP 8.8.8.8 > 192.168.64.32: ICMP echo reply, id 2, seq 1, length 64
22:22:17.514760 IP 192.168.64.32 > 8.8.8.8: ICMP echo request, id 2, seq 2, length 64
22:22:17.529402 IP 8.8.8.8 > 192.168.64.32: ICMP echo reply, id 2, seq 2, length 64
22:22:18.526196 IP 192.168.64.32 > 8.8.8.8: ICMP echo request, id 2, seq 3, length 64
22:22:18.534793 IP 8.8.8.8 > 192.168.64.32: ICMP echo reply, id 2, seq 3, length 64
```

- routing table
    - matcher nexthop という構造

<div></div>

```
ubuntu@learning-sapsucker:~$ ip route show
default via 192.168.64.1 dev enp0s1 proto dhcp src 192.168.64.32 metric 100
192.168.64.0/24 dev enp0s1 proto kernel scope link src 192.168.64.32
192.168.64.1 dev enp0s1 proto dhcp scope link src 192.168.64.32 metric 100
```

## 3. Network Namespace

- ip サブコマンドの address と link は何が違うんだろう
    - たぶん `link` = L2、`address` = L3 っぽい
    - `link show` だと mac address まで出る。`address show` だと ip まで出る

```sh
# インタフェース作成直後はどちらも同じ表示。mac アドレスは既に割り振られているが、ip アドレスはまだない
ubuntu@learning-sapsucker:~$ sudo ip link add ns1-veth0 type veth peer name gw-veth0
ubuntu@learning-sapsucker:~$ sudo ip link set ns1-veth0 netns ns1
ubuntu@learning-sapsucker:~$ sudo ip link set gw-veth0 netns router

ubuntu@learning-sapsucker:~$ sudo ip netns exec ns1 ip link show
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
6: ns1-veth0@if5: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/ether 26:12:a1:69:0b:4c brd ff:ff:ff:ff:ff:ff link-netns router
ubuntu@learning-sapsucker:~$ sudo ip netns exec ns1 ip address show
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
6: ns1-veth0@if5: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN group default qlen 1000
    link/ether 26:12:a1:69:0b:4c brd ff:ff:ff:ff:ff:ff link-netns router

# インタフェースを UP にすると、v6 の ip は勝手に割り振られる?
ubuntu@learning-sapsucker:~$ sudo ip netns exec ns1 ip link set ns1-veth0 up

ubuntu@learning-sapsucker:~$ sudo ip netns exec ns1 ip address show
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
6: ns1-veth0@if5: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether 26:12:a1:69:0b:4c brd ff:ff:ff:ff:ff:ff link-netns router
    inet6 fe80::2412:a1ff:fe69:b4c/64 scope link
       valid_lft forever preferred_lft forever

# v4 の ip を割り当てるとそれも表示される
ubuntu@learning-sapsucker:~$ sudo ip netns exec ns1 ip address add 192.168.8.1/24 dev ns1-veth0

ubuntu@learning-sapsucker:~$ sudo ip netns exec ns1 ip address show
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
6: ns1-veth0@if5: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether 26:12:a1:69:0b:4c brd ff:ff:ff:ff:ff:ff link-netns router
    inet 192.168.8.1/24 scope global ns1-veth0
       valid_lft forever preferred_lft forever
    inet6 fe80::2412:a1ff:fe69:b4c/64 scope link
       valid_lft forever preferred_lft forever
ubuntu@learning-sapsucker:~$ sudo ip netns exec ns1 ip link show
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
6: ns1-veth0@if5: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP mode DEFAULT group default qlen 1000
    link/ether 26:12:a1:69:0b:4c brd ff:ff:ff:ff:ff:ff link-netns router
```


- 同じネットワークセグメントのホスト同士はルータを介さす直接通信できる
    - 逆を言うとルータはセグメントを越える通信をするための機器
- CIDR とサブネットマスク
	- 本質的には同じ物
	- CIDR はネットワークホスト部の区切りを指定
		- 左から数えてその範囲までがネットワーク部
	- サブネットマスクは and でネットワークアドレスを取り出せるように記載
	- `/24` と `255.255.255.0` は同じ
	- 同じネットワークアドレスは同じセグメント
- ルーターの ip アドレスの慣習
    - そのセグメントの最大値にすることが多い
    - ただしホストアドレスの全ビットが `1` のものはブロードキャストアドレスとして、`0` のものはネットワークアドレスの区別として予約されているので、それ以外の範囲になる
    - i.e.,
        - 192.168.1.0/24 のセグメントの場合、ホストアドレスは `1` から `254` で、255 はブロードキャスト用
- ping の `-I` はインタフェース指定
- ping のエラーメッセージ

```sh
# ルーターに設定ミス (ip_forward を有効化していない) があり到達できないケース
ubuntu@learning-sapsucker:~$ sudo ip netns exec ns1 ping -c 3 192.168.99.1 -I 192.168.8.1
PING 192.168.99.1 (192.168.99.1) from 192.168.8.1 : 56(84) bytes of data.
ping: sendmsg: Network is unreachable
ping: sendmsg: Network is unreachable
ping: sendmsg: Network is unreachable

--- 192.168.99.1 ping statistics ---
3 packets transmitted, 0 received, 100% packet loss, time 2105ms

# 手元のインタフェースが DOWN のケース
ubuntu@learning-sapsucker:~$ ping -c 3 8.8.8.8 -I tmp-veth0
ping: connect: Network is unreachable

# 手元のインタフェースは UP state だが ip を割り当てていないケース
ubuntu@learning-sapsucker:~$ ping -c 3 8.8.8.8 -I tmp-veth0
ping: Warning: source address might be selected on device other than: tmp-veth0
PING 8.8.8.8 (8.8.8.8) from 192.168.64.32 tmp-veth0: 56(84) bytes of data.

--- 8.8.8.8 ping statistics ---
3 packets transmitted, 0 received, 100% packet loss, time 2127ms

# ルーター自体は起動しているがルーティング先がわからず失敗した場合 (routing table に転送先がない場合)
# ルーターの ip から返事が返ってきている
ubuntu@learning-sapsucker:~$ sudo ip netns exec ns1 ping -c 3 198.51.100.1 -I 192.0.2.1
PING 198.51.100.1 (198.51.100.1) from 192.0.2.1 : 56(84) bytes of data.
From 192.0.2.254 icmp_seq=1 Destination Net Unreachable
From 192.0.2.254 icmp_seq=2 Destination Net Unreachable
From 192.0.2.254 icmp_seq=3 Destination Net Unreachable

--- 198.51.100.1 ping statistics ---
3 packets transmitted, 0 received, +3 errors, 100% packet loss, time 2107ms
```

- `-I` の指定方法の違い (ip or name) でエラーメッセージが違う?
    - 名前で指定すると `ping: sendmsg: Network is unreachable` が出ない

```sh
ubuntu@learning-sapsucker:~$ sudo ip netns exec ns1 ping -c 3 192.168.99.1 -I ns1-veth0
PING 192.168.99.1 (192.168.99.1) from 192.168.8.1 ns1-veth0: 56(84) bytes of data.

--- 192.168.99.1 ping statistics ---
3 packets transmitted, 0 received, 100% packet loss, time 2096ms

ubuntu@learning-sapsucker:~$ sudo ip netns exec ns1 ping -c 3 192.168.99.1 -I 192.168.8.1
PING 192.168.99.1 (192.168.99.1) from 192.168.8.1 : 56(84) bytes of data.
ping: sendmsg: Network is unreachable
ping: sendmsg: Network is unreachable
ping: sendmsg: Network is unreachable

--- 192.168.99.1 ping statistics ---
3 packets transmitted, 0 received, 100% packet loss, time 2105ms
```

- `net.ipv4.ip_forward` でカーネルレベルでルーターとして動作するようになる
    - 名前のとおりだけど ip forwarding という機能で、ある nic に届いたパケットを別の nic に転送する
    - routing table の設定に従って転送する

```sh
ubuntu@learning-sapsucker:~$ sudo ip netns exec router sysctl net.ipv4.ip_forward
net.ipv4.ip_forward = 0
ubuntu@learning-sapsucker:~$ sudo ip netns exec router sysctl net.ipv4.ip_forward=1
net.ipv4.ip_forward = 1
ubuntu@learning-sapsucker:~$ sudo ip netns exec router sysctl net.ipv4.ip_forward
net.ipv4.ip_forward = 1
```

- routing table の設定
    - 今回は学習のため手動で設定したがこれは静的ルーティング
    - 実用はそれでは当然まわらないので動的ルーティングが行われている
        - BGP, OSPF というプロトコルがある

## 4. Ethernet

- ip link set で mac アドレスを設定できる

```sh
ubuntu@learning-sapsucker:~$ sudo ip netns exec ns1 ip link set dev ns1-veth0 address 00:00:5E:00:53:01
ubuntu@learning-sapsucker:~$ sudo ip netns exec ns2 ip link set dev ns2-veth0 address 00:00:5E:00:53:02
```

- tcpdump `-el` オプションで ethernet の情報も表示できる

```sh
# -el をつけると ethernet ヘッダ情報も表示される
ubuntu@learning-sapsucker:~$ sudo ip netns exec ns1 tcpdump -tnel -i ns1-veth0 icmp
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
listening on ns1-veth0, link-type EN10MB (Ethernet), capture size 262144 bytes
00:00:5e:00:53:01 > 00:00:5e:00:53:02, ethertype IPv4 (0x0800), length 98: 192.0.2.1 > 192.0.2.2: ICMP echo request, id 31082, seq 1, length 64
00:00:5e:00:53:02 > 00:00:5e:00:53:01, ethertype IPv4 (0x0800), length 98: 192.0.2.2 > 192.0.2.1: ICMP echo reply, id 31082, seq 1, length 64
00:00:5e:00:53:01 > 00:00:5e:00:53:02, ethertype IPv4 (0x0800), length 98: 192.0.2.1 > 192.0.2.2: ICMP echo request, id 31082, seq 2, length 64
00:00:5e:00:53:02 > 00:00:5e:00:53:01, ethertype IPv4 (0x0800), length 98: 192.0.2.2 > 192.0.2.1: ICMP echo reply, id 31082, seq 2, length 64
00:00:5e:00:53:01 > 00:00:5e:00:53:02, ethertype IPv4 (0x0800), length 98: 192.0.2.1 > 192.0.2.2: ICMP echo request, id 31082, seq 3, length 64
00:00:5e:00:53:02 > 00:00:5e:00:53:01, ethertype IPv4 (0x0800), length 98: 192.0.2.2 > 192.0.2.1: ICMP echo reply, id 31082, seq 3, length 64
^C
6 packets captured
6 packets received by filter
0 packets dropped by kernel

# -el をつけなかった場合
ubuntu@learning-sapsucker:~$ sudo ip netns exec ns1 tcpdump -tnl -i ns1-veth0 icmp
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
listening on ns1-veth0, link-type EN10MB (Ethernet), capture size 262144 bytes
IP 192.0.2.1 > 192.0.2.2: ICMP echo request, id 31089, seq 1, length 64
IP 192.0.2.2 > 192.0.2.1: ICMP echo reply, id 31089, seq 1, length 64
IP 192.0.2.1 > 192.0.2.2: ICMP echo request, id 31089, seq 2, length 64
IP 192.0.2.2 > 192.0.2.1: ICMP echo reply, id 31089, seq 2, length 64
IP 192.0.2.1 > 192.0.2.2: ICMP echo request, id 31089, seq 3, length 64
IP 192.0.2.2 > 192.0.2.1: ICMP echo reply, id 31089, seq 3, length 64
```

- `ip neigh`
    - mac アドレスのキャッシュの操作
    - たぶん `neigh = neighber` か?

```sh
ubuntu@learning-sapsucker:~$ sudo ip netns exec ns1 ip neigh
192.0.2.2 dev ns1-veth0 lladdr 00:00:5e:00:53:02 STALE
ubuntu@learning-sapsucker:~$ sudo ip netns exec ns1 ip neigh flush all
```

- tcpdump で arp パケットのキャプチャ
    - なお arp は ipv4 で使われるプロトコルで v6 の場合は ICMPv6 の Neighbor Discovery

```sh
ubuntu@learning-sapsucker:~$ sudo ip netns exec ns1 tcpdump -nel -i ns1-veth0 icmp or arp
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
listening on ns1-veth0, link-type EN10MB (Ethernet), capture size 262144 bytes
21:50:54.721493 00:00:5e:00:53:01 > ff:ff:ff:ff:ff:ff, ethertype ARP (0x0806), length 42: Request who-has 192.0.2.2 tell 192.0.2.1, length 28
21:50:54.721554 00:00:5e:00:53:02 > 00:00:5e:00:53:01, ethertype ARP (0x0806), length 42: Reply 192.0.2.2 is-at 00:00:5e:00:53:02, length 28
```

- スイッチ、ブリッジ
    - 分配のイーサネット版
    - ブリッジのポートに繋がっている機器の mac アドレスを記憶しておき、そこへの通信ができるようにしている
        - もっと単純な機器としてリピータハブというものもあり、それは単に全パケットを全ポートにコピーして送信する
    - L1の情報でデータを転送するのがリピータハブ、L2 は スイッチングハブ、ブリッジ、L3 は ルーター
- ブリッジはネットワークインタフェース。`type bridge`

```sh
ubuntu@learning-sapsucker:~$ sudo ip netns exec bridge ip link add dev br0 type bridge
ubuntu@learning-sapsucker:~$ sudo ip netns exec bridge ip link set dev br0 up
ubuntu@learning-sapsucker:~$ sudo ip netns exec bridge ip link set ns1-br0 master br0
ubuntu@learning-sapsucker:~$ sudo ip netns exec bridge ip link set ns2-br0 master br0
ubuntu@learning-sapsucker:~$ sudo ip netns exec bridge ip link set ns3-br0 master br0
```

- mac アドレステーブルの確認

```sh
ubuntu@learning-sapsucker:~$ sudo ip netns exec bridge bridge fdb show br br0 | grep -i 00:00:5e
00:00:5e:00:53:01 dev ns1-br0 master br0
00:00:5e:00:53:02 dev ns2-br0 master br0
00:00:5e:00:53:03 dev ns3-br0 master br0
```

## 5. Transport, 6. Application

- tcpdump の `Flags [S]` はコントロールビット
- dhcp はネットワークインタフェースに IP 設定、デフォルトルートをルーティングテーブル追加、ネームサーバ指定などを行う
- dnsmasq などを使って dhcp サーバを検証できる

```sh
# server
ubuntu@learning-sapsucker:~$ sudo ip netns exec server dnsmasq --dhcp-range=192.0.2.100,192.0.2.200,255.255.255.0 --interface=s-veth0 --port 0 --no-resolv --no-daemon
dnsmasq: started, version 2.80 DNS disabled
dnsmasq: compile time options: IPv6 GNU-getopt DBus i18n IDN DHCP DHCPv6 no-Lua TFTP conntrack ipset auth nettlehash DNSSEC loop-detect inotify dumpfile
dnsmasq-dhcp: DHCP, IP range 192.0.2.100 -- 192.0.2.200, lease time 1h
dnsmasq-dhcp: DHCPDISCOVER(s-veth0) 7e:6b:a9:f5:61:e3
dnsmasq-dhcp: DHCPOFFER(s-veth0) 192.0.2.153 7e:6b:a9:f5:61:e3
dnsmasq-dhcp: DHCPREQUEST(s-veth0) 192.0.2.153 7e:6b:a9:f5:61:e3
dnsmasq-dhcp: DHCPACK(s-veth0) 192.0.2.153 7e:6b:a9:f5:61:e3 learning-sapsucker

# client
ubuntu@learning-sapsucker:~$ sudo ip netns exec client dhclient -d c-veth0
Internet Systems Consortium DHCP Client 4.4.1
Copyright 2004-2018 Internet Systems Consortium.
All rights reserved.
For info, please visit https://www.isc.org/software/dhcp/

Listening on LPF/c-veth0/7e:6b:a9:f5:61:e3
Sending on   LPF/c-veth0/7e:6b:a9:f5:61:e3
Sending on   Socket/fallback
DHCPDISCOVER on c-veth0 to 255.255.255.255 port 67 interval 3 (xid=0x2a702e29)
DHCPOFFER of 192.0.2.153 from 192.0.2.254
DHCPREQUEST for 192.0.2.153 on c-veth0 to 255.255.255.255 port 67 (xid=0x292e702a)
DHCPACK of 192.0.2.153 from 192.0.2.254 (xid=0x2a702e29)
bound to 192.0.2.153 -- renewal in 1432 seconds.

# 確認
ubuntu@learning-sapsucker:~$ sudo ip netns exec client ip a show
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
29: c-veth0@if30: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether 7e:6b:a9:f5:61:e3 brd ff:ff:ff:ff:ff:ff link-netns server
    inet 192.0.2.153/24 brd 192.0.2.255 scope global dynamic c-veth0
       valid_lft 3519sec preferred_lft 3519sec
    inet6 fe80::7c6b:a9ff:fef5:61e3/64 scope link
       valid_lft forever preferred_lft forever
ubuntu@learning-sapsucker:~$ sudo ip netns exec client ip route show
default via 192.0.2.254 dev c-veth0
192.0.2.0/24 dev c-veth0 proto kernel scope link src 192.0.2.153
```

## 7. NAT

- source nat
    - 送信元の ip, port を書き換える方式
        - たぶんルーターが ip, port の組をかぶらないように記憶して書き換えている
        - チェックサムなども書き換えていい感じに整合性をとっている
        - icmp の場合はポート番号の代わりに Identifier というフィールドを使っている
- iptables
    - chain は度のタイミングで処理を差し込むか
    - MASQUERADE = source NAT

```sh
ubuntu@learning-sapsucker:~$ sudo ip netns exec router iptables -t nat -L
Chain PREROUTING (policy ACCEPT)
target     prot opt source               destination

Chain INPUT (policy ACCEPT)
target     prot opt source               destination

Chain OUTPUT (policy ACCEPT)
target     prot opt source               destination

Chain POSTROUTING (policy ACCEPT)
target     prot opt source               destination

ubuntu@learning-sapsucker:~$ sudo ip netns exec router iptables -t nat -A POSTROUTING -s 192.0.2.0/24 -o gw-veth1 -j MASQUERADE
ubuntu@learning-sapsucker:~$ sudo ip netns exec router iptables -t nat -L
Chain PREROUTING (policy ACCEPT)
target     prot opt source               destination

Chain INPUT (policy ACCEPT)
target     prot opt source               destination

Chain OUTPUT (policy ACCEPT)
target     prot opt source               destination

Chain POSTROUTING (policy ACCEPT)
target     prot opt source               destination
MASQUERADE  all  --  192.0.2.0/24         anywhere
```

- destination NAT
    - wan からのパケットを lan の特定のノードに届ける設定をする
        - いわゆる「ポートを開ける」という操作
    - `-j DNAT`

```sh
ubuntu@learning-sapsucker:~$ sudo ip netns exec router iptables -t nat -A PREROUTING -p tcp --dport 54321 -d 203.0.113.254 -j DNAT --to-destination 192.0.2.1
ubuntu@learning-sapsucker:~$ sudo ip netns exec router iptables -t nat -L
Chain PREROUTING (policy ACCEPT)
target     prot opt source               destination
DNAT       tcp  --  anywhere             203.0.113.254        tcp dpt:54321 to:192.0.2.1

Chain INPUT (policy ACCEPT)
target     prot opt source               destination

Chain OUTPUT (policy ACCEPT)
target     prot opt source               destination

Chain POSTROUTING (policy ACCEPT)
target     prot opt source               destination
MASQUERADE  all  --  192.0.2.0/24         anywhere
```

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B085BG8CH5/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/51kU2EFP5UL.jpg" alt="Linuxで動かしながら学ぶTCP/IPネットワーク入門" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B085BG8CH5/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Linuxで動かしながら学ぶTCP/IPネットワーク入門</a></div><div class="amazlet-detail">もみじあめ  (著)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B085BG8CH5/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
