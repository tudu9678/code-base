package http

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	de "myapp/core/initialize/transport/http"
	"myapp/core/utils"
	"myapp/internal/user/dto"
	"myapp/internal/user/service"
)

// Define a struct to hold registration requests
type registerRequest struct {
	ctx context.Context
	req *dto.CreateUserReq
	ch  chan<- *dto.RegisterRes // Channel to send the registration result
}

type userHandler struct {
	UserSvc         service.User          `inject:"user-svc"`
	Server          *echo.Echo            `inject:"server"`
	registerChannel chan *registerRequest // Channel to handle registration requests

}

func NewUserHandler(server *echo.Echo) *userHandler {
	// Create the registration channel
	registerChannel := make(chan *registerRequest)

	return &userHandler{registerChannel: registerChannel}
}

func (uh *userHandler) MapRouters() error {
	route := uh.Server.Group("")
	route.GET("/health-check", uh.Health())

	userGroup := route.Group("/v1/users")
	userGroup.POST("", uh.Register())
	userGroup.POST("/login", uh.Login())

	go uh.RegisterHandler()

	return nil
}

func (uh *userHandler) Health() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "i'm ok")
	}
}

func (uh *userHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			userReq *dto.CreateUserReq
			//ve      validator.ValidationErrors
		)

		err := c.Bind(&userReq)
		if err != nil {
			return c.JSON(utils.GetError(err).ID, de.EncodeJsonResponse(nil, nil, err))
		}

		// Create a channel to receive registration result
		ch := make(chan *dto.RegisterRes)

		// Create a registerRequest instance
		registerReq := &registerRequest{
			ctx: c.Request().Context(),
			req: userReq,
			ch:  ch,
		}

		// Send the registerRequest to the registration channel
		uh.registerChannel <- registerReq

		// Wait for the registration result with a timeout
		select {
		case res := <-ch:
			// Close the channel
			close(ch)
			// Check if there's an error in the registration result
			if res.Error != nil {
				return c.JSON(utils.GetError(res.Error).ID, de.EncodeJsonResponse(nil, nil, res.Error))
			}
			// Return the registration response
			return c.JSON(http.StatusCreated, res)

		case <-time.After(10 * time.Second): // Timeout after 5 seconds
			close(ch)
			return c.JSON(http.StatusInternalServerError, "Registration timeout")
		}
	}
}

func (uh *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req *dto.LoginReq
		err := c.Bind(&req)
		if err != nil {
			return c.JSON(utils.GetError(err).ID, de.EncodeJsonResponse(nil, nil, err))
		}

		ctx := c.Request().Context()
		resp, err := uh.UserSvc.Login(ctx, req)
		if err != nil {
			return c.JSON(utils.GetError(err).ID, de.EncodeJsonResponse(nil, nil, err))
		}

		return c.JSON(http.StatusOK, resp)
	}
}

func (uh *userHandler) RegisterHandler() {
	for req := range uh.registerChannel {
		res, err := uh.UserSvc.Register(req.ctx, req.req)
		if err != nil {
			// Send error response to the channel
			res = &dto.RegisterRes{Error: err}
		}
		req.ch <- res
	}
}
