/*
【TRY】画像変換コマンドを作ろう
▼次の􏰁仕様を満たすコマンドを作って下さい
ディレクトリを指定する //flag
指定したディレクトリ以下􏰁JPGファイルをPNGに変換(デフォルト)　？
ディレクトリ以下􏰀再帰的に処理する　
変換前と変換後􏰁画像形式を指定できる(オプション) //flag
▼以下を満たすように開発してください
mainパッケージと分離する
自作パッケージと標準パッケージと準標準パッケージ􏰁のみ使う
準標準パッケージ:golang.org/x以下􏰁パッケージ
ユーザ定義型を作ってみる
GoDocを生成してみる
*/

package main

import (
	"flag"
	"os"
	"fmt"
	"path/filepath"
	//"try_changepicture/traverse"
)

//go -> png
var before = flag.String("before", ".go", "変換前画像形式")
var after = flag.String("after", ".png", "変換後画像形式")

func main(){
	flag.Parse()
	dir := flag.Args() //ディレクトリ //相対パス指定

	for _, path := range dir {
		// ファイルを探し出す
		if err := filepath.Walk(path, func(path string,info os.FileInfo, err error)error{
			if filepath.Ext(path) == *before{
				fmt.Println(path)
			}
			return nil
		}); err != nil {
			fmt.Println(err)
		}
	}

}

//filepath.Walk(ディレクトリのパス string, ファイルまたはディレクトリが見つかったときに行う関数)

//メモ
//まだgoのファイルを探し出して表示する動作しかできていない
//traverseディレクトリは現状あってもなくても同じ

//知りたいこと
//指定したディレクトリ以下􏰁JPGファイルをPNGに変換
