package repository

import (
	"github.com/MrClean-code/dayling"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user dayling.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
