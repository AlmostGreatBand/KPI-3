package main

import (
	"database/sql"
	"flag"
	"github.com/AlmostGreatBand/KPI-3/server/db"
	"log"
	"net/http"
	"os"
	"os/signal"
)

var portNumber = flag.Int("p", 8080, "HTTP port number")

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

	if server, err := ComposeApiServer(HttpPortNumber(*portNumber)); err == nil {
		go func() {
			log.Println("Start server")
			if err := server.StartServer(); err != nil {
				if err == http.ErrServerClosed {
					log.Println("error: server has stopped")
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
