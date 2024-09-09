package entity

import (
	"myapp/pkg/utils"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Password string
	Name     string
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var err error

	u.Password, err = utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	return nil
}
