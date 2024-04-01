package auth

import (
	"time"

	"myapp/core/initialize/auth/errors"
	"myapp/core/initialize/helpers"

	"github.com/google/uuid"

	"github.com/golang-jwt/jwt"
)

type JWTAuth struct {
	iss        string
	signingKey string
	tokenTTL   uint64
}

func NewJWTAuth(iss, key string, tokenTTL uint64) *JWTAuth {
	return &JWTAuth{
		iss:        iss,
		signingKey: key,
		tokenTTL:   tokenTTL,
	}
}

func (jwtAuth *JWTAuth) CreateLoginToken(user IUser, expiredIn string) (*Auth, error) {
	mySigningKey := []byte(jwtAuth.signingKey)

	// generate access_token
	tokenTTL := time.Duration(600)

	if expiredIn != "" {
		tokenTTL = helpers.StringToTimeDuration(expiredIn)
	}
	duration := tokenTTL * time.Second
	expirationTime := time.Now().Add(duration)

	claims := jwt.MapClaims{
		"userName":  user.GetUserName(),
		"exp":       expirationTime.Unix(),
		"iss":       jwtAuth.iss,
		"createdAt": time.Now().Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenStr, err := accessToken.SignedString(mySigningKey)
	if err != nil {
		return nil, err
	}

	return &Auth{
		AccessToken:         accessTokenStr,
		AccessTokenDuration: int64(duration.Seconds()),
	}, nil
}

func (jwtAuth *JWTAuth) CreateToken(user *User, expiredIn string) (*Auth, error) {
	mySigningKey := []byte(jwtAuth.signingKey)

	// generate access_token
	tokenTTL := time.Duration(jwtAuth.tokenTTL)

	if expiredIn != "" {
		tokenTTL = helpers.StringToTimeDuration(expiredIn)
	}
	duration := tokenTTL * time.Second
	expirationTime := time.Now().Add(duration)
	sub := user.GetID()
	subID := ""
	claims := jwt.MapClaims{
		"id":        uuid.New().String(),
		"userName":  user.GetUserName(),
		"sub":       sub,
		"exp":       expirationTime.Unix(),
		"iss":       jwtAuth.iss,
		"createdAt": time.Now().Unix(),
		"subId":     subID,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenStr, err := accessToken.SignedString(mySigningKey)
	if err != nil {
		return nil, err
	}

	// generate refresh_token
	claims["id"] = uuid.New().String()
	claims["exp"] = time.Now().Add((time.Duration(jwtAuth.tokenTTL) * time.Second)).Unix()
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshTokenStr, err := refreshToken.SignedString(mySigningKey)
	if err != nil {
		return nil, err
	}

	return &Auth{
		UserId:               user.GetID(),
		AccessToken:          accessTokenStr,
		RefreshToken:         refreshTokenStr,
		AccessTokenDuration:  int64(duration.Seconds()),
		RefreshTokenDuration: int64(duration.Seconds() * 2),
	}, nil
}

// func (jwtAuth *JWTAuth) RefreshToken(refreshTokenStr string, expiredIn string) (*Auth, error) {
// 	refreshToken, err := jwt.Parse(refreshTokenStr, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, &errors.AuthError{Code: errors.ErrorInvalidToken}
// 		}

// 		return []byte(jwtAuth.signingKey), nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	// claims, ok := refreshToken.Claims.(jwt.MapClaims)
// 	// if !ok || !refreshToken.Valid {
// 	// 	return nil, &errors.AuthError{Code: errors.ErrorInvalidToken}
// 	// }

// 	// sub, ok := claims["sub"].(string)
// 	// if !ok {
// 	// 	return nil, &errors.AuthError{Code: errors.ErrorInvalidToken}
// 	// }
// 	// userEmail, ok := claims["email"].(string)
// 	// if !ok {
// 	// 	return nil, &errors.AuthError{Code: errors.ErrorInvalidToken}
// 	// }

// 	//Create new tokens
// 	return jwtAuth.CreateToken(nil, expiredIn)
// }

// ParseToken
func (jwtAuth *JWTAuth) ParseToken(tokenStr string) (*TokenInfo, error) {
	claims, err := jwtAuth.ClaimToken(tokenStr)
	if err != nil {
		return nil, err
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return nil, &errors.AuthError{Code: errors.ErrorInvalidToken}
	}

	userName, ok := claims["userName"].(string)
	if !ok {
		return nil, &errors.AuthError{Code: errors.ErrorInvalidToken}
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return nil, &errors.AuthError{Code: errors.ErrorInvalidToken}
	}

	createdAt, ok := claims["createdAt"].(float64)
	if !ok {
		return nil, &errors.AuthError{Code: errors.ErrorInvalidToken}
	}
	subId, ok := claims["subId"].(string)
	if !ok {
		return nil, &errors.AuthError{Code: errors.ErrorInvalidToken}
	}
	return &TokenInfo{
		Iss:       jwtAuth.iss,
		Sub:       sub,
		UserName:  userName,
		Exp:       int64(exp),
		CreatedAt: int64(createdAt),
		SubID:     subId,
	}, nil
}

// ClaimToken
func (jwtAuth *JWTAuth) ClaimToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, &errors.AuthError{Code: errors.ErrorInvalidToken}
		}

		return []byte(jwtAuth.signingKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, &errors.AuthError{Code: errors.ErrorInvalidToken}
	}
	if err := claims.Valid(); err != nil {
		return nil, &errors.AuthError{Code: errors.ErrorInvalidToken}
	}
	if !claims.VerifyIssuer(jwtAuth.iss, true) {
		return nil, &errors.AuthError{Code: errors.ErrorUnauthenticated}
	}

	return claims, nil
}

// ParseLoginToken
func (jwtAuth *JWTAuth) ParseLoginToken(tokenStr string) (*TokenInfo, error) {
	claims, err := jwtAuth.ClaimToken(tokenStr)
	if err != nil {
		return nil, err
	}

	userName, ok := claims["userName"].(string)
	if !ok {
		return nil, &errors.AuthError{Code: errors.ErrorInvalidLoginToken}
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return nil, &errors.AuthError{Code: errors.ErrorInvalidLoginToken}
	}

	createdAt, ok := claims["createdAt"].(float64)
	if !ok {
		return nil, &errors.AuthError{Code: errors.ErrorInvalidLoginToken}
	}
	return &TokenInfo{
		Iss:       jwtAuth.iss,
		UserName:  userName,
		Exp:       int64(exp),
		CreatedAt: int64(createdAt),
	}, nil
}
