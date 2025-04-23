package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	HarperDBHost     string
	HarperDBPort     int
	HarperDBUsername string
	HarperDBPassword string
	HarperDBSchema   string
	APIServerPort    string
}

func LoadConfig() *Config {
	cfg := &Config{
		HarperDBHost:     getEnv("HARPERDB_HOST", "localhost"),
		HarperDBPort:     getEnvAsInt("HARPERDB_PORT", 9925),
		HarperDBUsername: getEnv("HARPERDB_USERNAME", "admin"),
		HarperDBPassword: getEnv("HARPERDB_PASSWORD", "password"),
		HarperDBSchema:   getEnv("HARPERDB_SCHEMA", "kyc_data"),
		APIServerPort:    getEnv("API_SERVER_PORT", ":8080"),
	}
	fmt.Printf("Loaded Config: Host=%s, Port=%d, Username=%s, Password=%s, Schema=%s\n", cfg.HarperDBHost, cfg.HarperDBPort, cfg.HarperDBUsername, cfg.HarperDBPassword, cfg.HarperDBSchema)
	return cfg
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return fallback
}
