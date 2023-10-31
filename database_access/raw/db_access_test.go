package raw_test

import (
	"context"
	"database/sql"
	"esol/database_access/test"
	"esol/must"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type Car struct {
	Id              string
	Model           string
	FabricationDate time.Time
}

func TestDatabaseAccess(t *testing.T) {
	var db *sql.DB
	db = test.SetupDBContainer(t)
	ctx := context.Background()
	initializeSchema(ctx, t, db)

	// READ MORE: https://go.dev/doc/database/
	tx, err := db.Begin()
	require.NoError(t, err)

	q, err := tx.PrepareContext(
		ctx,
		"SELECT id, model, fabrication_date FROM cars WHERE fabrication_date > $1 ORDER BY fabrication_date DESC LIMIT 10 OFFSET $2",
	)
	require.NoError(t, err)
	t.Cleanup(func() {
		_ = q.Close()
	})

	rows, err := q.Query("2023-05-01", 0)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, rows.Close())
	})

	var cars []Car

	for rows.Next() {
		var car Car
		require.NoError(t, rows.Scan(&car.Id, &car.Model, &car.FabricationDate))
		cars = append(cars, car)
	}
	require.NoError(t, rows.Err())

	require.Equal(t, 2, len(cars))
	eqCar(t, Car{
		Model:           "Mazda CX-5 2022",
		FabricationDate: must.ParseDate("2023-05-21"),
	}, cars[0])
	eqCar(t, Car{
		Model:           "Toyota Yaris XP210",
		FabricationDate: must.ParseDate("2023-05-02"),
	}, cars[1])
}

func initializeSchema(ctx context.Context, t *testing.T, db *sql.DB) {
	test.ExecContextRaw(ctx, t, db, `
CREATE TABLE cars
(
    id               uuid PRIMARY KEY,
    model            text,
    fabrication_date date
)`, 0)

	test.ExecContextRaw(ctx, t, db,
		`
INSERT INTO cars VALUES (gen_random_uuid(), 'Honda Civic 2021', date '2021-01-01'),
                        (gen_random_uuid(), 'Toyota Yaris XP210', date '2023-04-23'),
                        (gen_random_uuid(), 'Mazda CX-5 2022', date '2023-05-21'),
                        (gen_random_uuid(), 'Toyota Yaris XP210', date '2023-05-02')
`, 4)
}

func eqCar(t *testing.T, c1, c2 Car) {
	c1.Id, c2.Id = "", ""
	require.Equal(t, c1, c2)
}
