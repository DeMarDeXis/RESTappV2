package storage

import (
	"fmt"
	"strings"

	gorestapiv2 "github.com/DeMarDeXis/RESTV1"
	"github.com/jmoiron/sqlx"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (t *TodoItemPostgres) Create(listID int, item gorestapiv2.TodoItem) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemID int
	createItemQ := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", todoItemsTable)

	row := tx.QueryRow(createItemQ, item.Title, item.Description)
	err = row.Scan(&itemID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemsQ := fmt.Sprintf("INSERT INTO %s (list_id, item_id) values ($1, $2)", listsItemsTable)
	_, err = tx.Exec(createListItemsQ, listID, itemID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemID, tx.Commit()
}

func (t *TodoItemPostgres) GetAll(userID, listID int) ([]gorestapiv2.TodoItem, error) {
	var items []gorestapiv2.TodoItem
	q := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li ON li.item_id = ti.id 
						INNER JOIN %s ul on ul.list_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2`,
		todoItemsTable, listsItemsTable, usersListsTable)

	if err := t.db.Select(&items, q, listID, userID); err != nil {
		return nil, err
	}

	return items, nil
}

func (t *TodoItemPostgres) GetByID(userID, itemID int) (gorestapiv2.TodoItem, error) {
	var item gorestapiv2.TodoItem
	q := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li ON li.item_id = ti.id
				INNER JOIN %s ul on ul.list_id = li.list_id WHERE ti.id = $1 AND ul.user_id = $2`,
		todoItemsTable, listsItemsTable, usersListsTable)
	if err := t.db.Get(&item, q, itemID, userID); err != nil {
		return item, err
	}

	return item, nil
}

func (t *TodoItemPostgres) Delete(userID, itemID int) error {
	q := fmt.Sprintf(`DELETE FROM %s ti USING %s li, %s ul
						WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND ti.id = $2`,
		todoItemsTable, listsItemsTable, usersListsTable)
	_, err := t.db.Exec(q, userID, itemID)
	return err
}

func (t *TodoItemPostgres) Update(userID, itemID int, input gorestapiv2.UpdateItemInput) error {
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

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done = $%d", argID))
		args = append(args, *input.Done)
		argID++
	}

	setQ := strings.Join(setValues, ", ")

	q := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s li, %s ul
						WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND ti.id = $%d`,
		todoItemsTable, setQ, listsItemsTable, usersListsTable, argID, argID+1)

	args = append(args, userID, itemID)

	_, err := t.db.Exec(q, args...)
	return err
}
