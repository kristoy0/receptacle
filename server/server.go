package main

import (
	"context"
	"log"
	"time"

	proto "github.com/kristoy0/receptacle/server/proto"
	micro "github.com/micro/go-micro"
)

type TaskHandler interface {
	Deploy(context.Context, *proto.DeployRequest, *proto.DeployResponse) error
	Undeploy(context.Context, *proto.UndeployRequest, *proto.UndeployResponse) error
	List(context.Context, *proto.ListRequest, *proto.ListResponse) error
}

type Task struct{}

func (t *Task) Deploy(ctx context.Context, req *proto.DeployRequest, res *proto.DeployResponse) error {
	res.Status = "Job" + req.Name + "deployed"
	service := getService()
	var err error
	for _, val := range req.Nodes {
		tasks := proto.NewTasksClient(val, service.Client())
		res, err = tasks.Deploy(ctx, req)
		if err != nil {
			return err
		}
	}
	return nil
}

func (*Task) Undeploy(ctx context.Context, req *proto.UndeployRequest, res *proto.UndeployResponse) error {
	res.Status = "Job" + req.Name + "undeployed"
	return nil
}

func (*Task) List(ctx context.Context, req *proto.ListRequest, res *proto.ListResponse) error {
	return nil
}

func main() {
	service := getService()

	service.Init()

	proto.RegisterTasksHandler(service.Server(), new(Task))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func getService() micro.Service {
	return micro.NewService(
		micro.Name("receptacle-server"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)
}
