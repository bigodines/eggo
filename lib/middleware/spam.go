package middleware

import (
	"fmt"
	lib "github.com/bigodines/eggo/lib"
	"github.com/gempir/go-twitch-irc/v2"
	"regexp"
)

type (
	Cmd struct{}
)

// implement the Command interface
func (c Cmd) Execute() error {
	return nil
}

// email validation regexp by W3C
// https://www.w3.org/TR/2016/REC-html51-20161101/sec-forms.html#email-state-typeemail
var SpamEmail = regexp.MustCompile("[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*")

func Spam() func(twitch.Message, chan lib.Command) error {
	return func(m twitch.Message, ch chan lib.Command) error {
		msg := m.(*twitch.PrivateMessage)
		if SpamEmail.MatchString(msg.Message) {
			// TODO: send command to kick/ban
			cmd := Cmd{}
			ch <- cmd
		}

		fmt.Printf("%d", m.GetType())
		return nil
	}
}
