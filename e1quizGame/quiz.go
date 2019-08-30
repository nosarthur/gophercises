package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"

	arg "github.com/alexflint/go-arg"
)

func main() {
	var args struct {
		FPath   string `arg:"-f" help:"csv file path"`
		Timeout int    `help:"answer time window, default 10 seconds."`
	}
	args.FPath = "problems.csv" // default
	args.Timeout = 10           // seconds
	arg.MustParse(&args)

	// read csv

	f, err := os.Open(args.FPath)
	if err != nil {
		panic("cannot open")
	}
	r := csv.NewReader(f)
	numQ := 0
	correctA := 0
	var ans string
	scanner := bufio.NewScanner(os.Stdin)
	t := time.NewTicker(time.Second * time.Duration(args.Timeout))
	defer t.Stop()

	fmt.Println("Total time: %d seconds!", args.Timeout)
loop:
	for {
		select {
		case <-t.C:
			fmt.Println("Time is up!")
			break loop
		default:
		}
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("what is %s?\n", record[0])
		scanner.Scan()
		ans = scanner.Text()
		if ans == record[1] {
			correctA++
		}
		numQ++
	}
	fmt.Printf("%d / %d\n", correctA, numQ)
}
