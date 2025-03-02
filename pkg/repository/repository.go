package repository

import (
	"github.com/jmoiron/sqlx"
	todo_app "todo-app"
)

type Authorization interface {
	CreateUser(user todo_app.User) (int, error)
	GetUser(username string, password string) (todo_app.User, error)
}

type ToDoList interface {
	Create(userId int, list todo_app.ToDoList) (int, error)
	GetAll(id int) ([]todo_app.ToDoList, error)
	GetById(userId int, listId int) (todo_app.ToDoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todo_app.UpdateListInput) error
}

type ToDoItem interface {
	Create(listId int, item todo_app.ToDoItem) (int, error)
	GetAll(userId, listId int) ([]todo_app.ToDoItem, error)
	GetById(userId, itemId int) (todo_app.ToDoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todo_app.UpdateItemInput) error
}

type Repository struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		ToDoList:      NewTodoListPostgres(db),
		ToDoItem:      NewTodoItemPostgres(db),
	}
}
