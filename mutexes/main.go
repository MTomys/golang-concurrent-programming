package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var o sync.Once

func main() {
	ctext()
}

func lock() {
	s := []int{}
	wg := sync.WaitGroup{}
	const iterations = 1000
	wg.Add(iterations)
	for i := 0; i < iterations; i++ {
		go func() {
			s = append(s, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(len(s))
}

func Run() {
	o.Do(func() {
		log.Println("opening connection to database")
		var err error
		db, err = sql.Open("sqlite3", "./mydb.db")
		if err != nil {
			log.Fatal(err)
		}
	})
}

func ctext() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(1)

	go func(ctx context.Context) {
		defer wg.Done()
		for range time.Tick(500 * time.Millisecond) {
			if ctx.Err() != nil {
				log.Println(ctx.Err())
				return
			}
			fmt.Println("Tick!")
		}
	}(ctx)

	wg.Wait()
}
