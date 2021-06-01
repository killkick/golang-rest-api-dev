package store_test

import (
	"os"
	"testing"
)

var (
	databaseUrl string
)

func testMain(m *testing.M) {
	databaseUrl = os.Getenv("DATABASE_URL")
	if databaseUrl == " " {
		databaseUrl = "host=localhost dbname=restapi_dev sslmode=disable password=password"
	}

	os.Exit(m.Run())
}
