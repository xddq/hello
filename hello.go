// Simply follows the official go getting-started tutorial and adds idiot-proof
// comments and questions/infos for me to remember.
package main

import (
	"fmt"
	"github.com/xddq/greetings"
	"log"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	// NOTE: I guess a string array/list is called a 'slice of strings' in go?
	names := []string{"Pierre", "Samantha", "Darrin"}

	// Request greeting messages for the names.
	greetings, err := greetings.Hellos(names)

	// Prints an error to the console and exits the program if an error was
	// returned by the greeings.Hello function.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, we print the map of greetings to stdout.
	fmt.Println(greetings)
}
