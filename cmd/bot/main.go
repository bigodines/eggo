package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/gempir/go-twitch-irc/v2"

	"github.com/bigodines/eggo/config"
	libbot "github.com/bigodines/eggo/lib"
)

type ()

func main() {
	conf, err := config.Load(getEnv())
	if err != nil {
		panic(err)
	}

	if conf.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	twitchClient := twitch.NewClient(conf.Username, conf.OAuthToken)
	twitchClient.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println(message.Message)
	})

	twitchClient.Join("gempir")

	err = twitchClient.Connect()
	if err != nil {
		panic(err)
	}

	bot := libbot.New(conf)
	log.Debug().Msg("Running bot")
	// this is my pet project, I name methods as I want!!!
	err = bot.Unleash()
	if err != nil {
		panic("sorry j. cannot do it")
	}

	log.Info().Msg("Done")

}

// helper function to figure environment
// defaults to "development"
func getEnv() string {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
	}
	return env
}
