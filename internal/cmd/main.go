package main

import (
	"context"

	"github.com/go-nostr/nostr"
	"github.com/go-nostr/nostr/internal/client"
	"github.com/go-nostr/nostr/internal/docs"
	"golang.org/x/sync/errgroup"
)

func main() {
	// NOTE: Create a new context for managing the lifetime of the main function
	ctx := context.Background()

	// NOTE: Initialize a struct containing all the service servers
	serviceCollection := struct {
		clientServer *client.Server
		docsServer   *docs.Server
		relay        *nostr.Relay
	}{
		clientServer: buildClientServer(),
		docsServer:   buildDocsServer(),
		relay:        buildRelay(),
	}

	// NOTE: Create an error group to manage the concurrent execution of service servers
	group, _ := errgroup.WithContext(ctx)

	// NOTE: Add service server Serve functions to the error group
	group.Go(serviceCollection.clientServer.ListenAndServe)
	group.Go(serviceCollection.docsServer.ListenAndServe)
	group.Go(serviceCollection.relay.ListenAndServe)

	// NOTE: Wait for all service servers to complete or return an error
	if err := group.Wait(); err != nil {
		panic(err)
	}
}
