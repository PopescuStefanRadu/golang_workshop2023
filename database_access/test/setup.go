package test

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib" // drivername "pgx" atunci cand dati sql.Open("pgx"....)
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/require"
	"testing"
)

func SetupDBContainer(t *testing.T) *sql.DB {
	// READ MORE: TestMain
	pool, err := dockertest.NewPool("")
	require.NoError(t, err)
	require.NoError(t, pool.Client.Ping())

	resource, err := pool.Run("postgres", "16", []string{
		"POSTGRES_USER=superuser",
		"POSTGRES_PASSWORD=pass",
		"POSTGRES_DB=test",
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, pool.Purge(resource))
	})

	var db *sql.DB
	err = pool.Retry(func() error {
		var err error
		db, err = sql.Open("pgx", fmt.Sprintf("postgres://superuser:pass@%s/test", resource.GetHostPort("5432/tcp")))
		if err != nil {
			return err
		}
		return db.Ping()
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, db.Close())
	})
	return db
}

func ExecContextRaw(ctx context.Context, t *testing.T, db *sql.DB, stmt string, expectedRowsAffected int64) {
	res, err := db.ExecContext(ctx, stmt)
	require.NoError(t, err)
	rowsAff, err := res.RowsAffected()
	require.NoError(t, err)
	require.Equal(t, expectedRowsAffected, rowsAff)
}
