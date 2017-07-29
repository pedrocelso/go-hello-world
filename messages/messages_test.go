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
