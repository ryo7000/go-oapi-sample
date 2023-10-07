package repository

import (
	"context"

	"github.com/ryo7000/go-oapi-sample/domain/model"
)

type UserRepository interface {
	Create(c context.Context, user *model.User) error
	Find(c context.Context, id uint64) (*model.User, error)
	FindByEmail(c context.Context, email string) (*model.User, error)
	Update(c context.Context, user *model.User) (*model.User, error)
	Delete(c context.Context, id uint64) error
}
