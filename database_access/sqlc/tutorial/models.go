// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package tutorial

import (
	"database/sql"
)

type Author struct {
	ID   int32
	Name string
	Bio  sql.NullString
}

type Book struct {
	ID       int32
	AuthorID int32
	Title    sql.NullString
}
