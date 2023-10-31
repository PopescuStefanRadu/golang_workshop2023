// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: query.question.sql

package tutorial

import (
	"context"
	"database/sql"
)

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO authors (
    name, bio
) VALUES (
             $1, $2
         )
    RETURNING id, name, bio
`

type CreateAuthorParams struct {
	Name string
	Bio  sql.NullString
}

func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) (Author, error) {
	row := q.db.QueryRowContext(ctx, createAuthor, arg.Name, arg.Bio)
	var i Author
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}

const deleteAuthor = `-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = $1
`

func (q *Queries) DeleteAuthor(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteAuthor, id)
	return err
}

const findBooksForAuthor = `-- name: FindBooksForAuthor :many
SELECT authors.id, authors.name, authors.bio, books.id, books.author_id, books.title FROM authors
LEFT JOIN books ON authors.id = books.author_id
WHERE authors.id = $1
`

type FindBooksForAuthorRow struct {
	Author Author
	Book   Book
}

func (q *Queries) FindBooksForAuthor(ctx context.Context, id int32) ([]FindBooksForAuthorRow, error) {
	rows, err := q.db.QueryContext(ctx, findBooksForAuthor, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindBooksForAuthorRow
	for rows.Next() {
		var i FindBooksForAuthorRow
		if err := rows.Scan(
			&i.Author.ID,
			&i.Author.Name,
			&i.Author.Bio,
			&i.Book.ID,
			&i.Book.AuthorID,
			&i.Book.Title,
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

const getAuthor = `-- name: GetAuthor :one
SELECT id, name, bio FROM authors
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAuthor(ctx context.Context, id int32) (Author, error) {
	row := q.db.QueryRowContext(ctx, getAuthor, id)
	var i Author
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}

const listAuthors = `-- name: ListAuthors :many
SELECT id, name, bio FROM authors
ORDER BY name
`

func (q *Queries) ListAuthors(ctx context.Context) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, listAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(&i.ID, &i.Name, &i.Bio); err != nil {
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

const updateAuthor = `-- name: UpdateAuthor :exec
UPDATE authors
set name = $1,
    bio = $2
WHERE id = $3
`

type UpdateAuthorParams struct {
	Name string
	Bio  sql.NullString
	ID   int32
}

func (q *Queries) UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) error {
	_, err := q.db.ExecContext(ctx, updateAuthor, arg.Name, arg.Bio, arg.ID)
	return err
}
