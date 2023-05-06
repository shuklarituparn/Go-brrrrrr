package main

import "fmt"

/*

Algorithm for tower of hanoi is:


function hanoi(n, source, destination, auxiliary):
    if n == 1:
        move disk from source to destination
    else:
        hanoi(n-1, source, auxiliary, destination)
        move disk from source to destination
        hanoi(n-1, auxiliary, destination, source)

*/

// TowerOfHanoi
//
/*The algorithm can be implemented as follows:

If there is only one disk, move it from the source peg to the destination peg.
Otherwise, recursively move the top n-1 disks from the source peg to the auxiliary peg using the
destination peg as the temporary peg.
Move the largest disk from the source peg to the destination peg.
Recursively move the n-1 disks from the auxiliary peg to the destination peg using the source peg
as the temporary peg.*/
///*
func TowerOfHanoi(n int, A int, B int, C int) { //we hav
	if n > 0 {
		TowerOfHanoi(n-1, A, C, B)
		fmt.Println(A, C)
		TowerOfHanoi(n-1, B, A, C)
	}

}

func main() {
	TowerOfHanoi(4, 1, 2, 3)

}

/*
Tower of hanoi uses the recursive approach to move the "n" disks from peg to peg c using peg B
Base case is n=0 which means there is no disk to move.

The function calls itself with n-1 disks from peg 'A' to peg 'B' using peg 'C' as temporary peg
Once, the n-1 disks have been, the function prints the move required to move the nth disk from peg "A" to peg "C"

Finally the function calls itself to move the disks from B to C using A as the temporary peg

The output is in the form of "(A,C)" where A is the peg the disk is moved from and "C" is the peg the disk is moved to
*/
