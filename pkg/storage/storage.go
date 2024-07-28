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
	Create(userID int, list gorestapiv2.TodoList) (int, error)
	GetAll(userID int) ([]gorestapiv2.TodoList, error)
	GetByID(userID, listID int) (gorestapiv2.TodoList, error)
	Delete(userID, listID int) error
	Update(userID, listID int, input gorestapiv2.UpdateListInput) error
}

type TodoItem interface {
	Create(listID int, item gorestapiv2.TodoItem) (int, error)
	GetAll(userID, listID int) ([]gorestapiv2.TodoItem, error)
	GetByID(userID, itemID int) (gorestapiv2.TodoItem, error)
	Delete(userID, itemID int) error
	Update(userID, itemID int, input gorestapiv2.UpdateItemInput) error
}

type Storage struct {
	Authorization
	TodoList
	TodoItem
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
