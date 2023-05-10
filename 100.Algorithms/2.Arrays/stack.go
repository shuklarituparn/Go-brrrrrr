package main

import "fmt"

type Stack struct {
	size int
	top  int
	data []int
}

func create(stak *Stack) {
	fmt.Print("Enter Size: ")
	fmt.Scan(&stak.size)
	stak.top = -1                      //declaring the top as -1
	stak.data = make([]int, stak.size) //creating an int array of size given
}

func Display(stak Stack) {
	for i := stak.top; i >= 0; i-- { //looping through the stack
		fmt.Printf("%d ", stak.data[i]) //printing the value of stack
	}
	fmt.Println()
}

func push(stak *Stack, element int) {
	if stak.top == stak.size-1 {
		fmt.Println("Stack overflow")
	} else {
		stak.top++
		stak.data[stak.top] = element
	}
}

func pop(stak *Stack) int {
	x := -1
	if stak.top == -1 {
		fmt.Println("Stack underflow")
	} else {
		x = stak.data[stak.top]
		stak.top--
	}
	return x
}

func peek(stak Stack, index int) int {
	x := -1
	if stak.top-index+1 < 0 {
		fmt.Println("Invalid index")
	} else {
		x = stak.data[stak.top-index+1]
	}
	return x
}

func isEmpty(stak Stack) bool {
	return stak.top == -1
}

func isFull(stak Stack) bool {
	return stak.top == stak.size-1
}

func stackTop(stak Stack) int {
	if !isEmpty(stak) {
		return stak.data[stak.top]
	}
	return -1
}

func main() {
	var stak Stack
	create(&stak)
	push(&stak, 10)
	push(&stak, 20)
	push(&stak, 30)
	push(&stak, 40)
	fmt.Println(peek(stak, 2))
	Display(stak)
}
