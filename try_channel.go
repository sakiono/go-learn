/*
【TRY】チャンネルを使ってみよう
次􏰀のコード􏰀のTODOを実装してみよう
標準入力から受け取った文字列を出力するコードになります
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func input(r io.Reader) <-chan string { //io.Readerインタフェース

	// TODO: チャネルを作る
	ch1 := make(chan string)

	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			// TODO: チャネルに読み込んだ文字列を送る
			ch1 <- s.Text()
		}
		// TODO: チャネルを閉じる
		close(ch1)
	}()
	// TODO: チャネルを返す
	return ch1
}

func main() {
	ch := input(os.Stdin)
	for {
		fmt.Print(">")
		fmt.Println(<-ch) //空だから、送信くるまでブロックされている
	}
}

/*
//memo

ch1 <-chan int //受信専用
ch2 chan<- int //送信専用

//疑問
capacity0でも、受信できるのはなぜ？
ch1 := make(chan string)このcapacityは？
ch1 := make(chan string)とvar ch2 chan stringの違いは？
処理の流れがよく分からない
*/