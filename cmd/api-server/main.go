package main

import (
	"myapp/internal/config"
)

func main() {
	envConfig := config.NewViper()
	log := config.NewLogger(envConfig)
	db := config.NewDatabase(envConfig, log)
	gin := config.NewGin()

	config.Bootstrap(&config.BootstrapConfig{
		DB:     db,
		Gin:    gin,
		Config: envConfig,
	})

	webPort := envConfig.Port
	log.Printf("App is running at port %s", webPort)
	err := gin.Run(":" + webPort)
	if err != nil {
		log.Printf("Failed to start server: %v", err)
	}
}
