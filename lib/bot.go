package lib

import (
	"github.com/bigodines/eggo/config"
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/rs/zerolog/log"
)

type (
	botService struct {
		conf config.Config
	}
)

func New(c config.Config) *botService {
	return &botService{
		conf: c,
	}
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
