package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sum(a int, b int) int {
	var c = a + b
	return c
}

func difference(a int, b int) int {
	var c = a - b
	return c
}

func multiplication(a int, b int) int {
	var c = a * b
	return c
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Welcome to our little program: ")
	fmt.Print("Press 1 to add \n")
	fmt.Print("Press 2 to subtract\n ")
	fmt.Print("Press 3 to multiply \n")
	options, _ := reader.ReadString('\n')
	name, _ := strconv.Atoi(strings.TrimSpace(options))
	switch name {

	case 1:
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Println("Enter two numbers:")
		scanner.Scan()
		input1, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid input")
			return
		}

		scanner.Scan()
		input2, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid input")
			return
		}
		fmt.Println("The sum is:", sum(input1, input2))

	case 2:
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Println("Enter two numbers:")
		scanner.Scan()
		input1, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid input")
			return
		}

		scanner.Scan()
		input2, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid input")
			return
		}
		fmt.Println("The difference is:", difference(input1, input2))

	case 3:
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Println("Enter two numbers:")
		scanner.Scan()
		input1, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid input")
			return
		}

		scanner.Scan()
		input2, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid input")
			return
		}
		fmt.Println("The product is:", multiplication(input1, input2))

	default:
		fmt.Println("Wrong option entered, please try again:(")
	}
}
