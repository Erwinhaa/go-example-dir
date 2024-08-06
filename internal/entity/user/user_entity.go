package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Password string
	Name     string
}
