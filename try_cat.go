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
)

func main() {
	boo := flag.Bool("n", false, "オプション設定の有無")
	flag.Parse()

	memo := flag.Args() //フラグの分を除外してプログラム引数を取得 //os.Argsはフラグ,実行ファイル名含め全ての要素取得
	//flag.Parse()

	for num, text := range memo {
		if *boo == true {
			fmt.Printf("%d: %s\n", num+1, text)
		} else {
			fmt.Printf("%s\n", text)
		}
	}

}

//scan系との違い
//22行目と19行目のflag.Parseの違い
//catで打てるためにはどうしたらいいのか
