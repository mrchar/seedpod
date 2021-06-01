package server

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	server *http.Server
	router *http.Handler
}

func Default() *Server {
	router := newRouter()
	server, err := New(router)
	if err != nil {
		logrus.Fatal(err)
	}

	return server
}

func New(router http.Handler) (*Server, error) {
	var config Config
	if err := config.Load(); err != nil {
		return nil, err
	}

	return &Server{
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", config.Address, config.Port),
			Handler: router,
		},
	}, nil
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}
