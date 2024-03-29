package middleware

// import (
// 	"errors"
// 	"fmt"
// 	"path"

// 	"myapp/configs"
// 	"myapp/infrastructure/cache"
// 	"myapp/pkg/auth"
// 	"myapp/pkg/constants"
// 	ce "myapp/pkg/errors"

// 	"github.com/go-redis/redis/v8"
// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"
// 	echoJWT "github.com/labstack/echo-jwt"
// 	"github.com/labstack/gommon/log"
// 	"go.uber.org/zap"
// )

// const (
// 	UserIDKey      string = "user_id"
// 	UserEmailKey   string = "user_email"
// 	UserTypeKey    string = "user_type"
// 	ActivityLogKey string = "activity_log"
// 	UserSubIDKey   string = "sub_id"
// )

// const (
// 	UserStatusInActiveRedisKey = `is_user_inactive_%s`
// )

// type ApiMiddleware struct {
// 	conf    *configs.Config
// 	jwtAuth *auth.JWTAuth
// 	cache   cache.Cache
// }

// func NewApiMiddleware(conf *configs.Config, cache cache.Cache, jwtAuth *auth.JWTAuth, log *zap.Logger) *ApiMiddleware {
// 	return &ApiMiddleware{
// 		cache:   cache,
// 		conf:    conf,
// 		jwtAuth: jwtAuth,
// 	}
// }

// func (am *ApiMiddleware) CustomJWTAuth(skippers []string) echo.MiddlewareFunc {
// 	skipperFunc := middleware.DefaultSkipper
// 	if len(skippers) > 0 {
// 		skipperFunc = func(c echo.Context) bool {
// 			rURL, err := c.Request().URL.Parse(c.Path())
// 			if err != nil {
// 				log.Error("failed apply CustomJWTAuth ", zap.Any("error", err))
// 			}
// 			s := path.Base(rURL.Path)
// 			for _, v := range skippers {
// 				if s == v {
// 					return true
// 				}
// 			}

// 			return false
// 		}
// 	}

// 	jwtConf := echoJWT.Config{
// 		Skipper:        skipperFunc,
// 	}

// 	return echoJWT.WithConfig(jwtConf)
// }

// func (am *ApiMiddleware) parseTokenFunc(auth string, c echo.Context) (interface{}, error) {
// 	token, err := am.jwtAuth.ParseToken(auth)
// 	if err != nil {
// 		return nil, err
// 	}

// 	isAliveToken := ""

// 	err = am.cache.Get(c.Request().Context(), fmt.Sprintf(constants.INACTIVE_TTL, token.Sub), &isAliveToken)
// 	if err != nil && !errors.Is(err, redis.Nil) {
// 		return nil, fmt.Errorf("a.Cache.Get : %w", err)
// 	}

// 	if isAliveToken == "" {
// 		var validAfterTime int64
// 		err = am.cache.Get(c.Request().Context(), constants.TOKEN_VALID_AFTER+":"+token.Sub, &validAfterTime)
// 		if err != nil && !errors.Is(err, redis.Nil) {
// 			return nil, fmt.Errorf("a.Cache.Get : %w", err)
// 		}
// 		if int64(token.CreatedAt) < validAfterTime {
// 			return nil, ce.New(ce.ErrorUnauthorized, "invalid token", nil)
// 		}
// 	}

// 	c.Set(UserIDKey, token.Sub)
// 	c.Set(UserTypeKey, token.Type)
// 	c.Set(UserEmailKey, token.Email)
// 	c.Set(UserSubIDKey, token.SubID)

// 	return token, nil
// }

// func (am *ApiMiddleware) ResetIDFormMainToSubAccount(skippers []string) echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) (err error) {
// 			id := c.Get(UserSubIDKey).(string)
// 			if id == "" {
// 				return next(c)
// 			}
// 			rURL, err := c.Request().URL.Parse(c.Path())
// 			if err != nil {
// 				log.Error("failed apply CustomJWTAuth ", zap.Any("error", err))
// 			}
// 			s := path.Base(rURL.Path)
// 			skip := false
// 			for _, v := range skippers {
// 				if s == v {
// 					skip = true
// 				}
// 			}
// 			if !skip {
// 				c.Set(UserIDKey, id)
// 			}
// 			return next(c)
// 		}
// 	}
// }
