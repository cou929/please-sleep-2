{"title":"Rails6 + GCP CloudSQL MySQL で Query Insights を使う際の Yak Shaving","date":"2022-12-25T23:20:00+09:00","tags":["rails", "gcp", "mysql"]}

GCP の CloudSQL には [Query Insights](https://cloud.google.com/sql/docs/mysql/using-query-insights?hl=ja) というクエリパフォーマンス分析ツールがある。その中の一機能に、クエリが発行されたコントローラーやアクション別にタグを振り、タグ別にパフォーマンスを集計してくれるというものがある。クエリへのタグ付けには [sqlcommenter](https://google.github.io/sqlcommenter/) というツールが公式に準備されている。この sqlcommenter の Rails6 以前への対応がいまいちで、導入に一工夫が必要だった。

## Query Insights とは

- クエリパフォーマンスを分析できる便利なツール
- 例えば AWS の RDS でいうところの [Performance Insights](https://aws.amazon.com/rds/performance-insights/) のようなもの

<figure>

<img src=images/query-insights-dashboard.png />

<figcaption>Query Insights を使用してクエリのパフォーマンスを向上させる https://cloud.google.com/sql/docs/mysql/using-query-insights より引用</figcaption>
</figure>

- 当初は Postgresql にしか対応していなかったが 2022-09-29 に [MySQL 対応](https://cloud.google.com/blog/topics/developers-practitioners/cloud-sql-query-insights-ga-mysql-query-load-tags-query-plans) も [GA になった](https://cloud.google.com/sql/docs/mysql/release-notes#September_29_2022)
- アプリケーション側に手を入れなくても利用できるが、クエリにタグを埋め込むとそのタグごとに分析ができるという便利な機能もある
    - 次のようにそのクエリが発行されたコントローラーやアクションを SQL コメントに書き出すと、タグごとに集計された統計を見ることができるようになる

```sql
SELECT * from USERS /*action='run+this',
controller='foo%3',
traceparent='00-01',
tracestate='rojo%2'*/
```

<figure>

<img src=images/query-insights-application-tags.png />

<figcaption>Query Insights を使用してクエリのパフォーマンスを向上させる https://cloud.google.com/sql/docs/mysql/using-query-insights より引用</figcaption>
</figure>

## sqlcommenter

- アプリケーションからのクエリコメントへのタグ情報埋め込みには [sqlcommenter](https://google.github.io/sqlcommenter/) というツールを使うよう [公式ドキュメントから案内されている](https://cloud.google.com/sql/docs/mysql/using-query-insights?hl=ja#adding-tags-to-sql-queries)
- sqlcommenter は有名所の Web フレームワーク・ORM には [対応していて](https://google.github.io/sqlcommenter/#frameworks)、Rails・ActiveRecord もサポートされている
- [Rails の場合](https://google.github.io/sqlcommenter/ruby/rails/) は、[sqlcommenter_rails](https://github.com/google/sqlcommenter/tree/master/ruby/sqlcommenter-ruby/sqlcommenter_rails) というパッケージを導入すればよい
    - 内部では [basecamp/marginalia](https://github.com/basecamp/marginalia/) を使って機能を実現している
- この sqlcommenter_rails の対応状況が微妙で、すんなりと導入できなかった
    - なお後述するが Rails7 以降ではこのような苦労をする必要はない

## 問題点

- 一言でいうと sqlcommenter_rails は rubygems に公開されていない
    - sqlcommenter_rails は marginalia 本家の [未マージの PR](https://github.com/basecamp/marginalia/pull/130) に依存しているためだと思われる
    - marginalia の [出力フォーマット](https://github.com/basecamp/marginalia/blob/adc34ade565164273fcc256cb40b871cf790a252/lib/marginalia/comment.rb#L27) は [sqlcommenter の spec](https://google.github.io/sqlcommenter/spec/) に合致せず、またフォーマットのカスタマイズができない作りになっている
    - 上記の PR でフォーマットをカスタマイズ可能にしようとしている
- また他の依存パッケージである [marginalia-opencensus](https://github.com/google/sqlcommenter/tree/master/ruby/sqlcommenter-ruby/marginalia-opencensus) も rubygmens に公開されていない
    - こちらはやろうと思えば公開できそうだが、単体で公開しても意味は無いので未対応なのだと思う
- そのせいか、[sqlcommenter_rails](https://github.com/google/sqlcommenter/tree/master/ruby/sqlcommenter-ruby/sqlcommenter_rails#license) も [marginalia-opencensus](https://github.com/google/sqlcommenter/tree/master/ruby/sqlcommenter-ruby/marginalia-opencensus#license) も License が明示されておらず、商用利用するにあたって懸念が残る

## とりあえずの解決

- sqlcommenter_rails は使わず、marginalia を直接導入して、必要な部分だけ自分でパッチを当てるのが手っ取り早かった
- `config/initializers/marginalia.rb` で出したいタグの種別 (components) を指定する
    - [spec](https://google.github.io/sqlcommenter/spec/#sorting) では key が辞書順になるよう規定されているので sort している
    - このサンプルは controller, action だけを指定しているが、Query Insights は他にも [framework, route, application, db_driver](https://cloud.google.com/sql/docs/mysql/using-query-insights?hl=ja#adding-tags-to-sql-queries) というタグにも対応している

```ruby
# config/initializers/marginalia.rb
Marginalia::Comment.components = [:controller, :action].sort
```

- `Marginalia::Comment` に monkey patch をあてる
    - marginalia のデフォルトは `key:value` というフォーマットだが、これを `'key'='value'` と出力するようにする
    - 参考
        - https://github.com/basecamp/marginalia/pull/130/files#diff-47bf3d98885a31ab30af60ac7e2c5373847c79e8aadcf0fe6610cca12e825ccbR30
        - https://github.com/google/sqlcommenter/blob/36ce7d3bcd80b0700907fa7eb93a108df7a598c6/ruby/sqlcommenter-ruby/sqlcommenter_rails/lib/sqlcommenter_rails/marginalia_components.rb#L37-L38
        - https://google.github.io/sqlcommenter/spec/
    - ちなみに、sidekiq などのジョブワーカーを使っている場合、marginalia は `job` というキーでワーカーのクラス名を出力するが、[Query Insights はこれに対応していない](https://cloud.google.com/sql/docs/mysql/using-query-insights?hl=ja#adding-tags-to-sql-queries)。そのためジョブワーカーの場合は `controller` キーにワーカーのクラス名を入れるようにしてみた
        - こういったワークアラウンドで正しいのかわからないが

```ruby
# config/initializers/marginalia_patch.rb など
module Marginalia
  module Comment
    def self.construct_comment
      ret = String.new
      is_job = !self.send(:job).nil?
      self.components.each do |c|
        component_value = self.send(c)
        component_value = self.send(:job) if is_job && c == :controller  # Query Insights に集計してもらうため job の場合 controller キーでもジョブ名を出力する
        if component_value.present?
          quoted = "'#{component_value.gsub("'", "\\\\'")}'"
          ret << "#{c}=#{quoted},"
        end
      end
      ret.chop!
      ret = self.escape_sql_comment(ret)
      ret
    end
  end
end
```

- この他にも framework, db_driver, route といったキーも出力したい場合は [sqlcommenter_rails のこのあたり](https://github.com/google/sqlcommenter/blob/36ce7d3bcd80b0700907fa7eb93a108df7a598c6/ruby/sqlcommenter-ruby/sqlcommenter_rails/lib/sqlcommenter_rails/marginalia_components.rb#L22-L33) を参考にするとよい

## Rails7 対応

- このような Yak Shaving が必要なのは Rils6 系までで、出たばかりの 7 系以降はしっかり対応がされている
    - ref. https://github.com/google/sqlcommenter/tree/master/ruby/sqlcommenter-ruby#sqlcommenter-support-in-rails
- 7.0 では [query_log_tags](https://guides.rubyonrails.org/configuring.html#config-active-record-query-log-tags-enabled) を有効化し、[PlanetScale の sqlcommenter gem](https://github.com/planetscale/activerecord-sql_commenter#installation) を導入すれば良い
    - Rails7 では [marginalia が Rails 本体に導入されていて](https://github.com/rails/rails/pull/42240)、ActiveRecord の [QueryLogs というモジュール](https://github.com/rails/rails/blob/83009f94e7539c7db89f83446b1478bdaec9d1a0/activerecord/lib/active_record/query_logs.rb) が該当処理を担当している
    - PlanetScale の gem は前述の [marginalia への PR](https://github.com/planetscale/activerecord-sql_commenter) に近い内容を `ActiveRecord::QueryLogs` に適用している
        - よって Rails6 以前で使うことはできない
- 7.1 以降では `query_log_tags` を有効化するだけでよい
    - sqlcommenter フォーマットへの対応がすでに [公式に取り込ま出ている](https://github.com/rails/rails/pull/45081)
- このような事情もあって現状の sqlcommenter_rails をはじめ Rails6 以前への対応が積極的にされていないのだと思われる
    - なぜこんな中途半端な状態で放置されているのか疑問だったが、これで合点がいった
    - この情報は GCP や sqlcommenter のドキュメントサイトには無く、sqlcommenter の GitHub 上の README に気づかないとわからない状態に現状なっている

## OpenTelemetry

- ちなみに sqlcommenter プロジェクトはもともと Google がホストしていたものだったが、現在では CNCF の [OpenTelemetry](https://opentelemetry.io/) に譲渡されているらしい
- よって sqlcommenter_rails 自体を改善したい場合はこちらのルートから進んだ方がおそらく良い
- 前述のように Rails7 では対応済みなので、Rails ではない Ruby のプロジェクトでくらいしか用途がなさそうだが
- 現状の OpenTelemetry 側の対応状況はちょっとよくわからないが、おそらく解決はしていなそうだった
    - https://github.com/google/sqlcommenter/#sqlcommenter
    - https://github.com/open-telemetry/opentelemetry-sqlcommenter

## PR

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B09ZQ6FHTT/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/I/41Sbx79V4jL.jpg" alt="Observability Engineering (English Edition)" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B09ZQ6FHTT/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Observability Engineering (English Edition)</a></div><div class="amazlet-detail">英語版  Charity Majors  (著), Liz Fong-Jones  (著), George Miranda  (著)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B09ZQ6FHTT/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
