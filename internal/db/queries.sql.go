// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: queries.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO users (id, email, password) VALUES($1,$2,$3)
`

type CreateUserParams struct {
	ID       int64
	Email    string
	Password string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser, arg.ID, arg.Email, arg.Password)
	return err
}

const listUsers = `-- name: ListUsers :many
SELECT id, email, password, created_at, updated_at FROM users ORDER BY id
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Password,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
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
