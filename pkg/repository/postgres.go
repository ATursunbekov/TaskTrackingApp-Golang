package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	userTable      = "users"
	todoListTable  = "todo_lists"
	usersListTable = "users_lists"
	todoItemsTable = "todo_items"
	listItemsTable = "list_items"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	error := db.Ping()
	if error != nil {
		return nil, error
	}

	return db, nil
}
