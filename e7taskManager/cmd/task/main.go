package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	arg "github.com/alexflint/go-arg"
	"github.com/boltdb/bolt"
	"github.com/nosarthur/gophercises/e7taskManager/todo"
)

// init the db and bucket
func init() {

}

// AddCmd is a subcommand
type AddCmd struct {
	Task []string `arg:"positional, required"`
}

// ListCmd lists all todos
type ListCmd struct {
}

// DoCmd finishes a todo
type DoCmd struct {
	Number int `arg:"positional, required"`
}

var args struct {
	Add  *AddCmd  `arg:"subcommand:add"`
	List *ListCmd `arg:"subcommand:list"`
	Do   *DoCmd   `arg:"subcommand:do"`
}

var dbLoc = "my.db"

func main() {
	db, err := bolt.Open(dbLoc, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("todos"))
		if err != nil {
			return fmt.Errorf("Cannot create bucket: %s", err)
		}
		return nil
	})

	arg.MustParse(&args)

	switch {
	case args.Add != nil:
		if err := todo.Add(db, strings.Join(args.Add.Task, " ")); err != nil {
			panic("Fail to add task!")
		}
	case args.List != nil:
		fmt.Println("list")
	case args.Do != nil:
		fmt.Printf("do %d\n", args.Do.Number)
	}
}
