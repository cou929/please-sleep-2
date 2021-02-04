{"title":"net.Conn, net/http と io, bufio の整理","date":"2021-02-04T10:00:00+09:00","tags":["golang"]}

[net.Conn](https://pkg.go.dev/net#Conn) や [net/http の Server.Serve](https://pkg.go.dev/net/http#Server.Serve) あたりを利用する際、自分の理解が甘い部分があったので改めて整理したメモ。

## [net.Conn](https://pkg.go.dev/net#Conn)

- [net.Conn](https://pkg.go.dev/net#Conn) はストリーム指向のネットワーク接続
    - 今回の用途では TCP 接続を現している
    - `io.Reader` をはじめ `io.Writer` `io.Closer` を満たしている
- [net/http の Server.Serve](https://pkg.go.dev/net/http#Server.Serve) は HTTP サーバ実装の入り口
    - [net.Listener](https://pkg.go.dev/net#Listener) (listen の結果返される tcp 接続の構造体) を受け取り、goroutine を起動し、リクエストのパースなどを行い、Handler を呼び出すなどする
- リクエストのパースは [readRequest](https://github.com/golang/go/blob/e491c6eea9ad599a0ae766a3217bd9a16ca3a25a/src/net/http/server.go#L957) というプライベートメソッドが担当する
    - `net.Conn` から Read して ([こちらなど](https://github.com/golang/go/blob/e491c6eea9ad599a0ae766a3217bd9a16ca3a25a/src/net/textproto/reader.go#L57))、それをパースしている
- `net/http` では [net.Conn に bufio をかませている](https://github.com/golang/go/blob/0e85fd7561de869add933801c531bf25dee9561c/src/net/http/server.go#L1861-L1862)
    - [bufio.ReadLine](https://pkg.go.dev/bufio#Reader.ReadLine) など、生の Read よりも便利に利用している
- `net.Conn.Read` は最終的に [read(2)](https://man7.org/linux/man-pages/man2/read.2.html) を呼んでいる
    - [1](https://github.com/golang/go/blob/0e85fd7561de869add933801c531bf25dee9561c/src/net/net.go#L183), [2](https://github.com/golang/go/blob/0e85fd7561de869add933801c531bf25dee9561c/src/net/fd_posix.go#L54)

ここまでで、io, bufio の関係やシステムコールレベルでの tcp 接続の read について理解が曖昧だったというのが今回の背景。

## [io](https://pkg.go.dev/io)

- IO に関する基本的なインタフェースを提供しているパッケージ
    - `Package io provides basic interfaces to I/O primitives`
- [io.Reader.Read](https://pkg.go.dev/io#Reader)
    - 読み取れたデータがバッファのバイト数未満でも、残りのバッファ領域は内部的に利用される可能性がある
    - まだ読み取れるデータがあり、かつ読み取り済みデータ量がバッファ以下の場合、通常 Read は残りのデータを待つのではなく利用可能になった分だけを呼び出し元へ返す
    - ファイル末尾 EOF に達した場合、読み込めたバイト数 n と err == EOF を同時に返すか、err == nil を返し次の呼び出しが `n == 0 && err == EOF` を返すか、実装によって挙動にバリエーションがある
        - どちらのパターンでもその後の呼び出しが必ず `n == 0 && err == EOF` になるのは保証されている
    - 呼び出し元は err の内容に関わらず、読み取れた n byte を必ず処理しないといけない
    - Read の実装側は `n == 0 && err == nil` を返さないようにする必要がある。呼び出し側はそのようなケースでは何も起こっていない (EOF でもない) とみなすべき
- [io.Writer.Write](https://pkg.go.dev/io#Writer)
    - 書き込めたバイト数 n が指定よりも少ない場合、必ず `err != nil`
- いろいろな関数
    - [io.Copy](https://pkg.go.dev/io#Copy)
        - EOF まで読み込んで書き込む
        - 正常終了時に EOF を err として返さない (`err == nil` を返す)
        - 内部でバッファ用にメモリを確保する。それを避けたければ CopyBuffer を使えば良い
    - [io.Pipe](https://pkg.go.dev/io#Pipe)
        - つながっている reader, writer のセットを作る
        - バッファはなく、並列に読み書きしても内部で直列化
    - [io.ReadFull](https://pkg.go.dev/io#ReadFull)
        - 指定したバイト数の読み取り
        - EOF が返されるのは読み込めるデータが無いとき
        - データを読み取れたが指定したバイト数に満たずに EOF となった場合は ErrUnexpectedEOF
        - 指定したバイト数読み取れた場合は `err == nil`

### [io.EOF](https://pkg.go.dev/io#pkg-variables)

- [EOF は読み取りの正常終了](https://pkg.go.dev/io#pkg-variables) を意味する
    - そのため Reader を実装する際、規定されたデータストリームの途中で読み取り不能になった場合は EOF ではなく ErrUnexpectedEOF などを返すべき
- [EOF は Go のレイヤーで規定しているもの](https://github.com/golang/go/blob/fca94ab3ab113ceddb7934f76d0f1660cad98260/src/internal/poll/fd_posix.go#L16-L21)
    - `n == 0 && err == nil && fd.ZeroReadIsEOF` なら EOF (`n` は読み取れたバイト数、err は読み取りエラー)
    - [ZeroReadIsEOF](https://github.com/golang/go/blob/fca94ab3ab113ceddb7934f76d0f1660cad98260/src/internal/poll/fd_unix.go#L40-L42) というフラグもある
        - [SOCK_STREAM は true](https://github.com/golang/go/blob/0e85fd7561de869add933801c531bf25dee9561c/src/net/fd_unix.go#L31)

## [bufio](https://pkg.go.dev/bufio)

- 読み書きのバッファリング機能を提供 + いくつかのユーティリティ
    - NewReader, NewWriter で io.Reader, io.Writer にバッファリングを追加できる
- バッファリングすると?
    - 書き込み
        - 一定のデータを内部で貯めてからまとめて書き込みを行う
        - 書きこみする先のもの (デバイスやネットワーク等) に細かく何度も IO することを防ぐことができる
            - 極端に言うと、そうでない場合 (たいした量や回数の読み書きをしない場合) はバッファリングし無くてもシンプルでいいのかも?
    - 読み込み
        - 一定のデータを内部に貯めて、そこで充足するなら読み出し要求にはバッファから返す
        - 小さく頻度の多い Read 要求が多い場合、バッファを導入したほうが実 IO 数を減らせる
        - 加えてバッファに一度貯めることで ReadLine や Scanner 系など、アプリケーションからのデータ読み取りに便利なツールも載せられている
    - [APUE](http://www.amazon.co.jp/exec/obidos/ASIN/B00KRB9U8K/pleasesleep-22/ref=nosim/) の 5.4 より引用

> 標準入出力ライブラリでバッファリングする目標は、readとwriteの呼び出し回数を最小にすることです。(図3.6では、さまざまなバッファサイズを用いた入出力動作に必要なCPU時間を示した。)さらに、各入出力ストリームのバッファリングを自動化し、アプリケーションで気にする必要がないようにします。残念ながら、もっとも混乱しやすい標準入出力ライブラリの一面がバッファリングです。

- 例えば [Reader](https://github.com/golang/go/blob/32e789f4fb45b6296b9283ab80e126287eab4db5/src/bufio/bufio.go#L32-L39) だとラップ対象の `io.Reader` (underlying Reader) とバッファ (`[]byte`)、あとはバッファ内の読み取りカーソルなどを保持している

```go
type Reader struct {
	buf          []byte
	rd           io.Reader // reader provided by the client
	r, w         int       // buf read and write positions
	err          error
	lastByte     int // last byte read for UnreadByte; -1 means invalid
	lastRuneSize int // size of last rune read for UnreadRune; -1 means invalid
}
```

- [bufio.Reader.Read](https://pkg.go.dev/bufio#Reader.Read)
    - 最大 1 回 underlying Reader の Read を読む
        - [バッファにあればそれを使う](https://github.com/golang/go/blob/32e789f4fb45b6296b9283ab80e126287eab4db5/src/bufio/bufio.go#L238)
    - 任意のバイト数確実に読みたい場合は `io.ReadFull` などを使う
    - EOF の場合は n は必ず 0
- [Scanner](https://pkg.go.dev/bufio#Scanner)
    - 改行区切りのテキストなど、特定のデリミタ (SplitFunc) ごとにデータを読み取るための便利なインタフェース
    - なんとなくユーザーからすると、bufio というパッケージ名なのでこの機能は別パッケージなっていたほうがわかりやすい気がした
        - けれども bufio の上に Scanner が乗っているし、他に良い置き場所もなさそう、独立させるには分量が少なそうなので、今が最適解という気もする

## tcp と read(2)

[APUE](http://www.amazon.co.jp/exec/obidos/ASIN/B00KRB9U8K/pleasesleep-22/ref=nosim/) と [UNIXネットワークプログラミング](http://www.amazon.co.jp/exec/obidos/ASIN/4894712059/pleasesleep-22/ref=nosim/) を読み返し。

- APUE 16.2ソケット記述子 より

> ソケットSOCK_STREAMはバイトストリームサービスを提供します。アプリケーションにはメッセージ境界は分りません。つまり、ソケットSOCK_STREAMからデータを読み取ると、送り手が書き出したバイト数と同じ数を返さないこともあります。最終的には送られたものをすべて受け取りますが、それには数回の関数呼び出しが必要になるでしょう。

- [`man socket(2)` より](https://man7.org/linux/man-pages/man2/socket.2.html)
    - 第二引数でタイプ指定 (`SOCK_STREAM`、`SOCK_DGRAM` など)
- tcp はバイトストリーム指向のプロトコル
    - つまり接続している間は区切り無くデータを read できる
    - 接続がクローズされているときに read できたバイト数が 0 になる
        - このケースを Go が `io.EOF` として扱う
    - tcp よりは上のレイヤーのプロトコル目線で、アプリケーションが気にするべきこと
        - 一度の read 呼び出しで、アプリケーションが必要とするデータが読み取れない場合が正常であると認識する必要がある
        - ループで複数回 read を呼び出すような実装が必要
            - 当たり前だけど buf より長いメッセージが来たら複数回の read 呼び出しが必要
        - データの「区切り」は自分で処理する必要がある
            - http/1x ならテキストとして改行区切りで扱うなど
- なお udp はデータグラム通信
    - バイトストリームに対して、固定最大長、接続という概念がない、非信頼性の通信など

## ここまでの実験

[delve, lsof, strace \(dtruss\), gops で Go のプロセスを解析する \- Please Sleep](https://please-sleep.cou929.nu/diagnosis-tools-of-go-process.html) とほぼ同じだが、自分の理解のため改めて確認した。

- 以下のサンプルコードの動作を追ってみる
    - 2000 番ポートで tcp 接続を待ち受け
    - 接続確立したら read
    - strace でシステムコールの状況確認
    - 今回は ubuntu16.04 でやってみた

```sh
vagrant@ubuntu-xenial:~$ cat /etc/lsb-release
DISTRIB_ID=Ubuntu
DISTRIB_RELEASE=16.04
DISTRIB_CODENAME=xenial
DISTRIB_DESCRIPTION="Ubuntu 16.04.7 LTS"
```

```go
// main.go
package main

import (
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go func(c net.Conn) {
			for {
				buf := make([]byte, 100)
				n, err := c.Read(buf)
				if err != nil {
					log.Println(n, "read error", err)
					return
				}
				log.Println(n, string(buf))
			}
		}(conn)
	}
}
```

- 実行手順例

```sh
$ go build -o svr . 
$ sudo strace -t -f ./svr

# 別セッションで
$ telnet 127.0.0.1 2000
  // 適当にデータを送信する
  // 最後にコネクションを切る
^]
telnet> quit
Connection closed.
```

- strace の実行
    - 今回のサンプルでは接続確立後に goroutine (スレッド) を起動しているので [`-f` オプション](https://man7.org/linux/man-pages/man1/strace.1.html) も必要
        - `sudo strace -t -f -p <PID>` など
- strace の結果
    - socket は NONBLOCK でオープンされ、[epoll](https://man7.org/linux/man-pages/man2/epoll_ctl.2.html) と組み合わせて使われている
    - Go 側で指定した 100 バイトずつ read している
    - close すると read の読み取り結果が 0 バイトで、Go 側では EOF としてハンドリングしている
    - 100 バイト以上のデータをクライアントから送信した場合、複数回 read が走る

```c
// socket は SOCK_NONBLOCK で開かれる
// fd は 3 番
[pid  1997] 04:54:51 socket(PF_INET6, SOCK_STREAM|SOCK_CLOEXEC|SOCK_NONBLOCK, IPPROTO_IP) = 3

// bind に fd 3 番を渡す
[pid  1997] 04:54:51 bind(3, {sa_family=AF_INET6, sin6_port=htons(2000), inet_pton(AF_INET6, "::", &sin6_addr), sin6_flowinfo=0, sin6_scope_id=0}, 28) = 0

// listen も fd 3 番を渡す
[pid  1997] 04:54:51 listen(3, 128)     = 0

// epoll
// epoll_ctl で fd 3 を EPOLL_CTL_ADD
// socket が nonblock なので accept からは EAGAIN が即座に返っている (はず)
[pid  1997] 04:54:51 epoll_create1(EPOLL_CLOEXEC) = 4
[pid  1997] 04:54:51 epoll_ctl(4, EPOLL_CTL_ADD, 3, {EPOLLIN|EPOLLOUT|EPOLLRDHUP|EPOLLET, {u32=1796096696, u64=139940420534968}}) = 0
[pid  1997] 04:54:51 getsockname(3, {sa_family=AF_INET6, sin6_port=htons(2000), inet_pton(AF_INET6, "::", &sin6_addr), sin6_flowinfo=0, sin6_scope_id=0}, [28]) = 0
[pid  1997] 04:54:51 accept4(3, 0xc820043b28, 0xc820043b24, SOCK_CLOEXEC|SOCK_NONBLOCK) = -1 EAGAIN (Resource temporarily unavailable)  
[pid  1997] 04:54:51 epoll_wait(4, [], 128, 0) = 0

// 通信が確立できた際の accept
// fd 5 番が返される
[pid  1997] 05:00:03 accept4(3, {sa_family=AF_INET6, sin6_port=htons(60258), inet_pton(AF_INET6, "::ffff:127.0.0.1", &sin6_addr), sin6_flowinfo=0, sin6_scope_id=0}, [28], SOCK_CLOEXEC|SOCK_NONBLOCK) = 5

// 細かく見ていないが accept 後は select も多用されているのが見える
[pid  1998] 05:00:03 select(0, NULL, NULL, NULL, {0, 20}) = 0 (Timeout)

// fd 5 から read
// クライアントから送られた文字列を読み取ることができる
[pid  1997] 05:01:00 read(5, "abc abc\r\n", 100) = 9

// 100 バイト以上のデータをクライアントから送った場合
// 2 度の read で全体が読み込まれている
[pid  1999] 05:03:18 read(5, "01234567890123456789012345678901"..., 100) = 100
[pid  1999] 05:03:18 read(5, "123\r\n", 100) = 5

// このときの Go からの標準出力
// こちらも 100 バイトずつ 2 度の Read 呼び出しに分かれている
2021/02/04 05:03:18 100 012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678
2021/02/04 05:03:18 5 123

// 接続の close 時には読み取りバイト数が 0
[pid  1999] 05:06:10 read(5, "", 100) = 0

// このときの Go からの標準出力
// EOF として扱われている
2021/02/04 05:06:10 0 read error EOF
```

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00KRB9U8K/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/51L6CwNG11L.jpg" alt="詳解UNIXプログラミング 第3版" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00KRB9U8K/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">詳解UNIXプログラミング 第3版</a></div><div class="amazlet-detail">W. Richard Stevens  (著), Stephen A. Rago (著), 大木 敦雄 (監修)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00KRB9U8K/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4894712059/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://images-na.ssl-images-amazon.com/images/I/213B9PVJD1L._BO1,204,203,200_.jpg" alt="UNIXネットワークプログラミング〈Vol.1〉ネットワークAPI:ソケットとXTI" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4894712059/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">UNIXネットワークプログラミング〈Vol.1〉ネットワークAPI:ソケットとXTI</a></div><div class="amazlet-detail">W.リチャード スティーヴンス (著), W.Richard Stevens (原著), 篠田 陽一 (翻訳)<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4894712059/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
