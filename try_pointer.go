//【TRY】ポインタ
//値を入れ替えるswap2関数を実装してください
//次のコードが正しく動作するように実装してください

package main

func main() {
	n, m := 10, 20
	swap2(&n, &m) //n,mのアドレスを引数に指定
	println(n, m)
}

func swap2(np, mp *int) { //アドレスをintのポイント型にいれる
	*np, *mp = *mp, *np //アドレスが指している変数を、入れ替える
}
