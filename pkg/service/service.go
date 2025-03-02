package service

import (
	todo_app "todo-app"
	"todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo_app.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type ToDoList interface {
	Create(userId int, list todo_app.ToDoList) (int, error)
	GetAll(id int) ([]todo_app.ToDoList, error)
	GetById(userId int, listId int) (todo_app.ToDoList, error)
	Delete(userId int, listId int) error
	Update(userId, listId int, input todo_app.UpdateListInput) error
}

type ToDoItem interface {
	Create(userId, listId int, item todo_app.ToDoItem) (int, error)
	GetAll(userId, listId int) ([]todo_app.ToDoItem, error)
	GetById(userId, itemId int) (todo_app.ToDoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todo_app.UpdateItemInput) error
}

type Service struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		ToDoList:      NewTodoListService(repos.ToDoList),
		ToDoItem:      newTodoItemService(repos.ToDoItem, repos.ToDoList),
	}
}
