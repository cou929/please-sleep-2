{"title":"Sphinx のドキュメントを Travis CI でビルドし S3 にデプロイ・配信する","date":"2014-09-15T12:55:00+09:00","tags":["nix"]}

いままでさくらの VPS で配信していた [Google JavaScript Style Guide](http://cou929.nu/data/google_javascript_style_guide/) と [How Browsers Work](http://cou929.nu/docs/how-browsers-work/)。これらは [Sphinx](http://sphinx-doc.org/index.html) で書かれていて静的ファイルだけなので、S3 から配信することにした。ついでにビルドとデプロイを [Travis](https://travis-ci.org/) にやらせて、リポジトリにプッシュすると最新版をビルドして S3 へデプロイしてくれる仕組みにした。

### Travis CI で Sphinx のドキュメントをビルドする

まずはいままで自前のスクリプトで頑張っていたドキュメントのビルドを Travis に任せるところからはじめる。

`.travis.yml` に次の内容を記述する。

<pre><code data-language="generic">language: python
python:
- '2.7'
install: pip install -q -r requirements.txt
script: sphinx-build -nW -b html -d _build/doctrees . _build/html</code></pre>

ふつうに python のプロジェクトとして扱う。

`requirements.txt` の中身は Sphinx への依存を一行書いてあるだけ。バージョンは適宜変更する。

    Sphinx==1.2.3

ビルドコマンドには `Makefile` の内容に `-n` と `-W` のオプションを追加している。`-n` はターゲットのないリファレンスがあった場合に警告するオプション。`-W` は warning があった場合もエラーとして扱い、コマンドの戻り値を 0 以外にするオプション。これで警告があった場合にはビルドが落ちるようになる。

言うまでもないけれど、`.travis.yml` と `requirements.txt` をリポジトリに追加したあとは [Travis の設定画面](https://travis-ci.org/profile/) から対象のリポジトリを有効化する。以降 push や p-r ごとにビルドが走るようになる。

### S3 で静的ファイルを配信できるようにする

次に S3 から静的ファイルを配信できるよう設定する。これに関しては次のスライドで 100 % 解決する。

[Amazon S3による静的Webサイトホスティング](http://www.slideshare.net/horiyasu/amazon-s3web-27138902)

- S3 上の静的ファイルの配信
- 一部ドキュメントのリダイレクト
- 自ドメインでの配信 (Route53 使用)
- アクセスログの取得
- アクセスログのローテーション (Glacier へのアーカイブや消し込みなど)
- ルートドメインでのアクセス

といったように、必要なことはすべてこのスライドで説明されている。

しいて補足すると、自ドメインで配信する場合はバケット名をそのドメイン名にする必要がある。たとえば今回だと `cou929.nu` というバケット名になる。SSL で S3 へ直接アクセスする URL は `https://cou929.nu.s3-ap-northeast-1.amazonaws.com/` といったものになるが、その際に証明書のエラーが出てしまう。S3 の SSL 証明書は `*.s3-ap-northeast-1.amazonaws.com` というワイルドカードなので、バケット名にドットが入るとエラーが出てしまうようだ。

自分の場合は [Cyberduck](http://cyberduck.io/) というクライアントを使っているが、該当のバケットにアクセスすると警告がでる。SSL 経由でクライアントライブラリなどから接続する場合も同様だろう。

自ドメインでの配信をしたい場合はドットを避けられないので、とりあえず警告を無視してしまうのが一番手っ取り早そうだった。Cyberduck の場合はエラーダイアログの詳細表示から「以降は無視する」チェックボックスをチェックすればいいし、その他のスクリプトでも無視するオプション、例えば `curl` でいう `-k`、をつければいい。

### Travis CI でビルド後に S3 へデプロイする

最後に Travis からのデプロイの設定をする。

まずはデプロイに使う IAM を発行する。[Management Console](https://console.aws.amazon.com/iam/home#home) から適当なユーザーを作成する。そのユーザーの Permission として S3 だけへのアクセス許可を与える。該当ユーザーの詳細ページから `Permissions`、`Attach User Policy` ボタンをクリック、問題なければ `Amazon S3 Full Access` のテンプレートをそのまま使ってしまう。

次にアクセスキーを発行する。`Security Credentials` の項目から `Manage Access Keys`、`Create Access Key` とクリックしていく。アクセスキーとシークレットキーは発行時にしか参照できないので、ダウンロードしておくなりする。

AWS 側の設定はこれで完了。次に Travis 側の設定を行う。

まず、`travis` コマンドがない場合はインストールしておく。

    $ gem install travis

`.travis.yml` をパブリックなリポジトリに保存することになるので、すくなくともシークレットキーは暗号化して記載したい。[Travis が提供している暗号化の仕組み](http://docs.travis-ci.com/user/encryption-keys/) があるので、まずはこれを使って設定する。

リポジトリのルートで次のように `travis encrypt` というコマンドで設定する。一応アクセスキーも暗号化する例になっている。

    $ travis encrypt --add deploy.access_key_id 'YOUR_ACCESS_KEY'
    $ travis encrypt --add deploy.secret_access_key 'YOUR_SECRET_KEY'

このように `--add` オプションをつけると結果を `.travis.yml` に追記してくれる。

あとは以下のように各項目を設定していけばよい。


<pre><code data-language="generic">deploy:
  # s3 固定
  provider: s3

  # アクセスキー。自動で挿入されている
  access_key_id:
    secure: ...

  # シークレットキー。自動で挿入されている
  secret_access_key:
    secure: ...

  # 対象のバケット
  bucket: cou929.nu

  # ビルド後のクリーンアップを抑制する  
  skip_cleanup: true

  # エンドポイント。静的ファイル配信設定がされているバケットの場合は以下のように設定する必要がある
  endpoint: cou929.nu.s3-website-ap-northeast-1.amazonaws.com

  # リージョン。これは Tokyo リージョンの例
  region: ap-northeast-1

  # Travis 側のビルド結果のディレクトリ。この内容を S3 に put する
  local-dir: _build/html/

  # S3 側のディレクトリ。このディレクトリに put する。
  upload-dir: docs/how-browsers-work</code></pre>

リージョンとエンドポイントの内容については [AWS のドキュメント](http://docs.aws.amazon.com/general/latest/gr/rande.html) に記述がある。

ちなみに `travis setup s3` というコマンドを使うと、対話形式でこのような設定のテンプレートを挿入してくれる。

この他にも、タグをうったときにだけデプロイする `tags: true` という設定など、痒いところに手が届く仕組みになっている。詳しくは公式のドキュメントを参照。

[Travis CI: S3 Deployment](http://docs.travis-ci.com/user/deployment/s3/)

### 参考

- [Amazon S3による静的Webサイトホスティング](http://www.slideshare.net/horiyasu/amazon-s3web-27138902)
- [Travis CI: S3 Deployment](http://docs.travis-ci.com/user/deployment/s3/)
- [Mike Fiedler : Have Travis-CI test your Sphinx docs](https://coderwall.com/p/wws2uq)
- [code.rock: S3のバケット名はよく考えて命名しましょう！](http://blog.dateofrock.com/2012/02/s3.html)

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00MGGW3MY/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="http://ecx.images-amazon.com/images/I/51p-m4Jj1tL._SL160_.jpg" alt="Amazon Web Services クラウドデザインパターン実装ガイド（日経BP Next ICT選書）" style="border: none;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00MGGW3MY/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon Web Services クラウドデザインパターン実装ガイド（日経BP Next ICT選書）</a><div class="amazlet-powered-date" style="font-size:80%;margin-top:5px;line-height:120%">posted with <a href="http://www.amazlet.com/" title="amazlet" target="_blank">amazlet</a> at 14.09.18</div></div><div class="amazlet-detail">日経BP社 (2014-08-08)<br />売り上げランキング: 1,294<br /></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00MGGW3MY/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
