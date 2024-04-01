package errors

import "fmt"

type AuthError struct {
	Code    AuthErrorCode `json:"code"`
	Message string        `json:"message"`
}

func (ae *AuthError) Error() string {
	return fmt.Sprintf("%s: %d", ae.Code, ae.Message)

}
