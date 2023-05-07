package main

import "fmt"

func main() {
	/*	i := 0
		for {
			i++
			if i%2 == 0 {
				continue
			}
			fmt.Println(i)
			if i >= 50 {
				break
			}
		}*/

	i := 0
	for {
		i++
		if i <= 10 {
			fmt.Print(" counting till 10\n", i) //here we'll print i before "counting till 10"
			//starting the count from i=0
		} else {
			fmt.Print("Till 10, not a number more")
			break
		}
	}

}
