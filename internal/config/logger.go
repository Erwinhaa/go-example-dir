package config

import (
	"log"
	"myapp/internal/model"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

type (
	Logger struct {
		standard *log.Logger
		zerolog  *zerolog.Logger
	}
)

func NewZeroLog(envConfig *model.Config) *Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if envConfig.LogMode == "debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	zlog := zerolog.New(os.Stdout).With().Timestamp().Logger()

	return &Logger{
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		&zlog,
	}
}

func (l *Logger) Std() *log.Logger { return l.standard }

func (l *Logger) Zlog() *zerolog.Logger { return l.zerolog }

func (l *Logger) Level(level string) *Logger {
	lv, err := zerolog.ParseLevel(strings.ToLower(level))
	if err == nil {
		*l.zerolog = l.zerolog.Level(lv)
		l.standard.SetOutput(l.zerolog)
	}

	return l
}
