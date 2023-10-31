package sqlc_test

import (
	"context"
	"database/sql"
	_ "embed"
	"esol/database_access/sqlc/tutorial"
	"esol/database_access/test"
	"github.com/stretchr/testify/require"
	"testing"
)

//go:embed schema.sql
var schema string

func TestAndDemoSQLCUsage(t *testing.T) {
	db := test.SetupDBContainer(t)
	ctx := context.Background()
	initializeSchema(ctx, t, db)

	queries := tutorial.New(db)
	tx, err := db.Begin()
	require.NoError(t, err)
	queriesInTx := queries.WithTx(tx)
	author, err := queriesInTx.GetAuthor(ctx, 1)
	require.NoError(t, err)
	require.Equal(t, tutorial.Author{
		ID:   1,
		Name: "Jon Olav Fosse",
		Bio: sql.NullString{
			String: "Norwegian author, translator, and playwright.",
			Valid:  true,
		},
	}, author)
}

func initializeSchema(ctx context.Context, t *testing.T, db *sql.DB) {
	test.ExecContextRaw(ctx, t, db, schema, 0)

	test.ExecContextRaw(ctx, t, db,
		`
INSERT INTO authors(id,name,bio) VALUES (1, 'Jon Olav Fosse', 'Norwegian author, translator, and playwright.'),
                                        (2, 'Annie Thérèse Blanche Ernaux', 'French writer who was awarded the 2022 Nobel Prize in Literature "for the courage and clinical acuity with which she uncovers the roots, estrangements and collective restraints of personal memory".'),
                                        (3, 'Abdulrazak Gurnah','Tanzanian-born British novelist and academic.');
`, 3)
}
