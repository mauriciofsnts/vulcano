package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mauriciofsnts/bot/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New(cfg config.Config) (*gorm.DB, error) {
	if cfg.DB.Type == config.DatabasePostgres {
		return NewPostgresConnection(cfg.DB.Postgres.Host, cfg.DB.Postgres.Port, cfg.DB.Postgres.User, cfg.DB.Postgres.Password, cfg.DB.Postgres.Database)
	}

	return NewSqliteConnection(cfg.DB.Sqlite.Path)
}

func NewPostgresConnection(host string, port int, user, password, database string) (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		host, user, password, database, port)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,  // Slow SQL threshold
			LogLevel:                  logger.Error, // Log level
			IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,         // Don't include params in the SQL log
			Colorful:                  true,         // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewSqliteConnection(path string) (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,  // Slow SQL threshold
			LogLevel:                  logger.Error, // Log level
			IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,         // Don't include params in the SQL log
			Colorful:                  true,         // Disable color
		},
	)

	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{Logger: newLogger})

	if err != nil {
		return nil, err
	}

	return db, nil
}
