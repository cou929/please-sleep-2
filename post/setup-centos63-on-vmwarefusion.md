{"title":"VMWare Fusion \u0026 CentOS6.3 環境構築","date":"2013-01-18T23:21:40+09:00","tags":["nix"]}

### install

[CentOS6.3 の netinstall](http://ftp.jaist.ac.jp/pub/Linux/CentOS/6.3/isos/x86_64/CentOS-6.3-x86_64-netinstall.iso) をいれる. 詳細はつぎを参照:

[Mac OS X Mountain Lion + VMware Fusion 5 + CentOS 6.3 + Apache + mod_proxy + PSGI + Movable Type 5.2 のローカル環境を構築した \| かたつむりくんのWWW](http://www.tinybeans.net/blog/2012/10/09-161209.html)

### ユーザ作成と鍵の準備
useradd して, authorized_keys の準備

### 固定 IP 化
デフォルトでは dhcp なのでこれを固定する.

ホスト OS 側 (Mac) の以下のパスにローカルネットワークのレンジやゲートウェイなどネットワーク設定に必要な値がかいてある.

    /Library/Preferences/VMware\ Fusion/vmnet8/dhcpd.conf

あとはゲスト OS 側 (CentOS) で必要な設定をすればよい.

#### eth0 の設定

- `/etc/sysconfig/network-scripts/ifcfg-eth0`
- `BOOTPROTO` を `static` に
- `IPADDR`, `NETMASK`, `GATEWAY` を追加

        DEVICE="eth0"
        BOOTPROTO="static"
        HWADDR="00:0C:29:4A:48:27"
        NM_CONTROLLED="yes"
        ONBOOT="yes"
        TYPE="Ethernet"
        UUID="411a5059-05d2-4af3-a567-8ef02e55d5f6"
        IPADDR=192.168.69.134
        NETMASK=255.255.255.0
        GATEWAY=192.168.69.2

#### ゲートウェイの設定

- `/etc/sysconfig/network`
- `GATEWAY` を追加

        NETWORKING=yes
        HOSTNAME=dev.centos63
        GATEWAY=192.168.69.2

#### dns の設定

- `/etc/resolv.conf`
- たぶんデフォルトでホスト OS の ip が設定されているのでそのままで OK

あとはネットワークを再起動して ifcofig, route で設定内容を確認しておく

    # /etc/init.d/network restart
    # ifconfig -a
    # route

### SELinux Off

`/etc/selinux/config` の `SELINUX` を `disabled` にしてリブート

### yum
まずは update

    $ sudo yum clean all
    $ sudo yum update

epel, rpmforge, remi を追加

    $ curl -O http://dl.fedoraproject.org/pub/epel/6/x86_64/epel-release-6-8.noarch.rpm
    $ sudo rpm -Uvh epel-release-6-8.noarch.rpm
    $ curl -O http://rpms.famillecollet.com/enterprise/remi-release-6.rpm
    $ sudo rpm -Uvh remi-release-6.rpm
    $ curl -OL http://pkgs.repoforge.org/rpmforge-release/rpmforge-release-0.5.2-2.el6.rf.x86_64.rpm
    $ sudo rpm -Uvh rpmforge-release-0.5.2-2.el6.rf.x86_64.rpm

`/etc/yum.repos.d/epel.repo`, `/etc/yum.repos.d/rpmforge.repo` を編集して `enable=0` にしておく.

### packages

#### make, gcc

    $ sudo yum install -y make gcc

#### zlib-devel (for python)

    $ sudo yum install -y zlib-devel

#### git

    $ sudo yum install  -y git

ついでに `~/.gitconfig` も

    [user]
            name = Kosei Moriyama
            email = cou929@gmail.com
    [color]
            ui = auto
    [alias]
            co = checkout
            st = status -sb
            pr = pull --rebase
            fo = fetch origin
            ro = rebase origin
            rc = rebase --continue
            wd = diff --word-diff
            gp = grep -n
            lg = log --graph --pretty=oneline --decorate --date=short --abbrev-commit --branches
            ci = commit
            br = branch
    [push]
            default = tracking

#### emacs
yum には 23.1 があったのでこれで妥協する.

    $ sudo yum install -y emacs

`.emacs` もおこのみで

#### nginx

    $ sudo yum install -y nginx.x86_64 --enablerepo=epel
    $ sudo chkconfig nginx on
    $ sudo /etc/init.d/nginx start

#### mysql

    $ sudo yum install -y mysql-server --enablerepo=remi
    $ sudo chkconfig mysqld on
    $ sudo /etc/init.d/mysqld start
    $ sudo /usr/bin/mysql_secure_installation

以下の設定をデフォルトの my.cnf にマージ

    [mysqld]
    character-set-server=utf8
    skip-character-set-client-handshake
    
    [client]
    default-character-set=utf8
    
    [mysql]
    default-character-set=utf8
    
    [mysqldump]
    default-character-set=utf8

restart

    $ sudo /etc/init.d/mysqld restart

#### memcahced

    $ sudo yum install -y memcached --enablerepo=remi
    $ sudo chkconfig memcached on
    $ sudo /etc/init.d/memcached start

### python
2.7.3 を自前で入れる.

    $ curl -O http://www.python.org/ftp/python/2.7.3/Python-2.7.3.tgz
    $ tar -zxvf Python-2.7.3.tgz
    $ cd Python-2.7.3
    $ ./configure --prefix=/usr/local
    $ make
    $ sudo make altinstall

distribute

    $ curl -O http://pypi.python.org/packages/source/d/distribute/distribute-0.6.34.tar.gz
    $ tar -zxvf distribute-0.6.34h.tar.gz
    $ cd distribute-0.6.34
    $ sudo /usr/local/bin/python2.7 setup.py install

virtualenv

    $ sudo /usr/local/bin/easy_install-2.7 virtualenv
    $ sudo /usr/local/bin/easy_install-2.7 virtualenvwrapper

.bashrc に設定

    ## python env
    export WORKON_HOME=$HOME/.virtualenvs
    export VIRTUALENVWRAPPER_PYTHON=/usr/local/bin/python2.7
    . virtualenvwrapper.sh
    
    mkvenv () {
        base_python=`which python$1`
        mkvirtualenv --distribute --python=$base_python $2
    }

### perl
perlbrew, cpanm, carton な構成で. carton は開発版の 0.9.7 にする

    $ curl -kL http://install.perlbrew.pl | bash
    $ source ~/perl5/perlbrew/etc/bashrc
    $ perlbrew install -n perl-5.16.2
    $ perlbrew use 5.16.2
    $ perlbrew install-cpanm
    $ cpanm MIYAGAWA/carton-v0.9_7.tar.gz

### node

    $ git clone git://github.com/creationix/nvm.git ~/nvm
    $ . ~/nvm/nvm.sh
    $ nvm install v0.8.17
    $ nvm install v0.9.6

### snapshot

ここでスナップショットをとっておく

Virtual Machine > Snapshots > Take Snapshot
