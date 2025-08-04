package routes

import (
	"go-todo/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getAllTodos(context *gin.Context) {
	todos, err := models.GetAllTodos()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch todos.Try again later."})
		return
	}

	context.JSON(http.StatusOK, todos)

}

func getTodo(context *gin.Context) {
	todoId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse todo id"})
	}

	todo, err := models.GetTodoById(todoId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request data"})
	}

	context.JSON(http.StatusOK, todo)
}

func createTodo(context *gin.Context) {
	var todo models.Todo
	err := context.ShouldBindJSON(&todo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	todo.UserId = 3

	err = todo.CreateTodo()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create todo.Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Todo created", "todoId": todo})
}

func updateTodo(context *gin.Context) {
	todoId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse event id"})
		return
	}

	_, err = models.GetTodoById(todoId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event."})
		return
	}

	var updateTodo models.Todo
	err = context.ShouldBindJSON(&updateTodo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	updateTodo.Id = todoId

	err = updateTodo.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update the todo."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully!"})
}

func deleteTodo(context *gin.Context) {
	todoId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse todo id"})
		return
	}

	todo, err := models.GetTodoById(todoId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the todo."})
		return
	}

	err = todo.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the todo."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully!"})
}
