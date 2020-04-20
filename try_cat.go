//catコマンド
/*
作成するcatコマンドの仕様
・引数でファイルパスの一覧を貰い、そのファイルを与えられた順に標準出力にそのまま出力するコマンドを作ってください。
・また、-nオプションを指定すると、行番号を各行につけて表示されるようにしてください。
・なお、行番号のすべてのファイルで通し番号にしてください。
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"bufio"
)
var boo = flag.Bool("n", false, "オプション設定の有無")
var num = 1

func main() {
	flag.Parse()

	fileNames := flag.Args() //フラグの分を除外してプログラム引数を取得 //os.Argsはフラグ,実行ファイル名含め全ての要素取得

	for _, fn := range fileNames {
		err := readLine(fn) //エラーの場合の受け取るための変数

		if err != nil { //エラーの場合の処理
			fmt.Println(err)
			return
		}
	}
}

func readLine (fn string) error {
	f, err := os.Open(fn) //ファイル開ける

	if err != nil { //エラーならerrを返す//終了
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if *boo == true {
			fmt.Printf("%d: ",num)
			num++
		}
		fmt.Println(scanner.Text())
	}
	return nil
}


//「catコマンド」を作る

//bufio.newscannerとscanner.scanが何してるのかいまいちわからない
//画像
//-nの設定がないと動かないのは何で
//52はreturn nilでいいのか
//30のreturnはいる？

//22行目と19行目のflag.Parseの違い ok
//catで打てるためにはどうしたらいいのか ok
