package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"task-management-system/internal/pagination"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseUrl string
	JwtSecret   string
	Pagination  pagination.PaginationConfig
}

func getPort() string {
	port := os.Getenv("HTTP_PORT")
	_, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("HTTP_PORT is not an int: %v\n", err)
	}

	return port
}

func getDatabaseUrl() string {
	dbUrl := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DATABASE_HOST"),
		"5432",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
	)

	return dbUrl
}

func getJwtSecret() string {
	return os.Getenv("JWT_SECRET")
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}

	cfg := &Config{
		Port:        getPort(),
		DatabaseUrl: getDatabaseUrl(),
		JwtSecret:   getJwtSecret(),
	}

	// Apply pagination configuration
	cfg.LoadPaginationConfig()

	return cfg
}
