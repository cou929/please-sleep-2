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

## アーカイブの内容

ソースコードは git 形式で、issue, pull request, milestone 等々はすべて json 形式で保存されている。また issue 等に添付されたファイルもそのまま入っている。ヒューマンリーダブルではないが、バックアップにはなっている。

```
attachments/
attachments_000001.json
commit_comments_000001.json
issue_comments_000001.json
issue_events_000001.json
issues_000001.json
milestones_000001.json
organizations_000001.json
projects_000001.json
protected_branches_000001.json
pull_request_review_comments_000001.json
pull_request_reviews_000001.json
pull_requests_000001.json
releases_000001.json
repositories/
repositories_000001.json
repository_files/
repository_files_000001.json
schema.json
teams_000001.json
users_000001.json
```

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B07JLJSDMJ/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/51tcyTz82UL.jpg" alt="GitHub実践入門──Pull Requestによる開発の変革 WEB+DB PRESS plus" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B07JLJSDMJ/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">GitHub実践入門──Pull Requestによる開発の変革 WEB+DB PRESS plus</a></div><div class="amazlet-detail">大塚 弘記  (著)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B07JLJSDMJ/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
