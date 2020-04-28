package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/bigodines/eggo/config"
	_ "github.com/bigodines/eggo/lib"
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

	//bot := libbot.New()
	// this is my pet project, I name methods as I want!!!
	//bot.Unleash()
	fmt.Printf("%+v\n", conf)

	fmt.Println("vim-go")
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
