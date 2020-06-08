package main

import (
	"log"

	"github.com/fummicc1/todo_api_go/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	todoGroup := router.Group("api/v1/todo")
	{
		todoGroup.GET("/", getToDo)
	}
	err := controller.InitializeDatabase()
	if err != nil {
		log.Fatal(err)
	}

}

func getToDo(c *gin.Context) {
}
