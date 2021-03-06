package bridge

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorNotification_Notify(t *testing.T) {
	sender := NewEmailMessageSender([]string{"test@test.com"})
	n := NewErrorNotification(sender)
	err := n.Notify("test msg")

	assert.Nil(t, err)
}
