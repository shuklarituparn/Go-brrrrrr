package main

import "fmt"

func main() {

	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := `Do you like my hat?`
	g := 'M'

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)

	/*Inside a function, the := short assignment statement can be used in place of a var
	declaration with implicit type.

	Outside a function, every statement begins with a keyword (var, func, and so on) and so the
	:= construct is not available.
	*/
}
