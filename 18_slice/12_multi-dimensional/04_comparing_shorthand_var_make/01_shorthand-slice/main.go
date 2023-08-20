package main

import (
	"fmt"
)

func main() {
	student := []string{} 
	students := [][]string{}
	//student[0] = "Todd"  //thsis doesn't work and gives out of bounds error
	student = append(student, "Todd")   //this works
	fmt.Println(student)     
	fmt.Println(students)  //students is an empty slice containing a slice
}
