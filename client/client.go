package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/boltdb/bolt"
	"github.com/kristoy0/receptacle/store"
	"github.com/urfave/cli"
)

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
	}

	app.Run(os.Args)
}
