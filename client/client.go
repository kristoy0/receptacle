package main

import (
	"context"
	"log"

	proto "github.com/kristoy0/receptacle/server/proto"
	micro "github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(micro.Name("receptacle-client"))

	tasks := proto.NewTasksClient("receptacle-server", service.Client())

	res, err := tasks.Deploy(context.TODO(), &proto.DeployRequest{
		Name:    "test-container",
		Image:   "python:latest",
		Command: []string{"python", "app.py"},
		Resources: &proto.Resources{
			Memory:    "512MB",
			CPU:       "0.5",
			Instances: 5,
			Volumes: []string{
				"/home/kristo/pycode:/app/",
			},
		},
		Env: []string{
			"FOO=bar",
			"BAR=foo",
		},
		Nodes: []string{
			"node-1", "node-2",
		},
	})
	if err != nil {
		log.Println(err)
	}

	log.Println(res)
}
