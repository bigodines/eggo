package middleware

import (
	lib "github.com/bigodines/eggo/lib"
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDetect(t *testing.T) {
	msg := twitch.PrivateMessage{
		Message: "hello. i'm here to avertise my email spammy@badguys.com",
	}
	cmd := make(chan lib.Command, 1)
	middleware := Spam()
	err := middleware(&msg, cmd)
	assert.Nil(t, err)

	select {
	case c := <-cmd:
		assert.NotNil(t, c)
	case <-time.After(2 * time.Second):
		assert.Fail(t, "bye")
	}

	assert.True(t, true)
}
