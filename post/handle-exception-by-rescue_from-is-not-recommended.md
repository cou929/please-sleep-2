{"title":"Rails の rescue_from で Exception や StandardError をキャッチすることは推奨されていない","date":"2022-08-22T23:30:00+09:00","tags":["ruby", "rails"]}

rescue_from で `Exception` をキャッチすると予期しない挙動をして驚いたことがあった。Rails ガイドを見てみると以下のようにそのような使い方は非推奨のようだった。

[Action Controller Overview — Ruby on Rails Guides](https://guides.rubyonrails.org/action_controller_overview.html#rescue-from)

> Using rescue_from with Exception or StandardError would cause serious side-effects as it prevents Rails from handling exceptions properly. As such, it is not recommended to do so unless there is a strong reason.
>
> Exception や StandardError を指定すると Rails による正しい例外ハンドリングが阻害されて深刻な副作用が起こる可能性がある。強い理由がない限りはそうしないほうがよい。

## rescue_from のおさらい

[ActiveSupport::Rescuable::ClassMethods](https://edgeapi.rubyonrails.org/classes/ActiveSupport/Rescuable/ClassMethods.html#method-i-rescue_from) の通りで、例えば次のように記述すると、指定した例外を指定したハンドラで処理できる。`User::NotAuthorized` が発生すると deny_access が、`ActiveRecord::RecordInvalid` には show_record_errors が、`MyApp::BaseError` にはその場で登録している処理が対応する。

```ruby
class ApplicationController < ActionController::Base
  rescue_from User::NotAuthorized, with: :deny_access
  rescue_from ActiveRecord::RecordInvalid, with: :show_record_errors

  rescue_from "MyApp::BaseError" do |exception|
    redirect_to root_url, alert: exception.message
  end

  private
    def deny_access
      head :forbidden
    end

    def show_record_errors(exception)
      redirect_back_or_to root_url, alert: exception.record.errors.full_messages.to_sentence
    end
end
```

## 評価は下から上にされる

これは [ドキュメント](https://edgeapi.rubyonrails.org/classes/ActiveSupport/Rescuable/ClassMethods.html#method-i-rescue_from) にも記載されているが、登録したハンドラは下から上の順に検索され、例外がマッチしたらそこで終了する。マッチするかどうかは `exception.is_a?(klass)` でチェックする。

> Handlers are inherited. They are searched from right to left, from bottom to top, and up the hierarchy. The handler of the first class for which exception.is_a?(klass) holds true is the one invoked, if any.

例えば次のように記載すると、例外はまず [is_a](https://docs.ruby-lang.org/ja/latest/method/Object/i/is_a=3f.html) StandardError かどうか (StandardError 及びそのサブクラスかどうか) をチェックされる。ほとんどのケースでここでマッチするので、実質的に `:deny_access`, `:show_record_errors` を通ることは無くなる。よって範囲の広い例外をキャッチする場合は「上」に記載する必要がある。

```ruby
class ApplicationController < ActionController::Base
  rescue_from User::NotAuthorized, with: :deny_access
  rescue_from ActiveRecord::RecordInvalid, with: :show_record_errors
  rescue_from StandardError, with :other_error_handler

  # ...
end
```

[実装](https://github.com/rails/rails/blob/82bab92cfe9ab62793cc82d25e6662e4359352e9/activesupport/lib/active_support/rescuable.rb#L124-L127) は次のようになっている。

- rescue_handlers は `rescue_from` で記載したハンドラが、[登録順に入っている](https://github.com/rails/rails/blob/82bab92cfe9ab62793cc82d25e6662e4359352e9/activesupport/lib/active_support/rescuable.rb#L70)
- これを `reverse_each` で操作している。つまり「下から上へ」調べていくという挙動になっている

```ruby
def find_rescue_handler(exception)
  if exception
    # Handlers are in order of declaration but the most recently declared
    # is the highest priority match, so we search for matching handlers
    # in reverse.
    _, handler = rescue_handlers.reverse_each.detect do |class_or_name, _|  # ここで reverse_each している
      if klass = constantize_rescue_handler_class(class_or_name)
        klass === exception
      end
    end

    handler
  end
end
```

## マッチする例外がない場合 cause を取り出して再度評価する

知る限りはドキュメントには記載がないが、発生した例外にマッチするハンドラがなかった場合 [cause](https://docs.ruby-lang.org/ja/latest/method/Exception/i/cause.html) でラップされた元の例外を取り出して再度すべてのハンドラを調査するという挙動をする。

例えば MySQL を使っているアプリケーションで DB インスタンスとの接続エラーが起きた場合を考える。その場合、まずは `Mysql2::Error::ConnectionError` が投げられるが、ActiveRecord の ConnectionAdapters は [ActiveRecord::ConnectionNotEstablished](https://github.com/rails/rails/blob/82bab92cfe9ab62793cc82d25e6662e4359352e9/activerecord/lib/active_record/connection_adapters/mysql2_adapter.rb#L39) でそれをラップしている。

ここでコントローラでは次のように `Mysql2::Error::ConnectionError` を捕捉しているとする。

```ruby
class ApplicationController < ActionController::Base
  rescue_from User::NotAuthorized, with: :deny_access
  rescue_from ActiveRecord::RecordInvalid, with: :show_record_errors
  rescue_from Mysql2::Error::ConnectionError, with: :show_internal_server_error

  # ...
end
```

その場合次のような挙動になる。

- まず `ActiveRecord::ConnectionNotEstablished` に対して対応するハンドラを探す
- マッチするものがないので `cause` を呼び出し、`Mysql2::Error::ConnectionError` が得られる
- `Mysql2::Error::ConnectionError` に対して、再度ハンドラを探す
- 最初の `:show_internal_server_error` が該当するのでこちらに処理が任される

[実装](https://github.com/rails/rails/blob/82bab92cfe9ab62793cc82d25e6662e4359352e9/activesupport/lib/active_support/rescuable.rb#L98) は次のようになっている。

- `visited_exceptions` に確認済みの例外をメモしながらハンドラをチェックしていく
- ハンドラが見つからず、かつ `exception.cause` がチェック済みでない場合、再帰的に rescue_with_handler を呼び出して再度はじめから検索している


```ruby
def rescue_with_handler(exception, object: self, visited_exceptions: [])
  visited_exceptions << exception

  if handler = handler_for_rescue(exception, object: object)
    handler.call exception
    exception
  elsif exception
    if visited_exceptions.include?(exception.cause)
      nil
    else
      rescue_with_handler(exception.cause, object: object, visited_exceptions: visited_exceptions)  # 見つからなかった場合は exception.cause を取り出して再帰的に rescue_with_handler を呼び出す
    end
  end
end
```

## Exception を指定するとどうなるか

以上の挙動を踏まえると、Rails ガイドが注意喚起しているように、下手に Exception や StandardError を rescue_from しようとすると、意図しない挙動につながってしまう恐れがあるため、よく注意する必要がある。

### ハンドラの登録順

まずは rescue_from でハンドラを登録する順番に注意する必要がある。前述のように「下」で Exception を指定するとすべての例外がそのハンドラに渡ってしまい、`:deny_access`、`:show_record_errors` を通ることは無くなってしまう。

```ruby
class ApplicationController < ActionController::Base
  rescue_from User::NotAuthorized, with: :deny_access
  rescue_from ActiveRecord::RecordInvalid, with: :show_record_errors
  rescue_from StandardError, with :other_error_handler  # ほとんどの例外がここに吸収される

  # ...
end
```

設定するとしたら「上」に持っていく必要がある。記載順が変わるだけで容易に壊れてしまう恐れがあるので、今後のメンテナンス性にも注意する必要がある。


```ruby
class ApplicationController < ActionController::Base
  rescue_from Exception, with :other_error_handler  # 広い範囲を補足したい場合は「一番上」に
  rescue_from User::NotAuthorized, with: :deny_access
  rescue_from ActiveRecord::RecordInvalid, with: :show_record_errors

  # ...
end
```

### ラップされた例外

また発生した例外がどうラップされているかによって挙動が意図せず変わってしまう問題もある。ラップされている例外が cause で取り出されていた挙動が、Exception のハンドラを登録することで必ず起こらなくなる。

前述の例のように `Mysql2::Error::ConnectionError` を補足しようとしている場合、データベースとの接続エラーが発生すると `:show_internal_server_error` がそれを処理することになる。

```ruby
class ApplicationController < ActionController::Base
  rescue_from User::NotAuthorized, with: :deny_access
  rescue_from ActiveRecord::RecordInvalid, with: :show_record_errors
  rescue_from Mysql2::Error::ConnectionError, with: :show_internal_server_error

  # ...
end
```

ここで次のように `Exception` のハンドラも追加したとする。

```ruby
class ApplicationController < ActionController::Base
  rescue_from Exception, with :other_error_handler  # 1 周目の検索で必ずここにマッチする
  rescue_from User::NotAuthorized, with: :deny_access
  rescue_from ActiveRecord::RecordInvalid, with: :show_record_errors
  rescue_from Mysql2::Error::ConnectionError, with: :show_internal_server_error  # cause で取り出した 2 週目の検索でここにマッチしていた

  # ...
end
```

この状態で接続エラーが発生すると、これまでとは異なり `:other_error_handler` に `ActiveRecord::ConnectionNotEstablished` が渡ってくる挙動になる。

- これまでは 1 周目の rescue_handlers の検索で `ActiveRecord::ConnectionNotEstablished` にマッチするハンドラが見つからず、cause で `Mysql2::Error::ConnectionError` が取り出され、2 周目で `:show_internal_server_error` にマッチしていた
- 変更後は 1 周目で `ActiveRecord::ConnectionNotEstablished` が `:other_error_handler` にマッチしそこで処理されるように変わった

注意してハンドラを登録すれば問題は避けられる。ただ投げられるすべての例外がどうラップされているのかを把握するのは現実的ではなく、思いもよらない挙動につながってしまう恐れは今後も残り続けてしまう。

## まとめ

- rescure_from は以下の挙動をする
    - 登録したハンドラの逆順で評価される
    - 例外がどのハンドラにもマッチしなかった場合、cause で元の例外を取り出して再度検索する
- そのため Exception や StandardError を rescue_from すると意図しない挙動になってしまう恐れがある
    - Exception, StandardError に対応するハンドラに意図せずに処理が寄ってしまうことがある
- 落とし穴を避けて実装することもできるが、変更に弱いコードになってしまうリスクがある
    - rescue_from を書く位置や例外の内容によって、容易に意図しない挙動になってしまうリスクがある

## 参考

- [Action Controller Overview — Ruby on Rails Guides](https://guides.rubyonrails.org/action_controller_overview.html#rescue-from)
- [ActiveSupport::Rescuable::ClassMethods](https://edgeapi.rubyonrails.org/classes/ActiveSupport/Rescuable/ClassMethods.html#method-i-rescue_from)
- [rails/rescuable\.rb at 82bab92cfe9ab62793cc82d25e6662e4359352e9 · rails/rails](https://github.com/rails/rails/blob/82bab92cfe9ab62793cc82d25e6662e4359352e9/activesupport/lib/active_support/rescuable.rb)
- [Object\#is\_a? \(Ruby 3\.1 リファレンスマニュアル\)](https://docs.ruby-lang.org/ja/latest/method/Object/i/is_a=3f.html)
- [Exception\#cause \(Ruby 3\.1 リファレンスマニュアル\)](https://docs.ruby-lang.org/ja/latest/method/Exception/i/cause.html)

## PR

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B08D3DW7LP/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/51hCV4olz-L.jpg" alt="パーフェクト Ruby on Rails　【増補改訂版】" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B08D3DW7LP/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">パーフェクト Ruby on Rails　【増補改訂版】</a></div><div class="amazlet-detail">by すがわら まさのり  (著), 前島 真一  (著), 橋立 友宏 (著), 五十嵐 邦明  (著), 後藤 優一 (著)  Format: Kindle Edition<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B08D3DW7LP/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
