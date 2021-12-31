{"title":"multipass で開発環境を構築したメモ","date":"2022-01-01T01:00:00+09:00","tags":["nix"]}

[Multipass](https://multipass.run/) 上に開発環境を構築したメモ。

## 背景

- もともと Vagrant + VirtualBox (クライアントは MacOS) で開発環境を構築していた
- M1 チップの Mac に乗り換えたところ、VirtualBox が未対応だったので、EC2 (ubuntu) 上に環境を移した
- 今回 multipass がなかなか良さそうなので環境の移植を試した

## multipass とは

- canonical が開発している ubuntu の仮想環境がさくっと立ち上げられるもの
- cli のみだったり、コマンドもシンプルだったり、cloud-init に対応していたり、なかなか developer friendly というかセンスが良い第一印象

## やってみた感想の雑多なメモ

- (当たり前かもだが) M1 Mac 上ではアーキテクチャが aarc64 の ubuntu のイメージしかない
    - 自分の場合これまで x86 前提の環境だったので、結局ここの移行コストが一番高かった (まだ全部終わっていない)

```sh
2022-01-01 00:31 ubuntu@t2 $ uname -a
Linux t2 5.4.0-91-generic #102-Ubuntu SMP Fri Nov 5 16:30:45 UTC 2021 aarch64 aarch64 aarch64 GNU/Linux
```

- cloud-init の [ユーザーデータフォーマット](https://cloudinit.readthedocs.io/en/latest/topics/format.html) として [cloud-config の yaml 形式にしか対応していない](https://ubuntu.com/blog/using-cloud-init-with-multipass)
    - これまで shell script を使っていたので yaml に移植する必要が出てしまった
    - 簡潔な仕様なのでそんなに大変ではないけど、cloud-init 対応が謳われていて既存の資産がそのまま使えるかなと期待していたのは裏切られた
- 次のようなコマンドで直接 ssh アクセスができるらしい
    - multipassd が管理している秘密鍵を使えば良いらしい
    - このあたりを `multipass shell` などが隠蔽してくれているのだと思う
    - ssh port forward もこれでできる
    - 秘密鍵の所有者が root なので、適宜パーミッションの調整が必要
        - 例えば multipass のインスタンスに vscode remote でつなぐため、とりあえず秘密鍵を別の場所にコピーし一般ユーザーに chown した
            - もっと適切なやり方がありそうだが一旦こういう方法もあるということで
    - ref. [Port forwarding · Issue #309 · canonical/multipass](https://github.com/canonical/multipass/issues/309)

```sh
sudo ssh -i /var/root/Library/Application\ Support/multipassd/ssh-keys/id_rsa ubuntu@<multipass instance ip>

# port forward するならこんな感じ
sudo ssh -N -i /var/root/Library/Application\ Support/multipassd/ssh-keys/id_rsa -L 0.0.0.0:3000:localhost:3000 ubuntu@<multipass instance ip>
```

- cloud-config の [packages](https://cloudinit.readthedocs.io/en/latest/topics/examples.html#install-arbitrary-packages) に存在しないパッケージ名が混ざるとすべてのインストールが行われない
    - package のインストールは [apt-get instsall](https://github.com/canonical/cloud-init/blob/bae9b11da9ed7dd0b16fe5adeaf4774b7cc628cf/cloudinit/distros/debian.py#L246) に [yaml で指定したすべてのパッケージを一括で渡している](https://github.com/canonical/cloud-init/blob/bae9b11da9ed7dd0b16fe5adeaf4774b7cc628cf/cloudinit/distros/debian.py#L138) ため
    - yaml 上は一列に並んでいるので無意識に「問題のパッケージは失敗するがそれ以前のパッケージはインストール済みの状態になっている」と思ったがそうではなかった
    - (multipass は関係ない cloud-init の話)
- shebang が `#! /bin/sh` の場合、source コマンドが無い
    - まず `/bin/sh` の実態は他のシェルで、bash が使われることが多いが ubuntu は dash らしい
    - dash には source が無いので、`#! /bin/sh` ... `source ~/.bashrc` などとすると失敗する
    - `.` の方が汎用的なのでそちらのほうがいいらしい (i.e., `. ~/.bashrc`)
    - ref. [bash - source command not found in sh shell - Stack Overflow](https://stackoverflow.com/questions/13702425/source-command-not-found-in-sh-shell)
    - (multipass は関係ない ubuntu の話)

## PR

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4297116901/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://images-na.ssl-images-amazon.com/images/I/51jIGaXilNL._SX352_BO1,204,203,200_.jpg" alt="図解即戦力 仮想化&コンテナがこれ1冊でしっかりわかる教科書" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4297116901/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">図解即戦力 仮想化&コンテナがこれ1冊でしっかりわかる教科書</a></div><div class="amazlet-detail">五十嵐 貴之  (著), 薄田 達哉  (著)<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4297116901/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
