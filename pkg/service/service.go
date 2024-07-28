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
	Create(userID int, list gorestapiv2.TodoList) (int, error)
	GetAll(userID int) ([]gorestapiv2.TodoList, error)
	GetByID(userID, listID int) (gorestapiv2.TodoList, error)
	Delete(userID, listID int) error
	Update(userID, listID int, input gorestapiv2.UpdateListInput) error
}

type TodoItem interface {
	Create(userID, listID int, item gorestapiv2.TodoItem) (int, error)
	GetAll(userID, listID int) ([]gorestapiv2.TodoItem, error)
	GetByID(userID, itemID int) (gorestapiv2.TodoItem, error)
	Delete(userID, itemID int) error
	Update(userID, itemID int, input gorestapiv2.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(strg *storage.Storage) *Service {
	return &Service{
		Authorization: NewAuthService(strg.Authorization),
		TodoList:      NewTodoListService(strg.TodoList),
		TodoItem:      NewTodoItemService(strg.TodoItem, strg.TodoList),
	}
}
