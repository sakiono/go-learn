/*
【TRY】電話帳を作ろう
■ 次の仕様を満たすコマンドラインツールを作ろう
	● ID、名前、電話番号を保持できるツール
		○ IDはAUTOINCREMENT
	● プログラムを再起動してもデータが保持される
	● 起動すると現在登録されている情報がすべて表示される
	● その後入力モードになり、1人分の情報を入力する
	● 1人分を入力するごとに現在保存されている情報をすべて表示する
	● データベースにはSQLiteを用いる
■ 余裕があれば改造する
	● IDを指定してデータを更新できるようにする
*/

package main

import (
	"database/sql"
	"fmt"
	"os"

	//_ "modernc.org/sqlite"
	"github.com/tenntenn/sqlite"
)

//const DriverName = "sqlite"

type Record struct {
	ID    int64
	Name  string
	Phone int64
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}
}

func run() error {
	db, err := sql.Open(sqlite.DriverName, "addressbook.db") //goから直接sql呼び出すのが大変(,やり取りするデータベース新規or既存)
	if err != nil {
		return err
	}

	if err := createTable(db); err != nil {
		return err
	}

	for {
		if err := showRecord(db); err != nil {
			return err
		}
		if err := inputRecord(db); err != nil {
			return err
		}
	}
	return nil

}

func createTable(db *sql.DB) error {
	const sql = `
	CREATE TABLE IF NOT EXISTS addressbook (
			id   INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			phone INTEGER NOT NULL
	);`

	if _, err := db.Exec(sql); err != nil {
		return err
	}
	return nil
}

//表示する
func showRecord(db *sql.DB) error {
	fmt.Println("------------")
	fmt.Println("全件表示")
	fmt.Println("------------")

	rows, err := db.Query("SELECT * FROM addressbook")
	if err != nil {
		return err
	}
	for rows.Next() {
		var r Record
		if err := rows.Scan(&r.ID, &r.Name, &r.Phone); err != nil {
			return err
		}
		fmt.Println(r) //これできないかも
	}
	fmt.Println("------------")

	return nil
}

//入力する
func inputRecord(db *sql.DB) error {

	var r Record

	fmt.Print("Name >")
	fmt.Scan(&r.Name)
	fmt.Print("TEL >")
	fmt.Scan(&r.Phone)

	const sql = `
	INSERT INTO addressbook(name, phone) values (?, ?)
	`

	_, err := db.Exec(sql, r.Name, r.Phone)
	if err != nil {
		return err
	}

	return nil
}

//memos
/*
104行、fmt.Scanの時、&rにするのは何でだっけ
*/
