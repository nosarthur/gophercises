package main

import (
	"fmt"
	"strings"

	arg "github.com/alexflint/go-arg"
	"github.com/nosarthur/gophercises/e7taskManager/db"
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

// RmCmd removes a todo
type RmCmd struct {
	Number int `arg:"positional, required"`
}

var args struct {
	Add  *AddCmd  `arg:"subcommand:add"`
	List *ListCmd `arg:"subcommand:list"`
	Do   *DoCmd   `arg:"subcommand:do"`
	Rm   *RmCmd   `arg:"subcommand:rm"`
}

var dbPath = "my.db"

func main() {
	db.MustInit(dbPath)
	arg.MustParse(&args)

	switch {
	case args.Add != nil:
		if err := db.Add(strings.Join(args.Add.Task, " ")); err != nil {
			fmt.Println("%s", err)
			panic("Fail to add task!")
		}
	case args.List != nil:
		if err := db.List(); err != nil {
			fmt.Println("%s", err)
			panic("cannot list")
		}
	case args.Do != nil:
		fmt.Printf("do %d\n", args.Do.Number)
	case args.Rm != nil:
		fmt.Println("rm")
	}

}
