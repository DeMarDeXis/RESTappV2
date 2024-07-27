package storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	DB_NAME = "postgres"
)

const (
	usersTable      = "users"
	todoListsTable  = "todo_lists"
	usersListsTable = "users_lists"
	todoItemsTable  = "todo_items"
	listsItemsTable = "lists_items"
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
	db, err := sqlx.Open(DB_NAME, buildStringConnections(cfg))
	fmt.Println(buildStringConnections(cfg))
	if err != nil {
		return nil, fmt.Errorf("failed open db %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping db %w", err)
	}

	return db, nil
}

func buildStringConnections(cfg Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.DBName,
		cfg.Password,
		cfg.SSLMode)
}
