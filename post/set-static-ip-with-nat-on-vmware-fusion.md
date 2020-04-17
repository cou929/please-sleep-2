{"title":"VMWare Fusion でゲスト OS の IP を固定する","date":"2013-11-01T19:33:30+09:00","tags":["nix"]}

NAT でゲスト OS の IP を固定し、ホスト OS の hosts でその IP を指定するという方針。

- NAT Gateway Address を確認する
  - `/Library/Preferences/VMware Fusion/vmnet8/nat.conf` に仮想ルータの IP が記述されているので確認する。以下のような行をみればよい。VMWare 特有な手順はここだけ

            # NAT gateway address
            ip = 192.168.112.2

- ゲスト OS のネットワーク設定から IP を固定する
  - デフォルトではおそらく DHCP になっているので static にする
  - GATEWAY には先ほど調べた NAT Gateway Address の ip を設定する
  - IPADDR には固定したい IP を記述
  - たとえば CentOS の場合このようにする。これは `192.168.112.200` に固定する例

            $ cat /etc/sysconfig/network-scripts/ifcfg-eth0
            DEVICE=eth0
            ONBOOT=yes
            HWADDR=00:0c:29:30:9d:96
            TYPE=Ethernet
            BOOTPROTO=static
            DNS1=8.8.8.8
            USERCTL=no
            PEERDNS=yes
            IPV6INIT=no
            IPADDR=192.168.112.200
            NETMASK=255.255.255.0
            GATEWAY=192.168.112.2

- ホスト OS 側の hosts ファイルに固定した IP を追加
  - `/etc/hosts` に書くだけ

            192.168.112.200 localdev

  - この例は仮に localdev というホスト名にしている


- 参考サイト
  - [ゲストOSのIPを固定したい件 | taichino.com](http://taichino.com/engineer-life/linux/430)
  - [vmware で centos をコピーするとネットワーク(eth0)にアクセスできなくなる \- Yahoo!ジオシティーズ](http://geocities.yahoo.co.jp/gl/ds301b/view/20110928/1317221758)
    - eth0 が見えなくなってしまうことがあるらしい
