package errors

type AuthErrorCode int

const (
	ErrorInvalidToken = iota + 1000
	ErrorUnauthenticated
	ErrorInvalidLoginToken
)
