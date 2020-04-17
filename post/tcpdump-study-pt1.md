{"title":"tcpdump の見方を勉強","date":"2013-11-10T12:00:00+09:00","tags":["nix"]}

例として www.example.com への HTTP GET のパケットキャプチャをやってみる。

    $ sudo tcpdump host www.example.com and port 80
    tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
    listening on en0, link-type EN10MB (Ethernet), capture size 65535 bytes
    09:36:20.760358 IP 10.0.1.23.65428 > 93.184.216.119.http: Flags [S], seq 2250708012, win 65535, options [mss 1460,nop,wscale 4,nop,nop,TS val 938288017 ecr 0,sackOK,eol], length 0
    09:36:20.885412 IP 93.184.216.119.http > 10.0.1.23.65428: Flags [S.], seq 1676582138, ack 2250708013, win 14600, options [mss 1400,nop,nop,sackOK,nop,wscale 6], length 0
    09:36:20.885482 IP 10.0.1.23.65428 > 93.184.216.119.http: Flags [.], ack 1, win 16384, length 0
    09:36:20.885587 IP 10.0.1.23.65428 > 93.184.216.119.http: Flags [P.], seq 1:147, ack 1, win 16384, length 146
    09:36:21.010182 IP 93.184.216.119.http > 10.0.1.23.65428: Flags [.], ack 147, win 245, length 0
    09:36:21.016537 IP 93.184.216.119.http > 10.0.1.23.65428: Flags [.], seq 1:1401, ack 147, win 245, length 1400
    09:36:21.016817 IP 93.184.216.119.http > 10.0.1.23.65428: Flags [P.], seq 1401:1592, ack 147, win 245, length 191
    09:36:21.016863 IP 10.0.1.23.65428 > 93.184.216.119.http: Flags [.], ack 1592, win 16372, length 0
    09:36:21.017103 IP 10.0.1.23.65428 > 93.184.216.119.http: Flags [F.], seq 147, ack 1592, win 16384, length 0
    09:36:21.146290 IP 93.184.216.119.http > 10.0.1.23.65428: Flags [F.], seq 1592, ack 148, win 245, length 0
    09:36:21.146399 IP 10.0.1.23.65428 > 93.184.216.119.http: Flags [.], ack 1593, win 16384, length 0

まずは各行の見方について。左から順に、

- タイムスタンプ
- IP (?)
- 送信元 (IP アドレスとポート)
- 送信先 (IP アドレスとポート)
- フラグ (ACK, SYN, FIN など。どれも立っていなければ ".")
- seq (シーケンス番号)
  - "1:1401" のような表記の場合は "開始シーケンス番号:終了シーケンス番号" という意味
- ack (ACK 番号)
- win (ウィンドウサイズ)
- blacket に囲まれている部分はその他オプション
- length (送信バイト数)

### 3 way handshake

最初の 3 行はいわゆるスリーウェイハンドシェイク。

    09:36:20.760358 IP 10.0.1.23.65428 > 93.184.216.119.http: Flags [S], seq 2250708012, win 65535, options [mss 1460,nop,wscale 4,nop,nop,TS val 938288017 ecr 0,sackOK,eol], length 0
    09:36:20.885412 IP 93.184.216.119.http > 10.0.1.23.65428: Flags [S.], seq 1676582138, ack 2250708013, win 14600, options [mss 1400,nop,nop,sackOK,nop,wscale 6], length 0
    09:36:20.885482 IP 10.0.1.23.65428 > 93.184.216.119.http: Flags [.], ack 1, win 16384, length 0

まずはローカル (10.0.1.23) からサーバ (93.184.216.119, www.example.com) へ SYN パケットが送られる。ポート番号はサーバ側は 80 番へ、クライアント側は適当に割り振られた空きポートである 65428 だ。

最初の行はクライアントからサーバへの SYN リクエスト。シーケンス番号は 2250708012。length は 0 になっている。

次の行はサーバからの ACK と SYN。ACK の番号がさきほどクライアントから送られたシーケンス番号に +1 したものになっている。またサーバ側のシーケンス番号として 1676582138 が送られる。

三行目はクライアントからの ACK。番号が 1 になっているが、tcpdump がわかりやすく表示しているもので、本当は受け取ったシーケンス番号 +1 である 1676582139 だ。これ以降のシーケンス番号はすべて 1 が起点で表示される。tcpdump に `-S` オプションをつけるとそのままのシーケンス番号を表示してくれる。

