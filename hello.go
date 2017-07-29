package main

import "fmt"

func main() {
	var input string
	fmt.Print("Enter Your Name: ")
	fmt.Scanf("%s", &input)
	fmt.Printf("\nHello, %s!\n\n", input)
}
