package test

import (
	"database/sql"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

func TestDB(t *testing.T) {
	time.Sleep(5 * time.Second)
	db, err := sql.Open("postgres", "postgres://postgres@postgres/postgres?sslmode=disable")
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
