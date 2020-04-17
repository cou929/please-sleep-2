{"title":"CentOS6 でタイムゾーンを変更する","date":"2014-02-18T19:19:40+09:00","tags":["nix"]}

`/etc/localtime` がタイムゾーンを司るファイル。元となるファイル地域別に `/usr/share/zoneinfo/` 以下にあるので、適切なものを指すようにしてあげればよい。

Japan にする場合は、こんなかんじでOK

    $ sudo rm /etc/localtime
    $ sudo ln -s /usr/share/zoneinfo/Japan /etc/localtime

