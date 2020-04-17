{"title":"linux で apache のツールなしに htpasswd ファイルをつくる","date":"2012-12-03T00:34:26+09:00","tags":["nix"]}

nginx などで basic 認証をかけたい時にいちいちそれだけのために htpasswd を持ってくるのもださいから, 次のワンライナーで htpasswd ファイル用のエントリを生成できるのでこれを使うと良い.

    $ printf "John:$(openssl passwd -crypt V3Ry)\n" >> .htpasswd

これは `John` というユーザ名で `V3Ry` というパスワードのエントリを作って `.htpasswd` に追加している例.

[How do I generate an .htpasswd file without having Apache tools installed?](http://wiki.nginx.org/Faq#How_do_I_generate_an_.htpasswd_file_without_having_Apache_tools_installed.3F)
