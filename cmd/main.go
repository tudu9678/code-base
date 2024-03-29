package main

import (
	"fmt"
	"net/http"

	"myapp/config"
	in "myapp/core/initialize"
	repo "myapp/internal/user/repository"
	svc "myapp/internal/user/service"
	handler "myapp/internal/user/transport/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// Register recover middleware
	e.Use(recoverMiddleware)

	e.GET("/health-check", func(c echo.Context) error {
		return c.String(http.StatusOK, "i'm ok")
	})

	var (
		internalCfg = &config.Config{}
		dbCfg       = &in.Postgres{}
	)
	if err := in.LoadConfiguration(internalCfg, dbCfg); err != nil {
		e.Logger.Error(err)
	}

	logger := in.NewLogRus(internalCfg.LogLevel, fmt.Sprintf("%s-logger", "auth-service"), internalCfg.ENV)

	in.RegisterIOCs("logger", logger)
	in.RegisterIOCs("handler", handler.NewUserHandler())
	//in.RegisterIOCs("middleware", md)
	in.RegisterIOCs("user-svc", svc.NewUserService())
	in.RegisterIOCs("user-repo", repo.NewUserRepo())

	if err := in.InitIOCs(); err != nil {
		e.Logger.Error(err)
	}

	// Start the server
	e.Start(internalCfg.Port)

}

func recoverMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if err := recover(); err != nil {
				// Log the panic
				e := echo.NewHTTPError(http.StatusInternalServerError)
				e.Message = "Internal server error"
				c.Logger().Error(err)
				return // Re-throw the error to stop further processing
			}
		}()
		return next(c)
	}
}
