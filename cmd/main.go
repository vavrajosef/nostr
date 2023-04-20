package main

import (
	"context"

	"github.com/go-nostr/go-nostr/api/grpc/nostrgrpc"
	"github.com/go-nostr/go-nostr/api/http/nostrhttp"
	"github.com/go-nostr/go-nostr/api/wss/nostrwss"
	"github.com/go-nostr/go-nostr/docs"
	"github.com/go-nostr/go-nostr/web"
	"golang.org/x/sync/errgroup"
)

func main() {
	// NOTE: Create a new context for managing the lifetime of the main function
	ctx := context.Background()

	// NOTE: Initialize a struct containing all the service servers
	serviceCollection := struct {
		docsServer      *docs.Server
		nostrGRPCServer *nostrgrpc.Server
		nostrHTTPServer *nostrhttp.Server
		webServer       *web.Server
		wssServer       *nostrwss.Server
	}{
		docsServer:      buildDocsServer(),
		nostrGRPCServer: buildGRPCServer(),
		nostrHTTPServer: buildHTTPServer(),
		webServer:       buildWebServer(),
		wssServer:       buildWSSServer(),
	}

	// NOTE: Create an error group to manage the concurrent execution of service servers
	group, _ := errgroup.WithContext(ctx)

	// NOTE: Add service server Serve functions to the error group
	group.Go(serviceCollection.docsServer.Serve)
	group.Go(serviceCollection.nostrGRPCServer.Serve)
	group.Go(serviceCollection.nostrHTTPServer.Serve)
	group.Go(serviceCollection.webServer.Serve)
	group.Go(serviceCollection.wssServer.Serve)

	// NOTE: Wait for all service servers to complete or return an error
	if err := group.Wait(); err != nil {
		panic(err)
	}
}
