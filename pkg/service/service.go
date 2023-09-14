package service

import (
	"github.com/MrClean-code/dayling"
	"github.com/MrClean-code/dayling/pkg/repository"
)

type Authorization interface {
	CreateUser(user dayling.User) (int, error)
	GenerateToken(name, username string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
