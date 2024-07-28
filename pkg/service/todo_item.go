package service

import (
	gorestapiv2 "github.com/DeMarDeXis/RESTV1"
	"github.com/DeMarDeXis/RESTV1/pkg/storage"
)

type TodoItemService struct {
	storage  storage.TodoItem
	listStrg storage.TodoList
}

func NewTodoItemService(storage storage.TodoItem, listStrg storage.TodoList) *TodoItemService {
	return &TodoItemService{
		storage:  storage,
		listStrg: listStrg,
	}
}

func (s *TodoItemService) Create(userID, listID int, item gorestapiv2.TodoItem) (int, error) {
	_, err := s.listStrg.GetByID(userID, listID)
	if err != nil {
		return 0, err
	}

	return s.storage.Create(listID, item)
}

func (s *TodoItemService) GetAll(userID, listID int) ([]gorestapiv2.TodoItem, error) {
	return s.storage.GetAll(userID, listID)
}

func (s *TodoItemService) GetByID(userID, itemID int) (gorestapiv2.TodoItem, error) {
	return s.storage.GetByID(userID, itemID)
}

func (s *TodoItemService) Delete(userID, itemID int) error {
	return s.storage.Delete(userID, itemID)
}

func (s *TodoItemService) Update(userID, itemID int, input gorestapiv2.UpdateItemInput) error {
	return s.storage.Update(userID, itemID, input)
}
