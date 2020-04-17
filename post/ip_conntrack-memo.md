{"title":"ip_conntrack メモ","date":"2012-04-04T21:59:53+09:00","tags":["nix"]}

- iptables を使うとき, ip_conntrack というテーブルで tcp のセッションを管理している
- 上限値は 65536 で, 最近のスペックのマシンだとこの数はネックになりがち
  - ソフトウエアロードバランサにしているサーバとか, とにかくリクエストをうけるサーバは注意しておく必要がある
- 上限値は procfs でみる
  - /proc/sys/net/ipv4/ip_conntrack_max
  - ipv4 ってついてるので, v6 はまた別なのかも
- ip_conntrack_max を超えると syslog にログがでる
  - /var/log/message か dmseg あたりにこんなのが出るらしい

            ip_conntrack: maximum limit of 65536 entries exceeded

- 設定は /etc/sysctl.conf に書いて sysctl -p か,  procfs に直接書きこんでもいいらしい

        $ sudo sysctl -p
        $ echo '100000' >  /proc/sys/net/ipv4/ip_conntrack_max

- 現状値確認はこれらしい. ちょっと自信なし

        $ cat /proc/net/ip_conntrack | wc -l

- 350 byte / 1接続 (swap 領域は使用できない) 使うらしいのでメモリも考慮
- 再起動しても設定がクリアされないように, /etc/sysctl.conf に書いておくといいみたい

        net.ipv4.netfilter.ip_conntrack_max = 2000000

  - ただうまくいかないこともあるっぽいので, 監視も入れとくといいかも
    - [ip_conntrack_maxの設定値を監視するNagiosプラグインを書いた - Lism.in * blog - nekoya (id:studio-m)](http://d.hatena.ne.jp/studio-m/20110505/1304570658)

### iptables

- よくわかってないから実装追いたいけど優先度は低い.
- パケットフィルタリングだけじゃなくて NAT もやるらしい
- Linux のもの
- kernel 2.4 以降だとnetfilter ってのが下にいるらしい

### Refs

- [中〜大規模サーバーを運用するときの勘所 – iptablesとip_conntrack \| cyano](http://www.onflow.jp/cyano/archives/103)
- [iptables の ip_conntrack の最大値を変更する方法 \| Carpe Diem](http://www.sssg.org/blogs/naoya/archives/1454)
- [ip_conntrack_maxの設定値を監視するNagiosプラグインを書いた - Lism.in * blog - nekoya (id:studio-m)](http://d.hatena.ne.jp/studio-m/20110505/1304570658)
- [iptables - Wikipedia](http://ja.wikipedia.org/wiki/Iptables)
