package api

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"github.com/gin-gonic/gin"
)

// CreateService - creates a new docker service
func CreateService(c *gin.Context) {
	ctx, cli, err := FetchContextAndClient()
	if err != nil {
		c.JSON(200, err)
	}
	var service swarm.ServiceSpec
	c.BindJSON(&service)
	cli.ServiceCreate(ctx, service, types.ServiceCreateOptions{})
	c.JSON(200, service)
}

// UpdateService - updates an existing service
func UpdateService(c *gin.Context) {
	ctx, cli, err := FetchContextAndClient()
	if err != nil {
		c.JSON(200, err)
	}
	var service swarm.ServiceSpec
	serviceID := c.Param("serviceID")
	svc, _, err := cli.ServiceInspectWithRaw(ctx, serviceID)
	if err != nil {
		c.JSON(200, err)
	}
	c.BindJSON(&service)
	cli.ServiceUpdate(ctx, serviceID, svc.Version, service, types.ServiceUpdateOptions{})

}

// InspectService - inspects a service
func InspectService(c *gin.Context) {
	ctx, cli, err := FetchContextAndClient()
	if err != nil {
		c.JSON(200, err)
	}
	serviceID := c.Param("serviceID")
	svc, _, err := cli.ServiceInspectWithRaw(ctx, serviceID)
	if err != nil {
		c.JSON(200, err)
	}
	c.JSON(200, svc)
}

// ListServices - lists all services
func ListServices(c *gin.Context) {
	ctx, cli, err := FetchContextAndClient()
	if err != nil {
		c.JSON(200, err)
	}
	svc, err := cli.ServiceList(ctx, types.ServiceListOptions{})
	c.JSON(200, svc)
}

// RemoveService - removes a service
func RemoveService(c *gin.Context) {
	ctx, cli, err := FetchContextAndClient()
	if err != nil {
		c.JSON(200, err)
	}
	serviceID := c.Param("serviceID")
	err = cli.ServiceRemove(ctx, serviceID)
}
