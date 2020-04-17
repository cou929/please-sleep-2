{"title":"SPF について勉強","date":"2012-10-02T23:41:58+09:00","tags":["memo"]}

Sender Policy Framework. メール送信者の認証方法のひとつ. 正しい送信者であることを認証してスパム・フィッシングを防ぐのが目的

### 方式

メール受信側のクライアントが以下の方法で認証を行う.

1. メールヘッダの from から送信元のドメインを取り出す
2. 上記ドメインの dns の spf レコードを問い合わせる
3. そのドメインの spf レコードに定義されている ip にメール送信元の ip アドレスが含まれるかをチェックする.

### spf レコードの指定例

- `example.com` への指定は, すべての ip が無効になる. このドメインからはメール送信はおこなっていないという表明.
- 残りの 2 つは `192.0.2.1` は valid. それ以外は無効という指定方法.
- all の前の `-` は, 当該ドメインの送信メールサーバとして認証しないという意味になる. `+` は逆に認証するという意味. `~` は `-` と `+` の中間で, 積極的に認証されてはいないが即座に受信拒否すべきでないという意味になる.

        example.com. IN TXT "v=spf1 -all"
        example.net. IN TXT "v=spf1 ip4:192.0.2.1 -all"
        example.jp.  IN SPF "v=spf1 ip4:192.0.2.1 -all"

### 認証結果

- Pass
  - 認証 OK
- Fail
  - 認証 ng. spf レコードの `-` に合致した場合
- SoftFail
  - spf レコードの `~` に合致した場合. 注意
- Neutral
  - spf レコードの `?` に合致した場合. None に近く, SoftFail よりは正しさが上として扱う
- None
  - 当該ドメインの dns が spf 情報を公開していない

### やってみる

例えば gmail の Web UI だとメニューの `show original` からメールヘッダ込みの全文が見られる. 通常のメールの場合 spf 関連の項目は以下のように `pass` となっている

    Received-SPF: pass (google.com: domain of z622b6c93epbh929+f=tznvy.pbz@postmaster.twitter.com designates 199.59.148.236 as permitted sender) client-ip=199.59.148.236;
    Authentication-Results: mx.google.com; spf=pass (google.com: domain of z622b6c93epbh929+f=tznvy.pbz@postmaster.twitter.com designates 199.59.148.236 as permitted sender) smtp.mail=z622b6c93epbh929+f=tznvy.pbz@postmaster.twitter.com; dkim=pass header.i=@twitter.com

例えば手元のサーバから `gmail.com` や `yahoo.co.jp` などを from に指定してメール送信すると `neutral` や `softfail` になる

    // 送信例
    $ echo 'test' | mail -s 'test mail' foo@gmail.com -- -f 'test@gmail.com'
    $ echo 'test' | mail -s 'test mail' foo@gmail.com -- -f 'test@yahoo.co.jp'

    // 受信例
    Received-SPF: neutral (google.com: 49.212.15.82 is neither permitted nor denied by domain of test@gmail.com) client-ip=49.212.15.82;
    Received-SPF: softfail (google.com: domain of transitioning test@yahoo.co.jp does not designate 49.212.15.82 as permitted sender) client-ip=49.212.15.82;

gmail の場合は spf 認証結果をあくまでメールの扱いを決める位置変数として扱っているようだ. 上記のように pass していなくても受信自体はできていて inbox に格納される. ドメインによっては以下のようにフィッシングの警告が出る.

![gmail フィッシング警告](http://gyazo.com/90b55d7d62b20810f29c973b2d8449ce.png?1349187415)

ほかの `不正らしさ` の判定結果 (例えばそのドメインから日常的に大量のメールが送られてきている, など) によってはスパム判定されたり, そもそも受信拒否されるなどの挙動も見られた.


### Ref.

- [SPF（Sender Policy Framework） : 迷惑メール対策委員会](http://salt.iajapan.org/wpmu/anti_spam/admin/tech/explanation/spf/)
