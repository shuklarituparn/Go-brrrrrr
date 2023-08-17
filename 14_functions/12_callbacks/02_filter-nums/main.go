package main

import "fmt"

func filter(numbers []int, callback func(int) bool) []int {
	var xs []int          //the empty slice to store the filtered values
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)  //appending the filtered values to the slice
		}
	}
	return xs
}

func main() {
	xs := filter([]int{1, 2, 3, 4}, func(n int) bool { 
		return n > 1
	})                 //passsing an int slice and a function that returns bool
	fmt.Println(xs) // [2 3 4]
}
