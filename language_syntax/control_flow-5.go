package main 
import (
	"fmt"
)

func main() {

	// If Conditional Statements. 
	number := 50
	guess := 50 
	threshold := 30 

	if guess < threshold {
		fmt.Println("Below Threshold") 
	}else if guess > number {
		fmt.Println("Guess Value is Greater than the number")
	}else if guess < number {
		fmt.Println("Guess Value is lesser than the number")
	} else {
		fmt.Println("Guess Value is same as Number")
	}

	// Switch with Tag/Label. 
	value := 9

	switch value1 := value + 1; value1 { // Dynamic Calculation ; syntax is similar in if statement. 
		case 0: 
			fmt.Println("\nI recieved 0")
		case 1: 
			fmt.Println("I recieved 1")
		case 2,3,4,5:
			fmt.Println("I recieved a value between 2 and 5")
		default:
			fmt.Println("I recieved other than >5")
	}

	// Tag Less Switch  , you can use variable & comparison operators in each case statement.
	switch {
		case value < 10:
			fmt.Println("\n I recieved < 10")
		case value > 10: 
			fmt.Println("\n I recieved > 10")
		default:  
			fmt.Println("\n I recieved  10")
	}

	// Type Switch  

	var Val interface{} = []int {1,2,3,4}
	switch Val.(type) {
		case int:
			fmt.Println("\nI Recieved INT")
		case float64:
			fmt.Println("I Recieved Float") 
		case string:
			fmt.Println("I Recieved String")
		case []int:
			fmt.Println("I Recieved a Slice")
		default:
			fmt.Println("I Recieved Other than Int, Float, String")
	}

}