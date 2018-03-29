package main

import (
	"github.com/docker/libkv"
	kvstore "github.com/docker/libkv/store"
	"github.com/docker/libkv/store/consul"
	"github.com/kristoy0/receptacle/store"

	"encoding/json"
	"log"
	"time"
)

// WatchServiceDiscovery -
// This function is used to watch the service
// discovery kv store
func WatchServiceDiscovery() error {
	// register a consul instance
	consul.Register()
	client := "127.0.0.1:8500"

	kv, err := libkv.NewStore(
		kvstore.CONSUL,
		[]string{client},
		&kvstore.Config{
			ConnectionTimeout: 10 * time.Second,
		},
	)

	stopCh := make(<-chan struct{})

	// fetch all events from the receptacle tree
	// on the kv store
	events, err := kv.WatchTree("receptacle", stopCh)
	if err != nil {
		return err
	}

	// loop over the events endlessly until
	// the program is closed
	for {
		select {
		case pairs := <-events:
			for _, pair := range pairs {
				// check if the key and value arent empty
				if pair.Key != "" && string(pair.Value) != "" {
					// handler for creating containers
					createHandler(pair.Key, pair.Value)
				}
			}
		}
	}
}

func createHandler(key string, value []byte) error {
	// strip directory from given key
	skey, err := stripDirectory(string(key))
	if err != nil {
		return err
	}
	log.Println(skey)

	// check if container exists already
	cexists, err := ContainerExists(skey)
	if err != nil {
		return err
	}

	// if container doesnt exist already
	if !cexists {
		// create a new Task instance
		task := store.Task{}
		if err := json.Unmarshal(value, &task); err != nil {
			return err
		}
		log.Println(task)

		// check if image exists on the host already
		iexists, err := ImageExists(task.Image)
		if err != nil {
			return err
		}
		// if image doesnt already exist, pull it
		if !iexists {
			err = PullImage(task.Image)
			if err != nil {
				return err
			}
		}

		// create the container with given specification
		err = CreateContainer(task)
		if err != nil {
			return err
		}
	}

	return nil
}
