{"title":"Go の interface と channel の良いサンプル","date":"2020-06-18T22:32:00+09:00","tags":["go"]}

[revel/cmd](https://github.com/revel/cmd) の [`harness.AppCmd.Start` という関数](https://github.com/revel/cmd/blob/531aa1e209463d09e3c1d6602d7ad4f2218e742c/harness/app.go#L63-L91) が Go らしさのある、Go の機能を説明するのに良いサンプルだなと思ったのでメモ。interface (io.Writer) と並列処理 (channel, select) を理解するのに良さそうだった。

## revel/cmd とは

[revel/cmd](https://github.com/revel/cmd) は ウェブアプリケーションフレームワーク [Revel](http://revel.github.io/) の CLI コマンド。世の中のフレームワークによくある、アプリケーションを new したり run したりする機能を提供している。Revel は Go では珍しいフルスタック系のフレームワークなので、こういうツールが提供されている。

`$ revel run ...` (サーバアプリケーションの起動) コマンドを実行すると、ビルドとサーバプロセスの起動が行われる。これは次のような、少々大がかりな仕組みになっている。

- サーバアプリケーションの main package 部分を [テンプレート](https://github.com/revel/cmd/blob/531aa1e209463d09e3c1d6602d7ad4f2218e742c/harness/build.go#L447-L527) から生成する
- ビルドしバイナリを作る
- バイナリを子プロセスで起動する
- 起動モードが `prod` 以外の場合さらに以下も行う
    - ソースコードの変更を監視し、変更があれば再ビルド・サーバ再起動を行う
    - リバースプロキシを立て、子プロセスへリクエストを中継する

## harness.AppCmd.Start

`harness.AppCmd.Start` はこの中でも、ビルドされたバイナリを別プロセスでの起動する処理。[os/exec.Cmd](https://pkg.go.dev/os/exec?tab=doc#Cmd) で子プロセスを起動、バイナリを実行している。子プロセスの起動状況を監視し、正常であれば後続処理を続行、異常であればエラーを報告する。このとき、正常起動したことをは、次のように検知している。

- サーバアプリケーションが正常に起動すると `Revel engine is listening on..` というメッセージが標準出力に書き込まれる
- 親プロセス (revel/cmd) は子プロセス (サーバアプリケーション) の標準出力を監視し、上記のメッセージが書き込まると正常に起動したとみなしている

この要件を interface (io.Writer) と並列処理 (channel, select) ですっきりと実現していた。

- io.Writer を実装した独自の Writer を子プロセスの Stdout に設定する
- 独自 Writer はメッセージを検知すると channel で親プロセスにメッセージを送る
- 親プロセスは select でメッセージを待ち受ける
    - タイムアウトや異常終了も同時に待ち受けている

Go らしくて良いなと思ったのがこれらのポイント。それぞれもう少し詳細に見ていく。

## io.Writer interface

[os/exec.Cmd](https://pkg.go.dev/os/exec?tab=doc#Cmd) 構造体の `Stdout` フィールドは [io.Writer](https://pkg.go.dev/io?tab=doc#Writer) 型。

revel/cmd の [startupListeningWriter](https://github.com/revel/cmd/blob/531aa1e209463d09e3c1d6602d7ad4f2218e742c/harness/app.go#L182-L187) 構造体は Write メソッドを実装し、io.Writer 型を満たしている。

`startupListeningWriter.Write` メソッドは次のように、特定のメッセージがあれば `notifyReady` という channel に true を送信し、いずれの場合でも来たメッセージを `os.Stdout` に書き込んでいる。つまり、基本的には来た内容をそのまま標準出力にバイパスし、特定のメッセージがあったときだけ channel で親プロセスにメッセージを送っている。

```go
// A io.Writer that copies to the destination, and listens for "Revel engine is listening on.."
// in the stream.  (Which tells us when the revel server has finished starting up)
// This is super ghetto, but by far the simplest thing that should work.
type startupListeningWriter struct {
	dest        io.Writer
	notifyReady chan bool
	c           *model.CommandConfig
	buffer      *bytes.Buffer
}

// Writes to this output stream
func (w *startupListeningWriter) Write(p []byte) (int, error) {
	if w.notifyReady != nil && bytes.Contains(p, []byte("Revel engine is listening on")) {
		w.notifyReady <- true
		w.notifyReady = nil
	}
	if w.c.HistoricMode {
		if w.notifyReady != nil && bytes.Contains(p, []byte("Listening on")) {
			w.notifyReady <- true
			w.notifyReady = nil
		}
	}
	if w.notifyReady!=nil {
		w.buffer.Write(p)
	}
	return w.dest.Write(p)
}

func (cmd AppCmd) Start(c *model.CommandConfig) error {
	listeningWriter := &startupListeningWriter{os.Stdout, make(chan bool), c, &bytes.Buffer{}}
	cmd.Stdout = listeningWriter

	...
```

`cmd.Stdout` に listeningWriter を代入するだけで良いのが、らしいポイントだなと思った。

## 並列処理

[os/exec.Cmd.Start](https://pkg.go.dev/os/exec?tab=doc#Cmd.Start) で子プロセスを起動した後に、select で次のイベントをブロックして待っている。

- 子プロセスの終了
    - 子プロセスはサーバプロセス (リクエストを Listen している) ので、すぐに終了コードが返ってくるのは異常ということになる
    - goroutine で [os/exec.Cmd.Wait](https://pkg.go.dev/os/exec?tab=doc#Cmd.Wait) で [ブロックして待ち、終了したら exitStatus channel に内容を送っている](https://github.com/revel/cmd/blob/531aa1e209463d09e3c1d6602d7ad4f2218e742c/harness/app.go#L164-L177)
- タイムアウト
    - [time.After](https://pkg.go.dev/time?tab=doc#After) で一定時間後に通知を受け取る
    - Go では [割と有名な idiom だと思う](https://talks.golang.org/2012/concurrency.slide#35)
- 子プロセスの正常起動
    - 前述の `startupListeningWriter.Write` が notifyReady channel に書き込んだものを受け取る

```go
func (cmd AppCmd) Start(c *model.CommandConfig) error {
	...

	if err := cmd.Cmd.Start(); err != nil {
		utils.Logger.Fatal("Error running:", "error", err)
	}

	select {
	case exitState := <-cmd.waitChan():
		fmt.Println("Startup failure view previous messages, \n Proxy is listening :", c.Run.Port)
		err := utils.NewError("","Revel Run Error", "starting your application there was an exception. See terminal output, " + exitState,"")
		// TODO pretiffy command line output
		// err.MetaError = listeningWriter.getLastOutput()
		return err

	case <-time.After(60 * time.Second):
		println("Revel proxy is listening, point your browser to :", c.Run.Port)
		utils.Logger.Error("Killing revel server process did not respond after wait timeout.", "processid", cmd.Process.Pid)
		cmd.Kill()
		return errors.New("revel/harness: app timed out")

	case <-listeningWriter.notifyReady:
		println("Revel proxy is listening, point your browser to :", c.Run.Port)
		return nil
	}
```

いろいろな並列処理の結果をシーケンシャルに受け取るのを、channel と select を使ってシンプルに書けていて、Go らしさがわかりやすい気がした。ちなみに goroutine 間のコミュニケーションを channel で行うのが普通だが、プロセス間の通信でも同じようにやっているのが、個人的にちょっと面白かった。

## 全体の骨格

今回のポイントだけをざっくりとまとめると次のようなコードになる。

```go
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

type startupListeningWriter struct {
	dest        io.Writer
	notifyReady chan bool
}

func (w *startupListeningWriter) Write(p []byte) (int, error) {
	if bytes.Contains(p, []byte("example target line")) {
		w.notifyReady <- true
	}
	return w.dest.Write(p)
}

func main() {
	cmd := exec.Command(`echo`, `example target line`)
	listeningWriter := &startupListeningWriter{os.Stdout, make(chan bool)}
	cmd.Stdout = listeningWriter

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	select {
	case <-time.After(5 * time.Second):
		cmd.Process.Kill()
		panic("timeout")
	case <-listeningWriter.notifyReady:
		fmt.Println("child process invoked correctly")
	}
}
```

## さいごに

interface、なかでも io.Writer/Reader や groutine 間の channel でのやりとり (今回はプロセスだったが) は Go の良さがわかりやすい機能だと思う。今回のサンプルこれらの恩恵を理解しやすい気がしたので簡単にメモした。
