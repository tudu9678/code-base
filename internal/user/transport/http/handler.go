package http

import (
	"net/http"

	"myapp/internal/user/dto"
	"myapp/internal/user/service"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	service service.User `inject:"user-svc"`
}

func NewUserHandler() *userHandler {
	return &userHandler{}
}

func (uh *userHandler) Register(c echo.Context) error {
	var (
		userReq *dto.CreateUserReq
		//ve      validator.ValidationErrors
	)

	err := c.Bind(&userReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// uv := NewUserValidator()
	// err = uv.validate.Struct(userReq)
	// if errors.As(err, &ve) {
	// 	for _, fe := range ve {
	// 		return r.Resp(http.StatusBadRequest, c, ce.New(uv.GetError(fe).CodeIdx, uv.GetError(fe).Message, nil))
	// 	}
	// }

	ctx := c.Request().Context()
	resp, err := uh.service.Register(ctx, userReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, resp)
}

func (uh *userHandler) Login(c echo.Context) error {
	var req *dto.LoginReq
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	resp, err := uh.service.Login(ctx, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, resp)
}
