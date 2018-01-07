package main

import (
	"context"
	"log"

	proto "github.com/kristoy0/receptacle/server/proto"
	micro "github.com/micro/go-micro"
)

type TaskHandler interface {
	Deploy(context.Context, *proto.DeployRequest, *proto.DeployResponse) error
}

type Task struct{}

func (t *Task) Deploy(ctx context.Context, req *proto.DeployRequest, res *proto.DeployResponse) error {
	res.Status = "Vist töötab"
	log.Println(req)
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("receptacle-server"),
		micro.Version("latest"),
	)

	service.Init()

	proto.RegisterTasksHandler(service.Server(), new(Task))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
