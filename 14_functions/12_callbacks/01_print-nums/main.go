package main

import "fmt"

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers { //first index is blank as we just want to know num
		callback(n) //call the callback function
	}
}

func main() {
	visit([]int{1, 2, 3, 4}, func(n int) {  
		fmt.Println(n)    //passing int slice and  a function as a parameter
	})
}

// callback: passing a func as an argument
