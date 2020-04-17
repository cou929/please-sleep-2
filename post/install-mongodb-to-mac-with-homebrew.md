{"title":"homebrew で mac に mongodb を入れる","date":"2012-11-14T22:50:38+09:00","tags":["mac"]}

悩むことはない

1. brew update
2. brew install mongodb
3. ガイドにも出てるけど, ログイン時に自動でスタートさせたい場合は

        ln -s /usr/local/opt/mongodb/*.plist ~/Library/LaunchAgents/

4. 後は $ mongod で OK

ログは /usr/local/var/log/mongodb/mongo.log に出る
