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
	"net/http"
)

//クリエストが来たらおみくじ結果を返す
func handler3(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("p")
	
	if name == ""{
		name = "ゲスト"
	}

	fmt.Fprintf(w,"%sさんの運勢は「大吉」です！\n", name)

}

func main() {
	//ルーティング
	http.HandleFunc("/", handler3)

	//サーバ起動
	http.ListenAndServe(":8080", nil)

}

/*
//疑問
 デフォルトで設定もっと上手いやり方ある？
 */