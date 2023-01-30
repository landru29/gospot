package main

import (
	"github.com/landru29/gospot/internal/app"
	"github.com/landru29/gospot/internal/server"
	"github.com/landru29/gospot/internal/server/api"
	"github.com/landru29/gospot/internal/server/client"
	"github.com/sirupsen/logrus"
)

func providePublicAPI(log logrus.FieldLogger, conf *app.Config, toLaunch apiList) []server.Router {
	out := []server.Router{}

	if toLaunch.Has("api") {
		out = append(out, api.New(log, conf))
	}

	if toLaunch.Has("client") {
		out = append(out, client.New(log, conf))
	}

	return out
}

func provideLogger() logrus.FieldLogger {
	return logrus.New()
}
