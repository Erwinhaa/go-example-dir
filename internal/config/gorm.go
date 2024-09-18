package config

import (
	"fmt"
	"myapp/internal/model"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewDatabase(envConfig *model.Config, log *Logger) *gorm.DB {
	dbUser := envConfig.DBUser
	dbPassword := envConfig.DBPassword
	dbHost := envConfig.DBHost
	dbPort := envConfig.DBPort
	database := envConfig.DBDatabase

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.New(log.standard, logger.Config{
			SlowThreshold:             time.Second * 5,
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
			LogLevel:                  logger.LogLevel(logger.Info),
		}),
	})
	if err != nil {
		log.Zlog().Fatal().Msgf("failed to connect database: %v", err)
	}

	_, err = db.DB()
	if err != nil {
		log.Zlog().Fatal().Msgf("failed to connect database: %v", err)
	}

	return db
}

// type zerologWriter struct {
// 	Logger *zerolog.Logger
// }

// func (l *zerologWriter) Printf(message string, args ...interface{}) {
// 	l.Logger.Printf(message, args...)
// }
