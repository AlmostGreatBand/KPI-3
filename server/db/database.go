package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/url"
)

type Connection struct {
	DbName string
	Schema string
	User, Password string
	Host string
	DisableSSL bool
}

func (c *Connection) ConnectionUrl() string {
	dbUrl := &url.URL{
		Scheme: "postgres",
		Host:   c.Host,
		User:   url.UserPassword(c.User, c.Password),
		Path:   c.DbName,
	}
	query := url.Values{}
	if c.DisableSSL {
		query.Set("sslmode", "disable")
	}
	query.Set("search_path", c.Schema)

	dbUrl.RawQuery = query.Encode()
	log.Println(dbUrl.String())
	return dbUrl.String()
}

func (c *Connection) Open() (*pgxpool.Pool, error) {
	db, err := pgxpool.Connect(context.Background(), c.ConnectionUrl())
	return db, err
}
