{"title":"さくら VPS の CentOS6 でメールを送信できるようにする","date":"2014-01-03T16:25:31+09:00","tags":["nix"]}

送信だけできるよう postfix を設定する。postfix は最初からインストールされていたっぽいが、なければ `sudo yum install postfix` でいいはず。

`/etc/postfix/main.cf` を編集する。

    ## すでにある項目を編集
    myhostname = foo.bar.com
    mydomain = foo.bar.com
    myorigin = $mydomain

    ## ファイル末尾に追記

    # SMTP-Auth configuration
    smtpd_sasl_auth_enable = yes
    smtpd_sasl_local_domain = $myhostname
    smtpd_recipient_restrictions =
        permit_mynetworks
        permit_sasl_authenticated
        reject_unauth_destination
    
    # limit
    message_size_limit = 10485760

saslauthd の起動 & 自動起動設定。

    $ sudo /etc/rc.d/init.d/saslauthd start
    $ sudo /sbin/chkconfig saslauthd on
    $ sudo /sbin/chkconfig --list saslauthd
    # saslauthd       0:off   1:off   2:on    3:on    4:on    5:on    6:off  のようになっていることを確認

postfix の起動 & 自動起動設定。

    $ sudo /etc/rc.d/init.d/postfix start
    $ sudo /sbin/chkconfig postfix on
    $ sudo /sbin/chkconfig --list postfix
    # postifx       0:off   1:off   2:on    3:on    4:on    5:on    6:off  のようになっていることを確認

### 参考

- [さくらVPSで、メールを送信出来るようにする - ちゃまぐの備忘録](http://d.hatena.ne.jp/tyamaguc07/20110413/p1)

