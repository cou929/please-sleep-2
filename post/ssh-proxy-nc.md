{"title":"踏み台サーバ経由の ssh を nc で proxy する","date":"2012-12-14T23:56:27+09:00","tags":["nix"]}

よくやるが毎回ググっている気がする.

踏み台サーバのユーザーの鍵を手元に持ってきておいて, 手元の `.ssh/config` に以下を設定すれば OK

<script src="https://gist.github.com/4263989.js"></script>
