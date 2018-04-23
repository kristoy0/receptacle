package main

import (
	"log"
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

var (
	cmds    []cli.Command
	srcFile string
	dstIP   string
)

func init() {
	db, err := DbConn()
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

	cmds = []cli.Command{
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
			Action: Create,
		},
		{
			Name:   "remove",
			Usage:  "Remove a task",
			Action: Remove,
		},
		{
			Name:   "tasks",
			Usage:  "View tasks from local database",
			Action: Tasks,
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
			Action: Deploy,
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
			Action: Undeploy,
		},
		{
			Name:   "hosts",
			Usage:  "List all hosts",
			Action: Hosts,
		},
		{
			Name:   "services",
			Usage:  "List all services",
			Action: Services,
		},
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "rectl"
	app.Usage = "Receptacle command line client"

	app.Commands = cmds
	app.Run(os.Args)
}
