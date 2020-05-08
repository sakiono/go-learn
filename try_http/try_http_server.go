package main

import (
	"math/rand"
	"net/http"
	"html/template"
	"time"
)

type Result struct {
	Name    string
	Omikuji string
}

//クリエストが来たらおみくじ結果を返す
func handler4(w http.ResponseWriter, r *http.Request) {

	result := Result{
		Name:    r.FormValue("p"), //パラメータの値(名前)
		Omikuji: omikuji4(), //おみくじ結果
	}


	if result.Name == ""{
		result.Name = "ゲスト"
	}

	tmpl := template.Must(template.New("omikuji").Parse("<html><body>{{.Name}}さんの運勢は「<b>{{.Omikuji}}</b>」です </body></html>"))
	tmpl.Execute(w, result)

}

func omikuji4()string {
	randomNum := rand.Intn(5)

	switch randomNum {
	case 0:
		return "凶"
	case 1:
		return "吉"
	case 2:
		return "小吉"
	case 3:
		return "中吉"
	default:
		return "大吉"
	}

}

func main() {
	//乱数の種
	rand.Seed(time.Now().UnixNano())

	//ルーティング
	http.HandleFunc("/", handler4)

	//サーバ起動
	http.ListenAndServe(":8080", nil)

}