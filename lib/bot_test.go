package lib

import (
	"testing"

	"github.com/bigodines/eggo/config"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/stretchr/testify/assert"
)

func TestUse(t *testing.T) {
	fakeMW := func(m twitch.Message) error {
		return nil
	}
	c := config.Config{}
	b := NewBot(c)
	b.Use("onPrivateMessage", fakeMW)
	assert.True(t, true)
}
