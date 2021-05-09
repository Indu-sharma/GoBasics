package main 
import (
	"fmt"
)


func main() {
	myPointer := "Hi"
	s := sum(&myPointer, "Hello",10,20,30,40,50)
	fmt.Println("The Sum is", s) 

	s1 := sum1(&myPointer, "Hello",10,20,30,40,50)
	fmt.Println("The Sum when Pointer to Integer is Returned:", *s1) 
	// Can be defined only within another Function
	func(i int) {
		fmt.Println("Im anonymous function!  ", i)
	}(10)
}

// Demonstration of Call by value, call by Reference and variable Arguments. 
func sum(Ptr *string, msg string, values ...int) int {
	fmt.Println("Got the Message::", msg) 
	fmt.Println("Got the Pointer with value::", *Ptr) 
	res := 0
	for _,v := range values {
		res += v
	}
	return res
}

func sum1(Ptr *string, msg string, values ...int) *int {
	fmt.Println("Got the Message::", msg) 
	fmt.Println("Got the Pointer with value::", *Ptr) 
	res := 0
	for _,v := range values {
		res += v
	}
	return &res
}



