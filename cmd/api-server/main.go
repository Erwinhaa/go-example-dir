package main

import (
	"context"
	"fmt"
	"myapp/internal/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	envConfig := config.NewViper()
	log := config.NewZeroLog(envConfig)
	db := config.NewDatabase(envConfig, log)
	gin := config.NewGin()

	config.Bootstrap(&config.BootstrapConfig{
		DB:     db,
		Gin:    gin,
		Config: envConfig,
	})

	webPort := envConfig.Port

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: gin.Handler(),
	}

	go func() {
		log.Std().Printf("App is running at port %s", webPort)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Zlog().Fatal().Err(err).Msgf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Zlog().Info().Msg("Server shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Zlog().Fatal().Err(err).Msg("Server Shutdown: ")
	}
}
