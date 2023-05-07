package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// create a new scanner object to scan input from standard input
	scanner := bufio.NewScanner(os.Stdin)

	// set the scanner's split function to scan words instead of lines
	scanner.Split(bufio.ScanWords)

	// use Scan() method to read input word by word
	for scanner.Scan() {
		word := scanner.Text()
		// process the word of input
		fmt.Println("Word:", word)
	}

	if err := scanner.Err(); err != nil {
		// handle any errors that occurred during scanning
		fmt.Println("Error:", err)
	}
}
