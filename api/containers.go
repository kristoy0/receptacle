package api

import (
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
)

// InspectContainer - inspects a container
func InspectContainer(c *gin.Context) {
	ctx, cli, err := FetchContextAndClient()
	if err != nil {
		c.JSON(200, err)
	}
	containerID := c.Param("containerID")
	cont, err := cli.ContainerInspect(ctx, containerID)
	if err != nil {
		c.JSON(200, err)
	}
	c.JSON(200, cont)
}

// ListContainers - lists all containers
func ListContainers(c *gin.Context) {
	ctx, cli, err := FetchContextAndClient()
	if err != nil {
		c.JSON(200, err)
	}
	conts, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		c.JSON(200, err)
	}
	c.JSON(200, conts)
}

// RemoveContainer - deletes a container
func RemoveContainer(c *gin.Context) {
	ctx, cli, err := FetchContextAndClient()
	if err != nil {
		c.JSON(200, err)
	}
	containerID := c.Param("containerID")
	err = cli.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{Force: true})
	c.JSON(200, fmt.Sprintf("Container %s removed", containerID))
}
