package main

import (
	"go-todo/db"
	"go-todo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	routes.RegiseteRouters(server)

	server.Run(":8080")
}