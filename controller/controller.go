package controller

import (
	"database/sql"
	"log"

	"github.com/fummicc1/todo_api_go/model"
)

var db *sql.DB

// InitializeDatabase sets up database.
func InitializeDatabase() (err error) {
	// データベースを開く
	// paramのparseTimをtrueにしてあげると、MySQLのTimeをtime.Time型にパースしてくれる
	db, err = sql.Open("mysql", "fummicc1:fummicc1@/sample_todo_db?parseTime=true")
	return
}

// GetToDo : fetch todo from mysql.
func GetToDo() (result []model.Todo, err error) {
	rows, err := db.Query("SELECT * from " + GetToDoTableName())

	for rows.Next() {
		todo := &model.Todo{}
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

// AddToDo adds todo in database.
func AddToDo(todo model.Todo) error {
	result, err := db.Exec("INSERT INTO"+GetToDoTableName()+"VALUES(?, ?, ?, ?)", todo.ID, todo.Title, todo.Content, todo.Due)
	if err != nil {
		return err
	}
	_, err = result.LastInsertId()
	return err
}
