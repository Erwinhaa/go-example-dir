package config

import (
	"fmt"
	"myapp/internal/model"
	"time"

	"github.com/rs/zerolog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(envConfig *model.Config, log *zerolog.Logger) *gorm.DB {
	dbUser := envConfig.DBUser
	dbPassword := envConfig.DBPassword
	dbHost := envConfig.DBHost
	dbPort := envConfig.DBPort
	database := envConfig.DBDatabase

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(&zerologWriter{Logger: log}, logger.Config{
			SlowThreshold:             time.Second * 5,
			Colorful:                  false,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			LogLevel:                  logger.LogLevel(log.GetLevel()),
		}),
	})
	if err != nil {
		log.Fatal().Msgf("failed to connect database: %v", err)
	}

	connection, err := db.DB()
	if err != nil {
		log.Fatal().Msgf("failed to connect database: %v", err)
	}
}

type zerologWriter struct {
	Logger *zerolog.Logger
}

func (l *zerologWriter) Printf(message string, args ...interface{}) {
	l.Logger.Printf(message, args...)
}
