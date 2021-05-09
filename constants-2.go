// Constants variables names are either PascalCase or camelCase. 
package main

import (
	"fmt"
	"math"
)

const (
	a = iota 
	b = iota
	c   // After first Constat variable rest need not to be initialized to iota. 
)
func main() {
	fmt.Println("Experiment of Contsants!")
	const infer_int = 10 // This will be inferred in case of Constants. 
	const myConst int32 = 42 // This can't be changed. This
	const myConst1 int = math.MaxInt64  // This is Ok as no methods are being called/used.  
	// const myConst2 int = math.Sin(30) // Const initializer math.Sin(30) is not a constant. This is invalid & program fails. 
	fmt.Printf("%v %T\n", myConst, myConst)
	fmt.Printf("%v %T\n", myConst1, myConst1)
	// fmt.Printf("%v %T", myConst2, myConst2)
	fmt.Printf("%v %T\n", myConst1, myConst1)
	fmt.Printf("%v %T\n", infer_int, infer_int)
	// Enumerated Constants.  
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
}
