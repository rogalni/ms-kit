package config

import "os"

const (
	EnvServiceName = "MS_KIT_SERVICE_NAME"
	EnvPort        = "MS_KIT_PORT"
	EnvJsonLog     = "MS_KIT_JSON_LOG"
	EnvLogLevel    = "MS_KIT_LOG_LEVEL"
	EnvJwkSetUri   = "MS_KIT_JWK_SET_URI"
)

func EnvOr(key, defaultValue string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return defaultValue
}
