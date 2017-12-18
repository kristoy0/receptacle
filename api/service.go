package api

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"github.com/gin-gonic/gin"
)

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

func UpdateService(c *gin.Context) {
	ctx, cli, err := FetchContextAndClient()
	if err != nil {
		c.JSON(200, err)
	}
	var service swarm.ServiceSpec
	id := c.Param("id")
	svc, _, err := cli.ServiceInspectWithRaw(ctx, id)
	if err != nil {
		c.JSON(200, err)
	}
	c.BindJSON(&service)
	cli.ServiceUpdate(ctx, id, svc.Version, service, types.ServiceUpdateOptions{})

}

func InspectService(c *gin.Context) {
	ctx, cli, err := FetchContextAndClient()
	if err != nil {
		c.JSON(200, err)
	}
	id := c.Param("id")
	svc, _, err := cli.ServiceInspectWithRaw(ctx, id)
	if err != nil {
		c.JSON(200, err)
	}
	c.JSON(200, svc)
}

func ListServices(c *gin.Context) {
	ctx, cli, err := FetchContextAndClient()
	if err != nil {
		c.JSON(200, err)
	}
	svc, err := cli.ServiceList(ctx, types.ServiceListOptions{})
	c.JSON(200, svc)
}

func RemoveService(c *gin.Context) {
	ctx, cli, err := FetchContextAndClient()
	if err != nil {
		c.JSON(200, err)
	}
	id := c.Param("id")
	err = cli.ServiceRemove(ctx, id)
}
