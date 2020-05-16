package middleware

import (
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUse(t *testing.T) {
	msg := twitch.PrivateMessage{}
	cmd := make(chan Cmd, 1)
	middleware := Spam()
	err := middleware(&msg, cmd)
	assert.Nil(t, err)

	assert.True(t, true)
}
