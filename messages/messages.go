package messages

import (
	"fmt"
	"strconv"
	"strings"
)

// GetGreetingMessage returns an Hello Message to the user
func GetGreetingMessage(name string) (message string) {
	return fmt.Sprintf("Hello, %s!", name)
}

// GetSum sums up comma separated numbers
func GetSum(numbersString string) (string, error) {
	var numbers = strings.Split(numbersString, ",")
	var sum = 0
	for _, v := range numbers {
		i, err := strconv.Atoi(v)
		if err != nil {
			return "", err
		}
		sum += i
	}

	return fmt.Sprintf("The sum is: %v", sum), nil
}
