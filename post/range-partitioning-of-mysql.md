{"title":"MySQL でログ系テーブルのレンジパーティション","date":"2013-11-24T12:57:41+09:00","tags":["nix"]}

MySQL にはパーティションという機能がある。データ全体を特定の範囲ごとに小分けにしてインデックスをそれぞれ別に作ることができる機能と自分は理解している。たとえば id が 10,000 以下のものとそれ以上のものでパーティションをわけると、それぞれひとつずつインデックスが作成される。操作対象がひとつのパーティション内で完結する場合は、検索・更新・削除ともにより高速になる。全体でひとつの大きなインデックスを使うのではなく、部分ごとに小さなインデックスファイルを作成するという戦略だ。

全体のうちある特定の部分にアクセスがかたよる傾向のあるテーブルに対してパーティションをはるのが有効だ。適切な単位でパーティションをきることで、アクセスがひとつのパーティション内の小さなインデックスに収まり、更新も検索も高速に保てる。

例としてわかりやすいのはログ系のテーブルだろう。基本的に時系列に増えていくもののため、日付があたらしい部分にのみ挿入される。検索に関しても、現在に近いレコードへの参照がほとんどになるだろう。古いレコードを参照する場合でも特定の期間を指定して SELECT を行うユースケースがほとんどのはずだ。よって時系列のログを貯めるテーブルには日付でパーティションを切る戦略が有効だ。

またパーティションはその DROP が高速という特徴もある。`ALTER TABLE table_name DROP PARTITION partition_name` というクエリを発行すると対象のパーティション全体を削除してくれる。この操作が DELETE 文よりも非常に高速に動作するため、大量のデータをオンラインで落とす場合に有効だ。前述のようなログテーブルの場合、一定期間経過したデータはファイルに落として保管しておき、テーブルからは DROP PARTITION で消しこんでおくという運用が比較的簡単に実現できる。

注意点はパーティションのキーとして指定したすべてのカラムは、そのテーブルのすべてのユニーク制約に含まれている必要がある。今回のような例の場合、日付のカラムだけではユニークにすることができない。この制約を回避するために、id カラムを用意して日付とあわせて複合の主キーにするというテクニックがある。たとえば以下の操作ログをとるテーブルの例は、本来 id は必要ないがパーティションを設定するために追加している。

    CREATE TABLE operation_logs (
      id          BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
      operator    VARCHAR(255) NOT NULL,
      operation   VARCHAR(255) NOT NULL,
      operated_at datetime NOT NULL
      PRIMARY KEY (id, operated_at),
      INDEX operated_at,
    ) ENGINE=InnoDB
    PARTITION BY RANGE ( TO_DAYS(operated_at) ) (
      PARTITION p201301 VALUES LESS THAN ( TO_DAYS('2013-01-01') ),
      PARTITION p201302 VALUES LESS THAN ( TO_DAYS('2013-01-02') ),
      PARTITION p201303 VALUES LESS THAN ( TO_DAYS('2013-01-03') )
    );

### 参考

- [MySQL :: MySQL 5.1 Reference Manual :: 18.5.1 Partitioning Keys, Primary Keys, and Unique Keys](http://dev.mysql.com/doc/refman/5.1/en/partitioning-limitations-partitioning-keys-unique-keys.html)
- [ソーシャルゲームのためのMySQL入門 - Technology of DeNA](http://engineer.dena.jp/2010/11/mysql-for-socialgame.html)

