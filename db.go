package main

import (
	"database/sql"
	"log"
)

func initializeDatabase() {
	// データベースを開く
	// paramのparseTimをtrueにしてあげると、MySQLのTimeをtime.Time型にパースしてくれる
	db, err := sql.Open("mysql", "fummicc1:fummicc1@/sample_todo_db?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	// // 複数レコードの取得
	// rows, err := db.Query("SELECT * FROM todos")
	// // QueryRowで1レコードを取得できる
	// // db.QueryRow("SELECT * FROM todos WHERE id = ?LIMIT 1", 1)

	// // カラムを取得
	// columns, err := rows.Columns()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // スライスの作成
	// values := make([]sql.RawBytes, len(columns))

	// scanArgs := make([]interface{}, len(values))
	// for i := range values {
	// 	scanArgs[i] = &values[i]
	// }

	// for rwos.Next() {
	// }
}

func getToDo() []Todo, error {
}
