package config

import "strconv"

type Config struct {
	Port          string `json:"PORT"`
	ENV           string `json:"ENV"`
	LogLevel      string `json:"LOG_LEVEL"`
	AppName       string `json:"APP_NAME"`
	AuthSecretKey string `json:"AUTH_SECRET_KEY"`
	TokenTTL      string `json:"TOKEN_TTL"`
}

func ParseStringToUint64(str string) uint64 {
	var parsedUint64 uint64
	var err error

	// Use strconv.ParseUint to convert the string to uint64
	parsedUint64, err = strconv.ParseUint(str, 10, 64)
	if err != nil {
		// If parsing fails, return 0 and an error
		return 0
	}

	return parsedUint64
}
