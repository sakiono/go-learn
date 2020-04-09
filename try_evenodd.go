//奇数偶数判定関数を作成
//以下のプログラムを条件式で利用

package main

func main() {
	for i := 1; i <= 100; i++ {
		print(i)
		if evenodd(i) {
			println("-偶数")
		} else {
			println("-奇数")
		}
	}
}

func evenodd(num int) bool {
	return num%2 == 0
}
