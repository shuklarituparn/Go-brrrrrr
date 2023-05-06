package main

import "fmt"

const (
	a = iota // 0
	b        // 1
	c        // 2

	//only if a is iota then the rest value will be increased sequentially

	//if we declared "a" as 0 and didn't declare others explicitly then they'll become 0 too
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
