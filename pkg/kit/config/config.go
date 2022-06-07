package config

var Kit *KitConfig

type KitConfig struct {
	ServiceName      string
	Port             string
	JwkSetUri        string
	CorsUrls         string
	LogLevel         string
	IsJsonLogging    bool
	IsTracingEnabled bool
	IsMetricsEnabled bool
	IsDevMode        bool
	JaegerEndpoint   string
}

func Apply(config KitConfig) {
	Kit = &config
}

func Default() KitConfig {
	return KitConfig{
		ServiceName:      "ms-kit-service",
		Port:             "8080",
		IsJsonLogging:    true,
		IsMetricsEnabled: true,
		IsTracingEnabled: true,
		JaegerEndpoint:   "http://jaeger:14268/api/traces",
	}
}

func Development() KitConfig {
	return KitConfig{
		ServiceName:      "ms-kit-service",
		Port:             "8080",
		IsJsonLogging:    false,
		IsTracingEnabled: false,
		IsDevMode:        true,
		IsMetricsEnabled: false,
	}
}
