package repository

import (
	"context"

	"github.com/ryo7000/go-oapi-sample/domain/model"
	"github.com/ryo7000/go-oapi-sample/domain/repository"
	"gorm.io/gorm"
)

type userRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(Conn *gorm.DB) repository.UserRepository {
	return &userRepository{Conn}
}

func (r *userRepository) Create(c context.Context, user *model.User) error {
	return r.Conn.WithContext(c).Create(user).Error
}

func (r *userRepository) Find(c context.Context, id uint64) (*model.User, error) {
	u := &model.User{ID: id}
	err := r.Conn.WithContext(c).First(u).Error
	return u, err
}

func (r *userRepository) FindByEmail(c context.Context, email string) (*model.User, error) {
	u := &model.User{Email: email}
	err := r.Conn.WithContext(c).First(u).Error
	return u, err
}

func (r *userRepository) Update(c context.Context, user *model.User) (*model.User, error) {
	err := r.Conn.WithContext(c).Model(user).Updates(user).Error
	return user, err
}

func (r *userRepository) Delete(c context.Context, id uint64) error {
	u := &model.User{ID: id}
	err := r.Conn.WithContext(c).Delete(u).Error
	return err
}
