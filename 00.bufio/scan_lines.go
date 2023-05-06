package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// create a new scanner object to scan input from standard input
	scanner := bufio.NewScanner(os.Stdin)

	// use Scan() method to read input line by line
	for scanner.Scan() {
		line := scanner.Text()
		// process the line of input
		fmt.Println("Input:", line)
	}

	if err := scanner.Err(); err != nil {
		// handle any errors that occurred during scanning
		fmt.Println("Error:", err)
	}
}
