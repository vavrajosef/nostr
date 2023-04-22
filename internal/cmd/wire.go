//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-nostr/nostr/internal/client"
	"github.com/go-nostr/nostr/internal/docs"
	"github.com/go-nostr/nostr/internal/relay"
	"github.com/google/wire"
)

func buildClientServer() *client.Server {
	wire.Build(
		client.NewServer,
	)
	return &client.Server{}
}

func buildDocsServer() *docs.Server {
	wire.Build(
		docs.NewServer,
	)
	return &docs.Server{}
}

func buildRelayServer() *relay.Server {
	wire.Build(
		relay.NewServer,
	)
	return &relay.Server{}
}
