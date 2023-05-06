package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// create a new scanner object to scan input from standard input
	scanner := bufio.NewScanner(os.Stdin)

	// set the scanner's split function to scan runes instead of lines
	scanner.Split(bufio.ScanRunes)

	// use Scan() method to read input rune by rune
	for scanner.Scan() {
		r := scanner.Text()
		// process the rune of input
		fmt.Println("Rune:", r)
	}

	if err := scanner.Err(); err != nil {
		// handle any errors that occurred during scanning
		fmt.Println("Error:", err)
	}
}
