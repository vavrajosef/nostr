// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-nostr/nostr/internal/client"
	"github.com/go-nostr/nostr/internal/docs"
	"github.com/go-nostr/nostr/internal/relay"
)

// Injectors from wire.go:

func buildClientServer() *client.Server {
	server := client.NewServer()
	return server
}

func buildDocsServer() *docs.Server {
	server := docs.NewServer()
	return server
}

func buildRelayServer() *relay.Server {
	server := relay.NewServer()
	return server
}
