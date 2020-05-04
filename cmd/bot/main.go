package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

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

	log.Debug().Interface("Config", conf).Msg("Config loaded")
	bot := libbot.New(conf)
	log.Debug().Msg("Running bot")
	// this is my pet project, I name methods as I want!!!
	err = bot.Unleash()
	if err != nil {
		log.Fatal().Err(err).Msg("Sorry j. cannot do it")
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
