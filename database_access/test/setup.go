package test

import (
	"database/sql"
	"fmt"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/require"
	"testing"
)

func SetupDBContainer(t *testing.T) (db *sql.DB) {
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
	return
}
