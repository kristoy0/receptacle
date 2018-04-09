package main

import (
	"github.com/boltdb/bolt"
)

func DbConn() (*bolt.DB, error) {
	db, err := bolt.Open("rectl.db", 0600, nil)
	if err != nil {
		return nil, err
	}
	return db, nil
}
