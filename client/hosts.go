package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kristoy0/receptacle/store"
	"github.com/ryanuber/columnize"
	"github.com/urfave/cli"
)

func Hosts(c *cli.Context) error {
	req := APIRequest{
		"go.receptacle.server",
		"Tasks.Hosts",
		store.Task{},
	}

	byt, err := json.Marshal(req)
	if err != nil {
		return err
	}

	httpReq, err := http.NewRequest("POST", "http://"+dstIP+":8080/rpc", bytes.NewBuffer(byt))
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		log.Println(err)
	}

	read, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	data := store.Nodes{}

	err = json.Unmarshal(read, &data)
	if err != nil {
		log.Println(err)
	}

	out := []string{
		"ID | NODE | ADDRESS",
	}

	for _, node := range data.Hosts {
		body := fmt.Sprintf("%s | %s | %s", node.ID[:8], node.Node, node.Address)
		out = append(out, body)
	}

	fmt.Println(columnize.SimpleFormat(out))
	return nil
}
