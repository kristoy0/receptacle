package main

import (
	"log"
	"time"

	micro "github.com/micro/go-micro"
)

func main() {
	// Initialize the microservice
	// This will register the agent to service discovery
	service := micro.NewService(
		micro.Name("go.receptacle.agent"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)

	service.Init()

	// Start watching for changes in the key/value store
	// of the service discovery
	go WatchServiceDiscovery()

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
