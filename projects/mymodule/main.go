package main

import (
	"fmt"
	"log"
	"mymodule/mypackage"
	"mymodule/greetings"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, Modules!")
			mypackage.PrintHello()
			// Get a greeting message and print it
			//message := greetings.Hello("Gladys")
			//fmt.Println(message)


		},
	}
	fmt.Println("Calling cmd.Execute()!")
	cmd.Execute()
	
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(messages)

	//// Request a greeting message.
	//message, err := greetings.Hello("Gladys")
	//// If an error was returned, print it to the console and
	//// exit the program.
	//if err != nil {
	//	log.Fatal(err)
	//}
	//// If no error was returned, print the returned message
	//// to the console.
	//fmt.Println(message)
}
