{"title":"ChangeLog メモを Posterous に 1 コマンドで投稿する","date":"2012-02-25T14:30:52+09:00","tags":["tips"]}

普段メモは [ChangeLog](http://0xcc.net/unimag/1/) (+ Dropbox + markdown) でとっています. ChangeLog メモを web 公開するツールは [chalow](http://chalow.org/) が有名ですが, ここでは Posterous にポストして公開する方法について説明します.

### 概要

公開したいメモをメールで posterous に送信するだけです

### スクリプト

スクリプトが 3 つ登場します.

- [clgrep.py](https://gist.github.com/959428)
  - ChangeLog メモをメモごとに grep するツール. もとは [高林さんが公開している ruby スクリプト](http://0xcc.net/unimag/1/). python に移植して, タブを削除するオプションなどを追加した.

- [sendgmail.py](https://gist.github.com/1582468)
  - gmail の smtp を使ってメールを送信するスクリプト. コマンドラインで posterous に登録してあるメールアドレスでメールが送れれば何でもいい. 自分でメールサーバを立ち上げるのは嫌だったので, gmail を使わせてもらうことにした

- [changelog2posterous.rb](https://gist.github.com/1580812)
  - clgrep.py の出力を受け取って posterous にメールを送るスクリプト. ChangeLog メモは markdown で書かれていることを前提としているので, markdown タグを追加する.

### 手順

clgrep.py の出力を changelog2posterous.rb に渡してあげます. sendgmail.py は changelog2posterous.rb の中で呼ばれています.

    $ clgrep.py 'search_ward' -n 1 -t | changelog2posterous.rb

パイプ左側の出力はこんな感じで, メモを抜き出して行頭タブを除いて出力しています.

    % clgrep.py 'javascript:' -n 5 -t
    * javascript: 正規表現のコンストラクタとリテラル
    それぞれ若干書き方が異なってしまう.
    
        #!javascript
        /foo\b/g;
        // equal
        new RegExp('foo\\b', 'g');
    
    バックスラッシュはエスケープが必要. モードもクオートで囲う

### その他
- 文字コードが依然わからなくて時間かかった
  - utf-8 (のはず) の文字列を iso-2022-jp にしてメールを送信したい
  - unicode(orig_str, 'utf-8').encode('iso-2022-jp')
  - こういうふうにしないとコーデックがなんとかと言われて怒られる

- posterous はメールサブジェクトのエンコーディングがおかしいと投稿を受け付けてくれないらしい
- ruby はノリでかけた
