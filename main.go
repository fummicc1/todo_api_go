package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	todoGroup := router.Group("api/v1/todo")
	{
		todoGroup.GET("/", getToDo)
	}
	err := InitializeDatabase()
	if err != nil {
		log.Fatal(err)
	}

}

func getToDo(c *gin.Context) {
}
