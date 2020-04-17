{"title":"aws-cli s3 sync の --exact-timestamps","date":"2015-02-04T00:12:01+09:00","tags":["nix"]}

[awc-cli](https://github.com/aws/aws-cli) の s3 sync で s3 からローカルへファイルを同期する場合、変更前後のファイルサイズが同じだった場合は、変更なしとみなされ、そのファイルは同期されない。`--exact-timestamps` オプションをつけると、ファイルサイズだけでなくファイルの更新日時をみてファイルの変更判定をしてくれるので、このようなケースでも同期が走る。

### sync 挙動の概要

[この issue](https://github.com/aws/aws-cli/issues/599) によると、現在の sync の際のファイル変更判定ロジックは、ファイルサイズと最終更新日時で判定しているらしい。s3 からローカルへのファイル同期の場合、ファイルサイズが異なっていれば当然変更があったものとみなされる。もしファイルサイズが同じだった場合、デフォルトでは更新日時を見ずに変更がなかったものと判定されるようだ。

ここでの「最終更新日時」(`ListObjects` の `LastModified`) はそのファイルが s3 にアップロードされた時間で、いわゆる mtime とは異なる。たとえば、ローカルから s3 へ sync し、そのすぐあとに s3 から同じローカルへ sync した場合。ローカルのファイルと s3 上のファイルは同じにもかかわらず、ローカルの mtime と s3 上の `LastModified` が異なるために、変更があったとみなされる可能性がある。そのため、無駄な通信をなくすために、デフォルトではタイムスタンプを見ないようにしているようだ。

たとえば、あるインスタンス A から s3 に sync し、それを別のインスタンス B で sync して取得する場合。ファイルサイズが変わらない変更はインスタンス A から B へは伝わらないということが起こりえる。そんな場合は `--exact-timestamps` をつけて、タイムスタンプ込みで判定するようにすれば良い。

[前述の issue](https://github.com/aws/aws-cli/issues/599) を見ていると ETag での変更判定なども議論されているけれど、現行のバージョンにはまだ入っていないようだ。

### 参考

- [sync — AWS CLI 1.7.4 documentation](http://docs.aws.amazon.com/cli/latest/reference/s3/sync.html)
- [AWS S3 Sync Issues · Issue #599 · aws/aws-cli](https://github.com/aws/aws-cli/issues/599)
- [s3 sync not syncing files locally with that have the same size but have been modified · Issue #1074 · aws/aws-cli](https://github.com/aws/aws-cli/issues/1074)
- [Merge branch 'release-1.3.23' · 7526a9d · aws/aws-cli](https://github.com/aws/aws-cli/commit/7526a9de170ccd2d0d558dcf049146903230d4d9)
