package api

import (
	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
)

// InspectTask - returns task information
func InspectTask(c *gin.Context) {
	ctx, cli, err := FetchContextAndClient()
	if err != nil {
		c.JSON(200, err)
	}
	taskID := c.Param("taskID")
	task, _, err := cli.TaskInspectWithRaw(ctx, taskID)
	c.JSON(200, task)
}

// ListTasks - returns a list of all tasks
func ListTasks(c *gin.Context) {
	ctx, cli, err := FetchContextAndClient()
	if err != nil {
		c.JSON(200, err)
	}
	task, err := cli.TaskList(ctx, types.TaskListOptions{})
	c.JSON(200, task)
}
