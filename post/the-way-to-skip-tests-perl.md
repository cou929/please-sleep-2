{"title":"perl で状況に応じてテストをスキップする仕組み","date":"2012-10-06T13:04:26+09:00","tags":["perl"]}

perl の話です.

### Test::Requires

    use Test::Requires {
        'HTTP::MobileAttribute' => 0.01, # skip all if HTTP::MobileAttribute doesn't installed
    };
    
    use Test::Requires 'Text::SimpleTable';

こんな風に書いておくと, そのモジュールがなかった場合そのファイルのテストすべてをスキップする

prove するとこんな出力

    t/06_jslint.t ....................... skipped: Test requires module 'Text::SimpleTable' but it's not found

### Test::More の skip_all

    use Test::More skip_all => $skip_reason;
    
    plan skip_all => 'skip reason' if <conditional>;

skip_all ですべてのテストをスキップできる. スキップ理由も.

呼んだ時点ですぐにスクリプトは終了する. 終了コードは 0 (成功).

### Test::More の skip

    SKIP: {
         skip $why, $how_many if $condition;
    
         ...normal testing code goes here...
    }

そのブロックの中の `$how_many` 個のテストをスキップするらしい.
