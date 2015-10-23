package test

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

func TestDB(t *testing.T) {
	db, err := sql.Open("postgres", "postgres://postgres@postgres/pg_test?sslmode=disable")
	if err != nil {
		println(err.Error())
		t.FailNow()
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		println(err.Error())
		t.FailNow()
	}
}
