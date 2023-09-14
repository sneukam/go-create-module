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

// Hellos() reutrns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {

	// a map to associate names with messages
	messages := make(map[string]string)

	// loop through received slice of names,
	// calling the Hello() function to get a message for each name.
	for _, name := range names {
		greeting_message, err := Hello(name)

		// return if error (error = empty name)
		if err != nil {
			return nil, err
		}

		// otherwise, capture the greeting for the user's name
		messages[name] = greeting_message
	}

	return messages, nil
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
