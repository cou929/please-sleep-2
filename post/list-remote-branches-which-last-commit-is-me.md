{"title":"最終コミットが自分のリモートブランチを一覧で出す","date":"2013-09-08T16:25:55+09:00","tags":["git"]}

    PATTERN='cou929@'; for br in $(git branch -a | grep remote | grep -v HEAD | awk '{print $1}'); do git log -n1 $br | grep $PATTERN > /dev/null && echo $br; done

`PATTERN` には自分のメールアドレスなど自分を特定できるパターンをいれておく。ブランチ一覧のすべてに対して `git log -n 1` で直近のコミットを取得。PATTERN にマッチするものがあればブランチ名を echo するだけ。
