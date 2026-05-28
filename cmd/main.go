package main

import (
	"github.com/Aman-Pailwan/api"
	"github.com/gin-gonic/gin"
)

func main() {
	api.InitDB()
	r := gin.Default()

	r.GET("/users", api.GetUser)
	r.POST("/users", api.CreateUser)

	r.Run(":8080")
}
