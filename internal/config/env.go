package config

import (
	"os"
	"strconv"
)

const (
	EnvServiceName    = "MS_KIT_SERVICE_NAME"
	EnvPort           = "MS_KIT_PORT"
	EnvJwkSetUri      = "MS_KIT_JWK_SET_URI"
	EnvCors           = "MS_KIT_CORS"
	EnvLogLevel       = "MS_KIT_LOG_LEVEL"
	EnvTracingEnabled = "MS_KIT_TRACING_ENABLED"
	EnvDev            = "MS_KIT_DEV"
)

func EnvOr(key, defaultValue string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return defaultValue
}

func BEnv(key string) bool {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		b, err := strconv.ParseBool(v)
		if err != nil {
			return false
		}
		return b
	}
	return false
}
