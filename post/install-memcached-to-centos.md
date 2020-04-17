{"title":"CentOS 5.8 で memcached インストール","date":"2012-05-12T19:30:20+09:00","tags":["nix"]}

CentOS 5.8 で x86_64 な環境 (さくら vps) に memcached を入れる

    % cat /etc/redhat-release
    CentOS release 5.8 (Final)
    cou929:kosei% uname -a
    Linux cou929.nu 2.6.18-238.19.1.el5 #1 SMP Fri Jul 15 07:31:24 EDT 2011 x86_64 x86_64 x86_64 GNU/Linux

yum だと rpmforge-extras にあるので, まだの場合はリポジトリの設定をする

    $ rpm -Uhv http://apt.sw.be/redhat/el5/en/x86_64/rpmforge/RPMS//rpmforge-release-0.3.6-1.el5.rf.x86_64.rpm

ほかのアーキテクチャだとここを参照: [DAG: Frequently Asked Questions](http://dag.wieers.com/rpm/FAQ.php#B2)

rpmforge-extras に 1.4.7 がある

    % yum info memcached --enablerepo rpmforge-extras
    Loaded plugins: downloadonly, fastestmirror, priorities
    172 packages excluded due to repository priority protections
    Available Packages
    Name       : memcached
    Arch       : x86_64
    Version    : 1.4.7
    Release    : 1.el5.rfx
    Size       : 81 k
    Repo       : rpmforge-extras
    Summary    : Distributed memory object caching system
    URL        : http://memcached.org/
    License    : BSD
    Description: memcached is a high-performance, distributed memory object caching system,
               : generic in nature, but intended for use in speeding up dynamic web
               : applications by alleviating database load.

インストール

    % sudo yum install -y memcached

ここらへんのパッケージに依存している

     perl-AnyEvent                                    x86_64                             5.340-1.el5.rfx                                rpmforge-extras                             360 k
     perl-Async-Interrupt                             x86_64                             1.05-1.el5.rf                                  rpmforge                                     76 k
     perl-EV                                          x86_64                             4.03-1.el5.rf                                  rpmforge                                    376 k
     perl-Guard                                       x86_64                             1.021-1.el5.rf                                 rpmforge                                     39 k
     perl-JSON                                        noarch                             2.50-1.el5.rf                                  rpmforge                                    100 k
     perl-JSON-XS                                     x86_64                             2.30-1.el5.rf                                  rpmforge                                    152 k
     perl-Net-SSLeay                                  x86_64                             1.36-1.el5.rfx                                 rpmforge-extras                             334 k
     perl-TermReadKey                                 x86_64                             2.30-3.el5.rf                                  rpmforge                                     57 k
     perl-YAML                                        noarch                             0.72-1.el5.rf                                  rpmforge                                     84 k
     perl-common-sense                                x86_64                             3.0-1.el5.rf                                   rpmforge                                     22 k

Net::SSLeay は base リポジトリにも古いバージョンがあったりするので, yum-(plugin)-repositories をいれておいて rpmforge の priority を上げておかないと依存が解決できないと怒られるかもしれない

    % sudo cat /etc/yum.repos.d/rpmforge.repo                                                                                                                   [/home/kosei]
    ### Name: RPMforge RPM Repository for RHEL 5 - dag
    ### URL: http://rpmforge.net/
    [rpmforge]
    name = RHEL $releasever - RPMforge.net - dag
    baseurl = http://apt.sw.be/redhat/el5/en/$basearch/rpmforge
    mirrorlist = http://apt.sw.be/redhat/el5/en/mirrors-rpmforge
    #mirrorlist = file:///etc/yum.repos.d/mirrors-rpmforge
    enabled = 1
    protect = 0
    gpgkey = file:///etc/pki/rpm-gpg/RPM-GPG-KEY-rpmforge-dag
    gpgcheck = 1
    priority=1
    
    [rpmforge-extras]
    name = RHEL $releasever - RPMforge.net - extras
    baseurl = http://apt.sw.be/redhat/el5/en/$basearch/extras
    mirrorlist = http://apt.sw.be/redhat/el5/en/mirrors-rpmforge-extras
    #mirrorlist = file:///etc/yum.repos.d/mirrors-rpmforge-extras
    enabled = 1
    protect = 0
    gpgkey = file:///etc/pki/rpm-gpg/RPM-GPG-KEY-rpmforge-dag
    gpgcheck = 1
    priority=1
    
    [rpmforge-testing]
    name = RHEL $releasever - RPMforge.net - testing
    baseurl = http://apt.sw.be/redhat/el5/en/$basearch/testing
    mirrorlist = http://apt.sw.be/redhat/el5/en/mirrors-rpmforge-testing
    #mirrorlist = file:///etc/yum.repos.d/mirrors-rpmforge-testing
    enabled = 1
    protect = 0
    gpgkey = file:///etc/pki/rpm-gpg/RPM-GPG-KEY-rpmforge-dag
    gpgcheck = 1

バージョンチェック

    % memcached -h | head -n 1                                                                                                                                  [/home/kosei]
    memcached 1.4.7

設定確認

    % cat /etc/sysconfig/memcached
    PORT="11211"
    USER="nobody"
    MAXCONN="1024"
    CACHESIZE="64"
    OPTIONS="-l 127.0.0.1"

デーモン起動

    % sudo /etc/init.d/memcached strat

確認

    % ps aux | grep mem
    nobody    9989  0.0  0.1 140884   956 ?        Ssl  17:02   0:00 memcached -d -p 11211 -u nobody -c 1024 -m 64
    % telnet localhost 11211
    Trying 127.0.0.1...
    Connected to localhost.localdomain (127.0.0.1).
    Escape character is '^]'.
    stats
    STAT pid 10218
    STAT uptime 20
    STAT time 1336810604
    STAT version 1.4.7
    STAT libevent 1.4.13-stable
    STAT pointer_size 64
    STAT rusage_user 0.000000
    STAT rusage_system 0.000999
    STAT curr_connections 5
    STAT total_connections 6
    STAT connection_structures 6
    STAT cmd_get 0
    STAT cmd_set 0
    STAT cmd_flush 0
    STAT get_hits 0
    STAT get_misses 0
    STAT delete_misses 0
    STAT delete_hits 0
    STAT incr_misses 0
    STAT incr_hits 0
    STAT decr_misses 0
    STAT decr_hits 0
    STAT cas_misses 0
    STAT cas_hits 0
    STAT cas_badval 0
    STAT auth_cmds 0
    STAT auth_errors 0
    STAT bytes_read 7
    STAT bytes_written 0
    STAT limit_maxbytes 67108864
    STAT accepting_conns 1
    STAT listen_disabled_num 0
    STAT threads 4
    STAT conn_yields 0
    STAT bytes 0
    STAT curr_items 0
    STAT total_items 0
    STAT evictions 0
    STAT reclaimed 0
    END

chkconfig

    % sudo /sbin/chkconfig --add memcached
    % sudo /sbin/chkconfig --list | grep mem
    memcached       0:off   1:off   2:on    3:on    4:on    5:on    6:off
