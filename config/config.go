package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"task-management-system/internal/logger"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	Port        string
	DatabaseUrl string
	JwtSecret   string
	Logger      *logger.Logger
	Pagination  PaginationConfig
	DB          *sqlx.DB
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

	// Initialize logger
	development := os.Getenv("ENVIRONMENT") == "development"

	loggers, err := logger.Init(development)
	if err != nil {
		log.Fatalf("Error initializing logger: %v\n", err)
	}

	// Initialize the database
	dbConn, err := sqlx.Connect("postgres", getDatabaseUrl())
	if err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
	}

	cfg := &Config{
		Port:        getPort(),
		DatabaseUrl: getDatabaseUrl(),
		JwtSecret:   getJwtSecret(),
		Logger:      loggers,
		DB:          dbConn,
	}
	cfg.Logger = loggers

	// Apply pagination configuration
	cfg.LoadPaginationConfig()

	return cfg
}
