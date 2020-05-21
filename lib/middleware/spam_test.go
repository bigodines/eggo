package middleware

import (
	lib "github.com/bigodines/eggo/lib"
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUse(t *testing.T) {
	msg := twitch.PrivateMessage{}
	cmd := make(chan lib.Command, 1)
	middleware := Spam()
	err := middleware(&msg, cmd)
	assert.Nil(t, err)

	assert.True(t, true)
}
