package middleware

import (
	"fmt"
	"github.com/gempir/go-twitch-irc/v2"
	"regexp"
)

type (
	Cmd struct{}
)

// implement the Command interface
func (c *Cmd) Execute() {
}

// email validation regexp by W3C
// https://www.w3.org/TR/2016/REC-html51-20161101/sec-forms.html#email-state-typeemailconst
var SpamEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func Spam() func(twitch.Message, chan Cmd) error {
	return func(m twitch.Message, ch chan Cmd) error {
		fmt.Printf("%d", m.GetType())
		return nil
	}
}
