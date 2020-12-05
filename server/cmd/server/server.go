package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

type HttpPortNumber int

type BalancersApiServer struct {
	Port HttpPortNumber
	BalancersHandler http.HandlerFunc
	server *http.Server
}

func (s *BalancersApiServer) StartServer() error {
	if s.BalancersHandler == nil {
		return errors.New("error: http handler is not provided")
	}
	if s.Port == 0 {
		return errors.New("error: port is not provided")
	}

	handler := http.NewServeMux()
	handler.HandleFunc("/balancers", s.BalancersHandler)

	s.server = &http.Server{
		Addr: fmt.Sprintf(":%d", s.Port),
		Handler: handler,
	}

	return s.server.ListenAndServe()
}

func (s *BalancersApiServer) StopServer() error {
	if s.server == nil {
		return errors.New("error: server has already been stopped")
	}

	return s.server.Shutdown(context.Background())
}
