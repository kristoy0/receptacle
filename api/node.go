package api

import (
	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
)

// InspectNode - inspects a node
func InspectNode(c *gin.Context) {
	ctx, cli, err := FetchContextAndClient()
	if err != nil {
		c.JSON(200, err)
	}
	nodeID := c.Param("nodeID")
	node, _, err := cli.NodeInspectWithRaw(ctx, nodeID)
	if err != nil {
		c.JSON(200, err)
	}
	c.JSON(200, node)
}

// ManagerInfo - shows info from the docker host
func ManagerInfo(c *gin.Context) {
	ctx, cli, err := FetchContextAndClient()
	if err != nil {
		c.JSON(200, err)
	}
	info, err := cli.Info(ctx)
	if err != nil {
		c.JSON(200, err)
	}
	c.JSON(200, info)

}

// ListNodes - lists all nodes
func ListNodes(c *gin.Context) {
	ctx, cli, err := FetchContextAndClient()
	if err != nil {
		c.JSON(200, err)
	}
	nodes, err := cli.NodeList(ctx, types.NodeListOptions{})
	if err != nil {
		c.JSON(200, err)
	}
	c.JSON(200, nodes)
}
