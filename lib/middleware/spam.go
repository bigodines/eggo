package middleware

import (
	"fmt"
	"github.com/gempir/go-twitch-irc/v2"
)

func Spam() func(twitch.Message) error {
	return func(m twitch.Message) error {
		fmt.Printf("%d", m.GetType())
		return nil
	}
}
