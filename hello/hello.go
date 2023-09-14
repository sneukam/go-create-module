/*
Creating a hello module.
Hopefully I can get it to connect to greetings.go this time.
*/

package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main() {
	// set properties of the predefined Logger
	// including the log entry prefix and a flag to disable printing the :
	// the time, source file, and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greetings message and print it.
	// we are intentionally throwing an error below to practice handling it.
	message, err := greetings.Hello("Spencer")

	// throw error or print to terminal.
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
