{"title":"イラストでわかる Docker と Kubernetes を読んだ","date":"2020-12-30T23:57:00+09:00"}

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B08PNMRXKN/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/41g+F7WohJL.jpg" alt="イラストでわかるDockerとKubernetes Software Design plus" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B08PNMRXKN/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">イラストでわかるDockerとKubernetes Software Design plus</a></div><div class="amazlet-detail">徳永 航平  (著)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B08PNMRXKN/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

Docker, Kubanetes のアーキテクチャがざっくり把握できて良かった。分量も 130p くらいですぐ読めるので、はじめの一冊としてよさそう。

前提として自分のステータスはこんな感じで、その自分が概観を掴むのにちょうどよかった。

- Docker など関連技術はユーザーとしてなんとなく使っているだけで、内部構造や最新動向は全然追っていない
- k8s の本番環境での利用経験はなし
- 先日の [kubernetes が Dockershim を Depricated にする話](https://kubernetes.io/blog/2020/12/02/dont-panic-kubernetes-and-docker/) は、解説ブログをよんでなんとなくしかわかっていない
    - この本を読んで、ざっくりだがどの部分に影響がある話なのかわかった気になれた

`第2章 Dockerの概要` と `第4章 コンテナランタイムとコンテナの標準仕様の概要` は特にわかりやすく理解が進んだ。一方で `第3章 Kubernetesの概要` はアーキテクチャを理解するにも使い方を理解するにも若干中途半端な感があり、チュートリアルや類書で補完したほうがよさそうだった。

以下いろいろと試しながら読んだ際のメモ。

## イメージのレイヤ構造

これまでユーザーとして Docker を利用している中で、Docker image がレイヤー構造になっていることは、次のような挙動や仕様からなんとなくイメージできていた。

- Docker image をビルドするとき、Dockerfile の一行ごとにキャッシュされている
- Base image に変更を上乗せして Docker image を作るという構成

この構造を、実際の image の中身を見ながらより具体的に理解できた。

[`docker save`](https://docs.docker.com/engine/reference/commandline/save/) で指定した image の内容を出力できる。

```
docker save myimage:v1 | tar -xC ./dumpimage
```

(`myimage:v1` は本書のサンプル image と同じ内容で、`ubuntu:20.04` に `hello.sh` というファイルを追加したというもの)

つぎのように設定ファイルやファイルそのものが確認できる。

```
$ cat dumpimage/manifest.json | jq .
[
  {
    "Config": "99561f028873afc6a69163adcdc6fe08c6dfca5f6d2de448c14ff3094cafec1a.json",
    "RepoTags": [
      "myimage:v1"
    ],
    "Layers": [
      "4fc2b30c28ffa7c7fc054e025ae58d7e9b147044f83cdb713e2772d284de6152/layer.tar",
      "a8b100bdf94d6b2883b6cca2063de4991311320ac3a20cf7b2144fbb6d9d50d5/layer.tar",
      "3f415c4e8f619fa8d4651a9f77e2b9ba0ec893c398bf7bcacf8e1ab89b7cfcfc/layer.tar",
      "d3cde48172a0bc804b673d925fe9ac291635a10ca6f6bc38b8202c5e59bcf6f2/layer.tar"
    ]
  }
]
```

各 layer.tar がそれぞれのレイヤのファイルに相当する。一番上層には Dockerfile で追加した `hello.sh` が見える。

```
$ tar --list -f dumpimage/4fc2b30c28ffa7c7fc054e025ae58d7e9b147044f83cdb713e2772d284de6152/layer.tar | head
bin
boot/
dev/
etc/
etc/.pwd.lock
etc/adduser.conf
etc/alternatives/
etc/alternatives/README
etc/alternatives/awk
etc/alternatives/nawk

$ tar --list -f dumpimage/a8b100bdf94d6b2883b6cca2063de4991311320ac3a20cf7b2144fbb6d9d50d5/layer.tar | head
etc/
etc/apt/
etc/apt/apt.conf.d/
etc/apt/apt.conf.d/docker-autoremove-suggests
etc/apt/apt.conf.d/docker-clean
etc/apt/apt.conf.d/docker-gzip-indexes
etc/apt/apt.conf.d/docker-no-languages
etc/dpkg/
etc/dpkg/dpkg.cfg.d/
etc/dpkg/dpkg.cfg.d/docker-apt-speedup

$ tar --list -f dumpimage/3f415c4e8f619fa8d4651a9f77e2b9ba0ec893c398bf7bcacf8e1ab89b7cfcfc/layer.tar | head
run/
run/systemd/
run/systemd/container

$ tar --list -f dumpimage/d3cde48172a0bc804b673d925fe9ac291635a10ca6f6bc38b8202c5e59bcf6f2/layer.tar | head
hello.sh
```

`docker save` で出力される形式は Docker Image Specification V1 というもので、OCI (コミュニティの標準) の仕様ではないらしい。Docker Image Specification の [README には `docker save, load` 以外では使われていないよという記載](https://github.com/moby/moby/tree/master/image/spec#docker-image-specification-v1) がある。

[containers/skopeo](https://github.com/containers/skopeo) というツールでこれを [OCI 標準形式](https://github.com/opencontainers/image-spec) に変換できるらしい。

```
$ skopeo copy docker-daemon:myimage:v1 oci:dumpociimage
$ ls dumpociimage/
blobs/      index.json  oci-layout
```

フォーマットこそ違うが、ファイルの構成がレイヤ構造になっていて、概念的には似たようなものだと思って大丈夫そう。

```
$ cat dumpociimage/index.json | jq .
{
  "schemaVersion": 2,
  "manifests": [
    {
      "mediaType": "application/vnd.oci.image.manifest.v1+json",
      "digest": "sha256:1bed7c632aa802773a06589531baf96f91f0e441822479ff9dd8dceec3921c8d",
      "size": 812
    }
  ]
}

$ cat dumpociimage/blobs/sha256/1bed7c632aa802773a06589531baf96f91f0e441822479ff9dd8dceec3921c8d | jq .
{
  "schemaVersion": 2,
  "config": {
    "mediaType": "application/vnd.oci.image.config.v1+json",
    "digest": "sha256:79977be18007ca282b6a4886744f044f03a76beadb526168d507896c5a5c7cde",
    "size": 2800
  },
  "layers": [
    {
      "mediaType": "application/vnd.oci.image.layer.v1.tar+gzip",
      "digest": "sha256:f6291d8887317896606687305f6661ecfd884bc4e4523d948c39af71d3cf3a5d",
      "size": 29977988
    },
    {
      "mediaType": "application/vnd.oci.image.layer.v1.tar+gzip",
      "digest": "sha256:9e0775ca9a2f3180f9cd3ca1532446a8d6dc58f4970ab9babd6168fc117ddee7",
      "size": 925
    },
    {
      "mediaType": "application/vnd.oci.image.layer.v1.tar+gzip",
      "digest": "sha256:247a9afb7564837564ec955789095de51dec696556c9707b3a816919b54190be",
      "size": 173
    },
    {
      "mediaType": "application/vnd.oci.image.layer.v1.tar+gzip",
      "digest": "sha256:f8c4c898ed357f87c4e3da939664a4d84d460ee00a27000cfdc0d2a8923d0a47",
      "size": 141
    }
  ]
}
```

Docker Image Spec 形式同様に、上層に `hello.sh` が登場していた。

```
$ tar --list -f dumpociimage/blobs/sha256/f6291d8887317896606687305f6661ecfd884bc4e4523d948c39af71d3cf3a5d | head
bin
boot/
dev/
etc/
etc/.pwd.lock
etc/adduser.conf
etc/alternatives/
etc/alternatives/README
etc/alternatives/awk
etc/alternatives/nawk

$ tar --list -f dumpociimage/blobs/sha256/f8c4c898ed357f87c4e3da939664a4d84d460ee00a27000cfdc0d2a8923d0a47 | head
hello.sh
```

## overlay2 Storage Driver

Docker image の構成の概要がわかったところで、それがどう実現されているのかに話が進む。

Docker は各モジュールがプラガブルな設計になっており、こうした image のファイル構造の管理などを担当しているのが [Storage Driver](https://docs.docker.com/storage/storagedriver/select-storage-driver/) らしい。現在推奨されている Storage Driver が [overlay2](https://docs.docker.com/storage/storagedriver/overlayfs-driver/) でその説明がされていた。

`overlay2` は [OverlayFS](https://en.wikipedia.org/wiki/OverlayFS) というファイルシステムのうえに構築されている。OverlayFS は [`union mount filesystem`](https://en.wikipedia.org/wiki/Union_mount) という、複数のディレクトリをひとつにマージして見せるファイルシステムに分類され、すでにカーネルにもマージ済みらしい。

レイヤ構造になっている各ディレクトリについて、下層は Read Only になっている。下層のファイルを編集する際は、それをコピーして上層に配置する (つまり Copy On Write)。Docker では同じ Docker image から複数のコンテナを立ち上げた際、Read only な低層は全くおなじなので、コンテナ間で共有しているらしい。

現在の Storage Driver は [`docker info`](https://docs.docker.com/engine/reference/commandline/info/) で確認できる。

```
$ docker info | grep -A 3 'Storage'
 Storage Driver: overlay2
  Backing Filesystem: extfs
  Supports d_type: true
  Native Overlay Diff: true
```

コンテナの詳細情報は [`docker inspect`](https://docs.docker.com/engine/reference/commandline/inspect/) で確認できる。ここから実際のファイルパスがわかる。

```
$ docker inspect mycontainer | jq '.[0].GraphDriver'
{
  "Data": {
    "LowerDir": "/var/lib/docker/overlay2/26561d1a763d94798e0e652c8a6f760f39b04e802b580f7812e3e35d3627382a-init/diff:/var/lib/docker/overlay2/da9d83e6f12122882e7fbc2e7e53b1b10b8201a4f34f60f62ae1dc3fa4a227b2/diff:/var/lib/docker/overlay2/59f72c61e5e742c2c0ce1afd5da4f67598ef2da6a6940985fdf36fff75e28b22/diff:/var/lib/docker/overlay2/951f80a22d8235c982c36c43b2987da55cefa5cf76c54aebd93b5d664f47a6c8/diff:/var/lib/docker/overlay2/e49a7c6f6ecec3ae0ded8decba94ef26682b111b8741a1bd1b7af9960dc756d6/diff",
    "MergedDir": "/var/lib/docker/overlay2/26561d1a763d94798e0e652c8a6f760f39b04e802b580f7812e3e35d3627382a/merged",
    "UpperDir": "/var/lib/docker/overlay2/26561d1a763d94798e0e652c8a6f760f39b04e802b580f7812e3e35d3627382a/diff",
    "WorkDir": "/var/lib/docker/overlay2/26561d1a763d94798e0e652c8a6f760f39b04e802b580f7812e3e35d3627382a/work"
  },
  "Name": "overlay2"
}
```

実際に `/var/lib/docker/overlay2/` 以下を見に行ってみる。

```
$ docker run -it --privileged --pid=host debian nsenter -t 1 -m -u -n -i sh
```

```
# みやすさのために改行をいれた
# docker inspect の内容と一致している

/ # mount | grep 26561d1a763d94798e0e652c8a6f760f39b04e802b580f7812e3e35d3627382a
overlay on /var/lib/docker/overlay2/26561d1a763d94798e0e652c8a6f760f39b04e802b580f7812e3e35d3627382a/merged type overlay (rw,relatime,
lowerdir=/var/lib/docker/overlay2/l/OQKBCJO6UQAT6ZP72CGDZGVBYP:/var/lib/docker/overlay2/l/EBZIYV75XFPERCXHX23BU2REPN:/var/lib/docker/overlay2/l/AFU74I2HKXON542KA65LFHDEMD:/var/lib/docker/overlay2/l/QGDWNWNNUACYYU4CNFZYDGZLUH:/var/lib/docker/overlay2/l/STX4TR3EDB7P72N4Z6D6KDJZSY,
upperdir=/var/lib/docker/overlay2/26561d1a763d94798e0e652c8a6f760f39b04e802b580f7812e3e35d3627382a/diff,
workdir=/var/lib/docker/overlay2/26561d1a763d94798e0e652c8a6f760f39b04e802b580f7812e3e35d3627382a/work)
```

lowerdir をいくつか見てみる。ちなみに [コマンドライン引数の長さ制限を回避するため](https://docs.docker.com/storage/storagedriver/overlayfs-driver/#image-and-container-layers-on-disk) に `overlay2` は各レイヤーに対して `STX4TR3EDB7P72N4Z6D6KDJZSY` のような短縮表現を用意しているらしい。

```
/ # ls -l /var/lib/docker/overlay2/l/STX4TR3EDB7P72N4Z6D6KDJZSY/
total 60
lrwxrwxrwx    1 root     root             7 Nov  6 01:21 bin -> usr/bin
drwxr-xr-x    2 root     root          4096 Apr 15  2020 boot
drwxr-xr-x    2 root     root          4096 Nov  6 01:25 dev
drwxr-xr-x   30 root     root          4096 Nov  6 01:25 etc
drwxr-xr-x    2 root     root          4096 Apr 15  2020 home
lrwxrwxrwx    1 root     root             7 Nov  6 01:21 lib -> usr/lib
lrwxrwxrwx    1 root     root             9 Nov  6 01:21 lib32 -> usr/lib32
lrwxrwxrwx    1 root     root             9 Nov  6 01:21 lib64 -> usr/lib64
lrwxrwxrwx    1 root     root            10 Nov  6 01:21 libx32 -> usr/libx32
drwxr-xr-x    2 root     root          4096 Nov  6 01:21 media
drwxr-xr-x    2 root     root          4096 Nov  6 01:21 mnt
drwxr-xr-x    2 root     root          4096 Nov  6 01:21 opt
drwxr-xr-x    2 root     root          4096 Apr 15  2020 proc
drwx------    2 root     root          4096 Nov  6 01:25 root
drwxr-xr-x    4 root     root          4096 Nov  6 01:21 run
lrwxrwxrwx    1 root     root             8 Nov  6 01:21 sbin -> usr/sbin
drwxr-xr-x    2 root     root          4096 Nov  6 01:21 srv
drwxr-xr-x    2 root     root          4096 Apr 15  2020 sys
drwxrwxrwt    2 root     root          4096 Nov  6 01:25 tmp
drwxr-xr-x   13 root     root          4096 Nov  6 01:21 usr
drwxr-xr-x   11 root     root          4096 Nov  6 01:25 var
/ # ls -l /var/lib/docker/overlay2/l/STX4TR3EDB7P72N4Z6D6KDJZSY/..
total 8
-rw-------    1 root     root             0 Dec 29 04:27 committed
drwxr-xr-x   17 root     root          4096 Dec 29 04:27 diff
-rw-r--r--    1 root     root            26 Dec 29 04:27 link

# hello.sh も登場
/ # ls -l /var/lib/docker/overlay2/l/EBZIYV75XFPERCXHX23BU2REPN/
total 4
-rwxr-xr-x    1 root     root            36 Dec 29 04:23 hello.sh
```

upperdir を見てみる。該当のコンテナには `docker exec` で bash を起動していたため、すでに `.bash_history` ができていた。

```
/ # ls -la /var/lib/docker/overlay2/26561d1a763d94798e0e652c8a6f760f39b04e802b580f7812e3e35d3627382a/diff/*
total 12
drwx------    2 root     root          4096 Dec 30 05:59 .
drwxr-xr-x    3 root     root          4096 Dec 29 04:29 ..
-rw-------    1 root     root            26 Dec 30 05:59 .bash_history
```

workdir はからっぽ。

```
/ # ls -la /var/lib/docker/overlay2/26561d1a763d94798e0e652c8a6f760f39b04e802b580f7812e3e35d3627382a/work/
total 12
drwx------    3 root     root          4096 Dec 30 00:27 .
drwx------    5 root     root          4096 Dec 30 00:27 ..
d---------    2 root     root          4096 Dec 30 05:59 work
/ # ls -la /var/lib/docker/overlay2/26561d1a763d94798e0e652c8a6f760f39b04e802b580f7812e3e35d3627382a/work/*
total 8
d---------    2 root     root          4096 Dec 30 05:59 .
drwx------    3 root     root          4096 Dec 30 00:27 ..
```

merged を確認。`hello.sh` もちゃんとあった。

```
/ # ls -la /var/lib/docker/overlay2/26561d1a763d94798e0e652c8a6f760f39b04e802b580f7812e3e35d3627382a/merged/
total 72
drwxr-xr-x    1 root     root          4096 Dec 29 04:29 .
drwx------    5 root     root          4096 Dec 30 00:27 ..
-rwxr-xr-x    1 root     root             0 Dec 29 04:29 .dockerenv
lrwxrwxrwx    1 root     root             7 Nov  6 01:21 bin -> usr/bin
drwxr-xr-x    2 root     root          4096 Apr 15  2020 boot
drwxr-xr-x    1 root     root          4096 Dec 29 04:29 dev
drwxr-xr-x    1 root     root          4096 Dec 29 04:29 etc
-rwxr-xr-x    1 root     root            36 Dec 29 04:23 hello.sh
drwxr-xr-x    2 root     root          4096 Apr 15  2020 home
lrwxrwxrwx    1 root     root             7 Nov  6 01:21 lib -> usr/lib
lrwxrwxrwx    1 root     root             9 Nov  6 01:21 lib32 -> usr/lib32
lrwxrwxrwx    1 root     root             9 Nov  6 01:21 lib64 -> usr/lib64
lrwxrwxrwx    1 root     root            10 Nov  6 01:21 libx32 -> usr/libx32
drwxr-xr-x    2 root     root          4096 Nov  6 01:21 media
drwxr-xr-x    2 root     root          4096 Nov  6 01:21 mnt
drwxr-xr-x    2 root     root          4096 Nov  6 01:21 opt
drwxr-xr-x    2 root     root          4096 Apr 15  2020 proc
drwx------    1 root     root          4096 Dec 30 05:59 root
drwxr-xr-x    1 root     root          4096 Nov 25 22:25 run
lrwxrwxrwx    1 root     root             8 Nov  6 01:21 sbin -> usr/sbin
drwxr-xr-x    2 root     root          4096 Nov  6 01:21 srv
drwxr-xr-x    2 root     root          4096 Apr 15  2020 sys
drwxrwxrwt    2 root     root          4096 Nov  6 01:25 tmp
drwxr-xr-x    1 root     root          4096 Nov  6 01:21 usr
drwxr-xr-x    1 root     root          4096 Nov  6 01:25 var
```

ファイル作成してみる。

```$
$ docker exec -it mycontainer bash
root@d447742a26fe:/# touch /tmp/foo
```

upperdir に追加され、merged にも反映されていた。

```
# upperdir
/ # ls -la /var/lib/docker/overlay2/26561d1a763d94798e0e652c8a6f760f39b04e802b580f7812e3e35d3627382a/diff/*
/var/lib/docker/overlay2/26561d1a763d94798e0e652c8a6f760f39b04e802b580f7812e3e35d3627382a/diff/root:
total 12
drwx------    2 root     root          4096 Dec 30 05:59 .
drwxr-xr-x    4 root     root          4096 Dec 29 04:29 ..
-rw-------    1 root     root            26 Dec 30 05:59 .bash_history

/var/lib/docker/overlay2/26561d1a763d94798e0e652c8a6f760f39b04e802b580f7812e3e35d3627382a/diff/tmp:
total 8
drwxrwxrwt    2 root     root          4096 Dec 30 06:32 .
drwxr-xr-x    4 root     root          4096 Dec 29 04:29 ..
-rw-r--r--    1 root     root             0 Dec 30 06:32 foo

# merged にも反映
/ # ls -la /var/lib/docker/overlay2/26561d1a763d94798e0e652c8a6f760f39b04e802b580f7812e3e35d3627382a/merged/tmp/
total 8
drwxrwxrwt    1 root     root          4096 Dec 30 06:32 .
drwxr-xr-x    1 root     root          4096 Dec 29 04:29 ..
-rw-r--r--    1 root     root             0 Dec 30 06:32 foo
```

docker stop してみる。

```
$ docker stop mycontainer

# merged がなくなっている
/ # ls -la /var/lib/docker/overlay2/26561d1a763d94798e0e652c8a6f760f39b04e802b580f7812e3e35d3627382a/
total 40
drwx------    4 root     root          4096 Dec 30 06:34 .
drwx------  187 root     root         20480 Dec 30 00:40 ..
drwxr-xr-x    4 root     root          4096 Dec 29 04:29 diff
-rw-r--r--    1 root     root            26 Dec 29 04:29 link
-rw-r--r--    1 root     root           144 Dec 29 04:29 lower
drwx------    3 root     root          4096 Dec 30 00:27 work

# upper には反映済み
/ # ls -la /var/lib/docker/overlay2/26561d1a763d94798e0e652c8a6f760f39b04e802b580f7812e3e35d3627382a/diff/tmp/
total 8
drwxrwxrwt    2 root     root          4096 Dec 30 06:32 .
drwxr-xr-x    4 root     root          4096 Dec 29 04:29 ..
-rw-r--r--    1 root     root             0 Dec 30 06:32 foo
```

docker rm してみる。

```
$ docker rm mycontainer

# なくなっている
/ # ls -la /var/lib/docker/overlay2/26561d1a763d94798e0e652c8a6f760f39b04e802b580f7812e3e35d3627382a/
ls: /var/lib/docker/overlay2/26561d1a763d94798e0e652c8a6f760f39b04e802b580f7812e3e35d3627382a/: No such file or directory
```

## Container runtime

コンテナを動作、管理するランタイムには高レベル (CRI) と低レベル (OCI) の二種類がある。

- kubernetes のクラスタ全体を管理しているのが Controll Plane で、クラスタ情報を提供する API も Controll Plane 内にある
- 各ノードでコンテナを管理しているのが kubelet で、kubelet は Controll Plane の API を通じて情報を取得し動作する
- kubelet がコンテナを管理する際のインタフェースが CRI、個別のコンテナをホストから隔離して動作させるのが OCI という役割分担で標準化されている
    - いわゆる namespace や cgroups でリソースをホストから分けているんですよ、という部分を主に担当しているのが OCI
- それぞれ containerd, CRI-O や runc, gVisor など複数の実装がある

本書では OCI Runtime の [runc](https://github.com/opencontainers/runc) について解説している。

- コンテナ動作に必要な設定とファイル群を入力としてうけとる
- コンテナを start, list, delete などのサブコマンドで操作する

簡単に試してみる。

デフォルトの OCI Runtime は `docker info` で確認できる。

```
$ docker info | grep 'Default Runtime'
 Default Runtime: runc
```

`centos:8` の image を pull して export し、それを `rootfs` というディレクトリに保存する。

```
vagrant@ubuntu-xenial:~$ mkdir bundle
vagrant@ubuntu-xenial:~$ mkdir bundle/rootfs
vagrant@ubuntu-xenial:~$ sudo docker pull centos:8
8: Pulling from library/centos
7a0437f04f83: Pull complete
Digest: sha256:5528e8b1b1719d34604c87e11dcd1c0a20bedf46e83b5632cdeac91b8c04efc1
Status: Downloaded newer image for centos:8
docker.io/library/centos:8
vagrant@ubuntu-xenial:~$ sudo docker run --rm --name tmp -d centos:8 sleep infinity
915b765161fc62d69c208e41052343b16aa59d016358600abd80aabe3bf90bca
vagrant@ubuntu-xenial:~$ sudo docker export tmp | tar -xC bundle/rootfs
vagrant@ubuntu-xenial:~$ ls bundle/rootfs
bin  dev  etc  home  lib  lib64  lost+found  media  mnt  opt  proc  root  run  sbin  srv  sys  tmp  usr  var
```

file bundle の設定ファイルは雛形を `runc spec` で出力できるらしい。今回はこれをそのまま使う。

```
vagrant@ubuntu-xenial:~$ runc spec -b bundle
vagrant@ubuntu-xenial:~$ cat bundle/config.json | head
{
        "ociVersion": "1.0.2-dev",
        "process": {
                "terminal": true,
                "user": {
                        "uid": 0,
                        "gid": 0
                },
                "args": [
                        "sh"
vagrant@ubuntu-xenial:~$ ls -l bundle
total 8
-rw-rw-r--  1 vagrant vagrant 2652 Dec 30 07:39 config.json
drwxrwxr-x 17 vagrant vagrant 4096 Dec 30 07:39 rootfs
```

これを引数に渡し、`runc run` でコンテナを起動できる。

```
vagrant@ubuntu-xenial:~$ sudo runc run -b bundle mycentos
sh-4.4# cat /etc/os-release
NAME="CentOS Linux"
VERSION="8"
ID="centos"
ID_LIKE="rhel fedora"
VERSION_ID="8"
PLATFORM_ID="platform:el8"
PRETTY_NAME="CentOS Linux 8"
ANSI_COLOR="0;31"
CPE_NAME="cpe:/o:centos:centos:8"
HOME_URL="https://centos.org/"
BUG_REPORT_URL="https://bugs.centos.org/"
CENTOS_MANTISBT_PROJECT="CentOS-8"
CENTOS_MANTISBT_PROJECT_VERSION="8"
sh-4.4# ls
bin  dev  etc  home  lib  lib64  lost+found  media  mnt  opt  proc  root  run  sbin  srv  sys  tmp  usr  var
sh-4.4# ps aux
USER       PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
root         1  0.0  0.3  12020  3196 pts/0    Ss   07:40   0:00 sh
root         8  0.0  0.3  44628  3436 pts/0    R+   07:41   0:00 ps aux
```

list, kill なども可能。

```
vagrant@ubuntu-xenial:~$ sudo runc list
ID          PID         STATUS      BUNDLE                 CREATED                          OWNER
mycentos    4854        running     /home/vagrant/bundle   2020-12-30T07:41:37.192015773Z   root
vagrant@ubuntu-xenial:~$ sudo runc kill mycentos
vagrant@ubuntu-xenial:~$ sudo runc list
ID          PID         STATUS      BUNDLE      CREATED     OWNER
```

## 参考

- [Docker Image Specification v1.](https://github.com/moby/moby/tree/master/image/spec)
- [opencontainers/image\-spec: OCI Image Format](https://github.com/opencontainers/image-spec)
- [containers/skopeo: Work with remote images registries \- retrieving information, images, signing content](https://github.com/containers/skopeo)
- [OCI Image Format Specification v1\.0\.1を読んで \| うなすけとあれこれ](https://blog.unasuke.com/2018/read-oci-image-spec-v101/)
- [Docker storage drivers \| Docker Documentation](https://docs.docker.com/storage/storagedriver/select-storage-driver/)
- [Use the OverlayFS storage driver \| Docker Documentation](https://docs.docker.com/storage/storagedriver/overlayfs-driver/)
- [Overlay2 について調べてみた \- Qiita](https://qiita.com/toshihirock/items/e99889e4a77a76f28455)
- [OverlayFS \- Wikipedia](https://en.wikipedia.org/wiki/OverlayFS)
- [Kubernetes Components \| Kubernetes](https://kubernetes.io/docs/concepts/overview/components/)
- [コンテナユーザなら誰もが使っているランタイム「runc」を俯瞰する\[Container Runtime Meetup \#1発表レポート\] \| by Kohei Tokunaga \| nttlabs \| Medium](https://medium.com/nttlabs/runc-overview-263b83164c98)
- [opencontainers/runc: CLI tool for spawning and running containers according to the OCI specification](https://github.com/opencontainers/runc)
- [opencontainers/runtime\-spec: OCI Runtime Specification](https://github.com/opencontainers/runtime-spec)
    - [runtime\-spec/bundle\.md at master · opencontainers/runtime\-spec](https://github.com/opencontainers/runtime-spec/blob/master/bundle.md)
    - [runtime\-spec/runtime\.md at master · opencontainers/runtime\-spec](https://github.com/opencontainers/runtime-spec/blob/master/runtime.md)
