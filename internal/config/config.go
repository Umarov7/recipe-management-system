package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	ROUTER_ADDRESS string
	CACHE_SIZE     int
	CACHE_TIME     int // in seconds
	DB_HOST        string
	DB_PORT        string
	DB_USER        string
	DB_NAME        string
	DB_PASSWORD    string
}

func Load() *Config {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("error while getting current working directory: %v", err)
	}

	if err := godotenv.Load(path + "/.env"); err != nil {
		log.Printf("error while loading .env file: %v", err)
	}

	return &Config{
		ROUTER_ADDRESS: cast.ToString(coalesce("ROUTER_ADDRESS", "localhost:8080")),
		CACHE_SIZE:     cast.ToInt(coalesce("CACHE_SIZE", 1024)),
		CACHE_TIME:     cast.ToInt(coalesce("CACHE_TIME", 60)),
		DB_HOST:        cast.ToString(coalesce("DB_HOST", "localhost")),
		DB_PORT:        cast.ToString(coalesce("DB_PORT", "5432")),
		DB_USER:        cast.ToString(coalesce("DB_USER", "postgres")),
		DB_NAME:        cast.ToString(coalesce("DB_NAME", "postgres")),
		DB_PASSWORD:    cast.ToString(coalesce("DB_PASSWORD", "password")),
	}
}

func coalesce(key string, value interface{}) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return value
}
