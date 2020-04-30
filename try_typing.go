/*
【TRY】タイピングゲームを作ろう
■ 標準出力に英単語を出す(出すも􏰀􏰂のは自由) ■ 標準入力から1行受け取る
■ 制限時間内に何問解けたか表示する

//ヒント
制限時間にはtime.After関数を使う
	context.Withtimeoutでも良い
select構文を用いる
	制限時間と入力を同時に待つ
 */

package main

import(
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
	"context"
)

var members = []string{"uehara","ono","okazawa","tsuji","ito","inoue","sakuma","nakagawa","tanaka","tsuchida"}
var point = 0

func judge(word string, i int) string{
	if word == members[i] {
		point += 1
		return "正解:)"
	}else{
		return "残念;("
	}
}

func input(r io.Reader, ch chan string) <-chan string {

	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text() //チャネルに読み込んだ文字列を送る
		}
		close(ch) //受信側で検知できる、受信側に終了を伝えるために使う //チャネルを閉じる
	}()
	return ch //チャネルを返す
}

func typingLoop(ctx context.Context){
	typingCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	ch := make(chan string)

	for i := 0; i < len(members); i++{

		//タイピング処理
		fmt.Println(">" + members[i])
		fmt.Print(">")
		input(os.Stdin,ch)
		word := <-ch //受信するまで待機 //受信したものを代入
		fmt.Println(judge(word, i))

		//キャンセルされた時終了
		select{
			case <- typingCtx.Done():
				return //typingLoop関数終了
			default:
		}

	}
}

func main() {

	bc := context.Background()
	tm := 15 * time.Second
	ctx, cancel := context.WithTimeout(bc,tm) //15秒待ってキャンセルする
	defer cancel()

	go typingLoop(ctx)

	select{
	case <-ctx.Done():
		//キャンセルされた時の処理
		fmt.Printf("15秒間で %d人 入力できました！\n", point) //得点表示
		return //main関数終了
	}
}


/*
//疑問
mainとtypingloop、両方の関数でselect-caseしなきゃいけない？
全体的な構造はこれでよい？
標準入力の読み込み方合ってる？
input関数のclose(),returnはこの場合必要？
小さい機能を作る→組み合わせるで作ったが、本当はどうやって作ったらいいか？

//memo
//入力された時の挙動が違う
//caseの条件、入力されたときに動くようにする
//
 */
