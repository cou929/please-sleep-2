{"title":"delve, lsof, strace (dtruss), gops で Go のプロセスを解析する","date":"2020-06-22T13:30:00+09:00","tags":["go","nix"]}

Go のプロセスやバイナリの動きを調べるいろいろなツールについて、使い方を確認したメモ。

なお、標準パッケージ ~ システムコールレベルの動きを確認したいというモチベーションだったので、pprof などのプロファイリング系のツールは対象外。また OSX 環境を想定しています (そのせいで若干のヤクの毛刈りも発生した)。

## サンプルコード

次のような、listen, accept し、同じ内容を返す tcp のエコーサーバを題材にした。[net.Listener のサンプルコードそのまま](https://pkg.go.dev/net?tab=doc#example-Listener) のコード。

```go
// main.go
package main

import (
	"io"
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
		go func(c net.Conn) {
			io.Copy(c, c)
			c.Close()
		}(conn)
	}
}
```

このサーバプロセスを起動し、telnet などで接続すると、リクエストがそのまま返ってくる様子が確認できる。

```
$ go build -o svr .
$ ./svr
```

```
$ telnet 127.0.0.1 2000
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
foo bar baz
foo bar baz
^]
telnet> quit
Connection closed.
```

## delve

何はともあれデバッガでステップインしていけば、深い部分の処理を見やすい。半ばスタンダードなデバッガは [Delve](https://github.com/go-delve/delve) で、公式にも [gdb の better alternative](https://golang.org/doc/gdb) と紹介されている。

### ステップ実行

ステップ実行はふつうに gdb のように操作できる。

```
$ dlv debug main.go
Type 'help' for list of commands.

(dlv) b main.main:1
Breakpoint 1 set at 0x11348e2 for main.main() ./main.go:10
(dlv) c
> main.main() ./main.go:10 (hits goroutine(1):1 total:1) (PC: 0x11348e2)
     5:         "log"
     6:         "net"
     7: )
     8:
     9: func main() {
=>  10:         l, err := net.Listen("tcp", ":2000")
    11:         if err != nil {
    12:                 log.Fatal(err)
    13:         }
    14:         defer l.Close()
    15:
(dlv) s
> net.Listen() /usr/local/Cellar/go/1.14.2_1/libexec/src/net/dial.go:705 (PC: 0x10fa188)
   700: // The Addr method of Listener can be used to discover the chosen
   701: // port.
   702: //
   703: // See func Dial for a description of the network and address
   704: // parameters.
=> 705: func Listen(network, address string) (Listener, error) {
   706:         var lc ListenConfig
   707:         return lc.Listen(context.Background(), network, address)
   708: }
   709:
   710: // ListenPacket announces on the local network address.
```

なお [vscode とのインテグレーション](https://github.com/golang/vscode-go/blob/master/docs/debugging.md) もちゃんとしているし、ステップ実行はエディタからやったほうが便利。

### goroutine や thread の確認

delve が良いのは、Go のランタイム、データ構造、構文を gdb よりもちゃんとサポートしているところ。例えば次のように goroutine の諸々を見ることもできる。

- 現在の goroutine の stack の表示

```
(dlv) goroutine
Thread 322976 at /usr/local/Cellar/go/1.14.2_1/libexec/src/net/dial.go:707
Goroutine 1:
        Runtime: /usr/local/Cellar/go/1.14.2_1/libexec/src/net/dial.go:707 net.Listen (0x10fa1bd)
        User: /usr/local/Cellar/go/1.14.2_1/libexec/src/net/dial.go:707 net.Listen (0x10fa1bd)
        Go: /usr/local/Cellar/go/1.14.2_1/libexec/src/runtime/asm_amd64.s:220 runtime.rt0_go (0x1067986)
        Start: /usr/local/Cellar/go/1.14.2_1/libexec/src/runtime/proc.go:113 runtime.main (0x103a2a0)
```

- 現在起動している goroutine の一覧

```
(dlv) goroutines
* Goroutine 1 - User: /usr/local/Cellar/go/1.14.2_1/libexec/src/net/dial.go:707 net.Listen (0x10fa1bd) (thread 322976)
  Goroutine 2 - User: /usr/local/Cellar/go/1.14.2_1/libexec/src/runtime/proc.go:305 runtime.gopark (0x103a82b)
  Goroutine 3 - User: /usr/local/Cellar/go/1.14.2_1/libexec/src/runtime/proc.go:305 runtime.gopark (0x103a82b)
  Goroutine 4 - User: /usr/local/Cellar/go/1.14.2_1/libexec/src/runtime/proc.go:305 runtime.gopark (0x103a82b)
  Goroutine 18 - User: /usr/local/Cellar/go/1.14.2_1/libexec/src/runtime/proc.go:305 runtime.gopark (0x103a82b)
[5 goroutines]
```

- 現在起動しているスレッドの一覧

```
(dlv) threads
* Thread 322976 at 0x10fa1bd /usr/local/Cellar/go/1.14.2_1/libexec/src/net/dial.go:707 net.Listen
  Thread 323098 at :0
  Thread 323099 at :0
  Thread 323100 at :0
  Thread 323101 at :0
  Thread 323102 at :0
```

- 特定の goroutine に切り替えて stack を表示

```
(dlv) goroutine 2
Switched from 1 to 2 (thread 322976)
(dlv) bt
0  0x000000000103a82b in runtime.gopark
   at /usr/local/Cellar/go/1.14.2_1/libexec/src/runtime/proc.go:305
1  0x000000000103a8e3 in runtime.goparkunlock
   at /usr/local/Cellar/go/1.14.2_1/libexec/src/runtime/proc.go:310
2  0x000000000103a6ca in runtime.forcegchelper
   at /usr/local/Cellar/go/1.14.2_1/libexec/src/runtime/proc.go:253
3  0x0000000001069b71 in runtime.goexit
   at /usr/local/Cellar/go/1.14.2_1/libexec/src/runtime/asm_amd64.s:1373
```

## lsof

[lsof](https://linux.die.net/man/8/lsof) は一般的な Linux/Unix 系ツールだが、あるプロセスが開いているファイルを調べることができる。

今回の例題では、`Listen` を呼び出した時点で LISTEN 状態のリスニングソケットが、クライアントとして telnet などで接続すると、その接続の数だけサーバは `Accept` し、ESTABLISHED 状態の接続済みソケットができていると予想される。この挙動を実際にサーバプロセスがオープンしているファイルを見て調べてみる。

### リスニングソケットの確認

まずはビルドしてサーバを起動。

```
$ go build -o svr .
$ ./svr
```

このプロセスを見てみると、予想通り port 2000 番で LISTEN 状態のリスニングソケットができている。特に他に処理もしていないせいか、ファイルディスクリプタは (標準エラー出力の次の) 3 だった。

```
2020-06-22 10:44 go master$ lsof -p 16329 -P -n
COMMAND   PID   USER   FD     TYPE             DEVICE SIZE/OFF                NODE NAME
...
svr     16329 cou929    3u    IPv6 0xb7376df7b15138e1      0t0                 TCP *:2000 (LISTEN)
```

なお FD の数字あとにある `u` というアルファベットは、読み書き可能ということを表しているらしい。

```
// man lsof

FD is followed by one of these characters, describing the mode under which the file is open:

    r for read access;
    w for write access;
    u for read and write access;
    space if mode unknown and no lock
        character follows;
    `-' if mode unknown and lock
        character follows.
```

### コネクションの確立とクローズの確認

ここから、telnet で 2 本コネクションをはってみる。

```
$ telnet 127.0.0.1 2000
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.

$ telnet 127.0.0.1 2000
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
```

予想通り ESTABLISHED なソケットが 2 つできている。

```
2020-06-22 10:44 go master$ lsof -p 16329 -P -n
COMMAND   PID   USER   FD     TYPE             DEVICE SIZE/OFF                NODE NAME
...
svr     16329 cou929    3u    IPv6 0xb7376df7b15138e1      0t0                 TCP *:2000 (LISTEN)
...
svr     16329 cou929    7u    IPv6 0xb7376df7b1511a41      0t0                 TCP 127.0.0.1:2000->127.0.0.1:57032 (ESTABLISHED)
svr     16329 cou929    8u    IPv6 0xb7376df7b1514521      0t0                 TCP 127.0.0.1:2000->127.0.0.1:57034 (ESTABLISHED)
```

当然、クライアントからコネクションを閉じると、サーバが開いているソケットも一つ減る。

```
$ telnet 127.0.0.1 2000
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
^]
telnet> quit
Connection closed.
```

```
$ lsof -p 16329 -P -n
COMMAND   PID   USER   FD     TYPE             DEVICE SIZE/OFF                NODE NAME
...
svr     16329 cou929    3u    IPv6 0xb7376df7b15138e1      0t0                 TCP *:2000 (LISTEN)
...
svr     16329 cou929    7u    IPv6 0xb7376df7b1511a41      0t0                 TCP 127.0.0.1:2000->127.0.0.1:57032 (ESTABLISHED)
```

## strace (dtruss)

こちらも Go に限らない一般的なツールだが、[strace](https://linux.die.net/man/1/strace) でそのプロセスが発行しているシステムコールやシグナルの状況を見ることができる。OSX だと [dtruss](https://opensource.apple.com/source/dtrace/dtrace-147/DTTk/dtruss.auto.html) が代替となるそうだ。

今回の題材は、Go のレイヤーでは [Listen](https://pkg.go.dev/net?tab=doc#Listen) と [Accept](https://pkg.go.dev/net?tab=doc#TCPListener.Accept) の 2 つのメソッドに抽象化されているが、システムコールだと [socket](https://linux.die.net/man/2/socket)、[bind](https://linux.die.net/man/2/bind)、[listen](https://linux.die.net/man/2/listen)、[accept](https://linux.die.net/man/2/accept) が呼ばれていると予想される。この挙動を dtruss でみてみる。

### SIP の無効化

OSX 特有の話だが、El Capitan 以降では [System Integrity Protection (SIP)](https://support.apple.com/en-us/HT204899) という、一部の機能を制限するセキュリティ機能が導入されているらしい。dtruss の利用に必要な一部機能もこの仕組みにより制約を受けている。まずはこの SIP を部分的に無効化する。

SIP 設定状況は [`csrutil`](https://developer.apple.com/library/archive/documentation/Security/Conceptual/System_Integrity_Protection_Guide/ConfiguringSystemIntegrityProtection/ConfiguringSystemIntegrityProtection.html) というツールで確認できる。

```
$ csrutil status
System Integrity Protection status: unknown (Custom Configuration).

Configuration:
        Apple Internal: disabled
        Kext Signing: enabled
        Filesystem Protections: enabled
        Debugging Restrictions: enabled
        DTrace Restrictions: disabled
        NVRAM Protections: enabled
        BaseSystem Verification: enabled

This is an unsupported configuration, likely to break in the future and leave your machine in an unknown state.

# 変更後にコマンドを実行したので一部の configuration が disabled になっている
```

一方で設定変更はセーフモード下でしかできないので、次の手順を踏む必要がある。

- PC をセーフモードで起動する
    - 起動時に cmd+R を押しておく
- セーフモード起動後にコンソールを開く
- 次のコマンドで DTrace Restrictions のみ無効化する
    - `$ csrutil enable --without dtrace`
- リスタートし通常モードで起動する

なお DTrace Restrictions が enabled なまま dtruss を起動すると、次のようなワーニングが出るのでわかる。

```
dtrace: system integrity protection is on, some features will not be available
```

このワーニングは出るものの、一応コマンドの実行自体はできるが、制限をうけていない機能しか使えない。制限をうけている機能を使おうとした際は次のようなエラーメッセージになる。

```
dtrace: error on enabled probe ID ...
```

### dtruss でシステムコールの呼び出しを見る

起動済みのプロセスにアタッチすることもできるが、今回は起動時の初期化処理で呼んでいるシステムコールを見たいので、dtruss 経由でプロセスを立ち上げる。

`sudo dtruss -a <COMMAND>` とすると、呼び出したシステムコールとその引数、返り値、所要時間、経過時間などが確認できる。

```
$ sudo dtruss -a ./svr
        PID/THRD  RELATIVE  ELAPSD    CPU SYSCALL(args)                  = return
29738/0x5d6ef:      1178     159     89 open("/dev/dtracehelper\0", 0x2, 0xFFFFFFFFEFBFEED0)             = 3 0
29738/0x5d6ef:      2358    2378   1177 ioctl(0x3, 0x80086804, 0x7FFEEFBFEDE0)           = 0 0
29738/0x5d6ef:      2378      70     13 close(0x3)               = 0 0
...
```

サーバプロセスを起動し、クライアントから接続しない状態で、次のように `socket`、`bind`、`listen` が呼び出されている様子が確認できる。

```
29738/0x5d6ef:      6295      16     14 socket(0x1E, 0x1, 0x0)           = 3 0
...
29738/0x5d6ef:      6415       9      8 bind(0x3, 0xC00012411C, 0x1C)            = 0 0
29738/0x5d6ef:      6424      13      7 listen(0x3, 0x80, 0x0)           = 0 0
...
```

### man やヘッダファイルから仕様とフラグを調べる

`socket` の呼び出しをもうすこし細かく見てみる。

```
socket(0x1E, 0x1, 0x0)           = 3 0
```

引数や返り値の仕様はふつうに `man 2 socket` などとして調べる。

```
SYNOPSIS
     #include <sys/socket.h>

     int
     socket(int domain, int type, int protocol);

...

DESCRIPTION
...
     The domain parameter specifies a communications domain within which communication will take place; this selects the protocol family which should be used.  These families are defined in the include file <sys/socket.h>.  The currently understood formats are

           PF_LOCAL        Host-internal protocols, formerly called PF_UNIX,
           PF_UNIX         Host-internal protocols, deprecated, use PF_LOCAL,
           PF_INET         Internet version 4 protocols,
           PF_ROUTE        Internal Routing protocol,
           PF_KEY          Internal key-management function,
           PF_INET6        Internet version 6 protocols,
           PF_SYSTEM       System domain,
           PF_NDRV         Raw access to network device
...
RETURN VALUES
     A -1 is returned if an error occurs, otherwise the return value is a descriptor referencing the socket.
```

戻り値はエラーでなければファイルディスクリプタを返す仕様のようだ。

- 今回は `3` を返していることがわかる
- その後の bind, listen に `3` を渡している様子も確認できる

第一引数は `domain` で内容は `sys/socket.h` に定義されているようだ。

- 今回は `0x1E` (= 10 進数で 30) で呼び出していることがわかる
- `sys/socket.h` によるとこれは `PF_INET6` らしい

```c
#define AF_INET6        30              /* IPv6 */
...
#define PF_INET6        AF_INET6
```

- 実際に [net パッケージの実装](https://github.com/golang/go/blob/bd486c39bad4c9f90190ae58de8a592bb9a2aae9/src/net/ipsock_posix.go#L119-L122) もその様になっている
    - デバッガで追うとわかりやすい

<div></div>

```go
func favoriteAddrFamily(network string, laddr, raddr sockaddr, mode string) (family int, ipv6only bool) {
	...

	if mode == "listen" && (laddr == nil || laddr.isWildcard()) {
		if supportsIPv4map() || !supportsIPv4() {
			return syscall.AF_INET6, false
		}
```

- また前述の lsof では TYPE カラムが `IPv6` となっていたこととも整合している

<div></div>

```
svr     16329 cou929    3u    IPv6 0xb7376df7b15138e1      0t0                 TCP *:2000 (LISTEN)
```

第二引数の `type` も同じ要領で見てみる。

- 今回は `0x1`
- これは `SOCK_STREAM` らしい

```c
#define SOCK_STREAM     1               /* stream socket */
```

なおヘッダファイルのインクルードパスがわからない場合は、適当な c のソースコードを gcc 等で `-v` コマンドをつけてコンパイルすると、次のような出力が得られ、あたりがつけられるらしい。

```
#include <...> search starts here:
/usr/local/include
/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/lib/clang/11.0.3/include
/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX.sdk/usr/include
/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/include
/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX.sdk/System/Library/Frameworks (framework directory)
```

例えば今回の `sys/socket.h` は以下の位置にあった。

```
/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX.sdk/usr/include/sys/socket.h
```

### クライアントから接続してみる

このプロセスにクライアントから接続してみる。`accept` が呼び出され、新たな確立済みのソケットが返されること、そこへの read / write が行われることが予想できる。

まずは接続。

```
2020-06-22 12:26 go master$ telnet 127.0.0.1 2000
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
```

予想通り `accept` が呼ばれている。先程のリスニングソケット `3` が第一引数に指定されていること、`7` が接続済みのソケットとして返されていることも確認できる。

```
29738/0x5d6ef:      6803      13     10 accept(0x3, 0xC00003AC14, 0xC00003ABF4)          = 7 0
```

また一度 read が行われ、それがエラーを返しているのもわかる。

```
29738/0x5d6ef:      7155       5      2 read(0x7, "\0", 0x8000)          = -1 Err#35
```

`sys/errno.h` によると 35 は `EAGAIN` らしい。

```c
#define EAGAIN          35              /* Resource temporarily unavailable */
```

Go 側の該当箇所は [internal/poll.FD.Accept](https://github.com/golang/go/blob/bd486c39bad4c9f90190ae58de8a592bb9a2aae9/src/internal/poll/fd_unix.go#L392-L397) だと思われる。

```go
for {
    s, rsa, errcall, err := accept(fd.Sysfd)
    if err == nil {
        return s, rsa, "", err
    }
    switch err {
    case syscall.EINTR:
        continue
    case syscall.EAGAIN:
        if fd.pd.pollable() {
            if err = fd.pd.waitRead(fd.isFile); err == nil {
                continue
            }
        }
```

### クライアントから書き込みしてみる

クライアントから文字列を送信すると同じ文字列がサーバから返される。

```
$ telnet 127.0.0.1 2000
Trying 127.0.0.1...
Connected to localhost.
...
test test test  // 送信内容
test test test  // 受信内容
```

次のようにファイルディスクリプタ `7` からの読み込みと、同内容の書き込みが確認できる。返り値の `16` は読み書きしたバイト数。書き込み後には先程と同様に EAGAIN で読み込みをブロッキングしながら待っている様子も見える。

```
29738/0x5d702:       252      17     12 read(0x7, "test test test\r\n\340AB%*t\0", 0x8000)               = 16 0
29738/0x5d702:       283      80     18 write(0x7, "test test test\r\n\0", 0x10)                 = 16 0
29738/0x5d702:       302       6      1 read(0x7, "\0", 0x8000)          = -1 Err#35
```

## gops

[gops](https://github.com/google/gops) は動作中の Go プロセスの診断や統計情報取得ができるツール。

今回の例題は accept でソケットを受け取るごとに、goroutine を起動してそのリクエストを処理している。この様子を見てみる。

### agent の設定

利用するには、次のようにアプリケーションに `gops` とやりとりする口を用意しないといけないらしい。

```go
package main

import (
	"io"
	"log"
	"net"

	"github.com/google/gops/agent"
)

func main() {
    // gops からのリクエストを待つ口を開ける。
	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen("tcp", ":2000")
    ...
```

gops の実行に必須ではないが、これがないと有益な情報 (そのプロセスの goroutine やスレッドの数や、スタックなど) が参照できないらしい。初めは外から取れる情報だけでうまくやるのかと期待していたが、流石にそんなことはなかった。

### 一通りの gops コマンド

`gops` でマシン上の Go プロセスが一覧される。

```
$ gops
75660 75638 go                 go1.14.2  /usr/local/Cellar/go/1.14.2_1/libexec/bin/go
126   1     com.docker.vmnetd  go1.12.16 /Library/PrivilegedHelperTools/com.docker.vmnetd
47708 2213  gops               go1.14.2  /Users/cou929/bin/gops
47670 1373  svr              * go1.14.2  /Users/cou929/src/github.com/cou929/smonw/svr
```

`gops <PID>` で以下のような統計情報が取得できる。lsof のように開いているソケットの情報も出してくれていた。port 2000 で待ち受けているのが例題の主題のエコーサーバで、58630 が gops のエージェント。

```
$ gops 47670
parent PID:     1373
threads:        6
memory usage:   0.030%
cpu usage:      0.000%
username:       cou929
cmd+args:       ./svr
elapsed time:   00:33
local/remote:   127.0.0.1:58630 <-> :0 (LISTEN)
local/remote:   *:2000 <-> :0 (LISTEN)
```

`gops stack <PID>` で goroutine ごとのスタックを参照できる。`goroutine 19` が gops のエージェント、`goroutine 1` が accept 待ちをしており、状態が `IO wait` 担っていることがわかる。

```
2020-06-22 12:52 go master$ gops stack 47670
goroutine 19 [running]:
runtime/pprof.writeGoroutineStacks(0x1190080, 0xc00010e000, 0x1000000000030, 0xd0)
        /usr/local/Cellar/go/1.14.2_1/libexec/src/runtime/pprof/pprof.go:665 +0x9d
runtime/pprof.writeGoroutine(0x1190080, 0xc00010e000, 0x2, 0x0, 0x0)
        /usr/local/Cellar/go/1.14.2_1/libexec/src/runtime/pprof/pprof.go:654 +0x44
runtime/pprof.(*Profile).WriteTo(0x1274c40, 0x1190080, 0xc00010e000, 0x2, 0xc00010e000, 0x0)
        /usr/local/Cellar/go/1.14.2_1/libexec/src/runtime/pprof/pprof.go:329 +0x3da
github.com/google/gops/agent.handle(0x2b00008, 0xc00010e000, 0xc00001c0c0, 0x1, 0x1, 0x0, 0x0)
        /Users/cou929/pkg/mod/github.com/google/gops@v0.3.10/agent/agent.go:189 +0x1af
github.com/google/gops/agent.listen()
        /Users/cou929/pkg/mod/github.com/google/gops@v0.3.10/agent/agent.go:133 +0x2bf
created by github.com/google/gops/agent.Listen
        /Users/cou929/pkg/mod/github.com/google/gops@v0.3.10/agent/agent.go:111 +0x36b

goroutine 1 [IO wait]:
internal/poll.runtime_pollWait(0x1789e38, 0x72, 0x0)
        /usr/local/Cellar/go/1.14.2_1/libexec/src/runtime/netpoll.go:203 +0x55
internal/poll.(*pollDesc).wait(0xc0000d2098, 0x72, 0x0, 0x0, 0x11630c3)
        /usr/local/Cellar/go/1.14.2_1/libexec/src/internal/poll/fd_poll_runtime.go:87 +0x45
internal/poll.(*pollDesc).waitRead(...)
        /usr/local/Cellar/go/1.14.2_1/libexec/src/internal/poll/fd_poll_runtime.go:92
internal/poll.(*FD).Accept(0xc0000d2080, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0)
        /usr/local/Cellar/go/1.14.2_1/libexec/src/internal/poll/fd_unix.go:384 +0x1d4
net.(*netFD).accept(0xc0000d2080, 0xc000040ed0, 0x1191a20, 0xc0000b2008)
        /usr/local/Cellar/go/1.14.2_1/libexec/src/net/fd_unix.go:238 +0x42
net.(*TCPListener).accept(0xc0000a8140, 0xc0000a8140, 0x0, 0x0)
        /usr/local/Cellar/go/1.14.2_1/libexec/src/net/tcpsock_posix.go:139 +0x32
net.(*TCPListener).Accept(0xc0000a8140, 0x3, 0x1162e04, 0x5, 0x1191720)
        /usr/local/Cellar/go/1.14.2_1/libexec/src/net/tcpsock.go:261 +0x64
main.main()
        /Users/cou929/src/github.com/cou929/smonw/main.go:23 +0x198
```

なお、前述の dlv で goroutine をみたときは、ランタイム側が用意したのか goroutine が 5 つ起動していた。gops ではユーザー側が起動した goroutine のみが表示され、ランタイム側のものは隠されているようだ。

`gops memstats <PID>` でメモリ使用状況の確認。

```
$ gops memstats 47670
alloc: 1.11MB (1161512 bytes)
total-alloc: 1.11MB (1161512 bytes)
sys: 69.45MB (72827136 bytes)
lookups: 0
mallocs: 397
frees: 6
heap-alloc: 1.11MB (1161512 bytes)
heap-sys: 63.69MB (66781184 bytes)
heap-idle: 62.18MB (65200128 bytes)
heap-in-use: 1.51MB (1581056 bytes)
heap-released: 62.12MB (65134592 bytes)
heap-objects: 391
stack-in-use: 320.00KB (327680 bytes)
stack-sys: 320.00KB (327680 bytes)
stack-mspan-inuse: 27.62KB (28288 bytes)
stack-mspan-sys: 32.00KB (32768 bytes)
stack-mcache-inuse: 6.78KB (6944 bytes)
stack-mcache-sys: 16.00KB (16384 bytes)
other-sys: 770.43KB (788916 bytes)
gc-sys: 3.28MB (3436808 bytes)
next-gc: when heap-alloc >= 4.27MB (4473924 bytes)
last-gc: -
gc-pause-total: 0s
gc-pause: 0
num-gc: 0
enable-gc: true
debug-gc: false
```

`gops stats <PID>` で goroutine やスレッド、cpu 数などを確認できる。

```
$ gops stats 47670
goroutines: 2
OS threads: 7
GOMAXPROCS: 4
num CPU: 4
```

### クライアントから接続してみる

telnet で接続すると、goroutine が起動し、accept 待ちをする様子が確認できる。

まずはクライアントから接続。

```
$ telnet 127.0.0.1 2000
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
```

goroutine の数が増えていることを確認。

```
$ gops stats 47670
goroutines: 3
OS threads: 7
GOMAXPROCS: 4
num CPU: 4
```

stack を見ると新しい `goroutine 34` が read 待ちをしているのが見える。

```
$ gops stack 47670
...
goroutine 34 [IO wait]:
internal/poll.runtime_pollWait(0x1789d58, 0x72, 0xffffffffffffffff)
        /usr/local/Cellar/go/1.14.2_1/libexec/src/runtime/netpoll.go:203 +0x55
internal/poll.(*pollDesc).wait(0xc00010a198, 0x72, 0x8000, 0x8000, 0xffffffffffffffff)
        /usr/local/Cellar/go/1.14.2_1/libexec/src/internal/poll/fd_poll_runtime.go:87 +0x45
internal/poll.(*pollDesc).waitRead(...)
        /usr/local/Cellar/go/1.14.2_1/libexec/src/internal/poll/fd_poll_runtime.go:92
internal/poll.(*FD).Read(0xc00010a180, 0xc000122000, 0x8000, 0x8000, 0x0, 0x0, 0x0)
        /usr/local/Cellar/go/1.14.2_1/libexec/src/internal/poll/fd_unix.go:169 +0x201
net.(*netFD).Read(0xc00010a180, 0xc000122000, 0x8000, 0x8000, 0xc000047d50, 0x104b4ec, 0x8000)
        /usr/local/Cellar/go/1.14.2_1/libexec/src/net/fd_unix.go:202 +0x4f
net.(*conn).Read(0xc00010e018, 0xc000122000, 0x8000, 0x8000, 0x0, 0x0, 0x0)
        /usr/local/Cellar/go/1.14.2_1/libexec/src/net/net.go:184 +0x8e
io.copyBuffer(0x11903c0, 0xc0000988c0, 0x2b04170, 0xc00010e018, 0xc000122000, 0x8000, 0x8000, 0x100a909, 0x127a320, 0x2b041b0)
        /usr/local/Cellar/go/1.14.2_1/libexec/src/io/io.go:405 +0x122
io.Copy(...)
        /usr/local/Cellar/go/1.14.2_1/libexec/src/io/io.go:364
net.genericReadFrom(0x1190080, 0xc00010e018, 0x2b04170, 0xc00010e018, 0x12a5898, 0x0, 0x0)
        /usr/local/Cellar/go/1.14.2_1/libexec/src/net/net.go:626 +0x9a
net.(*TCPConn).readFrom(0xc00010e018, 0x2b04170, 0xc00010e018, 0xc00002eed8, 0x100b66a, 0x1138a60)
        /usr/local/Cellar/go/1.14.2_1/libexec/src/net/tcpsock_posix.go:54 +0x4d
net.(*TCPConn).ReadFrom(0xc00010e018, 0x2b04170, 0xc00010e018, 0x2b041b0, 0xc00010e018, 0x2b04101)
        /usr/local/Cellar/go/1.14.2_1/libexec/src/net/tcpsock.go:103 +0x4d
io.copyBuffer(0x1190080, 0xc00010e018, 0x2b04170, 0xc00010e018, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0)
        /usr/local/Cellar/go/1.14.2_1/libexec/src/io/io.go:391 +0x2fc
io.Copy(...)
        /usr/local/Cellar/go/1.14.2_1/libexec/src/io/io.go:364
main.main.func1(0x11928a0, 0xc00010e018)
        /Users/cou929/src/github.com/cou929/smonw/main.go:28 +0xba
created by main.main
        /Users/cou929/src/github.com/cou929/smonw/main.go:27 +0x184
```

## 最後に

delve, lsof, dtruss, gops の基本的な使い方を確認した。静的なコードリーディングだけでなく、こうしたツールをつかった動的な検証も合わせて行うと、効率的にデバッグや動作理解を深めることができそうだった。

なおソケットやシステムコールそのものの知識は、以下の `UNIXネットワークプログラミング` という本が個人的にはおすすめ。分厚い本だが 2 章と 4 章をさっと読むだけで、この記事に出てくる内容には十分。へたに Web 上の記事などを読み漁るよりも、トータルのコスパは良くなると思う。

## 参考

- [go\-delve/delve: Delve is a debugger for the Go programming language\.](https://github.com/go-delve/delve)
- [google/gops: A tool to list and diagnose Go processes currently running on your system](https://github.com/google/gops)
- [dtrussでgolangのシステムコールをトレースしたい \- Qiita](https://qiita.com/yurakawa/items/2ab084a2f3f0aa0bf0c7)
- [Diagnostics \- The Go Programming Language](https://golang.org/doc/diagnostics.html)

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4894712059/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://images-na.ssl-images-amazon.com/images/I/213B9PVJD1L._BO1,204,203,200_.jpg" alt="UNIXネットワークプログラミング〈Vol.1〉ネットワークAPI:ソケットとXTI" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4894712059/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">UNIXネットワークプログラミング〈Vol.1〉ネットワークAPI:ソケットとXTI</a></div><div class="amazlet-detail">W.リチャード スティーヴンス (著), W.Richard Stevens (原著), 篠田 陽一 (翻訳)<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4894712059/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4908686033/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://images-na.ssl-images-amazon.com/images/I/51RKK5+6bpL._SX348_BO1,204,203,200_.jpg" alt="Goならわかるシステムプログラミング" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4908686033/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Goならわかるシステムプログラミング</a></div><div class="amazlet-detail">渋川 よしき  (著), ごっちん (イラスト)<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4908686033/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
