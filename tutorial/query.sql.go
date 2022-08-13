// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package tutorial

import (
	"context"
	"database/sql"
)

const createAuthor = `-- name: CreateAuthor :execresult
INSERT INTO category (
    name
) VALUES (
    ?
)
`

func (q *Queries) CreateAuthor(ctx context.Context, name string) (sql.Result, error) {
	return q.db.ExecContext(ctx, createAuthor, name)
}

const deleteAuthor = `-- name: DeleteAuthor :exec
DELETE FROM category
WHERE id = ?
`

func (q *Queries) DeleteAuthor(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAuthor, id)
	return err
}

const getAuthor = `-- name: GetAuthor :one
SELECT id, name FROM category
WHERE id = ? LIMIT 1
`

func (q *Queries) GetAuthor(ctx context.Context, id int64) (Category, error) {
	row := q.db.QueryRowContext(ctx, getAuthor, id)
	var i Category
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const listAuthors = `-- name: ListAuthors :many
SELECT id, name FROM category
ORDER BY id desc
`

func (q *Queries) ListAuthors(ctx context.Context) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, listAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
