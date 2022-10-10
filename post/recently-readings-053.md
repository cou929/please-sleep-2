{"title":"最近読んだもの 53 - Efficient MySQL Performance 4 章まで","date":"2022-10-10T23:30:00+09:00","tags":["readings"]}

- [Efficient MySQL Performance](http://www.amazon.co.jp/exec/obidos/ASIN/B09N5NWKR1/pleasesleep-22/ref=nosim/) の 4 章まで
    - まだ半分くらいだがいい本
    - MySQL の初級・上級の本は既刊であるが、その間を埋めるものがないので書かれたというもので、難易度を 1 ~ 5 で表すと 4 くらい、難易度 5 は [実践ハイパフォーマンスMySQL](http://www.amazon.co.jp/exec/obidos/ASIN/4873116384/pleasesleep-22/ref=nosim/) とのことだった
    - あくまで Deep dive したいアプリケーションエンジニア向けの本で、DBA 向けではないと明記されていた
    - まず最初の章で North Star Metrics としてクエリのレスポンスタイムを定義し、その改善にひつような項目を体系立てて説明している。この構成がかなり良い
        - レスポンスタイム以外の指標は measurable でなかったり actionable でなかったりなどするので、ここに絞っている
        - こうした明確な指標があることで対応がぶれないし、この本の内容も簡潔かつ実践的になっている
        - またチューニングも以降の章で紹介する順序で行うことを推奨している
            - 例えばいきなりハードウェアのスケールアップをするのは良くないなど
    - 2 章ではまずクエリのチューニングということで、クエリとインデックスについて扱う
        - 純粋なクエリとインデックスのチューニングがこの章のテーマ
        - Innodb のクラスタードインデックスとセカンダリインデックスの構造が説明される
        - そこからインデックスの leftmost prefix ルール、EXPLAIN の見方と順序立てて解説され、とても読みやすい
    - 3 章ではデータのチューニングを扱う
        - まず大きなデータセットではいくらインデックスを使っても限界があり、最小限のデータサイズや QPS を達成することが大切という原則が確認される
        - クエリが不要なデータをフェッチすること防ぐこと、必要最低限のデータだけを保持すること、それぞれチェックリストが提示される
        - データ削除の際の注意点 (バッチサイズなど) も紹介がある
    - 4 章ではアクセスパターンの改善を扱う
        - 言い換えるとアプリケーションの変更を伴う変更
        - リードのオフロード (キャッシュまたはリードレプリカの利用)、ライトのキューイング (非同期 API やイベントストリームベースの設計など)、データのパーティション (主にホットとコールドデータで分けること、シャーディングは次章で扱う)
        - 最後にここまで検討して改善が難しければハードウェアのスケールアップも検討の余地があるなど
    - それ以降はまだ呼んでいるところだが、まず 5 章ではシャーディングが扱われている
        - その後はメトリクス、レプリケーションラグ、トランザクションといったトピックが並んでおり楽しみ

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B09N5NWKR1/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/W/IMAGERENDERING_521856-T2/images/I/4151TBrJq1L.jpg" alt="Efficient MySQL Performance (English Edition)" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B09N5NWKR1/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Efficient MySQL Performance (English Edition)</a></div><div class="amazlet-detail">英語版  Daniel Nichter  (著)  形式: Kindle版<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B09N5NWKR1/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873116384/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://m.media-amazon.com/images/W/IMAGERENDERING_521856-T2/images/I/51li6EcTU+L._SX389_BO1,204,203,200_.jpg" alt="実践ハイパフォーマンスMySQL 第3版" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873116384/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">実践ハイパフォーマンスMySQL 第3版</a></div><div class="amazlet-detail">Baron Schwartz (著), Peter Zaitsev (著), Vadim Tkachenko (著), 菊池 研自 (監修), 株式会社クイープ (翻訳)<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873116384/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
