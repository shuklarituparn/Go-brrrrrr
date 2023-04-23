package main

import (
	"fmt"
	"github.com/shuklarituparn/Go-brrrrrr/02_package/icomefromalaska"
	"github.com/shuklarituparn/Go-brrrrrr/02_package/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	fmt.Println(stringutil.MyName)
	fmt.Println(winniepooh.BearName)
}

/*In order to implement a file or a package in GO and to use
it, go to the folder, if you're hosting that file on the github
then initialize the go mod there

go mod init github.com/shuklarituparn/Go-brrrrrr/02_package

it will create a go.mod file in the repo containing the code that we wanna use and then we can use
using just import keyword.
*/
