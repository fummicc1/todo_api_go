package main

import (
	"log"
	"net/http"

	"github.com/fummicc1/todo_api_go/controller"
	"github.com/fummicc1/todo_api_go/model"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	todoGroup := router.Group("api/v1/todo")
	{
		todoGroup.GET("/", getToDos)
		todoGroup.POST("/", createToDo)
	}
	err := controller.InitializeDatabase()
	if err != nil {
		log.Fatal(err)
	}
	router.Run(":8080")
}

func getToDos(c *gin.Context) {
	todos, err := controller.GetToDo()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, todos)
	}
}

func createToDo(c *gin.Context) {
	todo := model.Todo{}
	err := c.BindJSON(todo)
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, err.Error())
	}
	err = controller.AddToDo(todo)
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, err.Error())
	}
	c.String(http.StatusOK, "Success saving todo")
}
