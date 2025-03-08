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

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
