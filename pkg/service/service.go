package service

import (
	gorestapiv2 "github.com/DeMarDeXis/RESTV1"
	"github.com/DeMarDeXis/RESTV1/pkg/storage"
)

type Authorization interface {
	CreateUser(user gorestapiv2.User) (int, error)
	GenerateToken(username, password string) (string, error)
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

func NewService(strg *storage.Storage) *Service {
	return &Service{
		Authorization: NewAuthService(strg.Authorization),
	}
}
