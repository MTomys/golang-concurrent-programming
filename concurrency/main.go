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

	go func(validOrderCh <-chan order) {
		select {
		case order := <-validOrderCh:
			fmt.Printf("Valid order received: %v\n", order)
		case order := <-invalidOrderCh:
			fmt.Printf("Invalid order received: %v with error %v\n", order.order, order.err)
		}
		wg.Done()
	}(validOrderCh)

	go func(invalidOrderCh <-chan invalidOrder) {
		wg.Done()
	}(invalidOrderCh)

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
}

func validateOrder(in <-chan order, out chan<- order, errCh chan<- invalidOrder) {
	order := <-in
	if order.Quantity <= 0 {
		errCh <- invalidOrder{order: order, err: errors.New("Quantity can not be less or equal to 0")}
	} else {
		out <- order
	}
}

var rawOrders = []string{
	`{"productCode": 1111, "quantity": -1, "status": 1}`,
	`{"productCode": 1234, "quantity": 2, "status": 1}`,
	`{"productCode": 4321, "quantity": 9, "status": 1}`,
	`{"productCode": 4321, "quantity": -2, "status": 1}`,
}
