package repository

import (
	"fmt"

	"github.com/Nurt0re/chatik"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	usersTable = "users"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*gorm.DB, error) {

	dsn := fmt.Sprintf("host = %s port = %s user = %s dbname = %s password = %s sslmode = %s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	db.Migrator().DropTable(&chatik.User{})
	err = db.AutoMigrate(&chatik.User{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {

		return nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, err

	}
	return db, nil
}
