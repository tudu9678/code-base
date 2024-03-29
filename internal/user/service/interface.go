package service

import (
	"context"

	"myapp/internal/user/dto"
)

type User interface {
	Register(ctx context.Context, req *dto.CreateUserReq) (*dto.RegisterRes, error)
	Login(ctx context.Context, req *dto.LoginReq) (*dto.AuthRes, error)
}
