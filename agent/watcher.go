package main

import (
	"github.com/docker/libkv"
	"github.com/docker/libkv/store"
	"github.com/docker/libkv/store/consul"

	"encoding/json"
	"log"
	"time"
)

func WatchServiceDiscovery() error {
	consul.Register()
	client := "127.0.0.1:8500"

	kv, err := libkv.NewStore(
		store.CONSUL,
		[]string{client},
		&store.Config{
			ConnectionTimeout: 10 * time.Second,
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
					createHandler(pair.Key, pair.Value)
				}
			}
		}
	}
}

func createHandler(key string, value []byte) error {
	skey, err := stripDirectory(string(key))
	if err != nil {
		return err
	}
	log.Println(skey)

	cexists, err := ContainerExists(skey)
	if err != nil {
		return err
	}

	if !cexists {
		task := Task{}
		if err := json.Unmarshal(value, &task); err != nil {
			return err
		}
		log.Println(task)
		iexists, err := ImageExists(task.Image)
		if err != nil {
			return err
		}
		if !iexists {
			err = PullImage(task.Image)
			if err != nil {
				return err
			}
		}

		err = CreateContainer(task)
		if err != nil {
			return err
		}
	}

	return nil
}
