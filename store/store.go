package store

import (
	"os"

	"github.com/hashicorp/consul/api"
)

func GetKV() (*api.KV, error) {
	client, err := api.NewClient(&api.Config{Address: os.Getenv("CONSUL_ADDR"), Scheme: "http"})
	if err != nil {
		return &api.KV{}, err
	}

	kv := client.KV()
	return kv, nil
}

func GetCatalog() (*api.Catalog, error) {
	client, err := api.NewClient(&api.Config{Address: os.Getenv("CONSUL_ADDR"), Scheme: "http"})
	if err != nil {
		return &api.Catalog{}, err
	}

	catalog := client.Catalog()

	return catalog, nil
}

func GetAgent() (*api.Agent, error) {
	client, err := api.NewClient(&api.Config{Address: os.Getenv("CONSUL_ADDR"), Scheme: "http"})
	if err != nil {
		return &api.Agent{}, err
	}

	agent := client.Agent()

	return agent, nil
}
