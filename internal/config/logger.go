package config

import (
	"myapp/internal/model"
	"os"

	"github.com/rs/zerolog"
)

func NewLogger(envConfig *model.Config) *zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if envConfig.LogMode == "debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	logger := zerolog.New(os.Stdout)

	return &logger
}
