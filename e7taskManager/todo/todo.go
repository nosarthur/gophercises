package todo

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

// FIXME: We should not know the bucket name
var bucketName = []byte("todos")

// Add a task
func Add(db *bolt.DB, task string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		t := time.Now()
		err := b.Put([]byte(t.String()), []byte(task))
		return err
	})
}

// List all tasks
func List(db *bolt.DB) error {

	return db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket(bucketName)

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
		}

		return nil
	})
}
