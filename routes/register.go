package routes

import (
	"go-todo/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForTodo(context *gin.Context) {
	userId := context.GetInt64("userId")
	todoId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse todo id"})
		return
	}

	todo, err := models.GetTodoById(todoId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch todo."})
		return
	}

	err = todo.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for event."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered"})

}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	todoId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse todo id"})
		return
	}

	var todo models.Todo
	todo.Id = todoId

	err = todo.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration user for todo."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Cancelled"})

}
