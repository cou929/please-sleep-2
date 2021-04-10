{"title":"Go の append の挙動の確認","date":"2021-04-10T17:30:00+09:00","tags":["go"]}

スライスに要素を append した際に、もとのスライスが破壊的に変更されるのかどうかの理解が曖昧ではまったので確認したメモ。わりと何周もしているネタだと思われる。

## 今回はまったケース

今回はまったケースは次のようなもの。

- map の値としてスライスを保持
- map のあるキー A に入っているスライスに append した結果を別のキー B に代入、同じように B のスライスに append して C に...と繰り返す処理をしていた
- この例でキー A の値が、後に書きかわったり、書きかわらなかったりする挙動だった

append が第一引数の配列を破壊的に変更しているのだろうけど、変化がある場合と無い場合があり不思議だった。

## append の挙動

これについてはドキュメントで明確に説明されている。

[builtin \- The Go Programming Language](https://golang.org/pkg/builtin/#append)

- そのスライスのキャパシティに余裕があれば、もとの underlying array に追加する
- キャパシティに余裕がなければ新たにメモリを確保する

以前スライスの内部構造をまとめた際にこの挙動は確認済みだったけれど、解決までにちょっと時間がかかってしまった。知識としてはあっても手に馴染んでいない状態だった。

[Go の Slice の内部構造 \- Please Sleep](https://please-sleep.cou929.nu/golang-array-and-slice-note.html)

## 試してみる

前提として、スライスは内部的に配列を保持していて、以降はそれを underlying array と記載する。(前述の記事も参照)

[go \- How to inspect slice header? \- Stack Overflow](https://stackoverflow.com/questions/54195834/how-to-inspect-slice-header) によると、slice のポインタを [unsafe.Pointer](https://golang.org/pkg/unsafe/#Pointer) に変換し、さらにそれを [reflect.SliceHeader](https://golang.org/pkg/reflect/#SliceHeader) に変換させることができるらしい。これを使って確認してみる。

```go
package main

import (
    "fmt"
    "reflect"
    "unsafe"
)

func main() {
    slice := []int{1, 2, 3}
    fmt.Printf("%p, %p, %+v\n", &slice, slice, (*reflect.SliceHeader)(unsafe.Pointer(&slice)))
    // => 0xc00000c018, 0xc000014018, &{Data:824633802776 Len:3 Cap:3}
    // 初期状態

    slice = append(slice, 4)
    fmt.Printf("%p, %p, %+v\n", &slice, slice, (*reflect.SliceHeader)(unsafe.Pointer(&slice)))
    // => 0xc00000c018, 0xc000078000, &{Data:824634212352 Len:4 Cap:6}
    // cap に余裕がなかったのでメモリ再確保、Data のアドレスが変わっている、Cap は 6 まで伸びている

    slice = append(slice, 5)
    fmt.Printf("%p, %p, %+v\n", &slice, slice, (*reflect.SliceHeader)(unsafe.Pointer(&slice)))
    // => 0xc00000c018, 0xc000078000, &{Data:824634212352 Len:5 Cap:6}
    // cap に余裕があったので 824634482736 にそのまま append している。Len だけが増え、Cap はそのまま
}
```

[The Go Playground](https://play.golang.org/p/FqLXdzeS6v6)

キャパシティを広げるロジックは [ここ](https://github.com/golang/go/blob/11f159456b1dba3ec499da916852dd188d1e04a7/src/runtime/slice.go#L115-L240) のようだが、ある程度までは現状の長さの 2 倍ずつ拡張していく。

今回はまったケースについては、繰り返し回数が少ないテストケースで動作確認し、その後繰り返しが多い実際のケースに適用すると問題が発生した。今回の場合は初回の append は必ず underlying array のリアロケートが走るので、テスト時点では問題に気づけなかったというオチだった。

なお append そのものの実装については、コンパイルフェーズで直接アセンブリを生成しているらしく、ハードルが高いので今回は詳細の確認を割愛した。

[Where is the implementation of func append in Go? \- Stack Overflow](https://stackoverflow.com/questions/33405327/where-is-the-implementation-of-func-append-in-go)

# 参考

- [builtin \- The Go Programming Language](https://golang.org/pkg/builtin/#append)
- [go \- How to inspect slice header? \- Stack Overflow](https://stackoverflow.com/questions/54195834/how-to-inspect-slice-header)
- [unsafe \- The Go Programming Language](https://golang.org/pkg/unsafe/#Pointer)
- [reflect \- The Go Programming Language](https://golang.org/pkg/reflect/#SliceHeader)
- [go/slice\.go at master · golang/go](https://github.com/golang/go/blob/11f159456b1dba3ec499da916852dd188d1e04a7/src/runtime/slice.go#L115-L240)
- [Where is the implementation of func append in Go? \- Stack Overflow](https://stackoverflow.com/questions/33405327/where-is-the-implementation-of-func-append-in-go)

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://images-na.ssl-images-amazon.com/images/I/41meaSLNFfL._SX382_BO1,204,203,200_.jpg" alt="プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></div><div class="amazlet-detail">Alan A.A. Donovan (著), Brian W. Kernighan (著), 柴田 芳樹 (翻訳)<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
