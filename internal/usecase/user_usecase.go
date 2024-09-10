package usecase

import (
	"context"
	"myapp/internal/dto"
	entity "myapp/internal/entity/user"
	"myapp/internal/repository"
	"myapp/pkg/utils"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, payload dto.CreateUserRequest) (user *entity.User, err error)
	GetUserById(ctx context.Context, id int) (*entity.User, error)
}

type usecase struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &usecase{repo}
}

func (uc *usecase) CreateUser(ctx context.Context, request dto.CreateUserRequest) (user *entity.User, err error) {
	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	user = &entity.User{
		Name:     request.FullName,
		Password: hashedPassword,
	}

	return uc.repo.CreateUser(ctx, user)
}

func (uc *usecase) GetUserById(ctx context.Context, id int) (*entity.User, error) {
	return uc.repo.GetUserById(ctx, id)
}
