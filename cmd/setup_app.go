//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/landru29/gospot/internal/app"

	"github.com/landru29/gospot/internal/music"
	"github.com/landru29/gospot/internal/music/spotify"
)

func setupApplication(conf *app.Config, toLaunch apiList) (*app.Controller, func(), error) {
	wire.Build(
		app.New,
		providePublicAPI,
		provideLogger,
		provideSpotify,
		wire.Bind(new(music.Cataloger), new(*spotify.Client)),
	)

	return &app.Controller{}, func() {}, nil
}
