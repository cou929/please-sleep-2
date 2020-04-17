{"title":"travis-ci でのビルドが No Rakefile found で失敗する","date":"2015-05-31T22:46:03+09:00","tags":["nix"]}

静的ファイルだけを置くリポジトリを GitHub に作って、マージするたびにそれを S3 に置く作業を travis にやらせようと思ったところ、`No Rakefile found` というエラーでビルドが失敗して少しはまった。

ちょっとググると [travis.yml のファイル名を typo しているせいだ](http://stackoverflow.com/questions/15953739/error-on-travis-ci-build-no-rakefile-found) とか、[yaml の構文エラーのせいだ](http://stackoverflow.com/questions/13489301/failed-to-build-and-deploy-node-js-project-with-travis-ci-no-rakefile-found) とかいう情報がでてくる。名前も間違っていないし [lint](http://yaml.travis-ci.org/) も通っているのでこういう話ではない。

原因は、当たり前だけど「Travis は (Ruby の場合は `Rake` というように) デフォルトではビルドのシステムを必要とする」ということだった。S3 に置くだけでいいので、[deploy](http://docs.travis-ci.com/user/deployment/s3/) ディレクティブだけを書いた `.travis.yml` を準備していたが、それだと travis が rake でビルドしようとするので失敗していた。

とりあえず `script` ディレクティブを空にしてなにもしないようにさせて回避した。

    script: ""
