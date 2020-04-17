{"title":"Mavericks で Local Items のパスワードを何度も尋ねられる","date":"2013-12-30T23:20:59+09:00","tags":["mac"]}

Mac (Mavericks) を起動すると "Local Items" という項目の keychain パスワード尋ねるプロンプトが何度も立ち上がる問題があった。admin のパスワードを入れても違うと出るし、キャンセルしても何度も出続ける。"Keychain Access" アプリで first aid をしても修復できない。

[Mavericks Keychain keeps asking to...: Apple Support Communities](https://discussions.apple.com/thread/5467304?tstart=0)

こちらのフォーラムの投稿のとおりにするとなんとなく解決したっぽい。

- Keychain のアプリをひらく
- 左カラム login の keychain を選択
- 左上の南京錠アイコンをクリック。ロック・アンロックを一度ずつ行う
- iCloud Keychain の設定プロンプトがでる
- そこで新旧のパスワードを設定
- mac を再起動

別の記事でこういう情報も合ったが試していない。"~/Library/Keychains/" のファイルを消すという作業。また同様の問題が起こったら試したい。

[OS X Mavericks v10.9.1: Repeated prompts to unlock "Local Items" keychain](http://support.apple.com/kb/TS5362)

