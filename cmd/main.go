package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
	gLog "gorm.io/gorm/logger"

	"myapp/config"
	in "myapp/core/initialize"
	"myapp/core/initialize/auth"
	repo "myapp/internal/user/repository"
	svc "myapp/internal/user/service"
	handler "myapp/internal/user/transport/http"
)

func main() {
	e := echo.New()
	// Register recover middleware
	e.Use(recoverMiddleware)
	e.Use(middleware.Logger())

	var (
		internalCfg = &config.Config{}
		dbMasterCfg = &in.Postgres{}
	)
	if err := in.LoadConfiguration(internalCfg, dbMasterCfg); err != nil {
		e.Logger.Error(err)
	}
	handler := handler.NewUserHandler(e)

	conn := in.NewConnector(context.Background(), internalCfg.AppName, internalCfg.ENV)
	dbMasterCfg.GormConfig = gorm.Config{
		Logger:      gLog.Default.LogMode(gLog.Error),
		PrepareStmt: false,
	}
	logger := in.NewLogRus(internalCfg.LogLevel, fmt.Sprintf("%s-logger", "auth-service"), internalCfg.ENV)

	in.RegisterIOCs("logger", logger)
	in.RegisterIOCs("server", e)
	in.RegisterIOCs("jwt-auth", auth.NewJWTAuth(internalCfg.AppName, internalCfg.AuthSecretKey, config.ParseStringToUint64(internalCfg.TokenTTL)))
	in.RegisterIOCs("db-master", conn.InitPostgres(dbMasterCfg))
	in.RegisterIOCs("user-repo", repo.NewUserRepo())
	in.RegisterIOCs("user-svc", svc.NewUserService())
	in.RegisterIOCs("user-handler", handler)

	//in.RegisterIOCs("middleware", md)

	if err := in.InitIOCs(); err != nil {
		e.Logger.Error(err)
	}

	// Start the server
	if err := handler.MapRouters(); err != nil {
		log.Fatalf("MapHandlers Error: %v", err)
	}
	if err := e.Start(fmt.Sprintf(":%s", internalCfg.Port)); err != nil {
		log.Fatalf("Running HTTP server: %v", err)
	}

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
