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
	if s, ok := v.(Stringer); ok{ //インタフェース型の値を任意の型にキャストする。第二引数はキャストできるかどうかの返り値 //okがtrueの時の処理
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
	if s, err :=ToStringer(v); err != nil{ //ToStringer(v)の返り値をs,errに代入
		fmt.Fprintln(os.Stderr, "ERROR:",err)
	}else{
		fmt.Println(s.String())
	}
}

//25行目v.(Stringer)は型アサーションをしている
//型アサーションは、そのinterfaceを実装しているかどうかのチェック
//第一引数アサーション後の型,第二引数成功したかどうかのbool型

/*
33行目
type error interface {
    Error() string
}
*/

//メソッドは、主語述語
//関数は、いろいろ書いてある文章