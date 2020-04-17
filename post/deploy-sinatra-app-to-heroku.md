{"title":"heroku で Sinatra のアプリを動かす","date":"2013-12-20T21:55:43+09:00","tags":["ruby"]}

- [アカウントをつくる](https://www.heroku.com/)
- [heroku toolbelt](https://toolbelt.heroku.com/) をインストールする
  - インストール後 heroku コマンドが使えるようになる

            $ heroku --version
            heroku-toolbelt/3.2.1 (x86_64-darwin10.8.0) ruby/1.9.3

- sinatra のアプリを書く
  - Gemfile を準備する。モジュールをインストールする

            $ cat Gemfile
            source 'https://rubygems.org'
            gem 'sinatra'
            $ bundle install --path vendor/bundle

    - `vendor`、`.bundle` は gitignore する。Gemfile.lock はバージョン管理する。
  - プロジェクトのルートに app.rb などを作ってアプリを書く。テンプレートは views 以下におく。今回は erb を使う。

        <pre><code data-language="ruby"># app.rb
        
        require 'sinatra'
        
        get '/' do
          @text = 'hi'
          erb :index
        end</code></pre>

        <pre><code data-language="html"># views/index.erb
        
        &lt;!DOCTYPE html&gt;
        &lt;html&gt;
        &lt;body&gt;
        &lt;%= @text %&gt;
        &lt;/body&gt;
        &lt;/html&gt;</code></pre>

- `bundle exec ruby app.rb` で手元で動作確認できる
- config.ru ファイルを準備。よく理解していないけど、config.ru は rack アプリ起動時のエンドポイントみたいなものかな?

    <pre><code data-language="ruby">require 'bundler'
    Bundler.require
    
    require './app'
    run Sinatra::Application</code></pre>

- heroku のアプリとして初期化する。プロジェクトのルートで以下を実行。初回は id/pass や公開鍵の設置などの入力をプロンプトで求められるはず

        $ heroku create APP_NAME

  - こうすると git remote heroku が追加されている
- heroku にデプロイする。push するだけ。

        $ git push heroku master

- アプリを開く

        $ heroku open

- ログを見る

        $ heroku logs

### 雑感

paas 全般に言えることなんだけど、そのサービス固有の知識が求められたり特有の制限にひっかかったりするとやる気が萎えてしまう。90 % の時間は快適に使わせてもらっていても、ちょっとそういうことがあると実際以上にマイナスイメージをいだいてしまうような気がする。(無料でつかわせてもらっているのにせこい話だけど…)。AWS なり VPS を契約しているのなら、はじめからそっちを使って自分でやったほうが最終的にはやかったりするということもあると再確認することになった。

今回はでかいクッキーを食わせた場合のブラウザの挙動をしらべていて、単に任意の Set-Cookie ヘッダを発行するだけのアプリをインターネットから見えるところにおいておきたいというだけの要件だった。多くの web サーバはヘッダサイズのリミットを持っているが、当然ながらそんな設定値を触らせてくれるわけがないので、そもそも最初から自分の選択がおかしいという話だった。

ちなみにリクエストヘッダのサイズは heroku の場合 8KB くらいが限界のようだ。

[HTTP Routing | Heroku Dev Center](https://devcenter.heroku.com/articles/http-routing#request-buffering)

