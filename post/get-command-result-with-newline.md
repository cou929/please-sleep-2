{"title":"シェルスクリプトで実行したコマンドの結果の改行をそのままにする","date":"2012-12-09T18:49:16+09:00","tags":["nix"]}

bash 系の話だけど, `$()` とかバッククオートでコマンドを実行してその標準出力を得るさいに, そのままだと改行がなくなってしまう. そんな時は全体をダブルクオートでくくれば改行を残せる

```
$ echo $(ls -l /tmp/)
total 56 -rw------- 1 kosei wheel 849 12 7 09:13 Config-pAimbL -rw-r--r-- 1 kosei wheel 23020 12 4 22:40 debug.log drwx------ 4 kosei wheel 136 12 6 00:16 launch-8y9r7a drwx------ 3 kosei wheel 102 12 1 16:28 launch-O688DH drwx------ 3 kosei wheel 102 12 1 16:28 launch-Obe9Hd drwx------ 3 kosei wheel 102 12 1 16:28 launch-mH6LiV drwx------ 3 kosei wheel 102 12 1 16:28 launchd-156.GLXAJQ srwxrwxrwx 1 kosei wheel 0 12 1 16:29 mongodb-27017.sock drwxr-xr-x 67 kosei wheel 2278 12 4 22:53 please-sleep
```
```
bash-3.2$ echo "$(ls -l /tmp/)"
total 56
-rw-------   1 kosei  wheel    849 12  7 09:13 Config-pAimbL
-rw-r--r--   1 kosei  wheel  23020 12  4 22:40 debug.log
drwx------   4 kosei  wheel    136 12  6 00:16 launch-8y9r7a
drwx------   3 kosei  wheel    102 12  1 16:28 launch-O688DH
drwx------   3 kosei  wheel    102 12  1 16:28 launch-Obe9Hd
drwx------   3 kosei  wheel    102 12  1 16:28 launch-mH6LiV
drwx------   3 kosei  wheel    102 12  1 16:28 launchd-156.GLXAJQ
srwxrwxrwx   1 kosei  wheel      0 12  1 16:29 mongodb-27017.sock
drwxr-xr-x  67 kosei  wheel   2278 12  4 22:53 please-sleep
```

<iframe style="width:120px;height:240px;" marginwidth="0" marginheight="0" scrolling="no" frameborder="0" src="//rcm-fe.amazon-adsystem.com/e/cm?lt1=_blank&bc1=000000&IS2=1&bg1=FFFFFF&fc1=000000&lc1=0000FF&t=pleasesleep-22&language=ja_JP&o=9&p=8&l=as4&m=amazon&f=ifr&ref=as_ss_li_til&asins=B07JGYV4Q8&linkId=f9738b0d0792b41b1652a3e05580b7ee"></iframe>
