package main

import "fmt"

func main() {
	var x [58]string

	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)      //converting the num values to string
	}

	fmt.Println(x)
	fmt.Println(x[42])
}
