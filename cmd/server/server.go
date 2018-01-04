package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kristoy0/receptacle/server"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.POST("/create", server.Create)
		v1.GET("/list", server.List)
		v1.GET("/list/:id", server.Inspect)
	}
	router.Run(":5237")
}
