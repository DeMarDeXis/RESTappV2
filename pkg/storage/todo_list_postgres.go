package storage

import (
	"fmt"
	"strings"

	gorestapiv2 "github.com/DeMarDeXis/RESTV1"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (t *TodoListPostgres) Create(userID int, list gorestapiv2.TodoList) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUsersListQuery, userID, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (t *TodoListPostgres) GetAll(userID int) ([]gorestapiv2.TodoList, error) {
	var lists []gorestapiv2.TodoList
	q := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.user_id = $1",
		todoListsTable, usersListsTable)
	if err := t.db.Select(&lists, q, userID); err != nil {
		return nil, err
	}

	return lists, nil
}

func (t *TodoListPostgres) GetByID(userID, listID int) (gorestapiv2.TodoList, error) {
	var list gorestapiv2.TodoList

	q := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description 
					FROM %s tl 
					INNER JOIN %s ul on tl.id = ul.list_id 
					WHERE ul.user_id = $1 AND ul.list_id = $2`,
		todoListsTable, usersListsTable)
	err := t.db.Get(&list, q, userID, listID)
	if err != nil {
		return gorestapiv2.TodoList{}, err
	}

	return list, nil
}

func (t *TodoListPostgres) Delete(userID, listID int) error {
	q := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id = $1 AND ul.list_id = $2",
		todoListsTable, usersListsTable)
	_, err := t.db.Exec(q, userID, listID)

	return err
}

func (t *TodoListPostgres) Update(userID, listID int, input gorestapiv2.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title = $%d", argID))
		args = append(args, *input.Title)
		argID++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description = $%d", argID))
		args = append(args, *input.Description)
		argID++
	}

	//title=$1
	//description=$1
	//title=$1, description=$2
	setQ := strings.Join(setValues, ", ")

	q := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		todoListsTable, setQ, usersListsTable, argID, argID+1)
	args = append(args, listID, userID)

	logrus.Debugf("Update query: %s", q)
	logrus.Debugf("Update args: %s", args)

	_, err := t.db.Exec(q, args...)

	return err
}
