package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kristoy0/receptacle-worker/api"
)

func main() {
	r := gin.Default()
	service := r.Group("/api/service")
	{
		service.GET("/", api.ListServices)
		service.GET("/inspect/:id", api.InspectService)
		service.POST("/create", api.CreateService)
		service.POST("/update/:id", api.UpdateService)
		service.POST("/remove/:id", api.RemoveService)
	}

	r.Run(":8080")
}
