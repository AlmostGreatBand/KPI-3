package db

import (
	"fmt"
	"testing"
)

func TestDatabaseConnection(t *testing.T) {
	connection := &Connection {
		DbName: "kpi3",
		User: "balance_admin",
		Password: "3kpi",
		Host: "localhost",
		DisableSSL: true,
	}
	if connection.ConnectionURL() != "lab3://balance_admin:3kpi@localhost/kpi3?sslmode=disable" {
		t.Error("Unexpected connection string")
	}
}
