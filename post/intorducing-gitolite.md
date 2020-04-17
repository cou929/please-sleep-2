{"title":"gitolite 導入","date":"2012-02-25T19:07:29+09:00","tags":["nix"]}

以下リポジトリをたてるサーバを gitserver, ローカルマシンを local として, プロンプトで示してます.

    gitserver$ uname -a
    Linux gitserver 2.6.32-5-xen-686 #1 SMP Thu Nov 3 09:08:23 UTC 2011 i686 GNU/Linux
    gitserver$ cat /etc/debian_version
    6.0.3

### gitolite ユーザ作成

    gitserver$ sudo /usr/sbin/useradd -m -s /bin/sh gitolite
    gitserver$ sudo passwd gitolite
    gitserver$ sudo mkdir -m 700 /home/gitolite/.ssh
    gitserver$ sudo chown -R gitolite:gitolite /home/gitolite/.ssh/

今回の場合は sshd_config で AllowUser を設定していたので, gitolite ユーザにも権限を与えてあげる.

    gitserver$ sudo vi /etc/ssh/sshd_config
    gitserver$ sudo /etc/init.d/ssh restart
    gitserver$ sudo /etc/init.d/ssh status
    sshd is running.

### gitolite インストール

    gitserver$ cp .ssh/authorized_keys /tmp/kosei.pub
    gitserver$ git clone git://github.com/sitaramc/gitolite
    gitserver$ sudo gitolite/src/gl-system-install
    gitserver$ ls /usr/local/bin/gl*
    /usr/local/bin/gl-admin-push    /usr/local/bin/gl-conf-convert  /usr/local/bin/gl-mirror-push   /usr/local/bin/gl-setup           /usr/local/bin/gl-time
    /usr/local/bin/gl-auth-command  /usr/local/bin/gl-dryrun        /usr/local/bin/gl-mirror-shell  /usr/local/bin/gl-setup-authkeys  /usr/local/bin/gl-tool
    /usr/local/bin/gl-compile-conf  /usr/local/bin/gl-install       /usr/local/bin/gl-query-rc      /usr/local/bin/gl-system-install
    gitserver$ ls /var/gitolite/
    conf  hooks
    gitserver$ su gitolite
    gitserver$ gl-setup -q /tmp/kosei.pub
    gitserver$ exit

### admin 用リポジトリのクローン

    local$ git clone ssh://gitolite@gitserver/gitolite-admin.git
    edward:kosei% ls gitolite-admin/*
    gitolite-admin/conf:
    gitolite.conf
    
    gitolite-admin/keydir:
    kosei.pub

### 新しいリポジトリを作ってみる

    local$ cd gitolite-admin/
    local$ vi conf/gitolite.conf
    
    repo    gitolite-admin
            RW+     =   kosei
    
    repo    testing
            RW+     =   @all
    
    repo    testrepo            # この2行を
            RW+     =   @all    # 追加

    local$ git diff
    diff --git a/conf/gitolite.conf b/conf/gitolite.conf
    index 684deec..449fd55 100644
    --- a/conf/gitolite.conf
    +++ b/conf/gitolite.conf
    @@ -3,3 +3,6 @@ repo    gitolite-admin
    
     repo    testing
             RW+     =   @all
    +
    +repo   testrepo
    +       RW+     =   @all
    
    local$ git commit -am 'test to add repo'
    local$ git push origin master
    
    Koseis-MBA:kosei% git clone ssh://gitolite@cou929.no-ip.org:1234/testrepo.git
    Cloning into testrepo...
    warning: You appear to have cloned an empty repository.

ユーザの追加は keydir に username.pub という公開鍵を置いてコミットすればいいらしい
