{"title":"foursquare のソフトウエアスタック","date":"2012-03-17T18:15:01+09:00","tags":["infra"]}

[About - foursquare](https://foursquare.com/about/)

4sq のアバウトページに 4sq のシステム構成が割りとガチに書いてあるらしいので読んでみた.

### サーバ構成

- EC2
- CentOS
- nginx でリバースプロキシ. リクエストを static なものとそうでないものに振り分け
- HAProxy でロードバランス

### データストア

- live site data (?) は MongoDB に保存
- memcached でキャッシュ
- オフラインでのデータ解析は, 定期的にデータを Hadoop クラスタにインポート
  - 基本的に Hive
- Solr と Elasticsearch で検索. 位置, Tips, ユーザ, イベントの検索に.
- geo-indexing 検索は Google s2 library, サーチインデックス内に cellid を保管 (?)
- PostGIS と geonames.org のデータセットで geocode address を座標に変換
- Kestrel を非同期タスクのためのキューとして使用
- 写真は Amazon S3 に保存して Akamai で配信

### アプリケーションフレームワーク

- Web サイト, API, バッチの多くは Scala
  - サイトと API は Lift web framework
- ビルドの自動化やデプロイ, オペレーションの自動化には Python や シェルスクリプト
- フロントエンドは Backbone.js + jQuery + Soy (templating)

### 地図データ

- 地図データは MapBox と OpenStreetMap のデータ

### Refs

- [Amazon Elastic Compute Cloud (Amazon EC2)](http://aws.amazon.com/ec2/)
- [www.centos.org - The Community ENTerprise Operating System](http://www.centos.org/)
- [nginx](http://nginx.org/en/)
- [HAProxy - The Reliable, High Performance TCP/HTTP Load Balancer](http://haproxy.1wt.eu/)
- [MongoDB](http://www.mongodb.org/)
- [memcached - a distributed memory object caching system](http://memcached.org/)
- [Welcome to Apache™ Hadoop™!](http://hadoop.apache.org/)
- [Welcome to Hive!](http://hive.apache.org/)
- [Apache Lucene - Apache Solr](http://lucene.apache.org/solr/)
- [elasticsearch - - Open Source, Distributed, RESTful, Search Engine](http://www.elasticsearch.org/)
- [s2-geometry-library - A Library for Spherical Math - Google Project Hosting](http://code.google.com/p/s2-geometry-library/)
- [PostGIS : Home](http://postgis.refractions.net/)
- [GeoNames](http://www.geonames.org/)
- [robey/kestrel](https://github.com/robey/kestrel)
- [Amazon Simple Storage Service (Amazon S3)](http://aws.amazon.com/s3/)
- [アカマイ：ウェブアプリケーションの高速化、パフォーマンス管理、ストリーミング・メディア・サービスならびにコンテンツデリバリーにおけるグローバルリーダー](http://www.akamai.co.jp/enja/)
- [Lift :: Home](http://liftweb.net/)
- [The Scala Programming Language](http://www.scala-lang.org/)
- [jQuery: The Write Less, Do More, JavaScript Library](http://jquery.com/)
- [Backbone.js](http://documentcloud.github.com/backbone/)
- [Closure Tools — Google Developers](https://developers.google.com/closure/templates/)
- [MapBox](http://mapbox.com/)
- [OpenStreetMap](http://www.openstreetmap.org/)
