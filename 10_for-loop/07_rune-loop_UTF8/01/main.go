package main

import "fmt"

func main() {
	for i := 250; i <= 340; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i))) //string(i) converts the int to string
	} //the []byte(string(i)) then converts it to byte slice
	foo := "a"
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)
}

/*
NOTE:
Some operating systems (Windows) might not print characters where i < 256

If you have this issue, you can use this code:

fmt.Println(i, " - ", string(i), " - ", []int32(string(i)))

UTF-8 is the text coding scheme used by Go.

UTF-8 works with 1 - 4 bytes.

A byte is 8 bits.

[]byte deals with bytes, that is, only 1 byte (8 bits) at a time.

[]int32 allows us to store the value of 4 bytes, that is, 4 bytes * 8 bits per byte = 32 bits.
*/
/*
A byte slice is a sequence of bytes in Go, similar to an array of bytes. It is represented by the []byte type.
Byte slices are commonly used to store and manipulate binary data, such as image or audio files.

Byte slices can be created from a string using the []byte() conversion function. For example, []byte("hello") would
create a byte slice containing the ASCII values for the characters 'h', 'e', 'l', 'l', and 'o'.

Basically creating an array of ACII values of the char of a string

Byte slices can be manipulated using various built-in functions and methods, such as len() to get the length of the
slice, copy() to copy data between slices, and append() to add elements to the slice.

250  -  ú  -  [195 186]

we get 195 and 186 as character acute u is represented in UTF-8 by two bytes
*/
