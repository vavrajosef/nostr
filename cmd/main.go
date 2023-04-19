package main

import (
	"context"

	"github.com/go-nostr/go-nostr/api/grpc/nostrgrpc"
	"github.com/go-nostr/go-nostr/api/http/nostrhttp"
	"github.com/go-nostr/go-nostr/api/wss"
	"github.com/go-nostr/go-nostr/docs"
	"github.com/go-nostr/go-nostr/web"
	"golang.org/x/sync/errgroup"
)

func main() {
	// NOTE: TBD
	ctx := context.Background()

	// NOTE: TBD
	serviceCollection := struct {
		docsServer      *docs.Server
		nostrGRPCServer *nostrgrpc.Server
		nostrHTTPServer *nostrhttp.Server
		webServer       *web.Server
		wssServer       *wss.Server
	}{
		docsServer:      buildDocsServer(),
		nostrGRPCServer: buildGRPCServer(),
		nostrHTTPServer: buildHTTPServer(),
		webServer:       buildWebServer(),
		wssServer:       buildWSSServer(),
	}

	// TODO: improve error handling
	group, _ := errgroup.WithContext(ctx)

	// TODO: add #Serve()
	group.Go(serviceCollection.docsServer.Serve)
	group.Go(serviceCollection.nostrGRPCServer.Serve)
	group.Go(serviceCollection.nostrHTTPServer.Serve)
	group.Go(serviceCollection.webServer.Serve)
	group.Go(serviceCollection.wssServer.Serve)

	// NOTE: TBD
	if err := group.Wait(); err != nil {
		panic(err)
	}
}
