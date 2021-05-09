package main
import (
	"fmt"
)
func main() {
	for i := 0; i < 5 ; i++ {
		fmt.Println(i)
	}

	mySlice := []int {10,20,40,50}
	for k, v := range mySlice {
		fmt.Println(k, v)
	}

	Students := map[string]int {
		"Indu": 100,
		"Sharma": 200,
	}
	// Loop through Map. 

	for k, v := range Students {
		println("My value", v)
		println("My Key", k)
	}
	// Note Replace k or v with _ in case you dont want either of key, value. 
}


