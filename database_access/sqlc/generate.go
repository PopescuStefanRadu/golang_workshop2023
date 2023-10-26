package sqlc

//go:generate rm -rf tutorial
//go:generate sqlc generate
// TODO fix issue with sqlc generate using ? instead of $1,$2,$3 by replacing them appropriately
