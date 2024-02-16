package database

import (
	"database/sql"
	"log"
	"news_api/config"
	"os"
	"time"

	_ "github.com/lib/pq"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
)

var DB *reform.DB
var sqlDB *sql.DB

func Connect() error {
	config, err := config.LoadConfig("./")
	if err != nil {
		return err
	}

	sqlDB, err := sql.Open("postgres", config.DBConnectionString)

	if err != nil {
		return err
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	logger := log.New(os.Stderr, "SQL: ", log.Flags())

	DB = reform.NewDB(sqlDB, postgresql.Dialect, reform.NewPrintfLogger(logger.Printf))

	return nil
}

func Close() {
	if sqlDB != nil {
		sqlDB.Close()
	}
}