### データの転送

    09:36:20.885587 IP 10.0.1.23.65428 > 93.184.216.119.http: Flags [P.], seq 1:147, ack 1, win 16384, length 146
    09:36:21.010182 IP 93.184.216.119.http > 10.0.1.23.65428: Flags [.], ack 147, win 245, length 0
    09:36:21.016537 IP 93.184.216.119.http > 10.0.1.23.65428: Flags [.], seq 1:1401, ack 147, win 245, length 1400
    09:36:21.016817 IP 93.184.216.119.http > 10.0.1.23.65428: Flags [P.], seq 1401:1592, ack 147, win 245, length 191
    09:36:21.016863 IP 10.0.1.23.65428 > 93.184.216.119.http: Flags [.], ack 1592, win 16372, length 0

1 行目。クライアントからサーバへ 146 バイトのデータが送られている。リクエストの中身を表示していないが、これは HTTP の GET リクエストだ。シーケンス番号の部分、`1:147` となっているのは "開始シーケンス番号:終了シーケンス番号" という意味になる。シーケンス番号は転送されたバイト数増えるので、ここでは 1 + 146 の値が終了シーケンス番号として示されている。[TCP のセグメント構造](http://ja.wikipedia.org/wiki/Transmission_Control_Protocol#TCP.E3.82.BB.E3.82.B0.E3.83.A1.E3.83.B3.E3.83.88.E6.A7.8B.E9.80.A0) をみると、"終了シーケンス番号" というものは含まれていないし、seq とlength から計算で導くことができる情報だ。よっておそらくこれも見やすさのため tcpdump が表示してくれているのだろう。

2 行目でサーバが ack を返したあと、3、4 行目でサーバからクライアントへデータが転送されている。サーバからの ack 番号 147 について、シーケンス番号 1 から 1592 まで、1591 バイトの情報が 2 回に分けて転送されていることがわかる。こちらも中身を表示していないが、www.example.com の HTML が転送されていることになる。ちなみに tcpdump に -A というオプションをつけるとパケットの内容を ASCII で表示してくれる。

5 行目では転送が完了し ack がサーバへ返されている。

1、4 行目では PSH フラグがたっていることがわかる。PSH は "バッファに蓄えたデータをアプリケーションにプッシュする（押しやる）ことを依頼する。" というフラグだ。それぞれデータ転送が伴う最後の通信の際に立っているように見える。すべてのデータが転送し終わったため、バッファのデータをすべてアプリへ送ることを依頼していると考えてよいだろうか。

### コネクションのクローズ

    09:36:21.017103 IP 10.0.1.23.65428 > 93.184.216.119.http: Flags [F.], seq 147, ack 1592, win 16384, length 0
    09:36:21.146290 IP 93.184.216.119.http > 10.0.1.23.65428: Flags [F.], seq 1592, ack 148, win 245, length 0
    09:36:21.146399 IP 10.0.1.23.65428 > 93.184.216.119.http: Flags [.], ack 1593, win 16384, length 0

まずはクライアントからサーバへ FIN を送信。seq, ack の番号はさきほどみた転送作業がすべておわった時点のものだ。つぎにサーバは ack と FIN を返す。ack の番号は送られたものの +1。最後にクライアントが FIN に対する ack を返して完了だ。

### 参考

- [Transmission Control Protocol - Wikipedia](http://ja.wikipedia.org/wiki/Transmission_Control_Protocol)
- [UNIXの部屋 コマンド検索:tcpdump (*BSD/Linux)](http://x68000.q-e-d.net/~68user/unix/pickup?tcpdump)
- [TCPDUMPの出力を見てみよう](http://net-newbie.com/tcpip/tcp/tcpdump.html)

今回は web の情報と実際にコマンドをたたいてみた情報のみ参考にしたが、きちんと体系的に勉強したい。

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0827QNDNT/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://images-fe.ssl-images-amazon.com/images/I/41GcGJkZ-6L._SL160_.jpg" alt="マスタリングTCP/IP　入門編（第6版）" style="border: none;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0827QNDNT/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">マスタリングTCP/IP　入門編（第6版）</a><div class="amazlet-powered-date" style="font-size:80%;margin-top:5px;line-height:120%">posted with <a href="http://www.amazlet.com/" title="amazlet" target="_blank">amazlet</a> at 20.03.08</div></div><div class="amazlet-detail">オーム社 (2019-12-06)<br />売り上げランキング: 4,995<br /></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0827QNDNT/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
