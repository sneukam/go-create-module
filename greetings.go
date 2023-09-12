/*
Remember, a module is made up of packages.
A package is all the files and functions within the same directory.

A module can be published online for others to download and use.

A function whose name starts with a capital letter can be called by a function not in the same pacakge.
This is called an exported function.

:= is a shortcut for decarling and initializing a variable in the same line.

*/

package greetings

import "fmt"

// Return a greeting for the named person.
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
