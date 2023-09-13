/*
Creating modules.
*/

package main // this needs to be in each file under the /hello folder to denote they are part of the main package.

import (
	"fmt"
	"golang.org/x/example/hello/reverse"
)

func main() {
	fmt.Println(reverse.String("Hello"))
}
