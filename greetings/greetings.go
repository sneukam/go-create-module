/*
Starting over.
Creating a go module.
*/

package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello() returns a greeting for the named person.
func Hello(name string) (string, error) {

	// if no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("Error: name argument cannot be an empty string.")
	}
	// return a greeting
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// return a random greeting message
func randomFormat() string {

	// a slice (not an array) of messages
	msg_format := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v! Well met!",
	}

	// return a randomly selected message format by specifying
	// a random index for the slice of message formats
	return msg_format[rand.Intn(len(msg_format))]
}
