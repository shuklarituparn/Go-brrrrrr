package main

import (
	"bufio" //we import bufio package for reading of the input
	"fmt"
	"os" //we import the os package for accessing the standard input stream
)

func main() {
	reader := bufio.NewReader(os.Stdin) //we create new bufio.reader object by passing os.stdin

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n') //we tell function to stop reading when it encounters newline or enter
	fmt.Printf("Hello, %s\n", name)    //we print the greet message

	//here after the name, we use _ to denote that we'll discard the second value that this expression will return

	//in our case it's the error check, as bufio when encountering an error throws the error

	//or we can implement error check as

	/*
		if err!= nil{
				//handle the error
		}
	*/
}
