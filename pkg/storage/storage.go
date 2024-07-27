package storage

import (
	gorestapiv2 "github.com/DeMarDeXis/RESTV1"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user gorestapiv2.User) (int, error)
	GetUser(username, password string) (gorestapiv2.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Storage struct {
	Authorization
	TodoList
	TodoItem
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		Authorization: NewAuthPostgres(db),
	}
}
