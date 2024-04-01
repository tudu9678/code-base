package repository

import (
	"context"

	"gorm.io/gorm"

	"myapp/internal/user/model"
)

type userRepo struct {
	DB *gorm.DB `inject:"db-master"`
}

// NewUserRepo
func NewUserRepo() UserRepo {
	return &userRepo{}
}

func (r *userRepo) WithTnx() *gorm.DB {
	return r.DB
}

func (r *userRepo) WithReplicaTnx() *gorm.DB {
	return r.DB
	//return r.dbRep
}

func (r *userRepo) CreateUser(user *model.User) error {
	if err := r.DB.Create(&user).Error; err != nil {
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
	if err := r.DB.WithContext(ctx).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) CheckUserExist(ctx context.Context, conds model.User) (*model.User, error) {
	var user *model.User
	q := r.DB.Debug().WithContext(ctx)
	if conds.Email != "" {
		q = q.Or("email = ?", conds.Email)
	}
	if conds.PhoneNumber != "" {
		q = q.Or("phone_number = ?", conds.PhoneNumber)
	}
	if conds.UserName != "" {
		q = q.Or("user_name = ?", conds.UserName)
	}
	if err := q.First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}
