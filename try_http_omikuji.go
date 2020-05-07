/*
【TRY】おみくじWebアプリの作成
	Webアプリ化してみよう ok
	HTTPサーバを作成する ok
	リクエストが来たらおみくじの結果を返す ok
	乱数の種は1回だけ初期化する
		HTTPサーバを起動する前に初期化する
*/

package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

//クリエストが来たらおみくじ結果を返す
func handler2(w http.ResponseWriter, r *http.Request) {

	randomNum := rand.Intn(5)
	fmt.Println(randomNum)

	switch randomNum {
	case 0:
		fmt.Fprint(w, "今日の運勢は凶です")
	case 1:
		fmt.Fprint(w, "今日の運勢は吉です")
	case 2:
		fmt.Fprint(w, "今日の運勢は小吉です")
	case 3:
		fmt.Fprint(w, "今日の運勢は中吉です")
	case 4:
		fmt.Fprint(w, "今日の運勢は大吉です")
	}

}

func main() {
	//ルーティング
	http.HandleFunc("/", handler2) //ここで何か処理はしてる？

	//乱数の種の初期化 //1回だけ //サーバ起動前
	rand.Seed(time.Now().UnixNano())

	//サーバ起動
	http.ListenAndServe(":8080", nil)

}

/*
//疑問
ページを更新すると、randomNumの値が更新されるのは、どう言う処理手順？
 */