package main

import (
	/*
	micro "github.com/micro/go-micro"
	*/
	"log"
	/*"time"*/
)

func main() {
/*	service := micro.NewService(
		micro.Name("go.receptacle.agent"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)

	service.Init()

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}*/

	err := WatchServiceDiscovery()
	if err != nil {
		log.Fatal(err)
	}
}
