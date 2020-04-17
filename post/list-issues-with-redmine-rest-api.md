{"title":"redmine で一週間分の完了した担当チケットをリストアップするワンライナー","date":"2013-07-05T15:31:39+09:00","tags":["nix"]}

自分が担当しているチケットのうち、ここ 1 週間で完了したものを "チケット名 #チケット番号" というスタイルでリストアップするワンライナー。

    APIKEY=XXX; PROJECT_ID=1; USER_ID=1; AWEEKAGO=`date -v-7d +%Y-%m-%d`; curl -s "https://your.redmine.host/issues.json?key=${APIKEY}&project_id=${PROJECT_ID}&assigned_to_id=${USER_ID}&status_id=5&updated_on=%3E%3D${AWEEKAGO}" | perl -MJSON::PP -nle 'binmode(STDOUT, ":utf8"); for my $issue (@{decode_json($_)->{issues}}) { printf "* %s #%d\n", $issue->{subject}, $issue->{id} }'

APIKEY は redmine の `/my/page` で確認できる。ユーザー ID は redmine の UI 上で自分の名前のリンクをクリックすると `/users/<ID>` という url に飛ぶので、それで確認するのがはやい。redmine の リストは textile 形式なのでアスタリスクを行頭に置いている。

### 参考

[Rest Issues - Redmine](http://www.redmine.org/projects/redmine/wiki/Rest_Issues)
