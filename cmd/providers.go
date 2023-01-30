package main

import (
	"github.com/landru29/gospot/internal/app"
	"github.com/landru29/gospot/internal/server"
	"github.com/landru29/gospot/internal/server/api"
	"github.com/landru29/gospot/internal/server/client"
	"github.com/sirupsen/logrus"
)

func providePublicAPI(log logrus.FieldLogger, conf *app.Config, toLaunch apiList) ([]server.Router, error) {
	out := []server.Router{}

	if toLaunch.Has("api") {
		srv, err := api.New(log, conf)
		if err != nil {
			return nil, err
		}

		out = append(out, srv)
	}

	if toLaunch.Has("client") {
		srv, err := client.New(log, conf)
		if err != nil {
			return nil, err
		}

		out = append(out, srv)
	}

	return out, nil
}

func provideLogger() logrus.FieldLogger {
	return logrus.New()
}
