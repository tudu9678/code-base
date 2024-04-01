// Package service handles the business logic of authentication.
package service

import (
	"context"
	"time"

	_ "myapp/config"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	cCm "myapp/core/common"
	in "myapp/core/initialize"
	"myapp/core/initialize/auth"
	"myapp/internal/user/dto"
	"myapp/internal/user/model"
	"myapp/internal/user/repository"
)

type user struct {
	//conf   config.Config       `inject:"configs"`
	repo    repository.UserRepo `inject:"user-repo"`
	logger  in.Logging          `inject:"logging"`
	jwtAuth *auth.JWTAuth       `inject:"jwt-auth"`
	//jwtAuth *auth.JWTAuth
	//zapLog  *zap.Logger
}

func NewUserService() User {
	return &user{}
}

func (u *user) Register(ctx context.Context, req *dto.CreateUserReq) (*dto.RegisterRes, error) {

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		u.logger.Error(ctx, "[Register] GenerateFromPassword err: %v", err)

		return nil, err
	}

	dob, err := time.Parse(cCm.DateFm, req.Dob)
	if err != nil {
		u.logger.Error(ctx, "[Register] Dob time.Parse err: %v", err)
		return nil, err
	}

	userModel := &model.User{
		FullName:    req.FullName,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		UserName:    req.UserName,
		Password:    string(password),
		Dob:         &dob,
		LatestLogin: &time.Time{},
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

	if err := u.repo.CreateUser(userModel); err != nil {
		return nil, err
	}

	tokenExpire := "3600"
	auth, err := u.jwtAuth.CreateToken(&auth.User{
		ID:       userModel.ID.String(),
		UserName: userModel.UserName,
		FullName: userModel.FullName,
	}, tokenExpire)
	if err != nil {
		u.logger.Info(ctx, "[Login] username %v - CreateToken err %v", req.UserName, err)
		return nil, err
	}

	return &dto.RegisterRes{
		UserInfo: &dto.UserRes{
			ID:          userModel.ID.String(),
			Email:       userModel.Email,
			PhoneNumber: userModel.PhoneNumber,
			FullName:    userModel.FullName,
		},
		Auth: &dto.AuthRes{
			UserID:       userModel.ID.String(),
			AccessToken:  auth.AccessToken,
			RefreshToken: auth.RefreshToken,
			ExpiresIn:    uint64(auth.AccessTokenDuration),
		},
	}, nil
}

func (u *user) Login(ctx context.Context, req *dto.LoginReq) (*dto.AuthRes, error) {
	u.logger.Info(ctx, "[Login] username: %v", zap.Any("req.Username", req.Username))
	//userModel, err := u.repo.FindByCondition(ctx, model.User{})

	// if err := bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(req.Password)); err != nil && u.conf.Env != "local" {
	// 	return nil, ce.New(ce.ErrInvalidCredentials, "invalid password", nil)
	// }

	// tokenExpire := "3600"

	// if err != nil {
	// 	u.logger.Error("[Login] get Env %s err: %v", zap.Any("constants.OTP_EXPIRE_TIME", constants.OTP_EXPIRE_TIME), zap.Error(err))
	// }

	// auth, err := u.jwtAuth.CreateToken(userModel, tokenExpire)
	// if err != nil {
	// 	u.logger.Info("[Login] username %v - CreateToken err %v", zap.Any("req.Username", req.Username), zap.Error(err))
	// 	return nil, ce.New(ce.ErrInvalidCredentials, "something went wrong", nil)
	// }

	//u.notiService.ExternalPushNoti(ctx, enum.NotificationActionLogin, userModel.ID, nil)

	//ctx = context.WithValue(ctx, middleware.UserIDKey, userModel.ID.String())

	return &dto.AuthRes{
		// UserID:       userModel.ID.String(),
		// AccessToken:  auth.AccessToken,
		// RefreshToken: auth.RefreshToken,
		// ExpiresIn:    uint64(auth.AccessTokenDuration),
	}, nil
}
