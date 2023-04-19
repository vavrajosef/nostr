package main

import (
	"context"

	"github.com/go-nostr/go-nostr/api/grpc"
	"github.com/go-nostr/go-nostr/api/http"
	"github.com/go-nostr/go-nostr/docs"
	"github.com/go-nostr/go-nostr/web"
	"golang.org/x/sync/errgroup"
)

func main() {
	// NOTE: TBD
	ctx := context.Background()

	// NOTE: TBD
	serviceCollection := struct {
		docsServer *docs.Server
		grpcServer *grpc.Server
		httpServer *http.Server
		webServer  *web.Server
	}{
		docsServer: buildDocsServer(),
		grpcServer: buildGRPCServer(),
		httpServer: buildHTTPServer(),
		webServer:  buildWebServer(),
	}

	// TODO: improve error handling
	group, _ := errgroup.WithContext(ctx)

	// TODO: add #Serve()
	group.Go(serviceCollection.docsServer.Serve)
	group.Go(serviceCollection.grpcServer.Serve)
	group.Go(serviceCollection.httpServer.Serve)
	group.Go(serviceCollection.webServer.Serve)

	// NOTE: TBD
	if err := group.Wait(); err != nil {
		panic(err)
	}
}
