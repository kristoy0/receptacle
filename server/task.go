package main

import (
	"context"
	"log"

	proto "github.com/kristoy0/receptacle/server/proto"
)

type TaskHandler interface {
	Deploy(context.Context, *proto.DeployRequest, *proto.DeployResponse) error
	Undeploy(context.Context, *proto.UndeployRequest, *proto.UndeployResponse) error
	List(context.Context, *proto.ListRequest, *proto.ListResponse) error
}

type Task struct{}

func (t *Task) Deploy(ctx context.Context, req *proto.DeployRequest, res *proto.DeployResponse) error {
	log.Println(req)
	return nil
}

func (*Task) Undeploy(ctx context.Context, req *proto.UndeployRequest, res *proto.UndeployResponse) error {
	res.Status = "Job" + req.Name + "undeployed"
	return nil
}

func (*Task) List(ctx context.Context, req *proto.ListRequest, res *proto.ListResponse) error {
	return nil
}
