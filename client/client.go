package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/boltdb/bolt"
	"github.com/kristoy0/receptacle/store"
	"github.com/urfave/cli"
)

type APIRequest struct {
	Service string
	Method  string
	Request store.Task
}

func main() {
	db, err := bolt.Open("rectl.db", 0600, nil)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("tasks"))
		if err != nil {
			log.Fatalf("failed to create bucket: %v", err)
		}
		return nil
	})

	srcFile := ""
	dstIP := ""

	app := cli.NewApp()
	app.Name = "rectl"
	app.Usage = "Receptacle command line client"

	app.Commands = []cli.Command{
		{
			Name:  "create",
			Usage: "Create a task",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "file, f",
					Usage:       "Pass JSON file to add command",
					Destination: &srcFile,
				},
			},
			Action: func(c *cli.Context) error {
				data, err := ioutil.ReadFile(srcFile)
				if err != nil {
					return err
				}
				task := store.Task{}
				json.Unmarshal(data, &task)

				db.Update(func(tx *bolt.Tx) error {
					b := tx.Bucket([]byte("tasks"))
					err := b.Put([]byte(task.Name), data)
					return err
				})

				return nil
			},
		},
		{
			Name:  "remove",
			Usage: "Remove a task",
			Action: func(c *cli.Context) error {
				if c.Args().First() != "" {
					db.Update(func(tx *bolt.Tx) error {
						b := tx.Bucket([]byte("tasks"))
						fmt.Println(c.Args().First())
						err := b.Delete([]byte(c.Args().First()))
						return err
					})
				}

				return nil
			},
		},
		{
			Name:  "tasks",
			Usage: "View tasks from local database",
			Action: func(c *cli.Context) error {
				if c.Args().First() == "" {
					db.View(func(tx *bolt.Tx) error {
						b := tx.Bucket([]byte("tasks"))

						c := b.Cursor()

						for k, _ := c.First(); k != nil; k, _ = c.Next() {
							fmt.Printf("%s", k)
						}

						return nil
					})
				} else {
					db.View(func(tx *bolt.Tx) error {
						b := tx.Bucket([]byte("tasks"))
						v := b.Get([]byte(c.Args().First()))

						fmt.Printf("%s", v)
						return nil
					})
				}
				return nil
			},
		},
		{
			Name:  "deploy",
			Usage: "Deploy a task from the local database",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "master, m",
					Usage:       "Pass master address",
					Destination: &dstIP,
				},
			},
			Action: func(c *cli.Context) error {
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
			},
		},
		{
			Name:  "undeploy",
			Usage: "Undeploy a task from the cluster",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "master, m",
					Usage:       "Pass master address",
					Destination: &dstIP,
				},
			},
			Action: func(c *cli.Context) error {
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
			},
		},
	}

	app.Run(os.Args)
}
