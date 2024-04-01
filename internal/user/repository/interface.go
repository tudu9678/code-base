package repository

import (
	"context"

	"myapp/internal/user/model"

	"gorm.io/gorm"
)

// UserRepo
type UserRepo interface {
	WithTnx() *gorm.DB

	// User
	CreateUser(user *model.User) error
	UpdateUser(tx *gorm.DB, user *model.User) error
	FindByCondition(ctx context.Context, conds model.User) (*model.User, error)
	CheckUserExist(ctx context.Context, conds model.User) (*model.User, error)
}
