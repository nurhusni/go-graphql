package service

import (
	"context"

	"github.com/nurhusni/go-graphql/internal/app/entity"
	"github.com/nurhusni/go-graphql/internal/app/repository"
)

type UserService interface {
	GetUsers(ctx context.Context, params entity.User) (result []entity.User, err error)
	CreateUser(ctx context.Context, params entity.User) (err error)
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(UserRepository repository.UserRepository) UserService {
	return UserServiceImpl{
		UserRepository: UserRepository,
	}
}

func (u UserServiceImpl) GetUsers(ctx context.Context, params entity.User) (result []entity.User, err error) {
	return u.UserRepository.GetUsers(ctx, params)
}

func (u UserServiceImpl) CreateUser(ctx context.Context, params entity.User) (err error) {
	return u.UserRepository.CreateUser(ctx, params)
}
