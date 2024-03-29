package repository

import (
	"context"

	"myapp/internal/user/model"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB `inject:"db"`
}

// NewUserRepo
func NewUserRepo() UserRepo {
	return &userRepo{}
}

func (r *userRepo) WithTnx() *gorm.DB {
	return r.db
}

func (r *userRepo) CreateUser(user *model.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

// UpdateUser
func (r *userRepo) UpdateUser(tx *gorm.DB, user *model.User) error {
	res := tx.Save(user)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *userRepo) FindByCondition(ctx context.Context, user model.User) (*model.User, error) {
	if err := r.db.WithContext(ctx).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
