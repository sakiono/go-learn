package main

import (
	"fmt"
	"net/http"
)



//HTTPハンドラはインターフェースである
/*
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
 }
*/

//①1つのHTTPリクエストを処理する関数(ハンドラ)作成
/*
func handler(レスポンスを書き込む先,クライアントからのリクエスト){
	レスポンスの書き込み
}
*/
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, HTTPサーバ")
}


func main() {

	//②"/index"などのエントリポイントとハンドラを結びつける //これがルーティング
	http.HandleFunc("/", handler)

	//③ホスト名、ポート番号、ハンドラを指定してHTTPサーバを起動する
	http.ListenAndServe(":8080", nil)
}