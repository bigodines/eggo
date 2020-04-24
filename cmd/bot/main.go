package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/bigodines/eggo/config"
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

	fmt.Println("vim-go")
}

// helper function to figure environment
func getEnv() string {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
	}
	return env
}
