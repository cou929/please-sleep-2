{"title":"BigQuery の job_id から実行したクエリを調査する","date":"2016-10-01T22:22:40+09:00","tags":["bigquery"]}

Web UI から実行したクエリは、左側の `Query History` メニューから検索できるけれど、API 経由で実行したクエリはこの方法では無理らしい。バッチなどで定期実行するクエリはこのケースに該当する認識。

こういう場合は `jobs: get` API で情報を取得する必要がある。(CLI ツールの [bq コマンド](https://cloud.google.com/bigquery/bq-command-line-tool) にこのインタフェースがあるといいんだけど、見当たらなかった)

[Jobs: get  |  BigQuery  |  Google Cloud Platform](https://cloud.google.com/bigquery/docs/reference/v2/jobs/get)

次のように `job_id` をパラメータに GET リクエストを投げると、そのジョブの詳細が返ってくる。レスポンスの一部に投げたクエリが入っている。

<pre><code data-language="sh">curl -H 'Content-Type: application/json' -H 'Authorization: Bearer xxx' -v 'https://www.googleapis.com/bigquery/v2/projects/turnkey-conduit-708/jobs/job_xxx' | jq .
</code></pre>

- `Authorization` には認証済みの access token を入れる
- path の `job_xxx` の部分を調査したい job_id にする
    - 過去に実行した job の一覧は [list api](https://cloud.google.com/bigquery/docs/reference/v2/jobs/list) で取得可能

レスポンスとして [かなりいろいろな情報](https://cloud.google.com/bigquery/docs/reference/v2/jobs#resource) が返ってくる。クエリは `$.configuration.query.query` を見れば良い。

また、tips としては

- サポートの問い合わせには必ず job_id が必要になるので、バッチの実行ログには job_id を出しておいたほうが良い
- access token は perl では次のようにして取得している

<pre><code data-language="perl">use Furl;
use JSON::WebToken;
use JSON::XS qw/decode_json/;

sub _build_access_token {
    my $time = time;
    my $client_email = 'foobarbaz@developer.gserviceaccount.com';
    my $expiration_span = 3600;
    my $private_key = '-----BEGIN PRIVATE KEY-----...';
    my $ua = Furl->new(
        agent   => 'MyAPIClient',
        timeout => 60 * 5,  # 5 min
    );

    my $jwt = JSON::WebToken->encode({
        iss => $client_email,
        scope => 'https://www.googleapis.com/auth/bigquery',
        aud => 'https://accounts.google.com/o/oauth2/token',
        exp => $time + $expiration_span,
        iat => $time,
    }, $private_key, 'RS256', {typ => 'JWT'});

    my $res = $ua->post(
        'https://accounts.google.com/o/oauth2/token',
        [],
        [grant_type => 'urn:ietf:params:oauth:grant-type:jwt-bearer',
         assertion => $jwt],
    );

    Carp::croak sprintf "Auth request failed\tcode:%s\tcontent:%s", $res->code, $res->content unless $res->is_success;

    my $data = decode_json $res->content;

    return $data->{access_token};
}
</code></pre>

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774180955/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="http://ecx.images-amazon.com/images/I/61tT2KMx%2BqL._SL160_.jpg" alt="クラウド開発徹底攻略 (WEB+DB PRESS plus)" style="border: none;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774180955/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">クラウド開発徹底攻略 (WEB+DB PRESS plus)</a><div class="amazlet-powered-date" style="font-size:80%;margin-top:5px;line-height:120%">posted with <a href="http://www.amazlet.com/" title="amazlet" target="_blank">amazlet</a> at 16.10.01</div></div><div class="amazlet-detail">菅原 元気 磯辺 和彦 山口 与力 澤登 亨彦 内田 誠悟 小林 明大 石村 真吾 相澤 歩 柴田 博志 伊藤 直也 登尾 徳誠 <br />技術評論社 <br />売り上げランキング: 132,345<br /></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774180955/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

