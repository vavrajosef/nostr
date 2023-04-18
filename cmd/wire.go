//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-nostr/go-nostr/web"
	"github.com/google/wire"
)

func buildWebServer() *web.Server {
	wire.Build(
		web.NewServer,
	)
	return &web.Server{}
}
