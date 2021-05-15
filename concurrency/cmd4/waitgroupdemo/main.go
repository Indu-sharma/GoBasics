package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type producer struct {
	i int
}

func (a producer) produce(j chan  int) {
	for jj := 0;jj <a.i;jj++ {
		j <- jj
	}
	close(j)
}
type consumer struct {

}
var wgG = sync.WaitGroup{}
func (a consumer) consume(j chan  int, cons int, timeval int) {
	for val := range j {
		fmt.Println("cons ", cons ," the value " , val)
		time.Sleep(time.Duration(timeval) *time.Second)
	}
	wgG.Done()
}
func main() {
	product := make(chan int)
	p1 := producer{10}
	c1 := consumer{}
	wgG.Add(1)
	//c2 := consumer{}
	go p1.produce(product)
	go c1.consume(product,1, rand.Intn(6))
	//go c2.consume(product,2, rand.Intn(3))
	wgG.Wait()
}
