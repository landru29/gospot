package main

import (
	"github.com/landru29/gospot/internal/app"
	"github.com/landru29/gospot/internal/music"
	"github.com/landru29/gospot/internal/music/spotify"
	"github.com/landru29/gospot/internal/server"
	"github.com/landru29/gospot/internal/server/api"
	"github.com/landru29/gospot/internal/server/client"
	"github.com/sirupsen/logrus"
)

func provideSpotify() (*spotify.Client, error) {
	return spotify.New("https://api.spotify.com/v1")
}

func providePublicAPI(log logrus.FieldLogger, conf *app.Config, toLaunch apiList, catalog music.Cataloger) ([]server.Router, error) {
	out := []server.Router{}

	if toLaunch.Has("api") {
		srv, err := api.New(log, conf, catalog)
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
