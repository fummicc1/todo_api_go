package controller

import (
	"database/sql"
	"log"
	"time"

	"github.com/fummicc1/todo_api_go/model"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

// InitializeDatabase sets up database.
func InitializeDatabase() error {
	// データベースを開く
	// paramのparseTimをtrueにしてあげると、MySQLのTime周りをtime.Time型にパースしてくれる
	db, err = sql.Open("mysql", "fummicc1:fummicc1@tcp(127.0.0.1:3306)/sample_todo_db?parseTime=true&loc=Asia%2FTokyo")
	return err
}

// GetToDo : fetch todo from mysql.
func GetToDo() (result []model.Todo, err error) {
	var rows *sql.Rows
	rows, err = db.Query("SELECT * from " + GetToDoTableName())

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
	todo.Due = time.Now()
	var stmt *sql.Stmt
	stmt, err = db.Prepare("INSERT INTO " + GetToDoTableName() + " VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	var result sql.Result
	result, err = stmt.Exec(todo.ID, todo.Title, todo.Content, todo.Due)
	_, err = result.LastInsertId()
	return err
}
