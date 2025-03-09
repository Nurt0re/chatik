package service

import (
	"github.com/Nurt0re/chatik"
	"github.com/Nurt0re/chatik/pkg/repository"
)

type Authorization interface {
	CreateUser(user chatik.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Updater interface {
	UpdateUser(id int, input chatik.User) error
	DeleteUser(id int) error
	GetUser(id int) (chatik.User, error)
	GetAllUsers() ([]chatik.User, error)
}
type Service struct {
	Authorization
	Updater
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Updater:       NewUpdater(repos.Updater),
	}
}
