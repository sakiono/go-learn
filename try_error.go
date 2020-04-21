/*
//【TRY】エラー処理をしてみよう
★Stringerインタフェースに変換する関数を作ろう
・任意􏰀値をStringer型に変換する関数
	type Stringer{String() string}
・引数にempty interfaceを取り、Stringerとエラーを返す
	func ToStringer(v interface{}) (Stringer, error)
・返すエラー型􏰁errorインタフェースを実装したユーザ定義型にする
・実際に呼び出してみてエラー処理をしてみよう
	エラーが発生した場合􏰁エラーが発生した旨を表示する
*/

package main

import (
	"fmt"
	"os"
)

type Stringer interface{
	String() string
}

func ToStringer(v interface{}) (Stringer, error){ //S型が送られて来ているが受取可能
	if s, ok := v.(Stringer); ok{ //v.(Stringer)が分からない //okがtrueの時の処理
		return s,nil
	}
	return nil, MyError("CastError")
}

type MyError string

func (e MyError) Error() string{
	return string (e)
}

type S string

func (s S) String() string{
	return string(s)
}

func main(){
	v := S("Hoge") //string型"Hoge"をS型にキャスト
	if s, err :=ToStringer(v); err != nil{ //ToStringer(v)の結果をe,errに代入
		fmt.Fprintln(os.Stderr, "ERROR:",err)
	}else{
		fmt.Println(s.String())
	}
}

//25行目v.(Stringer)