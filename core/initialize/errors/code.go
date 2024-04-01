package errors

type AuthErrorCode int

const (
	ErrCommonBadRequest          = "BadRequest"
	ErrCommonUnauthorized        = "Unauthorized"
	ErrCommonForbidden           = "Forbidden"
	ErrCommonNotFound            = "NotFound"
	ErrCommonMethodNotAllowed    = "MethodNotAllowed"
	ErrCommonTimeout             = "TimeOut"
	ErrCommonConflict            = "Conflict"
	ErrCommonInternalServerError = "InternalServerError"

	ErrorInvalidToken = iota + 1000
	ErrorUnauthenticated
	ErrorInvalidLoginToken
)
