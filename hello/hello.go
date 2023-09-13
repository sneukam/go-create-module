/*
Ok, this code will call functions from the greetings package.

I noticed that the example has package main declared here.
My guess is that the main package is what `go run .` looks for to execute the program.

Also, it looks like import statements can be consolidated.
*/

package main

import (
	"example.com/greetings"
	"fmt"
)

func main() {
	// get a greeting message and print it.
	fmt.Println(greetings.Hello("Spencer"))
}
