package main

import (
	"fmt"
	"time"
)

type producer struct {
	i int
}

func (a producer) produce(j chan  int) {
	for jj := 0;jj <5;jj++ {
		j <- jj
	}
	close(j)
}

var i = 0
func main() {
	//ch := make(chan int)
	//go func() {
	//	producer{}.produce(ch)
	//}()
	//selecClose(ch)
	ch := make(chan struct{},1)
	for j := 0;j < 1000;j++{
		go updateIByOne(ch)
	}
	time.Sleep(6*time.Second)
	fmt.Println(i)
}

func updateIByOne(ch chan struct{}){
	ch <- struct{}{}
	i = i + 1
	<-ch
}



//tabnine
func selecClose(ch chan int) {
	raj:
	for {
		select {
		   case producedItems,ok := <- ch:
		   	fmt.Println(ok)
		   	if !ok{
		   		break raj
			}
		   	fmt.Println(producedItems)
		   	time.Sleep(1*time.Second)
		default:
			fmt.Println("nothing produced")
			time.Sleep(1*time.Second)
		}
	}
	fmt.Println("out of the loop")
}

