package config

import (
	"myapp/internal/delivery/http/v1/controller"
	"myapp/internal/delivery/http/v1/route"
	"myapp/internal/model"
	"myapp/internal/repository"
	"myapp/internal/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB     *gorm.DB
	Config *model.Config
	Gin    *gin.Engine
}

func Bootstrap(config *BootstrapConfig) {
	userRepo := repository.NewUserRepository(config.DB)
	userUC := usecase.NewUserUseCase(userRepo)
	userController := controller.NewUserController(userUC)
	domain := config.Gin.Group("users")

	route.UserRoutes(domain, userController)
}
