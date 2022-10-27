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

	greeting, err := greetings.Hello("Pierre")

	// Prints an error to the console and exits the program if an error was
	// returned by the greeings.Hello function.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, we print the greeting to stdout.
	fmt.Println(greeting)
}
