//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/go-nostr/nostr"
	"github.com/google/wire"
)

func buildClientServer() *http.Server {
	wire.Build(wire.NewSet(
		nostr.NewClient,
		wire.Struct(new(http.Server), "Handler"),
	))
	return &http.Server{}
}
