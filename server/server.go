package main

import (
	"log"
	"time"

	proto "github.com/kristoy0/receptacle/server/proto"
	micro "github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("go.receptacle.server"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)

	service.Init()

	proto.RegisterTasksHandler(service.Server(), new(Task))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
