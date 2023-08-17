package main

import "fmt"

func hello() {
	fmt.Print("hello ")
}

func world() {
	fmt.Println("world")
}

func main() {
	world()   //here we don't have the defer keyword here so it be called first
	hello()
}
