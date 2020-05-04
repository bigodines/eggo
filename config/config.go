package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	yaml "github.com/ghodss/yaml"
	"github.com/vrischmann/envconfig"
)

type (
	// Configuration options
	Config struct {
		Environment string `envconfig:"default=development"`
		// channel name
		Channel string `json:"channel" envconfig:"-"`
		// oauth token can't be set in the yml file to prevent people from checking in their tokens
		OAuthToken string `json:"oauth_token" envconfig:"default=THIS_IS_READ_FROM_ENV"`
		// bot username (optional)
		Name string `json:"name" envconfig:"-"`
		// Flood (too many messages) protection configuration
		Flood FloodConfig `json:"flood" envconfig:"-"`
	}

	FloodConfig struct {
		Enabled bool `json:"enabled" envconfig:"-"`
		Lines   int  `json:"lines" envconfig:"-"`
		// Interval in seconds
		Interval int `json:"interval" envconfig:"-"`
	}
)

// Load configuration struct into memory
func Load(env string) (Config, error) {
	config, err := configFromFile(env)
	if err != nil {
		return config, err
	}
	config.Environment = env

	// override with environment variables as needed and accepted by `Config`
	err = envconfig.Init(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}

// configFromFile reads configuration file for environment and return a Config struct
func configFromFile(env string) (Config, error) {
	env = strings.ToLower(env)
	// attempt to read config from {{environment}}.yml (defaults to development.yml)
	var fname string
	var conf Config
	if _, err := os.Stat(fmt.Sprintf("config/%s.yml", env)); err == nil {
		fname = fmt.Sprintf("config/%s.yml", env)
	} else {
		fname = fmt.Sprintf("config/development.yml")
	}

	ymlFile, err := ioutil.ReadFile(fname)
	if err != nil {
		return conf, err
	}

	err = yaml.Unmarshal(ymlFile, &conf)
	if err != nil {
		return conf, err
	}

	return conf, nil

}
