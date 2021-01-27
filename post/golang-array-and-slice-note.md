{"title":"Go の Slice の内部構造","date":"2021-01-27T21:57:00+09:00","tags":["go"]}

Go で [お楽しみで競プロ問題を解いたり](https://please-sleep.cou929.nu/leetcode-with-mobile.html)、バイナリを触るコードを書いたりしていたところ、slice と array をあらためて理解しておいたほうが良い気がしてきた。そのため軽くドキュメントなどから内部構造を確認したメモ。

([Go Slices: usage and internals \- The Go Blog](https://blog.golang.org/slices-intro) から関連ドキュメントをたどって読んだだけです)

## 要点

使うだけならこれだけ理解しておけば大丈夫そう。

### Array

- array は長さも含めて型
- array は値

```go
a := [4]int{1, 2, 3, 4}  // array
```

### Slice

- slice は内部的な array (underlying array) へのポインタと len, cap を持っている
    - 概念的にはこんなかんじ
    - Capacity は undelying array の長さ

```go
// []byte のイメージ。Go のコンパイラが実際にこういう構造体を管理しているわけではない
type sliceHeader struct {
    Length        int
    Capacity      int
    ZerothElement *byte
}
```

- slice 記法で部分配列・部分スライスを取り出した slice を作ることができる
    - その際の undelying array は元となる array (または元となる slice の undelying array) と同じ

```go
arr := [3]int{1,2,3}
s := arr[:]
s[1] = 999
fmt.Println(arr)  // [1 999 3]  (s の underlying array なのでこちらも変わっている)
fmt.Println(s)    // [1 999 3]
```

- copy は新たな underlying array 用のメモリを確保する
- append は追加要素が cap 以内なら同じ underlying array を使い回し、そうでなければ拡張する
    - underlying array の拡張はより大きなメモリを確保し内容をそちらへコピーする

```go
arr := [2]int{1,2}
s := arr[:]
fmt.Println(s, len(s), cap(s))  // [1 2] 2 2

s[0] = 99
fmt.Println(arr)  // [99 2] (s の underlying array は arr なのでそちらも更新されている)

s = append(s, 3)
s[0] = 999
fmt.Println(s, len(s), cap(s))  // [999 2 3] 3 4
fmt.Println(arr)                // append での拡張時に underlying array 用に別のメモリ領域が確保されているので、999 に上書きされていない
```

## 実装

- このあたりが該当のコードらしい
    - [go/slice\.go at cd176b361591420f84fcbcaaf0cf24351aed0995 · golang/go](https://github.com/golang/go/blob/cd176b361591420f84fcbcaaf0cf24351aed0995/src/runtime/slice.go)
    - [go/ssa\.go at cd176b361591420f84fcbcaaf0cf24351aed0995 · golang/go](https://github.com/golang/go/blob/cd176b361591420f84fcbcaaf0cf24351aed0995/src/cmd/compile/internal/gc/ssa.go#L2841)
- [copy](https://github.com/golang/go/blob/cd176b361591420f84fcbcaaf0cf24351aed0995/src/runtime/slice.go#L247)
     - 最終的に [memmove](https://github.com/golang/go/blob/cd176b361591420f84fcbcaaf0cf24351aed0995/src/runtime/slice.go#L277) している
     - 恐らく [memmove(3)](https://man7.org/linux/man-pages/man3/memmove.3.html) だと思う (linux の場合)
        - memmove 自体が src, dst がオーバーラップしていても ok な仕様らしい
        - 計算量はたぶん `O(n)`
- [append 時の underlying array の拡張](https://github.com/golang/go/blob/cd176b361591420f84fcbcaaf0cf24351aed0995/src/runtime/slice.go#L144-L163)
    - 大雑把に言うと足りなくなったら現状の 2 倍のメモリを確保している
    - ただし、大きな slice では 2 倍のメモリを確保し続けないように制限が入っている
    - 既存 cap の 2 倍を確保すれば、要求される cap を満たせる場合
        - 既存 cap が 1024 未満
            - そのまま 2 倍のメモリを確保
        - 既存 cap が 1024 以上
            - `既存 cap / 4` の倍数で要求 cap を超えた最小値を確保
                - つまり要求 cap を少し超えた、でも 二倍未満の領域を確保している
    - 既存 cap を 2 倍しても要求 cap を満たせない場合
        - 要求 cap 分のメモリを確保

## その他のトピックや感想

- array から slice を作る場合は `slice := array[:]` というイディオムを覚えておくと便利
- 内部構造をこのように設計にすることで、部分配列 (部分スライス、部分文字列も) 取得や、cap 以内の要素の出し入れが高速に処理できるメリットがある
    - slice の拡張やコピーはメモリ確保、要素のコピー (`O(n)`) が入るのでコストは高いが、ユーザーがそれを防ぐようコードを書く手立てがある
- この設計のデメリットとしては、巨大な underlying array から実際に利用する小さな部分スライスを使うようなケースで無駄が多いこと
    - 例えば以下では `FindDigits` の戻り値の slice の underlying array に読み込んだファイル全体が入っている
    - しかし、アクセスするのはその中の正規表現にマッチした部分だけで無駄が多いし、戻り値が使われている限り gc の対象にもならない
    - FindDigits は copy なり append なりで、正規表現マッチ部分だけのメモリを新たに確保してそれを返すようにしたほうがよい

```go
var digitRegexp = regexp.MustCompile("[0-9]+")

func FindDigits(filename string) []byte {
    b, _ := ioutil.ReadFile(filename)
    return digitRegexp.Find(b)
}
```

- [research\!rsc: Go Data Structures](https://research.swtch.com/godata) によると、これは [Java など他の言語でもよく知られた問題 (gotcha)](https://bugs.java.com/bugdatabase/view_bug.do?bug_id=4513622) らしい
    - 解決を試みたが、スライス作成時のコストが低いメリットが上回るため現状の設計が維持されているそう

> (As an aside, there is a well-known gotcha in Java and other languages that when you slice a string to save a small piece, the reference to the original keeps the entire original string in memory even though only a small amount is still needed. Go has this gotcha too. The alternative, which we tried and rejected, is to make string slicing so expensive—an allocation and a copy—that most programs avoid it.)

- 実装は雰囲気で読んだが、コンパイラの部分は全然わかっていない
    - AST から SSA を作るフェーズがあり、`append` などのビルトイン関数はそのときに処理しているようだった
    - [SSA](https://en.wikipedia.org/wiki/Static_single_assignment_form) は今回初めて知った
- `golang/go` の compile 以下の [`gc (src/cmd/compile/internal/gc)`](https://github.com/golang/go/tree/master/src/cmd/compile/internal/gc) は Garbage Collector ではなく Go Compiler の略らしい
    - [go/src/cmd/compile at master · golang/go](https://github.com/golang/go/tree/master/src/cmd/compile) に以下の記述があった
        - `It should be clarified that the name "gc" stands for "Go compiler", and has little to do with uppercase "GC", which stands for garbage collection.`
    - 紛らわしかった

## 参考

- [Go Slices: usage and internals \- The Go Blog](https://blog.golang.org/slices-intro)
- [The Go Programming Language Specification \- The Go Programming Language](https://golang.org/ref/spec)
    - [Slice types](https://golang.org/ref/spec#Slice_types)
    - [Length and capacity](https://golang.org/ref/spec#Length_and_capacity)
    - [Making slices, maps and channels](https://golang.org/ref/spec#Making_slices_maps_and_channels)
    - [Appending to and copying slices](https://golang.org/ref/spec#Appending_and_copying_slices)
- [Arrays, slices \(and strings\): The mechanics of 'append' \- The Go Blog](https://blog.golang.org/slices)
- [research\!rsc: Go Data Structures](https://research.swtch.com/godata)
- [Introduction to the Go compiler - go/src/cmd/compile at master · golang/go](https://github.com/golang/go/tree/master/src/cmd/compile)

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://images-na.ssl-images-amazon.com/images/I/41meaSLNFfL._SX382_BO1,204,203,200_.jpg" alt="プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></div><div class="amazlet-detail">Alan A.A. Donovan (著), Brian W. Kernighan (著), 柴田 芳樹 (翻訳)<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

