// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-nostr/go-nostr/api/grpc"
	"github.com/go-nostr/go-nostr/api/http"
	"github.com/go-nostr/go-nostr/api/wss"
	"github.com/go-nostr/go-nostr/docs"
	"github.com/go-nostr/go-nostr/web"
)

// Injectors from wire.go:

func buildDocsServer() *docs.Server {
	server := docs.NewServer()
	return server
}

func buildGRPCServer() *grpc.Server {
	server := grpc.NewServer()
	return server
}

func buildHTTPServer() *http.Server {
	server := http.NewServer()
	return server
}

func buildWebServer() *web.Server {
	server := web.NewServer()
	return server
}

func buildWSSServer() *wss.Server {
	server := wss.NewServer()
	return server
}
