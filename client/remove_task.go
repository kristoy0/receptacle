package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/urfave/cli"
)

func Remove(c *cli.Context) error {
	db, err := DbConn()
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

	if c.Args().First() != "" {
		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("tasks"))
			fmt.Println(c.Args().First())
			err := b.Delete([]byte(c.Args().First()))
			return err
		})
	}

	return nil
}
