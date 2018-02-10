package main

import (
	"github.com/docker/libkv"
	"github.com/docker/libkv/store"
	"github.com/docker/libkv/store/consul"

	"log"
	"time"
	"encoding/json"
)

func WatchServiceDiscovery() error {
	consul.Register()
	client := "127.0.0.1:8500"

	kv, err := libkv.NewStore(
		store.CONSUL,
		[]string{client},
		&store.Config{
			ConnectionTimeout: 10*time.Second,
		},
	)

	stopCh := make(<-chan struct{})

	events, err := kv.WatchTree("receptacle", stopCh)
	if err != nil {
		return err
	}
	for {
		select {
		case pairs := <-events:
			for _, pair := range pairs {
				if pair.Key != "" && string(pair.Value) != "" {
					key, err := stripDirectory(string(pair.Key))
					if err != nil {
						return err
					}
					log.Println(key)

					exists, err := ContainerExists(key)
					if err != nil {
						return err
					}
					if !exists {
						task := Task{}
						if err := json.Unmarshal(pair.Value, &task); err != nil {
							return err
						}
						log.Println(task)

						// Need to pull image first

						err := CreateContainer(task)
						if err != nil {
							return err
						}
					}
				}
			}
		}
	}

	return nil
}
