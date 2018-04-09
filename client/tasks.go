package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/urfave/cli"
)

func Tasks(c *cli.Context) error {
	db, err := DbConn()
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

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
}
