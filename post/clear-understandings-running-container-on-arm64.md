{"title":"Docker イメージ・実行環境と CPU アーキテクチャの理解整理","date":"2022-01-02T23:45:00+09:00","tags":["nix"]}

[multipass で開発環境を構築したメモ - Please Sleep](https://please-sleep.cou929.nu/multipass-cloud-init-aarch64.html) にて、M1 Mac で [Multipass](https://multipass.run/) を使って開発環境を作ってみていたところ、うまく動かないコンテナがあった。初見では不可解に感じられるポイントがあったので、理解を整理したメモ。

## 状況

- docker-compose でいくつかのコンテナをまとめた環境がある
- M1 Mac での直接の実行はできる
- multipass で起動した ubuntu 20.04 の仮想環境ではうまく動作しなかった

次のようなエラーでうまく実行できなかった。

- pull した一部のイメージの実行時にアーキテクチャが違うというエラーが出た
- arm64 のイメージを pull するようにしてみたが、該当のイメージが無いというエラーが出ることもあれば、別のイメージは pull はできるが実行時エラーということもあった

```sh
# 実行時エラーの例
standard_init_linux.go:228: exec user process caused: exec format error

# pull しようとしてアーキテクチャにマッチするイメージが無いというエラーの例
$ docker pull mysql:5.7
5.7: Pulling from library/mysql
no matching manifest for linux/arm64/v8 in the manifest list entries
```

## よくわからなかったポイント

- 同じ arm64 のアーキテクチャなので、両方動かないのなら納得だけど、M1 だと動作したのはなぜ?
- もともと mysql:5.7 イメージは `platform: linux/x86_64` を指定して M1 環境で動作していたが、これはなんだっけ?
- docker pull は実行環境のアーキテクチャに応じてマッチするイメージを持ってきてくれているようだが、マッチするイメージが無いというエラーではなく、pull はできてその後に実行時エラーになるのはなぜ?

このあたりが一見不思議に感じられたので、今回整理したという背景。

## 前提をいくつか

ホストの M1 Mac での docker version

```sh
$ docker version
Client:
 Cloud integration: v1.0.22
 Version:           20.10.11
 API version:       1.41
 Go version:        go1.16.10
 Git commit:        dea9396
 Built:             Thu Nov 18 00:36:09 2021
 OS/Arch:           darwin/arm64
 Context:           default
 Experimental:      true

Server: Docker Engine - Community
 Engine:
  Version:          20.10.11
  API version:      1.41 (minimum version 1.12)
  Go version:       go1.16.9
  Git commit:       847da18
  Built:            Thu Nov 18 00:34:44 2021
  OS/Arch:          linux/arm64
  Experimental:     false
 containerd:
  Version:          1.4.12
  GitCommit:        7b11cfaabd73bb80907dd23182b9347b4245eb5d
 runc:
  Version:          1.0.2
  GitCommit:        v1.0.2-0-g52b36a2
 docker-init:
  Version:          0.19.0
  GitCommit:        de40ad0
```

その上で `multipass launch` した ubuntu 20.04 の仮想環境での docker version

```sh
$ docker version
Client: Docker Engine - Community
 Version:           20.10.12
 API version:       1.41
 Go version:        go1.16.12
 Git commit:        e91ed57
 Built:             Mon Dec 13 11:44:28 2021
 OS/Arch:           linux/arm64
 Context:           default
 Experimental:      true

Server: Docker Engine - Community
 Engine:
  Version:          20.10.12
  API version:      1.41 (minimum version 1.12)
  Go version:       go1.16.12
  Git commit:       459d0df
  Built:            Mon Dec 13 11:43:05 2021
  OS/Arch:          linux/arm64
  Experimental:     false
 containerd:
  Version:          1.4.12
  GitCommit:        7b11cfaabd73bb80907dd23182b9347b4245eb5d
 runc:
  Version:          1.0.2
  GitCommit:        v1.0.2-0-g52b36a2
 docker-init:
  Version:          0.19.0
  GitCommit:        de40ad0
```

実行しようとした `dokcer-compose.yml` はこんな感じ。いくつかの公式イメージのミドルウェア、非公式イメージのミドルウェア、自分たちで作っている rails アプリという構成。

```yaml
version: "3"
services:
  db:
    image: mysql:5.7  # Docker 公式の mysql image https://hub.docker.com/_/mysql
    platform: linux/x86_64  # 以前 M1 Mac で動作させるために追記されたもの
    # ...
  redis:
    image: redis:4-alpine  # Docker 公式の redis image https://hub.docker.com/_/redis
    # ...
  api:
    build: .  # Docker 公式の ruby image https://hub.docker.com/_/ruby をもとに rails アプリをビルド
    # ...
  some-middleware:
    image: some-middleware/some-middleware:latest  # Docker 公式でないミドルウェアのイメージ
    # ...
```

ホスト側、仮想環境側のどちらも `OS/Arch: linux/arm64` だが、仮想環境では前述のようにうまく動作しなかった。

## Docker Desktop for Mac では x86 向けのイメージも動く

[Docker Desktop for Apple silicon | Docker Documentation](https://docs.docker.com/desktop/mac/apple-silicon/) や [Leverage multi-CPU architecture support | Docker Documentation](https://docs.docker.com/desktop/multi-arch/) によると、Docker Desktop for Mac では x86 向けのイメージでも動作させられるらしい。

- [qemu](https://wiki.qemu.org/Main_Page) を使って実現している
- 完全に動作するかはあくまでベストエフォートで、またエミュレーションしている関係上パフォーマンスも劣るので、arm 用にビルドしたイメージを使えるならそうした方がベストではある
- [mysql の公式イメージ](https://hub.docker.com/_/mysql?tab=tags&page=1&ordering=last_updated) は arm 向けのビルドが無いので、`--platform linux/amd64` を指定して x86 用イメージを明示的に取得して動かすか、[mariadb](https://hub.docker.com/_/mariadb?tab=tags&page=1&ordering=last_updated) なら arm 対応ビルドが用意されているからそちらを使うよう [案内している](https://docs.docker.com/desktop/mac/apple-silicon/#known-issues)

このエミュレーションの仕組みは Docker Desktop に実装されているものなので、linux 環境にて docker client, server 使っているケースでは利用できない。

## multi-arch images というひとつのイメージで複数のアーキテクチャに対応する仕組みがある

[Leverage multi-CPU architecture support | Docker Documentation](https://docs.docker.com/desktop/multi-arch/) によると、複数のアーキテクチャに対応したひとつのイメージを作ることができる、multi-arch images という仕組みがあるらしい。

- docker pull する際に、対象のイメージが multi-arch 対応だった場合、そのホストのアーキテクチャにマッチするイメージを自動的に pull してくれる
- Docker の公式イメージは [すべて multi-arch 対応済み](https://github.com/docker-library/official-images#what-are-official-images) らしい
- `docker buildx` というツールを使ってわりと簡単にビルドできる。使い方のチュートリアルが Web に結構ある

この multi-arch と前述の qemu による x86 のエミュレーションのおかげで、一般的な開発者は手元のマシンを M1 Mac に買い替えても、これまでと同じコンテナがほぼそのまま動いていると思う。なるほどという仕組みだった。

## docker pull は対象のイメージが multi-arch の時だけ対応するイメージの有無をチェックする

ここが個人的には一番ミスリードされていた部分だった。docker pull したときの動作が、対象のイメージが multi-arch かどうかで変わる。

- 対象のイメージが multi-arch 対応の場合
    - アーキテクチャにマッチするイメージがあるかを確認し、なければエラーで pull しない
    - platform オプションで指定したアーキテクチャ対応のイメージを強制的に pull することもできる
- 対象のイメージが multi-arch 未対応の場合 (single-arch image の場合)
    - 特にチェックはなく指定されたイメージを pull する
    - platform オプションは単に無視される（[warning は出る](https://github.com/docker/for-mac/issues/5625))

この仕様は次のように確認した。

- multi-arch と single-arch のイメージのメタ情報をそれぞれ見たところ、マニフェストファイルの MediaType がそれぞれ `application/vnd.docker.distribution.manifest.list.v2+json` と `application/vnd.docker.distribution.manifest.v2+json` というものだった

```sh
# multi-arch のイメージ
2022-01-02 14:26 ubuntu@t2 moby master$ docker buildx imagetools inspect busybox
Name:      docker.io/library/busybox:latest
MediaType: application/vnd.docker.distribution.manifest.list.v2+json
Digest:    sha256:5acba83a746c7608ed544dc1533b87c737a0b0fb730301639a0179f9344b1678

Manifests:
  Name:      docker.io/library/busybox:latest@sha256:62ffc2ed7554e4c6d360bce40bbcf196573dd27c4ce080641a2c59867e732dee
  MediaType: application/vnd.docker.distribution.manifest.v2+json
  Platform:  linux/amd64

  Name:      docker.io/library/busybox:latest@sha256:ca038f83e1a3a6a08b539830ca3beefb503a3989cc1f19c265ae4e624a45a9cc
  MediaType: application/vnd.docker.distribution.manifest.v2+json
  Platform:  linux/arm/v5
  ...

# single-arch のイメージ
$ docker buildx imagetools inspect circleci/golang
{
   "schemaVersion": 2,
   "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
   "config": {
      "mediaType": "application/vnd.docker.container.image.v1+json",
      "size": 12844,
      "digest": "sha256:3376bfde60d59ef847fa36fedb4c1695e8950c0db240c47f332fc7a7d5d124c6"
   },
   "layers": [
       ...
```

- ここから少しコードを追ってみる
- それぞれのマニフェストファイルを格納する構造体の型
    - multi-arch の mediaType `application/vnd.docker.distribution.manifest.list.v2+json` は [manifestlist.DeserializedManifestList 型](https://github.com/distribution/distribution/blob/6a977a5a754baa213041443f841705888107362a/manifest/manifestlist/manifestlist.go#L16)
    - single-arch の mediaType `application/vnd.docker.distribution.manifest.v2+json` は [schema2.DeserializedManifest 型](https://github.com/distribution/distribution/blob/6a977a5a754baa213041443f841705888107362a/manifest/schema2/manifest.go#L15)
- pull する処理はそれぞれ [前者には pullManifestList、後者には pullSchema2](https://github.com/moby/moby/blob/10aecb0e652d346130a37e5b4383eca28f594c21/distribution/pull_v2.go#L440-L454) というメソッドが使われる
    - 前者は [マニフェストのリストから platform に合致するものを選ぶという処理](https://github.com/moby/moby/blob/10aecb0e652d346130a37e5b4383eca28f594c21/distribution/pull_v2.go#L791-L797) が入っていて、マッチするものがなければエラーを返している
    - 後者は [そのまま pull している模様](https://github.com/moby/moby/blob/10aecb0e652d346130a37e5b4383eca28f594c21/distribution/pull_v2.go#L740)

つまり、multi-arch のマニフェストファイルには、各アーキテクチャに対応してマニフェストがリストで入っていて、必然的に pull 時にそれをチェックしている。一方で single-arch の場合はその必要がなく単に pull するだけ。これらは後から multi-arch 機能が追加されたことを考慮すると自然な流れだけど、初見ではややこしい仕様になっていると思う。`platform` というオプションがあるので、常にイメージのアーキテクチャが自動判定されるように一見思えるが、実は対象のイメージによって判定されたりされなかったりする。

もうひとつポイントなのは、[公式の mysql イメージ](https://hub.docker.com/_/mysql) は multi-arch 対応済みだが、x86 (amd64) にしか対応していないという状態だった。実際 inspect すると次のような出力になっている。

```sh
$ docker buildx imagetools inspect mysql
Name:      docker.io/library/mysql:latest
MediaType: application/vnd.docker.distribution.manifest.list.v2+json
Digest:    sha256:e9027fe4d91c0153429607251656806cc784e914937271037f7738bd5b8e7709

Manifests:
  Name:      docker.io/library/mysql:latest@sha256:238cf050a7270dd6940602e70f1e5a11eeaf4e02035f445b7f613ff5e0641f7d
  MediaType: application/vnd.docker.distribution.manifest.v2+json
  Platform:  linux/amd64
```

前述のように multi-arch のイメージは pull 時にアーキテクチャのチェックが走るので、例えば M1 Mac で mysql イメージを使おうとすると `no matching manifest for linux/arm64/v8 in the manifest list entries` というエラーが出ることになる。`platform: linux/x86_64` オプションを指定して強制的に pull し Docker Desktop のエミュレーションによって動作させるというワークアラウンドが可能になっている。

## 以上の情報で当初の疑問を整理する

- 同じ arm64 のアーキテクチャなので、両方動かないのなら納得だけど、M1 だと動作したのはなぜ?
    - Docker Desktop のエミュレーション機能によって amd64 用イメージも動作したから
    - また公式イメージは multi-arch 対応をしているため、イメージ側が対応していれば arm64 用のものが Dockerfile の書き換えなしに取得できていたから
- もともと mysql:5.7 イメージは `platform: linux/x86_64` を指定して M1 環境で動作していたが、これはなんだっけ?
    - Docker 公式の mysql イメージは multi-arch 対応しているが、x86 用のイメージしか用意されていない (マニフェストのフォーマットとしては multi-arch だが、1 アーキテクチャにしか対応していない)
    - そのため arm のマシンから pull しようとするとマッチするイメージが無いというエラーになっていまう
    - M1 Mac など Docker Desktop を使っている環境では x86 を明示的に指定して pull しエミュレーションで動作させるワークアラウンドが可能なのでこうしていた
- docker pull は実行環境のアーキテクチャに応じてマッチするイメージを持ってきてくれているようだが、マッチするイメージが無いというエラーではなく、pull はできてその後に実行時エラーになるのはなぜ?
    - multi-arch 対応のイメージの場合、arm64 にマッチするものがなければその旨のエラーが出ていた
    - single-arch の amd64 イメージの場合はそれが pull され、実行時にエラーになっていた
    - 今回の自分の例だと
        - redis 公式イメージは arm64 対応済みなので問題ない
        - rails アプリはもととなっている ruby 公式イメージが arm64 対応済みなので問題ない
        - mysql 公式イメージ (multi-arch) は arm64 未対応なのでエラー
            - マッチするイメージが無いというエラー、あるいは x86 を指定して pull すると実行時エラーになる
        - ミドルウェアの非公式イメージ (single-arch) は x86 用なのでエラー
            - 実行時にエラーになる

## まとめ

- Docker Desktop のエミュレーション機能と multi-arch images によって、ユーザーは M1 Mac でほとんど何も気にせずに既存の Dockerfile がつかえるようになっている
- ただその副作用で pull の仕様が複雑化しているように思われた

## PR

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B08PNMRXKN/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/41g+F7WohJL.jpg" alt="イラストでわかるDockerとKubernetes Software Design plus" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B08PNMRXKN/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">イラストでわかるDockerとKubernetes Software Design plus</a></div><div class="amazlet-detail">徳永 航平  (著)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B08PNMRXKN/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
