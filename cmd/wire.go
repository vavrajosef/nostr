//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-nostr/go-nostr/api/grpc/nostrgrpc"
	"github.com/go-nostr/go-nostr/api/http/nostrhttp"
	"github.com/go-nostr/go-nostr/api/wss/nostrwss"
	"github.com/go-nostr/go-nostr/docs"
	"github.com/go-nostr/go-nostr/web"
	"github.com/google/wire"
)

func buildDocsServer() *docs.Server {
	wire.Build(
		docs.NewServer,
	)
	return &docs.Server{}
}

func buildGRPCServer() *nostrgrpc.Server {
	wire.Build(
		nostrgrpc.NewServer,
	)
	return &nostrgrpc.Server{}
}

func buildHTTPServer() *nostrhttp.Server {
	wire.Build(
		nostrhttp.NewServer,
	)
	return &nostrhttp.Server{}
}

func buildWebServer() *web.Server {
	wire.Build(
		web.NewServer,
	)
	return &web.Server{}
}

func buildWSSServer() *nostrwss.Server {
	wire.Build(
		nostrwss.NewServer,
	)
	return &nostrwss.Server{}
}
