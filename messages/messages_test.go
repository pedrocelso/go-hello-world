package messages_test

import (
	"testing"

	msg "github.com/pedrocelso/go-hello-world/messages"
	"github.com/stretchr/testify/assert"
)

func TestMessage(t *testing.T) {
	var mockName = "Pedro"
	var message = msg.GetGreetingMessage(mockName)

	assert.NotNil(t, message)
	assert.Equal(t, "Hello, Pedro!", message)
}

// func TestValidSum(t *testing.T) {
// 	var mockNubers = "23,12,44,1,37"
// 	var sum, err = msg.GetSum(mockNubers)
//
// 	assert.NotNil(t, sum)
// 	assert.Nil(t, err)
// 	assert.Equal(t, "The sum is: 117", sum)
// }
//
// func TestInvalidSum(t *testing.T) {
// 	var mockNubers = "a,b,c,d,e"
// 	var sum, err = msg.GetSum(mockNubers)
//
// 	assert.NotNil(t, err)
// 	assert.Nil(t, sum)
// }
