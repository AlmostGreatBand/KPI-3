package main

import (
	"flag"
	"github.com/AlmostGreatBand/KPI-3/server/db"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"
	"os"
	"os/signal"
)

var portNumber = flag.Int("p", 8080, "HTTP port number")

func DatabaseConnection() (*pgxpool.Pool, error) {
	conn := &db.Connection{
		DbName: "kpi3",
		Schema: "lab3",
		User: "balance_admin",
		Password: "3kpi",
		Host: "localhost",
		DisableSSL: true,
	}
	return conn.Open()
}

func main() {
	flag.Parse()

	if server, err := ComposeApiServer(HttpPortNumber(*portNumber)); err == nil {
		go func() {
			log.Println("start server")
			if err := server.StartServer(); err != nil {
				if err == http.ErrServerClosed {
					log.Println("server has stopped")
				} else {
					log.Fatalf("error: cannot start server: %s", err)
				}
			}
		}()

		sigChannel := make(chan os.Signal, 1)
		signal.Notify(sigChannel, os.Interrupt)
		<-sigChannel

		if err := server.StopServer(); err != nil && err != http.ErrServerClosed {
			log.Printf("error: cannot stop server - %s", err)
		}
	} else {
		log.Fatalf("error: cannot initialize server - %s", err)
	}
}
