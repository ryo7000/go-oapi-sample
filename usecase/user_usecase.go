package usecase

import (
	"context"
	"errors"

	"github.com/ryo7000/go-oapi-sample/domain/model"
	"github.com/ryo7000/go-oapi-sample/domain/repository"
)

type UserUseCase interface {
	Create(c context.Context, user *model.User) error
	Login(c context.Context, email string, password string) (*model.User, error)
	FindById(c context.Context, id uint64) (*model.User, error)
}

type userUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (repo *userUseCase) Create(c context.Context, user *model.User) error {
	hashed, err := model.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashed

	return repo.userRepo.Create(c, user)
}

func (repo *userUseCase) Login(c context.Context, email string, password string) (*model.User, error) {
	user, err := repo.userRepo.FindByEmail(c, email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !user.CheckPassword(password) {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}

func (repo *userUseCase) FindById(c context.Context, id uint64) (*model.User, error) {
	return repo.userRepo.Find(c, id)
}
