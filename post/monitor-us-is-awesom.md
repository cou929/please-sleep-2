{"title":"monitor.us が便利だ","date":"2012-12-03T01:22:50+09:00","tags":["nix"]}

[FREE Website Monitoring & Monitoring Software from Monitor.Us](http://www.monitor.us/website-monitoring)

これすごい.

- 外から http, ping などを飛ばして監視 & 通知. (smokeping 的な)
  - mysql への接続も監視可能
- エージェントをインストールすれば cpu / mem / disk / load averate / プロセスごとの cpu や mem といった基本的な項目を監視可能
- 無料プランで上記はできる
- 有料プランにアップデートすれば監視間隔を細かくしたり (無料版では 30 分), さらにいろいろな監視項目が追加される.

![dashboard image](http://gyazo.com/cfdb9dbe07b97ecad644603182bcc420.png?1354464842)

個人でさくら vps でやるような小規模なサービスとか自分用システムはこれ一択と言ってもいい. 正直外形監視だけかと思ってたけどちゃんと中に一個クライアントを入れてサーバのリソースのモニタリングまでしれくれるところまでを無料プランでやらせてくれるのは驚いた.

自分で nagios やら munin やら設定するのは面倒だしもう一個サーバいるしそんな手間とお金をかけたくない
