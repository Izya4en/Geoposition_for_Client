package config

import "os"

type Config struct {
	KafkaBroker string
}

func Load() *Config {
	return &Config{
		KafkaBroker: getEnv("KAFKA_BROKER", "localhost:9092"),
	}
}

func getEnv(key, def string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return def
}
