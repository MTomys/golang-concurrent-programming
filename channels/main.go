package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg = sync.WaitGroup{}
	ch := make(chan string)
	wg.Add(1)
	go send(ch)
	go receive(ch, &wg)
	wg.Wait()
}

func send(ch chan<- string) {
	ch <- "Message test"
}

func receive(ch <-chan string, wg *sync.WaitGroup) {
	msg := <-ch
	fmt.Println(msg)
	wg.Done()
}
