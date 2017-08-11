package main

import (
	"fmt"

	msg "github.com/pedrocelso/go-hello-world/messages"
)

func main() {
	input := askForInput("Enter Your Name: ")
	fmt.Printf("\n%s\n\n", msg.GetGreetingMessage(input))
	numbersStr := askForInput("Enter some numbers to sum (comma separated): ")
	sum, err := msg.GetSum(numbersStr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n%s\n\n", sum)
}

func askForInput(message string) string {
	var input string
	fmt.Print(message)
	fmt.Scanln(&input)
	return input
}
