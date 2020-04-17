{"title":"jenkins に急にログインできなくなった件","date":"2014-02-21T00:53:17+09:00","tags":["nix"]}

理由はさっぱり不明だけれど、自前でたてている jenkins に急にログインできなくなった。

認証の設定は以下のように、閲覧・操作にはログイン必須にして、一つのアカウントに全権限を付与していた。

[さくら VPS CentOS6 に jenkins を導入 - Please Sleep](http://please-sleep.cou929.nu/install-jenkins-on-centos6-with-auth.html)

全権限をつけたアカウントで急にログインできなくなってしまった。

### 解決法

乱暴だが jenkins のサーバに入って config を書き換え、一時的に認証をしないようにする。その状態で Web UI からアカウントを作りなおして、認証設定を再度行った。

CentOS の場合設定ファイルは jenkins のルートディレクトリ直下にある。

`/var/lib/jenkins/config.xml`

この xml を開き、

- `useSecurity` を `false`
- `authorizationStrategy` を削除

して設定ファイルを保存。jenkins を再起動すると認証なしの状態に戻すことができた。

あとは [さくら VPS CentOS6 に jenkins を導入 - Please Sleep](http://please-sleep.cou929.nu/install-jenkins-on-centos6-with-auth.html) と同様の作業をしてアカウントを設定しなおした。

### まとめ

jenkins の設定ファイルを直接編集すれば最悪なんとかなる。今回の原因は不明で、ユーザーと同名の `開発者` ができたのがトリガーになったような気もするが、深追いはしていない。

