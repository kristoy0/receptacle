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

func Services(c *cli.Context) error {
	req := APIRequest{
		"go.receptacle.server",
		"Tasks.List",
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

	data := store.Service{}

	err = json.Unmarshal(read, &data)
	if err != nil {
		log.Println(err)
	}

	out := []string{
		"HOST | NAME | IMAGE | PORT | IP",
	}

	for _, svc := range data.List {
		body := fmt.Sprintf("%s | %s | %s | %s | %s", svc.Host, svc.Name, svc.Image, svc.Port, svc.IP)
		out = append(out, body)
	}

	fmt.Println(columnize.SimpleFormat(out))
	return nil
}
