{"title":"jenkins のアップデート","date":"2014-02-20T08:33:51+09:00","tags":["nix"]}

新しいバージョンがあった場合、jenkins の設定画面に通知がでる。そのリンクから `jenkins.war` をダウンロードできるので、それを配置するだけで良い。

念のためバージョンをつけてバイナリをとっておき、シンボリックリンクをはることにした。以下は CentOS6 での例。

    $ curl -OL http://updates.jenkins-ci.org/download/war/1.551/jenkins.war
    $ sudo /etc/init.d/jenkins stop
    $ sudo mv /usr/lib/jenkins/jenkins.war /usr/lib/jenkins/jenkins.war.1.541  # 既存の war を退避
    $ sudo mv jenkins.war /usr/lib/jenkins/jenkins.war.1.551                   # ダウンロードした war を配置
    $ sudo ln -s /usr/lib/jenkins/jenkins.war.1.551 /usr/lib/jenkins/jenkins.war
    $ sudo /etc/init.d/jenkins start

`jenkins.war` の配置がわからなかったので `/etc/init.d/jenkins` の中を覗いてしらべた。

    $ grep jenkins.war /etc/init.d/jenkins
    JENKINS_WAR="/usr/lib/jenkins/jenkins.war"

