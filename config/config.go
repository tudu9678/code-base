package config

type Config struct {
	Port     string `json:"PORT"`
	ENV      string `json:"ENV"`
	LogLevel string `json:"LOG_LEVEL"`
}
