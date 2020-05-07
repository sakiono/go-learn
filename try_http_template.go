/*
【TRY】テンプレートエンジン􏰀使用
html/templateを利用しよう
	http://localhost:8080?p=Gopher というアクセスがあった場合
		Gopherさん􏰀の運勢􏰁は「大吉」です!
	のように表示してください
	<html><body>Gopherさん􏰀の運勢􏰁は「<b>大吉</b>」です </body></html>
	テンプレート􏰀初期化􏰁1回だけ行いパッケージ変数にいれる
 */

package main

import (
	"net/http"
	"html/template"
)

//クリエストが来たらおみくじ結果を返す
func handler4(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("p")

	if name == ""{
		name = "ゲスト"
	}

	tmpl := template.Must(template.New("omikuji").Parse("<html><body>{{.}}さんの運勢は「<b>大吉</b>」です </body></html>"))
	tmpl.Execute(w, name)

}

func main() {
	//ルーティング
	http.HandleFunc("/", handler4)

	//サーバ起動
	http.ListenAndServe(":8080", nil)

}