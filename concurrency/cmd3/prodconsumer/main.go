package main

import (
	"fmt"
	"sync"
	"time"
)
//type consumerinteface interface {
//	consume(j chan int, i int , timeSleep int)
//}
type producer struct {
	i int
}

func (a producer) produce(j chan  int) {
	for jj := 0;jj <a.i;jj++ {
		// in channel j put int value thil a.i = 10
		j <- jj
	}
	close(j)
}
type consumer struct {

}
func (a consumer) consume(j chan int, cons int, timeval int,wgGroup *sync.WaitGroup) {
	for val := range j {
		fmt.Println("cons ", cons ," the value " , val)
		time.Sleep(time.Duration(timeval) *time.Second)
	}
	wgGroup.Done()
}
var i = 10
func main() {
	//wgGroup  := sync.WaitGroup{}
	//product := make(chan int)
	//p1 := producer{6}
	////runtime.GOMAXPROCS(2)
	//go p1.produce(product)
	//var consumers = []int{1,2,3}
	//wgGroup.Add(len(consumers))
	//for _, consum := range consumers {
	//	//wgGroup.Add(1)
	//	go consumer{}.consume(product,consum, rand.Intn(1),&wgGroup)
	//}
	//
	//
	//
	////wgGroup.Add(1)
	////go c2.consume(product,2, rand.Intn(1),&wgGroup)
	//
	//wgGroup.Wait()
	//fmt.Println("this is after consume")
	for i:=0;i<10;i++ {
		updateI()
	}


}

func updateI() {

	i = i + 1;
}
//1. race condition produce might not have startd but consumer consuming --- channel will be listened only if some value is put
//2. 37 line will get printed once all the varibale conusmed