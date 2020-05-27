package lib

import (
	"testing"

	"github.com/bigodines/eggo/config"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/stretchr/testify/assert"
)

var ()

func TestUse(t *testing.T) {
	called := false
	fakeMW := func(m twitch.Message, ch chan Command) error {
		called = true
		return nil
	}
	c := config.Config{}
	b := NewBot(c)
	assert.Equal(t, 0, len(b.mw["onPrivateMessage"]))
	b.Use("onPrivateMessage", fakeMW)
	b.onPvtMsg(twitch.PrivateMessage{})
	assert.True(t, called)
	assert.Equal(t, 1, len(b.mw["onPrivateMessage"]))

	// test invalid events
	b.Use("bogus", fakeMW)
	assert.Equal(t, 0, len(b.mw["bogus"]))
}
