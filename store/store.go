package store

import (
	"github.com/hashicorp/consul/api"
)

func GetKV() (*api.KV, error) {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	kv := client.KV()
	return kv, err
}