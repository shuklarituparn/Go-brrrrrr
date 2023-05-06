package main

import "fmt"

const (
	d = iota
	a = iota // 0
	b = iota // 1
	c = iota // 2
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d) //iota is the number of const specification in the const declaration

	//d is declared first so it has serial number of 0, every const declared after will be +1
}
