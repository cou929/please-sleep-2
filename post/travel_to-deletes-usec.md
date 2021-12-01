{"title":"travel_to はミリ秒以下の情報を切り捨てる","date":"2021-12-02T02:50:00+09:00","tags":["ruby"]}

次のように渡した時刻のミリ秒以下の精度の情報は切り捨てられる。

```ruby
[1] pry(main)> require 'active_support/testing/time_helpers'
=> true
[2] pry(main)> include ActiveSupport::Testing::TimeHelpers
=> Object
[3] pry(main)> t = Time.now
=> 2021-12-01 17:32:59.691290907 +0000
[4] pry(main)> t.to_f
=> 1638379979.6912909
[5] pry(main)> travel_to t
=> nil
[6] pry(main)> Time.now
=> 2021-12-01 17:32:59 +0000
[7] pry(main)> Time.now.to_f
=> 1638379979.0                   # <== [4] と比較して小数部がなくなっている
```

実際以下で usec を 0 に設定している。

https://github.com/rails/rails/blob/83217025a171593547d1268651b446d3533e2019/activesupport/lib/active_support/testing/time_helpers.rb#L160

```ruby
now = date_or_time.to_time.change(usec: 0)
```

この挙動は [ドキュメント](https://api.rubyonrails.org/classes/ActiveSupport/Testing/TimeHelpers.html#method-i-travel_to) にも明示されている仕様。

> Note that the usec for the time passed will be set to 0 to prevent rounding errors with external services, like MySQL (which will round instead of floor, leading to off-by-one-second errors).

今回自分が嵌ったのは、MySQL にミリ秒以下の精度の datetime を格納するよう意図して定義したカラムのテストで travel_to がうまく使えなかったというケースだった。

- 前述のドキュメントにもあるように MySQL の DATETIME 型のデフォルトは秒単位の精度だが、それ以上の精度を `DATETIME(6)` などとして指定できる (5.6.4 以降らしい)
    - [MySQL :: MySQL 5\.7 Reference Manual :: 11\.2\.1 Date and Time Data Type Syntax](https://dev.mysql.com/doc/refman/5.7/en/date-and-time-type-syntax.html)
    - [MySQL :: MySQL 5\.6 リファレンスマニュアル :: 11\.1\.2 日付と時間型の概要](https://dev.mysql.com/doc/refman/5.6/ja/date-and-time-type-overview.html)
- Rails の場合 migration ファイルに `precision` というオプションをつければ良い
    - [create\_table \| Railsドキュメント](https://railsdoc.com/page/create_table)

```ruby
t.datetime :your_column, precision: 6
```

- 今回は `Time.now` にだけ依存する処理のテストだったので、次のようにここだけをピンポイントに差し替えて対処した
    - travel_to がやっているように他のケースも stub する必要がある場合はもっと複雑になる

```ruby
allow(Time).to receive(:now).and_return(t)
```

```ruby
[1] pry(main)> require 'rspec/mocks/standalone'
=> true
[2] pry(main)> t = Time.now
=> 2021-12-01 17:59:01.306662808 +0000
[3] pry(main)> t.to_f
=> 1638381541.3066628
[4] pry(main)> allow(Time).to receive(:now).and_return(t)
=> #<RSpec::Mocks::MessageExpectation #<Time (class)>.now(any arguments)>
[5] pry(main)> t2 = Time.now
=> 2021-12-01 17:59:01.306662808 +0000
[6] pry(main)> t2.to_f
=> 1638381541.3066628      # <== [3] と同様の結果になっている
```

## PR

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B08D3DW7LP/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/51hCV4olz-L.jpg" alt="パーフェクト Ruby on Rails　【増補改訂版】" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B08D3DW7LP/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">パーフェクト Ruby on Rails　【増補改訂版】</a></div><div class="amazlet-detail">すがわら まさのり  (著), 前島 真一  (著), 橋立 友宏 (著), 五十嵐 邦明  (著), 後藤 優一 (著)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B08D3DW7LP/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
