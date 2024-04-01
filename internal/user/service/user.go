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
	errorKit "myapp/core/initialize/errors"
	"myapp/core/utils"
	"myapp/internal/user/dto"
	"myapp/internal/user/model"
	"myapp/internal/user/repository"
)

type user struct {
	//conf   config.Config       `inject:"configs"`
	Repo    repository.UserRepo `inject:"user-repo"`
	Logger  in.Logging          `inject:"logger"`
	JwtAuth auth.JWTAuth        `inject:"jwt-auth"`
	//RegisterChannel chan *registerRequest
}

func NewUserService() User {
	//registerChannel := make(chan *RegisterRequest)
	return &user{}
}

func (u *user) Register(ctx context.Context, req *dto.CreateUserReq) (*dto.RegisterRes, error) {

	dob, err := time.Parse(cCm.DateFm, req.Dob)
	if err != nil {
		u.Logger.Error(ctx, "[Register] Dob time.Parse err: %v", err)
		return nil, errorKit.BadRequest(utils.GetError(err).Message)
	}

	if err = u.checkUserExist(ctx, req.Email, req.PhoneNumber, req.UserName); err != nil {
		u.Logger.Error(ctx, "[Register] checkUserExist err: %v", err)
		return nil, err
	}

	userModel := &model.User{
		FullName:    req.FullName,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		UserName:    req.UserName,
		Password:    req.Password,
		Dob:         &dob,
	}

	if err := u.Repo.CreateUser(userModel); err != nil {
		return nil, errorKit.BadRequest(utils.GetError(err).Message)
	}
	u.Logger.Info(ctx, "[Register] : %v", userModel)

	tokenExpire := "3600"
	auth, err := u.JwtAuth.CreateToken(&auth.User{
		ID:       userModel.ID.String(),
		UserName: userModel.UserName,
		FullName: userModel.FullName,
	}, tokenExpire)
	if err != nil {
		u.Logger.Info(ctx, "[Login] username %v - CreateToken err %v", req.UserName, err)
		return nil, errorKit.InternalServerError(utils.GetError(err).Message)
	}

	// Wait for the registration result
	res := &dto.RegisterRes{
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
	}

	return res, nil
}

func (u *user) checkUserExist(ctx context.Context, email, phone, userName string) error {
	user, err := u.Repo.CheckUserExist(ctx, model.User{
		Email:       email,
		UserName:    userName,
		PhoneNumber: phone,
	})
	if err != nil {
		return errorKit.InternalServerError(utils.GetError(err).Message)
	}
	if user != nil {
		return errorKit.BadRequest("user exist")
	}

	return nil
}

func (u *user) Login(ctx context.Context, req *dto.LoginReq) (*dto.AuthRes, error) {
	u.Logger.Info(ctx, "[Login] username: %v", zap.Any("req.Username", req.Username))

	userModel, err := u.Repo.CheckUserExist(ctx, model.User{
		Email:       req.Username,
		PhoneNumber: req.Username,
		UserName:    req.Username,
	})
	if err != nil {
		return nil, errorKit.InternalServerError(utils.GetError(err).Message)
	}
	if userModel == nil {
		return nil, errorKit.Unauthorized(errorKit.ErrCommonUnauthorized)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(req.Password)); err != nil {
		return nil, errorKit.Unauthorized(errorKit.ErrCommonUnauthorized)
	}

	tokenExpire := "3600"
	auth, err := u.JwtAuth.CreateToken(&auth.User{
		ID:       userModel.ID.String(),
		UserName: userModel.UserName,
		FullName: userModel.FullName,
	}, tokenExpire)
	if err != nil {
		u.Logger.Info(ctx, "[Login] username %v - CreateToken err %v", userModel.UserName, err)
		return nil, errorKit.InternalServerError(utils.GetError(err).Message)
	}

	return &dto.AuthRes{
		UserID:       userModel.ID.String(),
		AccessToken:  auth.AccessToken,
		RefreshToken: auth.RefreshToken,
		ExpiresIn:    uint64(auth.AccessTokenDuration),
	}, nil
}
