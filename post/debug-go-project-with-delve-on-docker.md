{"title":"Docker コンテナ上で Go のプロジェクトを delve でデバッグするには ptrace の許可が必要","date":"2019-01-05T18:43:09+09:00","tags":["nix"]}

Docker コンテナ上で Go のプロジェクトを [delve](https://github.com/derekparker/delve) でデバッグしたいとき。Docker はデフォルトで `ptrace(2)` システムコールの呼び出しを制限しているので、これを緩和する必要がある。

具体的には `docker run` に次のように `--cap-add=SYS_PTRACE` というオプションを渡してあげるとよい。

    docker container run --cap-add=SYS_PTRACE -it your-image:latest bash

### 背景

Docker コンテナ上で Go のプログラムを `delve` を使ってデバッグしようとすると次のようなエラーで動かなかった。

    root@a135f59c96cb:/go# dlv debug
    could not launch process: fork/exec /go/debug: operation not permitted

ググっているとまず見つけたのが [この issue](https://github.com/go-delve/delve/issues/515#issuecomment-214911481) で、

> Alright, so you're running within a Docker container. Docker has security settings preventing ptrace(2) operations by default with in the container. Pass --security-opt=seccomp:unconfined to docker run when starting.

というコメントがあった。これで解決はしたものの意味がよくわからなかったので、調べたことをまとめたのが今回の内容です。

以降は調べたことのメモ。

### `ptrace(2)` とは

システムコール。
プロセスをアタッチしたり、そのプロセスでのシステムコールの発生をまったり、その時のレジスタの値を取得できる。
これを使って実装されたのが `strace` や `gdb`。

### Docker のセキュリティと `Linux capabilities` について

[Docker security | Docker Documentation](https://docs.docker.com/engine/security/security/#linux-kernel-capabilities) に Docker のセキュリティ保全のために使われている技術の概要が説明されている。Linux の namespace や cgroups と並んで、`capabilities` という機能も使われているとのこと。

[capabilities](https://linuxjm.osdn.jp/html/LDP_man-pages/man7/capabilities.7.html) は、スーパーユーザーなら使える各種の権限を細かくオンオフできるもののようだ。

このなかに ptrace を扱う `CAP_SYS_PTRACE` というものもある。

    CAP_SYS_PTRACE
            Trace arbitrary processes using ptrace(2); apply get_robust_list(2) to arbitrary processes; inspect processes using kcmp(2).

Docker の話に戻ると、

> This means that in most cases, containers do not need “real” root privileges at all. And therefore, containers can run with a reduced capability set; meaning that “root” within a container has much less privileges than the real “root”.

> Docker supports the addition and removal of capabilities, allowing use of a non-default profile. This may make Docker more secure through capability removal, or less secure through the addition of capabilities. The best practice for users would be to remove all capabilities except those explicitly required for their processes.

とのことで、

- コンテナ内の `root` は実は特定の権限だけが許可されたユーザー
- また `capabilities` の付け外しはオプションでできるようになっている

らしい。

デフォルトで許可されている `capabilities` は [moby/defaults\.go at master · moby/moby](https://github.com/moby/moby/blob/master/oci/defaults.go#L14-L30) によると以下で、`CAP_SYS_PTRACE` は入っていない。

    "CAP_CHOWN",
    "CAP_DAC_OVERRIDE",
    "CAP_FSETID",
    "CAP_FOWNER",
    "CAP_MKNOD",
    "CAP_NET_RAW",
    "CAP_SETGID",
    "CAP_SETUID",
    "CAP_SETFCAP",
    "CAP_SETPCAP",
    "CAP_NET_BIND_SERVICE",
    "CAP_SYS_CHROOT",
    "CAP_KILL",
    "CAP_AUDIT_WRITE",

[Runtime privilege and Linux capabilities](https://docs.docker.com/engine/reference/run/#runtime-privilege-and-linux-capabilities) によると、`docker run` のオプションに `--cap-add` というオプションがあり、こちらに `SYS_PTRACE` を渡すと良さそうだ。

    --cap-add=SYS_PTRACE

ちなみに `--cap-add=ALL` ですべての `capabilities` を付与したり、

> By default, Docker containers are “unprivileged” and cannot, for example, run a Docker daemon inside a Docker container. This is because by default a container is not allowed to access any devices, but a “privileged” container is given access to all devices (see the documentation on cgroups devices).

> When the operator executes docker run --privileged, Docker will enable access to all devices on the host as well as set some configuration in AppArmor or SELinux to allow the container nearly all the same access to the host as processes running outside containers on the host. Additional information about running with --privileged is available on the Docker Blog.

のように `--priviledged` オプションでまるっと権限を与えることもできるそう。

### `seccomp` について

先程の [Docker security | Docker Documentation](https://docs.docker.com/engine/security/security/#linux-kernel-capabilities) にあったように、Docker のセキュリティは Linux の各種仕組みを組み合わせて担保されている。そのひとつとして、`capabilities` とは別に `seccomop` という仕組みも利用されている。

[Seccomp security profiles for Docker | Docker Documentation](https://docs.docker.com/engine/security/seccomp/#run-without-the-default-seccomp-profile) によると、`seccomp` とは Linux kernel の機能で、呼び出すことができるシステムコールの制限ができるらしい。もともと Linux 上のサンドボックス環境の構築のために導入されたもので、Chrome や Android, OpenSSH なども利用しているようだ。

> The default seccomp profile provides a sane default for running containers with seccomp and disables around 44 system calls out of 300+. It is moderately protective while providing wide application compatibility. The default Docker profile can be found here).

> ptrace
> 
> Tracing/profiling syscall, which could leak a lot of information on the host. Already blocked by dropping CAP_PTRACE.

とのことで、デフォルトでは `ptrace` は制限されている。

`seccomp` でどのようにシステムコール呼び出しを制御するかは `seccomp profile` という設定ファイルで定義できるらしい。`seccomp profile` は `seccomp 2.2.1` 以上で使えるもののようで、それは `Ubuntu 14.04, Debian Wheezy, or Debian Jessie` には無いとのこと。Docker のデフォルトの `seccomp profile` の内容はこちら。

[moby/default\.json at master · moby/moby](https://github.com/moby/moby/blob/master/profiles/seccomp/default.json)

`seccomop` も `capabilities` 同様にオプションで制御できる。

`--security-opt seccomp=unconfined` というオプションを `docker run` に渡してあげると、デフォルトの `seccomp profile` を使わずに起動できるようだ。

    $ docker run --rm -it --security-opt seccomp=unconfined debian:jessie \
        unshare --map-root-user --user sh -c whoami

あるいは特定のシステムコール (今回は `ptrace`) だけを緩和するような `seccomop profile` を書いてそれを `docker run` 時に渡すこともできるらしい。

Docker を利用する際の `cpabilities` と `seccomp` の設定の仕方については、以下にあるように、

> The default seccomp profile will adjust to the selected capabilities, in order to allow use of facilities allowed by the capabilities, so you should not have to adjust this, since Docker 1.12. In Docker 1.10 and 1.11 this did not happen and it may be necessary to use a custom seccomp profile or use --security-opt seccomp=unconfined when adding capabilities.

`capabilities` の設定に応じて `seccomp` のパラメータも自動で調整してくれるらしい。Docker 1.12 以降はこうなっているとのことで、だいぶ前から対応しているようだ。

以上より、`docker run --cap-add=SYS_PTRACE` というふうにコンテナを起動すれば、`seccomp` は気にしなくてもよさそうだ。

### docker-compose を使う場合

docker-compose の場合は `docker-compose.yml` に同様の内容を記載すれば良い。

    version: 3

    services:
        your_service:
            # ...
            cap_add:
                - SYS_PTRACE

ちなみに `--security-opt` は以下のように指定できる。

    security_opt:
        - seccomp:unconfined

- [Compose file version 3 reference | Docker Documentation](https://docs.docker.com/compose/compose-file/#cap_add-cap_drop)


### 実験してみる

次のような Dockerfile と Go のプログラムを準備して `strace` や `dlv debug, attach` をやってみる。

Dockerfile の内容

    FROM golang:latest

    RUN apt-get update \
        && apt-get install -y strace

    RUN go get github.com/derekparker/delve/cmd/dlv

    ADD ./test.go /go/test.go

test.go の内容

    package main

    import (
        "fmt"
        "time"
    )

    func main() {
        time.Sleep(5 * time.Minute)
        fmt.Println("hello")
    }

#### まずは普通に起動する

    docker image build -t sec-test .
    docker container run -it sec-test:latest bash

`golang` コンテナの環境を確認。

    root@7128fc5edb57:/go# cat /etc/os-release
    PRETTY_NAME="Debian GNU/Linux 9 (stretch)"
    NAME="Debian GNU/Linux"
    VERSION_ID="9"
    VERSION="9 (stretch)"
    ID=debian
    HOME_URL="https://www.debian.org/"
    SUPPORT_URL="https://www.debian.org/support"
    BUG_REPORT_URL="https://bugs.debian.org/"

    root@7128fc5edb57:/go# cat /etc/debian_version
    9.6

ここは予想通り、`strace` も `dlv debug` もできないことを確認できた。

    # strace 実験

    root@a135f59c96cb:/go# ps aux
    USER       PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
    root         1  0.1  0.1  18188  3328 pts/0    Ss   05:14   0:00 bash
    root       327  0.0  0.1  36636  2792 pts/0    R+   05:14   0:00 ps aux

    root@a135f59c96cb:/go# strace -p 1
    strace: attach: ptrace(PTRACE_ATTACH, 1): Operation not permitted

    # dlv 実験

    root@a135f59c96cb:/go# ls
    bin  src  test.go

    root@a135f59c96cb:/go# dlv debug
    could not launch process: fork/exec /go/debug: operation not permitted

#### `--cap-add=SYS_PTRACE` を試してみる

次のように `--cap-add` オプションを渡して起動する。

    docker container run --cap-add=SYS_PTRACE -it sec-test:latest bash

`starce` で attach できた。

    root@edcbb46771f8:/go# ps aux
    USER       PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
    root         1  0.6  0.1  18188  3228 pts/0    Ss   05:16   0:00 bash
    root         6  0.0  0.1  36636  2780 pts/0    R+   05:16   0:00 ps aux
    root@edcbb46771f8:/go# strace -p 1
    strace: Process 1 attached
    wait4(-1,

`dlv debug` もできた。

    root@edcbb46771f8:/go# ls
    bin  src  test.go
    root@edcbb46771f8:/go# dlv debug
    Type 'help' for list of commands.
    (dlv) b main.main
    Breakpoint 1 set at 0x49c5ff for main.main() ./test.go:5
    (dlv) c
    > main.main() ./test.go:5 (hits goroutine(1):1 total:1) (PC: 0x49c5ff)
        1: package main
        2:
        3: import "fmt"
        4:
    =>   5: func main() {
        6:         fmt.Println("hello")
        7: }
    (dlv) n
    > main.main() ./test.go:6 (PC: 0x49c60d)
        1: package main
        2:
        3: import "fmt"
        4:
        5: func main() {
    =>   6:         fmt.Println("hello")
        7: }
    (dlv) locals
    (no locals)
    (dlv) c
    hello
    Process 325 has exited with status 0
    (dlv) q
    Process 325 has exited with status 0

`dlv attach` もできた

    root@093ee4748493:/go# go run test.go &
    [3] 135
    root@093ee4748493:/go# dlv attach 135
    Type 'help' for list of commands.
    (dlv)


#### `--security-opt=seccomp=unconfined` を試してみる

次のように `--security-opt` オプションを渡して起動する。

    docker container run --security-opt=seccomp=unconfined -it sec-test:latest bash

`capabilities` のほうで制限されたままなので、`strace` はできない。

    root@a1fe21c3f80b:/go# strace -p 1
    strace: attach: ptrace(PTRACE_SEIZE, 1): Operation not permitted

予想に反して `dlv debug` はできた。

    root@a1fe21c3f80b:/go# dlv debug
    Type 'help' for list of commands.
    (dlv) b main.main
    Breakpoint 1 set at 0x49c5ff for main.main() ./test.go:5
    (dlv) c
    > main.main() ./test.go:5 (hits goroutine(1):1 total:1) (PC: 0x49c5ff)
        1: package main
        2:
        3: import "fmt"
        4:
    =>   5: func main() {
        6:         fmt.Println("hello")
        7: }
    (dlv) n
    > main.main() ./test.go:6 (PC: 0x49c60d)
        1: package main
        2:
        3: import "fmt"
        4:
        5: func main() {
    =>   6:         fmt.Println("hello")
        7: }
    (dlv) c
    hello
    Process 328 has exited with status 0
    (dlv) q
    Process 328 has exited with status 0

`dlv attach` はできなかった。

    root@3f3f7098a69d:/go# go run test.go &
    [1] 90
    root@3f3f7098a69d:/go# ps aux
    USER       PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
    root         1  0.0  0.1  18188  3340 pts/0    Ss   05:26   0:00 bash
    root        90  4.3  0.8 315580 18200 pts/0    Sl   05:27   0:00 go run test.go
    root       124  0.0  0.0 102656  1352 pts/0    Sl   05:27   0:00 /tmp/go-build823881696/b001/exe/test
    root       128  0.0  0.1  36636  2796 pts/0    R+   05:27   0:00 ps aux
    root@3f3f7098a69d:/go# dlv attach 90
    Could not attach to pid 90: this could be caused by a kernel security setting, try writing "0" to /proc/sys/kernel/yama/ptrace_scope

というわけで、`capabilities の CAP_SYS_PTRACE` と `seccomp で制御している ptarace システムコール` とは完全に等しいわけではなさそう。あるいは `delve` の `debug` と `attach` で必要な権限が違うのかもしれない。ここから先はもっと掘る必要があるけれど、今回はしらべていない。

### まとめ

- Docker はデフォルトでいろいろな操作を制限している。
- `ptrace(2)` も制限されているので、`strace` やデバッガを使う場合は `--cap-add=SYS_PTRACE` オプションを `docker run` に渡して許可してあげる必要がある。
- `dlv debug` に関しては `--security-opt seccomp=unconfined` でも動作したが、細かい原因は今回は調べなかった。

### 参考

-  [ptraceシステムコール入門 ― プロセスの出力を覗き見してみよう！ \- プログラムモグモグ](https://itchyny.hatenablog.com/entry/2017/07/31/090000)
- [Docker security | Docker Documentation](https://docs.docker.com/engine/security/security/#linux-kernel-capabilities)
- [Docker run reference | Docker Documentation](https://docs.docker.com/engine/reference/run/)
- [明日使えない Linux の capabilities の話 \- @nojima's blog](https://nojima.hatenablog.com/entry/2016/12/03/000000)
- [Ubuntu, could not launch process: fork/exec \./debug: operation not permitted · Issue \#515 · go\-delve/delve](https://github.com/go-delve/delve/issues/515)
- [Seccomp security profiles for Docker | Docker Documentation](https://docs.docker.com/engine/security/seccomp/#run-without-the-default-seccomp-profile)
- [seccomp \- Wikipedia](https://en.wikipedia.org/wiki/Seccomp)
- [LinuxのBPF : \(2\) seccompでの利用 \- 睡分不足](http://mmi.hatenablog.com/entry/2016/08/01/044000)
- [Compose file version 3 reference | Docker Documentation](https://docs.docker.com/compose/compose-file/)
