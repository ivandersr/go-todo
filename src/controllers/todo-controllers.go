package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ivandersr/go-todo/src/config"
	"github.com/ivandersr/go-todo/src/models"
	"gorm.io/gorm"
)

var db *gorm.DB = config.ConnectDB()

type todoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type todoResponse struct {
	todoRequest
	ID uint `json:"id"`
}

func CreateTodo(context *gin.Context) {
	var data todoRequest

	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{}
	todo.Name = data.Name
	todo.Description = data.Description

	result := db.Create(&todo)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	var response todoResponse
	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description

	context.JSON(http.StatusCreated, response)
}

func GetAllTodos(context *gin.Context) {
	var todos []models.Todo

	err := db.Find(&todos)

	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error getting data"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    todos,
	})
}

func UpdateTodo(context *gin.Context) {
	var data todoRequest

	reqParamId := context.Param("id")
	id, convErr := strconv.ParseUint(reqParamId, 10, 32)

	if convErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": convErr.Error()})
		return
	}

	if err := context.BindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{}

	todoById := db.Where("id = ?", id).First(&todo)
	if todoById.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Todo not found"})
		return
	}

	todo.Name = data.Name
	todo.Description = data.Description

	result := db.Save(&todo)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	var response todoResponse
	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description

	context.JSON(http.StatusCreated, response)
}

func DeleteTodo(context *gin.Context) {
	todo := models.Todo{}

	reqParamId := context.Param("id")
	id, convErr := strconv.ParseUint(reqParamId, 10, 32)

	if convErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": convErr.Error()})
		return
	}

	delete := db.Where("id = ?", id).Unscoped().Delete(&todo)
	fmt.Println(delete)

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    id,
	})
}
