package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var orderIds = make([]int,0,4)
var lo = sync.Mutex{}

func main() {
	trigger := make(chan struct{})
	fetchOrder :=  make(chan []int)
	go orderIdProcessor(trigger,fetchOrder)
	i := 0
	for i < 7{
		fmt.Println(getOrderId(trigger))
		fmt.Println(orderIds)
		i++
		time.Sleep(1*time.Second)
	}

}
func orderIdProcessor(triggger chan struct{}, fetchOrder chan []int ) {
	for {
		select {
		case <-triggger:
			fetchOrderIdsFromDatabase(fetchOrder)
		case orderId := <- fetchOrder:
			//lo.Lock()
			orderIds = append(orderIds, orderId...)
			//lo.Unlock()
		default:
			fmt.Println("default")
			time.Sleep(1*time.Second)
		}
	}
}
func getOrderId(trigger chan struct{}) int {
	//go func() {
		if len(orderIds) < 3 {
			trigger <- struct {}{}
			fmt.Println("sent the trigger")
		}
	//}()
	for len(orderIds) <= 0 {
		time.Sleep(1*time.Second)
	}
	order := orderIds[0]
	//lo.Lock()
	orderIds = orderIds[1:]
	//lo.Unlock()
	return order
}


func fetchOrderIdsFromDatabase(fetch chan []int) {
	go func() {
		fetch <- []int{rand.Int(),rand.Int(),rand.Int(),rand.Int()}
	}()
}
