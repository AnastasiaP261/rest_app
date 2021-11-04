package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) { // будет запущена один раз
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "postgres://postgres:test@localhost:5437/postgres?sslmode=disable"
	}

	os.Exit(m.Run())
}
