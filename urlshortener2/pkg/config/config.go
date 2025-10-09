package config

import (
	"os"
	"strconv"
)

type Config struct {
	ServerAddress string
	DatabaseURL   string
	RedisURL      string
	RateLimiter   int
	BaseURL       string
}

func Load() *Config {
	return &Config{
		ServerAddress: getEnv("SERVER_ADDR", ":8000"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres"),
		RedisURL: getEnv("REDIS_URL","redis"),
		RateLimiter: getEnvInt("RATE_LIMIT",100),
		BaseURL: getEnv("BASE_URL","localhost:8000"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
