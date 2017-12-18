package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kristoy0/receptacle-worker/api"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		// Services
		v1.GET("/service", api.ListServices)
		v1.GET("/service/:serviceID", api.InspectService)
		v1.POST("/service/create", api.CreateService)
		v1.PUT("/service/:serviceID", api.UpdateService)
		v1.DELETE("/service/:id", api.RemoveService)

		// Containers
		v1.GET("/container", api.ListContainers)
		v1.GET("/container/:containerID", api.InspectContainer)
		v1.DELETE("/container/:containerID", api.DeleteContainer)
	}

	r.Run(":8080")
}
