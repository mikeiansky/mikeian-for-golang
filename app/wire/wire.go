//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import "github.com/google/wire"

// wireApp init kratos application.
func createApp() *App {
	wire.Build(CreateName, CreateWireDb, CreateWireConfig, NewApp)
	return new(App)
}

// cd this dir
// wire gen
