package repository

import (
	"github.com/Nurt0re/chatik"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser( user chatik.User) (int, error)
	GetUser(username, password string) (chatik.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
