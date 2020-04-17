{"title":"github で fork したリポジトリで本家に追従する","date":"2014-04-22T21:28:45+09:00","tags":["git"]}

シンプルに、リモートの本家リポジトリを fetch し手元にマージするという方法。

fork したリポジトリに移動して以下を行う

1. 本家を upstream という名前でリモートとして追加する

        git remote add upstream git://github.com/FOO/BAR.git
        git remote -v show

2. 本家の変更を fetch

        git fetch upstream

3. 本家の変更を merge

        git merge upstream/master

手順としてはシンプルで理解しやすい。

fork したリポジトリの master は常に本家に追従するようにしておいて、pull request を送る際には必ずトピックブランチを切ってから送るように運用すればよさそうだ。

### 参考

[GitHubでFork/cloneしたリポジトリを本家リポジトリに追従する - Qiita](http://qiita.com/xtetsuji/items/555a1ef19ed21ee42873)
