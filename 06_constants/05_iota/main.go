package main

import "fmt"

const (
	a = iota // 0
	b        // 1
	c        // 2
)

const (
	d = iota // 0
	e        // 1
	f        // 2
)

// after each const declaration iota values resets so d,e,f, will again start from 0
func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
