//値を入れ替えるswap関数を実装してください
//次のコードが正しく動作するように実装してください

package main

func main() {
	n, m := swap(10, 20)
	println(n, m)
}

func swap(x, y int) (x2, y2 int) {
	x2, y2 = y, x
	return
}
