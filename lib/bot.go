package lib

import (
	"github.com/bigodines/eggo/config"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/rs/zerolog/log"
)

type (
	// Middleware is the interface every message handler must implement
	Middleware func(twitch.Message, chan Command) error

	// Commands are actions the bot can take
	Command interface {
		Execute() error
	}

	botService struct {
		conf  config.Config
		mw    map[string][]Middleware
		cmdCh chan Command
	}
)

var (
	// TODO: implement missing events and have them as consts
	validEvents = []string{"onPrivateMessage"}
)

func NewBot(c config.Config) *botService {
	mw := make(map[string][]Middleware, 0)
	cmdCh := make(chan Command, 10)

	bot := &botService{
		conf:  c,
		mw:    mw,
		cmdCh: cmdCh,
	}

	return bot
}

// Use adds a middleware as a listener to an event
func (b *botService) Use(eventName string, mw Middleware) {
	if validEvent(eventName) == false {
		log.Warn().Str("event", eventName).Msg("you tried to add middleware to invalid event type")
		return
	}
	if len(b.mw[eventName]) == 0 {
		b.mw[eventName] = make([]Middleware, 0)
	}
	b.mw[eventName] = append(b.mw[eventName], mw)
}

func validEvent(name string) bool {
	for _, v := range validEvents {
		if v == name {
			return true
		}
	}
	return false
}

func (b *botService) onPvtMsg(m twitch.PrivateMessage) {
	for _, mw := range b.mw["onPrivateMessage"] {
		// TODO: err handle
		mw(&m, b.cmdCh)
	}

}

func (b *botService) Unleash() error {
	twitchClient := twitch.NewClient(b.conf.Name, b.conf.OAuthToken)
	twitchClient.OnPrivateMessage(b.onPvtMsg)

	twitchClient.Join(b.conf.Name)

	err := twitchClient.Connect()
	if err != nil {
		return err
	}

	log.Debug().Msg("Connected")

	return nil
}
