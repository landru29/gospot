// Package client is the web client.
package client

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
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
	Url   string
	Debug string
}

// Server is the Client server.
type Server struct {
	log    logrus.FieldLogger
	conf   *app.Config
	server *http.Server
}

// New creates a new Client.
func New(log logrus.FieldLogger, conf *app.Config) (*Server, error) {
	router := mux.NewRouter()

	srv := &http.Server{
		Handler:           router,
		Addr:              conf.ClientBind,
		ReadHeaderTimeout: time.Second * server.ReadHeaderTimeoutSeconds,
	}

	apiURL, err := url.Parse(conf.APIBaseURL)
	if err != nil {
		return nil, fmt.Errorf("cannot parse apiBaseUrl in configuration file: %w", err)
	}

	router.HandleFunc(`/{file:.*}`, templateFile("./client/build/", injection{
		Hash:  randomHash(),
		Url:   apiURL.String(),
		Debug: map[bool]string{true: "true", false: "false"}[conf.Debug],
	})).Methods(http.MethodGet)

	return &Server{
		log:    log,
		conf:   conf,
		server: srv,
	}, nil
}

func randomHash() string {
	dataBytes := make([]byte, hashSize)
	_, _ = rand.Read(dataBytes)

	return base64.StdEncoding.EncodeToString(dataBytes)
}

// Start implements the Router interface.
func (s *Server) Start(ctx context.Context) {
	s.log.WithField("addr", s.server.Addr).Info("launching client")

	go func() {
		_ = s.server.ListenAndServe()
	}()

	<-ctx.Done()

	_ = s.server.Shutdown(ctx)
}
