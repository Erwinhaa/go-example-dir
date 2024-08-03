package main

import "myapp/internal/config"

func main() {
	envConfig := config.NewViper()
	log := config.NewLogger(envConfig)
	db := config.NewDatabase(envConfig)

	config.Bootstrap()
}
