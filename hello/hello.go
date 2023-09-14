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

	singleName()
	manyNames()
}

func singleName() {

	msg, err := greetings.Hello("Spencer")

	if err != nil {
		log.Fatal("Error: Name in Hello() cannot be empty.")
	}

	fmt.Println(msg)
	fmt.Println()
}

func manyNames() {

	names := []string{
		"Bobby",
		"Jimmy",
		"Carl",
		"Ryan",
	}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal("Error: One or more of the names in manyNames() is empty.")
	}

	for _, name := range names {
		fmt.Println(messages[name])
	}

	// ok, this is cool.
	// bit more high level functionality.
	// we don't have to loop over the map.
	fmt.Println("\nHEDGIES ARE FUKT!!!")
}
