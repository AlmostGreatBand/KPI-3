package db

import (
	"log"
	"testing"
)

func TestDatabaseConnection(t *testing.T) {
	connection := &Connection {
		DbName: "kpi3",
		Schema: "lab3",
		User: "balance_admin",
		Password: "3kpi",
		Host: "localhost",
		DisableSSL: true,
	}
	log.Println(connection.ConnectionUrl())
	if connection.ConnectionUrl() != "postgres://balance_admin:3kpi@localhost/kpi3?search_path=lab3&sslmode=disable" {
		t.Error("Unexpected connection string")
	}
}
