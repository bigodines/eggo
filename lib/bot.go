package lib

import (
	"github.com/bigodines/eggo/config"
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/rs/zerolog/log"
)

type (
	botService struct {
		conf       config.Config
		middleware map[string][]Middleware
	}

	Middleware func(twitch.Message) error

	// MessageHandler is the interface shared by all message handlers
	MessageHandler interface {
		Handle(twitch.Message) error
	}
)

var (
	validEvents = []string{"onPrivateMessages"}
)

func New(c config.Config) *botService {
	bot := &botService{
		conf: c,
	}

	return bot
}

func (b *botService) Use(eventName string, mw Middleware) {
	if validEvent(eventName) == false {
		log.Warn().Str("event", eventName).Msg("you tried to add middleware to invalid event type")
		return
	}
	if len(b.middleware[eventName]) == 0 {
		b.middleware[eventName] = make([]Middleware, 0)
	}
	b.middleware[eventName] = append(b.middleware[eventName], mw)
}

func validEvent(name string) bool {
	for _, v := range validEvents {
		if v == name {
			return true
		}
	}
	return false
}

func (b *botService) Unleash() error {
	twitchClient := twitch.NewClient(b.conf.Name, b.conf.OAuthToken)
	twitchClient.OnPrivateMessage(func(message twitch.PrivateMessage) {
		log.Debug().Msg(message.Message)
	})

	twitchClient.Join(b.conf.Name)

	err := twitchClient.Connect()
	if err != nil {
		return err
	}

	log.Debug().Msg("Connected")

	return nil
}
