/*
【TRY】トランザクションを使う
■ 送金処理を考えよう
	● user1からuser2に10円送金したいと考える
	● テーブルの状態は次のようになっている
	● トランザクションを使って実装しよう
*/

package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/tenntenn/sqlite"
)

func main() {

	//データベース作成
	db, err := sql.Open(sqlite.DriverName, "passbook.db") //goから直接sql呼び出すのが大変(,やり取りするデータベース新規or既存)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	//テーブル作成
	type Record struct {
		ID     int64
		Name   string
		Amount int64
	}

	const sql = `
	CREATE TABLE IF NOT EXISTS passbook (
			id   INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			amount INTEGER NOT NULL
	);`
	if _, err := db.Exec(sql); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	//値をinsert //初期追加済み
	/*
		records := []*Record{{Name: "user1", Amount: 100}, {Name: "user2", Amount: 20}}
		for i := range records {
			const sql = "INSERT INTO passbook(name, amount) values (?,?)"
			if _, err := db.Exec(sql, records[i].Name, records[i].Amount); err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			/*
				id, err := r.LastInsertId()
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					return
				}
				records[i].ID = id
	*/

	//}

	//表を表示
	fmt.Println("----送金前残高表示----")
	rows, err := db.Query("SELECT * FROM passbook")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	for rows.Next() {
		var r Record
		if err := rows.Scan(&r.ID, &r.Name, &r.Amount); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		fmt.Println(r)
	}

	//送金処理
	////トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	////user1から10円を引く
	if _, err := tx.Exec("UPDATE passbook SET amount = amount-10 WHERE id = 1"); err != nil {
		tx.Rollback()
		fmt.Fprintln(os.Stderr, err)
		return
	}

	////user2に10円を足す
	if _, err := tx.Exec("UPDATE passbook SET amount = amount+10 WHERE id = 2"); err != nil {
		tx.Rollback()
		fmt.Fprintln(os.Stderr, err)
		return
	}

	////コミット
	if err := tx.Commit(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	//残金を表示
	fmt.Println("----送金後残高表示----")
	rows, err = db.Query("SELECT * FROM passbook")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	for rows.Next() {
		var r Record
		if err := rows.Scan(&r.ID, &r.Name, &r.Amount); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		fmt.Println(r)
	}

}

/*
//memos
・insertだと実行するたびに追加される。一回だけ最初に追加する方法は？
・テーブル表示もう少し見やすくなる？
*/
