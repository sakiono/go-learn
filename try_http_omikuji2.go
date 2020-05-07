/*
【TRY】リクエストパラメタ􏰀の使用
■ リクエストパラメタを取得しよう
● http://localhost:8080?p=Gopher というアクセスがあった場合
	Gopherさん􏰀の運勢􏰁は「大吉」です!
	のように表示してください
*/

package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

//クリエストが来たらおみくじ結果を返す
func handler3(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("p")

	if name == ""{
		name = "ゲスト"
	}

	fmt.Fprintf(w,"%sさんの運勢は「%s」です！\n", name,omikuji3())

}

func omikuji3() string {
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
	http.HandleFunc("/", handler3)

	//サーバ起動
	http.ListenAndServe(":8080", nil)

}

/*
//疑問
 デフォルトで設定もっと上手いやり方ある？
 */