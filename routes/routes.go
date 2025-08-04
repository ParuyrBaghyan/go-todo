package routes

import (
	"go-todo/middlewares"

	"github.com/gin-gonic/gin"
)

func RegiseteRouters(server *gin.Engine) {
	server.GET("/todos", getAllTodos)
	server.GET("/todos/:id", getTodo)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/createTodo", createTodo)
	authenticated.PUT("/updateTodo/:id", updateTodo)
	authenticated.DELETE("/deleteTodo/:id", deleteTodo)
	authenticated.POST("/todos/:id/register", registerForTodo)
	authenticated.POST("/todos/:id/cancelRegister", cancelRegistration)

	server.POST("/signup", signUp)
	server.POST("/signin", signIn)

}
