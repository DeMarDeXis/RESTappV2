package service

import (
	gorestapiv2 "github.com/DeMarDeXis/RESTV1"
	"github.com/DeMarDeXis/RESTV1/pkg/storage"
)

type TodoListService struct {
	storage storage.TodoList
}

func NewTodoListService(storage storage.TodoList) *TodoListService {
	return &TodoListService{storage: storage}
}

func (s *TodoListService) Create(userID int, list gorestapiv2.TodoList) (int, error) {
	return s.storage.Create(userID, list)
}

func (s *TodoListService) GetAll(userID int) ([]gorestapiv2.TodoList, error) {
	return s.storage.GetAll(userID)
}

func (s *TodoListService) GetByID(userID, listID int) (gorestapiv2.TodoList, error) {
	return s.storage.GetByID(userID, listID)
}

func (s *TodoListService) Delete(userID, listID int) error {
	return s.storage.Delete(userID, listID)
}

func (s *TodoListService) Update(userID, listID int, input gorestapiv2.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.storage.Update(userID, listID, input)
}
