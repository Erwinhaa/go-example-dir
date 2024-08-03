package config

import (
	"myapp/internal/model"

	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB     *gorm.DB
	Config *model.Config
}

func Bootstrap(config *BootstrapConfig) {

}
