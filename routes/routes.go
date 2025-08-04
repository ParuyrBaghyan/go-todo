package routes

import "github.com/gin-gonic/gin"

func RegiseteRouters(server *gin.Engine) {

	server.GET("/todos", getAllTodos)

	server.GET("/todos/:id", getTodo)

	server.POST("/createTodo", createTodo)

	server.PUT("/updateTodo/:id", updateTodo)

	server.DELETE("/deleteTodo/:id", deleteTodo)

	server.POST("/signup", signUp)

	server.POST("/signin", signIn)


}
