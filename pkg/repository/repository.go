package repository

import (
	"github.com/Nurt0re/chatik"

	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser( user chatik.User) (int, error)
	GetUser(email, password string) (chatik.User, error)
}

type Updater interface {
	UpdateUser(id int, input chatik.User) error
	DeleteUser(id int) error
	GetUser(id int) (chatik.User, error)
	GetAllUsers() ([]chatik.User, error)
}
type Repository struct {
	Authorization
	Updater
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Updater:	   NewUpdPostgres(db),
	}
}
