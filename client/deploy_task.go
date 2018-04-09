package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/boltdb/bolt"
	"github.com/kristoy0/receptacle/store"
	"github.com/urfave/cli"
)

func Deploy(c *cli.Context) error {
	db, err := DbConn()
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

	if c.Args().First() == "" {
		fmt.Println("Task name not specified")
	} else if dstIP == "" {
		fmt.Println("Master address not specified")
	} else {
		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("tasks"))
			v := b.Get([]byte(c.Args().First()))

			data := store.Task{}

			if err := json.Unmarshal(v, &data); err != nil {
				panic(err)
			}

			req := APIRequest{
				"go.receptacle.server",
				"Tasks.Deploy",
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
				fmt.Println("Deployed " + data.Name + " successfully.")
			}

			return nil
		})
	}

	return nil
}
