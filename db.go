package main

import (
	"database/sql"
	"log"
)

var db *sql.DB

func initializeDatabase() (err error) {
	// データベースを開く
	// paramのparseTimをtrueにしてあげると、MySQLのTimeをtime.Time型にパースしてくれる
	db, err = sql.Open("mysql", "fummicc1:fummicc1@/sample_todo_db?parseTime=true")
	return
}

// GetToDo : fetch todo from mysql.
func GetToDo() (result []Todo, err error) {
	rows, err := db.Query("SELECT * from " + GetToDoTableName())

	for rows.Next() {
		todo := &Todo{}
		if err = rows.Scan(&todo.ID, &todo.Title, &todo.Content, &todo.Due); err != nil {
			log.Fatal(err)
		}
		result = append(result, *todo)
	}
	return
}

// GetToDoTableName returns name of todo talbe.
func GetToDoTableName() string {
	return "todos"
}
