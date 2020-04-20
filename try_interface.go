/*
【TRY】インタフェースを作ろう
Stringerインタフェース
・String() string メソッドを持つインタフェースを作る ok
・そして3つ以上Stringerインタフェースを実装する型を作る ng
インタフェースを受け取る関数
・Stringerインタフェースを引数で受け取る関数を作る
・受け取った値を上記の3つの具象型によって分岐し、具象型の型名と値を表示する
*/

package main

import "fmt"

type Stringer interface {
	String() string
}

type I int

func (i I) String() string { //メソッドの定義
	return "I"
}

type B bool

func (b B) String() string { //メソッドの定義
	return "B"
}

type S string

func (s S) String() string { //メソッドの定義
	return "S"
}

func main() {
	var n I = I(100) //I型なら代入可能、I型にcastして代入
	F(n)             //
	F(B(true))
	F(S("saki"))
}

func F(s Stringer) {
	switch v := s.(type) {
	case I:
		fmt.Println(int(v), "I")
	case B:
		fmt.Println(bool(v), "B")
	case S:
		fmt.Println(string(v), "S")
	}
}

//46行目以降、caseがI,B,Sに分かれるが、v := s.(type)で何でこの結果になる？
//上の方、String()関数が3つあるが、名前同じでどうやってそれぞれの関数呼び出してる？いつこれ呼び出されてる？
//処理されている順番が分からない
