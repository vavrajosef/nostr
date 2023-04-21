//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-nostr/go-nostr/internal/client"
	"github.com/go-nostr/go-nostr/internal/docs"
	"github.com/google/wire"
)

func buildDocsServer() *docs.Server {
	wire.Build(
		docs.NewServer,
	)
	return &docs.Server{}
}

func buildWebServer() *client.Server {
	wire.Build(
		client.NewServer,
	)
	return &client.Server{}
}
