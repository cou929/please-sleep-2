{"title":"GitHub Organization の移行機能を使ったデータエクスポート","date":"2020-04-25T11:38:00+09:00","tags":["tips"]}

GitHub には Organization データ移行のための migration api がある。例えば .com から GHE への移行と言ったユースケースを想定した API らしい。これを援用して現状の Organization 全体のバックアップもできる。

- [GitHub\.comのOrganizationのリポジトリのエクスポート \- GitHub ヘルプ](https://help.github.com/ja/enterprise/2.17/admin/migrations/exporting-the-githubcom-organizations-repositories)
- [GitHub\.comのOrganizationのリポジトリのエクスポート \- GitHub ヘルプ](https://help.github.com/ja/enterprise/2.17/admin/migrations/exporting-the-githubcom-organizations-repositories)
- [Organization migrations \| GitHub Developer Guide](https://developer.github.com/v3/migrations/orgs/)

コードだけでなく、issue や wiki などの付随した資産を一括で取得できるようだ。デフォルトでは issue 等に添付したファイルも取得してくれるので便利そう。

作業の流れとしては、

- migration を作成
- 進捗を確認
- アーカイブをダウンロード
- 必要に応じてロックの解除

となる。

## migration を作成

```
curl -H "Authorization: token GITHUB_ACCESS_TOKEN" -X POST \
-H "Accept: application/vnd.github.wyandotte-preview+json" \
-d'{"lock_repositories":true,"repositories":["org_name/repo_name"] \
https://api.github.com/orgs/org_name/migrations
```

このとき `lock_repositories` を偽にするとロックせずにマイグレーションを作成する。本当に移行したい際はロックした方がよいが、バックアップをとりたいだけだったらロックしなくてもいいかもしれない。

ロックするとそのリポジトリへのアクセスが一切できなくなる。Web からアクセスしてもその旨のエラー画面が表示される。

（ロックしない場合にどういう挙動になるかは試していないのでわからない）

## 進捗を確認

```
curl -H "Authorization: token GITHUB_ACCESS_TOKEN" \
-H "Accept: application/vnd.github.wyandotte-preview+json" \
https://api.github.com/orgs/org_name/migrations/id
```

id は作成した migration の id。
state パラメータに状態が入っている。

## アーカイブをダウンロード

```
curl -H "Accept: application/vnd.github.wyandotte-preview+json" \
-u user_name:GITHUB_ACCESS_TOKEN \
-L -o archive.tar.gz \
https://api.github.com/orgs/org_name/migrations/id/archive
```

## 必要に応じてロックの解除

移行元の Organization は基本的にはずっとロックがかかっているようだ。通常のユースケースではロックしたまま削除する流れを想定しているのだと思う。

ロックを解除する API はあるが、リポジトリ単位のようだ。複数のリポジトリを移行する場合は一つ一つ解除する必要がある。

```
curl -H "Authorization: token GITHUB_ACCESS_TOKEN" -X DELETE \
-H "Accept: application/vnd.github.wyandotte-preview+json" \
https://api.github.com/orgs/org_name/migrations/id/repos/repo_name/lock
```

<iframe style="width:120px;height:240px;" marginwidth="0" marginheight="0" scrolling="no" frameborder="0" src="//rcm-fe.amazon-adsystem.com/e/cm?lt1=_blank&bc1=000000&IS2=1&bg1=FFFFFF&fc1=000000&lc1=0000FF&t=pleasesleep-22&language=ja_JP&o=9&p=8&l=as4&m=amazon&f=ifr&ref=as_ss_li_til&asins=B07JLJSDMJ&linkId=f50c7f8d9fd59c3bf1e9d7030f5a79ca"></iframe>
