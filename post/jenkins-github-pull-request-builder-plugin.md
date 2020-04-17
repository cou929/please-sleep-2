{"title":"jenkins で GitHub のプルリクエストをマージしてテストする","date":"2014-03-22T16:52:55+09:00","tags":["jenkins"]}

[GitHub pull request builder plugin - Jenkins - Jenkins Wiki](https://wiki.jenkins-ci.org/display/JENKINS/GitHub+pull+request+builder+plugin)

GitHub pull request builder plugin を使うと、

- プルリクエストを自動でマージしてテスト
- 結果を GitHub の Status API で通知
- プルリクエストにコメントしてテストをリトライ

などということができる。

### 設定手順

設定手順は他のシンプルなプラグインに比べてちょっと複雑。

#### ボットユーザの作成

事前にボット用のユーザーを作成して、対象のリポジトリのコラボレーターとして登録しておく。このボットユーザーがプルリクにコメントしたりするので、アイコンをそれっぽくしておくと楽しい。

#### 全体の設定

プラグイン全体の設定をする。ここの設定がデフォルトとなり、プロジェクトごとに上書きして使う。

- Jenkinsの管理 -> システムの設定 -> `GitHub pull requests builder`
- `Access Token` には先ほど作成したボットのアクセストークンを入れる。
  - GitHub のアカウントごとの設定画面より、`Applications` -> `Personal access tokens` -> `Generate new token`
  - スコープは `repo`、 `repo:status` あたりのチェックが入っていることを確認
- `Use comments to report results when updating commit status fails` のチェックを入れると、GitHub 側からのイベントをフックして動作する。ここをオフにして別の箇所 (`Crontab line`) で設定し、ポーリングにすることも可能。
- `Admin list` に admin のユーザ名を入れる。
  - ここで設定したユーザからのプルリクエストの場合のみテストが実行される
  - admin 以外の人からのプルリクエストは、admin ユーザが許可した場合のみ (`test ok` とコメントする) テストされる
  - だれでもテストが実行できてしまう状態を防ぐ意図の機能
  - この設定は、もちろん、プロジェクトごとに上書きすることができる
- `Published Jenkins URL` には jenkins 側の url (ホスト名) を入れる。
  - ボットがコメントする際のリンクに使われる
  - ボットがコメントを一切しないようにしたい場合はここを空にしておく。テストの結果が自体は status api で表示されているので、コメントは必須ではない。

その他、前述の `test ok` のようなコマンドコメントの正規表現を設定する項目や、テスト成功・失敗時の通知のメッセージを設定する項目がある。

<img style="width:90%" alt="" src="/images/jenkins_prp_global_conf.png"/>

#### プロジェクトごとの設定

プロジェクトごとの設定画面。

- `GitHub project` に GitHub のリポジトリホームの URL を入れる
- ソースコード管理は Git を選択
  - `Repository URL` はリポジトリの URL
  - `Credentials` も適宜設定 ([参考](http://please-sleep.cou929.nu/perl-and-grunt-project-ci-with-github-and-jenkins.html))
  - `Credentials` の高度な設定をひらき、`Refspec` に `+refs/pull/*:refs/remotes/origin/pr/*` と設定する
  - `Branch Specifier (blank for 'any')` は `${sha1}` と設定
- ビルド・トリガ は `GitHub pull requests builder` を選択
  - 全体設定の内容がデフォルトとして使われるので、プロジェクト独自に設定したい部分だけをいじれば良い

<img style="width:90%" alt="" src="/images/jenkins_prp_project_conf.png"/>

これで設定は完了。

プルリクエストには次のように表示されるようになる。

<img style="width:90%" alt="" src="/images/jenkins_prp_status.png"/>

`Details` は jenkins のビルドにリンクしている。

### ログの確認

設定がうまくいかない場合は、とにかく jenkins のログを見て頑張るしかない。

ログは Jenkinsの管理 -> システムログ で参照できる。

ここのログもそんなに親切ではないが、ほぼこれしか手がかりはない。

自分の場合ボットをコラボレーターに入れ忘れていたが、ログにはステータス API のリクエストに失敗しているとしかでておらず、どう失敗したかがわからない状態だった。気がつくまでに少し時間がかかってしまった。

### 参考

- [JenkinsプラグインのGitHub pull request builder pluginを使ってみる - 技術めも](http://d.hatena.ne.jp/oovu70/20130118/p1)
- [Jenkinsでプルリクエストをビルドする - Qiita](http://qiita.com/quattro_4/items/6b1962909ce868f12e5a)

<iframe src="http://rcm-fe.amazon-adsystem.com/e/cm?lt1=_blank&bc1=000000&IS2=1&bg1=FFFFFF&fc1=000000&lc1=0000FF&t=pleasesleep-22&o=9&p=8&l=as4&m=amazon&f=ifr&ref=ss_til&asins=4774148911" style="width:120px;height:240px;" scrolling="no" marginwidth="0" marginheight="0" frameborder="0"></iframe>

