package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kristoy0/receptacle/store"
	"github.com/urfave/cli"
)

func Undeploy(c *cli.Context) error {
	if c.Args().First() == "" {
		fmt.Println("Task name not specified")
	} else if dstIP == "" {
		fmt.Println("Master address not specified")
	} else {
		data := store.Task{
			Name: c.Args().First(),
		}

		req := APIRequest{
			"go.receptacle.server",
			"Tasks.Undeploy",
			data,
		}

		byt, err := json.Marshal(req)
		if err != nil {
			return err
		}

		httpReq, err := http.NewRequest("POST", "http://"+dstIP+":8080/rpc", bytes.NewBuffer(byt))
		httpReq.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		_, err = client.Do(httpReq)
		if err != nil {
			fmt.Println(err)
		}

		if err == nil {
			fmt.Println("Undeployed " + c.Args().First() + " successfully.")
		}
	}
	return nil
}
