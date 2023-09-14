/*
Starting over.
Creating a go module.
*/

package greetings

import "fmt"

func Hello(name string) string {
	// return a greeting
	message := fmt.Sprintf("Hello, %v. Welcome!", name)
	return message
}
