{"title":"browserstack で hosts を有効にしつつローカルファイルの動作チェック","date":"2013-05-18T21:43:35+09:00","tags":["web"]}

[Cross Browser Testing Tool. 200+ Browsers, Mobile, Real IE.](http://www.browserstack.com/)

ローカルにある html ファイルの動作確認を browserstack でしたい。その html はサードパーティの js をロードしているが、向き先を hosts で書き換えたい。今回に関しては hosts でサードパーティ js をローカルに向け、ダミーのスクリプトに差し替えて動作確認をする。

こういうテストをやる手順。

1. ローカルに web サーバを立てておく。80 を listen できるように。
   - Mac の場合付属の apache をあげるのが楽だと思う。
2. テストしたいファイル・スクリプトを見えるようにしておく。
   - 今回の場合テストしたい html、スクリプト、差し替えたいダミースクリプトを `/Library/WebServer/Documents/` に置く。
3. ローカルマシンの hosts で、必要な host の向き先を変えておく。
   - 今回の場合サードパーティ js のホストをローカルに向けておく

            127.0.0.1 tp1.sample.com
            127.0.0.1 tp2.sample.com

4. browserstack にログイン
   - [Sign into the Best Browser Testing Tool](https://www.browserstack.com/users/sign_in)
5. トンネリングのために必要な jar ファイルをダウンロード
   - BrowserStackTunnel.jar
   - [Test local and internal servers](http://www.browserstack.com/local-testing) を参照
6. この jar を使ってトンネルを設定
   - KEY はログイン後に左下の `command line` をクリックして出てくるモーダルに書いてある
   - オプションで localhost と hosts で向き先を変えたいドメインを指定する

            $ java -jar BrowserStackTunnel.jar <KEY> localhost,80,0,tp1.sample.com,80,0,tp2.sample.com,80,0

あとは browserstack の UI からテストしたい OS / Browser を選べばよい。`http://localhost/test.html` というかんじでローカルの (apache でサーブしている) ファイルにアクセスできる。hosts を書き換えたものもローカルに向いている。

browserstack にも API があり、わざわざ UI を操作せずテストを自動化できるらしいので、次は試してみたい。
