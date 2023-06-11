package main

//
// func main() {
// 	var wg sync.WaitGroup
//
// 	receivedOrdersCh := receiveOrders()
// 	validOrderCh, invalidOrderCh := validateOrder(receivedOrdersCh)
// 	reservedInventoryCh := reserveInventory(validOrderCh)
//
// 	wg.Add(2)
// 	go func(invalidOrderCh <-chan invalidOrder) {
// 		for invalidOrder := range invalidOrderCh {
// 			fmt.Printf("Invalid order received: %v with error %v\n", invalidOrder.order, invalidOrder.err)
// 		}
// 		wg.Done()
// 	}(invalidOrderCh)
//
// 	go func(reservedInventoryCh <-chan order) {
// 		for order := range reservedInventoryCh {
// 			fmt.Printf("Inventory reserved for: %v\n", order)
// 		}
// 		wg.Done()
// 	}(reservedInventoryCh)
//
// 	wg.Wait()
// }
//
// func receiveOrders() <-chan order {
// 	out := make(chan order)
// 	go func() {
// 		for _, rawOrder := range rawOrders {
// 			var newOrder order
// 			err := json.Unmarshal([]byte(rawOrder), &newOrder)
// 			if err != nil {
// 				log.Print(err)
// 				continue
// 			}
// 			out <- newOrder
// 		}
// 		close(out)
// 	}()
//
// 	return out
// }
//
// var rawOrders = []string{
// 	`{"productCode": 1111, "quantity": 1, "status": 1}`,
// 	`{"productCode": 1234, "quantity": 2, "status": 1}`,
// 	`{"productCode": 4321, "quantity": 9, "status": 1}`,
// 	`{"productCode": 4321, "quantity": -2, "status": 1}`,
// }
//
// func validateOrder(in <-chan order) (<-chan order, <-chan invalidOrder) {
// 	out := make(chan order)
// 	errCh := make(chan invalidOrder, 1)
// 	go func() {
// 		for order := range in {
// 			if order.Quantity <= 0 {
// 				errCh <- invalidOrder{order: order, err: errors.New("Quantity can not be less or equal to 0")}
// 			} else {
// 				out <- order
// 			}
// 		}
// 		close(out)
// 		close(errCh)
// 	}()
// 	return out, errCh
// }
//
// func reserveInventory(in <-chan order) <-chan order {
// 	out := make(chan order)
//
// 	go func() {
// 		for o := range in {
// 			o.Status = reserved
// 			out <- o
// 		}
// 		close(out)
// 	}()
//
// 	return out
// }
//
// func printOrders(validOrderCh <-chan order, invalidOrderCh <-chan invalidOrder, wg *sync.WaitGroup) {
// 	wg.Add(1)
// 	go func(validOrderCh <-chan order, invalidOrderCh <-chan invalidOrder) {
// 	loop:
// 		for {
// 			select {
// 			case order, ok := <-validOrderCh:
// 				if ok {
// 					fmt.Printf("Valid order received: %v\n", order)
// 				} else {
// 					break loop
// 				}
// 			case invalidOrder, ok := <-invalidOrderCh:
// 				if ok {
// 					fmt.Printf("Invalid order received: %v with error %v\n", invalidOrder.order, invalidOrder.err)
// 				} else {
// 					break loop
// 				}
// 			}
// 		}
// 		wg.Done()
// 	}(validOrderCh, invalidOrderCh)
// }
