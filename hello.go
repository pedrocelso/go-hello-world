package main

import (
	"fmt"

	msg "github.com/pedrocelso/go-hello-world/messages"
)

func main() {
	var input string
	fmt.Print("Enter Your Name: ")
	fmt.Scanf("%s", &input)
	fmt.Printf("\n%s\n\n", msg.GetGreetingMessage(input))
}
