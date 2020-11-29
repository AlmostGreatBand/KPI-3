package server

import (
	"database/sql"
	"flag"
	"github.com/AlmostGreatBand/KPI-3/server/db"
)

func DatabaseConnection() (*sql.DB, error) {
	conn := &db.Connection{
		DbName:     "chat-example",
		User:       "roman",
		Host:       "localhost",
		DisableSSL: true,
	}
	return conn.Open()
}

func main() {

	flag.Parse()



}