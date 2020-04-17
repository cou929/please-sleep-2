{"title":"GitHub に Deploy key で push した際のアカウント","date":"2015-05-31T23:06:06+09:00","tags":["nix"]}

### tl;dr

GitHub に Deploy key で push した場合、その deploy key を登録したアカウントが push したとして扱われる。

### 背景

外のシステムと連携するために、GitHub ではリポジトリごとに公開鍵を登録する、[deploy key](https://developer.github.com/guides/managing-deploy-keys/#deploy-keys) という機能がある。

この deploy key を使って自動でリポジトリに tag を push していたところ、その tag の Author や Commiter は意図したユーザになっていたんだけど、それを push したユーザーが全然違うアカウントとして GitHub に認識されていて、困っていた。具体的には Webhook で push を外のチャットに通知していたが、そのメッセージ (`tag FOO pushed by BAR` のようなかんじ) として表示される「push した人」が意図しないアカウントになっていた。`.ginconfig` を編集しても反映されなかった。

整理してみると、このような話だった。

- tag の Authoer や Committer は git が管理している。そのため当然、`~/.gitconfig` などに記載のあるユーザーとして扱われている
- 一方で「push した人」は git は関係なくて、GitHub が扱う範囲。普通に鍵認証で GitHub に接続している場合は、その鍵が紐付いているアカウントが push したアカウントとして、GitHub に認識される。

今回は deploy key をつかって push していたので、この鍵のアカウントが誰として認識されているのかという話になる。そこで GitHub のサポートに問い合わせてみたところ、その deploy key を追加したアカウントが「push した人」として認識されるという仕様ということだった。
