{"title":"curl でレスポンスタイムを計測","date":"2013-06-21T20:02:02+09:00","tags":["nix"]}

`w/--write-out` というオプションがあり、いろいろと細かい情報をフォーマットして出力できる。たとえばこうするとレスポンスタイムだけを出力できる。

    curl -kL 'https://example.com/' -o /dev/null -w "%{time_total}" 2> /dev/null

レスポンスのステータスコードも出したい場合は

    curl -kL 'https://example.com/' -o /dev/null -w "%{http_code}\t%{time_total}" 2> /dev/null

たとえばこれをログにだしておいたり、growthforecast などに投げ込んでおくと、エンドポイントの簡易的な計測に使えそうだ。

ほかにも色々なオプションがあるので man を参照。

### 参考

[curlのオプション勉強したのでまとめ - うまい棒blog](http://d.hatena.ne.jp/hogem/20091122/1258863440)
