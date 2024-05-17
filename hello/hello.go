package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of Logger
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message
	message, err := greetings.Hello("James")
	
	// Error handling
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}