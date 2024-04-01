package utils

import (
	errorkit "myapp/core/initialize/errors"
)

func GetError(err error) *errorkit.Error {
	return errorkit.Parse(err.Error())
}
