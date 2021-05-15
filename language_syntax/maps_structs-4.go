package main 

import (
	"fmt"
	"reflect"
)

// Struct  :: Create new Compositive Data Type in Go. 
type Student struct {
	name string 
	rollnumber int `required max:"300"`
	marks []int  // Slice 
}


func main() {
	// Key is String and Values can be any primitive Types and the slices 
	// Array can be Key but not Slice. 
	// Map is like Dictionary in Python   
	students := map[string]int{} 
	students["Indu"] = 100 
	students["Bibash"] = 105
	students["Renuka"] = 110
	fmt.Print(students)
	// Another way 
	students1 := make(map[string]int)
	students1 = map[string]int{
		"Indu": 100,
		"Bibash": 105, 
		"Renuka": 110,
	} 
	fmt.Print("\n", students1)
	// Accessing Values  of Maps.  
	key := "Indu"
	fmt.Print("\n key & value = ",key, ":", students1[key])

	// Struct 

	MyStudent := Student{
		"Indu",
		100, 
		[]int {
			10,
			20,
			30,
		},
	}
	fmt.Print("\n Student Details:", MyStudent)
	fmt.Print("\n Marks:", MyStudent.marks)
	MyStudent.marks = []int{100,200,300,}
	fmt.Print("\n Marks After Modification:", MyStudent.marks)

	MyStudentCopy := MyStudent // This Make totally different Copy. To Create Alias , use & 
	fmt.Print("\n Student Details Copy:", MyStudentCopy)

	// Accessing Tags from within the Struct Fields any using Reflect 

	MyType := reflect.TypeOf(Student{})
	field, _ := MyType.FieldByName("rollnumber")
	println("\n", field.Tag) 


}