/*
【TRY】単体テストを行おう
偶数判定関数􏰀テスト
・偶数が与えられた場合にtrueを返す関数を考える
・テーブル駆動テストを行う
・テストケースとしてふさわしいも􏰀のを考えテストを書こう
・各テストケースをサブテストとして実行しよう
・coverprofileをみてみよう
*/

package isEven

func IsEven(num int) bool{ //偶数の時trueを返す
	return num%2 == 0
}

/*
func main(){
	flag.Parse()
	args := flag.Args()

	//文字列を数字に変換
	num := toInt(args)

	//数字の偶奇判定
	result := evenOdd(num)
	fmt.Println(result)
}
 */

/*
func toInt(args []string) int{
	for _, str := range args {
		num, _ := strconv.Atoi(str)
		return num
	}
	return 0 //？？
}
 */



//疑問
//fmt.scan()とos.stdin()の違い

//フィードバック
//命名 toInt分かりづらい
//if文は作らず「return num%2 == 0」でok
