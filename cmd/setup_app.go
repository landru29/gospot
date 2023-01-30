//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/landru29/gospot/internal/app"
)

func setupApplication(conf *app.Config, toLaunch apiList) (*app.Controller, func(), error) {
	wire.Build(
		app.New,
		providePublicAPI,
		provideLogger,
	)

	return &app.Controller{}, func() {}, nil
}
