package main

import (
	"log"
	"time"

	micro "github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("go.receptacle.agent"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)

	service.Init()

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

	err := WatchServiceDiscovery()
	if err != nil {
		log.Fatal(err)
	}
}
