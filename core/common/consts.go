package common

import "time"

const (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvUat   = "uat"
	EnvProd  = "prod"

	APIsTimeOut = time.Second * 20
	// DB log mode
	DBLogModeSilent = "silent"
	DBLogModeError  = "error"
	DBLogModeWarn   = "warn"
	DBLogModeInfo   = "info"
)
