package service

import (
	todo_app "todo-app"
	"todo-app/pkg/repository"
)

type ToDoItemService struct {
	repo     repository.ToDoItem
	listRepo repository.ToDoList
}

func newTodoItemService(repo repository.ToDoItem, listRepo repository.ToDoList) *ToDoItemService {
	return &ToDoItemService{repo: repo, listRepo: listRepo}
}

func (s *ToDoItemService) Create(userId, listId int, item todo_app.ToDoItem) (int, error) {
	_, error := s.listRepo.GetById(userId, listId)
	if error != nil {
		return 0, error
	}

	return s.repo.Create(listId, item)
}

func (s *ToDoItemService) GetAll(userId, listId int) ([]todo_app.ToDoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *ToDoItemService) GetById(userId, itemId int) (todo_app.ToDoItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *ToDoItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *ToDoItemService) Update(userId, itemId int, input todo_app.UpdateItemInput) error {
	return s.repo.Update(userId, itemId, input)
}
