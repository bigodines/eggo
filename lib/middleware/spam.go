package middleware

import (
	"github.com/gempir/go-twitch-irc/v2"
)

func Spam() func(twitch.Message) error {
	return func(twitch.Message) error {
		return nil
	}
}
