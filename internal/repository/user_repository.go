package repository

import (
	"context"
	entity "myapp/internal/entity/user"

	"gorm.io/gorm"
)

type (
	UserRepository interface {
		CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
		GetUserById(ctx context.Context, id int) (*entity.User, error)
		GetUsers(ctx context.Context) ([]*entity.User, error)
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	if err := r.db.Omit("updated_at").Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetUserById(ctx context.Context, id int) (*entity.User, error) {
	var user entity.User

	if err := r.db.Where("id = ? ", id).Take(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetUsers(ctx context.Context) ([]*entity.User, error) {
	var users []*entity.User

	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
