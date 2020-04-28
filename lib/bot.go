package lib

import (
	"github.com/bigodines/eggo/config"
)

type (
	botService struct {
	}
)

func New(c config.Config) *botService {
	return &botService{}
}
