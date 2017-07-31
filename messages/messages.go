package messages

import (
	"fmt"
)

// GetGreetingMessage returns an Hello Message to the user
func GetGreetingMessage(name string) (message string) {
	return fmt.Sprintf("Hello, %s!", name)
}

// GetSum sums up comma separated numbers
// func GetSum(numbersString string) (*null.String, error) {
// 	var numbers = strings.Split(numbersString, ",")
// 	var sum = 0
// 	for _, v := range numbers {
// 		i, err := strconv.Atoi(v)
// 		if err != nil {
// 			return nil, err
// 		}
// 		sum += i
// 	}
//
// 	return fmt.Sprintf("The sum is: %v", sum), nil
// }
