package main

import (
	"myapp/internal/config"
	"myapp/internal/delivery/http/v1/route"
)

func main() {
	envConfig := config.NewViper()
	log := config.NewLogger(envConfig)
	db := config.NewDatabase(envConfig, log)
	gin := config.NewGin()

	config.Bootstrap(&config.BootstrapConfig{
		DB:     db,
		Config: envConfig,
	})

	route.UserRoutes(gin)

	webPort := envConfig.Port
	// err :=
}
