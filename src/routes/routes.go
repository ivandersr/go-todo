package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ivandersr/go-todo/src/controllers"
)

func Routes() {
	route := gin.Default()

	route.POST("/todo", controllers.CreateTodo)
	route.GET("/todo", controllers.GetAllTodos)
	route.PUT("/todo/:id", controllers.UpdateTodo)
	route.DELETE("/todo/:id", controllers.DeleteTodo)

	route.Run()
}
