package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	/*
		defer is used to schedule a function call to be executed immediately before the function that contains
		the "defer" statement returns

		Here we are using defer to schedule a function call at the end of this main function no matter what happens
		for example if something happens during the reading part of the file it'll still get closed and free up systems
		resource
	*/

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	str := string(data)
	fmt.Println("Contents of file:")
	fmt.Println(str)
}
