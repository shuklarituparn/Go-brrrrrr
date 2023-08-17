package main

import "fmt"

func main() {
	var x [58]int  //59 values int array
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
	x[42] = 777
	fmt.Println(x[42])
}
