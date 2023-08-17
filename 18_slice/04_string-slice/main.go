package main

import "fmt"

func main() {

	greeting := []string{
		"Good morning!",
		"Bonjour!",
		"dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",   //even the last string must have a comma
	}

	for i, currentEntry := range greeting {
		fmt.Println(i, currentEntry)    //print current index and value
	}

	for j := 0; j < len(greeting); j++ { //
		fmt.Println(greeting[j])
	}

}
