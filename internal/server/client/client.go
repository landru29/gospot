// Package client is the web client.
package client

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/landru29/gospot/internal/app"
	"github.com/landru29/gospot/internal/server"
	"github.com/sirupsen/logrus"
)

const (
	hashSize = 10
)

type injection struct {
	Hash  string
	Login string
}

// Server is the Client server.
type Server struct {
	log    logrus.FieldLogger
	conf   *app.Config
	server *http.Server
}

// New creates a new Client.
func New(log logrus.FieldLogger, conf *app.Config) *Server {
	router := mux.NewRouter()

	srv := &http.Server{
		Handler:           router,
		Addr:              conf.ClientBind,
		ReadHeaderTimeout: time.Second * server.ReadHeaderTimeoutSeconds,
	}

	router.HandleFunc(`/{file:.*}`, templateFile("./public/", injection{
		Hash:  randomHash(),
		Login: conf.Login,
	})).Methods(http.MethodGet)

	return &Server{
		log:    log,
		conf:   conf,
		server: srv,
	}
}

func randomHash() string {
	dataBytes := make([]byte, hashSize)
	_, _ = rand.Read(dataBytes)

	return base64.StdEncoding.EncodeToString(dataBytes)
}

// Start implements the Router interface.
func (s *Server) Start(ctx context.Context) {
	s.log.WithField("addr", s.conf.ClientBind).Info("launching client")

	go func() {
		_ = s.server.ListenAndServe()
	}()

	<-ctx.Done()

	_ = s.server.Shutdown(ctx)
}
