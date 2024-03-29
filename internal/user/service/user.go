// Package service handles the business logic of authentication.
package service

import (
	"context"
	"strings"

	_ "myapp/config"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	in "myapp/core/initialize"
	"myapp/internal/user/dto"
	"myapp/internal/user/model"
	"myapp/internal/user/repository"
)

type user struct {
	//conf   config.Config       `inject:"configs"`
	repo   repository.UserRepo `inject:"user-repo"`
	logger in.Logging          `inject:"logging"`
	//jwtAuth *auth.JWTAuth
	//zapLog  *zap.Logger
}

func NewUserService() User {
	return &user{}
}

func (u *user) Register(ctx context.Context, req *dto.CreateUserReq) (*dto.RegisterRes, error) {
	pCode := req.PhoneCode
	pNum := req.PhoneNumber
	// if err := u.validateCreateUser(req); err != nil {
	// 	return nil, err
	// }

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		u.logger.Error(ctx, "[Register] GenerateFromPassword err: %v", err)

		return nil, err
	}

	userModel := &model.User{
		Email:       strings.ToLower(req.Email),
		PhoneCode:   pCode,
		PhoneNumber: pNum,
		Password:    string(password),
		//Status:      model.UserStatusActive,
	}
	userModel.CreatedBy = strings.ToLower(req.Email)
	userModel.UpdatedBy = strings.ToLower(req.Email)

	if err := u.repo.CreateUser(userModel); err != nil {
		return nil, err
	}

	//go u.SendPhoneOtp(context.Background(), userModel.ID, pCode, pNum, model.OtpActionRegister)
	//go u.SendEmailOtp(context.Background(), userModel.ID, userModel.Email, model.MailTypeRegister)

	//tokenExpire := "3600"

	// if err != nil {
	// 	u.logger.Error("[Register] get Env %s err: %v", zap.Any("constants.TOKEN_TTL", constants.TOKEN_TTL), zap.Error(err))
	// }

	// auth, err := u.jwtAuth.CreateToken(userModel, tokenExpire)
	// if err != nil {
	// 	u.logger.Info("[Login] username %v - CreateToken err %v", zap.Any("userModel.Email", userModel.Email), zap.Error(err))
	// 	return nil, ce.New(ce.ErrorSomethingWentWrong, "something went wrong", nil)
	// }

	return &dto.RegisterRes{
		UserInfo: &dto.UserRes{
			ID:          userModel.ID.String(),
			Email:       userModel.Email,
			PhoneCode:   userModel.PhoneCode,
			PhoneNumber: userModel.PhoneNumber,
		},
		Auth: &dto.AuthRes{
			UserID: userModel.ID.String(),
			// AccessToken:  auth.AccessToken,
			// RefreshToken: auth.RefreshToken,
			// ExpiresIn:    uint64(auth.AccessTokenDuration),
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
