package main

import (
	"context"

	"github.com/go-nostr/go-nostr/web"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx := context.Background()
	sc := struct {
		ws *web.Server
	}{
		ws: buildWebServer(),
	}

	// TODO: improve error handling
	g, _ := errgroup.WithContext(ctx)

	g.Go(sc.ws.Serve)

	if err := g.Wait(); err != nil {
		panic(err)
	}
}
