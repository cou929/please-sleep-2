{"title":"revel のアプリケーションをデバッガでデバッグする","date":"2019-01-20T17:06:25+09:00","tags":["go"]}

- 一度 `revel run` する
- revel がビルドしたファイルが `tmp/main.go` にできる
- これをデバッガから起動する

例えば [delve](https://github.com/go-delve/delve) を使う場合はこんな感じ

    dlv debug --log path/to/tmp/main.go -- -port=YOUR_PORT -importPath=github.com/your/proj -runMode=dev

最初 `revel run` したプロセスに `dlv attach` しようとしていてうまく行かなかったのですこしはまった。細かい tips だがメモしておく。

### 参考

- [Debugging \| Revel \- A Web Application Framework for Go\!](https://revel.github.io/manual/debug.html)
- [go \- Revel debugging \- How to \- Stack Overflow](https://stackoverflow.com/questions/23952886/revel-debugging-how-to)
