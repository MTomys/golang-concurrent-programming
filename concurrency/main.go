package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	receivedOrdersCh := make(chan order)
	validOrderCh := make(chan order)
	invalidOrderCh := make(chan invalidOrder)

	go receiveOrders(receivedOrdersCh)
	go validateOrder(receivedOrdersCh, validOrderCh, invalidOrderCh)

	wg.Add(1)

	go func(validOrderCh <-chan order, invalidOrderCh <-chan invalidOrder) {
	loop:
		for {
			select {
			case order, ok := <-validOrderCh:
				if ok {
					fmt.Printf("Valid order received: %v\n", order)
				} else {
					break loop
				}
			case invalidOrder, ok := <-invalidOrderCh:
				if ok {
					fmt.Printf("Invalid order received: %v with error %v\n", invalidOrder.order, invalidOrder.err)
				} else {
					break loop
				}
			}
		}
		wg.Done()
	}(validOrderCh, invalidOrderCh)

	wg.Wait()
}

func receiveOrders(out chan<- order) {
	for _, rawOrder := range rawOrders {
		var newOrder order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			log.Print(err)
			continue
		}
		out <- newOrder
	}
	close(out)
}

func validateOrder(in <-chan order, out chan<- order, errCh chan<- invalidOrder) {
	for order := range in {
		if order.Quantity <= 0 {
			errCh <- invalidOrder{order: order, err: errors.New("Quantity can not be less or equal to 0")}
		} else {
			out <- order
		}
	}
	close(out)
	close(errCh)
}

var rawOrders = []string{
	`{"productCode": 1111, "quantity": 1, "status": 1}`,
	`{"productCode": 1234, "quantity": 2, "status": 1}`,
	`{"productCode": 4321, "quantity": 9, "status": 1}`,
	`{"productCode": 4321, "quantity": -2, "status": 1}`,
}
