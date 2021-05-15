package main 
import (
	"fmt"
	"strconv"
)
// Variables Naming Convention in Go are either PascalCase or camelCase. 
// Dynamic Declaration are not possible at the package level. 
// Instead of Multiple var declaration ; we can declare in Bulk within the Block. 
var ( 
	my_int int = 123 
 	my_string string = "Hello world!"
)
// Next line of code throws error as you can't shadow same variable within the same block/scope. 
// var my_int int = 45 
func main() {
	// How ever we can override same variable declaration as that of package scope here as -  var my_int int = 45  
	my_char_auto := 'A' // Single Quote is treated as Init or series of Bytes. 
	my_string_auto := "Check!"
	my_int_auto := 10 
	fmt.Printf("%v %T\n", my_int, my_int)
	fmt.Printf("%v %T\n", my_string, my_string)
	fmt.Printf("%v %T\n", my_char_auto, my_char_auto)
	fmt.Printf("%v %T\n", my_string_auto, my_string_auto)
	fmt.Printf("%v %T\n\n", my_int_auto, my_int_auto)
	//my_int_auto_ := string(my_int_auto) // Convert  from Integer  to String this way results in a string of runes. 
	//fmt.Printf("%v %T\n", my_int_auto_, my_int_auto_) 
	my_int_str := strconv.Itoa(my_int_auto) // This is the way to convert from String to Integer. 
	fmt.Printf("%v %T\n", my_int_str, my_int_str)
}
