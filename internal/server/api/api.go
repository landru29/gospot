// Package api is the Backend server.
package api

import (
	"context"
	"net/http"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	"github.com/landru29/gospot/internal/app"
	"github.com/landru29/gospot/internal/oauth"
	"github.com/landru29/gospot/internal/server"
	"github.com/sirupsen/logrus"
)

// Server is the API server.
type Server struct {
	log       logrus.FieldLogger
	conf      *app.Config
	server    *http.Server
	oauth     *oauth.Client
	templates *template.Template
}

// New creates a new API.
func New(log logrus.FieldLogger, conf *app.Config) (*Server, error) {
	router := mux.NewRouter()

	auth, err := oauth.New(conf)
	if err != nil {
		return nil, err
	}

	tmpl, err := template.ParseFiles("./public/error.tmpl.html")
	if err != nil {
		return nil, err
	}

	srv := &http.Server{
		Handler:           router,
		Addr:              conf.APIBind,
		ReadHeaderTimeout: time.Second * server.ReadHeaderTimeoutSeconds,
	}

	output := &Server{
		log:       log,
		conf:      conf,
		server:    srv,
		oauth:     auth,
		templates: tmpl,
	}

	router.HandleFunc("/login", output.processLogin).Methods(http.MethodGet)
	router.HandleFunc("/callback", output.processCallback).Methods(http.MethodGet)
	router.HandleFunc("/albums", func(http.ResponseWriter, *http.Request) {}).Methods(http.MethodGet)

	return output, nil
}

// Start implements the Router interface.
func (s *Server) Start(ctx context.Context) {
	s.log.WithField("addr", s.server.Addr).Info("launching api")

	go func() {
		_ = s.server.ListenAndServe()
	}()

	<-ctx.Done()

	_ = s.server.Shutdown(ctx)
}
