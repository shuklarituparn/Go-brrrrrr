package main

import (
	"fmt"
	"github.com/shuklarituparn/Go-brrrrrr/04_scope/01_package-scope/02_visibility/vis/vis"
)

func main() {
	fmt.Println(vis.MyName)
	vis.PrintVar()
}
