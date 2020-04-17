{"title":"さくら VPS CentOS6 に jenkins を導入","date":"2013-12-01T03:44:52+09:00","tags":["nix"]}

### インストール

- java のインストール

        $ sudo yum install java-1.7.0-openjdk.x86_64

- jenkins のインストール

        $ sudo wget -O /etc/yum.repos.d/jenkins.repo http://pkg.jenkins-ci.org/redhat/jenkins.repo
        $ sudo rpm --import http://pkg.jenkins-ci.org/redhat/jenkins-ci.org.key
        $ sudo yum install jenkins

[Installing Jenkins on RedHat distributions - Jenkins - Jenkins Wiki](https://wiki.jenkins-ci.org/display/JENKINS/Installing+Jenkins+on+RedHat+distributions) を参考に。

chkconfig をみたところすでに on になっている。

    $ sudo /sbin/chkconfig --list | grep jenkins
    jenkins         0:off   1:off   2:off   3:on    4:off   5:on    6:off

デーモンを起動。

    $ sudo /etc/init.d/jenkins start

これでポート 8080 で起動するようだ。

    $ ps aux | grep jenkins
    jenkins  27431  7.4 20.0 1950596 204384 ?      Ssl  03:23   1:00 /etc/alternatives/java -Dcom.sun.akuma.Daemon=daemonized -Djava.awt.headless=true -DJENKINS_HOME=/var/lib/jenkins -jar /usr/lib/jenkins/jenkins.war --logfile=/var/log/jenkins/jenkins.log --webroot=/var/cache/jenkins/war --daemon --httpPort=8080 --ajp13Port=8009 --debug=5 --handlerCountMax=100 --handlerCountMaxIdle=20

### 認証設定

外から見える場所に置きたいかつプライベートで使いたいので認証を設定する。Basic 認証などでもいいのだが Jenkins の認証機構を使ってみる。自分だけログインできて、そうでないアカウントは閲覧を含め一切何もできないようにしたい。

- Jenkins の Web UI トップから "Jenkinsの管理" -> "グローバルセキュリティの設定"
- "セキュリティを有効化" をチェック、"Jenkinsのユーザーデータベース", "ユーザーにサインアップを許可", "ログイン済みユーザーに許可" を選択し保存
  - これで認証はするが、アカウントさえあれば何でも出来る状態になった。またアカウントはだれでも作成できる状態。
- 一度トップに戻りアカウントを作成する。
- 作ったアカウントでログインする。
- ふたたび "グローバルセキュリティの設定" 画面へ
- "ユーザーにサインアップを許可" を未選択にする
  - これで新規アカウントの発行ができなくなった
- 権限管理を "行列による権限設定" に。匿名ユーザーのチェックをすべて外す。さきほど作成したユーザーを追加し、すべてのチェックをつける。
  - これで匿名ユーザーは何もできなく、自分はすべての権限を持っている状態になる。

