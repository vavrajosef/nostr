//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-nostr/go-nostr/api/grpc"
	"github.com/go-nostr/go-nostr/api/http"
	"github.com/go-nostr/go-nostr/api/wss"
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

func buildGRPCServer() *grpc.Server {
	wire.Build(
		grpc.NewServer,
	)
	return &grpc.Server{}
}

func buildHTTPServer() *http.Server {
	wire.Build(
		http.NewServer,
	)
	return &http.Server{}
}

func buildWebServer() *web.Server {
	wire.Build(
		web.NewServer,
	)
	return &web.Server{}
}

func buildWSSServer() *wss.Server {
	wire.Build(
		wss.NewServer,
	)
	return &wss.Server{}
}
