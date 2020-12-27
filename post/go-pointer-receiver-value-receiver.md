{"title":"Go の pointer receiver, value receiver の挙動の整理","date":"2020-12-27T22:30:00+09:00"}

よくこんがらがるので改めてドキュメントを確認したログ。

[Tour of go の method の章](https://tour.golang.org/methods/1) によると、

- あるメソッドは、そのレシーバ (ポインタか値か) に指定された型に属する

```go
type t struct {}

// このメソッドは `*T` 型に属していて、`T` には属していない
func (t *T) fn() { /* ... */ }
```

- よってインタフェースを満たしているかのチェックや型アサーションではポインタ型かそうでないかの違いは区別される

```go
type I interface {
    fn()
}

tv := T{}   // T 型
tp := &T{}  // *T 型

var i I
i = tv // これはインタフェースを満たしておらず、コンパイルエラー
i = tp // これは OK
```

- ただし利便性のため、メソッドの呼び出しはポインタ型かそうでないかの違いをコンパイラが吸収する

```go
t := T{}
t.fn() // これは OK。コンパイラが (&t).fn() とみなしてくれる
```

とのことだった。

これらを [spec](https://golang.org/ref/spec) をななめ読みしながら確認していたたところ、「[ポインタ型 `*T` の method set には `T` のものも含まれる](https://golang.org/ref/spec#Method_sets)」という記載があった。これは知らなかった。

> The method set of the corresponding pointer type *T is the set of all methods declared with receiver *T or T (that is, it also contains the method set of T).

## spec の記載内容

まず前提として、ポインタ型とそのもととなる値の型は、それぞれ別の型として扱われる。

```
// Types https://golang.org/ref/spec#Types
Type      = TypeName | TypeLit | "(" Type ")" .
TypeName  = identifier | QualifiedIdent .
TypeLit   = ArrayType | StructType | PointerType | FunctionType | InterfaceType |
	    SliceType | MapType | ChannelType .

// Pointer type https://golang.org/ref/spec#Pointer_types
PointerType = "*" BaseType .  // BaseType と PointerType は別
BaseType    = Type .
```

https://play.golang.org/p/AX4T2T-9eep

```go
package main

import "fmt"

type BaseType struct{}

func main() {
	b := BaseType{}
	bp := &BaseType{}
	fmt.Printf("b: %T, bp: %T\n", b, bp) // b: main.BaseType, bp: *main.BaseType

	var v BaseType
	var vp *BaseType

	v = b
	v = bp // コンパイルエラー: cannot use bp (type *BaseType) as type BaseType in assignment

	vp = b // コンパイルエラー: cannot use b (type BaseType) as type *BaseType in assignment
	vp = bp
}
```

型には複数のメソッドを定義でき、その集合を [Method set](https://golang.org/ref/spec#Method_sets) という。

- ある型があるインタフェースを満たしているかどうかは、この method set を調べることで判断する
- ある型 `T` の method set には、レシーバが `T` のメソッドが入る
- ポインタ型の場合、`*T` の method set はレシーバが `*T` のメソッドだけでなく `T` のものも含まれる
    - ここを知らなかった。なぜこういう仕様なんだろう

https://play.golang.org/p/b6T4zOXVUlV

```go
package main

import "fmt"

type I interface {
	Str() string
}

type A struct{}

// レシーバは *A 型 (pointer receiver)
func (a *A) Str() string {
	return "a"
}

type B struct{}

// レシーバは B 型 (value receiver)
func (b B) Str() string {
	return "b"
}

func main() {
	var i I

	av := A{}
	ap := &A{}
	bv := B{}
	bp := &B{}
	fmt.Printf("av: %T, tp: %T, bv: %T, bp: %T\n", av, ap, bv, bp) // av: main.A, tp: *main.A, bv: main.B, bp: *main.B

	// A 型は Str メソッドを実装していないので、I インタフェースを満たしていない。
	i = av // コンパイルエラー: cannot use av (type A) as type I in assignment: A does not implement I (Str method has pointer receiver)

	// *A 型は I インタフェースを満たしている
	i = ap

	// B 型は I インタフェースを満たしている
	i = bv

	// *B 型の method set には、レシーバが `B` 型のメソッドも含まれているので、コンパイルエラーにならない
	i = bp
}
```

メソッドの呼び出しでは [ポインタ型のメソッドか値型かの違いを、利便性のためコンパイラが吸収する](https://tour.golang.org/methods/6)。

- `fn` は `T` 型のメソッドで、`t` という `*T` 型の値があった場合、本来 `(*t).fn()` となりそうだが `t.fn()` という呼び出しで OK (A)
- 反対に `fn` は `*T` 型のメソッドで、`t` という `T` 型の値があった場合、本来 `(&t).fn()` となりそうだが `t.fn()` という呼び出しで OK (B)
- 仕様を読む限り、恐らくこういうことらしい
    - 前述のように `*T` 型の method set には `T` 型のメソッドも入っている
        - これで上記 `(B)` のケースをカバーできる
    - spec の [Calls](https://golang.org/ref/spec#Calls) によると、`x.m()` という呼び出しで `x` が addressable かつ `&x` (ポインタ型のほう) の method set に `m` が含まれていれば、`x.m()` を `(&x).m()` として扱うとのこと
        - これで上記 `(A)` のケースをカバーできる

https://play.golang.org/p/XMeEX0lDNwo

```go
package main

import "fmt"

type A struct{}

// レシーバは *A 型 (pointer receiver)
func (a *A) Str() string {
	return "a"
}

type B struct{}

// レシーバは B 型 (value receiver)
func (b B) Str() string {
	return "b"
}

func main() {
	a := A{}                           // A 型
	b := &B{}                          // *B 型
	fmt.Printf("a: %T, b: %T\n", a, b) // a: main.A, b: *main.B

	fmt.Println(a.Str()) // コンパイラが (&a).Str() と補完
	fmt.Println(b.Str()) // コンパイラが (*b).Str() と補完
}
```

型アサーションではポインタか値かの違いをふまえて型を指定する必要がある。

- メソッド呼び出しとは異なり、ポインタ型・値型の違いをいいかんじに吸収してはくれない

https://play.golang.org/p/B3WHkVifbcZ

```go
package main

import "fmt"

type I interface {
	Str() string
}

type A struct{}

// レシーバは *A 型 (pointer receiver)
func (a *A) Str() string {
	return "a"
}

func (a *A) Str2() string {
	return "a2"
}

type B struct{}

// レシーバは B 型 (value receiver)
func (b B) Str() string {
	return "b"
}

func (b B) Str2() string {
	return "b2"
}

func main() {
	var i I

	a := &A{}                          // *A 型
	b := B{}                           // B 型
	fmt.Printf("a: %T, b: %T\n", a, b) // a: main.A, b: *main.B

	i = a
	// コンパイルエラー: impossible type assertion: A does not implement I (Str method has pointer receiver)
	// A 型はインタフェース I を満たしていない
	fmt.Println(i.(A).Str2())

	// *A 型はインタフェース I を満たしている
	fmt.Println(i.(*A).Str2())

	i = b
	// B 型はインタフェース I を満たしている
	fmt.Println(i.(B).Str2())
	// *B 型の method set に `Str()` が含まれインタフェース I を満たしているので、コンパイルエラーにならない
	fmt.Println(i.(*B).Str2())
}
```

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://images-na.ssl-images-amazon.com/images/I/41meaSLNFfL._SX382_BO1,204,203,200_.jpg" alt="プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></div><div class="amazlet-detail">Alan A.A. Donovan (著), Brian W. Kernighan (著), 柴田 芳樹 (翻訳)<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>
