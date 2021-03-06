package main

import (
	"context"
	"log"
	"strconv"
	"strings"

	"encoding/json"
	"errors"

	"github.com/hashicorp/consul/api"
	proto "github.com/kristoy0/receptacle/server/proto"
	"github.com/kristoy0/receptacle/store"
)

type TaskHandler interface {
	Deploy(context.Context, *proto.DeployRequest, *proto.DeployResponse) error
	Undeploy(context.Context, *proto.UndeployRequest, *proto.UndeployResponse) error
	List(context.Context, *proto.ListRequest, *proto.ListResponse) error
	Hosts(context.Context, *proto.HostsRequest, *proto.HostsResponse) error
}

type Task struct{}

func (t *Task) Deploy(ctx context.Context, req *proto.DeployRequest, res *proto.DeployResponse) error {
	if req.Name == "" || req.Image == "" {
		return errors.New("Name and/or image missing")
	}

	mreq, err := json.Marshal(req)
	if err != nil {
		return err
	}

	kv, err := store.GetKV()
	if err != nil {
		return err
	}

	p := &api.KVPair{Key: "receptacle/" + req.Name, Value: []byte(mreq)}
	_, err = kv.Put(p, nil)
	if err != nil {
		return err
	}

	log.Println(req)

	return nil
}

func (*Task) Undeploy(ctx context.Context, req *proto.UndeployRequest, res *proto.UndeployResponse) error {
	res.Status = "Job " + req.Name + " undeployed"

	kv, err := store.GetKV()
	if err != nil {
		return err
	}

	_, err = kv.Delete("receptacle/"+req.Name, nil)
	if err != nil {
		return err
	}

	return nil
}

func (*Task) List(ctx context.Context, req *proto.ListRequest, res *proto.ListResponse) error {
	agent, err := store.GetAgent()
	if err != nil {
		return err
	}

	if req.Id == "" {
		svcs, err := agent.Services()
		if err != nil {
			return err
		}

		for _, svc := range svcs {
			id := strings.Split(svc.ID, ":")

			if len(id) > 2 && !strings.HasPrefix(id[1], "receptacle") {
				res.List = append(res.List, &proto.List{
					Host:  id[0],
					Name:  id[1],
					Image: svc.Service,
					Port:  strconv.Itoa(svc.Port),
					Ip:    svc.Address,
				})
			}

		}
	}

	return nil
}

func (*Task) Hosts(ctx context.Context, req *proto.HostsRequest, res *proto.HostsResponse) error {
	catalog, err := store.GetCatalog()
	if err != nil {
		return err
	}

	nodes, _, err := catalog.Nodes(&api.QueryOptions{})
	if err != nil {
		return err
	}

	for _, node := range nodes {
		res.Hosts = append(res.Hosts, &proto.Host{
			Id:      node.ID,
			Node:    node.Node,
			Address: node.Address,
		})
	}

	return nil
}
