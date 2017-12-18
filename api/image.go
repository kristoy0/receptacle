package api

import (
	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
)

// InspectImage - inspects an image
func InspectImage(c *gin.Context) {
	ctx, cli, err := FetchContextAndClient()
	if err != nil {
		c.JSON(200, err)
	}
	imageID := c.Param("imageID")
	img, _, err := cli.ImageInspectWithRaw(ctx, imageID)
	if err != nil {
		c.JSON(200, err)
	}
	c.JSON(200, img)
}

// ListImages - lists all images
func ListImages(c *gin.Context) {
	ctx, cli, err := FetchContextAndClient()
	if err != nil {
		c.JSON(200, err)
	}
	img, err := cli.ImageList(ctx, types.ImageListOptions{})
	c.JSON(200, img)
}

// RemoveImage - removes an image
func RemoveImage(c *gin.Context) {
	ctx, cli, err := FetchContextAndClient()
	if err != nil {
		c.JSON(200, err)
	}
	imageID := c.Param("imageID")
	resp, err := cli.ImageRemove(ctx, imageID, types.ImageRemoveOptions{Force: true})
	if err != nil {
		c.JSON(200, err)
	}
	c.JSON(200, resp)
}
