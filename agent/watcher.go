package main

import (
	"context"
	"os"
	"strings"

	docker "docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"github.com/docker/libkv"
	kvstore "github.com/docker/libkv/store"
	"github.com/docker/libkv/store/consul"
	"github.com/kristoy0/receptacle/store"

	"encoding/json"
	"log"
	"time"
)

var (
	kv kvstore.Store
)

func init() {
	// register a consul instance

	consul.Register()
	client := os.Getenv("CONSUL_ADDR")

	var err error

	kv, err = libkv.NewStore(
		kvstore.CONSUL,
		[]string{client},
		&kvstore.Config{
			ConnectionTimeout: 10 * time.Second,
		},
	)

	if err != nil {
		log.Fatalln(err)
	}
}

// WatchServiceDiscovery -
// This function is used to watch the service
// discovery kv store
func WatchServiceDiscovery() error {

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
		default:
			err := deleteHandler()
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func deleteHandler() error {
	cli, err := docker.NewEnvClient()
	if err != nil {
		return err
	}

	ctx := context.Background()

	conts, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return err
	}

	for _, cont := range conts {
		if !strings.Contains(cont.Names[0][1:], "receptacle") {
			exists, err := kv.Exists("receptacle/" + cont.Names[0][1:])

			if err != nil {
				return err
			}

			if !exists {
				_ = DeleteContainer(cont.ID)
			}
		}
	}

	return nil
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
