// Package api is the Backend server.
package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/landru29/gospot/internal/app"
	"github.com/landru29/gospot/internal/server"
	"github.com/sirupsen/logrus"
)

// Server is the API server.
type Server struct {
	log    logrus.FieldLogger
	conf   *app.Config
	server *http.Server
}

// New creates a new API.
func New(log logrus.FieldLogger, conf *app.Config) *Server {
	router := mux.NewRouter()

	srv := &http.Server{
		Handler:           router,
		Addr:              conf.ClientBind,
		ReadHeaderTimeout: time.Second * server.ReadHeaderTimeoutSeconds,
	}

	router.HandleFunc("/login", func(http.ResponseWriter, *http.Request) {}).Methods(http.MethodGet)
	router.HandleFunc("/callback", func(http.ResponseWriter, *http.Request) {}).Methods(http.MethodGet)
	router.HandleFunc("/albums", func(http.ResponseWriter, *http.Request) {}).Methods(http.MethodGet)

	return &Server{
		log:    log,
		conf:   conf,
		server: srv,
	}
}

// Start implements the Router interface.
func (s *Server) Start(ctx context.Context) {}
