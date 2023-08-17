package main

import "fmt"

func main() {

	mySlice := make([]int, 0, 3) //an int slice with capacity of 3 and initialized with 0

	fmt.Println("-----------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("-----------------")

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len:", len(mySlice), "Capacity:", cap(mySlice), "Value: ", mySlice[i])
	}
}
