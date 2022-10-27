package main

import (
	"fmt"
	"github.com/xddq/greetings"
)

func main() {
	message := greetings.Hello("Pierre")
	fmt.Println(message)
}
