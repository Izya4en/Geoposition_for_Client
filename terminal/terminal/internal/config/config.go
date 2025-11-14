package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	DSN           string
	Port          string
	AdminToken    string
	NumTerminals  int
	SimIntervalSec int
}

func Load() *Config {
	numTerm, _ := strconv.Atoi(getEnv("NUM_TERMINALS", "5"))
	simInt, _ := strconv.Atoi(getEnv("SIM_INTERVAL_SEC", "10"))

	return &Config{
		DSN:           getEnv("DATABASE_DSN", "host=localhost user=postgres password=pass dbname=terminaldb port=5432 sslmode=disable"),
		Port:          getEnv("PORT", "8080"),
		AdminToken:    getEnv("ADMIN_TOKEN", "secret"),
		NumTerminals:  numTerm,
		SimIntervalSec: simInt,
	}
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
