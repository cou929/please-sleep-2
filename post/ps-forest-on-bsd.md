{"title":"mac の ps コマンドで --forest したい","date":"2012-12-09T18:50:56+09:00","tags":["nix"]}

linux だと `ps auxf` or `ps aux --forest` でプロセスの親子関係をツリー状に表示してくれるが, mac (確認はしてないけどたぶん bsd も) に入ってる px にはそのオプションが無い. 同等のオプションはなくて, `pstree` コマンドをインストールするしかなさそうだ.

homebrew にあるので, mac & homebrew の人はそれでOK

    $ brew install pstree

出る情報は `ps auxf` とは同じではない.

    $ pstree
    -+= 00001 root /sbin/launchd
     |--= 00010 root /usr/libexec/kextd
     |--= 00011 root /usr/libexec/UserEventAgent -l System
     |--= 00012 _mdnsresponder /usr/sbin/mDNSResponder -launchd
     |--= 00013 root /usr/libexec/opendirectoryd
     |--= 00014 root /usr/sbin/notifyd
     |--= 00015 root /usr/sbin/diskarbitrationd
     |--= 00016 root /usr/libexec/configd
     |--= 00017 root /usr/sbin/syslogd
