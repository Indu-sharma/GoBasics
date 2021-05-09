package main 

import (
	"fmt"
)

func main() {
	// Demonstration of Array 
	my_array := [3]int{1, 2, 3}
	my_array1 := [...]int{10,20,30}
	fmt.Printf("My Array : \n%v %T", my_array, my_array) // [1 2 3] [3]int 
	fmt.Printf("\nMy Array1 : \n%v %T", my_array1, my_array1) // [1 2 3] [3]int 
	// my_array2 := my_array // Make Totally different copies  of my_array
	my_array3 := &my_array // Make same  copy/alias  of my_array 
	fmt.Printf("\n%v", len(my_array3))

	// Demonstration of Slices :: Works like Python Lists. 
	my_slice := []int{1, 2, 3}
	my_slice1 := []int{10,20,30}
	fmt.Printf("\nMy my_slice : \n%v %T", my_slice, my_slice) // [1 2 3] [3]int 
	fmt.Printf("\nMy my_slice : \n%v %T", my_slice1, my_slice1) // [1 2 3] [3]int 
	my_slice3 := my_slice // Make same  copy/alias  of my_slice 
	fmt.Printf("\n%v", len(my_slice3))
	fmt.Printf("\n%v", cap(my_slice3)) // Gives Capacity 
	my_slice3 = append(my_slice3, 10)
	fmt.Printf("\n My Slice after Append - %v", my_slice3)
}