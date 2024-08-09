package db

import (
	"log"
	"task-management-system/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sqlx.DB

func Init() error {
	databaseURL := config.GetDatabaseURL()
	var err error
	DB, err = sqlx.Open("postgres", databaseURL)
	if err != nil {
		return err
	}
	if err = DB.Ping(); err != nil {
		return err
	}
	log.Println("Database connection established")
	return nil
}
