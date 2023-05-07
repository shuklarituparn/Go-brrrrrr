package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == j {
				fmt.Println("same index")
				fmt.Println("Index of I-", i, " and J-", j)

			} //printing the value of i and j if they are at same index
			//fmt.Println(i, " - ", j)
		}
	}
}
