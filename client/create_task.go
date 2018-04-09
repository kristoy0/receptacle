package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/boltdb/bolt"
	"github.com/kristoy0/receptacle/store"
	"github.com/urfave/cli"
)

func Create(c *cli.Context) error {
	db, err := DbConn()
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

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
}
